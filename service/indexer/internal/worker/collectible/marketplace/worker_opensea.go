package marketplace

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20/weth"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/opensea"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func (i *internal) handleOpenSea(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleOpenSea")

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
		case opensea.EventHashOrderFulfilled:
			internalTransfers, err = i.handleOpenseaSeaportOrderFulfilled(ctx, transaction, *log)
			if err != nil {
				return nil, err
			}
		case opensea.EventHashOrdersMatched:
			internalTransfers, err = i.handleOpenSeaWyvernExchangeOrdersMatched(ctx, transaction, *log)
			if err != nil {
				return nil, err
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
	internalTransaction.Platform = opensea.Platform

	return &internalTransaction, nil
}

func (i *internal) handleOpenseaSeaportOrderFulfilled(ctx context.Context, transaction model.Transaction, log types.Log) ([]model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, err
	}

	seaport, err := opensea.NewSeaport(opensea.AddressSeaport, ethereumClient)
	if err != nil {
		return nil, err
	}

	event, err := seaport.ParseOrderFulfilled(log)
	if err != nil {
		return nil, err
	}

	internalTransfers := make([]model.Transfer, 0)

	for _, item := range event.Consideration {
		if item.ItemType == opensea.ItemTypeERC20 {
			continue
		}

		nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, item.Token.String(), item.Identifier)
		if err != nil {
			zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("contract_address", item.Token.String()), zap.Int64("token_id", item.Amount.Int64()))

			return nil, err
		}

		tokenValue := decimal.NewFromBigInt(item.Amount, 0)
		nft.Value = &tokenValue

		var cost *metadata.Token
		for _, offer := range event.Offer {
			if cost, err = i.buildCost(ctx, transaction.Network, offer.Token, offer.Amount); err != nil {
				zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("contract_address", offer.Token.String()), zap.Int64("token_id", offer.Amount.Int64()))

				return nil, err
			}

			break
		}

		var sourceData ethereum.SourceData
		if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
			return nil, err
		}

		var transferLogIndex int64
		for _, logForIndex := range sourceData.Receipt.Logs {
			if len(logForIndex.Topics) == 4 &&
				logForIndex.Topics[0] == erc721.EventHashTransfer &&
				strings.EqualFold(common.HexToAddress(logForIndex.Topics[1].Hex()).String(), event.Offerer.String()) &&
				strings.EqualFold(common.HexToAddress(logForIndex.Topics[2].Hex()).String(), event.Recipient.String()) &&
				strings.EqualFold(logForIndex.Topics[3].Big().String(), nft.ID) {
				transferLogIndex = int64(logForIndex.Index)

				break
			}
		}

		internalTransfer, err := i.buildTradeTransfer(transaction, transferLogIndex, opensea.Platform, event.Recipient, event.Offerer, nft, cost)
		if err != nil {
			return nil, err
		}

		internalTransfers = append(internalTransfers, *internalTransfer)
	}

	return internalTransfers, nil
}

func (i *internal) handleOpenSeaWyvernExchangeOrdersMatched(ctx context.Context, transaction model.Transaction, log types.Log) ([]model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, err
	}

	wyvernExchange, err := opensea.NewWyvernExchange(opensea.AddressWyvernExchange, ethereumClient)
	if err != nil {
		return nil, err
	}

	event, err := wyvernExchange.ParseOrdersMatched(log)
	if err != nil {
		return nil, err
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, err
	}

	var (
		standard    string
		transferLog *types.Log
	)

	for _, log := range sourceData.Receipt.Logs {
		switch log.Topics[0] {
		case erc721.EventHashTransfer:
			if log.Address == weth.AddressEthereum {
				continue
			}

			standard = protocol.TokenStandardERC721
			transferLog = log
		case erc1155.EventHashTransferSingle:
			standard = protocol.TokenStandardERC1155
			transferLog = log
		default:
			continue
		}

		break
	}

	if transferLog == nil {
		return nil, errors.New("transfer log not found")
	}

	var (
		tokenID    *big.Int
		tokenValue *big.Int
	)

	switch standard {
	case protocol.TokenStandardERC721:
		erc721, err := erc721.NewERC721(transferLog.Address, ethereumClient)
		if err != nil {
			return nil, err
		}

		event, err := erc721.ParseTransfer(*transferLog)
		if err != nil {
			return nil, err
		}

		tokenID = event.TokenId
		tokenValue = big.NewInt(1)
	case protocol.TokenStandardERC1155:
		erc1155, err := erc1155.NewERC1155(transferLog.Address, ethereumClient)
		if err != nil {
			return nil, err
		}

		event, err := erc1155.ParseTransferSingle(*transferLog)
		if err != nil {
			return nil, err
		}

		tokenID = event.Id
		tokenValue = event.Value
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, transferLog.Address.String(), tokenID)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("standard", standard), zap.String("contract_address", transferLog.Address.String()), zap.Int64("token_id", tokenID.Int64()))

		return nil, err
	}

	nftValue := decimal.NewFromBigInt(tokenValue, 0)
	nft.Value = &nftValue

	cost, err := i.buildCost(ctx, transaction.Network, ethereum.AddressGenesis, event.Price)
	if err != nil {
		zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("contract_address", ethereum.AddressGenesis.String()), zap.Int64("value", event.Price.Int64()))

		return nil, err
	}

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), opensea.Platform, event.Maker, event.Taker, nft, cost)
	if err != nil {
		return nil, err
	}

	return []model.Transfer{
		*tradeTransfer,
	}, nil
}
