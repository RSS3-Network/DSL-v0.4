package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"

	"go.opentelemetry.io/otel"
)

func (h *Handler) GetMastodonFunc(c echo.Context) error {
	go h.apiReport(model.GetMastodon, c)
	tracer := otel.Tracer("GetMastodonFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return ValidateFailed(c)
	}

	go h.filterReport(model.GetMastodon, request, c)

	contentList, err := h.service.GetMastodonContent(ctx, request)
	if err != nil {
		return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeInternalError)
	}

	total := int64(len(contentList))
	return c.JSON(http.StatusOK, &model.Response{
		Total:  &total,
		Result: contentList,
	})
}
