package db

import (
	"context"
	"path/filepath"

	"github.com/common-creation/sim-applet-manager/util/apppath"
	"github.com/common-creation/sim-applet-manager/util/i18n"

	"go.etcd.io/bbolt"
)

func openDB(ctx context.Context, i18n *i18n.I18n) (*bbolt.DB, error) {
	dbPath := filepath.Join(apppath.MustAppDirPath(ctx, i18n), "bolt.db")
	db, err := bbolt.Open(dbPath, 0660, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
