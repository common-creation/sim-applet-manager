package log

import (
	"context"
	"os"
	"path"

	"github.com/common-creation/sim-applet-manager/util/apppath"
	"github.com/common-creation/sim-applet-manager/util/i18n"
)

func Rotate(ctx context.Context, i18n *i18n.I18n) {
	appPath := apppath.MustAppDirPath(ctx, i18n)
	oldLogPath := path.Join(appPath, "log.1.txt")
	if _, err := os.Stat(oldLogPath); err == nil || !os.IsNotExist(err) {
		if err := os.Remove(oldLogPath); err != nil {
			// TODO: エラーダイアログ
		}
	}
	currentLogPath := path.Join(appPath, "log.txt")
	if _, err := os.Stat(currentLogPath); err == nil || !os.IsNotExist(err) {
		if err := os.Rename(currentLogPath, oldLogPath); err != nil {
			// TODO: エラーダイアログ
		}
	}
}

func WriteString(ctx context.Context, i18n *i18n.I18n, message string) {
	appPath := apppath.MustAppDirPath(ctx, i18n)
	currentLogPath := path.Join(appPath, "log.txt")
	f, err := os.OpenFile(currentLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// TODO: エラーダイアログ
		return
	}
	defer f.Close()
	if _, err := f.WriteString(message); err != nil {
		// TODO: エラーダイアログ
	}
}
