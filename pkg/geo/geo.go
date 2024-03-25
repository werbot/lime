package geo

import (
	"path/filepath"

	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/geo/ipinfo"
	"github.com/werbot/lime/pkg/geo/maxmind"
)

type Storage string

const (
	Maxmind Storage = "maxmind"
	Ipinfo  Storage = "ipinfo"
)

// Database is ...
type Database struct {
	DBPath  string         `toml:"db-path"`
	Storage Storage        `toml:"storage"`
	Maxmind maxmind.Config `toml:"maxmind"`
	Ipinfo  ipinfo.Config  `toml:"ipinfo"`
}

var _ Database

func (db *Database) Download() error {
	if err := fsutil.MkDirs(0o775, db.DBPath); err != nil {
		return err
	}

	var err error
	switch db.Storage {
	case Maxmind:
		err = db.Maxmind.Download(db.DBPath)
	case Ipinfo:
		err = db.Ipinfo.Download(db.DBPath)
	}
	if err != nil {
		return err
	}

	return nil
}

// Check is ...
func (db *Database) Check() bool {
	pathDB := db.MMDBPath()
	var condition bool

	switch db.Storage {
	case Maxmind:
		condition = !fsutil.IsFile(pathDB) && db.Maxmind.AccountID > 0 && len(db.Maxmind.LicenseKey) > 30
	case Ipinfo:
		condition = !fsutil.IsFile(pathDB) && len(db.Ipinfo.Token) > 10
	default:
		return false
	}

	return condition
}

// DBPath is ...
func (db *Database) MMDBPath() string {
	var pathDB string

	switch db.Storage {
	case Maxmind:
		pathDB = filepath.Join(db.DBPath, db.Maxmind.DBName)
	case Ipinfo:
		pathDB = filepath.Join(db.DBPath, db.Ipinfo.DBName)
	}

	return pathDB
}
