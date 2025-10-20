import { Show, Switch, Match, createResource } from "solid-js";

import RequestInfo from "../../utils/init_data";

import styles from "../../styles/pages/home/IpInfoCard.module.css";

const fetchIpInfo = async (ipAddress) => {
    const response = await fetch(`${window.location.protocol}//${window.location.host}/api/ip/${ipAddress}`);
    return response.json();
}

const getDefaultIp = () => {
    return "8.8.8.8";
}

const getIpAddress = () => {
    const requestInfo = RequestInfo();
    if (!requestInfo || !requestInfo.ipAddress)
        return getDefaultIp();
    return requestInfo.ipAddress;
}

export default function IpInfoCard() {
    const [ipInfo] = createResource(getIpAddress, fetchIpInfo);
    return (
    <div className={`${styles.ipInfoCard} ${styles.pit}`}>
        <Show when={ipInfo.loading}>
            <p>Loading...</p>
        </Show>
        <Switch>
            <Match when={ipInfo.error}>
                <p>Something go wrong...</p>
            </Match>
            <Match when={ipInfo()}>
                <p><strong>Counry name</strong>: {ipInfo().country_name}</p>
                <p><strong>Counry code</strong>: {ipInfo().country_code2}</p>
                <p><strong>Internet service provider</strong>: {ipInfo().isp}</p>
            </Match>
        </Switch>
    </div>
    )
}