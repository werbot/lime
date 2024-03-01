package queries

import (
	"embed"

	"github.com/rs/zerolog/log"
	"github.com/werbot/lime/pkg/storage"
)

var db *Base

// Base is ...
type Base struct {
	SettingQueries
	AuthQueries
}

// Init is ...
func Init(cfg storage.Database, embed embed.FS) error {
	database, err := storage.UseStorage(cfg, embed)
	if err != nil {
		log.Err(err).Send()
		return err
	}

	db = &Base{
		SettingQueries: SettingQueries{DB: database},
		AuthQueries:    AuthQueries{DB: database},
	}

	return nil
}

// DB is ...
func DB() *Base {
	if db == nil {
		db = &Base{}
	}
	return db
}
