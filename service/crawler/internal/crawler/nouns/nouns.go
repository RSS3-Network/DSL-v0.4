package nouns

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/element"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type service struct {
	config       *config.Config
	kuroraClient *kurora.Client
	tokenClient  *token.Client
}

var (
	_ crawler.Crawler = (*service)(nil)

	nounsCacheKey = "crawler:nouns"
)

func New(config *config.Config) crawler.Crawler {
	return &service{
		config:      config,
		tokenClient: token.New(),
	}
}

func (s *service) Name() string {
	return protocol.PlatformNouns
}

func (s *service) Network() string {
	return protocol.NetworkEthereum
}

func (s *service) Run() error {
	loggerx.Global().Info("nouns: run")

	ctx := context.Background()
	var err error

	s.kuroraClient, err = kurora.Dial(ctx, s.config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error("nouns: kurora.Dial error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))

		return err
	}

	for {
		transactions, cacheInfo, err := s.handleNounsAuctions(ctx)
		if err != nil {
			loggerx.Global().Error("nouns: handleNounsAuctions error", zap.Error(err))

			return err
		}

		if len(transactions) == 0 {
			time.Sleep(1 * time.Hour)

			continue
		}

		// insert db
		err = database.UpsertTransactions(ctx, transactions, true)
		if err != nil {
			continue
		}

		// set cache
		cache.Global().Set(ctx, nounsCacheKey, cacheInfo, 7*24*time.Hour)
	}
}

func (s *service) handleNounsAuctions(ctx context.Context) ([]*model.Transaction, string, error) {
	tracer := otel.Tracer("nouns")
	_, trace := tracer.Start(ctx, "nouns:handleNounsAuctions")
	var (
		transactions []*model.Transaction
		err          error

		transactionsMap = make(map[common.Hash]*model.Transaction, 0)
	)
	defer func() { opentelemetry.Log(trace, nil, transactions, err) }()

	cursor, _ := cache.Global().Get(ctx, nounsCacheKey).Result()

	query := kurora.DatasetNounsEventQuery{
		Limit: lo.ToPtr(200),
	}

	if len(cursor) > 0 {
		query.Cursor = &cursor
	}

	events, err := s.kuroraClient.FetchDatasetNounsEvents(ctx, query)
	if err != nil {
		loggerx.Global().Error("FetchDatasetNounsEvents error", zap.Error(err), zap.Any("query", query))

		return nil, "", err
	}

	loggerx.Global().Info("zora: kuroraClient FetchDatasetNounsEvents result", zap.Int("len", len(events)), zap.String("cursor", cursor))

	var (
		limiter     = make(chan struct{}, 10)
		metaDataMap = make(map[string]*metadata.Token, 0)
		nftMap      = make(map[string]*struct {
			address common.Address
			tokenId *big.Int
		}, 0)

		waitGroup sync.WaitGroup
		locker    sync.Mutex
	)

	for _, event := range events {
		key := fmt.Sprintf("%v-%v", event.NftAddress.String(), event.NftId.BigInt())
		if _, ok := nftMap[key]; ok {
			continue
		}
		nftMap[key] = &struct {
			address common.Address
			tokenId *big.Int
		}{address: event.NftAddress, tokenId: event.NftId.BigInt()}
	}

	for _, nftData := range nftMap {
		limiter <- struct{}{}
		waitGroup.Add(1)

		go func(address string, tokenId *big.Int) {
			defer func() {
				<-limiter
				waitGroup.Done()
			}()

			nft, err := s.tokenClient.NFTToMetadata(ctx, s.Network(), address, tokenId)
			if err != nil {
				zap.L().Error("nft to metadata", zap.Error(err), zap.String("contract_address", address), zap.Stringer("token_id", tokenId))

				return
			}

			locker.Lock()
			metaDataMap[fmt.Sprintf("%v-%v", address, tokenId)] = nft
			locker.Unlock()
		}(nftData.address.String(), nftData.tokenId)
	}

	waitGroup.Wait()
	for _, event := range events {
		key := fmt.Sprintf("%v-%v", event.NftAddress.String(), event.NftId.BigInt())
		if _, ok := metaDataMap[key]; !ok {
			continue
		}

		nft := metaDataMap[key]

		nft.Action = event.EventType
		nft.StartTime = &event.StartTime
		nft.EndTime = &event.EndTime

		cost, err := s.buildCost(ctx, s.Network(), ethereum.AddressGenesis, event.Amount.BigInt())
		if err != nil {
			zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", event.TransactionHash.String()), zap.Stringer("contract_address", event.NftAddress), zap.Stringer("id", event.NftId))

			continue
		}

		nft.Cost = cost

		metadataRaw, err := json.Marshal(nft)
		if err != nil {
			return nil, cursor, fmt.Errorf("marshal metadata: %w", err)
		}

		var from, to string
		switch event.EventType {
		case filter.CollectibleAuctionBid:
			from, to = strings.ToLower(event.From.String()), strings.ToLower(event.To.String())
		case filter.CollectibleAuctionFinalize:
			from, to = strings.ToLower(event.To.String()), strings.ToLower(event.From.String())
		}

		transfer := model.Transfer{
			TransactionHash: strings.ToLower(event.TransactionHash.String()),
			Timestamp:       event.Timestamp,
			Index:           int64(event.LogIndex),
			AddressFrom:     from,
			AddressTo:       to,
			Metadata:        metadataRaw,
			Network:         s.Network(),
			Platform:        protocol.PlatformNouns,
			Source:          protocol.SourceKurora,
			Tag:             filter.TagCollectible,
			Type:            filter.CollectibleAuction,
			RelatedUrls: ethereum.BuildURL(
				[]string{utils.GetTxHashURL(s.Network(), event.TransactionHash.String()), fmt.Sprintf("%s%v", "https://nouns.wtf/noun/", event.NftId)},
				ethereum.BuildTokenURL(s.Network(), event.NftAddress.String(), event.NftId.String())...,
			),
		}

		if _, ok := transactionsMap[event.TransactionHash]; !ok {
			// append new tx
			transaction := &model.Transaction{
				BlockNumber: int64(event.BlockNumber),
				Timestamp:   event.Timestamp,
				Owner:       strings.ToLower(event.Owner.String()),
				Hash:        strings.ToLower(event.TransactionHash.String()),
				AddressFrom: strings.ToLower(event.From.String()),
				AddressTo:   strings.ToLower(event.To.String()),
				Network:     s.Network(),
				Success:     lo.ToPtr(true),
				Source:      protocol.SourceKurora,
				Platform:    protocol.PlatformNouns,
				Tag:         filter.TagCollectible,
				Type:        filter.CollectibleAuction,
				Transfers: []model.Transfer{
					// This is a real transfer
					transfer,
				},
			}

			transactionsMap[event.TransactionHash] = transaction

		} else {
			// append transfers
			transactionsMap[event.TransactionHash].Transfers = append(transactionsMap[event.TransactionHash].Transfers, transfer)
		}
	}

	last, err := lo.Last(events)
	if err == nil {
		cursor = kurora.LogCursor(last.TransactionHash, last.LogIndex)
		cache.Global().Set(ctx, nounsCacheKey, cursor, 0)
	}

	for _, transaction := range transactionsMap {
		transactions = append(transactions, transaction)
	}

	return transactions, cursor, nil
}

func (s *service) buildCost(ctx context.Context, network string, address common.Address, value *big.Int) (*metadata.Token, error) {
	var costToken metadata.Token

	if address == ethereum.AddressGenesis || address == element.AddressNativeToken {
		nativeToken, err := s.tokenClient.Native(ctx, network)
		if err != nil {
			return nil, err
		}

		costValue := decimal.NewFromBigInt(value, 0)
		costValueDisplay := costValue.Shift(-int32(nativeToken.Decimals))

		costToken = metadata.Token{
			Name:         nativeToken.Name,
			Symbol:       nativeToken.Symbol,
			Decimals:     nativeToken.Decimals,
			Standard:     protocol.TokenStandardNative,
			Image:        nativeToken.Logo,
			Value:        &costValue,
			ValueDisplay: &costValueDisplay,
		}
	} else {
		erc20Token, err := s.tokenClient.ERC20(ctx, network, address.String())
		if err != nil {
			return nil, err
		}

		costValue := decimal.NewFromBigInt(value, 0)
		costValueDisplay := costValue.Shift(-int32(erc20Token.Decimals))

		costToken = metadata.Token{
			Name:            erc20Token.Name,
			Symbol:          erc20Token.Symbol,
			Decimals:        erc20Token.Decimals,
			Standard:        protocol.TokenStandardERC20,
			ContractAddress: erc20Token.ContractAddress,
			Image:           erc20Token.Logo,
			Value:           &costValue,
			ValueDisplay:    &costValueDisplay,
		}
	}

	return &costToken, nil
}
