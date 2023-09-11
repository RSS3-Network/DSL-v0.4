package liquidity

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/aave"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/balancer"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lido"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/polygon"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/exchange/liquidity/job/polygonstaking"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

const Name = "liquidity"

var _ worker.Worker = (*internal)(nil)

type internal struct {
	tokenClient *token.Client
	redisClient *redis.Client
}

func (i *internal) Name() string {
	return Name
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
		protocol.NetworkBinanceSmartChain,
		protocol.NetworkArbitrum,
		protocol.NetworkOptimism,
		protocol.NetworkAvalanche,
		protocol.NetworkXDAI,
		protocol.NetworkCelo,
		protocol.NetworkFantom,
		protocol.NetworkBase,
	}
}

func (i *internal) Initialize(ctx context.Context) error {
	return nil
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	var (
		result    = make([]model.Transaction, 0)
		waitGroup sync.WaitGroup
		locker    sync.Mutex
	)

	for _, transaction := range transactions {
		// Unsupported Platform
		platform, exists := platformMap[common.HexToAddress(transaction.AddressTo)]
		if !exists {
			// Polygon staking validators
			exists, err := i.redisClient.HExists(ctx, polygonstaking.Key, strings.ToLower(transaction.AddressTo)).Result()
			if err != nil {
				return nil, fmt.Errorf("hash exists: %w", err)
			}

			if !exists {
				continue
			}

			platform = platformPolygonStaking
		}

		// Initialize the transaction
		transaction := transaction
		transaction.Transfers = make([]model.Transfer, 0)
		transaction.Platform = platform.Name

		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			var sourceData ethereum.SourceData
			if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
				zap.L().Warn("unmarshal source data", zap.Error(err), zap.String("source_data", string(transaction.SourceData)))

				return
			}

			for _, log := range sourceData.Receipt.Logs {
				// Filter anonymous log
				if len(log.Topics) == 0 {
					continue
				}

				var (
					transfer *model.Transfer
					err      error
				)

				switch log.Topics[0] {
				// Uniswap Protocol V2
				case uniswap.EventHashMintV2:
					transfer, err = i.handleUniswapV2Mint(ctx, message, transaction, *log, platform)
				case uniswap.EventHashBurnV2:
					transfer, err = i.handleUniswapV2Burn(ctx, message, transaction, *log, platform)
				// Uniswap Protocol V3
				case uniswap.EventHashMintV3:
					transfer, err = i.handleUniswapV3Mint(ctx, message, transaction, *log, platform)
				case uniswap.EventHashBurnV3:
					transfer, err = i.handleUniswapV3Burn(ctx, message, transaction, *log, platform)
				case uniswap.EventHashCollectV3:
					transfer, err = i.handleUniswapV3Collect(ctx, message, transaction, *log, platform)
				case aave.EventHashSupplyV2:
					transfer, err = i.handleAAVEV2Deposit(ctx, message, transaction, *log, platform)
				case aave.EventHashBorrowV2:
					transfer, err = i.handleAAVEV2Borrow(ctx, message, transaction, *log, platform)
				case aave.EventHashRepayV2:
					transfer, err = i.handleAAVEV2Repay(ctx, message, transaction, *log, platform)
				case aave.EventHashWithdrawV2:
					transfer, err = i.handleAAVEV2Withdraw(ctx, message, transaction, *log, platform)
				case aave.EventHashSupplyV3:
					transfer, err = i.handleAAVEV3Supply(ctx, message, transaction, *log, platform)
				case aave.EventHashBorrowV3:
					transfer, err = i.handleAAVEV3Borrow(ctx, message, transaction, *log, platform)
				case aave.EventHashRepayV3:
					transfer, err = i.handleAAVEV3Repay(ctx, message, transaction, *log, platform)
				case aave.EventHashWithdrawV3:
					transfer, err = i.handleAAVEV3Withdraw(ctx, message, transaction, *log, platform)
				case lido.EventHashSubmitted:
					transfer, err = i.handleLidoSubmitted(ctx, message, transaction, *log, platform)
				case lido.EventHashSubmitEvent:
					transfer, err = i.handleLidoSubmitEvent(ctx, message, transaction, *log, platform)
				case polygon.EventHashShareMinted:
					transfer, err = i.handlePolygonStakingShareMinted(ctx, message, transaction, *log, platform)
				case polygon.EventHashDelegatorClaimedRewards:
					transfer, err = i.handlePolygonStakingDelegatorClaimedRewards(ctx, message, transaction, *log, platform)
				case balancer.EventPoolBalanceChanged:
					transfer, err = i.handleBalancerPoolBalanceChanged(ctx, message, transaction, *log, platform)
				default:
					continue
				}

				if err != nil {
					zap.L().Error("handle event", zap.Error(err), zap.Stringer("topic_first", log.Topics[0]))

					return
				}

				transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagExchange, transfer.Tag, filter.ExchangeLiquidity, transfer.Type)

				transaction.Tag, transaction.Type = filter.UpdateTagAndType(filter.TagExchange, transaction.Tag, filter.ExchangeLiquidity, transaction.Type)
				transaction.Owner = transaction.AddressFrom
				transaction.Transfers = append(transaction.Transfers, *transfer)
			}

			locker.Lock()
			defer locker.Unlock()

			result = append(result, transaction)
		}()
	}

	waitGroup.Wait()

	return result, nil
}

func (i *internal) buildLiquidityMetadata(ctx context.Context, platform Platform, action string, tokenMap map[*token.ERC20]*big.Int) (json.RawMessage, error) {
	liquidityMetadata := metadata.Liquidity{
		Protocol: platform.Protocol,
		Action:   action,
		Tokens:   make([]metadata.Token, 0),
	}

	for internalToken, value := range tokenMap {
		internalTokenValue := decimal.NewFromBigInt(value, 0)

		internalTokenDisplay := internalTokenValue.Shift(-int32(internalToken.Decimals))

		standard := protocol.TokenStandardERC20

		// Native token structure has an empty contract address
		if internalToken.ContractAddress == "" {
			standard = protocol.TokenStandardNative
		}

		liquidityMetadata.Tokens = append(liquidityMetadata.Tokens, metadata.Token{
			Name:            internalToken.Name,
			Symbol:          internalToken.Symbol,
			Decimals:        internalToken.Decimals,
			Image:           internalToken.Logo,
			Standard:        standard,
			ContractAddress: internalToken.ContractAddress,
			Value:           &internalTokenValue,
			ValueDisplay:    &internalTokenDisplay,
		})
	}

	return json.Marshal(&liquidityMetadata)
}

func (i *internal) Jobs() []worker.Job {
	return []worker.Job{
		polygonstaking.New(i.redisClient),
	}
}

func New() worker.Worker {
	return &internal{
		tokenClient: token.New(),
		redisClient: cache.Global(),
	}
}
