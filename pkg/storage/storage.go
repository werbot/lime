package storage

import (
	"context"
	"database/sql"
	"embed"
	"errors"

	"github.com/pressly/goose/v3"
	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/storage/postgres"
	"github.com/werbot/lime/pkg/storage/sqlite"
)

type Storage string

const (
	Sqlite   Storage = "sqlite"
	Postgres Storage = "postgres"
)

// Database is ...
type Database struct {
	Storage  Storage         `toml:"storage"`
	Sqlite   sqlite.Config   `toml:"sqlite"`
	Postgres postgres.Config `toml:"postgres"`
}

func UseStorage(database Database, migrations embed.FS) (*sql.DB, error) {
	switch database.Storage {
	case Sqlite:
		db, err := sqlite.New(&database.Sqlite)
		if err != nil {
			return nil, err
		}

		if !fsutil.IsFile(database.Sqlite.DBPath) {
			if _, err := fsutil.OpenFile(database.Sqlite.DBPath, fsutil.FsCWFlags, 0666); err != nil {
				return nil, err
			}

			goose.SetBaseFS(migrations)
			goose.SetTableName("migrate_db_version")
			if err := goose.SetDialect(string(Sqlite)); err != nil {
				return nil, err
			}
			if err := goose.Up(db, "."); err != nil {
				return nil, err
			}
		}

		return db, nil
	case Postgres:
		db, err := postgres.New(context.Background(), &database.Postgres)
		if err != nil {
			return nil, err
		}

		if _, err := db.Query(`SELECT * FROM migrate_db_version`); err != nil {
			goose.SetBaseFS(migrations)
			goose.SetTableName("migrate_db_version")
			if err := goose.SetDialect(string(Postgres)); err != nil {
				return nil, err
			}
			if err := goose.Up(db, "."); err != nil {
				return nil, err
			}
		}

		return db, nil
	}

	return nil, errors.New("Invalid storage")
}
