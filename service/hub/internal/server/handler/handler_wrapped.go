package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

func (h *Handler) GetWrappedFunc(c echo.Context) error {
	tracer := otel.Tracer("GetWrappedFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return ValidateFailed(c)
	}

	request.Address = name_service.ReverseResolveAll(c.Request().Context(), request.Address, false).Address

	if len(request.Address) == 0 {
		return AddressIsInvalid(c)
	}

	result := model.WrappedResult{}

	err := h.service.GetWrapped(ctx, request, &result)
	if err != nil {
		return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeInternalError)
	}

	return c.JSON(http.StatusOK, &model.Response{
		Result: result,
	})
}
