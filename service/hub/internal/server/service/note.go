package service

import (
	"context"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/types"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
)

func (s *Service) GetNotes(ctx context.Context, request model.GetRequest) ([]dbModel.Transaction, int64, error) {
	request.Address = strings.ToLower(request.Address)

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
		if utils.SupportLensPlatform[strings.ToLower(transfer.Platform)] {
			metadata := gjson.ParseBytes(transfer.Metadata)
			profileID := metadata.Get("profile_id").Int()
			publicationID := metadata.Get("publication_id").Int()

			if profileID > 0 && publicationID > 0 {
				lensterURL := utils.GetLensRelatedURL(big.NewInt(profileID), big.NewInt(publicationID))
				transfer.RelatedUrls = append(transfer.RelatedUrls, lensterURL)
			}
		}

		transfer.RelatedUrls = lo.Uniq(transfer.RelatedUrls)

		if len(transferMap[transfer.TransactionHash]) < request.ActionLimit {
			transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
		}
	}

	for index := range transactions {
		transactions[index].Transfers = transferMap[transactions[index].Hash]
	}

	// publish mq message
	if len(request.Cursor) == 0 && (request.Refresh || len(transactions) == 0) {
		s.PublishIndexerMessage(ctx, protocol.Message{Address: request.Address})
	}

	return transactions, total, nil
}

func (s *Service) GetNotesByPlatform(ctx context.Context, request model.GetNotesByPlatformRequest) ([]dbModel.Transaction, error) {
	// get transactions from database
	transactions, err := dao.GetTransactionsByPlatform(ctx, request)
	if err != nil {
		return nil, err
	}

	// get transfers from database
	transactionHashes := make([]string, 0)
	for _, transactionHash := range transactions {
		transactionHashes = append(transactionHashes, transactionHash.Hash)
	}

	transfers, err := dao.GetTransfers(ctx, transactionHashes)
	if err != nil {
		return nil, err
	}

	transferMap := make(map[string][]dbModel.Transfer)
	for _, transfer := range transfers {
		transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
	}

	for index := range transactions {
		transactions[index].Transfers = transferMap[transactions[index].Hash]
	}

	return transactions, nil
}

func (s *Service) BatchGetNotes(ctx context.Context, request model.BatchGetNotesRequest) ([]dbModel.Transaction, int64, error) {
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
		if len(transferMap[transfer.TransactionHash]) < request.ActionLimit {
			transferMap[transfer.TransactionHash] = append(transferMap[transfer.TransactionHash], transfer)
		}
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
	if request.Limit <= 0 || request.Limit > model.DefaultLimit {
		request.Limit = model.DefaultLimit
	}

	transactions, total, err := dao.BatchGetSocialTransactions(ctx, request)
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

	return transactions, total, nil
}

func (s *Service) CheckRequestTagAndType(reqTags []string, reqTypes []string) ([]string, []string, bool) {
	// support many-many relationship between tag and type
	var tags []string
	var typeList []string
	var includePoap bool

	for _, tag := range reqTags {
		tag = strings.ToLower(tag)
		if len(reqTypes) > 0 {
			for _, typeX := range reqTypes {
				typeX = strings.ToLower(typeX)

				if types.CheckTypeValid(tag, typeX) {
					tags = append(tags, tag)
					typeList = append(typeList, typeX)
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

	return tags, typeList, includePoap
}

func (s *Service) GetNftFeeds(ctx context.Context, request model.GetRequest) ([]dbModel.Transaction, int64, error) {
	request.Address = strings.ToLower(request.Address)

	if request.Limit <= 0 || request.Limit > model.DefaultLimit {
		request.Limit = model.DefaultLimit
	}

	if len(request.Tag) > 0 {
		request.Tag, request.Type, request.IncludePoap = s.CheckRequestTagAndType(request.Tag, request.Type)
	}

	raraQuery := kurora.DatasetRaraEntryQuery{
		NftAddress: lo.ToPtr(common.HexToAddress(request.Address)),
	}

	entries, err := s.kuroraClient.FetchDatasetRara(ctx, raraQuery)
	if err != nil {
		return nil, 0, err
	}

	var tokenID *big.Int
	if len(request.TokenId) > 0 {
		tokenID = big.NewInt(0)
		tokenID.SetString(request.TokenId, 0)
	}
	hashes := make([]string, 0)

	for _, entry := range entries {
		if len(request.TokenId) > 0 && tokenID.String() != entry.NftId.BigInt().String() {
			continue
		}
		hashes = append(hashes, strings.ToLower(entry.TransactionHash.String()))
	}

	if len(request.TokenId) == 0 {
		request.TokenId = "0"
	}

	transactions := make([]dbModel.Transaction, 0)
	request.HashList = hashes

	if len(request.HashList) == 0 {
		return transactions, 0, nil
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

	return transactions, total, nil
}

func (s *Service) GetTransactionByHash(ctx context.Context, request model.GetTransactionRequest) (dbModel.Transaction, error) {
	transaction, err := dao.GetTransactionByHash(ctx, request)
	if err != nil {
		return dbModel.Transaction{}, err
	}

	transfers, err := dao.GetTransfers(ctx, []string{transaction.Hash})
	if err != nil {
		return dbModel.Transaction{}, err
	}

	transaction.Transfers = transfers

	return transaction, nil
}
