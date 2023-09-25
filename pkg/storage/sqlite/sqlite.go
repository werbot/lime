package sqlite

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Config is ...
type Config struct {
	DBPath string `toml:"db-path"`
}

// New is ...
func New(conf *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s?_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)&_pragma=journal_size_limit(200000000)&_pragma=synchronous(NORMAL)&_pragma=foreign_keys(ON)", conf.DBPath)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	db.Query("PRAGMA auto_vacuum")

	return db, nil
}
