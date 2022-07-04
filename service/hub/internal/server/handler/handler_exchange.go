package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"go.opentelemetry.io/otel"
)

type GetExchangeListRequest struct {
	Type    string   `param:"type"`
	Cursor  int      `param:"cursor"`
	Name    []string `query:"name"`
	Network []string `query:"network"`
}

var pageLimit = 1000

func (h *Handler) GetExchangeListFunc(c echo.Context) error {
	tracer := otel.Tracer("GetExchangeListFunc")
	ctx, httpSnap := tracer.Start(c.Request().Context(), "http")

	defer httpSnap.End()

	request := GetExchangeListRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}

	switch request.Type {
	case "cex":
		result, total, err := h.getCexListDatabase(ctx, request)
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return c.JSON(http.StatusOK, &Response{
				Result: make([]interface{}, 0),
			})
		}

		return c.JSON(http.StatusOK, &Response{
			Total:  total,
			Cursor: strconv.Itoa(request.Cursor),
			Result: result,
		})

	case "dex":
		result, total, err := h.getDexListDatabase(ctx, request)
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return c.JSON(http.StatusOK, &Response{
				Result: make([]interface{}, 0),
			})
		}

		return c.JSON(http.StatusOK, &Response{
			Total:  total,
			Cursor: strconv.Itoa(request.Cursor),
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

func (h *Handler) getCexListDatabase(c context.Context, request GetExchangeListRequest) ([]CexResult, int64, error) {
	tracer := otel.Tracer("getCexListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]model.CexWallet, 0)
	total := int64(0)
	sql := h.DatabaseClient.
		Model(&model.CexWallet{})

	if len(request.Network) > 0 {
		sql = sql.Where("network IN ?", request.Network)
	}

	if request.Cursor > 0 {
		sql = sql.Offset(request.Cursor * pageLimit)
	}

	if err := sql.Count(&total).Limit(pageLimit).Find(&dbResult).Error; err != nil {
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

func (h *Handler) getDexListDatabase(c context.Context, request GetExchangeListRequest) ([]DexResult, int64, error) {
	tracer := otel.Tracer("getCexListDatabase")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]model.SwapPool, pageLimit)
	total := int64(0)
	sql := h.DatabaseClient.
		Model(&model.SwapPool{})

	h.DatabaseClient.
		Model(&model.SwapPool{}).Count(&total)

	if len(request.Network) > 0 {
		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.Name) > 0 {
		for i, v := range request.Name {
			request.Name[i] = strings.ToLower(v)
		}
		sql = sql.Where("LOWER(source) IN ?", request.Name)
	}

	if request.Cursor > 0 {
		offset := request.Cursor * pageLimit
		total -= int64(offset)
		sql = sql.Offset(offset)
	}

	if err := sql.Limit(pageLimit).Find(&dbResult).Error; err != nil {
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
