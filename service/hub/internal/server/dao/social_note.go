package dao

import (
	"context"
	"strings"

	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

func BatchGetSocialTransactions(ctx context.Context, request model.BatchGetSocialNotesRequest) ([]dbModel.Transaction, int64, error) {
	tracer := otel.Tracer("batchGetSocialTransactions")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	transactions := make([]dbModel.Transaction, 0)
	total := int64(0)

	sql := database.Social().
		Model(&dbModel.Transaction{}).
		Where("owner IN ?", request.Address).
		Where("success IS TRUE") // Hide failed transactions

	if len(request.Cursor) > 0 {
		var lastItem dbModel.Transaction

		// no need to lowercase
		if err := database.Social().Where("hash = ?", request.Cursor).First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
	}

	sql = sql.Where("tag = ?", "social")
	if len(request.Type) > 0 {
		// type was already converted to lowercase
		sql = sql.Where("\"type\" IN ?", request.Type)
	}

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}

		sql = sql.Where("network IN ?", request.Network)
	}

	if len(request.Platform) > 0 {
		for i, v := range request.Platform {
			request.Platform[i] = strings.ToLower(v)
		}
		sql = sql.Where("LOWER(platform) IN ?", request.Platform)
	}

	if request.Timestamp.Unix() > 0 {
		sql = sql.Where("timestamp > ?", request.Timestamp)
	}

	if err := sql.Count(&total).Limit(request.Limit).Offset(request.Page * request.Limit).Order("timestamp DESC, index DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

func GetSocialTransfers(c context.Context, transactionHashes []string) ([]dbModel.Transfer, error) {
	tracer := otel.Tracer("GetSocialTransfers")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transfers := make([]dbModel.Transfer, 0)

	sql := database.Social().Model(&dbModel.Transfer{})

	if err := sql.
		Where("transaction_hash IN (SELECT * FROM UNNEST(?::TEXT[]))", pq.Array(transactionHashes)).
		Find(&transfers).Error; err != nil {
		return nil, err
	}

	return transfers, nil
}
