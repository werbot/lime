package queries

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/pkg/strutil"
)

// SettingQueries wraps a sql.DB connection allowing for easy querying and interaction
// with the database related to application settings.
type SettingQueries struct {
	*sql.DB
}

// GroupFieldMap generates a map of fields based on the type of settings.
func (q *SettingQueries) GroupFieldMap(settings any) models.Metadata {
	switch s := settings.(type) {
	case *models.Site:
		return models.Metadata{
			"site_domain":        &s.Domain,
			"site_name":          &s.Name,
			"site_signature":     &s.Signature,
			"site_email_support": &s.EmailSupport,
		}
	case *models.Mail:
		return models.Metadata{
			"mail_sender_name":  &s.SenderName,
			"mail_sender_email": &s.SenderEmail,
			"smtp_host":         &s.SMTP.Host,
			"smtp_port":         &s.SMTP.Port,
			"smtp_username":     &s.SMTP.Username,
			"smtp_password":     &s.SMTP.Password,
			"smtp_encryption":   &s.SMTP.Encryption,
		}
	default:
		return nil
	}
}

// GetSettingByGroup is a generic function that retrieves a setting from the database.
// It takes a context and a pointer to the Base struct which holds the database methods.
// The function returns a pointer to the requested setting of type T or an error if any occurs.
func GetSettingByGroup[T any](ctx context.Context) (*T, error) {
	setting, err := db.GetSettingByGroup(ctx, new(T))
	if err != nil {
		return nil, err
	}
	return setting.(*T), nil
}

// GetSettingByGroup retrieves settings based on the provided `settings` struct, populating it with values from the database.
func (q *SettingQueries) GetSettingByGroup(ctx context.Context, settings any) (any, error) {
	fieldMap := q.GroupFieldMap(settings)

	if fieldMap == nil {
		return nil, errors.ErrSettingNotFound
	}

	queryAddon := ""
	keys := make([]any, 0, len(fieldMap))
	for k := range fieldMap {
		keys = append(keys, k)
		queryAddon += fmt.Sprintf("$%v, ", len(keys))
	}

	query := fmt.Sprintf("SELECT key, value FROM setting WHERE key IN (%s)", queryAddon[:len(queryAddon)-2])
	rows, err := q.DB.QueryContext(ctx, query, keys...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key, value string
		err := rows.Scan(&key, &value)
		if err != nil {
			return nil, err
		}

		if fieldPtr, ok := fieldMap[key]; ok {
			switch ptr := fieldPtr.(type) {
			case *string:
				*ptr = value
			case *bool:
				bValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, err
				}
				*ptr = bValue
			case *int:
				iValue, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				*ptr = iValue
			}
		}
	}

	return settings, nil
}

// UpdateSettingByGroup updates the settings in the database using a transaction.
// It takes a context and a settings object of any type as arguments.
func (q *SettingQueries) UpdateSettingByGroup(ctx context.Context, settings any) error {
	fieldMap := q.GroupFieldMap(settings)

	tx, err := q.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `UPDATE setting SET value = $1 WHERE key = $2`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for key, value := range fieldMap {
		if _, err = stmt.ExecContext(ctx, value, key); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetSettingByKey retrieves a setting by its key from the database.
// It accepts a context for cancellation and a string representing the key of the setting.
// Returns a pointer to a SettingName model if found, or an error if not found or any other issue occurs.
func (q *SettingQueries) GetSettingByKey(ctx context.Context, key ...string) (map[string]models.SettingName, error) {
	if len(key) == 0 {
		return nil, errors.ErrSettingNotFound
	}

	queryAddon := ""
	keys := make([]any, 0, len(key))
	for k := range key {
		keys = append(keys, k)
		queryAddon += fmt.Sprintf("$%v, ", len(keys))
	}

	query := fmt.Sprintf("SELECT id, key, value FROM setting WHERE key IN (%s)", queryAddon[:len(queryAddon)-2])
	rows, err := q.DB.QueryContext(ctx, query, strutil.ToAny(key...)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	settings := map[string]models.SettingName{}
	for rows.Next() {
		var key string
		setting := models.SettingName{}
		if err := rows.Scan(&setting.ID, &key, &setting.Value); err != nil {
			return nil, err
		}
		settings[key] = setting
	}

	return settings, nil
}

// UpdateSettingByKey updates the value of a setting in the database based on the provided key.
func (q *SettingQueries) UpdateSettingByKey(ctx context.Context, setting *models.SettingName) error {
	query := `UPDATE setting SET value = $1 WHERE key = $2`
	_, err := q.DB.ExecContext(ctx, query, setting.Value, setting.Key)
	return err
}
