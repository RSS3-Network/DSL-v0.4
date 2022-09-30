package dao

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

func (d *Dao) InitializeAddressStatus(ctx context.Context, address string) {
	status := dbModel.Address{
		Address:          address,
		IndexingNetworks: protocol.SupportNetworks,
		Status:           false,
	}

	if err := database.Global().
		Model(dbModel.Address{}).
		Clauses(clause.OnConflict{
			UpdateAll: true,
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).
		Create(map[string]interface{}{
			"address":           status.Address,
			"status":            status.Status,
			"done_networks":     status.DoneNetworks,
			"indexing_networks": status.IndexingNetworks,
		}).Error; err != nil {
		loggerx.Global().Error("failed to upsert address", zap.Error(err), zap.String("address", address))
		return
	}
}
