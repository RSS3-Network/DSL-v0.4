package middlewarex

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
)

func GetParamMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		address := c.Param("address")
		if address != "" {
			if address, err := HandleAddress(address); err != nil {
				return err
			} else {
				c.SetParamValues(address)
			}
		}

		return next(c)
	}
}

func HandleAddress(address string) (string, error) {
	result := address

	if strings.HasSuffix(address, ".eth") {
		nsResult, err := ens.Resolve(config.ConfigHub.RPC.General.Ethereum.HTTP, address)
		if err != nil {
			return "", err
		}

		result = nsResult.Address
	}
	return strings.ToLower(result), nil
}
