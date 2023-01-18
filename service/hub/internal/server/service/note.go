package service

import (
	"context"
	"math/rand"
	"strings"
	"sync"
	"time"

	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/samber/lo"
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

	if len(request.Tag) > 0 {
		request.Tag, request.Type, request.IncludePoap = s.CheckRequestTagAndType(request.Tag, request.Type)
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
		// limit 10 addresses each time, max 50 times
		// 如果一次来 500 个，我们就手动把任务曲线拉平，500 个任务分布到 5s 发送
		// 请求的地址越多我们响应的越慢，很合理吧
		for _, addresses := range lo.Chunk(request.Address, 10) {
			wg := &sync.WaitGroup{}
			for _, address := range addresses {
				wg.Add(1)
				go func(address string) {
					defer wg.Done()
					n := rand.Intn(100)
					time.Sleep(time.Duration(n) * time.Millisecond) // max 100ms * 50 = 5s
					s.PublishIndexerMessage(ctx, protocol.Message{Address: address})
				}(address)
			}
			wg.Wait()
		}
	}

	return transactions, total, nil
}

func (s *Service) BatchGetSocialNotes(ctx context.Context, request model.BatchGetSocialNotesRequest) ([]dbModel.Transaction, int64, error) {
	return s.BatchGetNotes(ctx, model.BatchGetNotesRequest{
		Address:        request.Address,
		Type:           request.Type,
		Tag:            []string{"social"},
		Network:        request.Network,
		Platform:       request.Platform,
		Timestamp:      request.Timestamp,
		Limit:          request.Limit,
		Cursor:         request.Cursor,
		Refresh:        false,
		IncludePoap:    false,
		Page:           request.Page,
		QueryStatus:    request.QueryStatus,
		CountOnly:      request.CountOnly,
		IgnoreContract: true,
	})
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
