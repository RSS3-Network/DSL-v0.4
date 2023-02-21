package database

import (
	"bytes"
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/samber/lo"
	"github.com/scylladb/go-set/strset"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

func UpsertTransactions(ctx context.Context, transactions []*model.Transaction, dedupTransfer bool) error {
	dbChunkSize := 800
	var transfers []model.Transfer
	var updatedTransactions []model.Transaction

	for _, transaction := range transactions {
		// ignore empty tag and type
		if transaction.Type == "" || transaction.Tag == "" {
			continue
		}

		// Ignore empty transactions
		internalTransfers := make([]model.Transfer, 0)

		for _, transfer := range transaction.Transfers {
			if bytes.Equal(transfer.Metadata, metadata.Default) || len(transfer.Metadata) == 0 {
				continue
			}

			if transfer.Type == "" || transfer.Tag == "" {
				continue
			}

			internalTransfers = append(internalTransfers, transfer)

			transfers = append(transfers, transfer)
		}

		if len(internalTransfers) == 0 {
			continue
		}

		updatedTransactions = append(updatedTransactions, *transaction)
	}

	// Deduplicate Transfers
	// eg: multi owners own one tx which is mapping the same transfer in farcaster
	if dedupTransfer {
		transfersMap := getTransfersMap(transfers)
		transfers = transfersMap2Array(transfersMap)
	}

	for _, ts := range lo.Chunk(updatedTransactions, dbChunkSize) {
		if err := Global().
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			Create(ts).Error; err != nil {
			loggerx.Global().Error("failed to upsert transactions", zap.Error(err))

			return err
		}
	}

	for _, ts := range lo.Chunk(transfers, dbChunkSize) {
		if err := Global().
			Clauses(clause.OnConflict{
				UpdateAll: true,
				DoUpdates: clause.AssignmentColumns([]string{"metadata"}),
			}).
			Create(ts).Error; err != nil {
			loggerx.Global().Error("failed to upsert transfers", zap.Error(err))

			return err
		}
	}

	// 这里出错不影响主流程，因此不 return
	addresses := strset.New()
	for _, transaction := range updatedTransactions {
		addresses.Add(transaction.Owner)
	}
	for _, address := range addresses.List() {
		_ = SaveLatestTxHashByAddress(ctx, address)
	}

	return nil
}

var (
	tagMask      = []string{"social", "collectible", "governance", "donation"}
	typeMask     = []string{"transfer", "trade", "mint", "burn", "poap", "post", "revise", "comment", "launch", "donate", "propose", "vote"}
	networkMask  = []string{"ethereum", "polygon", "arbitrum", "optimism", "binance_smart_chain", "farcaster", "eip1577", "arweave"}
	platformMask = []string{"Mirror", "Lens", "EIP-1577", "Gitcoin", "Snapshot", "Farcaster", ""}
)

func IsMaskReq(tag, _type, network, platform []string) bool {
	if strset.New(tag...).IsEqual(strset.New(tagMask...)) &&
		strset.New(_type...).IsEqual(strset.New(typeMask...)) &&
		strset.New(network...).IsEqual(strset.New(networkMask...)) &&
		strset.New(platform...).IsEqual(strset.New(platformMask...)) {
		return true
	}
	return false
}

// 存储最新的 500 条 tx hash，为了优化 mask 来频繁请求的场景，即请求 hub batch 时得到所有地址的 tx 中最新的 500 条
// 这个场景如果一次性通过所有地址的 tx 来筛选最新的 500 条，扫描行数会非常大（百万行级别）然后排序，此时并发稍高一点的话数据库就完全没有办法处理，会非常慢
// 如果有一百万个地址，每个地址使用 20k 的内存，这个场景下需要 20G 的 redis
func SaveLatestTxHashByAddress(ctx context.Context, address string) error {
	latest500Hash := []struct {
		Hash      string    `json:"hash"`
		Timestamp time.Time `json:"timestamp"`
	}{}
	if err := Global().
		Model((*model.Transaction)(nil)).
		Select("hash", "timestamp").
		Where("owner = ?", address).
		Where("tag IN (?)", tagMask).
		Where("type IN (?)", typeMask).
		Where("network IN (?)", networkMask).
		Where("platform IN (?)", platformMask).
		Order("timestamp DESC, index DESC").
		Limit(500).Find(&latest500Hash).Error; err != nil {
		loggerx.Global().Error("get latest 500 tx hash failed", zap.Error(err), zap.String("address", address))
		return err
	} else {
		key := strings.Join([]string{protocol.LATEST_TX_HASH_KEY, address}, ":")
		value := []*redis.Z{}

		for _, tx := range latest500Hash {
			value = append(value, &redis.Z{
				Score:  float64(tx.Timestamp.Unix()),
				Member: tx.Hash,
			})
		}

		if len(value) > 0 {
			if err := cache.Global().ZAdd(ctx, key, value...).Err(); err != nil {
				loggerx.Global().Error("save latest 500 tx hash failed", zap.Error(err), zap.String("address", address))
				return err
			}

			// only keep top 500 tx hash by ts
			if err := cache.Global().ZRemRangeByRank(ctx, key, 0, -501).Err(); err != nil {
				loggerx.Global().Error("rem old tx hash failed", zap.Error(err), zap.String("address", address))
				return err
			}
		}

		return nil
	}
}

func DeduplicateTransactions(ctx context.Context, transactions []*model.Transaction) ([]*model.Transaction, error) {
	var hashList []string
	for _, transaction := range transactions {
		hashList = append(hashList, transaction.Hash)
	}

	var data []*model.Transaction
	if err := Global().
		Model(&model.Transaction{}).
		Where("hash in (?)", hashList).
		Find(&data).Error; err != nil {
		logrus.Error("[lens] deduplicateTransactions: find transactions error, ", err)

		return transactions, err
	}

	dbMap := make(map[string]bool)
	for _, t := range data {
		dbMap[t.Hash] = true
	}

	for i := 0; i < len(transactions); i++ {
		if dbMap[transactions[i].Hash] {
			transactions = append(transactions[:i], transactions[i+1:]...)
			i--
		}
	}

	return transactions, nil
}

func getTransfersMap(transfers []model.Transfer) map[string]model.Transfer {
	transfersMap := make(map[string]model.Transfer)

	for _, t := range transfers {
		transfersMap[t.TransactionHash] = t
	}

	return transfersMap
}

func transfersMap2Array(transfersMap map[string]model.Transfer) []model.Transfer {
	var transfers []model.Transfer

	for _, t := range transfersMap {
		transfers = append(transfers, t)
	}

	return transfers
}
