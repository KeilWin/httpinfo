import styles from "../../styles/pages/home/IpInfoCard.module.css";

export default function IpInfoCard() {
    return (
    <div className={`${styles.ipInfoCard} ${styles.pit}`}>
        <p><strong>Counry name</strong>: Russia</p>
        <p><strong>Counry code</strong>: RU</p>
        <p><strong>Internet service provider</strong>: St.Petersburg Telephone Network 1</p>
    </div>
    )
}