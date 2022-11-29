package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
		return err
	}

	result := model.WrappedResult{}

	err := h.service.GetWrapped(ctx, request, &result)
	if err != nil {
		return InternalError(c)
	}

	return c.JSON(http.StatusOK, &model.Response{
		Result: result,
	})
}
