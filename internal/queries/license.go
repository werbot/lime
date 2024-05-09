package queries

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/werbot/lime/internal/config"
	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/license"
	"github.com/werbot/lime/pkg/security"
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
		queryAddonCustomer = `WHERE "payment"."customer_id" = '` + customerID + `'`
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
			"license"."data",
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
	lic := &models.License{}
	payment := &models.Payment{
		Customer: &models.Customer{},
		Pattern:  &models.Pattern{},
	}
	err := q.DB.QueryRowContext(ctx, query, request[1]).
		Scan(
			&lic.ID,
			&lic.Status,
			&lic.Created,
			&lic.Updated,
			&payment.ID,
			&payment.Customer.ID,
			&payment.Customer.Status,
			&payment.Customer.Email,
			&lic.Hash,
			&lic.Data,
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

	lic.Payment = payment
	return lic, nil
}

// AddLicense is ...
func (q *LicenseQueries) AddLicense(ctx context.Context, payment *models.Payment) error {
	query := `
		SELECT
			"id"
		FROM
			"license"
		WHERE
			"payment_id" = $1
	`

	var licenseID string
	q.DB.QueryRowContext(ctx, query, payment.ID).Scan(&licenseID)
	if licenseID != "" {
		return errors.ErrLicenseLinkedToPayment
	}

	checkByteSlice, err := json.Marshal(payment.Pattern.Check)
	if err != nil {
		return err
	}

	paymentInfo, err := db.Payment(ctx, payment.ID)
	if err != nil {
		return err
	}

	var limitsSlice license.Limits
	for key, value := range *paymentInfo.Pattern.Limit {
		limitsSlice.Limits = append(limitsSlice.Limits, license.Limit{Key: key, Value: int(value.(float64))})
	}

	cfg := config.Data()

	privKey := fsutil.MustReadFile(filepath.Join(cfg.Keys.KeyDir, cfg.Keys.License.PrivateKey))
	licenseData, err := license.DecodePrivateKey(privKey)
	if err != nil {
		return err
	}
	licenseData.License = license.License{
		IssuedBy:     paymentInfo.Customer.Email,
		CustomerID:   paymentInfo.Customer.ID,
		SubscriberID: payment.ID,
		Type:         paymentInfo.Pattern.Name,
		Limit:        limitsSlice,
		Metadata:     checkByteSlice,
		ExpiresAt:    paymentInfo.Transaction.Payment.UTC().Add(paymentInfo.Pattern.Term.ToDuration()),
		IssuedAt:     paymentInfo.Transaction.Payment.UTC(),
	}

	encoded, err := licenseData.Encode()
	if err != nil {
		return err
	}

	query = `
			INSERT INTO
				"license" (
					"id",
					"payment_id",
					"hash",
					"data",
					"check"
				)
			VALUES
				($1, $2, $3, $4, $5)
		`

	hash := md5.Sum(encoded)
	_, err = q.DB.ExecContext(ctx, query,
		security.NanoID(),
		payment.ID,
		hex.EncodeToString(hash[:]),
		base64.StdEncoding.EncodeToString(encoded),
		payment.Pattern.Check,
	)

	return err
}
