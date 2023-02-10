package transaction

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
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
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	tokenjob "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction/job/token"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"

	"go.uber.org/zap"
)

const (
	Name       = "transaction"
	threadSize = 200
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
		protocol.NetworkArbitrum,
		protocol.NetworkOptimism,
		protocol.NetworkAvalanche,
		protocol.NetworkCelo,
		protocol.NetworkFantom,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	if err := s.loadCentralizedExchangeWallets(ctx); err != nil {
		return fmt.Errorf("initialize centralized exchange wallets: %w", err)
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (transaction []model.Transaction, err error) {
	tracer := otel.Tracer("worker_token")
	ctx, handlerSpan := tracer.Start(ctx, "worker_token:Handle")

	defer handlerSpan.End()

	var internalTransactions []*model.Transaction

	opt := parallel.NewOption().WithConcurrency(200)

	switch message.Network {
	case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain, protocol.NetworkCrossbell, protocol.NetworkXDAI, protocol.NetworkOptimism, protocol.NetworkAvalanche:
		internalTransactions, err = parallel.MapWithError(transactions, s.makeEthereumHandlerFunc(ctx, message, transactions), opt)
	case protocol.NetworkZkSync:
		internalTransactions, err = parallel.MapWithError(transactions, s.makeZkSyncHandlerFunc(ctx, message, transactions), opt)
	case protocol.NetworkArweave:
		internalTransactions, err = parallel.MapWithError(transactions, s.makeArweaveHandlerFunc(ctx, message, transactions), opt)
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

	var locker sync.Mutex

	// Limit the maximum number of tasks for a single address.
	for _, transfers := range lo.Chunk(transaction.Transfers, threadSize) {
		var waitGroup sync.WaitGroup

		for _, transfer := range transfers {
			waitGroup.Add(1)

			go func(transfer model.Transfer) {
				defer waitGroup.Done()

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
					// EIP tokens
					internalTransfers, err = s.handleEthereumOriginToken(ctx, message, transaction, transfer)
				}

				if err != nil && !errors.Is(err, ErrorNativeTokenTransferValueIsZero) && !errors.Is(err, ErrorUnsupportedContractEvent) && !errors.Is(err, ErrorNonUserDirectApproval) /* Ignore actively aborted transactions */ {
					zap.L().Warn("handle ethereum origin", zap.Error(err), zap.String("network", transaction.Network), zap.String("address", message.Address), zap.String("transaction_hash", transaction.Hash))

					return
				}

				for _, internalTransfer := range internalTransfers {
					wallet, err := s.isCentralizedExchangeWallet(ctx, internalTransfer.AddressFrom, internalTransfer.AddressTo)
					if err != nil {
						zap.L().Error("check centralized exchange wallet", zap.Error(err))

						return
					}

					if wallet != nil {
						// Build deposit or withdraw transaction
						if err := s.buildCentralizedExchangeTransfer(ctx, &internalTransfer, lo.FromPtr(wallet), message.Address); err != nil {
							zap.L().Error("build centralized exchange transaction", zap.Error(err))

							return
						}

						locker.Lock()
						internalTransaction.Platform = internalTransfer.Platform
						locker.Unlock()
					}

					locker.Lock()
					internalTransaction, internalTransfer = s.buildType(internalTransaction, internalTransfer)
					internalTransaction.Transfers = append(internalTransaction.Transfers, internalTransfer)
					internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(internalTransfer.Tag, transaction.Tag, internalTransfer.Type, internalTransaction.Type)
					locker.Unlock()
				}
			}(transfer)
		}

		waitGroup.Wait()
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

	return s.buildEthereumTokenTransferMetadata(ctx, message, transaction, transfer, nil, nil, sourceData.Transaction.Value())
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
		tokenApproval bool
		tokenApproved bool
		tokenAddress  = strings.ToLower(log.Address.String())
		tokenIDs      = make([]*big.Int, 0)
		tokenValues   = make([]*big.Int, 0)

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
	case erc20.EventHashApproval:
		switch len(log.Topics) {
		case 3: // ERC-20
			filterer, err := erc20.NewERC20Filterer(log.Address, nil)
			if err != nil {
				return nil, fmt.Errorf("new erc20 filterer: %w", err)
			}

			event, err := filterer.ParseApproval(*log)
			if err != nil {
				return nil, fmt.Errorf("parse approval event: %w", err)
			}

			tokenApproval = true
			tokenValues = append(tokenValues, event.Value)
		case 4: // ERC-721
			filterer, err := erc721.NewERC721Filterer(log.Address, nil)
			if err != nil {
				return nil, fmt.Errorf("new erc20 filterer: %w", err)
			}

			event, err := filterer.ParseApproval(*log)
			if err != nil {
				return nil, fmt.Errorf("parse approval event: %w", err)
			}

			tokenApproval = true
			tokenIDs = append(tokenIDs, event.TokenId)
		}
	case erc721.EventHashApprovalForAll:
		filterer, err := erc721.NewERC721Filterer(log.Address, nil)
		if err != nil {
			return nil, fmt.Errorf("new erc721 filterer: %w", err)
		}

		event, err := filterer.ParseApprovalForAll(*log)
		if err != nil {
			return nil, fmt.Errorf("parse approval for all event: %w", err)
		}

		tokenApproval = true
		tokenApproved = event.Approved
	default:
		return nil, ErrorUnsupportedContractEvent
	}

	if tokenApproval {
		if common.HexToAddress(transaction.AddressTo) != log.Address {
			return nil, ErrorNonUserDirectApproval
		}

		switch {
		case len(tokenValues) > 0: // Approval ERC-20
			for _, tokenValue := range tokenValues {
				internalTransfer, err := s.buildEthereumTokenApprovalMetadata(ctx, transaction, transfer, &tokenAddress, nil, tokenValue, tokenApproved)
				if err != nil {
					return nil, fmt.Errorf("build ethereum token approval metadata: %w", err)
				}

				transfers = append(transfers, *internalTransfer)
			}
		case len(tokenIDs) > 0: // Approval ERC-721
			for _, tokenID := range tokenIDs {
				internalTransfer, err := s.buildEthereumTokenApprovalMetadata(ctx, transaction, transfer, &tokenAddress, tokenID, nil, tokenApproved)
				if err != nil {
					return nil, fmt.Errorf("build ethereum token approval metadata: %w", err)
				}

				transfers = append(transfers, *internalTransfer)
			}
		default: // Approval for all (ERC-721 and ERC-1155)
			internalTransfer, err := s.buildEthereumTokenApprovalMetadata(ctx, transaction, transfer, &tokenAddress, nil, nil, tokenApproved)
			if err != nil {
				return nil, fmt.Errorf("build ethereum token approval metadata: %w", err)
			}

			transfers = append(transfers, *internalTransfer)
		}
	} else {
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

			internalTransfer, err := s.buildEthereumTokenTransferMetadata(ctx, message, transaction, transfer, &tokenAddress, tokenID, tokenValue)
			if err != nil {
				return nil, err
			}

			transfers = append(transfers, *internalTransfer)
		}
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
func (s *service) buildEthereumTokenTransferMetadata(ctx context.Context, message *protocol.Message, transaction model.Transaction, transfer model.Transfer, address *string, id *big.Int, value *big.Int) (*model.Transfer, error) {
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

func (s *service) buildEthereumTokenApprovalMetadata(ctx context.Context, transaction model.Transaction, transfer model.Transfer, tokenAddress *string, tokenID *big.Int, tokenValue *big.Int, approved bool) (*model.Transfer, error) {
	_, snap := otel.Tracer("worker_token").Start(ctx, "worker_token:buildEthereumTokenApprovalMetadata")

	defer snap.End()

	var tokenMetadata *metadata.Token
	var err error

	switch {
	case tokenID == nil && tokenValue != nil: // ERC-20
		if tokenMetadata, err = s.tokenClient.ERC20ToMetadata(ctx, transaction.Network, *tokenAddress); err != nil {
			return nil, fmt.Errorf("get erc20 metadata %s: %w", *tokenAddress, err)
		}

		tokenMetadata.SetValue(decimal.NewFromBigInt(tokenValue, 0))

		if len(tokenValue.Bits()) > 0 {
			tokenMetadata.Action = filter.ActionApprove
		} else {
			tokenMetadata.Action = filter.ActionRevoke
		}

		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagTransaction, transfer.Tag, filter.TransactionApproval, transfer.Type)
	case tokenID != nil && tokenValue != nil: // ERC-721
		erc721, err := s.tokenClient.ERC721(ctx, transaction.Network, *tokenAddress, tokenID)
		if err != nil {
			return nil, fmt.Errorf("get erc-721 %s/%d: %w", *tokenAddress, tokenID, err)
		}

		nft, err := erc721.ToNFT(tokenID)
		if err != nil {
			return nil, fmt.Errorf("erc721 to nft %s/%d: %w", *tokenAddress, tokenID, err)
		}

		if tokenMetadata, err = nft.ToMetadata(); err != nil {
			return nil, fmt.Errorf("erc721 to metadata %s/%d: %w", *tokenAddress, tokenID, err)
		}

		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagCollectible, transfer.Tag, filter.CollectibleApproval, transfer.Type)

		// Not supported
		// transfer.RelatedUrls = ethereum.BuildURL(transfer.RelatedUrls, ethereum.BuildTokenURL(transaction.Network, *tokenAddress, tokenID.String())...)
	case tokenID == nil && tokenValue == nil: // ERC-721 and ERC-1155
		if tokenMetadata, err = s.tokenClient.NFTToMetadata(ctx, transaction.Network, *tokenAddress, nil); err != nil {
			return nil, fmt.Errorf("get nft metadata %s: %w", *tokenAddress, err)
		}

		if approved {
			tokenMetadata.Action = filter.ActionApprove
		} else {
			tokenMetadata.Action = filter.ActionRevoke
		}

		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagCollectible, transfer.Tag, filter.CollectibleApproval, transfer.Type)
	default:
		return nil, ErrorUnsupportedContractEvent
	}

	metadataRaw, err := json.Marshal(tokenMetadata)
	if err != nil {
		return nil, fmt.Errorf("marshal token metadata: %w", err)
	}

	transfer.Metadata = metadataRaw

	return &transfer, nil
}

func (s *service) buildZkSyncTokenMetadata(ctx context.Context, message *protocol.Message, transfer model.Transfer, tokenInfo *model.GetTokenInfo, value string) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_token")
	_, snap := tracer.Start(ctx, "worker_token:buildEthereumTokenTransferMetadata")

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
