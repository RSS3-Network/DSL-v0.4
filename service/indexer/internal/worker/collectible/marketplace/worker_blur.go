package marketplace

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

	var (
		collection      common.Address
		collectionID    *big.Int
		collectionValue *big.Int
		token           common.Address
		tokenValue      *big.Int
	)

	if event.Buy.Price.Cmp(decimal.Zero.BigInt()) == 0 {
		collection = event.Sell.Collection
		collectionID = event.Sell.TokenId
		collectionValue = event.Sell.Amount
		token = event.Buy.PaymentToken
		tokenValue = event.Buy.Price
	} else {
		collection = event.Buy.Collection
		collectionID = event.Buy.TokenId
		collectionValue = event.Buy.Amount
		token = event.Buy.PaymentToken
		tokenValue = event.Sell.Price
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, collection.String(), collectionID)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", collection), zap.Stringer("token_id", collectionID))

		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	nft.SetValue(decimal.NewFromBigInt(collectionValue, 0))

	// This contract can only buy NFTs with native tokens,
	// and only one at a time
	if nft.Cost, err = i.buildCost(ctx, transaction.Network, token, tokenValue); err != nil {
		zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", collection), zap.Stringer("value", collectionID))

		return nil, fmt.Errorf("build cost: %w", err)
	}

	internalTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Maker, event.Taker, nft, nft.Cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	return []model.Transfer{
		lo.FromPtr(internalTransfer),
	}, nil
}
