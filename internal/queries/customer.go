package queries

import (
	"context"
	"database/sql"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/pkg/security"
	"github.com/werbot/lime/pkg/webutil"
)

// CustomerQueries is ...
type CustomerQueries struct {
	*sql.DB
}

// Customers is ...
func (q *CustomerQueries) Customers(ctx context.Context, pagination *webutil.PaginationQuery) (*models.Customers, error) {
	query := `
		SELECT
			"id",
			"email",
			"status",
			"created_at",
			"updated_at",
			(
				SELECT
					COUNT(*)
				FROM
					"payment"
				WHERE
					"customer_id" = "customer"."id"
			) AS "total_payments"
		FROM
			"customer"
	`

	// paginator init
	query += DB().SQLPagination(webutil.PaginationQuery{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		SortBy: `"updated_at":DESC`,
	})

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.Customers{}
	for rows.Next() {
		customer := &models.Customer{}
		payments := &models.Payments{}
		err := rows.Scan(
			&customer.ID,
			&customer.Email,
			&customer.Status,
			&customer.Created,
			&customer.Updated,
			&payments.Total,
		)
		if err != nil {
			return nil, err
		}

		customer.Payments = payments

		response.Customers = append(response.Customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Count total records
	query = `SELECT COUNT(DISTINCT customer.id) FROM customer`
	err = q.DB.QueryRowContext(ctx, query).Scan(&response.Total)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return response, nil
}

// Customer is ...
func (q *CustomerQueries) Customer(ctx context.Context, id string) (*models.Customer, error) {
	query := `
		SELECT
			"id",
			"email",
			"status",
			"created_at",
			"updated_at",
			(
				SELECT
					COUNT(*)
				FROM
					"payment"
				WHERE
					"customer_id" = "customer"."id"
			) AS "total_payment"
		FROM
			"customer"
		WHERE
			"id" = $1
	`

	customer := &models.Customer{}
	payments := &models.Payments{}
	err := q.DB.QueryRowContext(ctx, query, id).
		Scan(
			&customer.ID,
			&customer.Email,
			&customer.Status,
			&customer.Created,
			&customer.Updated,
			&payments.Total,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrCustomerNotFound
		}
		return nil, err
	}

	customer.Payments = payments

	return customer, nil
}

// AddCustomer is ...
func (q *CustomerQueries) AddCustomer(ctx context.Context, customer *models.Customer) error {
	query := `
		INSERT INTO
			"customer" (
				"id",
				"email",
				"status"
			)
		VALUES
			($1, $2, $3)
	`

	_, err := q.DB.ExecContext(ctx, query,
		security.NanoID(),
		customer.Email,
		customer.Status,
	)

	return err
}

// UpdateCustomer is ...
func (q *CustomerQueries) UpdateCustomer(ctx context.Context, customer *models.Customer) error {
	query := `
		UPDATE "customer"
		SET
			"email" = $2,
			"status" = $3,
			"updated_at" = CURRENT_TIMESTAMP
		WHERE
			"id" = $1
`

	_, err := q.DB.ExecContext(ctx, query,
		customer.ID,
		customer.Email,
		customer.Status,
	)

	return err
}

// DeleteCustomer is ...
func (q *CustomerQueries) DeleteCustomer(ctx context.Context, id string) error {
	queryCount := `
		SELECT
			COUNT(*)
		FROM
			"payment"
		WHERE
			"customer_id" = $1
	`

	var count int
	q.DB.QueryRowContext(ctx, queryCount, id).Scan(&count)
	if count > 0 {
		return errors.ErrCustomerNotDeleted
	}

	query := `
		DELETE FROM "customer"
		WHERE
			"id" = $1
	`

	_, err := q.DB.ExecContext(ctx, query, id)
	return err
}
