package marketplace

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/zerox"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/shopspring/decimal"
)

func (i *internal) handleZeroXERC721OrderFilled(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := zerox.NewERC721OrdersFeatureFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new erc721 orders feature filterer: %w", err)
	}

	event, err := filterer.ParseERC721OrderFilled(*log)
	if err != nil {
		return nil, fmt.Errorf("parse erc 721 order filled event: %w", err)
	}

	nftMetadata, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Erc721Token.String(), event.Erc721TokenId)
	if err != nil {
		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	var cost *metadata.Token
	if cost, err = i.buildCost(ctx, transaction.Network, event.Erc20Token, event.Erc20TokenAmount); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromInt(1)
	nftMetadata.Value = &nftValue

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Maker, event.Taker, nftMetadata, cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	internalTransfers := []model.Transfer{
		*tradeTransfer,
	}

	return internalTransfers, nil
}

func (i *internal) handleZeroXERC1155OrderFilled(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := zerox.NewERC1155OrdersFeature(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new erc721 orders feature filterer: %w", err)
	}

	event, err := filterer.ParseERC1155OrderFilled(*log)
	if err != nil {
		return nil, fmt.Errorf("parse erc 1155 order filled event: %w", err)
	}

	nftMetadata, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Erc1155Token.String(), event.Erc1155TokenId)
	if err != nil {
		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	var cost *metadata.Token
	if cost, err = i.buildCost(ctx, transaction.Network, event.Erc20Token, event.Erc20FillAmount); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromBigInt(event.Erc1155FillAmount, 0)
	nftMetadata.Value = &nftValue

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Maker, event.Taker, nftMetadata, cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	internalTransfers := []model.Transfer{
		*tradeTransfer,
	}

	return internalTransfers, nil
}
