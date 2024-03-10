import "./App.css";
import useDidMount from "beautiful-react-hooks/useDidMount";
import { observer } from "mobx-react-lite";
import { useStore } from "./stores";
import TopSelector from "./TopSelector";
import AppletList from "./AppletList";
import SimConfig from "./SimConfig";
import Loading from "./Loading";
import { Box, Link, Typography } from "@mui/material";

function App() {
    const { SimStore, I18nStore: i18n, VersionStore } = useStore();

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
            {VersionStore.version && (
                <Box sx={{ position: "fixed", top: 0, right: 0, p: 1, zIndex: 1, cursor: "pointer" }}>
                    <Typography variant="caption" color="primary" onClick={() => VersionStore.openReleasePage()}>
                        {i18n.t("updateAvailable", VersionStore.version)}
                    </Typography>
                </Box>
            )}
            <Loading />
            <TopSelector />
            <AppletList />
            <SimConfig />
        </Box>
    );
}

export default observer(App);
