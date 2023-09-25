package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Config is ...
type Config struct {
	User            string `toml:"user"`
	Password        string `toml:"password"`
	Host            string `toml:"host"`
	Database        string `toml:"database"`
	MaxConn         int    `toml:"max-conn,commented"`
	MaxIdleConn     int    `toml:"max-idle-conn,commented"`
	MaxLifetimeConn int    `toml:"max-lifetime-conn,commented"`
}

// New creates a new Connect object using the given PgSQLConfig.
func New(ctx context.Context, conf *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", conf.User, conf.Password, conf.Host, conf.Database)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	// Configure the connection pool.
	db.SetMaxOpenConns(conf.MaxConn)
	db.SetMaxIdleConns(conf.MaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(conf.MaxLifetimeConn) * time.Second)

	// Ping the database to ensure connectivity.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}

	return db, nil
}
