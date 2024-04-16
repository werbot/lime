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
func (q *LicenseQueries) Licenses(ctx context.Context, pagination *webutil.PaginationQuery, customerID string) (*models.Licenses, error) {
	response := &models.Licenses{}

	queryAddonCustomer := ""
	if customerID != "" {
		queryAddonCustomer = `WHERE "payment"."customer_id" = '7v38n58hXHVsNxS'`
	}

	// Count total records
	queryTotal := `
		SELECT 
			COUNT(DISTINCT "license"."id") 
		FROM 
			"license"
			LEFT JOIN "payment" ON "license"."payment_id" = "payment"."id"
		` + queryAddonCustomer
	err := q.DB.QueryRowContext(ctx, queryTotal).Scan(&response.Total)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

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
	` + queryAddonCustomer

	// paginator init
	query += DB().SQLPagination(webutil.PaginationQuery{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		SortBy: `"license"."updated_at":DESC`,
	})

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		license := models.License{}
		payment := models.Payment{
			Customer: &models.Customer{},
			Pattern:  &models.Pattern{},
		}
		err := rows.Scan(
			&license.ID,
			&license.Status,
			&license.Created,
			&license.Updated,
			&payment.Customer.ID,
			&payment.Customer.Email,
			&payment.Customer.Status,
			&payment.Pattern.Name,
			&payment.Pattern.ID,
			&payment.Pattern.Term,
			&payment.Pattern.Price,
			&payment.Pattern.Currency,
		)
		if err != nil {
			return nil, err
		}

		license.Payment = &payment
		response.Licenses = append(response.Licenses, &license)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

// License is ...
func (q *LicenseQueries) License(ctx context.Context, id string, customerID string) (*models.License, error) {
	request := strings.Split(id, "_")

	query := `
		SELECT
			"license"."id",
			"license"."status",
			"license"."created_at",
			"license"."updated_at",
			"payment"."id",
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

	if customerID != "" {
		query += ``
	}

	var limit sql.NullString
	license := &models.License{}
	payment := &models.Payment{
		Customer: &models.Customer{},
		Pattern:  &models.Pattern{},
	}
	err := q.DB.QueryRowContext(ctx, query, request[1]).
		Scan(
			&license.ID,
			&license.Status,
			&license.Created,
			&license.Updated,
			&payment.ID,
			&payment.Customer.ID,
			&payment.Customer.Status,
			&payment.Customer.Email,
			&license.Hash,
			&payment.Pattern.Name,
			&payment.Pattern.ID,
			&limit,
			&payment.Pattern.Term,
			&payment.Pattern.Price,
			&payment.Pattern.Currency,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrLicenseNotFound
		}
		return nil, err
	}

	if limit.Valid {
		var meta *models.Metadata
		json.Unmarshal([]byte(limit.String), &meta)
		payment.Pattern.Limit = meta
	}

	license.Payment = payment

	return license, nil
}
