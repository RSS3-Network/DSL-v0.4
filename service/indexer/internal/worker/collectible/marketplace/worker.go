package marketplace

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/looksrare"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/opensea"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shopspring/decimal"
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
		case opensea.AddressSeaport, opensea.AddressWyvernExchange:
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

func (i *internal) buildTradeTransfer(transaction model.Transaction, index int64, platform string, seller, buyer common.Address, nft *metadata.Token, cost *metadata.Token) (*model.Transfer, error) {
	nft.Cost = cost

	metadataRaw, err := json.Marshal(nft)
	if err != nil {
		return nil, err
	}

	transfer := model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleTrade,
		Index:           index,
		AddressFrom:     strings.ToLower(seller.String()),
		AddressTo:       strings.ToLower(buyer.String()),
		Metadata:        metadataRaw,
		Network:         transaction.Network,
		Platform:        platform,
		RelatedUrls:     ethereum.BuildTokenURL(transaction.Network, nft.ContractAddress, nft.ID),
	}

	return &transfer, nil
}

func (i *internal) buildCost(ctx context.Context, network string, address common.Address, value *big.Int) (*metadata.Token, error) {
	var costToken metadata.Token

	if address == ethereum.AddressGenesis {
		nativeToken, err := i.tokenClient.Native(ctx, network)
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
		erc20Token, err := i.tokenClient.ERC20(ctx, network, address.String())
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
