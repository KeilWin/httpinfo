import styles from "~/styles/common/pit.module.css";

export default function PitContainer(props) {
    return (
        <div class={styles.pit}>
            {props.children}
        </div>
    )
}