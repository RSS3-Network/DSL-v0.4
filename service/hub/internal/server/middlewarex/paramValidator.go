package middlewarex

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ValidateParamsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Param("address") == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "address is required")
		}

		if c.QueryParam("tag") == "" && c.QueryParam("type") != "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Tag is required when Type is present")
		}

		return next(c)
	}
}
