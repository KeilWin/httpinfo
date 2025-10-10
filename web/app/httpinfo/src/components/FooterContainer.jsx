import styles from "~/styles/footer.module.css";

export default function FooterContainer(props) {
    return (
        <div class={`${styles.hill} ${styles.footer}`}>
            {props.children}
        </div>
    )
}