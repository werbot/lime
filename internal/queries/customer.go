package queries

import (
	"context"
	"database/sql"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
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
			"updated_at"
		FROM
			"customer"
	`

	// paginator init
	query += DB().SQLPagination(webutil.PaginationQuery{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		SortBy: `"created_at":DESC`,
	})

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.Customers{}
	for rows.Next() {
		var updated sql.NullTime
		customer := models.Customer{}
		err := rows.Scan(
			&customer.ID,
			&customer.Email,
			&customer.Status,
			&customer.Created,
			&updated,
		)
		if err != nil {
			return nil, err
		}

		if updated.Valid {
			customer.Updated = &updated.Time
		}

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
			"updated_at"
		FROM
			"customer"
		WHERE "id" = $1
	`

	var updated sql.NullTime
	customer := &models.Customer{}
	err := q.DB.QueryRowContext(ctx, query, id).
		Scan(
			&customer.ID,
			&customer.Email,
			&customer.Status,
			&customer.Created,
			&updated,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrCustomerNotFound
		}
		return nil, err
	}

	if updated.Valid {
		customer.Updated = &updated.Time
	}

	return customer, nil
}
