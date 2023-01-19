package marketplace

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/blur"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/element"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/looksrare"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/opensea"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/quix"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/tofunft"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

var _ worker.Worker = (*internal)(nil)

const (
	SourceName = "marketplace"
	ThreadSize = 200
)

type internal struct {
	tokenClient *token.Client
}

func (i *internal) Name() string {
	return SourceName
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
		protocol.NetworkBinanceSmartChain,
		protocol.NetworkOptimism,
		protocol.NetworkAvalanche,
		protocol.NetworkXDAI,
		protocol.NetworkCelo,
		protocol.NetworkFantom,
	}
}

func (i *internal) Initialize(ctx context.Context) error {
	return nil
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	var (
		internalTransactions = make([]model.Transaction, 0)
		limiter              = make(chan struct{}, ThreadSize)

		waitGroup sync.WaitGroup
		locker    sync.Mutex
	)

	for _, transaction := range transactions {
		// Filter unsupported platforms
		platform, exists := platformMap[common.HexToAddress(transaction.AddressTo)]
		if !exists {
			continue
		}

		limiter <- struct{}{}
		waitGroup.Add(1)

		go func(transaction model.Transaction) {
			defer func() {
				<-limiter
				waitGroup.Done()
			}()

			var sourceData ethereum.SourceData

			if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
				zap.L().Warn("unmarshal source data", zap.Error(err), zap.String("transaction_hash", transaction.Hash))

				return
			}

			internalTransaction := transaction
			internalTransaction.Transfers = make([]model.Transfer, 0)

			for _, log := range sourceData.Receipt.Logs {
				var (
					internalTransfers []model.Transfer
					err               error
				)

				switch log.Topics[0] {
				case opensea.EventHashOrderFulfilled:
					internalTransfers, err = i.handleOpenSeaOrderFulfilled(ctx, transaction, log, sourceData.Receipt.Logs)
				case opensea.EventHashOrdersMatched:
					internalTransfers, err = i.handleOpenSeaOrdersMatched(ctx, transaction, log)
				case quix.EventHashSellOrderFilled:
					internalTransfers, err = i.handleQuickSellOrderFilled(ctx, transaction, log)
				case looksrare.EventHashTakerAsk:
					internalTransfers, err = i.handleLooksRareTakerAsk(ctx, transaction, log)
				case looksrare.EventHashTakerBid:
					internalTransfers, err = i.handleLooksRareTakerBid(ctx, transaction, log)
				case tofunft.EventEvInventoryUpdate:
					internalTransfers, err = i.handleTofuNFTEvInventoryUpdate(ctx, transaction, log, sourceData.Receipt.Logs)
				case blur.EventOrdersMatched:
					internalTransfers, err = i.handleBlurOrdersMatched(ctx, transaction, log)
				case element.EventERC721SellOrderFilled:
					internalTransfers, err = i.handleElementERC721SellOrderFilled(ctx, transaction, log)
				case element.EventERC1155SellOrderFilled:
					internalTransfers, err = i.handleElementERC1155SellOrderFilled(ctx, transaction, log)
				default:
					continue
				}

				if err != nil {
					zap.L().Warn("handle marketplace event", zap.Error(err))

					continue
				}

				locker.Lock()
				internalTransaction.Transfers = append(internalTransaction.Transfers, internalTransfers...)
				locker.Unlock()
			}

			locker.Lock()
			internalTransaction.Platform = platform
			internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagCollectible, internalTransaction.Tag, filter.CollectibleTrade, internalTransaction.Type)
			internalTransactions = append(internalTransactions, internalTransaction)
			locker.Unlock()
		}(transaction)
	}

	waitGroup.Wait()

	return internalTransactions, nil
}

func (i *internal) buildTradeTransfer(transaction model.Transaction, index int64, platform string, seller, buyer common.Address, nft metadata.Token, cost metadata.Token) (*model.Transfer, error) {
	if nft.Value == nil {
		nft.Value = lo.ToPtr(decimal.Zero)
	}

	nft.SetValue(*nft.Value)

	nft.Cost = &cost

	metadataRaw, err := json.Marshal(nft)
	if err != nil {
		return nil, fmt.Errorf("marshal metadata: %w", err)
	}

	transfer := model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleTrade,
		Index:           index,
		AddressFrom:     strings.ToLower(seller.String()),
		AddressTo:       strings.ToLower(buyer.String()),
		Metadata:        metadataRaw,
		Network:         transaction.Network,
		Platform:        platform,
		RelatedUrls: ethereum.BuildURL(
			[]string{
				ethereum.BuildScanURL(transaction.Network, transaction.Hash),
			},
			ethereum.BuildTokenURL(transaction.Network, nft.ContractAddress, nft.ID)...,
		),
	}

	return &transfer, nil
}

func (i *internal) buildCost(ctx context.Context, network string, address common.Address, value *big.Int) (*metadata.Token, error) {
	var costToken metadata.Token

	if address == ethereum.AddressGenesis || address == element.AddressNativeToken {
		nativeToken, err := i.tokenClient.Native(ctx, network)
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
		erc20Token, err := i.tokenClient.ERC20(ctx, network, address.String())
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

func (i *internal) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &internal{}
}
