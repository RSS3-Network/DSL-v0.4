package middlewarex

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

const (
	ErrorCodeBadParams = 1003
)

func PathUnescapeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		names := c.ParamNames()
		values := c.ParamValues()
		for i, name := range names {
			escaped, err := url.QueryUnescape(values[i])
			if err != nil {
				return c.JSON(http.StatusBadRequest, &ErrorResponse{
					Error:     err.Error(),
					ErrorCode: ErrorCodeBadParams,
				})
			}
			c.ParamNames()[i] = name
			c.ParamValues()[i] = escaped
		}
		return next(c)
	}
}
