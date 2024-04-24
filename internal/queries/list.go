package queries

import (
	"context"
	"database/sql"

	"github.com/werbot/lime/internal/models"
)

// PatternQueries is ...
type ListQueries struct {
	*sql.DB
}

//  TODO remove this constructions !!!!

// ListPatterns is ...
func (q *ListQueries) ListPatterns(ctx context.Context) (*models.Patterns, error) {
	query := `
		SELECT
			"id",
			"name",
			"private",
			"status"
		FROM
			"pattern"
	`

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.Patterns{}
	for rows.Next() {
		pattern := &models.Pattern{}
		err := rows.Scan(
			&pattern.ID,
			&pattern.Name,
			&pattern.Private,
			&pattern.Status,
		)
		if err != nil {
			return nil, err
		}

		response.Patterns = append(response.Patterns, pattern)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Count total records
	query = `SELECT COUNT(DISTINCT pattern.id) FROM pattern`
	err = q.DB.QueryRowContext(ctx, query).Scan(&response.Total)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return response, nil
}

// ListCustomers is ...
func (q *ListQueries) ListCustomers(ctx context.Context) (*models.Customers, error) {
	query := `
		SELECT
			"id",
			"email",
			"status"
		FROM
			"customer"
	`

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.Customers{}
	for rows.Next() {
		customer := &models.Customer{}
		err := rows.Scan(
			&customer.ID,
			&customer.Email,
			&customer.Status,
		)
		if err != nil {
			return nil, err
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

// ListCountries is ...
func (q *ListQueries) ListCountries(ctx context.Context, name string) ([]models.SettingName, error) {
	query := `
		SELECT
			"code",
			"name"
		FROM
			"country"
		WHERE
			LOWER("name") LIKE LOWER($1)
		ORDER BY
			"name" ASC
		LIMIT
			15
		OFFSET
			0
	`

	rows, err := q.DB.QueryContext(ctx, query, name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := []models.SettingName{}
	for rows.Next() {
		country := models.SettingName{}
		err := rows.Scan(
			&country.Key,
			&country.Value,
		)
		if err != nil {
			return nil, err
		}

		response = append(response, country)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}
