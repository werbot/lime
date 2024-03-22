package handlers

import (
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/werbot/lime/internal/config"
	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/license"
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

// NewLicense is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/license [post]
func NewLicense(c *fiber.Ctx) error {
	cfg := config.Data()

	month := time.Hour * 24 * 31

	licenseInfo := &license.License{
		IssuedBy:     "customer.Name",
		CustomerID:   "subscription.StripeID",
		SubscriberID: 123,
		Type:         "tariff.Name",
		Limit: []license.Limits{
			{
				Key:   "key1",
				Value: "value1",
			},
			{
				Key:   "key2",
				Value: "value2",
			},
		},
		Metadata:  []byte(`{"message": "test message"}`),
		ExpiresAt: time.Now().UTC().Add(month),
		IssuedAt:  time.Now().UTC(),
	}

	privKey := fsutil.MustReadFile(filepath.Join(cfg.Keys.KeyDir, cfg.Keys.License.PrivateKey))
	licenseKey := license.DecodePrivateKey(privKey)

	encoded, err := licenseInfo.Encode(licenseKey)
	if err != nil {
		return webutil.StatusNotFound(c, utils.StatusMessage(fiber.StatusNotFound))
	}

	return webutil.StatusOK(c, "Create license", string(encoded))
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
