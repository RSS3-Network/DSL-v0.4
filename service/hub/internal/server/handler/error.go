package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BadRequest(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
}
