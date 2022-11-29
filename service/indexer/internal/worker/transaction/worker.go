package transaction

import (
	"bytes"
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
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/worker/zksync"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	tokenjob "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction/job/token"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"

	"go.uber.org/zap"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	Name = "transaction"
)

var _ worker.Worker = (*service)(nil)

//go:embed asset/*
var asseFS embed.FS

type service struct {
	zksyncClient *zksync.Client
	tokenClient  *token.Client
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
		protocol.NetworkOptimism,
		protocol.NetworkAvalanche,
		protocol.NetworkCelo,
		protocol.NetworkFantom,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	dir := "asset/cex_wallet"
	files, err := asseFS.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, path := range files {
		records, err := readFile(dir + "/" + path.Name())
		if err != nil {
			return err
		}

		walletModels := make([]exchange.CexWallet, 0)

		for i, record := range records {
			if i == 0 {
				continue
			}

			walletModels = append(walletModels, exchange.CexWallet{
				WalletAddress: record[0],
				Name:          record[1],
				Source:        record[2],
				Network:       record[3],
			})

			allowlist.AllowList.Add(record[0], record[1])
		}

		if len(walletModels) == 0 {
			return nil
		}

		if err := database.Global().
			Model((*exchange.CexWallet)(nil)).
			Clauses(clause.OnConflict{
				DoNothing: true,
			}).
			Create(walletModels).
			Error; err != nil {
			return err
		}
	}

	return nil
}

func readFile(path string) ([][]string, error) {
	file, err := asseFS.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (transaction []model.Transaction, err error) {
	tracer := otel.Tracer("worker_token")
	ctx, handlerSpan := tracer.Start(ctx, "worker_token:Handle")

	defer handlerSpan.End()

	var internalTransactions []*model.Transaction

	opt := lop.NewOption().WithConcurrency(200)

	switch message.Network {
	case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain, protocol.NetworkCrossbell, protocol.NetworkXDAI, protocol.NetworkOptimism, protocol.NetworkAvalanche:
		internalTransactions, err = lop.MapWithError(transactions, s.makeEthereumHandlerFunc(ctx, message, transactions), opt)
	case protocol.NetworkZkSync:
		internalTransactions, err = lop.MapWithError(transactions, s.makeZkSyncHandlerFunc(ctx, message, transactions), opt)
	case arweave.Source:
		internalTransactions, err = lop.MapWithError(transactions, s.makeArweaveHandlerFunc(ctx, message, transactions), opt)
	}

	if err != nil {
		loggerx.Global().Error("worker_token:Handle", zap.Error(err))
	}

	// Remake transactions list
	transactions = make([]model.Transaction, 0, len(internalTransactions))

	for _, internalTransaction := range internalTransactions {
		if internalTransaction == nil {
			continue
		}

		if len(internalTransaction.Owner) == 0 {
			internalTransaction.Owner = message.Address
		}

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

	// The transaction has already been handled by another worker
	if transaction.Type != "" || transaction.Tag != "" {
		return &transaction, nil
	}

	// The transaction used as the return value
	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	for _, transfer := range transaction.Transfers {
		var (
			internalTransfers []model.Transfer
			err               error
		)

		if transfer.Index == protocol.IndexVirtual {
			// Native token
			if internalTransfer, err := s.handleEthereumOriginNative(ctx, message, transaction, transfer); err == nil { // Require error as nil
				internalTransfers = append(internalTransfers, *internalTransfer)
			}
		} else {
			// Contract events
			internalTransfers, err = s.handleEthereumOriginToken(ctx, message, transaction, transfer)
		}

		if err != nil && !errors.Is(err, ErrorNativeTokenTransferValueIsZero) && !errors.Is(err, ErrorUnsupportedContractEvent) /* Ignore actively aborted transactions */ {
			zap.L().Warn("handle ethereum origin", zap.Error(err), zap.String("network", transaction.Network), zap.String("address", message.Address), zap.String("transaction_hash", transaction.Hash))

			continue
		}

		for _, internalTransfer := range internalTransfers {
			var wallet exchange.CexWallet

			if err := s.checkCexWallet(ctx, message.Address, &internalTransaction, &internalTransfer, &wallet); err != nil {
				return nil, err
			}

			internalTransaction, internalTransfer = s.buildType(internalTransaction, internalTransfer)
			internalTransaction.Transfers = append(internalTransaction.Transfers, internalTransfer)
			internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(internalTransfer.Tag, internalTransaction.Tag, internalTransfer.Type, internalTransaction.Type)
		}
	}

	return s.buildCostMetadata(ctx, internalTransaction)
}

// Used to handle native token transactions, such as ETH, MATIC, BNB
func (s *service) handleEthereumOriginNative(ctx context.Context, message *protocol.Message, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	// Unmarshal transaction and receipt
	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		loggerx.Global().Error("unmarshal source data", zap.Error(err), zap.String("source_data", string(transaction.SourceData)))

		return nil, err
	}

	// Ignore native token transfers with value 0
	if sourceData.Transaction.Value().Cmp(big.NewInt(0)) == 0 {
		return nil, ErrorNativeTokenTransferValueIsZero
	}

	return s.buildEthereumTokenMetadata(ctx, message, transaction, transfer, nil, nil, sourceData.Transaction.Value())
}

// Used to handle ERC20 and NFT token transactions, such as DAI, AZUKI
func (s *service) handleEthereumOriginToken(ctx context.Context, message *protocol.Message, transaction model.Transaction, transfer model.Transfer) ([]model.Transfer, error) {
	var log *types.Log
	if err := json.Unmarshal(transfer.SourceData, &log); err != nil {
		loggerx.Global().Error("unmarshal log", zap.Error(err), zap.String("source_data", string(transfer.SourceData)))

		return nil, err
	}

	// There may be malicious or defective topics in the data
	if len(log.Topics) == 0 {
		return nil, ErrorInvalidTopicsLength
	}

	var (
		tokenAddress = strings.ToLower(log.Address.String())
		tokenIDs     = make([]*big.Int, 0)
		tokenValues  = make([]*big.Int, 0)

		transfers = make([]model.Transfer, 0)
	)

	switch log.Topics[0] {
	case erc20.EventHashTransfer, erc721.EventHashTransfer:
		if len(log.Topics) == 4 { // ERC-721 added last topic to index
			filterer, err := erc721.NewERC721Filterer(log.Address, nil) // https://eips.ethereum.org/EIPS/eip-721
			if err != nil {
				return nil, err
			}

			event, err := filterer.ParseTransfer(*log)
			if err != nil {
				return nil, err
			}

			tokenIDs = append(tokenIDs, event.TokenId)
		} else {
			filterer, err := erc20.NewERC20Filterer(log.Address, nil) // https://eips.ethereum.org/EIPS/eip-20
			if err != nil {
				return nil, err
			}

			event, err := filterer.ParseTransfer(*log)
			if err != nil {
				return nil, err
			}

			tokenValues = append(tokenValues, event.Value)
		}
	case erc1155.EventHashTransferSingle: // https://eips.ethereum.org/EIPS/eip-1155
		filterer, err := erc1155.NewERC1155Filterer(log.Address, nil)
		if err != nil {
			return nil, err
		}

		event, err := filterer.ParseTransferSingle(*log)
		if err != nil {
			return nil, err
		}

		tokenIDs = append(tokenIDs, event.Id)
		tokenValues = append(tokenValues, event.Value)
	case erc1155.EventHashTransferBatch:
		filterer, err := erc1155.NewERC1155Filterer(log.Address, nil)
		if err != nil {
			return nil, err
		}

		event, err := filterer.ParseTransferBatch(*log)
		if err != nil {
			return nil, err
		}

		if len(event.Ids) != len(event.Values) {
			return nil, ErrorInvalidContractEvent
		}

		for i, id := range event.Ids {
			tokenIDs = append(tokenIDs, id)
			tokenValues = append(tokenValues, event.Values[i])
		}
	default:
		return nil, ErrorUnsupportedContractEvent
	}

	for i := 0; i < lo.Max([]int{len(tokenIDs), len(tokenValues)}); i++ {
		var (
			tokenID    *big.Int
			tokenValue *big.Int
		)

		if len(tokenIDs) > i {
			tokenID = tokenIDs[i]
		}

		if len(tokenValues) > i {
			tokenValue = tokenValues[i]
		}

		internalTransfer, err := s.buildEthereumTokenMetadata(ctx, message, transaction, transfer, &tokenAddress, tokenID, tokenValue)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, *internalTransfer)
	}

	return transfers, nil
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
		return s.handleZkSync(ctx, message, transaction)
	}
}

func (s *service) handleZkSync(ctx context.Context, message *protocol.Message, transaction model.Transaction) (*model.Transaction, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:handleZkSync")

	defer snap.End()

	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	for _, transfer := range transaction.Transfers {
		if !(transfer.Metadata == nil || bytes.Equal(transfer.Metadata, metadata.Default)) {
			internalTransaction.Transfers = append(internalTransaction.Transfers, transfer)

			continue
		}

		var data zksync.GetTransactionData

		if err := json.Unmarshal(transfer.SourceData, &data); err != nil {
			return nil, err
		}

		tokenInfo, _, err := s.zksyncClient.GetToken(ctx, uint(data.Transaction.Operation.Token))
		if err != nil {
			loggerx.Global().Error("failed to get token", zap.Error(err), zap.String("token_id", strconv.Itoa(int(data.Transaction.Operation.Token))))

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
			internalTransfer, err := s.buildZkSyncTokenMetadata(ctx, message, transfer, tokenInfo, data.Transaction.Operation.Amount)
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

// For build metadata.Token for native and ERC-20 and ERC-721 and ERC-1155 tokens
func (s *service) buildEthereumTokenMetadata(ctx context.Context, message *protocol.Message, transaction model.Transaction, transfer model.Transfer, address *string, id *big.Int, value *big.Int) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:buildEthereumTokenMetadata")

	defer snap.End()

	var (
		tokenMetadata *metadata.Token
		err           error
	)

	switch {
	case address == nil: // Native
		tokenMetadata, err = s.tokenClient.NatvieToMetadata(ctx, message.Network)
		if err != nil {
			return nil, fmt.Errorf("native token: %w", err)
		}

		tokenMetadata.SetValue(decimal.NewFromBigInt(value, 0))

		transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
	case address != nil && id == nil: // ERC-20
		tokenMetadata, err = s.tokenClient.ERC20ToMetadata(ctx, message.Network, *address)
		if err != nil || tokenMetadata.Symbol == "" {
			return nil, fmt.Errorf("unsupported erc20 token %s: %w", *address, err)
		}

		if value == nil {
			value = big.NewInt(0)
		}

		tokenMetadata.SetValue(decimal.NewFromBigInt(value, 0))

		transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
	case address != nil && id != nil && value == nil: // ERC-721
		erc721, err := s.tokenClient.ERC721(ctx, message.Network, *address, id)
		if err != nil {
			return nil, fmt.Errorf("erc721 %s/%d: %w", *address, id, err)
		}

		nft, err := erc721.ToNFT(id)
		if err != nil {
			return nil, fmt.Errorf("erc721 to nft %s/%d: %w", *address, id, err)
		}

		if tokenMetadata, err = nft.ToMetadata(); err != nil {
			return nil, fmt.Errorf("erc721 to metadata %s/%d: %w", *address, id, err)
		}
		if err != nil {
			return nil, err
		}

		tokenMetadata.SetValue(decimal.New(1, 0))

		if transfer.AddressTo == strings.ToLower(ens.EnsRegistrarController.String()) || transfer.AddressFrom == strings.ToLower(ens.EnsRegistrarController.String()) {
			transfer.Platform = protocol.PlatformEns
		}

		transfer.Tag = filter.UpdateTag(filter.TagCollectible, transfer.Tag)
		transfer.RelatedUrls = ethereum.BuildURL(transfer.RelatedUrls, ethereum.BuildTokenURL(message.Network, *address, id.String())...)
	case address != nil && id != nil && value != nil: // ERC-1155
		erc1155, err := s.tokenClient.ERC1155(ctx, message.Network, *address, id)
		if err != nil {
			return nil, fmt.Errorf("erc1155 %s/%d: %w", *address, id, err)
		}

		nft, err := erc1155.ToNFT(id)
		if err != nil {
			return nil, fmt.Errorf("erc1155 to nft %s/%d: %w", *address, id, err)
		}

		if tokenMetadata, err = nft.ToMetadata(); err != nil {
			return nil, fmt.Errorf("erc1155 to metadata %s/%d: %w", *address, id, err)
		}

		tokenMetadata.SetValue(decimal.NewFromBigInt(value, 0))

		transfer.Tag = filter.UpdateTag(filter.TagCollectible, transfer.Tag)
		transfer.RelatedUrls = ethereum.BuildURL(transfer.RelatedUrls, ethereum.BuildTokenURL(message.Network, *address, id.String())...)
	}

	if isBurn(&transfer, tokenMetadata) {
		transfer.Type = filter.CollectibleBurn
	}

	metadataRaw, err := json.Marshal(tokenMetadata)
	if err != nil {
		return nil, err
	}

	transfer.Metadata = metadataRaw

	return &transfer, nil
}

func (s *service) buildZkSyncTokenMetadata(ctx context.Context, message *protocol.Message, transfer model.Transfer, tokenInfo *model.GetTokenInfo, value string) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:buildEthereumTokenMetadata")

	defer snap.End()

	var tokenMetadata *metadata.Token
	var err error

	if tokenInfo.Address == ethereum.AddressGenesis.String() {
		tokenMetadata, err = s.tokenClient.NatvieToMetadata(ctx, message.Network)
	} else {
		tokenMetadata, err = s.tokenClient.ERC20ToMetadata(ctx, message.Network, tokenInfo.Address)
	}
	if err != nil {
		return nil, err
	}

	tokenValue, err := decimal.NewFromString(value)
	if err != nil {
		return nil, err
	}

	tokenMetadata.SetValue(tokenValue)

	transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)

	metadataRaw, err := json.Marshal(tokenMetadata)
	if err != nil {
		return nil, err
	}

	transfer.Metadata = metadataRaw

	return &transfer, nil
}

func (s *service) buildZkSyncNFTMetadata(ctx context.Context, message *protocol.Message, transfer model.Transfer, nftToken *model.GetNFTTokenInfo, id, value string) (*model.Transfer, error) {
	var tokenMetadata metadata.Token

	tokenMetadata.Symbol = nftToken.Symbol
	tokenMetadata.ContractAddress = nftToken.Address
	tokenID := big.NewInt(0)
	tokenID.SetString(id, 0)

	tokenMetadata.ID = tokenID.String()

	// TODO ERC-1155
	tokenMetadata.Standard = protocol.TokenStandardERC721

	transfer.Tag = filter.UpdateTag(filter.TagCollectible, transfer.Tag)

	metadataRaw, err := json.Marshal(tokenMetadata)
	if err != nil {
		return nil, err
	}

	transfer.Metadata = metadataRaw

	return &transfer, nil
}

func (s *service) buildCostMetadata(ctx context.Context, transaction model.Transaction) (*model.Transaction, error) {
	if !isMint(transaction.Tag, transaction.Type) {
		return &transaction, nil
	}

	var tokenCost *metadata.Token

	for _, transfer := range transaction.Transfers {
		// Unmarshal cost token metadata
		if transfer.Index == protocol.IndexVirtual {
			if err := json.Unmarshal(transfer.Metadata, &tokenCost); err != nil {
				return nil, err
			}

			break
		}
	}

	// Free mint
	if tokenCost == nil {
		return &transaction, nil
	}

	var err error

	internalTransaction := transaction
	internalTransaction.Transfers = make([]model.Transfer, 0)

	for _, transfer := range transaction.Transfers {
		if !isMint(transfer.Tag, transfer.Type) {
			continue
		}

		var tokenOrigin metadata.Token
		if err := json.Unmarshal(transfer.Metadata, &tokenOrigin); err != nil {
			return nil, err
		}

		tokenOrigin.Cost = tokenCost

		if transfer.Metadata, err = json.Marshal(tokenOrigin); err != nil {
			return nil, err
		}

		internalTransaction.Transfers = append(internalTransaction.Transfers, transfer)
	}

	return &internalTransaction, nil
}

func (s *service) buildType(transaction model.Transaction, transfer model.Transfer) (model.Transaction, model.Transfer) {
	// transfer.Type might not be empty
	if transfer.Type == "" {
		switch {
		case transfer.AddressFrom == ethereum.AddressGenesis.String():
			transfer.Type = filter.TransactionMint
			switch transfer.Tag {
			case filter.TagTransaction:
				transfer.Type = filter.TransactionMint
			case filter.TagCollectible:
				transfer.Type = filter.CollectibleMint
			}
		case ethereum.IsBlackHoleAddress(common.HexToAddress(transfer.AddressTo)):
			transfer.Type = filter.TransactionBurn
			switch transfer.Tag {
			case filter.TagTransaction:
				transfer.Type = filter.TransactionBurn
			case filter.TagCollectible:
				transfer.Type = filter.CollectibleBurn
			}
		case transfer.Tag == filter.TagExchange:
			transaction.Platform = transfer.Platform
		default:
			transfer.Type = filter.TransactionTransfer
			switch transfer.Tag {
			case filter.TagTransaction:
				transfer.Type = filter.TransactionTransfer
			case filter.TagCollectible:
				transfer.Type = filter.CollectibleTransfer
			}
		}
	}
	transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)

	return transaction, transfer
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		tokenjob.New(),
	}
}

// Check address (from / to) is a WalletAddress. If true, update transfer
func (s *service) checkCexWallet(ctx context.Context, address string, transaction *model.Transaction, transfer *model.Transfer, wallet *exchange.CexWallet) error {
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
		err := database.Global().Model((*exchange.CexWallet)(nil)).Where("wallet_address = ?", strings.ToLower(transfer.AddressTo)).Or("wallet_address = ?", strings.ToLower(transfer.AddressFrom)).First(&wallet).Error
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

	if len(wallet.Name) == 0 {
		return nil
	}

	if transfer.Tag = filter.UpdateTag(filter.TagExchange, transfer.Tag); transfer.Tag == filter.TagExchange {
		transfer.Platform = wallet.Name
		transaction.Platform = wallet.Name

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

func isMint(actionTag, actionType string) bool {
	return (actionTag == filter.TagTransaction && actionType == filter.TransactionMint) || (actionTag == filter.TagCollectible && actionType == filter.CollectibleMint)
}

func isBurn(transfer *model.Transfer, metadata *metadata.Token) bool {
	return transfer.AddressTo == metadata.ContractAddress
}

func New() worker.Worker {
	return &service{
		zksyncClient: zksync.New(),
		tokenClient:  token.New(),
	}
}
