package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/panutat-p/order-microservices-go/core/domain"
)

type ListOrdersRes struct {
	Data []domain.Order `json:"data"`
}

// @Summary      listOrders Route
// @Description  reflect back the request body
// @Tags         list
// @Accept       json
// @Produce      json
// @Param        type   	query       string  true    "Order Type"
// @Success      200  		{object}  	ListOrdersRes
// @Router       /list		[get]
func (h *Handler) listOrders(c echo.Context) error {
	orderType := c.QueryParam("type")
	fmt.Println("🟨 orderType:", orderType)
	orders, err := h.OrderService.List(orderType)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, ListOrdersRes{
		Data: orders,
	})
}
