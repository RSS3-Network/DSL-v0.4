package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"

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

	request := GetExchangeRequest{}
	if err := c.Bind(&request); err != nil {
		return InternalError(c)
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	var cursor string
	switch strings.ToLower(request.ExchangeType) {
	case "cex":
		result, total, err := h.getCexListDatabase(ctx, request)
		if err != nil {
			return InternalError(c)
		}
		if len(result) == 0 {
			return c.JSON(http.StatusOK, &Response{
				Result: make([]interface{}, 0),
			})
		}
		if total > int64(DefaultLimit) {
			cursor = strconv.Itoa(request.Cursor + 1)
		}

		return c.JSON(http.StatusOK, &Response{
			Total:  total,
			Cursor: cursor,
			Result: result,
		})

	case "dex":
		result, total, err := h.getDexListDatabase(ctx, request)
		if err != nil {
			return InternalError(c)
		}
		if len(result) == 0 {
			return c.JSON(http.StatusOK, &Response{
				Result: make([]interface{}, 0),
			})
		}
		if total > int64(DefaultLimit) {
			cursor = strconv.Itoa(request.Cursor + 1)
		}

		return c.JSON(http.StatusOK, &Response{
			Total:  total,
			Cursor: cursor,
			Result: result,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"status":  "error",
		"message": "Invalid exchange type, must be cex or dex",
	})
}

type CexResult struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Network string `json:"network"`
}

func (h *Handler) getCexListDatabase(c context.Context, request GetExchangeRequest) ([]CexResult, int64, error) {
	tracer := otel.Tracer("getCexListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]exchange.CexWallet, 0)
	total := int64(0)
	sql := h.DatabaseClient.
		Model(&exchange.CexWallet{})

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}

	if request.Cursor > 0 {
		sql = sql.Offset(request.Cursor * DefaultLimit)
	}

	if err := sql.Count(&total).Limit(DefaultLimit).Find(&dbResult).Error; err != nil {
		return nil, 0, err
	}

	result := make([]CexResult, len(dbResult))
	for i, v := range dbResult {
		result[i].Name = v.Name
		result[i].Address = v.WalletAddress
		result[i].Network = v.Network
	}

	return result, total, nil
}

type DexResult struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Network string `json:"network"`
	Pair    string `json:"pair"`
}

func (h *Handler) getDexListDatabase(c context.Context, request GetExchangeRequest) ([]DexResult, int64, error) {
	tracer := otel.Tracer("getDexListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]exchange.SwapPool, DefaultLimit)
	total := int64(0)
	sql := h.DatabaseClient.
		Model(&exchange.SwapPool{})

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.Name) > 0 {
		for i, v := range request.Name {
			request.Name[i] = strings.ToLower(v)
		}
		sql = sql.Where("LOWER(source) IN ?", request.Name)
	}

	sql.Count(&total)

	if request.Cursor > 0 {
		offset := request.Cursor * DefaultLimit

		if int64(offset) < total {
			total -= int64(offset)
		}
		sql = sql.Offset(offset)
	}

	if err := sql.Limit(DefaultLimit).Find(&dbResult).Error; err != nil {
		return nil, 0, err
	}

	result := make([]DexResult, len(dbResult))
	for i, v := range dbResult {
		result[i].Name = v.Source
		result[i].Address = v.ContractAddress
		result[i].Network = v.Network
		result[i].Pair = v.Token0 + "-" + v.Token1
	}

	return result, total, nil
}
