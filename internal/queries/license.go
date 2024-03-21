package queries

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/pkg/webutil"
)

// LicenseQueries is ...
type LicenseQueries struct {
	*sql.DB
}

// Licenses is ...
func (q *LicenseQueries) Licenses(ctx context.Context, pagination *webutil.PaginationQuery, admin bool) (*models.Licenses, error) {
	query := `
		SELECT
			"license"."id",
			"license"."status",
			"license"."created_at",
			"license"."updated_at",
			"payment"."customer_id",
			"customer"."email"      AS "customer_email",
			"customer"."status"     AS "customer_status",
			"pattern"."name"        AS "pattern_name",
			"pattern"."id"          AS "pattern_id",
			"pattern"."term"        AS "pattern_term",
			"pattern"."price"       AS "pattern_price",
			"pattern"."currency"    AS "pattern_currency"
		FROM
			"license"
			LEFT JOIN "payment" ON "license"."payment_id" = "payment"."id"
			LEFT JOIN "customer" ON "payment"."customer_id" = "customer"."id"
			LEFT JOIN "pattern" ON "payment"."pattern_id" = "pattern"."id"
	`

	if admin {
		query += ``
	} else {
		query += ``
	}

	// paginator init
	query += DB().SQLPagination(webutil.PaginationQuery{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		SortBy: `"license"."created_at":DESC`,
	})

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.Licenses{}
	for rows.Next() {
		var updated sql.NullTime
		license := models.License{}
		customer := models.Customer{}
		pattern := models.Pattern{}
		err := rows.Scan(
			&license.ID,
			&license.Status,
			&license.Created,
			&updated,
			&customer.ID,
			&customer.Email,
			&customer.Status,
			&pattern.Name,
			&pattern.ID,
			&pattern.Term,
			&pattern.Price,
			&pattern.Currency,
		)
		if err != nil {
			return nil, err
		}

		if updated.Valid {
			license.Updated = &updated.Time
		}

		license.Pattern = &pattern
		license.Customer = &customer
		response.Licenses = append(response.Licenses, &license)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Count total records
	query = `SELECT COUNT(DISTINCT license.id) FROM license`
	err = q.DB.QueryRowContext(ctx, query).Scan(&response.Total)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return response, nil
}

// License is ...
func (q *LicenseQueries) License(ctx context.Context, id string, admin bool) (*models.License, error) {
	request := strings.Split(id, "_")

	query := `
		SELECT
			"license"."id",
			"license"."status",
			"license"."created_at",
			"license"."updated_at",
			"payment"."customer_id",
			"customer"."status"     AS "customer_status",
			"customer"."email"      AS "customer_email",
			"license"."hash",
			"pattern"."name"        AS "pattern_name",
			"pattern"."id"          AS "pattern_id",
			"pattern"."limit"       AS "pattern_limit",
			"pattern"."term"        AS "pattern_term",
			"pattern"."price"       AS "pattern_price",
			"pattern"."currency"    AS "pattern_currency"
		FROM
			"license"
			LEFT JOIN "payment" ON "license"."payment_id" = "payment"."id"
			LEFT JOIN "customer" ON "payment"."customer_id" = "customer"."id"
			LEFT JOIN "pattern" ON "payment"."pattern_id" = "pattern"."id"
		WHERE "license"."id" = $1
	`

	if admin {
		query += ``
	} else {
		query += ``
	}

	var updated sql.NullTime
	var limit sql.NullString
	license := &models.License{}
	customer := &models.Customer{}
	pattern := &models.Pattern{}
	err := q.DB.QueryRowContext(ctx, query, request[1]).
		Scan(
			&license.ID,
			&license.Status,
			&license.Created,
			&updated,
			&customer.ID,
			&customer.Status,
			&customer.Email,
			&license.Hash,
			&pattern.Name,
			&pattern.ID,
			&limit,
			&pattern.Term,
			&pattern.Price,
			&pattern.Currency,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrLicenseNotFound
		}
		return nil, err
	}

	if updated.Valid {
		license.Updated = &updated.Time
	}
	if limit.Valid {
		var meta map[string]any
		json.Unmarshal([]byte(limit.String), &meta)
		pattern.Limit = meta
	}

	license.Pattern = pattern
	license.Customer = customer

	return license, nil
}
