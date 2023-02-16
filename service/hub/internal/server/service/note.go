package service

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/nft"
	nft_contract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/nft"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/samber/lo"

	"go.uber.org/zap"
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

func (s *Service) GetNftFeeds(ctx context.Context, request model.GetRequest) ([]nft.Feed, int64, error) {
	feeds := make([]nft.Feed, 0)
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

	if len(entries) == 0 {
		return feeds, 0, err
	}

	var tokenID *big.Int
	if len(request.TokenId) > 0 {
		tokenID = big.NewInt(0)
		tokenID.SetString(request.TokenId, 0)
	}

	nftMap := make(map[string]*nft.Feed)
	tokenClient := token.New()

	for _, entry := range entries {
		network := protocol.IdToNetwork(fmt.Sprintf("%#x", entry.NftChainId.CoefficientInt64()))
		// 存在，加 action
		if feed, ok := nftMap[entry.NftId.BigInt().String()]; ok {
			comment := &metadata.Post{
				Body:           entry.Comment,
				Tags:           entry.Tags,
				Author:         []string{strings.ToLower(entry.From.String()), fmt.Sprintf("%s%s", "https://app.rara.social/profile/", strings.ToLower(entry.From.String()))},
				TypeOnPlatform: []string{filter.SocialComment},
			}
			feed.Actions = append(feed.Actions, *comment)
		} else if len(request.TokenId) == 0 || tokenID.String() == entry.NftId.BigInt().String() {
			nftInfo := &token.NFT{
				ContractAddress: strings.ToLower(entry.NftAddress.String()),
				ID:              entry.NftId.BigInt(),
			}
			switch entry.NftAddress {
			case nft_contract.CryptoKitties:
				ckMetadata, err := nft_contract.GetCkMetadata(ctx, entry.NftId)
				if err != nil {
					loggerx.Global().Warn("failed to handle CK NFT metadata", zap.Error(err), zap.String("tokenID", entry.NftId.String()))

					continue
				}
				nftInfo.Name = ckMetadata.Name
				nftInfo.Description = ckMetadata.Bio
				nftInfo.Image = ckMetadata.Image
				nftInfo.Standard = "ERC721"
			case nft_contract.CryptoPunks:
				nftInfo.Name = nft_contract.CryptoPunksName
				nftInfo.Description = nft_contract.CryptoPunksDes
				nftInfo.Standard = "ERC20"
			default:
				nftInfo, err = tokenClient.NFT(ctx, network, entry.NftAddress.String(), entry.NftId.BigInt())
				if err != nil {
					return feeds, 0, err
				}
			}

			nftMap[entry.NftId.BigInt().String()] = &nft.Feed{
				Nft: *nftInfo,
				Actions: []metadata.Post{
					{
						Body:           entry.Comment,
						Tags:           entry.Tags,
						Author:         []string{strings.ToLower(entry.From.String()), fmt.Sprintf("%s%s", "https://app.rara.social/profile/", strings.ToLower(entry.From.String()))},
						TypeOnPlatform: []string{filter.SocialComment},
					},
				},
				RelatedUrls: []string{fmt.Sprintf("%s/%s/%s/%s", "https://app.rara.social/nft", entry.NftChainId, strings.ToLower(entry.NftAddress.String()), entry.NftId)},
			}
		}
	}

	for _, nftInfo := range nftMap {
		feeds = append(feeds, *nftInfo)
	}

	return feeds, int64(len(feeds)), nil
}
