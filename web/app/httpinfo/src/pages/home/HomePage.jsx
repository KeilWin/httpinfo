import IpInfoContainer from "./IpInfoContainer";
import RequestContainer from "./RequestContainer";

import styles from "../../styles/pages/home/HomePage.module.css";

export default function HomePage() {
    return (
    <div className={`${styles.homePage} ${styles.hill}`}>
        <IpInfoContainer />
        <RequestContainer />
    </div>
    )
}