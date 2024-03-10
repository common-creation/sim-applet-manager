package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/common-creation/sim-applet-manager/util/apppath"
	"github.com/common-creation/sim-applet-manager/util/cardreader"
	"github.com/common-creation/sim-applet-manager/util/db"
	"github.com/common-creation/sim-applet-manager/util/gp"
	"github.com/common-creation/sim-applet-manager/util/i18n"
	"github.com/common-creation/sim-applet-manager/util/log"
	"github.com/tidwall/gjson"

	"golang.org/x/mod/semver"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type (
	// App struct
	App struct {
		ctx     context.Context
		i18n    *i18n.I18n
		version string
	}
	SimInfo struct {
		ICCID  string `json:"iccid"`
		Config db.Sim `json:"config"`
	}
)

// NewApp creates a new App application struct
func NewApp(version string) *App {
	return &App{
		i18n:    i18n.NewI18n(),
		version: version,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	apppath.MustAppDirPath(ctx, a.i18n)
	log.Rotate(ctx, a.i18n)
}

func (a *App) GetGpPath() string {
	return gp.Path(a.ctx, a.i18n)
}

func (a *App) ListCardReader() []string {
	return cardreader.ListCardReader()
}

func (a *App) FetchSimInfo(cardReader string) SimInfo {
	simInfo := SimInfo{}
	for i := 5; i < 10; i += 2 {
		cardreader.ResetOnMacOS(cardReader)
		result := gp.GetICCID(a.ctx, a.i18n, cardReader)
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
	config, err := db.GetSimConfig(a.ctx, a.i18n, simInfo.ICCID)
	if err != nil {
		fmt.Printf("GetSimConfig err: %+v\n", err)
		return simInfo
	}
	simInfo.Config = *config

	return simInfo
}

func (a *App) SaveSimConfig(iccid string, config db.Sim) bool {
	println("SaveSimConfig", iccid)
	err := db.PutSimConfig(a.ctx, a.i18n, iccid, &config)
	fmt.Printf("SaveSimConfig err: %+v\n", err)
	return err == nil
}

func (a *App) ShowErrorDialog(title string, message string) {
	wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
		Type:          wailsRuntime.ErrorDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"OK"},
		DefaultButton: "OK",
	})
}

func (a *App) ShowConfirmDialog(title string, message string) string {
	result, _ := wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
		Type:          wailsRuntime.QuestionDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"OK", a.i18n.T("cancel")},
		DefaultButton: "OK",
	})
	return result
}

func (a *App) ListApplets(cardReader string, key db.Key) []gp.ListResult {
	return gp.ListApplets(a.ctx, a.i18n, cardReader, key)
}

func (a *App) UninstallApplet(cardReader string, key db.Key, aid string) gp.Result {
	return gp.UninstallApplet(a.ctx, a.i18n, cardReader, key, aid)
}

func (a *App) SelectCapFilePath() string {
	result, _ := wailsRuntime.OpenFileDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		Title:                a.i18n.T("selectCapFile"),
		Filters:              []wailsRuntime.FileFilter{{DisplayName: "CAP", Pattern: "*.cap"}},
		CanCreateDirectories: false,
	})
	return result
}

func (a *App) InstallApplet(cardReader string, key db.Key, capPath string, params string) gp.Result {
	return gp.InstallApplet(a.ctx, a.i18n, cardReader, key, capPath, params)
}

func formatVersion(v string) string {
	if !strings.HasPrefix(v, "v") {
		return semver.Canonical("v" + v)
	}
	return semver.Canonical(v)
}

func (a *App) CheckUpdate() string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Get("https://api.github.com/repos/common-creation/sim-applet-manager/releases/latest")
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	m := map[string]interface{}{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return ""
	}

	currentVersion := a.version
	remoteVersion := gjson.Get(string(b), "tag_name").String()
	println("currentVersion", formatVersion(currentVersion), ",", "v:", formatVersion(remoteVersion), semver.Compare(formatVersion(currentVersion), formatVersion(remoteVersion)))
	if semver.Compare(formatVersion(currentVersion), formatVersion(remoteVersion)) < 0 {
		return remoteVersion
	}
	return ""
}
