import styles from "../../styles/pages/home/RequestRow.module.css";

export default function RequestRow(props) {
    return (
    <p className={styles.requestRow}>
        <strong>{props.key}</strong>: <span>{props.value}</span>
    </p>
    )
}