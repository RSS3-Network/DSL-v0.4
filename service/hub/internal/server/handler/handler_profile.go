package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"go.opentelemetry.io/otel"
)

// GetProfilesFunc supported filter:
// - address
// - network
// - platform
func (h *Handler) GetProfilesFunc(c echo.Context) error {
	go h.apiReport(GetProfiles, c)
	tracer := otel.Tracer("GetProfilesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	go h.filterReport(GetProfiles, request)

	profileList, err := h.getProfiles(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	if len(profileList) == 0 || request.Refresh {
		// refresh profile
		h.initializeIndexerStatus(ctx, request.Address)
		go h.publishIndexerMessage(ctx, protocol.Message{Address: request.Address, IgnoreNote: true})
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  int64(len(profileList)),
		Result: profileList,
	})
}

func (h *Handler) BatchGetProfilesFunc(c echo.Context) error {
	go h.apiReport(PostProfiles, c)
	tracer := otel.Tracer("BatchGetProfilesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := BatchGetProfilesRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	go h.filterReport(PostProfiles, request)

	if len(request.Address) > DefaultLimit {
		request.Address = request.Address[:DefaultLimit]
	}

	profileList, err := h.batchGetProfiles(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	if len(profileList) == 0 || request.Refresh {
		// publish mq message
		go func() {
			for _, address := range request.Address {
				h.initializeIndexerStatus(ctx, address)
				h.publishIndexerMessage(ctx, protocol.Message{Address: address, IgnoreNote: true})
				time.Sleep(500 * time.Millisecond)
			}
		}()
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  int64(len(profileList)),
		Result: profileList,
	})
}

// getProfiles get profile data from database
func (h *Handler) getProfiles(c context.Context, request GetRequest) ([]dbModel.Profile, error) {
	tracer := otel.Tracer("getProfiles")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]dbModel.Profile, 0)

	sql := database.Global().Model(&dbModel.Profile{}).Where("LOWER(address) = ? ",
		strings.ToLower(request.Address),
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

// batchGetProfiles get profile data from database
func (h *Handler) batchGetProfiles(c context.Context, request BatchGetProfilesRequest) ([]dbModel.Profile, error) {
	tracer := otel.Tracer("batchGetProfiles")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]dbModel.Profile, 0)

	if len(request.Address) > DefaultLimit {
		request.Address = request.Address[:DefaultLimit]
	}

	for i, v := range request.Address {
		request.Address[i] = strings.ToLower(v)
	}

	sql := database.Global().Model(&dbModel.Profile{}).Where("LOWER(address) IN ? ", request.Address)

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

	if err := sql.Find(&dbResult).Error; err != nil {
		return nil, err
	}

	return dbResult, nil
}
