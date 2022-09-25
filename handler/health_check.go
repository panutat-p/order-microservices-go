package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary      Readiness Probe
// @Description  health check
// @Tags         health
// @Produce      plain
// @Success      200              {string}		string		"OK"
// @Router       /					[get]
func (h *Handler) healthCheck(c echo.Context) error {
	fmt.Println("🟦 OK")
	return c.String(http.StatusOK, "OK")
}
