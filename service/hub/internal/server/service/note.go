package service

import (
	"context"
	"strings"

	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
)

func (s *Service) GetNotes(ctx context.Context, request model.GetRequest) ([]dbModel.Transaction, int64, error) {
	request.Address = strings.ToLower(request.Address)

	if request.Limit <= 0 || request.Limit > model.DefaultLimit {
		request.Limit = model.DefaultLimit
	}

	if len(request.Tag) > 0 {
		request.Tag, request.Type, request.IncludePoap = s.CheckRequestTagAndType(request.Tag, request.Type)
	}

	// get transactions from database
	transactions, total, err := dao.GetTransactions(ctx, request)
	if err != nil {
		return nil, 0, err
	}

	// get transfers from database
	transactionHashes := make([]string, 0)
	for _, transactionHash := range transactions {
		transactionHashes = append(transactionHashes, transactionHash.Hash)
	}

	transfers, err := dao.GetTransfers(ctx, transactionHashes)
	if err != nil {
		return nil, 0, err
	}

	transferMap := make(map[string][]dbModel.Transfer)
	for _, transfer := range transfers {
		transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
	}

	for index := range transactions {
		transactions[index].Transfers = transferMap[transactions[index].Hash]
	}

	// publish mq message
	if len(request.Cursor) == 0 && (request.Refresh || len(transactions) == 0) {
		s.PublishIndexerMessage(ctx, protocol.Message{Address: request.Address, Reindex: request.Reindex})
	}

	return transactions, total, nil
}

func (s *Service) BatchGetNotes(ctx context.Context, request model.BatchGetNotesRequest) ([]dbModel.Transaction, int64, error) {
	if request.Limit <= 0 || request.Limit > model.DefaultLimit {
		request.Limit = model.DefaultLimit
	}

	if len(request.Address) > model.DefaultLimit {
		request.Address = request.Address[:model.DefaultLimit]
	}

	if len(request.Tag) > 0 {
		request.Tag, request.Type, request.IncludePoap = s.CheckRequestTagAndType(request.Tag, request.Type)
	}

	for i, v := range request.Address {
		request.Address[i] = strings.ToLower(v)
	}

	transactions, total, err := dao.BatchGetTransactions(ctx, request)
	if err != nil {
		return nil, 0, err
	}

	transactionHashes := make([]string, 0)
	for _, transactionHash := range transactions {
		transactionHashes = append(transactionHashes, transactionHash.Hash)
	}

	transfers, err := dao.GetTransfers(ctx, transactionHashes)
	if err != nil {
		return nil, 0, err
	}

	transferMap := make(map[string][]dbModel.Transfer)
	for _, transfer := range transfers {
		transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
	}

	for index := range transactions {
		transactions[index].Transfers = transferMap[transactions[index].Hash]
	}

	// publish mq message
	if len(request.Cursor) == 0 && (request.Refresh || len(transactions) == 0) {
		for _, address := range request.Address {
			s.PublishIndexerMessage(ctx, protocol.Message{Address: address})
		}
	}

	return transactions, total, nil
}

func (s *Service) CheckRequestTagAndType(reqTags []string, reqTypes []string) ([]string, []string, bool) {
	// support many-many relationship between tag and type
	var tags []string
	var types []string
	var includePoap bool

	for _, tag := range reqTags {
		tag = strings.ToLower(tag)
		if len(reqTypes) > 0 {
			for _, typeX := range reqTypes {
				typeX = strings.ToLower(typeX)

				if filter.CheckTypeValid(tag, typeX) {
					tags = append(tags, tag)
					types = append(types, typeX)
					// by default POAPs are not returned
					if typeX == filter.CollectiblePoap {
						includePoap = true
					}
				}
			}
		} else {
			tags = append(tags, tag)
		}
	}

	return tags, types, includePoap
}
