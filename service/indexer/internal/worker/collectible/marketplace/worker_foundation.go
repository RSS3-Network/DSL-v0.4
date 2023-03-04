package marketplace

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/foundation"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

func (i *internal) handleFoundationOfferAccept(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := foundation.NewFoundationFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new foundation filterer: %w", err)
	}

	event, err := filterer.ParseOfferAccepted(*log)
	if err != nil {
		return nil, fmt.Errorf("parse offer accept event: %w", err)
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.NftContract.String(), event.TokenId)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.NftContract), zap.Stringer("token_id", event.TokenId))

		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	nft.SetValue(decimal.NewFromBigInt(big.NewInt(1), 0))

	// This contract can only buy NFTs with native tokens,
	// and only one at a time
	if nft.Cost, err = i.buildCost(ctx, transaction.Network, ethereum.AddressGenesis, event.CreatorRev); err != nil {
		zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.NftContract), zap.Stringer("value", event.TokenId))

		return nil, fmt.Errorf("build cost: %w", err)
	}

	internalTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Seller, event.Buyer, *nft, nft.Cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	return []model.Transfer{
		lo.FromPtr(internalTransfer),
	}, nil
}

func (i *internal) handleFoundationBuyAccept(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := foundation.NewFoundationFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new foundation filterer: %w", err)
	}

	event, err := filterer.ParseBuyPriceAccepted(*log)
	if err != nil {
		return nil, fmt.Errorf("parse buy price accept event: %w", err)
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.NftContract.String(), event.TokenId)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.NftContract), zap.Stringer("token_id", event.TokenId))

		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	nft.SetValue(decimal.NewFromBigInt(big.NewInt(1), 0))

	// This contract can only buy NFTs with native tokens,
	// and only one at a time
	if nft.Cost, err = i.buildCost(ctx, transaction.Network, ethereum.AddressGenesis, event.CreatorRev); err != nil {
		zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.NftContract), zap.Stringer("value", event.TokenId))

		return nil, fmt.Errorf("build cost: %w", err)
	}

	internalTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Seller, event.Buyer, *nft, nft.Cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	return []model.Transfer{
		lo.FromPtr(internalTransfer),
	}, nil
}
