import styles from "../../styles/pages/home/RequestHeadersItem.module.css";

export default function RequestHeadersItem(props) {
    return (
    <li className={`${styles.requestHeadersItem}`}>
        <strong className={`${styles.headersKey}`}>{props.key}</strong>: <span className={`${styles.headersValue}`}>{props.value}</span>
    </li>
    )
}