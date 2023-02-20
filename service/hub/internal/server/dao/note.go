package dao

import (
	"context"
	"sort"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
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
		Model(&dbModel.Transaction{}).
		Where("owner = ?", request.Address). // address was already converted to lowercase
		Where("success IS TRUE")             // Hide failed transactions

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

	if err := sql.Count(&total).Limit(request.Limit).Offset(request.Page * request.Limit).Order("timestamp DESC, index DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

func BatchGetTransactions(ctx context.Context, request model.BatchGetNotesRequest) ([]dbModel.Transaction, int64, error) {
	tracer := otel.Tracer("batchGetTransactions")
	_, postgresSnap := tracer.Start(ctx, "postgres")

	defer postgresSnap.End()

	transactions := make([]dbModel.Transaction, 0)
	total := int64(0)
	usedRedis := false

	// 针对 mask 常来刷新的特殊优化，从 redis 里面取最新的 Limit 条 hash，而不是用 owner in
	//  redis 操作失败时，从数据库取
	// （从数据库取是没问题的，但几十个这样的并发 SQL 会导致每句都超过 60s）
	if database.IsMaskReq(request.Tag, request.Type, request.Network, request.Platform) {
		failed := false

		allTx := []struct {
			Hash      string
			Timestamp int64
		}{}

		for _, v := range request.Address {
			key := strings.Join([]string{protocol.LATEST_TX_HASH_KEY, v}, ":")

			hashAndTS, err := cache.Global().ZRevRange(ctx, key, 0, int64(request.Limit-1)).Result()
			if err != nil {
				failed = true
				break
			}
			for index, v := range hashAndTS {
				if index%2 == 0 { // hash
					allTx = append(allTx, struct {
						Hash      string
						Timestamp int64
					}{Hash: v})
				} else { // timestamp
					allTx[len(allTx)-1].Timestamp, _ = strconv.ParseInt(v, 10, 64)
				}
			}
		}
		if !failed {
			sort.SliceStable(allTx, func(i, j int) bool {
				return allTx[i].Timestamp > allTx[j].Timestamp
			})
			latestTx := allTx[:request.Limit]
			if err := database.Global().Where("hash IN ?", latestTx).Find(&transactions).Error; err != nil {
				return nil, 0, err
			}
			usedRedis = true
		}
	}

	sql := database.Global().
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

	if usedRedis { // only need total
		if err := sql.Count(&total).Error; err != nil {
			return nil, 0, err
		}
	} else { // default
		if err := sql.Count(&total).Limit(request.Limit).Offset(request.Page * request.Limit).Order("timestamp DESC, index DESC").Find(&transactions).Error; err != nil {
			return nil, 0, err
		}
	}

	return transactions, total, nil
}

// getTransfers get transfer data from database
func GetTransfers(c context.Context, transactionHashes []string) ([]dbModel.Transfer, error) {
	tracer := otel.Tracer("getTransfers")
	_, postgresSnap := tracer.Start(c, "postgres")

	defer postgresSnap.End()

	transfers := make([]dbModel.Transfer, 0)

	sql := database.Global().Model(&dbModel.Transfer{})

	if err := sql.
		Where("transaction_hash IN (SELECT * FROM UNNEST(?::TEXT[]))", pq.Array(transactionHashes)).
		Find(&transfers).Error; err != nil {
		return nil, err
	}

	return transfers, nil
}
