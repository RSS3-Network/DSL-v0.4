package service

import (
	"context"
	"strings"

	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.uber.org/zap"

	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
)

func (s *Service) GetAssets(ctx context.Context, request model.GetAssetRequest) ([]dbModel.Asset, int64, error) {
	if request.Limit <= 0 || request.Limit > model.DefaultLimit {
		request.Limit = model.DefaultLimit
	}

	request.Address = strings.ToLower(request.Address)

	assetList, total, err := dao.GetAssets(ctx, request)
	if err != nil {
		loggerx.Global().Error("GetAssets: query database err", zap.Error(err))
		return nil, 0, err
	}

	// publish mq message
	if request.Refresh || len(assetList) == 0 {
		go s.PublishIndexerAssetMessage(ctx, request.Address)
	}

	return assetList, total, nil
}
