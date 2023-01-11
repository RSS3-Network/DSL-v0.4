package marketplace

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/element"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"

	"go.opentelemetry.io/otel"
)

func (i *internal) handleElementERC721SellOrderFilled(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleElementERC721SellOrderFilled")

	defer opentelemetry.Log(span, transaction, nil, nil)

	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := element.NewERC721OrdersFeatureFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new erc1155 orders feature filterer: %w", err)
	}

	event, err := filterer.ParseERC721SellOrderFilled(*log)
	if err != nil {
		return nil, fmt.Errorf("parse erc1155 buy order fulfilled: %w", err)
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Erc721Token.String(), event.Erc721TokenId)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.Erc721Token), zap.Stringer("token_id", event.Erc721TokenId))

		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	nft.SetValue(decimal.NewFromInt(1))

	// This contract can only buy NFTs with native tokens,
	// and only one at a time
	if nft.Cost, err = i.buildCost(ctx, transaction.Network, event.Erc20Token, event.Erc20TokenAmount); err != nil {
		zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.Erc20Token), zap.Stringer("value", event.Erc20TokenAmount))

		return nil, fmt.Errorf("build cost: %w", err)
	}

	internalTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Maker, event.Taker, *nft, *nft.Cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	return []model.Transfer{
		lo.FromPtr(internalTransfer),
	}, nil
}

func (i *internal) handleElementERC1155SellOrderFilled(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleElementERC1155SellOrderFilled")

	defer opentelemetry.Log(span, transaction, nil, nil)

	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := element.NewERC1155OrdersFeatureFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new erc1155 orders feature filterer: %w", err)
	}

	event, err := filterer.ParseERC1155SellOrderFilled(*log)
	if err != nil {
		return nil, fmt.Errorf("parse erc1155 buy order fulfilled: %w", err)
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Erc1155Token.String(), event.Erc1155TokenId)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.Erc1155Token), zap.Stringer("token_id", event.Erc1155TokenId))

		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	nft.SetValue(decimal.NewFromBigInt(event.Erc1155FillAmount, 0))
	nft.Standard = protocol.TokenStandardERC1155

	// This contract can only buy NFTs with native tokens,
	// and only one at a time
	if nft.Cost, err = i.buildCost(ctx, transaction.Network, event.Erc20Token, event.Erc20FillAmount); err != nil {
		zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", event.Erc20Token), zap.Stringer("value", event.Erc20FillAmount))

		return nil, fmt.Errorf("build cost: %w", err)
	}

	internalTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Maker, event.Taker, *nft, *nft.Cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	return []model.Transfer{
		lo.FromPtr(internalTransfer),
	}, nil
}
