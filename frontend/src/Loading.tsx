import { observer } from "mobx-react-lite";
import { Box, CircularProgress, Paper, Typography } from "@mui/material";
import { useStore } from "./stores";

function Loading() {
    const { LogStore } = useStore();

    if (LogStore.logs.length === 0) {
        return (
            <Box
                id="Loading"
                position={"fixed"}
                left={0}
                right={0}
                top={0}
                bottom={0}
                display={"flex"}
                flexDirection={"column"}
                justifyContent={"center"}
                alignItems={"center"}
                sx={{ backgroundColor: "rgba(0, 0, 0, 0.5)" }}
                zIndex={10000}
            >
                <CircularProgress />
            </Box>
        );
    }

    return (
        <Box
            id="Loading"
            position={"fixed"}
            left={0}
            right={0}
            top={0}
            bottom={0}
            display={"flex"}
            flexDirection={"column"}
            sx={{ backgroundColor: "rgba(0, 0, 0, 0.5)" }}
            zIndex={10000}
        >
            <Box position={"absolute"} bottom={32} left={32} right={32} height={"calc(50% - 32px)"}>
                <Paper sx={{ background: "rgba(20, 22, 32, 0.66)", p: "8px", width: "calc(100% - 16px)", height: "calc(100% - 16px)", overflow: "auto", backdropFilter: "blur(8px)" }}>
                    {LogStore.logs.map((line, index) => <Typography key={index} id={`log_line_${index}`} variant="body2" color={"white"} sx={{ width: "100%", wordBreak: "break-all" }} align="left">{line}</Typography>)}
                    <div id={"eof"} />
                </Paper>
            </Box>
        </Box>

    );
}

export default observer(Loading);
