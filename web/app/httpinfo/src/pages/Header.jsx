import styles from "../styles/Header.module.css";

function Header() {
    return (
    <div className={`${styles.header} ${styles.hill}`}>
        <a href="/">Home</a>
    </div>
    )
}

export default Header