import styles from "../styles/common/hill.module.css";

export default function HillContainer(props) {
    return (
        <div class={styles.hill}>
            {props.children}
        </div>
    )
}