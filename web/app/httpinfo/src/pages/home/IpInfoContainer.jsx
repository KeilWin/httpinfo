import IpInfoCard from "./IpInfoCard";

import styles from "../../styles/pages/home/IpInfoContainer.module.css";

export default function IpInfoContainer() {
    return (
    <div className={`${styles.ipInfoContainer}`}>
        <IpInfoCard />
    </div>
    )
}