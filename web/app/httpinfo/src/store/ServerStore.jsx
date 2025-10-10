import { createStore } from "solid-js/store";

const [serverData, setServerData] = createStore({ headers: [[]]});

export {serverData, setServerData};