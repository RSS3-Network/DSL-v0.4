package marketplace

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/nswap"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func (i *internal) handleNSwapEventMatch(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	filterer, err := nswap.NewNSwapExchangeFilterer(log.Address, nil)
	if err != nil {
		return nil, fmt.Errorf("new NSwapExchange filterer: %w", err)
	}

	event, err := filterer.ParseEventMatch(*log)
	if err != nil {
		return nil, fmt.Errorf("parse EventMatch: %w", err)
	}

	nftMetadata, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, event.NftAssetType.TokenAddress.String(), event.NftAssetType.TokenID)
	if err != nil {
		return nil, fmt.Errorf("nft to metadata: %w", err)
	}

	var cost *metadata.Token
	if cost, err = i.buildCost(ctx, transaction.Network, event.PriceAssetType.TokenAddress, event.PriceFilled); err != nil {
		return nil, err
	}

	nftMetadata.Cost = cost

	nftValue := decimal.NewFromBigInt(event.NftFilled, 0)
	nftMetadata.Value = &nftValue

	internalTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.SellMaker, event.BuyMaker, *nftMetadata, nftMetadata.Cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	return []model.Transfer{
		lo.FromPtr(internalTransfer),
	}, nil
}
