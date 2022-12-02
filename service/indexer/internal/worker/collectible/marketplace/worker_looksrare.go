package marketplace

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/looksrare"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/shopspring/decimal"
)

func (i *internal) handleLooksRareTakerAsk(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := looksrare.NewExchangeFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new exchange filterer: %w", err)
	}

	event, err := filterer.ParseTakerAsk(*log)
	if err != nil {
		return nil, fmt.Errorf("parse taker ask: %w", err)
	}

	nftMetadata, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Collection.String(), event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	var cost *metadata.Token
	if cost, err = i.buildCost(ctx, transaction.Network, event.Currency, event.Price); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromBigInt(event.Amount, 0)
	nftMetadata.Value = &nftValue

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Maker, event.Taker, nftMetadata, cost)
	if err != nil {
		return nil, err
	}

	internalTransfers := []model.Transfer{
		*tradeTransfer,
	}

	return internalTransfers, nil
}

func (i *internal) handleLooksRareTakerBid(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := looksrare.NewExchangeFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new exchange filterer: %w", err)
	}

	event, err := filterer.ParseTakerBid(*log)
	if err != nil {
		return nil, fmt.Errorf("parse taker bid: %w", err)
	}

	nftMetadata, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Collection.String(), event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	var cost *metadata.Token
	if cost, err = i.buildCost(ctx, transaction.Network, event.Currency, event.Price); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromBigInt(event.Amount, 0)
	nftMetadata.Value = &nftValue

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Maker, event.Taker, nftMetadata, cost)
	if err != nil {
		return nil, err
	}

	return []model.Transfer{
		*tradeTransfer,
	}, nil
}
