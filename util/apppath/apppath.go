package apppath

import (
	"context"
	"os"
	"path/filepath"

	"github.com/common-creation/sim-applet-manager/util/i18n"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

func MustAppDirPath(ctx context.Context, i18n *i18n.I18n) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		wailsRuntime.MessageDialog(ctx, wailsRuntime.MessageDialogOptions{
			Title:         i18n.T("directoryError"),
			Type:          wailsRuntime.ErrorDialog,
			Message:       i18n.T("errorGetUserDir"),
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})
		panic(err)
	}
	gpnDir := filepath.Join(homeDir, ".simappletmanager")
	info, err := os.Stat(gpnDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(gpnDir, 0755)
		} else {
			wailsRuntime.MessageDialog(ctx, wailsRuntime.MessageDialogOptions{
				Title:         i18n.T("directoryError"),
				Type:          wailsRuntime.ErrorDialog,
				Message:       i18n.T("errorCreateAppDir"),
				Buttons:       []string{"OK"},
				DefaultButton: "OK",
			})
			panic(err)
		}
	} else if !info.IsDir() {
		wailsRuntime.MessageDialog(ctx, wailsRuntime.MessageDialogOptions{
			Title:         i18n.T("directoryError"),
			Type:          wailsRuntime.ErrorDialog,
			Message:       i18n.T("errorAppDirConflict"),
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})
		panic("directory already exists and is not a directory")
	}
	return gpnDir
}
