package marketplace

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/looksrare"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
)

func (i *internal) handleLooksRare(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleLooksRare")

	defer opentelemetry.Log(span, transaction, nil, nil)

	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, err
	}

	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	exchangeContract, err := looksrare.NewExchange(common.HexToAddress(transaction.AddressTo), ethclient)
	if err != nil {
		return nil, err
	}

	var internalTransfers []model.Transfer

	for _, log := range sourceData.Receipt.Logs {
		switch log.Topics[0] {
		case looksrare.EventHashTakerAsk:
			if internalTransfers, err = i.handleLooksRareTakerAsk(ctx, internalTransaction, *log, exchangeContract); err != nil {
				return nil, err
			}
		case looksrare.EventHashTakerBid:
			if internalTransfers, err = i.handleLooksRareTakerBid(ctx, internalTransaction, *log, exchangeContract); err != nil {
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
	internalTransaction.Platform = looksrare.Platform

	return &internalTransaction, nil
}

func (i *internal) handleLooksRareTakerAsk(ctx context.Context, transaction model.Transaction, log types.Log, exchangeContract *looksrare.Exchange) ([]model.Transfer, error) {
	event, err := exchangeContract.ParseTakerAsk(log)
	if err != nil {
		return nil, err
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Collection.String(), event.TokenId)
	if err != nil {
		return nil, err
	}

	var cost *metadata.Token
	if cost, err = i.buildCost(ctx, transaction.Network, event.Currency, event.Price); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromBigInt(event.Amount, 0)
	nft.Value = &nftValue

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), looksrare.Platform, event.Maker, event.Taker, nft, cost)
	if err != nil {
		return nil, err
	}

	return []model.Transfer{
		*tradeTransfer,
	}, nil
}

func (i *internal) handleLooksRareTakerBid(ctx context.Context, transaction model.Transaction, log types.Log, exchangeContract *looksrare.Exchange) ([]model.Transfer, error) {
	event, err := exchangeContract.ParseTakerBid(log)
	if err != nil {
		return nil, err
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Collection.String(), event.TokenId)
	if err != nil {
		return nil, err
	}

	var cost *metadata.Token
	if cost, err = i.buildCost(ctx, transaction.Network, event.Currency, event.Price); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromBigInt(event.Amount, 0)
	nft.Value = &nftValue

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), looksrare.Platform, event.Maker, event.Taker, nft, cost)
	if err != nil {
		return nil, err
	}

	return []model.Transfer{
		*tradeTransfer,
	}, nil
}
