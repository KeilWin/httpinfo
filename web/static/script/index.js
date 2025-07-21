
const GOOGLE_DNS_8 = "8.8.8.8";

function GetDefaultIp() {
    return GOOGLE_DNS_8;
}

function ExtractIp(ipElementId) {
    let ipElement = document.getElementById(ipElementId);
    if (!ipElement)
        throw new Error("No element with id.", ipElementId);
    let ip = ipElement.textContent;
    return ip.indexOf("[") !== -1 ? GetDefaultIp() : ip;
}

function CreateRequestUrl(ip) {
    return `https://api.iplocation.net/?ip=${ip}`;
}

function SetIpInfo(ipInfo) {
    const countryNameId = "countryName";
    let countryNameElement = document.getElementById(countryNameId);
    if (!countryNameElement)
        throw new Error("No element with id.", countryNameId);
    
    const countryCodeId = "countryCode";
    let countryCodeElement = document.getElementById(countryCodeId);
    if (!countryCodeElement)
        throw new Error("No element with id.", countryCodeId);
    
    const ispId = "internetServiceProvider";
    let ispElement = document.getElementById(ispId);
    if (!ispElement)
        throw new Error("No element with id.", ispId);

    countryNameElement.textContent = ipInfo["country_name"];
    countryCodeElement.textContent = ipInfo["country_code2"];
    ispElement.textContent = ipInfo["isp"];
}

async function GetIpInfoAsync(ipElementId) {
    const ip = ExtractIp(ipElementId);
    const url = CreateRequestUrl(ip);
    fetch(url).then(function(response) {
        return response.json();
      }).then(function(data) {
        if (data["response_code"] !== "200")
            throw new Error("Bad response code from iplocation.", data["response_code"]);
        SetIpInfo(data);
        return data;
      }).catch(function(err) {
        console.log('Fetch Error :-S', err);
      });
}

const ipDataElementId = "remoteAddr";
const ipData = GetIpInfoAsync(ipDataElementId);