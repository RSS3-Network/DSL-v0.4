package middlewarex

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ZapLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fields := make([]zap.Field, 0)

			if err := next(c); err != nil {
				c.Error(err)

				fields = append(fields, zap.Error(err))
			}

			status := c.Response().Status

			fields = append(
				fields,
				zap.Int("status", c.Response().Status),
				zap.String("method", c.Request().Method),
				zap.String("uri", c.Request().RequestURI),
				zap.String("user_agent", c.Request().UserAgent()),
				zap.String("client_ip", c.RealIP()),
			)

			switch {
			case status >= http.StatusInternalServerError:
				logger.Error("request returned an error", fields...)
			case status >= http.StatusBadRequest:
				logger.Warn("request is invalid", fields...)
			case status >= http.StatusMultipleChoices:
				logger.Info("request was redirected", fields...)
			default:
				logger.Info("request was processed", fields...)
			}

			return nil
		}
	}
}
