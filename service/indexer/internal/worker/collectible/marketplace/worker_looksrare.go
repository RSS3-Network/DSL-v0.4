package marketplace

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/looksrare"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
)

func (i *internal) handleLooksRare(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleLooksRare")

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

	exchangeContract, err := looksrare.NewExchange(common.HexToAddress(transaction.AddressTo), ethclient)
	if err != nil {
		return nil, err
	}

	var internalTransfer *model.Transfer

	for _, log := range sourceData.Receipt.Logs {
		switch log.Topics[0] {
		case looksrare.EventHashTakerAsk:
			internalTransfer, err = i.handleLooksRareTakerAsk(ctx, message, internalTransaction, *log, exchangeContract)
		case looksrare.EventHashTakerBid:
			internalTransfer, err = i.handleLooksRareTakerBid(ctx, message, internalTransaction, *log, exchangeContract)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
	}

	if len(internalTransaction.Transfers) == 0 {
		return nil, errors.New("not found nft trade")
	}

	internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagCollectible, internalTransaction.Tag, filter.CollectibleTrade, internalTransaction.Type)
	internalTransaction.Platform = looksrare.PlatformLooksrare

	return &internalTransaction, nil
}

func (i *internal) handleLooksRareTakerAsk(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, exchangeContract *looksrare.Exchange) (*model.Transfer, error) {
	event, err := exchangeContract.ParseTakerAsk(log)
	if err != nil {
		return nil, err
	}

	nft, err := i.tokenClient.NFT(ctx, transaction.Network, event.Collection.String(), event.TokenId)
	if err != nil {
		return nil, err
	}

	var costToken *metadata.Token

	if costToken, err = i.buildToken(ctx, message, event.Currency, nil, event.Price); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromBigInt(event.Amount, 0)

	var tokenAttributes []metadata.TokenAttribute

	for _, attribute := range nft.Attributes {
		tokenAttributes = append(tokenAttributes, metadata.TokenAttribute{
			TraitType: attribute.TraitType,
			Value:     attribute.Value,
		})
	}

	tokenMetadata, err := json.Marshal(metadata.Token{
		Name:            nft.Name,
		Collection:      nft.Collection,
		Symbol:          nft.Symbol,
		Standard:        nft.Standard,
		ContractAddress: nft.ContractAddress,
		Image:           nft.Image,
		ID:              nft.ID.String(),
		Value:           &nftValue,
		ValueDisplay:    &nftValue,
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

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleTrade,
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(event.Taker.String()),
		AddressTo:       strings.ToLower(event.Maker.String()),
		Metadata:        tokenMetadata,
		Network:         transaction.Network,
		Platform:        looksrare.PlatformLooksrare,
		Source:          ethereum.Source,
		RelatedUrls: ethereum.BuildURL(
			[]string{},
			append(ethereum.BuildTokenURL(transaction.Network, nft.ContractAddress, nft.ID.String()), ethereum.BuildScanURL(transaction.Network, transaction.Hash))...,
		),
	}, nil
}

func (i *internal) handleLooksRareTakerBid(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, exchangeContract *looksrare.Exchange) (*model.Transfer, error) {
	event, err := exchangeContract.ParseTakerBid(log)
	if err != nil {
		return nil, err
	}

	nft, err := i.tokenClient.NFT(ctx, transaction.Network, event.Collection.String(), event.TokenId)
	if err != nil {
		return nil, err
	}

	var costToken *metadata.Token

	if costToken, err = i.buildToken(ctx, message, event.Currency, nil, event.Price); err != nil {
		return nil, err
	}

	nftValue := decimal.NewFromBigInt(event.Amount, 0)

	var tokenAttributes []metadata.TokenAttribute

	for _, attribute := range nft.Attributes {
		tokenAttributes = append(tokenAttributes, metadata.TokenAttribute{
			TraitType: attribute.TraitType,
			Value:     attribute.Value,
		})
	}

	tokenMetadata, err := json.Marshal(metadata.Token{
		Name:            nft.Name,
		Collection:      nft.Collection,
		Symbol:          nft.Symbol,
		Standard:        nft.Standard,
		ContractAddress: nft.ContractAddress,
		Image:           nft.Image,
		ID:              nft.ID.String(),
		Value:           &nftValue,
		ValueDisplay:    &nftValue,
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

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleTrade,
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(event.Taker.String()),
		AddressTo:       strings.ToLower(event.Maker.String()),
		Metadata:        tokenMetadata,
		Network:         transaction.Network,
		Platform:        looksrare.PlatformLooksrare,
		Source:          ethereum.Source,
		RelatedUrls: ethereum.BuildURL(
			[]string{},
			append(ethereum.BuildTokenURL(transaction.Network, nft.ContractAddress, nft.ID.String()), ethereum.BuildScanURL(transaction.Network, transaction.Hash))...,
		),
	}, nil
}
