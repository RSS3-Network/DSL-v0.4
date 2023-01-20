package marketplace

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/tofunft"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

func (i *internal) handleTofuNFTEvInventoryUpdate(ctx context.Context, transaction model.Transaction, log *types.Log, relatedLogs []*types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := tofunft.NewMarketplaceFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new exchange filterer: %w", err)
	}

	event, err := filterer.ParseEvInventoryUpdate(*log)
	if err != nil {
		return nil, fmt.Errorf("parse taker ask: %w", err)
	}

	var nftMetadata *metadata.Token

	for _, relatedLog := range relatedLogs {
		if relatedLog.Topics[0] == erc721.EventHashTransfer && len(relatedLog.Topics) == 4 {
			filterer, err := erc721.NewERC721Filterer(relatedLog.Address, ethereumClient)
			if err != nil {
				return nil, fmt.Errorf("new erc721 filterer: %w", err)
			}

			event, err := filterer.ParseTransfer(*relatedLog)
			if err != nil {
				return nil, fmt.Errorf("parse transfer: %w", err)
			}

			if nftMetadata, err = i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.Raw.Address.String(), event.TokenId); err != nil {
				return nil, fmt.Errorf("nft to metadata: %w", err)
			}

			break
		}
	}

	if nftMetadata == nil {
		zap.L().Error("nft metadata not found", zap.String("network", transaction.Network), zap.String("transaction_hash", transaction.Hash))

		return nil, fmt.Errorf("nft metadata not found")
	}

	var cost *metadata.Token
	if cost, err = i.buildCost(ctx, transaction.Network, event.Inventory.Currency, event.Inventory.Price); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromInt(1)
	nftMetadata.Value = &nftValue

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Inventory.Seller, event.Inventory.Buyer, *nftMetadata, cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	internalTransfers := []model.Transfer{
		*tradeTransfer,
	}

	return internalTransfers, nil
}
