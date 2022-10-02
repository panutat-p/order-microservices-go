package handler

import (
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
	return c.String(http.StatusOK, "OK")
}
