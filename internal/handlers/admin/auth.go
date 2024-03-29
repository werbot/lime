package handlers

import (
	"math"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/config"
	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/jwtutil"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// SignIn is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/sign/in [post]]
func SignIn(c *fiber.Ctx) error {
	cfg := config.Data()
	log := logging.New()

	request := &models.SignIn{}
	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}
	if err := request.Validate(); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	if request.Email != cfg.Admin.Email || request.Password != cfg.Admin.Password {
		return webutil.StatusUnauthorized(c, errors.MsgWrongPassword)
	}

	token, err := jwtutil.NewToken("admin", cfg.Keys.JWT.Expire, nil)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	metadata, err := jwtutil.ExtractMetadata(token)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	sec, dec := math.Modf(metadata.ExpiresAt)
	expires := time.Unix(int64(sec), int64(dec*(1e9)))

	metaAudit := webutil.GetRequestInfo(c, nil)
	queries.DB().AddAudit(c.Context(), models.SectionSystem, "admin", models.OnSignIn, metaAudit)

	c.Cookie(&fiber.Cookie{
		Name:     "admin",
		Value:    token,
		Expires:  expires,
		SameSite: "lax",
	})

	return webutil.StatusOK(c, "SignIn", nil)
}

// SignOut is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/sign/out [post]
func SignOut(c *fiber.Ctx) error {
	metaAudit := webutil.GetRequestInfo(c, nil)
	queries.DB().AddAudit(c.Context(), models.SectionSystem, "admin", models.OnSignOut, metaAudit)

	c.Cookie(&fiber.Cookie{
		Name:    "admin",
		Expires: time.Now().Add(-(time.Hour * 2)),
		// HTTPOnly: true,
		SameSite: "lax",
	})

	return c.SendStatus(fiber.StatusNoContent)
}
