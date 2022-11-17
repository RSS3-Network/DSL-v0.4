package marketplace

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/quix"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"

	"go.uber.org/zap"
)

func (i *internal) handleQuix(ctx context.Context, _ *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleQuix")

	defer opentelemetry.Log(span, transaction, nil, nil)

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, err
	}

	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	for _, log := range sourceData.Receipt.Logs {
		if len(log.Topics) == 0 {
			return nil, errors.New("invalid length of topics")
		}

		var (
			internalTransfers []model.Transfer
			err               error
		)

		switch log.Topics[0] {
		case quix.EventHashOrderFulfilled:
			// Quix exchange's Seaport contract is a fork of OpenSea Seaport contract,
			// which may be reused later if there are no compatibility issues.
			internalTransfers, err = i.handleQuickSeaportOrderFulfilled(ctx, transaction, *log)
			if err != nil {
				return nil, fmt.Errorf("handle quick seaport order fulfilled: %w", err)
			}
		case quix.EventHashSellOrderFilled:
			internalTransfers, err = i.handleQuickExchangeV5SellOrderFilled(ctx, transaction, *log)
			if err != nil {
				return nil, fmt.Errorf("handle quick exchange v5 sell order filled: %w", err)
			}
		default:
			continue
		}

		internalTransaction.Transfers = append(internalTransaction.Transfers, internalTransfers...)
	}

	if len(internalTransaction.Transfers) == 0 {
		return nil, errors.New("not found nft trade")
	}

	internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagCollectible, internalTransaction.Tag, filter.CollectibleTrade, internalTransaction.Type)
	internalTransaction.Platform = quix.Platform

	return &internalTransaction, nil
}

func (i *internal) handleQuickSeaportOrderFulfilled(ctx context.Context, transaction model.Transaction, log types.Log) ([]model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := quix.NewSeaPortFilterer(quix.AddressSeaport, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new seaport filterer: %w", err)
	}

	event, err := filterer.ParseOrderFulfilled(log)
	if err != nil {
		return nil, fmt.Errorf("parse order fulfilled: %w", err)
	}

	internalTransfers := make([]model.Transfer, 0)

	for _, spentItem := range event.Offer {
		nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, spentItem.Token.String(), spentItem.Identifier)
		if err != nil {
			zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("contract_address", spentItem.Token.String()), zap.Int64("token_id", spentItem.Amount.Int64()), zap.Uint64("value", spentItem.Amount.Uint64()))

			return nil, fmt.Errorf("nft to metadata: %w", err)
		}

		nft.SetValue(decimal.NewFromBigInt(spentItem.Amount, 0))

		for _, receivedItem := range event.Consideration {
			if receivedItem.Recipient == event.Offerer {
				if nft.Cost, err = i.buildCost(ctx, transaction.Network, receivedItem.Token, receivedItem.Amount); err != nil {
					zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("contract_address", receivedItem.Token.String()), zap.Int64("token_id", receivedItem.Amount.Int64()))

					return nil, fmt.Errorf("build cost: %w", err)
				}

				break
			}
		}

		var sourceData ethereum.SourceData
		if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
			return nil, err
		}

		var transferLogIndex int64
		for _, logForIndex := range sourceData.Receipt.Logs {
			// If a user purchases multiple NFTs in the same transaction,
			// the log indexes will conflict and need to be matched with their transfer logs separately.
			if len(logForIndex.Topics) == 4 &&
				logForIndex.Topics[0] == erc721.EventHashTransfer &&
				strings.EqualFold(common.HexToAddress(logForIndex.Topics[1].Hex()).String(), event.Offerer.String()) &&
				strings.EqualFold(common.HexToAddress(logForIndex.Topics[2].Hex()).String(), event.Recipient.String()) &&
				strings.EqualFold(logForIndex.Topics[3].Big().String(), nft.ID) {
				transferLogIndex = int64(logForIndex.Index)

				break
			}
		}

		internalTransfer, err := i.buildTradeTransfer(transaction, transferLogIndex, quix.Platform, event.Recipient, event.Offerer, nft, nft.Cost)
		if err != nil {
			return nil, fmt.Errorf("build trade transfer: %w", err)
		}

		internalTransfers = append(internalTransfers, *internalTransfer)
	}

	return internalTransfers, nil
}

func (i *internal) handleQuickExchangeV5SellOrderFilled(ctx context.Context, transaction model.Transaction, log types.Log) ([]model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := quix.NewExchangeV5Filterer(quix.AddressSeaport, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new exchange v5 filterer: %w", err)
	}

	event, err := filterer.ParseSellOrderFilled(log)
	if err != nil {
		return nil, fmt.Errorf("parse sell order filled: %w", err)
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.ContractAddress.String(), event.TokenId)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("contract_address", event.ContractAddress.String()), zap.Uint64("token_id", event.TokenId.Uint64()))

		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	nft.SetValue(decimal.NewFromBigInt(event.Price, 0))

	// This contract can only buy NFTs with native tokens,
	// and only one at a time
	if nft.Cost, err = i.buildCost(ctx, transaction.Network, ethereum.AddressGenesis, event.Price); err != nil {
		zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.ContractAddress), zap.Uint64("value", event.Price.Uint64()))

		return nil, fmt.Errorf("build cost: %w", err)
	}

	internalTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), quix.Platform, event.Buyer, event.Seller, nft, nft.Cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	return []model.Transfer{
		lo.FromPtr(internalTransfer),
	}, nil
}
