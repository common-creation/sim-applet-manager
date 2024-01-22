import { observer } from "mobx-react-lite";
import { useStore } from "./stores";
import { Box, Fab, FormControl, Grid, IconButton, MenuItem, Paper, Select, Tooltip } from "@mui/material";
import { Refresh, Check, Warning, Error, Settings } from "@mui/icons-material";

function TopSelector() {
    const { SimStore } = useStore();

    return (
        <Box sx={{ position: "relative" }}>
            <Paper id="TopSelector" square >
                <Grid container spacing={1} alignItems="center">
                    <Grid item xs={3}>
                        <Box display={"flex"} alignItems="center">
                            {SimStore.gpPath ? (
                                <Check color="success" />
                            ) : (
                                <Error color="error" />
                            )}
                            <Box ml={1}>
                                GlobalPlatformPro パス
                            </Box>
                        </Box>
                    </Grid>
                    <Grid item xs={9}>
                        <Box display={"flex"} alignItems={"center"}>
                            <Box>
                                {SimStore.gpPath || "-"}
                            </Box>
                            <Box ml={1}>
                                <IconButton onClick={() => SimStore.fetchGpPath()}>
                                    <Refresh />
                                </IconButton>
                            </Box>
                        </Box>
                    </Grid>
                    <Grid item xs={3}>
                        <Box display={"flex"} alignItems="center">
                            {SimStore.cardReaders.length > SimStore.selectedCardReaderIndex && SimStore.cardReaders[SimStore.selectedCardReaderIndex] ? (
                                <Check color="success" />
                            ) : (
                                <Error color="error" />
                            )}
                            <Box ml={1}>
                                カードリーダー
                            </Box>
                        </Box>
                    </Grid>
                    <Grid item xs={9}>
                        <Box display={"flex"} alignItems="center" textAlign={"left"}>
                            <FormControl fullWidth>
                                <Select
                                    id="card-reader-select"
                                    value={SimStore.selectedCardReaderIndex}
                                    onChange={(ev) => SimStore.setSelected(ev.target.value as number)}
                                    sx={{ background: "white" }}
                                >
                                    {SimStore.cardReaders.map((cardReader, index) => (
                                        <MenuItem key={`card-reader-${index}`} value={index}>{cardReader}</MenuItem>
                                    ))}
                                </Select>
                            </FormControl>
                            <Box ml={1}>
                                <IconButton onClick={() => SimStore.listCardReader()}>
                                    <Refresh />
                                </IconButton>
                            </Box>
                        </Box>
                    </Grid>
                    <Grid item xs={3}>
                        <Box display={"flex"} alignItems="center">
                            {SimStore.iccid ? (
                                <>
                                    {SimStore.keys.length > 0 ? (
                                        <Check color="success" />
                                    ) : (
                                        <Tooltip title="AID, 各種キーが未設定です。歯車アイコンから設定してください">
                                            <Warning color="warning" />
                                        </Tooltip>
                                    )}
                                </>
                            ) : (
                                <Error color="error" />
                            )}
                            <Box ml={1}>
                                ICCID
                            </Box>
                        </Box>
                    </Grid>
                    <Grid item xs={9}>
                        <Box display={"flex"} alignItems="center">
                            <Box display={"flex"} alignItems={"center"}>
                                <Box>
                                    {SimStore.iccid || "-"}
                                </Box>
                                {SimStore.iccid && (
                                    <Box ml={1}>
                                        <IconButton onClick={() => SimStore.setOpenSimConfig(true)}>
                                            <Settings />
                                        </IconButton>

                                    </Box>
                                )}
                                <Box ml={1}>
                                    <IconButton onClick={() => SimStore.fetchSimInfo(SimStore.cardReaders[SimStore.selectedCardReaderIndex])}>
                                        <Refresh />
                                    </IconButton>
                                </Box>
                            </Box>
                            <Fab sx={{ visibility: "hidden" }} />
                        </Box>
                    </Grid>
                </Grid>
            </Paper>
        </Box>
    );
}

export default observer(TopSelector);
