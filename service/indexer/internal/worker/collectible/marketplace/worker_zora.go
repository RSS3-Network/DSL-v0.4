package marketplace

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/zora"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

func (i *internal) handleZoraAskFilled(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := zora.NewZoraFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new zora filterer: %w", err)
	}

	event, err := filterer.ParseAskFilled(*log)
	if err != nil {
		return nil, fmt.Errorf("parse ask filled: %w", err)
	}

	nftMetadata, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.TokenContract.String(), event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	if nftMetadata == nil {
		zap.L().Error("nft metadata not found", zap.String("network", transaction.Network), zap.String("transaction_hash", transaction.Hash))

		return nil, fmt.Errorf("nft metadata not found")
	}

	var cost *metadata.Token
	if cost, err = i.buildCost(ctx, transaction.Network, event.Ask.AskCurrency, event.Ask.AskPrice); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromInt(1)
	nftMetadata.Value = &nftValue

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Ask.Seller, event.Buyer, *nftMetadata, cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	internalTransfers := []model.Transfer{
		*tradeTransfer,
	}

	return internalTransfers, nil
}
