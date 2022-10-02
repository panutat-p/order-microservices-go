package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Req model info
// @Description Request body
type Req struct {
	Id  string `json:"id"`
	Msg string `json:"msg"`
}

// @Summary      reflectReq Route
// @Description  reflect back the request body
// @Tags         reflect
// @Accept       json
// @Produce      json
// @Param        request		body		Req		true	"Request body"
// @Success      200  {object}  Req
// @Router       /reflect					[post]
func (h *Handler) reflectReq(c echo.Context) error {
	var req Req

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, req)
}
