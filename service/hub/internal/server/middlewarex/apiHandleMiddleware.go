package middlewarex

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func APIMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		address := c.Param("address")
		if address != "" {
			if address, err := ResolveAddress(address); err != nil {
				return c.JSON(http.StatusOK, &ErrorResponse{
					Error: err.Error(),
				})
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
	if len(result.Address) == 0 {
		return "", fmt.Errorf("The address provided is invalid. You can use a 0x, ENS, Crossbell, or Lens address.")
	}

	// check valid
	valid := name_service.IsValidAddress(result.Address)
	if !valid {
		return "", fmt.Errorf("The address provided is invalid. You can use a 0x, ENS, Crossbell, or Lens address.")
	}

	// check contract
	isContract, _ := name_service.IsEthereumContract(result.Address)
	if isContract {
		return "", fmt.Errorf("Contract addresses are not currently supported.")
	}

	return strings.ToLower(result.Address), nil
}
