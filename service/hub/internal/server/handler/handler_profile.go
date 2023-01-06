package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

// GetProfilesFunc supported filter:
// - address
// - network
// - platform
func (h *Handler) GetProfilesFunc(c echo.Context) error {
	go h.apiReport(model.GetProfiles, c)
	tracer := otel.Tracer("GetProfilesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	go h.filterReport(model.GetProfiles, request, c)

	profileList, err := h.service.GetProfiles(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	total := int64(len(profileList))
	return c.JSON(http.StatusOK, &model.Response{
		Total:  &total,
		Result: profileList,
	})
}

func (h *Handler) BatchGetProfilesFunc(c echo.Context) error {
	go h.apiReport(model.PostProfiles, c)
	tracer := otel.Tracer("BatchGetProfilesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.BatchGetProfilesRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	go h.filterReport(model.PostProfiles, request, c)

	if len(request.Address) > model.DefaultLimit {
		request.Address = request.Address[:model.DefaultLimit]
	}

	profileList, err := h.service.BatchGetProfiles(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	total := int64(len(profileList))
	return c.JSON(http.StatusOK, &model.Response{
		Total:  &total,
		Result: profileList,
	})
}
