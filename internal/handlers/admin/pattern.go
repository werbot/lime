package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/models"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// Patterns is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern [get]
func Patterns(c *fiber.Ctx) error {
	log := logging.New()

	pagination := webutil.GetPaginationFromCtx(c)
	patterns, err := queries.DB().Patterns(c.Context(), pagination)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "Patterns", patterns)
}

// Pattern is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern/:id [get]
func Pattern(c *fiber.Ctx) error {
	log := logging.New()

	pattern, err := queries.DB().Pattern(c.Context(), c.Params("id"))
	if err != nil {
		if err == errors.ErrPatternNotFound {
			return webutil.StatusNotFound(c, errors.MsgPatternNotFound)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}
	return webutil.StatusOK(c, "Customer info", pattern)
}

// AddPattern is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern [post]
func AddPattern(c *fiber.Ctx) error {
	log := logging.New()
	request := &models.Pattern{}

	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	err := queries.DB().AddPattern(c.Context(), request)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	// adit section
	metaAudit := webutil.GetRequestInfo(c, request)
	queries.DB().AddAudit(c.Context(), models.SectionPattern, "admin", models.OnAdd, metaAudit)

	return webutil.StatusOK(c, "Pattern added", nil)
}

// ClonePattern is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern/:id [post]
func ClonePattern(c *fiber.Ctx) error {
	log := logging.New()
	request := &models.Pattern{}

	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	if request.ID != c.Params("id") {
		return webutil.StatusBadRequest(c, nil)
	}

	pattern, err := queries.DB().ClonePattern(c.Context(), request)
	if err != nil {
		if err == errors.ErrPatternNotFound {
			return webutil.StatusNotFound(c, errors.MsgPatternNotFound)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	// adit section
	auditData := models.Metadata{
		"clone_id": request.ID,
		"pattern": models.Metadata{
			"id":   pattern.ID,
			"name": pattern.Name,
		},
	}
	metaAudit := webutil.GetRequestInfo(c, auditData)
	queries.DB().AddAudit(c.Context(), models.SectionPattern, "admin", models.OnClone, metaAudit)

	return webutil.StatusOK(c, "Clone pattern", pattern)
}

// UpdatePattern is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern/:id [patch]
func UpdatePattern(c *fiber.Ctx) error {
	log := logging.New()
	request := &models.Pattern{}

	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	err := queries.DB().UpdatePattern(c.Context(), request)
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
	request.Licenses = nil
	metaAudit := webutil.GetRequestInfo(c, request)
	queries.DB().AddAudit(c.Context(), models.SectionPattern, "admin", models.OnUpdate, metaAudit)

	return webutil.StatusOK(c, "Pattern updated", nil)
}

// Delete is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern/:id [patch]
func DeletePattern(c *fiber.Ctx) error {
	log := logging.New()

	err := queries.DB().DeletePattern(c.Context(), c.Params("id"))
	if err != nil {
		if err == errors.ErrPatternNotDeleted {
			return webutil.StatusBadRequest(c, errors.MsgPatternNotDeleted)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	// adit section
	auditData := models.Metadata{
		"id": c.Params("id"),
	}
	metaAudit := webutil.GetRequestInfo(c, auditData)
	queries.DB().AddAudit(c.Context(), models.SectionPattern, "admin", models.OnDelete, metaAudit)

	return webutil.StatusOK(c, "Pattern deleted", nil)
}
