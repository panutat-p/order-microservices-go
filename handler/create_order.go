package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/panutat-p/order-microservices-go/core/domain"
)

// @Summary      createOrder	Route
// @Description  create a new order
// @Tags         reflect
// @Accept       json
// @Produce      json
// @Param        request		body		domain.Order	true	"Order Info"
// @Success      200  			{object}  Req
// @Router       /create		[post]
func (h *Handler) createOrder(c echo.Context) error {
	var req domain.Order

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.OrderService.Create(&req)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"msg": "success"})
}
