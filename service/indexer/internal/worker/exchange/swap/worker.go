package swap

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	Name = "swap"
)

type Router struct {
	Name     string
	Protocol string
}

var (
	routerUniswapV2 = Router{
		Name:     protocol.PlatformUniswap,
		Protocol: "UniSwapV2",
	}

	routerUniswapV3 = Router{
		Name:     protocol.PlatformUniswap,
		Protocol: "UniSwapV2",
	}

	routerSushiSwap = Router{
		Name:     protocol.PlatformSushiswap,
		Protocol: "UniSwapV3",
	}

	pancakeSwap = Router{
		Name:     protocol.PlatformPancakeswap,
		Protocol: "UniSwapV2",
	}

	routerMap = map[string]Router{
		// Uniswap V2
		strings.ToLower("0xf164fC0Ec4E93095b804a4795bBe1e041497b92a"): routerUniswapV2, // Uniswap V2 1
		strings.ToLower("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"): routerUniswapV2, // Uniswap V2 2
		// Uniswap V3
		// https://docs.uniswap.org/protocol/reference/deployments
		strings.ToLower("0xE592427A0AEce92De3Edee1F18E0157C05861564"): routerUniswapV3, // Uniswap V3 1
		strings.ToLower("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"): routerUniswapV3, // Uniswap V3 2
		// SushiSwap
		// https://docs.sushi.com/docs/Developers/Deployment%20Addresses
		strings.ToLower("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"): routerSushiSwap, // SushiSwap Ethereum
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushiSwap, // SushiSwap Polygon
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushiSwap, // SushiSwap Binance Smart Chain
		// PancakeSwap
		// https://docs.pancakeswap.finance/code/smart-contracts/pancakeswap-exchange/router-v2
		strings.ToLower("0x10ED43C718714eb63d5aA57B78B54704E256024E"): pancakeSwap, // PancakeSwap V2
	}

	ErrorMetadataDoesNotHaveTokenField   = errors.New("metadata doesn't have transaction field")
	ErrorRouterTransferIsNotYetSupported = errors.New("router transfer isn't yet supported")
)

var _ worker.Worker = &service{}

type service struct {
	employer          *shedlock.Employer
	databaseClient    *gorm.DB
	ethereumClientMap map[string]*ethclient.Client
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	job := &Job{
		databaseClient: s.databaseClient,
	}

	var count int64

	if err := s.databaseClient.
		Model((*exchange.SwapPool)(nil)).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	return job.Run(func(ctx context.Context, duration time.Duration) error {
		return s.employer.Renewal(ctx, job.Name(), duration)
	})
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("worker_swap")
	_, trace := tracer.Start(ctx, "worker_swap:Handle")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	switch message.Network {
	case protocol.NetworkZkSync:
		return s.handleZkSync(ctx, message, transactions)
	default:
		return s.handleEthereum(ctx, message, transactions)
	}
}

func (s *service) handleEthereum(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("worker_swap")
	_, trace := tracer.Start(ctx, "worker_swap:handleEthereum")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		address := message.Address

		// Exclude transfers to self
		if transaction.AddressTo == address {
			continue
		}

		// Handle swap entry
		_, exist := routerMap[transaction.AddressTo]
		if !exist {
			continue
		}

		if transaction.Transfers, err = s.handleEthereumTransaction(ctx, message, transaction); err != nil {
			continue
		}
	}

	// Lay the map flat
	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (s *service) handleEthereumTransaction(ctx context.Context, message *protocol.Message, transaction model.Transaction) ([]model.Transfer, error) {
	var (
		internalTransfer = make([]model.Transfer, 0)
		//poolMap          = map[common.Address]*uniswap.PoolV3{}
		//poolTransferMap  = map[common.Address][]model.Transfer{}
		//poolTokenMap = map[common.Address]map[common.Address]*big.Int{}
	)

	for _, transfer := range transaction.Transfers {
		if transfer.Index == protocol.IndexVirtual {
			internalTransfer = append(internalTransfer, transfer)

			continue
		}

		var sourceData types.Log

		if err := json.Unmarshal(transfer.SourceData, &sourceData); err != nil {
			logger.Global().Error("failed to unmarshal source data", zap.Error(err))

			continue
		}

		if sourceData.Topics[0] != uniswap.EventHashTransfer {
			continue
		}

		//var poolAddress common.Address
		//
		//switch message.Address {
		//case transfer.AddressFrom:
		//	poolAddress = common.HexToAddress(transfer.AddressTo)
		//case transfer.AddressTo:
		//	poolAddress = common.HexToAddress(transfer.AddressFrom)
		//default:
		//	continue
		//}

		//uniswapPoolV3, err := uniswap.NewPoolV3(poolAddress, s.ethereumClientMap[message.Network])
		//if err != nil {
		//	logger.Global().Error("failed to create uniswap pool v3", zap.Error(err))
		//
		//	continue
		//}

		//poolMap[poolAddress] = uniswapPoolV3
		//_, exist := poolTransferMap[poolAddress]
		//if !exist {
		//	poolTransferMap[poolAddress] = make([]model.Transfer, 0)
		//}
		//
		//poolTransferMap[poolAddress] = append(poolTransferMap[poolAddress], transfer)
	}

	//for poolAddress, tokenMap := range poolTokenMap {
	//	for tokenAddress, tokenValue := range tokenMap {
	//		tokenContract, err := erc20.NewERC20(tokenAddress, s.ethereumClientMap[message.Network])
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		internalTransfer = append(internalTransfer, model.Transfer{
	//			TransactionHash: transaction.Hash,
	//			Timestamp:       transaction.Timestamp,
	//			BlockNumber:     big.NewInt(transaction.BlockNumber),
	//			Index:           0,
	//			AddressFrom:     message.Address,
	//			// TODO Recipient
	//			AddressTo:   message.Address,
	//			Metadata:    nil,
	//			Network:     "",
	//			Platform:    "",
	//			Source:      "",
	//			SourceData:  nil,
	//			RelatedUrls: nil,
	//			CreatedAt:   time.Time{},
	//			UpdatedAt:   time.Time{},
	//		})
	//	}
	//}

	return internalTransfer, nil
}

func (s *service) handleZkSync(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	// TODO Not yet supported

	return nil, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&Job{
			databaseClient: s.databaseClient,
		},
	}
}

func New(config *configx.RPC, employer *shedlock.Employer, databaseClient *gorm.DB) (worker.Worker, error) {
	var err error

	svc := service{
		ethereumClientMap: make(map[string]*ethclient.Client),
		employer:          employer,
		databaseClient:    databaseClient,
	}

	if svc.ethereumClientMap[protocol.NetworkEthereum], err = ethclient.Dial(config.General.Ethereum.HTTP); err != nil {
		return nil, err
	}

	if svc.ethereumClientMap[protocol.NetworkPolygon], err = ethclient.Dial(config.General.Polygon.HTTP); err != nil {
		return nil, err
	}

	if svc.ethereumClientMap[protocol.NetworkBinanceSmartChain], err = ethclient.Dial(config.General.BinanceSmartChain.HTTP); err != nil {
		return nil, err
	}

	if svc.ethereumClientMap[protocol.NetworkXDAI], err = ethclient.Dial(config.General.XDAI.HTTP); err != nil {
		return nil, err
	}

	return &svc, nil
}
