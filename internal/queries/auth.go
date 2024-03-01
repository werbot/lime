package queries

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
)

// AuthQueries is ...
type AuthQueries struct {
	*sql.DB
}

// CustomerIDByEmail is ...
func (q *AuthQueries) CustomerIDByEmail(ctx context.Context, email string) (string, error) {
	var id sql.NullString

	query := `SELECT id FROM customer WHERE email = $1 AND status = true`
	err := q.DB.QueryRowContext(ctx, query, email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.ErrCustomerNotFound
		}
		return "", err
	}

	return id.String, nil
}

// AccessLinkLetter
func (q *AuthQueries) AccessLinkLetter(ctx context.Context, email, token, expires string) (*models.MessageMail, error) {
	siteSetting, err := GetSettingByGroup[models.Site](ctx)
	if err != nil {
		return nil, err
	}

	mailLetter, err := db.GetSettingByKey(ctx, "site", "mail_letter_access_link")
	if err != nil {
		return nil, err
	}
	letterTemplate := models.Letter{}
	if err := json.Unmarshal([]byte(mailLetter["mail_letter_access_link"].Value.(string)), &letterTemplate); err != nil {
		return nil, err
	}

	mail := &models.MessageMail{
		To:     email,
		Letter: letterTemplate,
		Data: map[string]string{
			"Domain":       siteSetting.Domain,
			"Name":         siteSetting.Name,
			"Signature":    siteSetting.Signature,
			"EmailSupport": siteSetting.EmailSupport,
			"Expire":       expires,
			"Token":        token,
		},
	}

	return mail, nil
}
