package blend

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

	blendCacheKey = "crawler:blend"
)

func New(config *config.Config) crawler.Crawler {
	return &service{
		config:      config,
		tokenClient: token.New(),
	}
}

func (s *service) Name() string {
	return Blend
}

func (s *service) Network() string {
	return protocol.NetworkEthereum
}

func (s *service) Run() error {
	loggerx.Global().Info("blend: run")

	ctx := context.Background()
	var err error

	s.kuroraClient, err = kurora.Dial(ctx, s.config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error("blend: kurora.Dial error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))

		return err
	}

	for {
		transactions, cacheInfo, err := s.handleBlendEvents(ctx)
		if err != nil {
			loggerx.Global().Error("blend: handleBlendEvents error", zap.Error(err))

			return err
		}

		if len(transactions) == 0 {
			time.Sleep(1 * time.Minute)

			continue
		}

		// insert db
		err = database.UpsertTransactions(ctx, transactions, false)
		if err != nil {
			continue
		}

		// set cache
		cache.Global().Set(ctx, blendCacheKey, cacheInfo, 7*24*time.Hour)
	}
}

type nftMetadata struct {
	Address  common.Address
	TokenId  *big.Int
	Metadata *metadata.Token
}

func (s *service) handleBlendEvents(ctx context.Context) ([]*model.Transaction, string, error) {
	tracer := otel.Tracer("blend")
	_, trace := tracer.Start(ctx, "blend:handleBlendEvents")
	var (
		transactions []*model.Transaction
		err          error
	)

	transactionMap := make(map[string]*model.Transaction, 0)

	defer func() { opentelemetry.Log(trace, nil, transactions, err) }()

	cursor, _ := cache.Global().Get(ctx, blendCacheKey).Result()

	query := kurora.DatasetBlendEventQuery{
		Limit: lo.ToPtr(100),
		From:  lo.ToPtr(common.HexToAddress("0xe1749558e716eedc94c5651ea78d921432724cea")),
	}

	if len(cursor) > 0 {
		query.Cursor = &cursor
	}

	events, err := s.kuroraClient.FetchDatasetBlendEvents(ctx, query)
	if err != nil {
		loggerx.Global().Error("FetchDatasetBlendEvents error", zap.Error(err), zap.Any("query", query))

		return nil, "", err
	}

	loggerx.Global().Info("blend: kuroraClient FetchDatasetBlendEvents result", zap.Int("len", len(events)), zap.String("cursor", cursor))

	var (
		metaDataMap = make(map[string]*nftMetadata, 0)
		nftList     = make([]*nftMetadata, 0)
		locker      sync.Mutex
	)

	nativeToken, err := s.tokenClient.Native(context.Background(), s.Network())
	if err != nil {
		return nil, "", err
	}

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

		var (
			feedTag  string
			feedType string
		)

		transfers := make([]model.Transfer, 0)

		switch event.TopicFirst {
		case EventLoanOfferTaken, EventRefinance:
			feedTag, feedType = filter.TagExchange, filter.ExchangeLiquidity

			if event.TopicFirst == EventRefinance {
				tmpQuery := kurora.DatasetBlendEventQuery{
					Limit:  lo.ToPtr(1),
					Cursor: lo.ToPtr(kurora.LogCursor(event.TransactionHash, event.LogIndex-1)),
				}

				data, err := s.kuroraClient.FetchDatasetBlendEvents(ctx, tmpQuery)
				if err != nil {
					loggerx.Global().Error("FetchEventLoanOfferTaken error", zap.Error(err), zap.Any("query", query))

					continue
				}

				if len(data) == 1 {
					continue
				}
			}

			if _, ok := metaDataMap[key]; !ok {
				continue
			}

			nftData := metaDataMap[key]

			if nftData.Metadata == nil {
				continue
			}
			nft := nftData.Metadata

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

			transfers = s.buildTransfers(transfers, strings.ToLower(event.TransactionHash.String()), strings.ToLower(event.Borrower.String()), strings.ToLower(event.Lender.String()), feedTag, feedType, event.NftAddress.String(), event.NftId, event.Timestamp, int64(event.LogIndex), metadataRaw)

			// borrow eth
			metadataRaw, err = s.buildLiquidityMetadata(ctx, filter.ExchangeLiquidityBorrow, map[*token.Native]*big.Int{
				nativeToken: event.Amount.BigInt(),
			})
			if err != nil {
				zap.L().Error("buildLiquidityMetadata error", zap.Error(err))

				continue
			}

			transfers = s.buildTransfers(transfers, strings.ToLower(event.TransactionHash.String()), strings.ToLower(event.Lender.String()), strings.ToLower(event.Borrower.String()), feedTag, feedType, event.NftAddress.String(), event.NftId, event.Timestamp, protocol.IndexVirtual, metadataRaw)
		case EventRepay:
			feedTag, feedType = filter.TagExchange, filter.ExchangeLiquidity

			metadataRaw, err := s.buildLiquidityMetadata(ctx, filter.ExchangeLiquidityRepay, map[*token.Native]*big.Int{
				nativeToken: event.Amount.BigInt(),
			})
			if err != nil {
				zap.L().Error("buildLiquidityMetadata error", zap.Error(err))

				continue
			}

			transfers = s.buildTransfers(transfers, strings.ToLower(event.TransactionHash.String()), strings.ToLower(event.Borrower.String()), strings.ToLower(event.Lender.String()), feedTag, feedType, "", decimal.Decimal{}, event.Timestamp, int64(event.LogIndex), metadataRaw)
		case EventStartAuction:
			if _, ok := metaDataMap[key]; !ok {
				continue
			}

			nftData := metaDataMap[key]

			if nftData.Metadata == nil {
				continue
			}

			nft := nftData.Metadata

			feedTag, feedType = filter.TagCollectible, filter.CollectibleAuction

			nft.Action = filter.ActionCreate
			nft.StartTime = &event.Timestamp

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

			transfers = s.buildTransfers(transfers, strings.ToLower(event.TransactionHash.String()), strings.ToLower(event.Lender.String()), strings.ToLower(AddressBlend.String()), feedTag, feedType, event.NftAddress.String(), event.NftId, event.Timestamp, int64(event.LogIndex), metadataRaw)
		case EventSeize:
			feedTag, feedType = filter.TagCollectible, filter.CollectibleTransfer

			if _, ok := metaDataMap[key]; !ok {
				continue
			}

			nftData := metaDataMap[key]

			if nftData.Metadata == nil {
				continue
			}
			nft := nftData.Metadata

			metadataRaw, err := json.Marshal(nft)
			if err != nil {
				zap.L().Error("buildCollectibleMetadata error", zap.Error(err), zap.Stringer("hash", event.TransactionHash))

				continue
			}

			transfers = s.buildTransfers(transfers, strings.ToLower(event.TransactionHash.String()), strings.ToLower(event.Borrower.String()), strings.ToLower(event.Lender.String()), feedTag, feedType, event.NftAddress.String(), event.NftId, event.Timestamp, int64(event.LogIndex), metadataRaw)
		case EventBuyLocked:
			feedTag, feedType = filter.TagCollectible, filter.CollectibleTrade

			if _, ok := metaDataMap[key]; !ok {
				continue
			}

			nftData := metaDataMap[key]

			if nftData.Metadata == nil {
				continue
			}
			nft := nftData.Metadata

			cost := &metadata.Token{
				Name:         nativeToken.Name,
				Symbol:       nativeToken.Symbol,
				Decimals:     nativeToken.Decimals,
				Standard:     protocol.TokenStandardNative,
				Image:        nativeToken.Logo,
				Value:        lo.ToPtr(decimal.Zero),
				ValueDisplay: lo.ToPtr(decimal.Zero),
			}

			nft.Cost = cost

			nft.Cost.SetValue(event.Amount)

			metadataRaw, err := json.Marshal(nft)
			if err != nil {
				zap.L().Error("buildCollectibleMetadata error", zap.Error(err), zap.Stringer("hash", event.TransactionHash))

				continue
			}

			transfers = s.buildTransfers(transfers, strings.ToLower(event.TransactionHash.String()), strings.ToLower(event.Lender.String()), strings.ToLower(event.Borrower.String()), feedTag, feedType, event.NftAddress.String(), event.NftId, event.Timestamp, int64(event.LogIndex), metadataRaw)

		}

		transaction := &model.Transaction{
			BlockNumber: int64(event.BlockNumber),
			Timestamp:   event.Timestamp,
			Owner:       strings.ToLower(event.From.String()),
			Hash:        strings.ToLower(event.TransactionHash.String()),
			AddressFrom: strings.ToLower(event.From.String()),
			AddressTo:   strings.ToLower(event.To.String()),
			Network:     s.Network(),
			Success:     lo.ToPtr(true),
			Platform:    protocol.PlatformBlur,
			Source:      protocol.SourceKurora,
			Tag:         feedTag,
			Type:        feedType,
			Transfers:   transfers,
		}

		if _, ok := transactionMap[strings.ToLower(event.TransactionHash.String())]; !ok {
			transactionMap[strings.ToLower(event.TransactionHash.String())] = transaction
		} else {
			transactionMap[strings.ToLower(event.TransactionHash.String())].Transfers = append(transactionMap[strings.ToLower(event.TransactionHash.String())].Transfers, transfers...)
		}
	}

	last, err := lo.Last(events)
	if err == nil {
		cursor = kurora.LogCursor(last.TransactionHash, last.LogIndex)
	}

	for _, tx := range transactionMap {
		transactions = append(transactions, tx)
	}

	return transactions, cursor, nil
}

func (s *service) buildTransfers(transfers []model.Transfer, hash, from, to, feedTag, feedType, nftAddress string, nftId decimal.Decimal, timestamp time.Time, index int64, metadata json.RawMessage) []model.Transfer {
	transfer := model.Transfer{
		TransactionHash: hash,
		Timestamp:       timestamp,
		Index:           index,
		AddressFrom:     from,
		AddressTo:       to,
		Metadata:        metadata,
		Network:         s.Network(),
		Platform:        protocol.PlatformBlur,
		Source:          protocol.SourceKurora,
		Tag:             feedTag,
		Type:            feedType,
		RelatedUrls: ethereum.BuildURL(
			[]string{utils.GetTxHashURL(s.Network(), hash)},
			fmt.Sprintf("%s/%s/%v", "https://blur.io/asset/", strings.ToLower(nftAddress), nftId),
		),
	}

	transfers = append(transfers, transfer)

	return transfers
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

func (s *service) buildLiquidityMetadata(ctx context.Context, action string, tokenMap map[*token.Native]*big.Int) (json.RawMessage, error) {
	liquidityMetadata := metadata.Liquidity{
		Protocol: s.Name(),
		Action:   action,
		Tokens:   make([]metadata.Token, 0),
	}

	for internalToken, value := range tokenMap {
		internalTokenValue := decimal.NewFromBigInt(value, 0)

		internalTokenDisplay := internalTokenValue.Shift(-int32(internalToken.Decimals))

		liquidityMetadata.Tokens = append(liquidityMetadata.Tokens, metadata.Token{
			Name:         internalToken.Name,
			Symbol:       internalToken.Symbol,
			Decimals:     internalToken.Decimals,
			Image:        internalToken.Logo,
			Standard:     protocol.TokenStandardNative,
			Value:        &internalTokenValue,
			ValueDisplay: &internalTokenDisplay,
		})
	}

	return json.Marshal(&liquidityMetadata)
}
