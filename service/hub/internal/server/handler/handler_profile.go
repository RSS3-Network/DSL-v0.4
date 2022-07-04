package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	sModel "github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

func (h *Handler) GetProfileListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetProfileListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := sModel.GetRequest{}

	request.Address = c.Param("address")

	profileList, total, err := h.getProfileListDatabase(ctx, request)
	if err != nil {
		return err
	}

	if len(profileList) == 0 {
		return c.JSON(http.StatusOK, &Response{
			Result: make([]interface{}, 0),
		})
	}

	var cursor string
	if total > int64(DefaultLimit) {
		cursor = profileList[len(profileList)-1].TransactionHash
	}

	return c.JSON(http.StatusOK, &Response{
		Total:  total,
		Cursor: cursor,
		Result: profileList,
	})
}

// getTransferListDatabase get transfer data from database
func (h *Handler) getProfileListDatabase(c context.Context, request sModel.GetRequest) ([]dbModel.Transfer, int64, error) {
	tracer := otel.Tracer("getProfileListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]dbModel.Transfer, 0)
	total := int64(0)
	sql := h.DatabaseClient.Model(&dbModel.Transfer{}).Where("LOWER(address_from) = ?",
		strings.ToLower(request.Address),
	)

	if len(request.Cursor) > 0 {
		cursorInt, err := strconv.Atoi(request.Cursor)
		if err != nil {
			return nil, total, err
		}

		if cursorInt > 0 {
			sql = sql.Offset(cursorInt * DefaultLimit)
		}
	}

	sql = sql.Where("tag = ?", filter.TagSocial)
	sql = sql.Where("\"type\" IN ? ", []string{filter.SocialProfile})
	sql = sql.Limit(DefaultLimit)

	if err := sql.Find(&dbResult).Error; err != nil {
		return nil, total, err
	}

	return dbResult, total, nil
}
