package middlewarex

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/constant"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/rara"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service"

	"go.opentelemetry.io/otel"
)

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

const (
	ErrorCodeAddressIsInvalid = 1005
)

func APIMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		address := c.Param("address")
		if address != "" {
			if address, err := ResolveAddress(c, address, false); err != nil {
				return c.JSON(http.StatusOK, &ErrorResponse{
					Error:     err.Error(),
					ErrorCode: ErrorCodeAddressIsInvalid,
				})
			} else {
				c.SetParamValues(address)
			}
		}

		apiKey := c.Request().Header.Get(constant.API_KEY_HEADER)
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
		apiKey := c.Request().Header.Get(constant.API_KEY_HEADER)
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
func ResolveAddress(c echo.Context, address string, ignoreContract bool) (string, error) {
	tracer := otel.Tracer("ResolveAddress")
	_, httpSnap := tracer.Start(c.Request().Context(), "middleware")
	defer httpSnap.End()

	result := name_service.ReverseResolveAll(c.Request().Context(), strings.ToLower(address), false)
	if len(result.Address) == 0 && result.Err != nil {
		return "", result.Err
	}

	// Get rara nft list cach from redis
	// TODO Need a standardized interface to check whether address is a nft
	cacheMap := make(map[string]struct{})
	_, err := cache.GetJson(c.Request().Context(), rara.MapKey, &cacheMap)

	if err == nil {
		if _, ok := cacheMap[strings.ToLower(result.Address)]; ok {
			return fmt.Sprintf("%s:%s", "nft", strings.ToLower(result.Address)), nil
		}
	}

	// check contract
	if !ignoreContract {
		isEthContract, _ := name_service.IsEthereumContract(c.Request().Context(), result.Address, protocol.NetworkEthereum)
		if isEthContract {
			return "", fmt.Errorf("%s", name_service.ErrNotSupportContract)
		}
	}

	return strings.ToLower(result.Address), nil
}
