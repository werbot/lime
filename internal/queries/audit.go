package queries

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/pkg/geo"
	"github.com/werbot/lime/pkg/security"
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
			"audit"."id",
			"audit"."section",
			"audit"."customer_id",
			"customer"."email"    AS "customer_email",
			"customer"."status"   AS "customer_status",
			"audit"."action",
			"audit"."created_at"
		FROM
			"audit"
			LEFT JOIN "customer" ON "audit"."customer_id" = "customer"."id"
	`

	// paginator init
	query += DB().SQLPagination(webutil.PaginationQuery{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		SortBy: `"audit"."created_at":DESC`,
	})

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &models.Audits{}
	for rows.Next() {
		var email sql.NullString
		var status sql.NullBool
		audit := models.Audit{}
		customer := models.Customer{}
		err := rows.Scan(
			&audit.ID,
			&audit.Section,
			&customer.ID,
			&email,
			&status,
			&audit.Action,
			&audit.Created,
		)
		if err != nil {
			return nil, err
		}

		if email.Valid {
			customer.Email = email.String
		}

		if customer.ID == "admin" {
			customer.Status = true
			customer.Email = "admin"
		} else if status.Valid {
			customer.Status = status.Bool
		}

		audit.Customer = customer
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
			"audit"."id",
			"audit"."section",
			"audit"."customer_id",
			"customer"."email"    AS "customer_email",
			"customer"."status"   AS "customer_status",
			"audit"."action",
			"audit"."metadata",
			"audit"."created_at"
		FROM
			"audit"
			LEFT JOIN "customer" ON "audit"."customer_id" = "customer"."id"
		WHERE
			"audit"."id" = $1
	`

	var status sql.NullBool
	var email, metadata sql.NullString
	audit := &models.Audit{}
	customer := models.Customer{}
	err := q.DB.QueryRowContext(ctx, query, id).
		Scan(
			&audit.ID,
			&audit.Section,
			&customer.ID,
			&email,
			&status,
			&audit.Action,
			&metadata,
			&audit.Created,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrCustomerNotFound
		}
		return nil, err
	}

	if email.Valid {
		customer.Email = email.String
	}

	if customer.ID == "admin" {
		customer.Status = true
		customer.Email = "admin"
	} else if status.Valid {
		customer.Status = status.Bool
	}

	if metadata.Valid {
		// var meta map[string]any
		var meta webutil.MetaInfo
		json.Unmarshal([]byte(metadata.String), &meta)
		isoCountry := meta.Request.UserCountry
		meta.Request.UserCountry = fmt.Sprintf("%s %s", geo.FlagEmoji(isoCountry), geo.FullName(isoCountry))
		audit.Metadata = meta
	}

	audit.Customer = customer

	return audit, nil
}

// AddAudit is ...
func (q *AuditQueries) AddAudit(ctx context.Context, section models.Section, customerID string, action models.AuditAction, metadata any) error {
	id := security.NanoID()
	sectionStr := strconv.Itoa(int(section))
	actionStr := strconv.Itoa(int(action))
	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	_, err = q.DB.ExecContext(ctx, `INSERT INTO audit (id, section, customer_id, action, metadata) VALUES ($1, $2, $3, $4, $5)`,
		id,
		sectionStr,
		customerID,
		actionStr,
		metadataJSON,
	)

	return err
}
