package marketplace

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20/weth"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/opensea"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/quix"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"

	"go.uber.org/zap"
)

func (i *internal) handleOpenSeaOrderFulfilled(ctx context.Context, transaction model.Transaction, log *types.Log, relatedLogs []*types.Log) ([]model.Transfer, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleOpenSea")

	defer opentelemetry.Log(span, transaction, nil, nil)

	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := quix.NewSeaPortFilterer(quix.AddressSeaport, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new seaport filterer: %w", err)
	}

	event, err := filterer.ParseOrderFulfilled(*log)
	if err != nil {
		return nil, fmt.Errorf("parse order fulfilled: %w", err)
	}

	internalTransfers := make([]model.Transfer, 0)

	// In most cases the offered is the seller and the recipient are the buyer.
	// However, if the matching nft in Consideration transfers, the opposite is true.
	seller, buyer := event.Offerer, event.Recipient

	// Quix is OpenSea's Fork, but it modifies the transaction logic.
	if platform == protocol.PlatformQuix {
		if seller == quix.AddressWrapperSeaportProxy {
			seller = common.HexToAddress(transaction.AddressFrom)
		}

		if buyer == quix.AddressWrapperSeaportProxy {
			buyer = common.HexToAddress(transaction.AddressFrom)
		}
	}

	var nft, cost *metadata.Token

	for _, offer := range event.Offer {
		switch offer.ItemType {
		case opensea.ItemTypeNative, opensea.ItemTypeERC20:
			if cost, err = i.buildCost(ctx, transaction.Network, offer.Token, offer.Amount); err != nil {
				zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", offer.Token), zap.Stringer("token_id", offer.Amount))

				return nil, fmt.Errorf("build cost: %w", err)
			}
		case opensea.ItemTypeERC721, opensea.ItemTypeERC1155:
			if nft, err = i.tokenClient.NFTToMetadata(ctx, transaction.Network, offer.Token.String(), offer.Identifier); err != nil {
				zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", offer.Token), zap.Stringer("token_id", offer.Amount), zap.Stringer("value", offer.Amount))

				return nil, fmt.Errorf("nft to metadata: %w", err)
			}

			nft.SetValue(decimal.NewFromBigInt(offer.Amount, 0))
		}
	}

	for _, consideration := range event.Consideration {
		switch consideration.ItemType {
		case opensea.ItemTypeNative, opensea.ItemTypeERC20:
			// Ignore token transfers where receipt is not seller.
			if consideration.Recipient != seller {
				continue
			}

			if cost, err = i.buildCost(ctx, transaction.Network, consideration.Token, consideration.Amount); err != nil {
				zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", consideration.Token), zap.Stringer("token_id", consideration.Amount))

				return nil, fmt.Errorf("build cost: %w", err)
			}
		case opensea.ItemTypeERC721, opensea.ItemTypeERC1155:
			// Offerer is a buyer, Recipient is a seller.
			seller, buyer = buyer, seller

			if nft, err = i.tokenClient.NFTToMetadata(ctx, transaction.Network, consideration.Token.String(), consideration.Identifier); err != nil {
				zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Stringer("contract_address", consideration.Token), zap.Stringer("token_id", consideration.Amount), zap.Stringer("value", consideration.Amount))

				return nil, fmt.Errorf("nft to metadata: %w", err)
			}

			nft.SetValue(decimal.NewFromBigInt(consideration.Amount, 0))
		}
	}

	if nft == nil {
		return nil, fmt.Errorf("this is not an nft transaction")
	}

	var index int64
	for _, logForIndex := range relatedLogs {
		// If a user purchases multiple NFTs in the same transaction,
		// the log indexes will conflict and need to be matched with their transfer logs separately.
		if len(logForIndex.Topics) == 4 &&
			logForIndex.Topics[0] == erc721.EventHashTransfer &&
			strings.EqualFold(common.HexToAddress(logForIndex.Topics[1].Hex()).String(), event.Offerer.String()) &&
			strings.EqualFold(common.HexToAddress(logForIndex.Topics[2].Hex()).String(), event.Recipient.String()) &&
			strings.EqualFold(logForIndex.Topics[3].Big().String(), nft.ID) {
			index = int64(logForIndex.Index)

			break
		}
	}

	internalTransfer, err := i.buildTradeTransfer(transaction, index, platform, seller, buyer, *nft, cost)
	if err != nil {
		return nil, fmt.Errorf("build trade transfer: %w", err)
	}

	internalTransfers = append(internalTransfers, *internalTransfer)

	return internalTransfers, nil
}

func (i *internal) handleOpenSeaOrdersMatched(ctx context.Context, transaction model.Transaction, log *types.Log) ([]model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("create ethereum client: %w", err)
	}

	filterer, err := opensea.NewWyvernExchangeFilterer(opensea.AddressWyvernExchangeV1, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("create filterer: %w", err)
	}

	event, err := filterer.ParseOrdersMatched(*log)
	if err != nil {
		return nil, fmt.Errorf("parse orders matched event: %w", err)
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("unmarshal source data: %w", err)
	}

	var (
		standard    string
		transferLog *types.Log
	)

	// The value of ERC1155 is variable, while the value of ERC721 is always 1
	for _, log := range sourceData.Receipt.Logs {
		switch log.Topics[0] {
		case erc721.EventHashTransfer:
			if log.Address == weth.AddressEthereum {
				continue
			}

			standard = protocol.TokenStandardERC721
			transferLog = log
		case erc1155.EventHashTransferSingle:
			standard = protocol.TokenStandardERC1155
			transferLog = log
		default:
			continue
		}

		break
	}

	// Defensive programming, not triggered by normal conditions
	if transferLog == nil {
		return nil, errors.New("transfer log not found")
	}

	var (
		tokenID    *big.Int
		tokenValue *big.Int
	)

	switch standard {
	case protocol.TokenStandardERC721:
		erc721, err := erc721.NewERC721(transferLog.Address, ethereumClient)
		if err != nil {
			return nil, fmt.Errorf("new erc721 contract: %w", err)
		}

		event, err := erc721.ParseTransfer(*transferLog)
		if err != nil {
			return nil, fmt.Errorf("parse transfer event: %w", err)
		}

		tokenID = event.TokenId
		tokenValue = big.NewInt(1)
	case protocol.TokenStandardERC1155:
		erc1155, err := erc1155.NewERC1155(transferLog.Address, ethereumClient)
		if err != nil {
			return nil, fmt.Errorf("new erc1155 contract: %w", err)
		}

		event, err := erc1155.ParseTransferSingle(*transferLog)
		if err != nil {
			return nil, fmt.Errorf("parse transfer single event: %w", err)
		}

		tokenID = event.Id
		tokenValue = event.Value
	}

	nft, err := i.tokenClient.NFTToMetadata(ctx, transaction.Network, transferLog.Address.String(), tokenID)
	if err != nil {
		zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("standard", standard), zap.String("contract_address", transferLog.Address.String()), zap.Int64("token_id", tokenID.Int64()))

		return nil, err
	}

	nftValue := decimal.NewFromBigInt(tokenValue, 0)
	nft.Value = &nftValue

	cost, err := i.buildCost(ctx, transaction.Network, ethereum.AddressGenesis, event.Price)
	if err != nil {
		zap.L().Error("build cost", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("contract_address", ethereum.AddressGenesis.String()), zap.Int64("value", event.Price.Int64()))

		return nil, err
	}

	tradeTransfer, err := i.buildTradeTransfer(transaction, int64(log.Index), platform, event.Maker, event.Taker, *nft, cost)
	if err != nil {
		return nil, err
	}

	return []model.Transfer{
		*tradeTransfer,
	}, nil
}
