package dao

import (
	"context"
	"fmt"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

func GetAssets(c context.Context, request model.GetAssetRequest) ([]dbModel.Asset, int64, error) {
	tracer := otel.Tracer("getAssets")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	assetList := make([]dbModel.Asset, 0)
	total := int64(0)
	sql := database.Global().
		Model(&dbModel.Asset{}).
		Where("owner = ?", request.Address)

	if request.BlockSpam == nil || *request.BlockSpam {
		sql = sql.Where("token_address NOT IN ?", allowlist.SpamList.Keys())
	}

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.TokenAddress) > 0 {
		sql = sql.Where("token_address = ?", strings.ToLower(request.TokenAddress))
	}

	if len(request.TokenId) > 0 {
		sql = sql.Where("token_id = ?", request.TokenId)
	}

	if len(request.Cursor) > 0 {
		param := strings.Split(request.Cursor, ":")
		if len(param) < 3 {
			return nil, 0, fmt.Errorf("cursor is wrong")
		}

		var lastItem dbModel.Asset

		if err := database.Global().
			Where("network = ?", param[0]).
			Where("token_address = ?", param[1]).
			Where("token_id", param[2]).
			First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND token_id < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.TokenID)
	}

	if err := sql.Count(&total).Limit(request.Limit).Order("timestamp DESC, token_id DESC").Find(&assetList).Error; err != nil {
		return nil, 0, err
	}

	return assetList, total, nil
}
