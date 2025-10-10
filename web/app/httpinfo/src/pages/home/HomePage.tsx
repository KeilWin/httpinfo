import HomeContent from "./HomeContent";

import styles from "~/styles/pages/HomePageLayout.module.css";

export default function HomePage() {
    return (
        <div class={`${styles.layout} ${styles.hill}`}>
            <HomeContent />
        </div>
    )
}