package queries

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/pkg/webutil"
)

// PatternQueries is ...
type PatternQueries struct {
	*sql.DB
}

// Patterns is ...
func (q *PatternQueries) Patterns(ctx context.Context, pagination *webutil.PaginationQuery) (*models.Patterns, error) {
	query := `
		SELECT
			"id",
			"name",
			"term",
			"price",
			"currency",
			"private",
			"status"
		FROM
			"pattern"
	`

	// paginator init
	query += DB().SQLPagination(webutil.PaginationQuery{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		SortBy: `"created_at":ASC`,
	})

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.Patterns{}
	for rows.Next() {
		pattern := models.Pattern{}
		err := rows.Scan(
			&pattern.ID,
			&pattern.Name,
			&pattern.Term,
			&pattern.Price,
			&pattern.Currency,
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

// Pattern is ...
func (q *PatternQueries) Pattern(ctx context.Context, id string) (*models.Pattern, error) {
	query := `
		SELECT
			"id",
			"name",
			"limit",
			"check",
			"term",
			"price",
			"currency",
			"private",
			"status",
			"created_at",
			"updated_at"
		FROM
			"pattern"
		WHERE "id" = $1
	`

	var updated sql.NullTime
	var limit, check sql.NullString
	pattern := &models.Pattern{}
	err := q.DB.QueryRowContext(ctx, query, id).
		Scan(
			&pattern.ID,
			&pattern.Name,
			&limit,
			&check,
			&pattern.Term,
			&pattern.Price,
			&pattern.Currency,
			&pattern.Private,
			&pattern.Status,
			&pattern.Created,
			&updated,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrPatternNotFound
		}
		return nil, err
	}

	if updated.Valid {
		pattern.Updated = &updated.Time
	}
	if limit.Valid {
		var meta map[string]any
		json.Unmarshal([]byte(limit.String), &meta)
		pattern.Limit = meta
	}
	if check.Valid {
		var meta map[string]any
		json.Unmarshal([]byte(check.String), &meta)
		pattern.Check = meta
	}

	return pattern, nil
}
