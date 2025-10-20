import { Show, Switch, Match, createResource } from "solid-js";

import styles from "../../styles/pages/home/IpInfoCard.module.css";

const fetchIpInfo = async (ipAddress) => {
    const response = await fetch(`${window.location.protocol}//${window.location.host}/api/ip/${ipAddress}`);
    return response.json();
}

export default function IpInfoCard() {
    const [ipInfo] = createResource("8.8.8.8", fetchIpInfo);
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