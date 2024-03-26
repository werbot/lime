package geo

import (
	"fmt"

	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/geo/geoopen"
	"github.com/werbot/lime/pkg/geo/ipinfo"
	"github.com/werbot/lime/pkg/geo/maxmind"
)

type Storage string

const (
	Maxmind Storage = "maxmind"
	Ipinfo  Storage = "ipinfo"
	GeoOpen Storage = "geoopen"
)

// Database is ...
type Database struct {
	DBPath  string         `toml:"db-path"`
	Storage Storage        `toml:"storage"`
	GeoOpen geoopen.Config `toml:"geoopen"`
	Maxmind maxmind.Config `toml:"maxmind,commented"`
	Ipinfo  ipinfo.Config  `toml:"ipinfo,commented"`
}

// var _ Database
func (db *Database) Download() error {
	if err := fsutil.MkDirs(0o775, db.DBPath); err != nil {
		return err
	}

	switch db.Storage {
	case GeoOpen:
		if err := geoopen.DB(db.DBPath, db.GeoOpen).Download(); err != nil {
			return err
		}
	case Maxmind:
		if err := maxmind.DB(db.DBPath, db.Maxmind).Download(); err != nil {
			return err
		}
	case Ipinfo:
		if err := ipinfo.DB(db.DBPath, db.Ipinfo).Download(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported storage type")
	}

	return nil
}

// Check is ...
func (db *Database) Check() bool {
	switch db.Storage {
	case GeoOpen:
		return geoopen.DB(db.DBPath, db.GeoOpen).Check()
	case Maxmind:
		return maxmind.DB(db.DBPath, db.Maxmind).Check()
	case Ipinfo:
		return ipinfo.DB(db.DBPath, db.Ipinfo).Check()
	}

	return false
}
