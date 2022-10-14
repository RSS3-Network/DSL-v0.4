package dao

import (
	"context"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

func GetCexList(c context.Context, request model.GetExchangeRequest) ([]model.CexResult, int64, error) {
	tracer := otel.Tracer("dao.GetCexList")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]exchange.CexWallet, 0)
	total := int64(0)
	sql := database.Global().
		Model(&exchange.CexWallet{})

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}

	if request.Cursor > 0 {
		sql = sql.Offset(request.Cursor * model.DefaultLimit)
	}

	if err := sql.Count(&total).Limit(model.DefaultLimit).Find(&dbResult).Error; err != nil {
		return nil, 0, err
	}

	result := make([]model.CexResult, len(dbResult))
	for i, v := range dbResult {
		result[i].Name = v.Name
		result[i].Address = v.WalletAddress
		result[i].Network = v.Network
	}

	return result, total, nil
}

func GetDexList(c context.Context, request model.GetExchangeRequest) ([]model.DexResult, int64, error) {
	tracer := otel.Tracer("dao.GetDexList")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	dbResult := make([]exchange.SwapPool, model.DefaultLimit)
	total := int64(0)
	sql := database.Global().
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
		offset := request.Cursor * model.DefaultLimit

		if int64(offset) < total {
			total -= int64(offset)
		}
		sql = sql.Offset(offset)
	}

	if err := sql.Limit(model.DefaultLimit).Find(&dbResult).Error; err != nil {
		return nil, 0, err
	}

	result := make([]model.DexResult, len(dbResult))
	for i, v := range dbResult {
		result[i].Name = v.Source
		result[i].Address = v.ContractAddress
		result[i].Network = v.Network
		result[i].Pair = v.Token0 + "-" + v.Token1
	}

	return result, total, nil
}

func GetCexPlatformList(c context.Context, request model.GetPlatformRequest) ([]model.PlatformResult, error) {
	tracer := otel.Tracer("dao.GetCexPlatformList")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	result := make([]model.PlatformResult, 0)
	sql := database.Global().
		Model(&exchange.CexWallet{})

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}
	columns := []string{"name", "network"}
	if err := sql.Limit(model.DefaultLimit).Distinct(columns).Select(columns).Find(&result).Error; err != nil {
		return nil, err
	}

	for i := range result {
		result[i].Tag = filter.TagExchange
		result[i].Type = "CEX"
	}

	return result, nil
}
