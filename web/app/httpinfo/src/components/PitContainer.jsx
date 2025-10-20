import styles from "../styles/components/PitContainer.module.css";

function PitContainer(props) {
    return (
    <div className={styles.pit}>
        {props.children}
    </div>
    )
}

export default PitContainer