import RequestHeadersItem from "./RequestHeadersItem";

import styles from "../../styles/pages/home/RequestHeadersList.module.css";

export default function RequestHeadersList() {
    return (
    <>
    <p><strong className={`${styles.headersKey}`}>Headers</strong>:</p>
    <ul className={`${styles.requestHeadersList}`}>
        <RequestHeadersItem />
        <RequestHeadersItem />
        <RequestHeadersItem />
    </ul>
    </>
    )
}