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
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/opensea"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ worker.Worker = (*internal)(nil)

const (
	SourceName = "marketplace"
)

type internal struct {
	tokenClient *token.Client
}

func (i *internal) Name() string {
	return SourceName
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon,
	}
}

func (i *internal) Initialize(ctx context.Context) error {
	return nil
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		addressTo := common.HexToAddress(transaction.AddressTo)

		switch addressTo {
		case opensea.AddressSeaport:
			internalTransaction, err := i.handleOpenSea(ctx, message, transaction)
			if err != nil {
				zap.L().Error("failed to handle opensea transaction", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

				continue
			}

			internalTransactions = append(internalTransactions, *internalTransaction)
		case looksrare.AddressExchange:
			internalTransaction, err := i.handleLooksRare(ctx, message, transaction)
			if err != nil {
				zap.L().Error("failed to handle looksrare transaction", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

				continue
			}

			internalTransactions = append(internalTransactions, *internalTransaction)
		default:
			continue
		}
	}

	return internalTransactions, nil
}

func (i *internal) handleOpenSea(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleOpenSea")

	defer opentelemetry.Log(span, transaction, nil, nil)

	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}
	receipt, err := ethclient.TransactionReceipt(context.Background(), common.HexToHash(transaction.Hash))
	if err != nil {
		return nil, err
	}

	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	seaportContract, err := opensea.NewSeaport(common.HexToAddress(transaction.AddressTo), ethclient)
	if err != nil {
		return nil, err
	}

	for _, log := range receipt.Logs {
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

				internalTransaction.Transfers = append(internalTransaction.Transfers, model.Transfer{
					TransactionHash: internalTransaction.Hash,
					Timestamp:       internalTransaction.Timestamp,
					BlockNumber:     big.NewInt(internalTransaction.BlockNumber),
					Tag:             filter.TagCollectible,
					Type:            filter.CollectibleTrade,
					Index:           int64(log.Index),
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

func (i *internal) handleLooksRare(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_marketplace")
	_, span := tracer.Start(ctx, "worker_marketplace:handleLooksRare")

	defer opentelemetry.Log(span, transaction, nil, nil)

	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	receipt, err := ethclient.TransactionReceipt(context.Background(), common.HexToHash(transaction.Hash))
	if err != nil {
		return nil, err
	}

	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	exchangeContract, err := looksrare.NewExchange(common.HexToAddress(transaction.AddressTo), ethclient)
	if err != nil {
		return nil, err
	}

	var internalTransfer *model.Transfer

	for _, log := range receipt.Logs {
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

func (i *internal) buildToken(ctx context.Context, message *protocol.Message, address common.Address, id *big.Int, value *big.Int) (*metadata.Token, error) {
	var costToken metadata.Token

	if address == ethereum.AddressGenesis {
		nativeToken, err := i.tokenClient.Native(ctx, message.Network)
		if err != nil {
			return nil, err
		}

		costValue := decimal.NewFromBigInt(value, 0)
		costValueDisplay := costValue.Shift(-int32(nativeToken.Decimals))

		costToken = metadata.Token{
			Name:         nativeToken.Name,
			Symbol:       nativeToken.Symbol,
			Decimals:     nativeToken.Decimals,
			Standard:     protocol.TokenStandardNative,
			Image:        nativeToken.Logo,
			Value:        &costValue,
			ValueDisplay: &costValueDisplay,
		}
	} else {
		erc20Token, err := i.tokenClient.ERC20(ctx, message.Network, address.String())
		if err != nil {
			return nil, err
		}

		costValue := decimal.NewFromBigInt(value, 0)
		costValueDisplay := costValue.Shift(-int32(erc20Token.Decimals))

		costToken = metadata.Token{
			Name:            erc20Token.Name,
			Symbol:          erc20Token.Symbol,
			Decimals:        erc20Token.Decimals,
			Standard:        protocol.TokenStandardERC20,
			ContractAddress: erc20Token.ContractAddress,
			Image:           erc20Token.Logo,
			Value:           &costValue,
			ValueDisplay:    &costValueDisplay,
		}
	}

	return &costToken, nil
}

func (i *internal) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &internal{}
}
