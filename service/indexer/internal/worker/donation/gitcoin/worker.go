package gitcoin

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/donation"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/gitcoin"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/donation/gitcoin/job"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	Name                          = protocol.PlatformGitcoin
	ContractAddressPolygon        = "0xb99080b9407436ebb2b8fe56d45ffa47e9bb8877"
	ContractAddressEth            = "0x7d655c57f71464b6f83811c55d84009cd9f5221c"
	ContractAddressEthereumNative = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
)

var _ worker.Worker = (*service)(nil)

type service struct {
	ethereumClientMap      map[string]*ethclient.Client
	tokenClient            *token.Client
	databaseClient         *gorm.DB
	redisClient            *redis.Client
	gitcoinProjectCacheKey string
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
		protocol.NetworkZkSync,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	gitcoinProjectJob := &job.GitcoinProjectJob{
		DatabaseClient:         s.databaseClient,
		RedisClient:            s.redisClient,
		GitcoinProjectCacheKey: s.gitcoinProjectCacheKey,
	}

	gitcoinProjectJob.SetCache()

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (internalTransactions []model.Transaction, err error) {
	tracer := otel.Tracer("gitcoin_worker")
	_, trace := tracer.Start(ctx, "gitcoin_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, internalTransactions, err) }()

	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		if transaction.AddressTo != ContractAddressPolygon && transaction.AddressTo != ContractAddressEth {
			internalTransactionMap[transaction.Hash] = transaction

			continue
		}

		transfers, err := s.handleGitcoin(ctx, message, transaction)
		if err != nil {
			logger.Global().Error("failed to handle gitcoin transaction", zap.Error(err))

			continue
		}

		for _, transfer := range transfers {
			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Tag, value.Type = filter.UpdateTagAndType(transfer.Tag, value.Tag, transfer.Type, value.Type)
			value.Transfers = append(value.Transfers, transfer)

			if value.Tag == filter.TagDonation {
				value.Platform = Name
			}

			value.Owner = transaction.AddressFrom

			internalTransactionMap[transaction.Hash] = value
		}
	}

	internalTransactions = make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, internalTransaction)
	}

	return internalTransactions, nil
}

func (s *service) handleGitcoin(ctx context.Context, message *protocol.Message, transaction model.Transaction) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_gitcoin")
	_, span := tracer.Start(ctx, "worker_gitcoin:handleGitcoin")

	defer opentelemetry.Log(span, transaction, transfers, err)

	// Unsupported network
	ethereumClient, exists := s.ethereumClientMap[transaction.Network]
	if !exists {
		return nil, errors.New("unsupported network")
	}

	receipt, err := ethereumClient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
	if err != nil {
		logger.Global().Error("get transaction receipt error", zap.Error(err))

		return nil, err
	}

	for _, log := range receipt.Logs {
		if log.Topics[0] != gitcoin.EventHashDonationSent {
			continue
		}

		gitcoinContract, err := gitcoin.NewGitcoin(log.Address, ethereumClient)
		if err != nil {
			logger.Global().Error("get gitcoin contract error", zap.Error(err))

			continue
		}

		event, err := gitcoinContract.ParseDonationSent(*log)
		if err != nil {
			logger.Global().Error("parse donation sent event error", zap.Error(err))

			continue
		}

		if event.Dest == ethereum.AddressGenesis {
			continue
		}

		var project donation.GitcoinProject

		if err := s.databaseClient.
			Model((*donation.GitcoinProject)(nil)).
			Where(map[string]any{
				"admin_address": strings.ToLower(event.Dest.Hex()),
			}).
			Order("id DESC").
			First(&project).
			Error; err != nil {
			logger.Global().Error("get gitcoin project error", zap.Error(err))

			continue
		}

		// TODO Always return nil in non-production environments, need to refactor code involving redis
		// projectStr, err := s.redisClient.HGet(ctx, s.gitcoinProjectCacheKey, transfer.AddressTo).Result()
		// if err != nil || len(projectStr) == 0 {
		// 	 return transfer, err
		// }
		//
		// project := &donation.GitcoinProject{}
		// if err := json.Unmarshal([]byte(projectStr), &project); err != nil {
		//	 logger.Global().Error("unmarshal gitcoin project error", zap.Error(err))
		//
		//	 return nil, err
		// }

		sourceData, err := json.Marshal(log)
		if err != nil {
			logger.Global().Error("marshal source data error", zap.Error(err))

			return nil, err
		}

		transfer := model.Transfer{
			TransactionHash: transaction.Hash,
			Timestamp:       transaction.Timestamp,
			BlockNumber:     big.NewInt(transaction.BlockNumber),
			Tag:             transaction.Tag,
			Index:           int64(log.Index),
			AddressFrom:     strings.ToLower(event.Donor.String()),
			AddressTo:       strings.ToLower(event.Dest.String()),
			Metadata:        metadata.Default,
			Network:         transaction.Network,
			Source:          ethereum.Source,
			SourceData:      sourceData,
		}

		transfer.RelatedUrls = append(transfer.RelatedUrls, "https://gitcoin.co/grants"+strconv.Itoa(project.ID)+"/"+project.Slug)

		var tokenMetadata metadata.Token

		if strings.ToLower(event.Token.String()) == ContractAddressEthereumNative {
			native, err := s.tokenClient.Native(ctx, transaction.Network)
			if err != nil {
				logger.Global().Error("get native token error", zap.Error(err))

				return nil, err
			}

			tokenMetadata = metadata.Token{
				Name:     native.Name,
				Symbol:   native.Symbol,
				Decimals: native.Decimals,
				Standard: protocol.TokenStandardNative,
				Image:    native.Logo,
			}
		} else {
			erc20, err := s.tokenClient.ERC20(ctx, transaction.Network, event.Token.String())
			if err != nil {
				logger.Global().Error("get erc20 error", zap.Error(err))

				continue
			}

			tokenMetadata = metadata.Token{
				Name:            erc20.Name,
				Symbol:          erc20.Symbol,
				Decimals:        erc20.Decimals,
				Standard:        protocol.TokenStandardERC20,
				ContractAddress: erc20.ContractAddress,
				Image:           erc20.Logo,
			}
		}

		tokenValue := decimal.NewFromBigInt(event.Amount, 0)
		tokenValueDisplay := tokenValue.Shift(-int32(tokenMetadata.Decimals))

		tokenMetadata.Value = &tokenValue
		tokenMetadata.ValueDisplay = &tokenValueDisplay

		metadataRaw, err := json.Marshal(&metadata.Donation{
			ID:          project.ID,
			Title:       project.Title,
			Description: project.Description,
			Logo:        project.Logo,
			Platform:    protocol.PlatformGitcoin,
			Token:       tokenMetadata,
		})
		if err != nil {
			logger.Global().Error("marshal metadata error", zap.Error(err))

			continue
		}

		transfer.Metadata = metadataRaw
		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagDonation, transfer.Tag, filter.DonationDonate, transfer.Type)

		if transfer.Tag == filter.TagDonation {
			transfer.Platform = Name
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&job.GitcoinProjectJob{
			DatabaseClient:         s.databaseClient,
			RedisClient:            s.redisClient,
			GitcoinProjectCacheKey: s.gitcoinProjectCacheKey,
		},
	}
}

func New(databaseClient *gorm.DB, redisClient *redis.Client, ethereumClientMap map[string]*ethclient.Client) worker.Worker {
	return &service{
		databaseClient:         databaseClient,
		redisClient:            redisClient,
		gitcoinProjectCacheKey: "gitcoin_project",
		ethereumClientMap:      ethereumClientMap,
		tokenClient:            token.New(databaseClient, ethereumClientMap),
	}
}
