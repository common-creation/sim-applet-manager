import { observer } from "mobx-react-lite";
import { useStore } from "./stores";
import { Accordion, AccordionDetails, AccordionSummary, Box, Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Fab, FormControl, IconButton, Paper, Tab, Tabs, TextField, Typography } from "@mui/material";
import { Add, Delete, ExpandMore, Refresh } from "@mui/icons-material";
import { useEffect, useState } from "react";
import { SelectCapFilePath } from "../wailsjs/go/main/App";

function AppletList() {
    const { SimStore } = useStore();
    const [tabIndex, setTabIndex] = useState(0);
    const [openInstallDialog, setOpenInstallDialog] = useState(false);
    const [capPath, setCapPath] = useState("");
    const [params, setParams] = useState("");

    useEffect(() => {
        setTabIndex(0);
    }, [SimStore.keys]);
    useEffect(() => {
        if (SimStore.keys.length > tabIndex) {
            SimStore.listApplets(SimStore.keys[tabIndex]);
        }
    }, [tabIndex]);
    useEffect(() => {
        if (openInstallDialog) {
            setCapPath("");
            setParams("");
        }
    }, [openInstallDialog]);

    return (
        <Box mt={0} display={"flex"} flexDirection={"column"} flex={1} sx={{ background: "#f7f9fb" }} overflow={"hidden"}>
            {SimStore.keys.length > 1 && (
                <Paper>
                    <Tabs variant="fullWidth" value={tabIndex} onChange={(e, v) => setTabIndex(v)} sx={{ bgcolor: 'background.paper' }}>
                        {
                            SimStore.keys.map((key, index) => (
                                <Tab label={key.name} key={key.name + index} />
                            ))
                        }
                    </Tabs>
                </Paper>
            )
            }

            {
                SimStore.applets.length === 0 ? (
                    <Box flex={1} display={"flex"} alignItems={"center"} justifyContent={"center"}>
                        <Typography variant="h4" color={"text.secondary"}>アプレットがインストールされていません</Typography>
                    </Box>
                ) : (
                    <Box overflow={"auto"} p={2}>
                        {SimStore.applets.map((pkg) => (
                            <Accordion key={pkg.package.hex}>
                                <AccordionSummary
                                    expandIcon={<ExpandMore />}
                                    aria-controls="panel1-content"
                                    id="panel1-header"
                                >
                                    <Box display={"flex"} alignItems={"center"}>
                                        <Typography sx={{ fontFamily: "monospace" }}>
                                            {pkg.package.hex}
                                        </Typography>
                                        <Typography variant="caption" color={"text.secondary"} sx={{ ml: 1, fontFamily: "monospace" }}>
                                            {pkg.package.fingerPrint}
                                        </Typography>
                                    </Box>
                                </AccordionSummary>
                                <AccordionDetails>
                                    <Box display={"flex"} flexDirection={"column"} alignItems={"flex-start"}>
                                        <Typography variant="caption">
                                            パッケージ内のアプレット
                                        </Typography>
                                        {pkg.applets.map((applet) => (
                                            <Box display={"flex"} alignItems={"center"} key={applet.hex}>
                                                <Typography sx={{ fontFamily: "monospace" }}>
                                                    {applet.hex}
                                                </Typography>
                                                <Typography variant="caption" color={"text.secondary"} sx={{ ml: 1, fontFamily: "monospace" }}>
                                                    {applet.fingerPrint}
                                                </Typography>
                                            </Box>
                                        ))}
                                    </Box>
                                    <Box display={"flex"} justifyContent={"flex-end"} mt={2}>
                                        <Button
                                            startIcon={<Delete />}
                                            color="error"
                                            onClick={() => SimStore.uninstallApplet(SimStore.keys[tabIndex], pkg.package.hex)}
                                        >
                                            削除
                                        </Button>
                                    </Box>
                                </AccordionDetails>
                            </Accordion>
                        ))}
                        <Box mb={"78px"} />
                    </Box>
                )
            }
            <Box position={"fixed"} bottom={16} right={16}>
                <Fab disabled={!SimStore.iccid || SimStore.keys.length === 0} size="small" sx={{ mr: 2 }} onClick={() => SimStore.listApplets(SimStore.keys[tabIndex])}><Refresh /></Fab>
                <Fab color="primary" disabled={!SimStore.iccid || SimStore.keys.length === 0} onClick={() => setOpenInstallDialog(true)}><Add /></Fab>
            </Box>
            <Dialog open={openInstallDialog} fullWidth maxWidth={"xl"} PaperProps={{ sx: { height: "100%" } }}>
                <DialogTitle alignItems={"flex-start"}>アプレット インストール</DialogTitle>
                <DialogContent dividers sx={{ background: "#f7f9fb" }}>
                    <FormControl fullWidth>
                        <Box display={"flex"} alignItems={"start"}>
                            <TextField
                                label={"CAP ファイル パス"}
                                value={capPath}
                                helperText={!capPath && "必須"}
                                error={!capPath}
                                disabled
                                sx={{ flex: 1 }}
                            />
                            <Button
                                sx={{ ml: 2, mt: 1 }}
                                onClick={async () => {
                                    const filePath = await SelectCapFilePath();
                                    if (!filePath) {
                                        return;
                                    }
                                    setCapPath(filePath);
                                }}
                            >
                                選択
                            </Button>
                        </Box>
                    </FormControl>
                    <FormControl fullWidth sx={{ mt: 2 }}>
                        <TextField
                            label={"C9 パラメータ"}
                            value={params}
                            onChange={(event) => setParams(event.target.value)}
                            helperText={!params && "必須"}
                            error={!params}
                        />
                    </FormControl>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => setOpenInstallDialog(false)} color="inherit">キャンセル</Button>
                    <Button
                        disabled={!capPath || !params}
                        onClick={async () => {
                            const ok = await SimStore.installApplet(SimStore.keys[tabIndex], capPath, params);
                            if (ok) {
                                setOpenInstallDialog(false);
                            }
                        }}
                    >
                        インストール
                    </Button>
                </DialogActions>
            </Dialog>
        </Box >
    );
}

export default observer(AppletList);
