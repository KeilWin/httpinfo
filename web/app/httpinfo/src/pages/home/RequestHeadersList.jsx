import RequestHeadersItem from "./RequestHeadersItem";

import styles from "../../styles/pages/home/RequestHeadersList.module.css";

import { For } from "solid-js";

export default function RequestHeadersList(props) {
    return (
    <>
    <p><strong className={`${styles.headersKey}`}>{props.name}</strong>:</p>
    <ul className={`${styles.requestHeadersList}`}>
        <For each={Object.entries(props.headers)}>
            {([key, value]) => <RequestHeadersItem key={key} value={value}/>}
        </For>
    </ul>
    </>
    )
}