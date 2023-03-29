package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

func (h *Handler) GetPlatformListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetPlatformListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetPlatformRequest{}
	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	request.PlatformType = strings.ToLower(request.PlatformType)

	result := make([]model.PlatformResult, 0)
	switch request.PlatformType {
	case "exchange":
		// get DEX from the list
		for _, v := range protocol.PlatformList[request.PlatformType] {
			result = append(result, model.PlatformResult{
				Name: v,
				Tag:  request.PlatformType,
				Type: "DEX",
			})
		}
		// get CEX from database
		dbResult, err := dao.GetCexPlatformList(ctx, request)
		if err != nil {
			logrus.Warnf("[Hub] /platform, GetPlatformListFunc: %s", err)
		} else {
			result = append(result, dbResult...)
		}
	case "all":
		for k, t := range protocol.PlatformList {
			for _, v := range t {
				temp := model.PlatformResult{
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

		dbResult, err := dao.GetCexPlatformList(ctx, request)
		if err != nil {
			logrus.Warnf("[Hub] /platform, GetPlatformListFunc: %s", err)
		} else {
			result = append(result, dbResult...)
		}
	default:
		for _, v := range protocol.PlatformList[request.PlatformType] {
			result = append(result, model.PlatformResult{
				Name: v,
				Tag:  request.PlatformType,
			})
		}
	}
	total := int64(len(result))

	return c.JSON(http.StatusOK, &model.Response{
		Total:  &total,
		Result: result,
	})
}
