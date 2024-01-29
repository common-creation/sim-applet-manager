package main

import (
	"context"
	"fmt"
	"time"

	"github.com/common-creation/sim-applet-manager/util/apppath"
	"github.com/common-creation/sim-applet-manager/util/cardreader"
	"github.com/common-creation/sim-applet-manager/util/db"
	"github.com/common-creation/sim-applet-manager/util/gp"

	wailsRutime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type (
	// App struct
	App struct {
		ctx context.Context
	}
	SimInfo struct {
		ICCID  string `json:"iccid"`
		Config db.Sim `json:"config"`
	}
)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	apppath.MustAppDirPath(ctx)
}

func (a *App) GetGpPath() string {
	return gp.Path(a.ctx)
}

func (a *App) ListCardReader() []string {
	return cardreader.ListCardReader()
}

func (a *App) FetchSimInfo(cardReader string) SimInfo {
	simInfo := SimInfo{}
	for i := 5; i < 10; i += 2 {
		cardreader.ResetOnMacOS(cardReader)
		result := gp.GetICCID(a.ctx, cardReader)
		cardreader.ResetOnMacOS(cardReader)
		if result == "" {
			return SimInfo{}
		}
		if len(result) <= 20 {
			simInfo.ICCID = result
			break
		}

		time.Sleep(time.Duration(i) * time.Second)
	}
	println("GetSimConfig", simInfo.ICCID)
	config, err := db.GetSimConfig(a.ctx, simInfo.ICCID)
	if err != nil {
		fmt.Printf("GetSimConfig err: %+v\n", err)
		return simInfo
	}
	simInfo.Config = *config

	return simInfo
}

func (a *App) SaveSimConfig(iccid string, config db.Sim) bool {
	println("SaveSimConfig", iccid)
	err := db.PutSimConfig(a.ctx, iccid, &config)
	fmt.Printf("SaveSimConfig err: %+v\n", err)
	return err == nil
}

func (a *App) ShowErrorDialog(title string, message string) {
	wailsRutime.MessageDialog(a.ctx, wailsRutime.MessageDialogOptions{
		Type:          wailsRutime.ErrorDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"OK"},
		DefaultButton: "OK",
	})
}

func (a *App) ShowConfirmDialog(title string, message string) string {
	result, _ := wailsRutime.MessageDialog(a.ctx, wailsRutime.MessageDialogOptions{
		Type:          wailsRutime.QuestionDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"OK", "キャンセル"},
		DefaultButton: "OK",
	})
	return result
}

func (a *App) ListApplets(cardReader string, key db.Key) []gp.ListResult {
	return gp.ListApplets(a.ctx, cardReader, key)
}

func (a *App) UninstallApplet(cardReader string, key db.Key, aid string) gp.Result {
	return gp.UninstallApplet(a.ctx, cardReader, key, aid)
}

func (a *App) SelectCapFilePath() string {
	result, _ := wailsRutime.OpenFileDialog(a.ctx, wailsRutime.OpenDialogOptions{
		Title:                "CAP ファイルを選択",
		Filters:              []wailsRutime.FileFilter{{DisplayName: "CAP", Pattern: "*.cap"}},
		CanCreateDirectories: false,
	})
	return result
}

func (a *App) InstallApplet(cardReader string, key db.Key, capPath string, params string) gp.Result {
	return gp.InstallApplet(a.ctx, cardReader, key, capPath, params)
}
