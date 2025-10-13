import RequestRow from "./RequestRow";
import RequestHeadersList from "./RequestHeadersList";
import RequestInfo from "../../utils/init_data";

import styles from "../../styles/pages/home/RequestContainer.module.css";

import { Show } from "solid-js";

export default function RequestInfoContainer() {
    const info = RequestInfo();
    return (
    <div className={`${styles.requestContainer} ${styles.pit}`}>
        <Show when={info} fallback={<p>No data</p>}>
            <RequestRow key="Address" value={info["ipAddress"]}/>
            <RequestRow key="Port" value={info["port"]}/>
            <RequestRow key="Protocol" value={info["protocol"]}/>
            <RequestRow key="Method" value={info["method"]}/>
            <RequestRow key="Host" value={info["host"]}/>
            <RequestRow key="Url" value={info["url"]}/>
            <RequestHeadersList name="Headers" headers={info["headers"]}/>
            <RequestRow key="ContentLength" value={info["contentLength"]}/>
            <Show when={info["body"]}>
                <pre>{info["body"]}</pre>
            </Show>
        </Show>
    </div>
    )
}