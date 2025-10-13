import styles from "../../styles/pages/home/RequestHeadersItem.module.css";

export default function RequestHeadersItem() {
    return (
    <li className={`${styles.requestHeadersItem}`}>
        <strong className="headers-row-key">Key</strong>: <span className="headers-row-value">Value</span>
    </li>
    )
}