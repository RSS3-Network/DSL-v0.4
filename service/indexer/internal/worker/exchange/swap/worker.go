package swap

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/aave"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/balancer"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/curve"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/dodo"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/masknetwork"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	lop "github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"

	"go.uber.org/zap"
)

var _ worker.Worker = &service{}

type service struct {
	employer    *shedlock.Employer
	tokenClient *token.Client
}

func (s *service) Name() string {
	return filter.ExchangeSwap
}

func (s *service) Networks() []string {
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

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("worker_swap")
	_, trace := tracer.Start(ctx, "worker_swap:Handle")

	defer opentelemetry.Log(trace, transactions, data, err)

	// Currently only the EVM network needs to be handled
	return s.handleEthereum(ctx, message, transactions)
}

func (s *service) handleEthereum(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("worker_swap")
	_, trace := tracer.Start(ctx, "worker_swap:handleEthereum")

	defer opentelemetry.Log(trace, transactions, data, err)

	internalTransactionMap := make(map[string]model.Transaction)
	opt := lop.NewOption().WithConcurrency(ethereum.RPCMaxConcurrency)

	var mu sync.Mutex

	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		address := message.Address

		// Exclude transfers to self
		if transaction.AddressTo == address {
			return
		}

		// Exclude indexed transactions
		for _, transfer := range transaction.Transfers {
			if !(transfer.Metadata == nil || bytes.Equal(transfer.Metadata, metadata.Default)) {
				return
			}
		}

		// Handle swap entry
		mu.Lock()
		router, exist := routerMap[transaction.AddressTo]
		mu.Unlock()
		if !exist {
			return
		}

		if transaction.Transfers, err = s.handleEthereumTransaction(ctx, message, &transaction, router); err != nil {
			loggerx.Global().Warn("failed to handle ethereum transaction", zap.Error(err), zap.String("network", message.Network), zap.String("transaction_hash", transaction.Hash), zap.String("address", message.Address))

			return
		}

		transaction.Tag, transaction.Type = filter.UpdateTagAndType(filter.TagExchange, transaction.Tag, filter.ExchangeSwap, transaction.Type)

		switch transaction.Platform {
		// CoW Swap uses offline signatures to offer gasless orders, aka signed orders.
		// https://docs.cow.fi/front-end/cowswap
		case protocol.PlatformCow:
			transaction.Owner = message.Address
		default:
			transaction.Owner = transaction.AddressFrom
		}

		mu.Lock()
		internalTransactionMap[transaction.Hash] = transaction
		mu.Unlock()
	}, opt)

	// Lay the map flat
	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (s *service) handleEthereumTransaction(ctx context.Context, message *protocol.Message, transaction *model.Transaction, router Router) (internalTransfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_swap")
	_, trace := tracer.Start(ctx, "worker_swap:handleEthereumTransaction")

	defer opentelemetry.Log(trace, transaction, internalTransfers, err)

	ethereumClient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, errors.New("not supported network")
	}

	tokenMap := map[common.Address]*big.Int{}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source data: %w", err)
	}

	for _, log := range sourceData.Receipt.Logs {
		// Filter anonymous log
		if len(log.Topics) == 0 {
			continue
		}

		var internalTokenMap map[common.Address]*big.Int

		switch log.Topics[0] {
		case uniswap.EventHashSwapV2:
			internalTokenMap, err = s.handleUniswapV2(ctx, message, *log, tokenMap, ethereumClient)
		case uniswap.EventHashSwapV3:
			internalTokenMap, err = s.handleUniswapV3(ctx, message, *log, tokenMap, ethereumClient)
		case dodo.EventHashDODOSwap:
			internalTokenMap, err = s.handleDODO(ctx, message, *log, tokenMap, ethereumClient)
		case curve.EventHashAddLiquidity:
			internalTokenMap, err = s.handleCurve3PoolAddLiquidity(ctx, message, *log, sourceData.Receipt.Logs, tokenMap, ethereumClient)
		case curve.EventHashRemoveLiquidityOne:
			internalTokenMap, err = s.handleCurve3PoolRemoveLiquidityOne(ctx, message, *log, sourceData.Receipt.Logs, tokenMap, ethereumClient)
		case curve.EventHashTokenExchange:
			internalTokenMap, err = s.handleCurve3PoolTokenExchange(ctx, message, *log, tokenMap, ethereumClient)
		case curve.EventHashTokenExchange2:
			internalTokenMap, err = s.handleCurve3PoolTokenExchange2(ctx, message, *log, tokenMap, ethereumClient)
		case balancer.EventHashSwap:
			internalTokenMap, err = s.handleBalancerSwap(ctx, message, *log, tokenMap, ethereumClient)
		case aave.EventHashMint:
			internalTokenMap, err = s.handleAAVEMint(ctx, message, *log, tokenMap, ethereumClient)
		case aave.EventHashBurn:
			internalTokenMap, err = s.handleAAVEBurn(ctx, message, *log, tokenMap, ethereumClient)
		case masknetwork.EventSwapSuccess:
			internalTokenMap, err = s.handleMaskEventSwapSuccess(ctx, message, *log, tokenMap, ethereumClient)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		tokenMap = internalTokenMap
	}

	if len(tokenMap) == 0 {
		return nil, errors.New("unsupported protocol")
	}

	var addressTo string

	switch {
	case router == cowSwap: // Agent
		addressTo = message.Address
	default:
		addressTo = transaction.AddressFrom
	}

	transfer := model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           0, // TODO
		AddressFrom:     transaction.AddressFrom,
		AddressTo:       addressTo,
		Metadata:        metadata.Default,
		Network:         message.Network,
		Source:          protocol.SourceOrigin,
		RelatedUrls: []string{
			ethereum.BuildScanURL(message.Network, transaction.Hash),
		},
	}

	if len(tokenMap) > 0 {
		swapMetadataModel := metadata.Swap{
			Protocol: router.Protocol,
		}

		for token, value := range tokenMap {
			erc20, err := s.tokenClient.ERC20(ctx, message.Network, token.String())
			if err != nil {
				return nil, err
			}

			tokenValueTo := decimal.NewFromBigInt(value, 0)
			tokenValueDisplayTo := tokenValueTo.Shift(-int32(erc20.Decimals))

			tokenValueFrom := tokenValueTo.Abs()
			tokenValueDisplayFrom := tokenValueFrom.Shift(-int32(erc20.Decimals))

			switch value.Cmp(big.NewInt(0)) {
			case -1:
				swapMetadataModel.TokenFrom = metadata.Token{
					Name:            erc20.Name,
					Symbol:          erc20.Symbol,
					Decimals:        erc20.Decimals,
					Value:           &tokenValueFrom,
					ValueDisplay:    &tokenValueDisplayFrom,
					ContractAddress: token.String(),
					Standard:        protocol.TokenStandardERC20,
					Image:           erc20.Logo,
				}
			case 0:
				continue
			case 1:
				swapMetadataModel.TokenTo = metadata.Token{
					Name:            erc20.Name,
					Symbol:          erc20.Symbol,
					Decimals:        erc20.Decimals,
					Value:           &tokenValueTo,
					ValueDisplay:    &tokenValueDisplayTo,
					ContractAddress: token.String(),
					Standard:        protocol.TokenStandardERC20,
					Image:           erc20.Logo,
				}
			}
		}

		swapMetadata, err := json.Marshal(swapMetadataModel)
		if err != nil {
			return nil, err
		}

		transfer.Metadata = swapMetadata
		transfer.Tag = filter.TagExchange
		transfer.Type = filter.ExchangeSwap
		transaction.Platform = router.Name
		internalTransfers = append(internalTransfers, transfer)
	} else {
		// let transaction worker take care of it
		if transaction.SourceData, err = json.Marshal(sourceData.Receipt.Logs); err != nil {
			return nil, err
		}
	}

	return internalTransfers, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&Job{},
	}
}

func New(employer *shedlock.Employer) (worker.Worker, error) {
	svc := service{
		employer:    employer,
		tokenClient: token.New(),
	}
	return &svc, nil
}
