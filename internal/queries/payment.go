package queries

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/pkg/security"
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
			"payment"."status",
			(
				SELECT
					COUNT(*)
				FROM
					"license"
				WHERE
					"payment_id" = "payment"."id"
			) AS "total_licenses"
		FROM
			"payment"
			LEFT JOIN "customer" ON "payment"."customer_id" = "customer"."id"
			LEFT JOIN "pattern" ON "payment"."pattern_id" = "pattern"."id"
	`

	// paginator init
	query += DB().SQLPagination(webutil.PaginationQuery{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		SortBy: `"payment"."updated_at":DESC`,
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
		licenses := &models.Licenses{}
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
			&licenses.Total,
		)
		if err != nil {
			return nil, err
		}

		payment.Pattern = pattern
		payment.Pattern.Licenses = licenses
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
			(
				SELECT
					json_agg(
						json_build_object(
							'id',
							"license"."id",
							'status',
							"license"."status"
						)
					)
				FROM
					"license"
				WHERE
					"payment_id" = "payment"."id"
			) AS "list_licenses",
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
	var licenseJSON sql.NullString
	payment := &models.Payment{}
	pattern := &models.Pattern{}
	licenses := &models.Licenses{}
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
			&licenseJSON,
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

	payment.Customer = customer
	payment.Pattern = pattern
	if licenseJSON.Valid {
		var licensesTMP []*models.License
		json.Unmarshal([]byte(licenseJSON.String), &licensesTMP)
		licenses.Licenses = licensesTMP
		licenses.Total = len(licensesTMP)
	}
	payment.Pattern.Licenses = licenses

	if metadata.Valid {
		var meta models.Metadata
		json.Unmarshal([]byte(metadata.String), &meta)
		transaction.Meta = meta
	}
	payment.Transaction = transaction

	return payment, nil
}

// AddPayment is ...
func (q *PaymentsQueries) AddPayment(ctx context.Context, payment *models.Payment) error {
	query := `
		INSERT INTO
			"payment" (
				"id",
				"pattern_id",
				"customer_id",
				"provider",
				"status",
				"metadata"
			)
		VALUES
			($1, $2, $3, $4, $5, $6)
	`

	_, err := q.DB.ExecContext(ctx, query,
		security.NanoID(),
		payment.Pattern.ID,
		payment.Customer.ID,
		strconv.Itoa(int(payment.Transaction.Provider)),
		strconv.Itoa(int(payment.Transaction.Status)),
		payment.Transaction.Meta,
	)

	return err
}

// UpdatePayment is ...
func (q *PaymentsQueries) UpdatePayment(ctx context.Context, payment *models.Payment) error {
	meta, err := json.Marshal(payment.Transaction.Meta)
	if err != nil {
		return err
	}

	query := `
		UPDATE "payment"
		SET
			"status" = $2,
			"metadata" = $3,
			"updated_at" = CURRENT_TIMESTAMP
		WHERE
			"id" = $1
`

	_, err = q.DB.ExecContext(ctx, query,
		payment.ID,
		strconv.Itoa(int(payment.Transaction.Status)),
		meta,
	)

	return err
}
