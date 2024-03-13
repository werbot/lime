package handlers

import (
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/internal/config"
	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/mailer"
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
// @Router /api/sign/in?token= [post]]
func SignIn(c *fiber.Ctx) error {
	cfg := config.Data()
	log := logging.New()

	// if in request token
	tokenReq := c.Query("token")
	if tokenReq != "" {
		metadata, err := jwtutil.ExtractMetadata(tokenReq)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusBadRequest(c, err.Error())
		}

		sec, dec := math.Modf(metadata.ExpiresAt)
		expires := time.Unix(int64(sec), int64(dec*(1e9)))

		c.Cookie(&fiber.Cookie{
			Name:     "manager",
			Value:    tokenReq,
			Expires:  expires,
			SameSite: "lax",
		})
		return webutil.StatusOK(c, "SignIn", metadata)
	}

	// if no token in request
	request := &models.SignIn{}
	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}
	if err := request.Validate(); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	customerID, err := queries.DB().CustomerIDByEmail(c.Context(), request.Email)
	if err != nil {
		if err == errors.ErrCustomerNotFound {
			return webutil.StatusNotFound(c, errors.MsgCustomerNotFound)
		}

		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	token, err := jwtutil.NewToken(customerID, cfg.Keys.JWT.Expire, nil)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	if err := mailer.SendAccessLinkLetter(request.Email, token, cfg.Keys.JWT.Expire); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "SignIn", "Message sent")
}

// SignOut is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/sign/out [post]
func SignOut(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "SignOut", nil)
}
