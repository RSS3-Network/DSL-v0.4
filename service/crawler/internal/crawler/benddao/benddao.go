package benddao

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
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/benddao"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/element"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
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

	benddaoCacheKey = "crawler:benddao"
)

func New(config *config.Config) crawler.Crawler {
	return &service{
		config:      config,
		tokenClient: token.New(),
	}
}

func (s *service) Name() string {
	return protocol.PlatformBendDAO
}

func (s *service) Network() string {
	return protocol.NetworkEthereum
}

func (s *service) Run() error {
	loggerx.Global().Info("benddao: run")

	ctx := context.Background()
	var err error

	s.kuroraClient, err = kurora.Dial(ctx, s.config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error("foundation: kurora.Dial error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))

		return err
	}

	for {
		transactions, cacheInfo, err := s.handleBendDAOEvents(ctx)
		if err != nil {
			loggerx.Global().Error("foundation: handleFoundationAuctions error", zap.Error(err))

			return err
		}

		if len(transactions) == 0 {
			time.Sleep(1 * time.Minute)

			continue
		}

		// insert db
		err = database.UpsertTransactions(ctx, transactions, true)
		if err != nil {
			continue
		}

		// set cache
		cache.Global().Set(ctx, benddaoCacheKey, cacheInfo, 7*24*time.Hour)
	}
}

type nftMetadata struct {
	Address  common.Address
	TokenId  *big.Int
	Metadata *metadata.Token
}

func (s *service) handleBendDAOEvents(ctx context.Context) ([]*model.Transaction, string, error) {
	tracer := otel.Tracer("benddao")
	_, trace := tracer.Start(ctx, "benddao:handleBendDAOEvents")
	var (
		transactions []*model.Transaction
		err          error
	)
	defer func() { opentelemetry.Log(trace, nil, transactions, err) }()

	cursor, _ := cache.Global().Get(ctx, benddaoCacheKey).Result()

	query := kurora.DatasetBendDAOEventQuery{
		Limit: lo.ToPtr(100),
	}

	if len(cursor) > 0 {
		query.Cursor = &cursor
	}

	events, err := s.kuroraClient.FetchDatasetBendDAOEvents(ctx, query)
	if err != nil {
		loggerx.Global().Error("FetchDatasetBendDAOEvents error", zap.Error(err), zap.Any("query", query))

		return nil, "", err
	}

	loggerx.Global().Info("foundation: kuroraClient FetchDatasetBendDAOEvents result", zap.Int("len", len(events)), zap.String("cursor", cursor))

	var (
		metaDataMap = make(map[string]*nftMetadata, 0)
		nftList     = make([]*nftMetadata, 0)
		locker      sync.Mutex
	)

	for _, event := range events {
		key := fmt.Sprintf("%v-%v", event.NftAddress.String(), event.NftId.BigInt())
		if _, ok := metaDataMap[key]; ok {
			continue
		}

		nft := &nftMetadata{event.NftAddress, event.NftId.BigInt(), nil}

		metaDataMap[key] = nft

		nftList = append(nftList, nft)
	}

	opt := lop.NewOption().WithConcurrency(10)
	lop.ForEach(nftList, func(nftData *nftMetadata, i int) {
		address := nftData.Address.String()
		tokenId := nftData.TokenId

		metaData, err := s.tokenClient.NFTToMetadata(ctx, s.Network(), address, tokenId)
		if err != nil {
			zap.L().Error("nft to metadata", zap.Error(err), zap.String("contract_address", address), zap.Stringer("token_id", tokenId))

			return
		}

		locker.Lock()
		metaDataMap[fmt.Sprintf("%v-%v", address, tokenId)].Metadata = metaData
		locker.Unlock()
	}, opt)

	for _, event := range events {
		key := fmt.Sprintf("%v-%v", event.NftAddress.String(), event.NftId.BigInt())
		if _, ok := metaDataMap[key]; !ok {
			continue
		}

		nftData := metaDataMap[key]

		if nftData.Metadata == nil {
			continue
		}

		nft := nftData.Metadata

		var (
			feedTag   string
			feedType  string
			addressTo string
		)

		transfers := make([]model.Transfer, 0)

		switch event.Type {
		case Auction:
			feedTag = filter.TagCollectible
			feedType = filter.CollectibleAuction

			nft.Action = event.EventType
			nft.StartTime = &event.Timestamp

			cost, err := s.buildCost(ctx, s.Network(), event.TokenAddress, event.AmountToken.BigInt())
			if err != nil {
				zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", event.TransactionHash.String()), zap.Stringer("contract_address", event.NftAddress), zap.Stringer("id", event.NftId))

				continue
			}

			nft.Cost = cost

			metadataRaw, err := json.Marshal(nft)
			if err != nil {
				return nil, cursor, fmt.Errorf("marshal metadata: %w", err)
			}

			switch event.EventType {
			case LoadAuction:
				addressTo = strings.ToLower(benddao.AddressBendDAO.String())
			case LoadLiquidate:
				addressTo = strings.ToLower(benddao.AddressLendPool.String())
			}

			transfers = append(transfers,
				model.Transfer{
					TransactionHash: strings.ToLower(event.TransactionHash.String()),
					Timestamp:       event.Timestamp,
					Index:           int64(event.LogIndex),
					AddressFrom:     strings.ToLower(event.Borrower.String()),
					AddressTo:       strings.ToLower(event.From.String()),
					Metadata:        metadataRaw,
					Network:         s.Network(),
					Platform:        s.Name(),
					Source:          protocol.SourceKurora,
					Tag:             feedTag,
					Type:            feedType,
					RelatedUrls: ethereum.BuildURL(
						[]string{fmt.Sprintf("%s/%s/%v/", "https://www.benddao.xyz/en/auctions/bid/", strings.ToLower(event.NftAddress.String()), event.NftId)},
						ethereum.BuildTokenURL(s.Network(), event.NftAddress.String(), event.NftId.String())...,
					),
				})
		case Liquidity:
			feedTag = filter.TagExchange
			feedType = filter.ExchangeLiquidity

			tokenSingle, err := s.tokenClient.ERC20(ctx, s.Network(), event.TokenAddress.String())
			if err != nil {
				zap.L().Error("get erc20 token error", zap.Error(err), zap.Stringer("token_address", event.TokenAddress))

				continue
			}

			switch event.EventType {
			case LoadCreate:
				addressTo = strings.ToLower(benddao.AddressBendDAO.String())

				// supply nft
				liquidityMetadata := metadata.Liquidity{
					Protocol: s.Name(),
					Action:   filter.ExchangeLiquiditySupply,
					Tokens:   []metadata.Token{*nft},
				}

				metadataRaw, err := json.Marshal(&liquidityMetadata)
				if err != nil {
					zap.L().Error("buildLiquidityMetadata error", zap.Error(err), zap.Stringer("nft_address", event.NftAddress))

					continue
				}

				transfers = append(transfers,
					model.Transfer{
						TransactionHash: strings.ToLower(event.TransactionHash.String()),
						Timestamp:       event.Timestamp,
						Index:           int64(event.LogIndex),
						AddressFrom:     strings.ToLower(event.From.String()),
						AddressTo:       addressTo,
						Metadata:        metadataRaw,
						Network:         s.Network(),
						Platform:        s.Name(),
						Source:          protocol.SourceKurora,
						Tag:             feedTag,
						Type:            feedType,
						RelatedUrls: ethereum.BuildURL(
							ethereum.BuildTokenURL(s.Network(), event.NftAddress.String(), event.NftId.String()),
						),
					})

				// borrow eth
				metadataRaw, err = s.buildLiquidityMetadata(ctx, filter.ExchangeLiquidityBorrow, map[*token.ERC20]*big.Int{
					tokenSingle: event.AmountToken.BigInt(),
				})
				if err != nil {
					zap.L().Error("buildLiquidityMetadata error", zap.Error(err), zap.Stringer("token_address", event.TokenAddress))

					continue
				}
				transfers = append(transfers,
					model.Transfer{
						TransactionHash: strings.ToLower(event.TransactionHash.String()),
						Timestamp:       event.Timestamp,
						Index:           int64(event.LogIndex) + 1,
						AddressFrom:     strings.ToLower(event.From.String()),
						AddressTo:       addressTo,
						Metadata:        metadataRaw,
						Network:         s.Network(),
						Platform:        s.Name(),
						Source:          protocol.SourceKurora,
						Tag:             feedTag,
						Type:            feedType,
						RelatedUrls: ethereum.BuildURL(
							ethereum.BuildTokenURL(s.Network(), event.NftAddress.String(), event.NftId.String()),
						),
					})

			case LoadRepay:
				addressTo = strings.ToLower(benddao.AddressBendDAO.String())
				metadataRaw, err := s.buildLiquidityMetadata(ctx, filter.ExchangeLiquidityRepay, map[*token.ERC20]*big.Int{
					tokenSingle: event.AmountToken.BigInt(),
				})
				if err != nil {
					zap.L().Error("buildLiquidityMetadata error", zap.Error(err), zap.Stringer("token_address", event.TokenAddress))

					continue
				}
				transfers = append(transfers,
					model.Transfer{
						TransactionHash: strings.ToLower(event.TransactionHash.String()),
						Timestamp:       event.Timestamp,
						Index:           int64(event.LogIndex),
						AddressFrom:     strings.ToLower(event.From.String()),
						AddressTo:       strings.ToLower(event.Borrower.String()),
						Metadata:        metadataRaw,
						Network:         s.Network(),
						Platform:        s.Name(),
						Source:          protocol.SourceKurora,
						Tag:             feedTag,
						Type:            feedType,
						RelatedUrls: ethereum.BuildURL(
							ethereum.BuildTokenURL(s.Network(), event.NftAddress.String(), event.NftId.String()),
						),
					})
			case LoadBorrow:
				addressTo = strings.ToLower(benddao.AddressBendDAO.String())
				metadataRaw, err := s.buildLiquidityMetadata(ctx, filter.ExchangeLiquidityBorrow, map[*token.ERC20]*big.Int{
					tokenSingle: event.AmountToken.BigInt(),
				})
				if err != nil {
					zap.L().Error("buildLiquidityMetadata error", zap.Error(err), zap.Stringer("token_address", event.TokenAddress))

					continue
				}
				transfers = append(transfers,
					model.Transfer{
						TransactionHash: strings.ToLower(event.TransactionHash.String()),
						Timestamp:       event.Timestamp,
						Index:           int64(event.LogIndex),
						AddressFrom:     strings.ToLower(event.From.String()),
						AddressTo:       strings.ToLower(event.Borrower.String()),
						Metadata:        metadataRaw,
						Network:         s.Network(),
						Platform:        s.Name(),
						Source:          protocol.SourceKurora,
						Tag:             feedTag,
						Type:            feedType,
						RelatedUrls: ethereum.BuildURL(
							ethereum.BuildTokenURL(s.Network(), event.NftAddress.String(), event.NftId.String()),
						),
					})
			}
		}

		transaction := &model.Transaction{
			BlockNumber: int64(event.BlockNumber),
			Timestamp:   event.Timestamp,
			Owner:       strings.ToLower(event.From.String()),
			Hash:        strings.ToLower(event.TransactionHash.String()),
			AddressFrom: strings.ToLower(event.From.String()),
			AddressTo:   addressTo,
			Network:     s.Network(),
			Success:     lo.ToPtr(true),
			Platform:    s.Name(),
			Source:      protocol.SourceKurora,
			Tag:         feedTag,
			Type:        feedType,
			Transfers:   transfers,
		}

		transactions = append(transactions, transaction)

	}

	last, err := lo.Last(events)
	if err == nil {
		cursor = kurora.LogCursor(last.TransactionHash, last.LogIndex)
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

func (s *service) buildLiquidityMetadata(ctx context.Context, action string, tokenMap map[*token.ERC20]*big.Int) (json.RawMessage, error) {
	liquidityMetadata := metadata.Liquidity{
		Protocol: s.Name(),
		Action:   action,
		Tokens:   make([]metadata.Token, 0),
	}

	for internalToken, value := range tokenMap {
		internalTokenValue := decimal.NewFromBigInt(value, 0)

		internalTokenDisplay := internalTokenValue.Shift(-int32(internalToken.Decimals))

		standard := protocol.TokenStandardERC20

		// Native token structure has an empty contract address
		if internalToken.ContractAddress == "" {
			standard = protocol.TokenStandardNative
		}

		liquidityMetadata.Tokens = append(liquidityMetadata.Tokens, metadata.Token{
			Name:            internalToken.Name,
			Symbol:          internalToken.Symbol,
			Decimals:        internalToken.Decimals,
			Image:           internalToken.Logo,
			Standard:        standard,
			ContractAddress: internalToken.ContractAddress,
			Value:           &internalTokenValue,
			ValueDisplay:    &internalTokenDisplay,
		})
	}

	return json.Marshal(&liquidityMetadata)
}
