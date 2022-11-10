package database

import (
	"bytes"
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

func UpsertTransactions(ctx context.Context, transactions []*model.Transaction) error {
	dbChunkSize := 800
	transfers := []model.Transfer{}
	updatedTransactions := []model.Transaction{}

	for _, transaction := range transactions {
		// ignore empty tag and type
		if transaction.Type == "" || transaction.Tag == "" {
			continue
		}

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

			if transfer.Type == "" || transfer.Tag == "" {
				continue
			}

			transfers = append(transfers, transfer)
		}

		updatedTransactions = append(updatedTransactions, *transaction)
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

	return nil
}

func DeduplicateTransactions(ctx context.Context, transactions []*model.Transaction) ([]*model.Transaction, error) {
	hashList := []string{}
	for _, transaction := range transactions {
		hashList = append(hashList, transaction.Hash)
	}

	data := []*model.Transaction{}
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
