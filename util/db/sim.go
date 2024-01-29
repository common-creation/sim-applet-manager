package db

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"go.etcd.io/bbolt"
)

type (
	Sim struct {
		Keys []Key `json:"keys"`
	}
	Key struct {
		Name   string `json:"name"`
		AID    string `json:"aid"`
		EncKey string `json:"encKey"`
		MacKey string `json:"macKey"`
		KekKey string `json:"kekKey"`
	}
)

var mutex = sync.Mutex{}

func GetSimConfig(ctx context.Context, iccid string) (*Sim, error) {
	mutex.Lock()
	defer mutex.Unlock()

	db, err := openDB(ctx)
	if err != nil {
		// TODO
		return nil, err
	}
	defer db.Close()

	var sim Sim
	err = db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("sims"))
		if err != nil {
			return err
		}
		simJson := b.Get([]byte(iccid))
		if simJson == nil {
			return errors.New("sim not found")
		}
		return json.Unmarshal(simJson, &sim)
	})
	if err != nil {
		return nil, err
	}
	return &sim, nil
}

func PutSimConfig(ctx context.Context, iccid string, sim *Sim) error {
	mutex.Lock()
	defer mutex.Unlock()

	db, err := openDB(ctx)
	if err != nil {
		// TODO
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("sims"))
		if err != nil {
			return err
		}
		simJson, err := json.Marshal(sim)
		if err != nil {
			return err
		}
		return b.Put([]byte(iccid), simJson)
	})
}
