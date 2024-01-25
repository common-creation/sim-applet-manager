import "./App.css";
import useDidMount from "beautiful-react-hooks/useDidMount";
import { observer } from "mobx-react-lite";
import { useStore } from "./stores";
import TopSelector from "./TopSelector";
import AppletList from "./AppletList";
import SimConfig from "./SimConfig";
import Loading from "./Loading";
import { Box } from "@mui/material";

function App() {
    const { SimStore } = useStore();

    useDidMount(async () => {
        try {
            document.body.classList.add("lockloading");
            await Promise.all([SimStore.fetchGpPath(), SimStore.listCardReader()])
        } finally {
            document.body.classList.remove("lockloading");
        }
    });

    return (
        <Box id="App" display={"flex"} flexDirection={"column"}>
            <Loading />
            <TopSelector />
            <AppletList />
            <SimConfig />
        </Box>
    );
}

export default observer(App);
