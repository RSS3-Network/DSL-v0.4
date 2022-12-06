package marketplace

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/quix"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

func (i *internal) handleQuickSellOrderFilled(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := quix.NewExchangeV5Filterer(quix.AddressSeaport, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new exchange v5 filterer: %w", err)
	}

	event, err := filterer.ParseSellOrderFilled(*log)
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

	internalTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Buyer, event.Seller, nft, nft.Cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	return []model.Transfer{
		lo.FromPtr(internalTransfer),
	}, nil
}
