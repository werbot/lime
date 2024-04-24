package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// Payments is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/payment/:filter? [get]
func Payments(c *fiber.Ctx) error {
	log := logging.New()

	var payments *models.Payments
	var err error

	if c.Params("filter") == "no_lic" {
		payments, err = queries.DB().PaymentsNoLicense(c.Context())
	} else {
		pagination := webutil.GetPaginationFromCtx(c)
		payments, err = queries.DB().Payments(c.Context(), pagination)
	}

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

// AddPayment is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/payment [post]
func AddPayment(c *fiber.Ctx) error {
	log := logging.New()
	request := &models.Payment{}

	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	err := queries.DB().AddPayment(c.Context(), request)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	// adit section
	metaAudit := webutil.GetRequestInfo(c, request)
	queries.DB().AddAudit(c.Context(), models.SectionPayment, "admin", models.OnAdd, metaAudit)

	return webutil.StatusOK(c, "Pattern added", nil)
}

// UpdatePayment is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/payment/:id [patch]
func UpdatePayment(c *fiber.Ctx) error {
	log := logging.New()
	request := &models.Payment{}

	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	err := queries.DB().UpdatePayment(c.Context(), request)
	if err != nil {
		if err == errors.ErrPatternNotFound {
			return webutil.StatusNotFound(c, errors.MsgPatternNotFound)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}
	// adit section
	request.Created = nil
	request.Updated = nil
	metaAudit := webutil.GetRequestInfo(c, request)
	queries.DB().AddAudit(c.Context(), models.SectionPayment, "admin", models.OnUpdate, metaAudit)

	return webutil.StatusOK(c, "Pattern updated", nil)
}
