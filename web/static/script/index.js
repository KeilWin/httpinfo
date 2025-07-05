
function GetDefaultIp() {
    return "8.8.8.8";
}

function ExtractIp(ipElementId) {
    let ipElement = document.getElementById(ipElementId);
    if (!ipElement)
        throw new Error("No element with id.", ipElementId);
    let ip = ipElement.textContent.split(":")[0];
    return ip.indexOf("[") !== -1 ? GetDefaultIp() : ip;
}

function CreateRequestUrl(ip) {
    return `https://api.iplocation.net/?ip=${ip}`;
}

function SetIpInfo(ipInfo) {
    const ipInfoId = "ipInfo";
    let ipElement = document.getElementById(ipInfoId);
    if (!ipElement)
        throw new Error("No element with id.", ipInfoId);
    ipElement.textContent = JSON.stringify(ipInfo, undefined, 2);
}

async function GetIpInfoAsync(ipElementId) {
    const ip = ExtractIp(ipElementId);
    const url = CreateRequestUrl(ip);
    fetch(url).then(function(response) {
        return response.json();
      }).then(function(data) {
        SetIpInfo(data);
        return data;
      }).catch(function(err) {
        console.log('Fetch Error :-S', err);
      });
}

const ipDataElementId = "remoteAddr";
const ipData = GetIpInfoAsync(ipDataElementId);