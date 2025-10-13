import styles from "../../styles/pages/home/RequestRow.module.css";

export default function RequestRow() {
    return (
    <p className={styles.requestRow}>
        <strong>Param</strong>: <span>some value</span>
    </p>
    )
}