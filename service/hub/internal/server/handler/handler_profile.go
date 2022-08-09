package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"go.opentelemetry.io/otel"
)

// GetProfileListFunc supported filter:
// - address
// - network
// - platform
func (h *Handler) GetProfileListFunc(c echo.Context) error {
	go APIReport(GetProfiles, c.Get("API-KEY").(string))
	tracer := otel.Tracer("GetProfileListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	go FilterReport(GetProfiles, request)

	profileList, err := h.getProfileListDatabase(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	if len(profileList) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]dbModel.Transfer, 0),
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  int64(len(profileList)),
		Result: profileList,
	})
}

// getTransferListDatabase get transfer data from database
func (h *Handler) getProfileListDatabase(c context.Context, request GetRequest) ([]dbModel.Profile, error) {
	tracer := otel.Tracer("getProfileListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]dbModel.Profile, 0)

	sql := h.DatabaseClient.Model(&dbModel.Profile{}).Where("LOWER(address) = ? ",
		request.Address,
	)

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.Platform) > 0 {
		for i, v := range request.Platform {
			request.Platform[i] = strings.ToLower(v)
		}
		sql = sql.Where("platform IN ?", request.Platform)
	}

	sql = sql.Limit(DefaultLimit)

	if err := sql.Find(&dbResult).Error; err != nil {
		return nil, err
	}

	return dbResult, nil
}
