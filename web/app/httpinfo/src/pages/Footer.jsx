import RequestInfo from "../utils/init_data";

import styles from "../styles/Footer.module.css";

const Footer = () => {
    const request = RequestInfo();
    return (
    <div className={`${styles.footer} ${styles.hill}`}>
        <p>Some day...</p>
    </div>
    )
}

export default Footer