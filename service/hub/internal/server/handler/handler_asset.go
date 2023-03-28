package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

var cursorKey = "%v:%v:%v"

// GetAssetsFunc HTTP handler for asset API
func (h *Handler) GetAssetsFunc(c echo.Context) error {
	tracer := otel.Tracer("GetAssetsFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetAssetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return ValidateFailed(c)
	}

	assetList, total, err := h.service.GetAssets(ctx, request)
	if err != nil {
		return ErrorResp(c, err, http.StatusInternalServerError, ErrorCodeInternalError)
	}

	if len(assetList) == 0 {
		return c.JSON(http.StatusOK, &model.Response{
			Result: make([]dbModel.Asset, 0),
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		last := assetList[len(assetList)-1]
		cursor = fmt.Sprintf(cursorKey, last.Network, last.TokenAddress, last.TokenID)
	}

	return c.JSON(http.StatusOK, &model.Response{
		Total:  &total,
		Cursor: cursor,
		Result: assetList,
	})
}
