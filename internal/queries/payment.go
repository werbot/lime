package queries

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/pkg/webutil"
)

// PaymentsQueries is ...
type PaymentsQueries struct {
	*sql.DB
}

// Payments is ...
func (q *PaymentsQueries) Payments(ctx context.Context, pagination *webutil.PaginationQuery) (*models.Payments, error) {
	query := `
		SELECT
			"payment"."id",
			"payment"."pattern_id",
			"pattern"."name"        AS "pattern_name",
			"pattern"."term"        AS "pattern_term",
			"pattern"."price"       AS "pattern_price",
			"pattern"."currency"    AS "pattern_currency",
			"payment"."customer_id",
			"customer"."email"      AS "customer_email",
			"customer"."status"     AS "customer_status",  
			"payment"."provider",
			"payment"."status"
		FROM
			"payment"
			LEFT JOIN "customer" ON "payment"."customer_id" = "customer"."id"
			LEFT JOIN "pattern" ON "payment"."pattern_id" = "pattern"."id"
	`

	// paginator init
	query += DB().SQLPagination(webutil.PaginationQuery{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		SortBy: `"payment"."created_at":DESC`,
	})

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.Payments{}
	for rows.Next() {
		payment := &models.Payment{}
		pattern := &models.Pattern{}
		customer := &models.Customer{}
		transaction := &models.Transaction{}
		err := rows.Scan(
			&payment.ID,
			&pattern.ID,
			&pattern.Name,
			&pattern.Term,
			&pattern.Price,
			&pattern.Currency,
			&customer.ID,
			&customer.Email,
			&customer.Status,
			&transaction.Provider,
			&transaction.Status,
		)
		if err != nil {
			return nil, err
		}

		payment.Pattern = pattern
		payment.Customer = customer
		payment.Transaction = transaction

		response.Payments = append(response.Payments, payment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Count total records
	query = `SELECT COUNT(DISTINCT payment.id) FROM payment`
	err = q.DB.QueryRowContext(ctx, query).Scan(&response.Total)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return response, nil
}

// Payment is ...
func (q *PaymentsQueries) Payment(ctx context.Context, id string) (*models.Payment, error) {
	query := `
		SELECT
			"payment"."id",
			"payment"."pattern_id",
			"pattern"."name"        AS "pattern_name",
			"pattern"."term"        AS "pattern_term",
			"pattern"."price"       AS "pattern_price",
			"pattern"."currency"    AS "pattern_currency",
			"payment"."customer_id",
			"customer"."email"      AS "customer_email",
			"customer"."status"     AS "customer_status",
			"payment"."provider",
			"payment"."status",
			"payment"."metadata",
			"payment"."created_at",
			"payment"."updated_at"
		FROM
			"payment"
			LEFT JOIN "customer" ON "payment"."customer_id" = "customer"."id"
			LEFT JOIN "pattern" ON "payment"."pattern_id" = "pattern"."id"
		WHERE
			"payment"."id" = $1
	`

	var metadata sql.NullString
	payment := &models.Payment{}
	pattern := &models.Pattern{}
	customer := &models.Customer{}
	transaction := &models.Transaction{}

	err := q.DB.QueryRowContext(ctx, query, id).
		Scan(
			&payment.ID,
			&pattern.ID,
			&pattern.Name,
			&pattern.Term,
			&pattern.Price,
			&pattern.Currency,
			&customer.ID,
			&customer.Email,
			&customer.Status,
			&transaction.Provider,
			&transaction.Status,
			&metadata,
			&payment.Created,
			&payment.Updated,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrPatternNotFound
		}
		return nil, err
	}

	payment.Pattern = pattern
	payment.Customer = customer

	if metadata.Valid {
		var meta models.Metadata
		json.Unmarshal([]byte(metadata.String), &meta)
		transaction.Meta = meta
	}
	payment.Transaction = transaction

	return payment, nil
}
