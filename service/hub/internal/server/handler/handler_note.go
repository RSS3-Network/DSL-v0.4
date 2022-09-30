package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

// GetNotesFunc HTTP handler for action API
// parse query parameters, query and assemble data
func (h *Handler) GetNotesFunc(c echo.Context) error {
	go h.apiReport(model.GetNotes, c)
	tracer := otel.Tracer("GetNotesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := model.GetRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	// api report
	if len(request.Cursor) == 0 {
		go h.filterReport(model.GetNotes, request)
	}

	transactions, total, err := h.service.GetNotes(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	if request.CountOnly {
		return c.JSON(http.StatusOK, &model.Response{
			Total: &total,
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactions[len(transactions)-1].Hash
	}

	var addressStatus []dbModel.Address
	if request.QueryStatus {
		addressStatus, _ = h.getAddress(ctx, []string{request.Address})
	}

	return c.JSON(http.StatusOK, &model.Response{
		Total:         &total,
		Cursor:        cursor,
		Result:        transactions,
		AddressStatus: addressStatus,
	})
}

// BatchGetNotesFunc query multiple addresses and filters
func (h *Handler) BatchGetNotesFunc(c echo.Context) error {
	go h.apiReport(model.PostNotes, c)
	tracer := otel.Tracer("BatchGetNotesFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "BatchGetNotesFunc:http")

	defer httpSnap.End()

	request := model.BatchGetNotesRequest{}

	if err := c.Bind(&request); err != nil {
		return BadRequest(c)
	}

	if len(request.Cursor) == 0 {
		go h.filterReport(model.PostNotes, request)
	}

	if len(request.Address) == 0 {
		return AddressIsEmpty(c)
	}

	transactions, total, err := h.service.BatchGetNotes(ctx, request)
	if err != nil {
		return InternalError(c)
	}

	var addressStatus []dbModel.Address
	if request.QueryStatus {
		addressStatus, _ = h.getAddress(ctx, request.Address)
	}

	if request.CountOnly {
		return c.JSON(http.StatusOK, &model.Response{
			Total: &total,
		})
	}

	var cursor string
	if total > int64(request.Limit) {
		cursor = transactions[len(transactions)-1].Hash
	}

	if total == 0 {
		return c.JSON(http.StatusOK, &model.Response{
			Result:        make([]dbModel.Transaction, 0),
			AddressStatus: addressStatus,
		})
	}

	return c.JSON(http.StatusOK, &model.Response{
		Total:         &total,
		Cursor:        cursor,
		Result:        transactions,
		AddressStatus: addressStatus,
	})
}

// getAddress
func (h *Handler) getAddress(ctx context.Context, address []string) ([]dbModel.Address, error) {
	tracer := otel.Tracer("getAddress")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	addressStatus := make([]dbModel.Address, 0)

	sql := database.Global().Model(&dbModel.Address{})

	if err := sql.
		Where("address IN ?", address).
		Find(&addressStatus).Error; err != nil {
		return nil, err
	}

	return addressStatus, nil
}
