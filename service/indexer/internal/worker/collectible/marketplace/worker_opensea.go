package marketplace

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/opensea"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
)

func (i *internal) handleOpenSea(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleOpenSea")

	defer opentelemetry.Log(span, transaction, nil, nil)

	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, err
	}

	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	seaportContract, err := opensea.NewSeaport(common.HexToAddress(transaction.AddressTo), ethclient)
	if err != nil {
		return nil, err
	}

	for _, log := range sourceData.Receipt.Logs {
		if log.Topics[0] == opensea.EventHashOrderFulfilled {
			event, err := seaportContract.ParseOrderFulfilled(*log)
			if err != nil {
				return nil, err
			}

			for _, item := range event.Offer {
				nft, err := i.tokenClient.NFT(ctx, transaction.Network, item.Token.String(), item.Identifier)
				if err != nil {
					return nil, err
				}

				tokenValue := decimal.NewFromBigInt(item.Amount, 0)

				tokenAttributes := make([]metadata.TokenAttribute, 0)

				for _, attribute := range nft.Attributes {
					tokenAttributes = append(tokenAttributes, metadata.TokenAttribute{
						TraitType: attribute.TraitType,
						Value:     attribute.Value,
					})
				}

				var costToken *metadata.Token

				for _, receivedItem := range event.Consideration {
					if receivedItem.Recipient == event.Offerer {
						if costToken, err = i.buildToken(ctx, message, receivedItem.Token, nil, receivedItem.Amount); err != nil {
							return nil, err
						}
					}
				}

				tokenMetadata, err := json.Marshal(metadata.Token{
					Name:            nft.Name,
					Collection:      nft.Collection,
					Symbol:          nft.Symbol,
					Standard:        nft.Standard,
					ContractAddress: nft.ContractAddress,
					Image:           nft.Image,
					ID:              nft.ID.String(),
					Value:           &tokenValue,
					ValueDisplay:    &tokenValue,
					Cost:            costToken,
					Description:     nft.Description,
					Attributes:      tokenAttributes,
					ExternalLink:    nft.ExternalLink,
					ExternalURL:     nft.ExternalURL,
					AnimationURL:    nft.AnimationURL,
				})
				if err != nil {
					return nil, err
				}

				var transferLogIndex int64
				for _, logForIndex := range sourceData.Receipt.Logs {
					if len(logForIndex.Topics) == 4 &&
						logForIndex.Topics[0] == erc721.EventHashTransfer &&
						strings.EqualFold(common.HexToAddress(logForIndex.Topics[1].Hex()).String(), event.Offerer.String()) &&
						strings.EqualFold(common.HexToAddress(logForIndex.Topics[2].Hex()).String(), event.Recipient.String()) &&
						strings.EqualFold(logForIndex.Topics[3].Big().String(), nft.ID.String()) {
						transferLogIndex = int64(logForIndex.Index)
					}
				}

				internalTransaction.Transfers = append(internalTransaction.Transfers, model.Transfer{
					TransactionHash: internalTransaction.Hash,
					Timestamp:       internalTransaction.Timestamp,
					BlockNumber:     big.NewInt(internalTransaction.BlockNumber),
					Tag:             filter.TagCollectible,
					Type:            filter.CollectibleTrade,
					Index:           transferLogIndex,
					AddressFrom:     strings.ToLower(event.Offerer.String()),
					AddressTo:       strings.ToLower(event.Recipient.String()),
					Metadata:        tokenMetadata,
					Network:         internalTransaction.Network,
					Platform:        opensea.PlatformOpenSea,
					Source:          ethereum.Source,
					RelatedUrls: ethereum.BuildURL(
						[]string{},
						append(ethereum.BuildTokenURL(internalTransaction.Network, nft.ContractAddress, nft.ID.String()), ethereum.BuildScanURL(internalTransaction.Network, internalTransaction.Hash))...,
					),
				})
			}
		}
	}

	if len(internalTransaction.Transfers) == 0 {
		return nil, errors.New("not found nft trade")
	}

	internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagCollectible, internalTransaction.Tag, filter.CollectibleTrade, internalTransaction.Type)
	internalTransaction.Platform = opensea.PlatformOpenSea

	return &internalTransaction, nil
}
