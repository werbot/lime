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
			"status",
			(
				SELECT
					COUNT(*)
				FROM
					"payment"
				WHERE
					"pattern_id" = "pattern"."id"
			) AS "total_licenses"
		FROM
			"pattern"
	`

	// paginator init
	query += DB().SQLPagination(webutil.PaginationQuery{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		// SortBy: `CASE WHEN "updated_at" IS NULL THEN 1 ELSE 0 END,"updated_at":DESC,"created_at":DESC`,
		SortBy: `"updated_at":DESC`,
	})

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.Patterns{}
	for rows.Next() {
		pattern := &models.Pattern{}
		licenses := &models.Licenses{}
		err := rows.Scan(
			&pattern.ID,
			&pattern.Name,
			&pattern.Term,
			&pattern.Price,
			&pattern.Currency,
			&pattern.Private,
			&pattern.Status,
			&licenses.Total,
		)
		if err != nil {
			return nil, err
		}

		pattern.Licenses = licenses
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
			"updated_at",
			(
				SELECT
					COUNT(*)
				FROM
					"payment"
				WHERE
					"pattern_id" = "pattern"."id"
			) AS "total_licenses"
		FROM
			"pattern"
		WHERE "id" = $1
	`

	var limit, check sql.NullString
	pattern := &models.Pattern{}
	licenses := &models.Licenses{}
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
			&pattern.Updated,
			&licenses.Total,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrPatternNotFound
		}
		return nil, err
	}

	if limit.Valid {
		var meta *models.Metadata
		json.Unmarshal([]byte(limit.String), &meta)
		pattern.Limit = meta
	}

	if check.Valid {
		var meta *models.Metadata
		json.Unmarshal([]byte(check.String), &meta)
		pattern.Check = meta
	}

	pattern.Licenses = licenses

	return pattern, nil
}

// AddPattern is ...
func (q *PatternQueries) AddPattern(ctx context.Context, pattern *models.Pattern) error {
	query := `
		INSERT INTO
			"pattern" (
				"id",
				"name",
				"limit",
				"term",
				"price",
				"currency",
				"check",
				"private",
				"status"
			)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := q.DB.ExecContext(ctx, query,
		security.NanoID(),
		pattern.Name,
		pattern.Limit,
		strconv.Itoa(int(*pattern.Term)),
		strconv.Itoa(*pattern.Price),
		strconv.Itoa(int(*pattern.Currency)),
		pattern.Check,
		pattern.Private,
		pattern.Status,
	)

	return err
}

// ClonePattern is ...
func (q *PatternQueries) ClonePattern(ctx context.Context, pattern *models.Pattern) (*models.Pattern, error) {
	query := `
		INSERT INTO
			"pattern" (
				"id",
				"name",
				"limit",
				"term",
				"price",
				"currency",
				"check",
				"private",
				"status"
			)
		SELECT
			$1,
			$2,
			"limit",
			"term",
			"price",
			"currency",
			"check",
			"private",
			"status"
		FROM
			"pattern"
		WHERE
			"pattern"."id" = $3
	`

	newID := security.NanoID()
	_, err := q.DB.ExecContext(ctx, query,
		newID,
		pattern.Name,
		pattern.ID,
	)
	if err != nil {
		return nil, err
	}

	return q.Pattern(ctx, newID)
}

// UpdatePattern is ...
func (q *PatternQueries) UpdatePattern(ctx context.Context, pattern *models.Pattern) error {
	limit, err := json.Marshal(pattern.Limit)
	if err != nil {
		return err
	}

	check, err := json.Marshal(pattern.Check)
	if err != nil {
		return err
	}

	query := `
		UPDATE "pattern"
		SET
			"name" = $2,
			"limit" = $3,
			"check" = $4,
			"term" = $5,
			"price" = $6,
			"currency" = $7,
			"private" = $8,
			"status" = $9,
			"updated_at" = CURRENT_TIMESTAMP
		WHERE
			"id" = $1
`

	_, err = q.DB.ExecContext(ctx, query,
		pattern.ID,
		pattern.Name,
		limit,
		check,
		strconv.Itoa(int(*pattern.Term)),
		strconv.Itoa(*pattern.Price),
		strconv.Itoa(int(*pattern.Currency)),
		pattern.Private,
		pattern.Status,
	)

	return err
}

// DeletePattern is ...
func (q *PatternQueries) DeletePattern(ctx context.Context, id string) error {
	queryCount := `
		SELECT
			COUNT(*)
		FROM
			"payment"
		WHERE
			"pattern_id" = $1
	`

	var count int
	q.DB.QueryRowContext(ctx, queryCount, id).Scan(&count)
	if count > 0 {
		return errors.ErrPatternNotDeleted
	}

	query := `
		DELETE FROM "pattern"
		WHERE
			"id" = $1
	`

	_, err := q.DB.ExecContext(ctx, query, id)
	return err
}
