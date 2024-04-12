package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// SettingGroup is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/setting/:group [get]
func SettingGroup(c *fiber.Ctx) error {
	log := logging.New()
	settingGroup := c.Params("group")

	var section any
	var err error

	switch settingGroup {
	case "site":
		section, err = queries.DB().GetSettingByGroup(c.Context(), &models.Site{})
	case "mail":
		section, err = queries.DB().GetSettingByGroup(c.Context(), &models.Mail{})
	default:
		section, err = queries.DB().GetSettingByKey(c.Context(), settingGroup)
	}

	if err != nil {
		if err == errors.ErrSettingNotFound {
			return webutil.StatusNotFound(c, errors.MsgSettingNotFound)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "Setting group", section)
}

// UpdateSettingGroup is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/setting/:group [patch]
func UpdateSettingGroup(c *fiber.Ctx) error {
	log := logging.New()
	settingGroup := c.Params("group")

	var request any

	switch settingGroup {
	case "site":
		request = &models.Site{}
	case "mail":
		request = &models.Mail{}
	default:
		request = &models.SettingName{}
	}

	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	if _, ok := request.(*models.SettingName); ok {
		_request := request.(*models.SettingName)
		_request.Key = settingGroup
		if err := queries.DB().UpdateSettingByKey(c.Context(), _request); err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c, nil)
		}
		return webutil.Response(c, fiber.StatusOK, "Setting key updated", nil)
	}

	if err := queries.DB().UpdateSettingByGroup(c.Context(), request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	// adit section
	metaAudit := webutil.GetRequestInfo(c, request)
	queries.DB().AddAudit(c.Context(), models.SectionSetting, "admin", models.OnUpdate, metaAudit)

	return webutil.StatusOK(c, "Setting group updated", nil)
}
