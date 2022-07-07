package token

import (
	"context"
	"encoding/json"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/logger"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/alchemy"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token/coinmarketcap"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	Name = "token"
)

var _ worker.Worker = (*service)(nil)

type service struct {
	databaseClient *gorm.DB
	zksyncClient   *zksync.Client
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
		protocol.NetworkBinanceSmartChain,
		protocol.NetworkCrossbell,
		protocol.NetworkZkSync,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_token")
	ctx, handlerSpan := tracer.Start(ctx, "worker_token:Handle")

	defer handlerSpan.End()

	switch message.Network {
	case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.PlatfromCrossbell:
		return s.handleEthereum(ctx, message, transactions)
	case protocol.NetworkZkSync:
		return s.handleZkSync(ctx, message, transactions)
	}

	return make([]model.Transaction, 0), nil
}

func (s *service) handleEthereum(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	for _, transaction := range transactions {
		switch transaction.Source {
		case alchemy.Source:
			return s.handleAlchemy(ctx, message, transactions)
		case moralis.Source:
			return s.handleMoralis(ctx, message, transactions)
		case blockscout.Source:
			return s.handleBlockscout(ctx, message, transactions)
		case arweave.Source:
			return s.handleArweave(ctx, message, transactions)
		}
	}

	return make([]model.Transaction, 0), nil
}

func (s *service) handleAlchemy(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:handleAlchemy")

	defer snap.End()

	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		internalTransaction := transaction
		internalTransaction.Transfers = make([]model.Transfer, 0)

		for _, transfer := range transaction.Transfers {
			if transfer.Index == protocol.IndexVirtual {
				var sourceData ethereum.SourceData

				if err := json.Unmarshal(transfer.SourceData, &sourceData); err != nil {
					logger.Global().Error("failed to unmarshal source data", zap.Error(err), zap.String("source_data", string(transfer.SourceData)))

					return nil, err
				}

				internalTransfer, err := s.buildEthereumTokenMetadata(ctx, message, transfer, nil, sourceData.Transaction.Value())
				if err != nil {
					return nil, err
				}

				transfer = *internalTransfer
			} else {
				var sourceData types.Log

				if err := json.Unmarshal(transfer.SourceData, &sourceData); err != nil {
					logger.Global().Error("failed to unmarshal source data", zap.Error(err), zap.String("source_data", string(transfer.SourceData)))

					return nil, err
				}

				if !(sourceData.Topics[0] == erc20.EventHashTransfer || sourceData.Topics[0] == erc721.EventHashTransfer) {
					continue
				}

				tokenAddress := sourceData.Address.String()

				erc20Filterer, err := erc20.NewERC20Filterer(sourceData.Address, nil)
				if err != nil {
					return nil, err
				}

				event, err := erc20Filterer.ParseTransfer(sourceData)
				if err != nil {
					return nil, err
				}

				internalTransfer, err := s.buildEthereumTokenMetadata(ctx, message, transfer, &tokenAddress, event.Value)
				if err != nil {
					return nil, err
				}

				transfer = *internalTransfer
			}

			internalTransaction, transfer = s.buildType(internalTransaction, transfer)
			internalTransaction.Transfers = append(internalTransaction.Transfers, transfer)
		}

		if len(internalTransaction.Transfers) > 0 {
			internalTransactions = append(internalTransactions, internalTransaction)
		}
	}

	return internalTransactions, nil
}

func (s *service) handleMoralis(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:handleMoralis")

	defer snap.End()

	return make([]model.Transaction, 0), nil
}

func (s *service) handleBlockscout(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:handleBlockscout")

	defer snap.End()

	return make([]model.Transaction, 0), nil
}

func (s *service) handleArweave(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	return make([]model.Transaction, 0), nil
}

func (s *service) handleZkSync(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:handleZkSync")

	defer snap.End()

	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		internalTransaction := transaction
		internalTransaction.Transfers = make([]model.Transfer, 0)

		for _, transfer := range transaction.Transfers {
			var data zksync.GetTransactionData

			if err := json.Unmarshal(transfer.SourceData, &data); err != nil {
				return nil, err
			}

			tokenInfo, _, err := s.zksyncClient.GetToken(ctx, uint(data.Transaction.Operation.Token))
			if err != nil {
				logger.Global().Error("failed to get token", zap.Error(err), zap.String("token_id", strconv.Itoa(int(data.Transaction.Operation.Token))))

				return nil, err
			}

			if strings.HasPrefix(tokenInfo.Symbol, "NFT") {
				nftToken, _, err := s.zksyncClient.GetNFTToken(ctx, uint(data.Transaction.Operation.Token))
				if err != nil {
					return nil, err
				}

				// TODO ERC-1155
				internalTransfer, err := s.buildZkSyncNFTMetadata(ctx, message, transfer, nftToken, strconv.Itoa(int(*tokenInfo.ID)), "1")
				if err != nil {
					return nil, err
				}

				transfer = *internalTransfer
			} else {
				//tokenMetadata
				erc20Token, _, err := s.zksyncClient.GetToken(ctx, uint(data.Transaction.Operation.Token))
				if err != nil {
					return nil, err
				}

				internalTransfer, err := s.buildZkSyncTokenMetadata(ctx, message, transfer, erc20Token, data.Transaction.Operation.Amount)
				if err != nil {
					return nil, err
				}

				transfer = *internalTransfer
			}

			internalTransaction, transfer = s.buildType(internalTransaction, transfer)
			internalTransaction.Transfers = append(internalTransaction.Transfers, transfer)
		}

		if len(internalTransaction.Transfers) > 0 {
			internalTransactions = append(internalTransactions, internalTransaction)
		}
	}

	return internalTransactions, nil
}

func (s *service) buildEthereumTokenMetadata(ctx context.Context, message *protocol.Message, transfer model.Transfer, tokenAddress *string, value *big.Int) (*model.Transfer, error) {
	var tokenMetadata metadata.Token

	if tokenAddress == nil {
		// Native
		nativeToken := NativeTokenMap[transfer.Network]
		tokenMetadata.Name = nativeToken.Name
		tokenMetadata.Symbol = nativeToken.Symbol
		tokenMetadata.Decimals = uint8(nativeToken.Decimal)
		tokenMetadata.TokenStandard = nativeToken.Standard

		transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
	} else {
		// ERC-20 / ERC-721 / ERC-1155
		tokenInfo, err := coinmarketcap.CachedGetCoinInfo(ctx, message.Network, *tokenAddress)
		if err != nil {
			logger.Global().Warn("failed to get token info from coinmarketcap", zap.Error(err), zap.String("token_address", *tokenAddress))

			return nil, err
		}

		// TODO NFT

		tokenMetadata.Name = tokenInfo.Name
		tokenMetadata.Symbol = tokenInfo.Symbol
		tokenMetadata.Logo = tokenInfo.Logo
		tokenMetadata.Decimals = tokenInfo.Decimals
		tokenMetadata.TokenAddress = *tokenAddress
		tokenMetadata.TokenStandard = protocol.TokenStandardERC20

		transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
	}

	tokenValue := decimal.NewFromBigInt(value, 0)
	tokenMetadata.TokenValue = &tokenValue

	var err error
	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &tokenMetadata); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (s *service) buildZkSyncTokenMetadata(ctx context.Context, message *protocol.Message, transfer model.Transfer, erc20Token *model.GetTokenInfo, value string) (*model.Transfer, error) {
	var tokenMetadata metadata.Token

	tokenMetadata.Symbol = erc20Token.Symbol
	tokenMetadata.Symbol = erc20Token.Symbol
	tokenMetadata.Decimals = erc20Token.Decimals
	tokenMetadata.TokenAddress = erc20Token.Address
	tokenMetadata.TokenStandard = protocol.TokenStandardERC20

	tokenValue, err := decimal.NewFromString(value)
	if err != nil {
		return nil, err
	}

	tokenMetadata.TokenValue = &tokenValue

	transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &tokenMetadata); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (s *service) buildZkSyncNFTMetadata(ctx context.Context, message *protocol.Message, transfer model.Transfer, nftToken *model.GetNFTTokenInfo, id, value string) (*model.Transfer, error) {
	var tokenMetadata metadata.Token

	tokenMetadata.Symbol = nftToken.Symbol
	tokenMetadata.TokenAddress = nftToken.Address
	tokenMetadata.NFTMetadata = nftToken.Bytes()
	tokenID, err := decimal.NewFromString(id)
	if err != nil {
		return nil, err
	}

	tokenMetadata.TokenID = &tokenID
	// TODO ERC-1155
	tokenMetadata.TokenStandard = protocol.TokenStandardERC721

	transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &tokenMetadata); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (s *service) buildType(transaction model.Transaction, transfer model.Transfer) (model.Transaction, model.Transfer) {
	switch {
	case transfer.AddressFrom == ethereum.AddressGenesis.String():
		transfer.Type = filter.TransactionMint
		switch transfer.Tag {
		case filter.TagTransaction:
			transfer.Type = filter.TransactionMint
		case filter.TagCollectible:
			transfer.Type = filter.CollectibleMint
		}
	case transfer.AddressTo == ethereum.AddressGenesis.String():
		transfer.Type = filter.TransactionBurn
		switch transfer.Tag {
		case filter.TagTransaction:
			transfer.Type = filter.TransactionBurn
		case filter.TagCollectible:
			transfer.Type = filter.CollectibleBurn
		}
	case transfer.AddressFrom == transfer.AddressTo:
		transfer.Type = filter.TransactionCancel
	case transfer.Tag == filter.TagExchange:
		transaction.Platform = transfer.Platform
	default:
		transfer.Type = filter.TransactionTransfer
		switch transfer.Tag {
		case filter.TagTransaction:
			transfer.Type = filter.TransactionTransfer
		case filter.TagCollectible:
			transfer.Type = filter.CollectibleTrade
		}
	}

	transaction.Tag = filter.UpdateTag(transfer.Tag, transaction.Tag)

	return transaction, transfer
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		zksyncClient:   zksync.New(),
	}
}
