package db

import (
	"context"
	"github.com/common-creation/sim-applet-manager/util/apppath"
	"path/filepath"

	"go.etcd.io/bbolt"
)

func openDB(ctx context.Context) (*bbolt.DB, error) {
	dbPath := filepath.Join(apppath.MustAppDirPath(ctx), "bolt.db")
	db, err := bbolt.Open(dbPath, 0660, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
