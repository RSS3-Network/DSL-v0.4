package marketplace

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/x2y2"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

func (i *internal) handleX2Y2EvInventory(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := x2y2.NewExchangeFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new exchange filterer: %w", err)
	}

	event, err := filterer.ParseEvInventory(*log)
	if err != nil {
		return nil, fmt.Errorf("parse ev inventory: %w", err)
	}

	cost, err := i.buildCost(ctx, transaction.Network, event.Currency, event.Detail.Price)
	if err != nil {
		return nil, fmt.Errorf("build cost: %w", err)
	}

	dataIndex := len(event.Item.Data) - 32

	tokenAddress := common.BytesToAddress(event.Item.Data[dataIndex-32 : dataIndex])
	tokenID := new(big.Int).SetBytes(event.Item.Data[dataIndex:])

	nftMetadata, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, tokenAddress.String(), tokenID)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", tokenAddress), zap.Stringer("token_id", tokenID))

		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	// Unsupported ERC-1155
	nftMetadata.SetValue(decimal.NewFromInt(1))

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(event.Raw.Index), platform, event.Maker, event.Taker, nftMetadata, cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	internalTransfers := []model.Transfer{
		*tradeTransfer,
	}

	return internalTransfers, nil
}
