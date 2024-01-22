import { observer } from "mobx-react-lite";
import { useStore } from "./stores";
import { Accordion, AccordionDetails, AccordionSummary, Box, Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, FormControl, IconButton, Paper, TextField, Typography } from "@mui/material";
import { Add, Delete, ExpandMore } from "@mui/icons-material";
import { useRef } from "react";

function AppletList() {
    const { SimStore } = useStore();
    const descriptionElementRef = useRef<HTMLElement>(null);

    return (
        <Dialog open={SimStore.openSimConfig} fullWidth maxWidth={"xl"} PaperProps={{ sx: { height: "100%" } }}>
            <DialogTitle alignItems={"flex-start"}>{SimStore.iccid}</DialogTitle>
            <DialogContent dividers sx={{ background: "#f7f9fb" }}>
                <DialogContentText
                    id="scroll-dialog-description"
                    ref={descriptionElementRef}
                    tabIndex={-1}
                >
                    {SimStore.editKeys.map((key, index) => (
                        <Accordion defaultExpanded={!SimStore.editKeys[index].name}>
                            <AccordionSummary
                                expandIcon={<ExpandMore />}
                                aria-controls="panel1-content"
                                id="panel1-header"
                            >
                                設定 {index + 1}
                            </AccordionSummary>
                            <AccordionDetails>
                                <FormControl fullWidth>
                                    <TextField
                                        label={"名前"}
                                        value={SimStore.editKeys[index].name}
                                        onChange={(event) => SimStore.editKeys[index].name = event.target.value}
                                        helperText={!SimStore.editKeys[index].name && "必須"}
                                        error={!SimStore.editKeys[index].name}
                                    />
                                </FormControl>
                                <FormControl fullWidth sx={{ mt: 2 }}>
                                    <TextField
                                        label={"AID"}
                                        value={SimStore.editKeys[index].aid}
                                        onChange={(event) => SimStore.editKeys[index].aid = event.target.value}
                                        helperText={!SimStore.editKeys[index].aid && "必須"}
                                        error={!SimStore.editKeys[index].aid}
                                    />
                                </FormControl>
                                <FormControl fullWidth sx={{ mt: 2 }}>
                                    <TextField
                                        label={"ENC Key"}
                                        value={SimStore.editKeys[index].encKey}
                                        onChange={(event) => SimStore.editKeys[index].encKey = event.target.value}
                                        helperText={!SimStore.editKeys[index].encKey && "必須"}
                                        error={!SimStore.editKeys[index].encKey}
                                    />
                                </FormControl>
                                <FormControl fullWidth sx={{ mt: 2 }}>
                                    <TextField
                                        label={"MAC Key"}
                                        value={SimStore.editKeys[index].macKey}
                                        onChange={(event) => SimStore.editKeys[index].macKey = event.target.value}
                                        helperText={!SimStore.editKeys[index].macKey && "必須"}
                                        error={!SimStore.editKeys[index].macKey}
                                    />
                                </FormControl>
                                <FormControl fullWidth sx={{ mt: 2 }}>
                                    <TextField
                                        label={"KEK Key"}
                                        value={SimStore.editKeys[index].kekKey}
                                        onChange={(event) => SimStore.editKeys[index].kekKey = event.target.value}
                                        helperText={!SimStore.editKeys[index].kekKey && "必須"}
                                        error={!SimStore.editKeys[index].kekKey}
                                    />
                                </FormControl>

                                <Box display={"flex"} justifyContent={"flex-end"} mt={2}>
                                    <Button
                                        startIcon={<Delete />}
                                        color="error"
                                        onClick={() => SimStore.editKeys.splice(index, 1)}
                                    >
                                        削除
                                    </Button>
                                </Box>
                            </AccordionDetails>
                        </Accordion>
                    ))}
                    <Box>
                        <IconButton onClick={() => SimStore.editKeys.push({ name: "", aid: "", encKey: "", macKey: "", kekKey: "" })}>
                            <Add />
                        </IconButton>
                    </Box>
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={() => SimStore.setOpenSimConfig(false)} color="inherit">キャンセル</Button>
                <Button
                    disabled={SimStore.editKeys.some(key => Object.values(key).includes(""))}
                    onClick={() => SimStore.saveSimConfig()}
                >
                    保存
                </Button>
            </DialogActions>
        </Dialog>
    );
}

export default observer(AppletList);
