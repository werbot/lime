package queries

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/pkg/webutil"
)

// AuditQueries is ...
type AuditQueries struct {
	*sql.DB
}

// Audits is ...
func (q *AuditQueries) Audits(ctx context.Context, pagination *webutil.PaginationQuery) (*models.Audits, error) {
	query := `
		SELECT
			"id",
			"section",
			"section_id",
			"action",
			"created_at"
		FROM
			"audit"
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

	response := &models.Audits{}
	for rows.Next() {
		audit := models.Audit{}
		err := rows.Scan(
			&audit.ID,
			&audit.Section,
			&audit.SectionID,
			&audit.Action,
			&audit.Created,
		)
		if err != nil {
			return nil, err
		}

		response.Audits = append(response.Audits, audit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Count total records
	query = `SELECT COUNT(DISTINCT audit.id) FROM audit`
	err = q.DB.QueryRowContext(ctx, query).Scan(&response.Total)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return response, nil
}

// Audit is ...
func (q *AuditQueries) Audit(ctx context.Context, id string) (*models.Audit, error) {
	query := `
		SELECT
			"id",
			"section",
			"section_id",
			"action",
			"metadata",
			"created_at"
		FROM
			"audit"
		WHERE 
			"id" = $1
	`
	var metadata sql.NullString
	audit := &models.Audit{}
	err := q.DB.QueryRowContext(ctx, query, id).
		Scan(
			&audit.ID,
			&audit.Section,
			&audit.SectionID,
			&audit.Action,
			&metadata,
			&audit.Created,
			&audit,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrCustomerNotFound
		}
		return nil, err
	}

	if metadata.Valid {
		var meta map[string]any
		json.Unmarshal([]byte(metadata.String), &meta)
		audit.Metadata = meta
	}

	return audit, nil
}
