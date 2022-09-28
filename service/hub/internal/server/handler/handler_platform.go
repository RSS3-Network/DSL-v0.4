package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

// NOTE: if you update the list, also update the list in the following files:
// - common/protocol/platform.go

var platformList = map[string][]string{
	filter.TagSocial: {
		protocol.PlatformMirror,
		// Lens support is not yet available
		// protocol.PlatformLens,
		protocol.PlatformCrossbell,
	},
	filter.TagCollectible: {
		protocol.PlatformPOAP,
		protocol.PlatformGalaxy,
	},
	filter.TagDonation: {
		protocol.PlatformGitcoin,
	},
	filter.TagGovernance: {
		protocol.PlatformSnapshot,
	},
	filter.TagExchange: {
		protocol.PlatformUniswap,
		protocol.PlatformSushiswap,
		protocol.PlatformPancakeswap,
		protocol.Platform1inch,
		protocol.PlatformMetamask,
	},
}

type PlatformResult struct {
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Type    string `json:"type,omitempty"`
	Network string `json:"network,omitempty"`
}

func (h *Handler) GetPlatformListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetPlatformListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetPlatformRequest{}
	if err := c.Bind(&request); err != nil {
		return InternalError(c)
	}

	request.PlatformType = strings.ToLower(request.PlatformType)

	result := make([]PlatformResult, 0)
	switch request.PlatformType {
	case "exchange":
		// get DEX from the list
		for _, v := range platformList[request.PlatformType] {
			result = append(result, PlatformResult{
				Name: v,
				Tag:  request.PlatformType,
				Type: "DEX",
			})
		}
		// get CEX from database
		dbResult, err := h.getCexPlatformListDatabase(ctx, request)
		if err != nil {
			logrus.Warnf("[Hub] /platform, GetPlatformListFunc: %s", err)
		} else {
			result = append(result, dbResult...)
		}
	case "all":
		for k, t := range platformList {
			for _, v := range t {
				temp := PlatformResult{
					Name: v,
					Tag:  k,
				}
				if k == filter.TagExchange {
					temp.Tag = filter.TagExchange
					temp.Type = "DEX"
				}
				result = append(result, temp)
			}
		}

		dbResult, err := h.getCexPlatformListDatabase(ctx, request)
		if err != nil {
			logrus.Warnf("[Hub] /platform, GetPlatformListFunc: %s", err)
		} else {
			result = append(result, dbResult...)
		}
	default:
		for _, v := range platformList[request.PlatformType] {
			result = append(result, PlatformResult{
				Name: v,
				Tag:  request.PlatformType,
			})
		}
	}
	total := int64(len(result))

	return c.JSON(http.StatusOK, &Response{
		Total:  &total,
		Result: result,
	})
}

func (h *Handler) getCexPlatformListDatabase(c context.Context, request GetPlatformRequest) ([]PlatformResult, error) {
	tracer := otel.Tracer("getCexListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	result := make([]PlatformResult, 0)
	sql := database.Global().
		Model(&exchange.CexWallet{})

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}
	columns := []string{"name", "network"}
	if err := sql.Limit(DefaultLimit).Distinct(columns).Select(columns).Find(&result).Error; err != nil {
		return nil, err
	}

	for i := range result {
		result[i].Tag = filter.TagExchange
		result[i].Type = "CEX"
	}

	return result, nil
}
