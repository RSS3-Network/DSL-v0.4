package swap

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/errorx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/sirupsen/logrus"
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
	routerUniswap = Router{
		Name:     "Uniswap",
		Protocol: "UniSwapV3",
	}

	routerSushiSwap = Router{
		Name:     "SushiSwap",
		Protocol: "UniSwapV3",
	}

	pancakeSwap = Router{
		Name:     "PancakeSwap",
		Protocol: "UniSwapV3",
	}

	routerMap = map[string]Router{
		// Uniswap
		// https://docs.uniswap.org/protocol/reference/deployments
		strings.ToLower("0xE592427A0AEce92De3Edee1F18E0157C05861564"): routerUniswap, // Uniswap V3 1
		strings.ToLower("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"): routerUniswap, // Uniswap V3 2
		// SushiSwap
		// https://docs.sushi.com/docs/Developers/Deployment%20Addresses
		strings.ToLower("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"): routerSushiSwap, // SushiSwap Ethereum
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushiSwap, // SushiSwap Polygon
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushiSwap, // SushiSwap Binance Smart Chain
		// PancakeSwap
		// https://docs.pancakeswap.finance/code/smart-contracts/pancakeswap-exchange/router-v2
		strings.ToLower("0x10ED43C718714eb63d5aA57B78B54704E256024E"): pancakeSwap, // PancakeSwap V2
	}

	ErrorMetadataDoesNotHaveTokenField   = errors.New("metadata doesn't have token field")
	ErrorRouterTransferIsNotYetSupported = errors.New("router transfer isn't yet supported")
)

var _ worker.Worker = &service{}

type service struct {
	employer       *shedlock.Employer
	databaseClient *gorm.DB
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

	return job.Run(func(ctx context.Context, duration time.Duration) error {
		return s.employer.Renewal(ctx, job.Name(), duration)
	})
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	switch message.Network {
	case protocol.NetworkZkSync:
		return s.handleZkSync(ctx, message, transactions)
	default:
		return s.handleEthereum(ctx, message, transactions)
	}
}

func (s *service) handleEthereum(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		address := strings.ToLower(message.Address)

		// Exclude transfers to self
		if transaction.AddressTo == address {
			continue
		}

		// Handle swap entry
		router, exist := routerMap[transaction.AddressTo]
		if !exist {
			continue
		}

		for _, transfer := range transaction.Transfers {
			internalTransfer, err := s.handleEthereumTransfer(ctx, message, transfer, transaction.AddressTo)
			if err != nil {
				if errorx.IsUnexpectedError(err, gorm.ErrRecordNotFound, ErrorRouterTransferIsNotYetSupported, ErrorRouterTransferIsNotYetSupported) {
					logrus.Error(err)
				}

				continue
			}

			internalTransfer.Tag = filter.UpdateTag(filter.TagExchange, internalTransfer.Tag)

			if internalTransfer.Tag == filter.TagExchange {
				internalTransfer.Type = filter.ExchangeSwap
				internalTransfer.Platform = router.Name
			}

			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Transfers = append(value.Transfers, *internalTransfer)

			// Update transaction tag
			value.Tag = filter.UpdateTag(internalTransfer.Tag, transaction.Tag)

			if value.Tag == filter.TagExchange {
				value.Platform = router.Name
			}

			internalTransactionMap[transaction.Hash] = value
		}
	}

	// Lay the map flat
	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (s *service) handleEthereumTransfer(ctx context.Context, message *protocol.Message, transfer model.Transfer, swapRouterAddress string) (*model.Transfer, error) {
	var metadataModel metadata.Metadata

	if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
		return nil, err
	}

	if metadataModel.Token == nil {
		return nil, ErrorMetadataDoesNotHaveTokenField
	}

	var swapPoolAddress string

	switch strings.ToLower(message.Address) {
	case transfer.AddressFrom:
		swapPoolAddress = transfer.AddressTo
	case transfer.AddressTo:
		swapPoolAddress = transfer.AddressFrom
	default:
		// TODO Router

		return nil, ErrorRouterTransferIsNotYetSupported
	}

	if transfer.AddressFrom == swapRouterAddress || transfer.AddressTo == swapRouterAddress {
		if err := s.handleEthereumTransferSwapRouter(ctx, message, &transfer, swapRouterAddress); err != nil {
			return nil, err
		}
	} else {
		if err := s.handleEthereumTransferSwapPool(ctx, message, &transfer, swapPoolAddress); err != nil {
			return nil, err
		}
	}

	return &transfer, nil
}

func (s *service) handleEthereumTransferSwapPool(ctx context.Context, message *protocol.Message, transfer *model.Transfer, swapPoolAddress string) error {
	var swapPoolModel model.SwapPool

	if err := s.databaseClient.
		Model((*model.SwapPool)(nil)).
		Where(map[string]interface{}{
			"contract_address": swapPoolAddress,
			"network":          message.Network,
		}).
		First(&swapPoolModel).
		Error; err != nil {
		return err
	}

	var err error

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.SwapPool{
		Name:     swapPoolModel.Source,
		Type:     metadata.SwapTypePool,
		Network:  swapPoolModel.Network,
		Token0:   swapPoolModel.Token0,
		Token1:   swapPoolModel.Token1,
		Protocol: swapPoolModel.Protocol,
	}); err != nil {
		return err
	}

	return nil
}

func (s *service) handleEthereumTransferSwapRouter(ctx context.Context, message *protocol.Message, transfer *model.Transfer, swapRouterAddress string) error {
	var err error

	router := routerMap[swapRouterAddress]

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.SwapPool{
		Name:     router.Name,
		Type:     metadata.SwapTypeRouter,
		Network:  message.Network,
		Protocol: router.Protocol,
	}); err != nil {
		return err
	}

	return nil
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

func New(employer *shedlock.Employer, databaseClient *gorm.DB) worker.Worker {
	return &service{
		employer:       employer,
		databaseClient: databaseClient,
	}
}
