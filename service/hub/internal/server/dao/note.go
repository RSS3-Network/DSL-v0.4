package dao

import (
	"context"
	"strings"

	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/database"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"go.opentelemetry.io/otel"
)

// getTransactions get transaction data from database
func GetTransactions(ctx context.Context, request model.GetRequest) ([]dbModel.Transaction, int64, error) {
	tracer := otel.Tracer("getTransactions")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	transactions := make([]dbModel.Transaction, 0)
	total := int64(0)
	sql := database.Global().
		WithContext(ctx).
		Model(&dbModel.Transaction{}).
		Where("success IS TRUE") // Hide failed transactions

	if len(request.TokenId) == 0 {
		sql.Where("owner = ?", request.Address) // address was already converted to lowercase
	}

	if len(request.Cursor) > 0 {
		var lastItem dbModel.Transaction

		// no need to lowercase
		if err := database.Global().Where("hash = ?", request.Cursor).First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
	}

	if len(request.Hash) > 0 {
		sql = sql.Where("hash = ?", request.Hash) // no need to lowercase
	}

	if len(request.Tag) > 0 {
		sql = sql.Where("tag IN ?", request.Tag)
		if len(request.Type) > 0 {
			// type was already converted to lowercase
			sql = sql.Where("\"type\" IN ?", request.Type)
		}
	}

	if !request.IncludePoap {
		sql = sql.Where("\"type\" != ?", filter.CollectiblePoap)
	}

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}

		sql = sql.Where("LOWER(network) IN ?", request.Network)
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

	if len(request.HashList) > 0 {
		sql = sql.Where("hash IN?", request.HashList)
	}

	if err := sql.Count(&total).Limit(request.Limit).Offset(request.Page * request.Limit).Order("timestamp DESC, index DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

// getTransactions get transaction data from database
func GetTransactionsByPlatform(ctx context.Context, request model.GetNotesByPlatformRequest) ([]dbModel.Transaction, error) {
	tracer := otel.Tracer("getTransactions")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	transactions := make([]dbModel.Transaction, 0)
	sql := database.Global().
		WithContext(ctx).
		Model(&dbModel.Transaction{}).
		Where("success IS TRUE"). // Hide failed transactions
		Where("platform = ?", request.Platform)

	if request.Platform == protocol.NetworkEIP1577 {
		sql = database.Global().
			WithContext(ctx).
			Model(&dbModel.Transaction{}).
			Where("success IS TRUE"). // Hide failed transactions
			Where("network = ?", request.Platform)
	}

	if len(request.Cursor) > 0 {
		var lastItem dbModel.Transaction

		// no need to lowercase
		if err := database.Global().Where("hash = ?", request.Cursor).First(&lastItem).Error; err != nil {
			return nil, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
	}

	if err := sql.Limit(request.Limit).Order("timestamp DESC, index DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func BatchGetTransactions(ctx context.Context, request model.BatchGetNotesRequest) ([]dbModel.Transaction, int64, error) {
	tracer := otel.Tracer("batchGetTransactions")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	transactions := make([]dbModel.Transaction, 0)
	total := int64(0)

	sql := database.Global().
		WithContext(ctx).
		Model(&dbModel.Transaction{}).
		Where("owner IN ?", request.Address).
		Where("success IS TRUE") // Hide failed transactions

	if len(request.Cursor) > 0 {
		var lastItem dbModel.Transaction

		// no need to lowercase
		if err := database.Global().Where("hash = ?", request.Cursor).First(&lastItem).Error; err != nil {
			return nil, 0, err
		}

		sql = sql.Where("timestamp < ? OR (timestamp = ? AND index < ?)", lastItem.Timestamp, lastItem.Timestamp, lastItem.Index)
	}

	if len(request.Tag) > 0 {
		sql = sql.Where("tag IN ?", request.Tag)
		if len(request.Type) > 0 {
			// type was already converted to lowercase
			sql = sql.Where("\"type\" IN ?", request.Type)
		}
	}

	if !request.IncludePoap {
		sql = sql.Where("\"type\" != ?", filter.CollectiblePoap)
	}

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}

		sql = sql.Where("LOWER(network) IN ?", request.Network)
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

// getTransfers get transfer data from database
func GetTransfers(ctx context.Context, transactionHashes []string) ([]dbModel.Transfer, error) {
	tracer := otel.Tracer("getTransfers")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	transfers := make([]dbModel.Transfer, 0)

	sql := database.Global().WithContext(ctx).Model(&dbModel.Transfer{})

	if err := sql.
		Where("transaction_hash IN (SELECT * FROM UNNEST(?::TEXT[]))", pq.Array(transactionHashes)).
		Find(&transfers).Error; err != nil {
		return nil, err
	}

	return transfers, nil
}

func GetTransactionByHash(ctx context.Context, request model.GetTransactionRequest) (dbModel.Transaction, error) {
	tracer := otel.Tracer("getTransactionByHash")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	var transaction dbModel.Transaction

	sql := database.Global().WithContext(ctx).Model(&dbModel.Transaction{})

	if err := sql.
		Where("hash = ?", request.Hash).
		First(&transaction).Error; err != nil {
		return dbModel.Transaction{}, err
	}

	return transaction, nil
}
