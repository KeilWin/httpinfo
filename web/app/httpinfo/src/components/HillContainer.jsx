import styles from "../styles/components/HillContainer.module.css";

function HillContainer(props) {
    return (
    <div className={styles.hill}>
        {props.children}
    </div>
    )
}

export default HillContainer