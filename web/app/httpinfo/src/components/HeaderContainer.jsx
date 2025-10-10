import styles from "../styles/header.module.css";

export default function HeaderContainer(props) {
    return (
        <div class={`${styles.hill} ${styles.header}`}>
            {props.children}
        </div>
    )
}