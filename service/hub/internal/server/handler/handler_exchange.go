package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
)

// GetExchangeListFunc supported filter:
// - network
// - name: string, the exchange name
// - exchange_type: string, `cex` or `dex`
func (h *Handler) GetExchangeListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetExchangeListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetExchangeRequest{}
	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return ValidateFailed(c)
	}

	var cursor string
	switch strings.ToLower(request.ExchangeType) {
	case "cex":
		result, total, err := dao.GetCexList(ctx, request)
		if err != nil {
			return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeInternalError)
		}
		if len(result) == 0 {
			return c.JSON(http.StatusOK, &model.Response{
				Result: make([]interface{}, 0),
			})
		}
		if total > int64(model.DefaultLimit) {
			cursor = strconv.Itoa(request.Cursor + 1)
		}

		return c.JSON(http.StatusOK, &model.Response{
			Total:  &total,
			Cursor: cursor,
			Result: result,
		})

	case "dex":
		result, total, err := dao.GetDexList(ctx, request)
		if err != nil {
			return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeInternalError)
		}
		if len(result) == 0 {
			return c.JSON(http.StatusOK, &model.Response{
				Result: make([]interface{}, 0),
			})
		}
		if total > int64(model.DefaultLimit) {
			cursor = strconv.Itoa(request.Cursor + 1)
		}

		return c.JSON(http.StatusOK, &model.Response{
			Total:  &total,
			Cursor: cursor,
			Result: result,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"status":     "error",
		"message":    "Invalid exchange type, must be cex or dex",
		"error_code": strconv.Itoa(ErrorCodeInvalidExchangeType),
	})
}
