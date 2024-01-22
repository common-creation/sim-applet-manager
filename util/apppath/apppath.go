package apppath

import (
	"context"
	"os"
	"path/filepath"

	wailsRutime "github.com/wailsapp/wails/v2/pkg/runtime"
)

func MustAppDirPath(ctx context.Context) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		wailsRutime.MessageDialog(ctx, wailsRutime.MessageDialogOptions{
			Title:         "ディレクトリ エラー",
			Type:          wailsRutime.ErrorDialog,
			Message:       "ユーザーホームディレクトリを取得できません",
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
			wailsRutime.MessageDialog(ctx, wailsRutime.MessageDialogOptions{
				Title:         "ディレクトリ エラー",
				Type:          wailsRutime.ErrorDialog,
				Message:       "アプリケーション ディレクトリを作成できません",
				Buttons:       []string{"OK"},
				DefaultButton: "OK",
			})
			panic(err)
		}
	} else if !info.IsDir() {
		wailsRutime.MessageDialog(ctx, wailsRutime.MessageDialogOptions{
			Title:         "ディレクトリ エラー",
			Type:          wailsRutime.ErrorDialog,
			Message:       "アプリケーション ディレクトリがコンフリクトしています",
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})
		panic("directory already exists and is not a directory")
	}
	return gpnDir
}
