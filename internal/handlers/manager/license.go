package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/jwtutil"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// Licenses is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/license [get]
func Licenses(c *fiber.Ctx) error {
	log := logging.New()
	meta, err := jwtutil.ExtractMetadataFiber(c)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	pagination := webutil.GetPaginationFromCtx(c)
	licenses, err := queries.DB().Licenses(c.Context(), pagination, meta.ID)
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
// @Router /api/license/:id [get]
func License(c *fiber.Ctx) error {
	log := logging.New()
	meta, err := jwtutil.ExtractMetadataFiber(c)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	license, err := queries.DB().License(c.Context(), c.Params("id"), meta.ID)
	if err != nil {
		if err == errors.ErrLicenseNotFound {
			return webutil.StatusNotFound(c, errors.MsgLicenseNotFound)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}
	return webutil.StatusOK(c, "License info", license)
}

// LicenseVerify is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/license/verify [post]
func LicenseVerify(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "License Verify", nil)
}

// DownloadLicense is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/license/download/:id [get]
func DownloadLicense(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Download License", nil)
}
