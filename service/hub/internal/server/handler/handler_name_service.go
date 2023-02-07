package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

func (h *Handler) GetNameResolveFunc(c echo.Context) error {
	go h.apiReport(model.GetNS, c)
	tracer := otel.Tracer("GetNameResolveFunc")
	_, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	if len(request.Address) == 0 {
		return AddressIsEmpty(c)
	}

	go h.filterReport(model.GetNS, request, c)

	result := name_service.ReverseResolveAll(c.Request().Context(), strings.ToLower(request.Address), true)

	if len(result.Address) == 0 {
		return AddressIsInvalid(c)
	}
	return c.JSON(http.StatusOK, result)
}
