package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// Licenses is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/license [get]
func Licenses(c *fiber.Ctx) error {
	log := logging.New()

	pagination := webutil.GetPaginationFromCtx(c)
	licenses, err := queries.DB().Licenses(c.Context(), pagination, "")
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "Licenses", licenses)
}

// License is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/license/:id [get]
func License(c *fiber.Ctx) error {
	log := logging.New()

	license, err := queries.DB().License(c.Context(), c.Params("id"), "")
	if err != nil {
		if err == errors.ErrLicenseNotFound {
			return webutil.StatusNotFound(c, errors.MsgLicenseNotFound)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}
	return webutil.StatusOK(c, "License info", license)
}

// AddLicense is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/license [post]
func AddLicense(c *fiber.Ctx) error {
	log := logging.New()
	// cfg := config.Data()

	request := &models.Payment{}
	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	err := queries.DB().AddLicense(c.Context(), request)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "Create license", nil)
}

// UpdateLicense is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/license/:id [patch]
func UpdateLicense(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Update license", nil)
}
