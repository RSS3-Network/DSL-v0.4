package swap

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	Name = "swap"
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
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("worker_swap")
	_, trace := tracer.Start(ctx, "worker_swap:Handle")

	defer opentelemetry.Log(trace, transactions, data, err)

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

	defer opentelemetry.Log(trace, transactions, data, err)

	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		address := message.Address

		// Exclude transfers to self
		if transaction.AddressTo == address {
			continue
		}

		// Handle swap entry
		router, exist := routerMap[transaction.AddressTo]
		if !exist {
			continue
		}

		if transaction.Transfers, err = s.handleEthereumTransaction(ctx, message, transaction, router); err != nil {
			logger.Global().Warn("failed to handle ethereum transaction", zap.Error(err), zap.String("network", message.Network), zap.String("transaction_hash", transaction.Hash), zap.String("address", message.Address))

			continue
		}

		transaction.Tag, transaction.Tag = filter.UpdateTagAndType(filter.TagExchange, transaction.Tag, filter.ExchangeSwap, transaction.Type)

		internalTransactionMap[transaction.Hash] = transaction
	}

	// Lay the map flat
	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (s *service) handleEthereumTransaction(ctx context.Context, message *protocol.Message, transaction model.Transaction, router Router) ([]model.Transfer, error) {
	ethereumClient, exist := s.ethereumClientMap[message.Network]
	if !exist {
		return nil, errors.New("not supported network")
	}

	internalTransfers := make([]model.Transfer, 0)
	tokenMap := map[common.Address]*big.Int{}

	for _, transfer := range transaction.Transfers {
		if transfer.Index == protocol.IndexVirtual {
			internalTransfers = append(internalTransfers, transfer)

			break
		}
	}

	receipt, err := ethereumClient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
	if err != nil {
		return nil, err
	}

	internalTokenMap := make(map[common.Address]*big.Int)

	for _, log := range receipt.Logs {
		internalTokenMap = make(map[common.Address]*big.Int)

		for _, topic := range log.Topics {
			switch topic {
			case uniswap.EventHashSwapV2:
				internalTokenMap, err = s.handleUniswapV2(ctx, message, *log, tokenMap, ethereumClient)
			case uniswap.EventHashSwapV3:
				internalTokenMap, err = s.handleUniswapV3(ctx, message, *log, tokenMap, ethereumClient)
			default:
				continue
			}

			if err != nil {
				return nil, err
			}

			tokenMap = internalTokenMap
		}
	}

	swapMetadataModel := metadata.Swap{
		Protocol: router.Protocol,
	}

	for token, value := range tokenMap {
		erc20Contract, err := erc20.NewERC20(token, ethereumClient)
		if err != nil {
			return nil, err
		}

		tokenSymbol, err := erc20Contract.Symbol(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}

		tokenDecimals, err := erc20Contract.Decimals(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}

		tokenValue := decimal.NewFromBigInt(value, 0).DivRound(decimal.NewFromBigInt(big.NewInt(1), int32(tokenDecimals)), int32(tokenDecimals))

		switch value.Cmp(big.NewInt(0)) {
		case -1:
			swapMetadataModel.TokenFrom = metadata.SwapToken{
				Address:  token.String(),
				Symbol:   tokenSymbol,
				Decimals: tokenDecimals,
				Value:    tokenValue.Abs(),
			}
		case 0:
			continue
		case 1:
			swapMetadataModel.TokenTo = metadata.SwapToken{
				Address:  token.String(),
				Symbol:   tokenSymbol,
				Decimals: tokenDecimals,
				Value:    tokenValue,
			}
		}
	}

	swapMetadata, err := json.Marshal(swapMetadataModel)
	if err != nil {
		return nil, err
	}

	internalTransfers = append(internalTransfers, model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Tag:             filter.TagExchange,
		Type:            filter.ExchangeSwap,
		Index:           0, // TODO
		AddressFrom:     transaction.AddressFrom,
		AddressTo:       transaction.AddressFrom,
		Metadata:        swapMetadata,
		Network:         message.Network,
		Platform:        router.Name,
		Source:          protocol.SourceOrigin,
		// SourceData:      sourceDa,
		RelatedUrls: []string{
			ethereum.BuildScanURL(message.Network, transaction.Hash),
		},
	})

	return internalTransfers, nil
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
