package mailer

import (
	"context"
	"time"

	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/internal/queries"
)

// SendAccessLinkLetter is ...
func SendAccessLinkLetter(email, token, expires string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	letter, err := queries.DB().AuthQueries.AccessLinkLetter(ctx, email, token, expires)
	if err != nil {
		return err
	}

	mailSetting, err := queries.GetSettingByGroup[models.Mail](ctx)
	if err != nil {
		return err
	}

	if err := SendMail(mailSetting, letter); err != nil {
		return err
	}

	return nil
}
