package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// Customers is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/customer [get]]
func Customers(c *fiber.Ctx) error {
	log := logging.New()

	pagination := webutil.GetPaginationFromCtx(c)
	customers, err := queries.DB().Customers(c.Context(), pagination)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "Customers", customers)
}

// Customer is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/customer/:id [get]
func Customer(c *fiber.Ctx) error {
	log := logging.New()

	customer, err := queries.DB().Customer(c.Context(), c.Params("id"))
	if err != nil {
		if err == errors.ErrCustomerNotFound {
			return webutil.StatusNotFound(c, errors.MsgCustomerNotFound)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}
	return webutil.StatusOK(c, "Customer info", customer)
}

// AddCustomer is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/customer [post]
func AddCustomer(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Create customer", nil)
}

// UpdateCustomer is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/customer/:id [patch]
func UpdateCustomer(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Update customer", nil)
}
