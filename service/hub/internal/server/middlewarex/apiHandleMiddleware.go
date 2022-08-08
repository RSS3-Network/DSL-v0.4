package middlewarex

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler"
)

func APIMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		address := c.Param("address")
		if address != "" {
			if address, err := TranslateAddress(address); err != nil {
				return err
			} else {
				c.SetParamValues(address)
			}
		}

		err := CheckAPIKey(c)
		if err != nil {
			// return err
		}

		return next(c)
	}
}

func CheckAPIKeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := CheckAPIKey(c)
		if err != nil {
			// return err
		}

		return next(c)
	}
}

func CheckAPIKey(c echo.Context) error {
	var item model.APIKey
	var apiKey = c.Request().Header.Get("X-API-KEY")

	if len(apiKey) == 0 {
		return fmt.Errorf("miss X-API-KEY header")
	}

	if err := config.ConfigHub.DatabaseClient.
		Where("uuid = ?", apiKey).
		First(&item).Error; err != nil {
		return fmt.Errorf("X-API-KEY is invaild")
	}

	return nil
}

// TranslateAddress translate handles into an address
func TranslateAddress(address string) (string, error) {
	result := address

	splits := strings.Split(address, ".")

	if len(splits) < 2 {
		return strings.ToLower(result), nil
	}

	switch splits[len(splits)-1] {
	case "eth":
		nsResult, err := ens.Resolve(config.ConfigHub.RPC.General.Ethereum.HTTP, address)
		if err != nil {
			return "", fmt.Errorf("failed to resolve ENS address %s: %s", address, err)
		}

		result = nsResult
	case "csb":
		nsResult, err := handler.ResolveCrossbell(address)
		if err != nil {
			return "", fmt.Errorf("failed to resolve Crossbell address %s: %s", address, err)
		}
		result = nsResult
	case "lens":
		nsResult, err := handler.ResolveLens(address)
		if err != nil {
			return "", fmt.Errorf("failed to resolve Lens address %s: %s", address, err)
		}
		result = nsResult
	}

	return strings.ToLower(result), nil
}
