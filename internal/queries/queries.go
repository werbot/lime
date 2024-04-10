package queries

import (
	"embed"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/werbot/lime/pkg/storage"
	"github.com/werbot/lime/pkg/webutil"
)

var db *Base

// Base is ...
type Base struct {
	AuditQueries
	SettingQueries
	AuthQueries
	LicenseQueries
	PatternQueries
	CustomerQueries
	PaymentsQueries
	ListQueries
}

// Init is ...
func Init(cfg storage.Database, embed embed.FS) error {
	database, err := storage.UseStorage(cfg, embed)
	if err != nil {
		log.Err(err).Send()
		return err
	}

	db = &Base{
		AuditQueries:    AuditQueries{DB: database},
		SettingQueries:  SettingQueries{DB: database},
		AuthQueries:     AuthQueries{DB: database},
		LicenseQueries:  LicenseQueries{DB: database},
		PatternQueries:  PatternQueries{DB: database},
		CustomerQueries: CustomerQueries{DB: database},
		PaymentsQueries: PaymentsQueries{DB: database},
		ListQueries:     ListQueries{DB: database},
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

// SQLPagination is ...
// example query's for sortBy - id:DESC or id:ASC
func (db *Base) SQLPagination(params webutil.PaginationQuery) string {
	if params.Offset < 0 {
		params.Offset = 0
	}

	if params.Limit <= 0 {
		params.Limit = 30
	}

	var showSortBy string
	if len(params.SortBy) > 0 {
		showSortBy = "ORDER BY "

		var orderParts []string
		sorts := strings.Split(params.SortBy, ",")
		for _, sort := range sorts {
			parts := strings.SplitN(sort, ":", 2)
			if len(parts) == 1 {
				orderParts = append(orderParts, parts[0])
			}
			if len(parts) == 2 {
				orderParts = append(orderParts, fmt.Sprintf("%s %s", parts[0], parts[1]))
			}
		}
		showSortBy = showSortBy + strings.Join(orderParts, ", ")
	}

	return fmt.Sprintf(" %s LIMIT %d OFFSET %d", showSortBy, params.Limit, params.Offset)
}
