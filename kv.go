package main

import (
	"github.com/danthegoodman1/RedKey/utils"
	"github.com/dgraph-io/badger/v4"
)

var (
	kv *badger.DB
)

func init() {
	var err error
	kv, err = badger.Open(badger.DefaultOptions(utils.Env_BadgerPath))
	if err != nil {
		logger.Fatal().Err(err).Msg("error opening badger")
	}
}
