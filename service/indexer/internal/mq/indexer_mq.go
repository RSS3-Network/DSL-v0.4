package mq

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/samber/lo"
	"github.com/scylladb/go-set/strset"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Handler struct{}

var h *Handler

func IndexHandler() {
	deliveryCh, err := RabbitmqChannel.Consume(RabbitmqQueue.Name, "", true, false, false, false, nil)
	if err != nil {
		return
	}

	for delivery := range deliveryCh {
		message := protocol.Message{}
		if err := json.Unmarshal(delivery.Body, &message); err != nil {
			loggerx.Global().Error("failed to unmarshal message", zap.Error(err))

			continue
		}

		go func() {
			if err := h.handle(context.Background(), &message); err != nil {
				loggerx.Global().Error("failed to handle message", zap.Error(err), zap.String("address", message.Address), zap.String("network", message.Network))
			}
		}()
	}
}

func (h *Handler) handle(ctx context.Context, message *protocol.Message) (err error) {
	lockKey := fmt.Sprintf("indexer:%v:%v", message.Address, message.Network)

	if !shedlock.Global().DoLock(lockKey, 2*time.Minute) {
		return fmt.Errorf("%v lock", lockKey)
	}

	cctx, cancel := context.WithCancel(context.Background())
	go func(cctx context.Context) {
		for {
			time.Sleep(time.Second)
			if err := shedlock.Global().Renewal(cctx, lockKey, time.Minute); err != nil {
				return
			}
		}
	}(cctx)

	defer cancel()

	var transactions []model.Transaction
	defer func() {
		shedlock.Global().UnLock(lockKey)

		transfers := 0

		for _, transaction := range transactions {
			transfers += len(transaction.Transfers)
		}

		loggerx.Global().Info("indexed data completion", zap.String("address", message.Address), zap.String("network", message.Network), zap.Int("transactions", len(transactions)), zap.Int("transfers", transfers))

		// upsert address status
		go h.upsertAddress(ctx, model.Address{
			Address: message.Address,
		})
	}()

	// convert address to lowercase
	message.Address = strings.ToLower(message.Address)
	tracer := otel.Tracer("indexer")

	ctx, handlerSpan := tracer.Start(ctx, "indexer:handler")

	handlerSpan.SetAttributes(
		attribute.String("network", message.Network),
	)

	defer handlerSpan.End()

	loggerx.Global().Info("start indexing data", zap.String("address", message.Address), zap.String("network", message.Network))

	// Ignore triggers
	if !message.IgnoreTrigger {
		if err := executeTriggers(ctx, message); err != nil {
			zap.L().Error("failed to execute trigger", zap.Error(err), zap.String("address", message.Address), zap.String("network", message.Network))
		}
	}

	// Ignore notes
	if message.IgnoreNote {
		return nil
	}

	// Open a database transaction
	tx := database.Global().WithContext(ctx).Begin()

	for _, datasource := range s.datasources {
		for _, network := range datasource.Networks() {
			if network == message.Network {
				// Get the time of the latest data for this address and network
				var result struct {
					Timestamp   time.Time `gorm:"column:timestamp"`
					BlockNumber int64     `gorm:"column:block_number"`
				}

				// Delete data from this address and reindex it
				if message.Reindex {
					var hashes []string

					// TODO Use the owner to replace hashes field
					// Get all hashes of this address on this network
					if err := tx.
						Model((*model.Transaction)(nil)).
						Where("network = ? AND owner = ?", message.Network, message.Address).
						Pluck("hash", &hashes).
						Error; err != nil {
						return err
					}

					if err := tx.Where("network = ? AND hash IN (SELECT * FROM UNNEST(?::TEXT[]))", message.Network, pq.Array(hashes)).Delete(&model.Transaction{}).Error; err != nil {
						tx.Rollback()

						return err
					}

					if err := tx.Where("network = ? AND transaction_hash IN (SELECT * FROM UNNEST(?::TEXT[]))", message.Network, pq.Array(hashes)).Delete(&model.Transfer{}).Error; err != nil {
						tx.Rollback()

						return err
					}
				}

				if err := tx.
					Model((*model.Transaction)(nil)).
					Select("COALESCE(timestamp, 'epoch'::timestamp) AS timestamp, COALESCE(block_number, 0) AS block_number").
					Where("owner = ?", message.Address).
					Where("network = ?", message.Network).
					Order("timestamp DESC").
					Limit(1).
					First(&result).
					Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					continue
				}

				message.Timestamp = result.Timestamp
				message.BlockNumber = result.BlockNumber

				internalTransactions, err := datasource.Handle(ctx, message)
				// Avoid blocking indexed workers
				if err != nil {
					loggerx.Global().Error("datasource handle failed", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address), zap.String("datasource", datasource.Name()))

					continue
				}

				transactions = append(transactions, internalTransactions...)
			}
		}
	}

	transactionsMap := getTransactionsMap(transactions)

	return s.handleWorkers(ctx, message, tx, transactions, transactionsMap)
}

func (h *Handler) handleWorkers(ctx context.Context, message *protocol.Message, tx *gorm.DB, transactions []model.Transaction, transactionsMap map[string]model.Transaction) (err error) {
	tracer := otel.Tracer("indexer")
	ctx, span := tracer.Start(ctx, "indexer:handleWorkers")

	defer opentelemetry.Log(span, message, transactions, err)

	// Using workers to clean data
	for _, worker := range s.workers {
		for _, network := range worker.Networks() {
			if network == message.Network {
				internalTransactions, err := worker.Handle(ctx, message, transactions)
				if err != nil {
					loggerx.Global().Error("worker handle failed", zap.Error(err), zap.String("worker", worker.Name()), zap.String("network", network))

					continue
				}

				if len(internalTransactions) == 0 {
					continue
				}

				newTransactionMap := getTransactionsMap(internalTransactions)
				for _, t := range newTransactionMap {
					transactionsMap[t.Hash] = t
				}

				transactions = transactionsMap2Array(transactionsMap)
			}
		}
	}

	return h.upsertTransactions(ctx, message, tx, transactions)
}

func (h *Handler) upsertTransactions(ctx context.Context, message *protocol.Message, tx *gorm.DB, transactions []model.Transaction) (err error) {
	tracer := otel.Tracer("indexer")
	_, span := tracer.Start(ctx, "indexer:upsertTransactions")

	defer opentelemetry.Log(span, len(transactions), nil, err)

	dbChunkSize := 800

	var (
		transfers           []model.Transfer
		updatedTransactions []model.Transaction
	)

	for _, transaction := range transactions {
		addresses := strset.New(transaction.AddressFrom, transaction.AddressTo)

		// Ignore empty transactions
		internalTransfers := make([]model.Transfer, 0)

		for _, transfer := range transaction.Transfers {
			if bytes.Equal(transfer.Metadata, metadata.Default) {
				continue
			}

			internalTransfers = append(internalTransfers, transfer)
		}

		if len(internalTransfers) == 0 {
			continue
		}

		// Handle all transfers
		for _, transfer := range transaction.Transfers {
			// Ignore empty transfer
			if bytes.Equal(transfer.Metadata, metadata.Default) {
				continue
			}

			transfers = append(transfers, transfer)

			if transfer.AddressFrom != "" && transfer.AddressFrom != ethereum.AddressGenesis.String() {
				addresses.Add(transfer.AddressFrom)
			}

			if transfer.AddressTo != "" && transfer.AddressTo != ethereum.AddressGenesis.String() {
				addresses.Add(transfer.AddressTo)
			}
		}

		transaction.Addresses = addresses.List()
		updatedTransactions = append(updatedTransactions, transaction)
	}

	for _, ts := range lo.Chunk(updatedTransactions, dbChunkSize) {
		if err = tx.
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			Create(ts).Error; err != nil {
			loggerx.Global().Error("failed to upsert transactions", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address))

			tx.Rollback()

			return err
		}
	}

	for _, ts := range lo.Chunk(transfers, dbChunkSize) {
		if err = tx.
			Clauses(clause.OnConflict{
				UpdateAll: true,
				DoUpdates: clause.AssignmentColumns([]string{"metadata"}),
			}).
			Create(ts).Error; err != nil {
			loggerx.Global().Error("failed to upsert transfers", zap.Error(err), zap.String("network", message.Network), zap.String("address", message.Address))

			tx.Rollback()

			return err
		}
	}

	return tx.Commit().Error
}

func (h *Handler) upsertAddress(ctx context.Context, address model.Address) {
	for _, network := range protocol.SupportNetworks {
		key := fmt.Sprintf("indexer:%v:%v", address.Address, network)
		n, err := cache.Global().Exists(ctx, key).Result()
		switch {
		case err != nil:
			return
		case n == 1: // exists
			address.IndexingNetworks = append(address.IndexingNetworks, network)
		default: // not exists (unlocked)
			address.DoneNetworks = append(address.DoneNetworks, network)
		}
	}

	if len(address.IndexingNetworks) == 0 {
		address.Status = true
	}

	if err := database.Global().
		Model(model.Address{}).
		Clauses(clause.OnConflict{
			UpdateAll: true,
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).
		Create(map[string]interface{}{
			"address":           address.Address,
			"status":            address.Status,
			"done_networks":     address.DoneNetworks,
			"indexing_networks": address.IndexingNetworks,
		}).Error; err != nil {
		loggerx.Global().Error("failed to upsert address", zap.Error(err), zap.String("address", address.Address))
	}
}

func getTransactionsMap(transactions []model.Transaction) map[string]model.Transaction {
	transactionsMap := make(map[string]model.Transaction)

	for _, t := range transactions {
		transactionsMap[t.Hash] = t
	}

	return transactionsMap
}

func transactionsMap2Array(transactionsMap map[string]model.Transaction) []model.Transaction {
	transactions := make([]model.Transaction, 0)

	for _, t := range transactionsMap {
		transactions = append(transactions, t)
	}

	return transactions
}
