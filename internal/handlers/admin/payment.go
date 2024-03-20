package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// Payments is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/payment [get]
func Payments(c *fiber.Ctx) error {
	log := logging.New()

	pagination := webutil.GetPaginationFromCtx(c)
	payments, err := queries.DB().Payments(c.Context(), pagination)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "Payments", payments)
}

// Payment is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/payment/:id [get]
func Payment(c *fiber.Ctx) error {
	log := logging.New()

	payment, err := queries.DB().Payment(c.Context(), c.Params("id"))
	if err != nil {
		if err == errors.ErrPaymentNotFound {
			return webutil.StatusNotFound(c, errors.MsgPaymentNotFound)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}
	return webutil.StatusOK(c, "Payment info", payment)
}
