package marketplace

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/blur"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

func (i *internal) handleBlurOrdersMatched(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := blur.NewExchangeFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new exchange v5 filterer: %w", err)
	}

	event, err := filterer.ParseOrdersMatched(*log)
	if err != nil {
		return nil, fmt.Errorf("parse orders mateched event: %w", err)
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Sell.Collection.String(), event.Sell.TokenId)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.Sell.Collection), zap.Stringer("token_id", event.Sell.TokenId))

		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	nft.SetValue(decimal.NewFromBigInt(event.Sell.Amount, 0))

	// This contract can only buy NFTs with native tokens,
	// and only one at a time
	if nft.Cost, err = i.buildCost(ctx, transaction.Network, event.Buy.PaymentToken, event.Buy.Price); err != nil {
		zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.Buy.PaymentToken), zap.Stringer("value", event.Buy.Amount))

		return nil, fmt.Errorf("build cost: %w", err)
	}

	internalTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Sell.Trader, event.Buy.Trader, *nft, nft.Cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	return []model.Transfer{
		lo.FromPtr(internalTransfer),
	}, nil
}
