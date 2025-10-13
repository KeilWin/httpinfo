import RequestRow from "./RequestRow";
import RequestHeadersList from "./RequestHeadersList";

import styles from "../../styles/pages/home/RequestContainer.module.css";

export default function RequestInfoContainer() {
    return (
    <div className={`${styles.requestContainer} ${styles.pit}`}>
        <RequestRow />
        <RequestRow />
        <RequestRow />
        <RequestRow />
        <RequestHeadersList />
    </div>
    )
}