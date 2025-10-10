import PitContainer from "../../components/PitContainer";

import { createResource, createSignal } from "solid-js";
import { Show, Switch, Match } from "solid-js";

const fetchIpInfo = async (ip: string) => {
    console.log("ip: ", ip);
    ip = "8.8.8.8";
    const response = await fetch(`https://api.iplocation.net/?ip=${ip}`);
    console.log("obj: %O", response);
    if (!response.ok) {
        throw new Error(`Fetch error: ${response.statusText}`);
    }
    return response.json().catch((e) => {console.log(e)});
}

type IpinfoCardProps = {

}

export default function IpInfoCard(props: IpinfoCardProps) {
    // const [data, { mutate, refetch }] = createResource("8.8.8.8", fetchIpInfo);
    // return (
    //     <PitContainer>
    //         <Switch>
    //             <Match when={data.error}>
    //                 <p>Error: data.error</p>
    //             </Match>
    //             <Match when={data()}>
    //                 <pre>{data()}</pre>
    //             </Match>
    //         </Switch>
    //     </PitContainer>
    // )
    return (
        <p>Some text</p>
    )
}