package middlewarex

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service"
)

func APIMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		address := c.Param("address")
		if address != "" {
			if address, err := ResolveAddress(address); err != nil {
				return err
			} else {
				c.SetParamValues(address)
			}
		}

		apiKey := c.Request().Header.Get("X-API-KEY")
		c.Set("API-KEY", apiKey)

		_ = CheckAPIKey(apiKey)

		// err := CheckAPIKey(apiKey)
		// if err != nil {
		// 	return err
		// }

		return next(c)
	}
}

func CheckAPIKeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-KEY")
		c.Set("API-KEY", apiKey)

		_ = CheckAPIKey(apiKey)
		// err := CheckAPIKey(apiKey)
		// if err != nil {
		// 	return err
		// }

		return next(c)
	}
}

func CheckAPIKey(apiKey string) error {
	var item model.APIKey

	if len(apiKey) == 0 {
		return fmt.Errorf("miss X-API-KEY header")
	}

	if err := database.Global().
		Where("uuid = ?", apiKey).
		First(&item).Error; err != nil {
		return fmt.Errorf("X-API-KEY is invaild")
	}

	return nil
}

// ResolveAddress resolve handles into an address
func ResolveAddress(address string) (string, error) {
	result := name_service.ReverseResolveAll(strings.ToLower(address), false)

	return strings.ToLower(result.Address), nil
}
