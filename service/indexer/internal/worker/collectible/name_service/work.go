package name_service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens"
	ensClient "github.com/naturalselectionlabs/pregod/common/datasource/subgraph/ens"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

var _ worker.Worker = (*internal)(nil)

const (
	ThreadSize = 200
)

type internal struct {
	tokenClient *token.Client
	ensClient   ensClient.Client
}

func (i *internal) Name() string {
	return filter.TagCollectible
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
	}
}

func (i *internal) Initialize(ctx context.Context) error {
	return nil
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	var (
		internalTransactions = make([]model.Transaction, 0)
		limiter              = make(chan struct{}, ThreadSize)

		waitGroup sync.WaitGroup
		locker    sync.Mutex

		handlerMap = map[common.Address]map[common.Hash]func(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform string) (*model.Transfer, error){
			ens.EnsRegistrarController: {
				ens.EventNameRenewed:    i.handleENSNameRenewed,
				ens.EventNameRegistered: i.handleENSNameRegistered,
			},
			ens.EnsRegistrarControllerV2: {
				ens.EventNameRenewed:      i.handleENSNameRenewedV2,
				ens.EventNameRegisteredV2: i.handleENSNameRegisteredV2,
			},
			ens.EnsNameWrapper: {
				ens.EventNameWrapper: i.handleENSNameWrapper,
			},
			ens.EnsPublicResolverV1: {
				ens.EventTextChangedV1: i.handleENSTextChangedV1,
			},
			ens.EnsPublicResolverV2: {
				ens.EventTextChangedV2: i.handleENSTextChangedV2,
			},
		}
	)

	for _, transaction := range transactions {
		if transaction.Type != "" && transaction.Tag != "" {
			continue
		}

		limiter <- struct{}{}
		waitGroup.Add(1)

		go func(transaction model.Transaction) {
			defer func() {
				<-limiter
				waitGroup.Done()
			}()

			var sourceData ethereum.SourceData

			if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
				zap.L().Warn("unmarshal source data", zap.Error(err), zap.String("source_data", string(transaction.SourceData)))

				return
			}

			internalTransaction := transaction
			internalTransaction.Transfers = make([]model.Transfer, 0)

			for _, log := range sourceData.Receipt.Logs {
				// Filter anonymous log
				if len(log.Topics) == 0 {
					continue
				}

				var (
					internalTransfer *model.Transfer
					err              error
				)

				handler, ok := handlerMap[log.Address][log.Topics[0]]
				if !ok {
					continue
				}

				platform, exists := platformMap[log.Address]
				if !exists {
					continue
				}

				internalTransfer, err = handler(ctx, message, transaction, *log, platform)

				if err != nil {
					zap.L().Error("handle event", zap.Error(err), zap.Stringer("topic_first", log.Topics[0]))

					continue
				}

				locker.Lock()
				internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
				internalTransaction.Type = internalTransfer.Type
				internalTransaction.Tag = internalTransfer.Tag
				internalTransaction.Platform = platform
				locker.Unlock()
			}

			locker.Lock()
			if len(internalTransaction.Transfers) > 0 {
				internalTransactions = append(internalTransactions, internalTransaction)
			}
			locker.Unlock()
		}(transaction)
	}

	waitGroup.Wait()

	return internalTransactions, nil
}

func (i *internal) buildNameServiceMetadata(ctx context.Context, network, name, action string, tokenType *token.Native, cost, expires *big.Int, key, value string) (json.RawMessage, error) {
	nameServiceMetadata := metadata.NameService{
		Action: action,
		Name:   name,
	}

	label := strings.Split(name, ".eth")[0]
	tokenId := common.BytesToHash(crypto.Keccak256([]byte(label))).Big()
	tokenMetadata, _ := i.buildToken(ctx, network, tokenId)
	nameServiceMetadata.Detail = tokenMetadata

	if expires != nil {
		nameServiceMetadata.Expiry = lo.ToPtr(time.Unix(expires.Int64(), 0))
	}

	if tokenType != nil && cost != nil {
		internalTokenValue := decimal.NewFromBigInt(cost, 0)

		internalTokenDisplay := internalTokenValue.Shift(-int32(tokenType.Decimals))

		standard := protocol.TokenStandardNative

		nameServiceMetadata.Cost = &metadata.Token{
			Name:         tokenType.Name,
			Symbol:       tokenType.Symbol,
			Decimals:     tokenType.Decimals,
			Image:        tokenType.Logo,
			Standard:     standard,
			Value:        &internalTokenValue,
			ValueDisplay: &internalTokenDisplay,
		}
	}

	if key != "" {
		nameServiceMetadata.Key = key
		nameServiceMetadata.Value = value
	}

	return json.Marshal(&nameServiceMetadata)
}

func (i *internal) buildToken(ctx context.Context, network string, tokenId *big.Int) (*metadata.Token, error) {
	var tokenMetadata *metadata.Token

	erc721, err := i.tokenClient.ERC721(ctx, network, ens.ENSContractAddress.String(), tokenId)
	if err != nil {
		return nil, fmt.Errorf("erc721 %s/%d: %w", ens.ENSContractAddress.String(), tokenId, err)
	}

	nft, err := erc721.ToNFT(tokenId)
	if err != nil {
		return nil, fmt.Errorf("erc721 to nft %s/%d: %w", ens.ENSContractAddress.String(), tokenId, err)
	}

	if tokenMetadata, err = nft.ToMetadata(); err != nil {
		return nil, fmt.Errorf("erc721 to metadata %s/%d: %w", ens.ENSContractAddress.String(), tokenId, err)
	}
	if err != nil {
		return nil, err
	}

	tokenMetadata.SetValue(decimal.New(1, 0))

	return tokenMetadata, nil
}

func (i *internal) buildTokenMetadata(ctx context.Context, network string, tokenId *big.Int) (json.RawMessage, error) {
	tokenMetadata, err := i.buildToken(ctx, network, tokenId)
	if err != nil {
		return nil, err
	}

	metadataRaw, err := json.Marshal(tokenMetadata)
	if err != nil {
		return nil, err
	}

	return metadataRaw, nil
}

func (i *internal) Jobs() []worker.Job {
	return []worker.Job{}
}

func New() worker.Worker {
	return &internal{
		tokenClient: token.New(),
		ensClient:   ensClient.New(),
	}
}
