import { observer } from "mobx-react-lite";
import { Box, CircularProgress, Typography } from "@mui/material";

function Loading() {
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
            sx={{ backgroundColor: "rgba(255, 255, 255, 0.65)", backdropFilter: "blur(2px)" }}
            zIndex={10000}
        >
            <CircularProgress />
            {/* <Typography variant="caption" color={"text.secondary"} sx={{ mt: 2 }}>カードリーダーの排他制御で時間がかかる場合があります</Typography> */}
        </Box>

    );
}

export default Loading;
