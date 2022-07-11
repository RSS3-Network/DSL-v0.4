package transaction

import (
	"context"
	"embed"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/mrc20"
	"github.com/naturalselectionlabs/pregod/common/datasource/nft"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/common/worker/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction/coinmarketcap"
	lop "github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	Name = "token"
)

var _ worker.Worker = (*service)(nil)

//go:embed asset/*
var asseFS embed.FS

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
		protocol.NetworkXDAI,
		protocol.NetworkZkSync,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	dir := "asset/cex_wallet"
	files, err := asseFS.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, path := range files {
		file, err := asseFS.Open(dir + "/" + path.Name())
		if err != nil {
			return err
		}

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			return err
		}

		file.Close()

		wallentModels := make([]exchange.CexWallet, 0)

		for i, record := range records {
			if i == 0 {
				continue
			}

			wallentModels = append(wallentModels, exchange.CexWallet{
				WalletAddress: record[0],
				Name:          record[1],
				Source:        record[2],
				Network:       record[3],
			})
		}

		if len(wallentModels) == 0 {
			return nil
		}

		if err := s.databaseClient.
			Model((*exchange.CexWallet)(nil)).
			Clauses(clause.OnConflict{
				DoNothing: true,
			}).
			Create(wallentModels).
			Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (transaction []model.Transaction, err error) {
	tracer := otel.Tracer("worker_token")
	ctx, handlerSpan := tracer.Start(ctx, "worker_token:Handle")

	defer handlerSpan.End()

	var internalTransactions []*model.Transaction

	switch message.Network {
	case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain, protocol.NetworkCrossbell, protocol.NetworkXDAI:
		internalTransactions, err = lop.MapWithError(transactions, s.makeEthereumHandlerFunc(ctx, message, transactions))
	case protocol.NetworkZkSync:
		internalTransactions, err = lop.MapWithError(transactions, s.makeZkSyncHandlerFunc(ctx, message, transactions))
	case arweave.Source:
		internalTransactions, err = lop.MapWithError(transactions, s.makeArweaveHandlerFunc(ctx, message, transactions))
	}

	if err != nil {
		logger.Global().Error("worker_token:Handle", zap.Error(err))
	}

	transactions = make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactions {
		transactions = append(transactions, *internalTransaction)
	}

	return transactions, nil
}

func (s *service) makeEthereumHandlerFunc(ctx context.Context, message *protocol.Message, transactions []model.Transaction) func(transaction model.Transaction, i int) (*model.Transaction, error) {
	return func(transaction model.Transaction, i int) (*model.Transaction, error) {
		return s.handleEthereumOrigin(ctx, message, transaction)
	}
}

func (s *service) handleEthereumOrigin(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:handleEthereumOrigin")

	defer snap.End()

	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	for _, transfer := range transaction.Transfers {
		if transfer.Index == protocol.IndexVirtual {
			var sourceData ethereum.SourceData

			if err := json.Unmarshal(transfer.SourceData, &sourceData); err != nil {
				logger.Global().Error("failed to unmarshal source data", zap.Error(err), zap.String("source_data", string(transfer.SourceData)))

				return nil, err
			}

			internalTransfer, err := s.buildEthereumTokenMetadata(ctx, message, transfer, nil, nil, sourceData.Transaction.Value())
			if err != nil {
				continue
			}

			transfer = *internalTransfer
		} else {
			var sourceData types.Log

			if err := json.Unmarshal(transfer.SourceData, &sourceData); err != nil {
				logger.Global().Error("failed to unmarshal source data", zap.Error(err), zap.String("source_data", string(transfer.SourceData)))

				return nil, err
			}

			var (
				tokenValue   *big.Int
				tokenID      *big.Int
				tokenAddress = strings.ToLower(sourceData.Address.String())
			)

			switch sourceData.Topics[0] {
			case erc20.EventHashTransfer, erc721.EventHashTransfer:
				switch len(sourceData.Topics) {
				case 3:
					filterer, err := erc20.NewERC20Filterer(sourceData.Address, nil)
					if err != nil {
						return nil, err
					}

					event, err := filterer.ParseTransfer(sourceData)
					if err != nil {
						return nil, err
					}

					tokenValue = event.Value
				case 4:
					filterer, err := erc721.NewERC721Filterer(sourceData.Address, nil)
					if err != nil {
						return nil, err
					}

					event, err := filterer.ParseTransfer(sourceData)
					if err != nil {
						return nil, err
					}

					tokenID = event.TokenId
				}
			case mrc20.EventHashLogTransfer:
				filterer, err := mrc20.NewMRC20Filterer(sourceData.Address, nil)
				if err != nil {
					return nil, err
				}

				event, err := filterer.ParseLogTransfer(sourceData)
				if err != nil {
					return nil, err
				}

				tokenValue = event.Amount
			case erc1155.EventHashTransferSingle:
				filterer, err := erc1155.NewERC1155Filterer(sourceData.Address, nil)
				if err != nil {
					return nil, err
				}

				event, err := filterer.ParseTransferSingle(sourceData)
				if err != nil {
					return nil, err
				}

				tokenID = event.Id
				tokenValue = event.Value
			default:
				continue
			}

			internalTransfer, err := s.buildEthereumTokenMetadata(ctx, message, transfer, &tokenAddress, tokenID, tokenValue)
			if err != nil {
				return nil, err
			}

			transfer = *internalTransfer
		}

		var wallet exchange.CexWallet

		if err := s.checkCexWallet(ctx, message.Address, &transfer, &wallet); err != nil {
			return nil, err
		}

		internalTransaction, transfer = s.buildType(internalTransaction, transfer)
		internalTransaction.Transfers = append(internalTransaction.Transfers, transfer)
		internalTransaction.Tag = filter.UpdateTag(transfer.Tag, internalTransaction.Tag)
	}

	return &internalTransaction, nil
}

func (s *service) makeArweaveHandlerFunc(ctx context.Context, message *protocol.Message, transactions []model.Transaction) func(transaction model.Transaction, i int) (*model.Transaction, error) {
	return func(transaction model.Transaction, i int) (*model.Transaction, error) {
		return s.handleArweave(ctx, message, transaction)
	}
}

func (s *service) handleArweave(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:handleArweave")

	defer snap.End()

	return &transaction, nil
}

func (s *service) makeZkSyncHandlerFunc(ctx context.Context, message *protocol.Message, transactions []model.Transaction) func(transaction model.Transaction, i int) (*model.Transaction, error) {
	return func(transaction model.Transaction, i int) (*model.Transaction, error) {
		return s.handleArweave(ctx, message, transaction)
	}
}

func (s *service) handleZkSync(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:handleZkSync")

	defer snap.End()

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

	return &internalTransaction, nil
}

func (s *service) buildEthereumTokenMetadata(ctx context.Context, message *protocol.Message, transfer model.Transfer, address *string, id *big.Int, value *big.Int) (*model.Transfer, error) {
	var tokenMetadata metadata.Token

	if address == nil {
		// Native
		nativeToken := NativeTokenMap[message.Network]
		tokenMetadata.Name = nativeToken.Name
		tokenMetadata.Symbol = nativeToken.Symbol
		tokenMetadata.Decimals = uint8(nativeToken.Decimal)
		tokenMetadata.TokenStandard = nativeToken.Standard
		tokenValue := decimal.NewFromBigInt(value, 0)
		tokenMetadata.TokenValue = &tokenValue

		transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
	} else {
		tokenInfo, err := coinmarketcap.CachedGetCoinInfo(ctx, message.Network, *address)
		if err == nil && tokenInfo.Symbol != "" {
			// Common ERC-20
			tokenMetadata.Name = tokenInfo.Name
			tokenMetadata.Symbol = tokenInfo.Symbol
			tokenMetadata.Logo = tokenInfo.Logo
			tokenMetadata.Decimals = tokenInfo.Decimals
			tokenMetadata.TokenStandard = protocol.TokenStandardERC20
			tokenValue := decimal.NewFromBigInt(value, 0)
			tokenMetadata.TokenValue = &tokenValue

			transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
		} else {
			// Uncommon ERC-20
			if id == nil {
				return &transfer, nil
			}

			// ERC-721 / ERC-1155
			nftMetadata, err := nft.GetMetadata(message.Network, common.HexToAddress(*address), id)
			if err != nil {
				return &transfer, nil
			}

			tokenMetadata.NFTMetadata = nftMetadata

			tokenID := decimal.NewFromBigInt(id, 0)
			tokenMetadata.TokenID = &tokenID

			var tokenValue decimal.Decimal

			if value == nil {
				tokenValue = decimal.NewFromInt(1)
				tokenMetadata.TokenStandard = protocol.TokenStandardERC721
			} else {
				tokenValue = decimal.NewFromBigInt(value, 0)
				tokenMetadata.TokenStandard = protocol.TokenStandardERC1155
			}

			tokenMetadata.TokenValue = &tokenValue

			transfer.Tag = filter.UpdateTag(filter.TagCollectible, transfer.Tag)

			transfer.RelatedUrls = append(transfer.RelatedUrls, ethereum.BuildTokenURL(message.Network, *address, id.String()))
		}

		tokenMetadata.TokenAddress = *address
	}

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
		transfer.Type = filter.TransactionSelf
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

// Check address (from / to) is a WalletAddress. If true, update transfer
func (s *service) checkCexWallet(ctx context.Context, address string, transfer *model.Transfer, wallet *exchange.CexWallet) error {
	// get from redis cache (to)
	exists, err := cache.GetMsgPack(ctx, keyOfCheckCexWallet(strings.ToLower(transfer.AddressTo)), wallet)
	if err != nil {
		return err
	}

	if !exists { // get from redis cache (from)
		if exists, err = cache.GetMsgPack(ctx, keyOfCheckCexWallet(strings.ToLower(transfer.AddressFrom)), wallet); err != nil {
			return nil
		}
	}

	if !exists {
		err := s.databaseClient.Model((*exchange.CexWallet)(nil)).Where("wallet_address = ?", strings.ToLower(transfer.AddressTo)).Or("wallet_address = ?", strings.ToLower(transfer.AddressFrom)).First(&wallet).Error
		switch {
		case err == nil: // exists, set WalletAddress' cache
			if err := cache.SetMsgPack(ctx, keyOfCheckCexWallet(wallet.WalletAddress), wallet, 7*24*time.Hour); err != nil {
				return err
			}
		case errors.Is(err, gorm.ErrRecordNotFound): // not exists, set `from` and `to` address' cache (empty)
			if err := cache.SetMsgPack(ctx, keyOfCheckCexWallet(transfer.AddressFrom), wallet, 7*24*time.Hour); err != nil {
				return err
			}
			if err := cache.SetMsgPack(ctx, keyOfCheckCexWallet(transfer.AddressTo), wallet, 7*24*time.Hour); err != nil {
				return err
			}
		default: // other err
			return err
		}
	}

	if wallet.Name != "" {
		transfer.Tag = filter.UpdateTag(filter.TagExchange, transfer.Tag)
		transfer.Platform = wallet.Name

		if transfer.AddressTo == address {
			transfer.Type = filter.ExchangeWithdraw
		} else if transfer.AddressFrom == address {
			transfer.Type = filter.ExchangeDeposit
		}
	}

	return nil
}

func keyOfCheckCexWallet(address string) string {
	return fmt.Sprintf("check_exchange_wallet.%s", address)
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		zksyncClient:   zksync.New(),
	}
}