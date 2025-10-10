import { For, Show, createSignal, useContext } from "solid-js";
import { getRequestEvent, render, RequestEvent } from "solid-js/web";

import PitContainer from "../../components/PitContainer";

import { serverData } from "~/store/ServerStore";

import styles from "~/styles/components/RequestInfoCard.module.css";

type RequestInfoCardProps = {

}

function CapitalizeFirstSymbol(s: string): string {
    if (!s || s?.length === 0)
        return s;
    return `${s[0].toUpperCase()}${s.slice(1)}`
}

export default function RequestInfoCard(props: RequestInfoCardProps) {
    return (
        <PitContainer>
            <p class={styles.headersTitle}><strong>Headers</strong>:</p>
            <ul>
                <For each={serverData.headers}>
                    {([key, value]) => (
                    <li class={styles.headersItem}>
                        <strong class={styles.headersRowKey}>{CapitalizeFirstSymbol(key)}: </strong>
                        <span class={styles.headersRowValue}>{value}</span>
                    </li>
                    )}
                </For>
            </ul>
        </PitContainer>
    )
}