package liquidity

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const Name = "liquidity"

var _ worker.Worker = (*internal)(nil)

type internal struct {
	tokenClient       *token.Client
	ethereumClientMap map[string]*ethclient.Client
}

func (i *internal) Name() string {
	return Name
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}
}

func (i *internal) Initialize(ctx context.Context) error {
	return nil
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		ethereumClient, exists := i.ethereumClientMap[message.Network]
		if !exists {
			return nil, errors.New("network not supported")
		}

		router, exists := routerMap[strings.ToLower(transaction.AddressTo)]
		if !exists {
			continue
		}

		internalTransaction := transaction
		internalTransaction.Transfers = make([]model.Transfer, 0)
		internalTransaction.Platform = router.Name

		receipt, err := ethereumClient.TransactionReceipt(context.Background(), common.HexToHash(transaction.Hash))
		if err != nil {
			return nil, err
		}

		for _, log := range receipt.Logs {
			switch log.Topics[0] {
			// Uniswap Protocol V2
			case uniswap.EventHashMintV2:
				internalTransfer, err := i.handleUniswapV2Mint(ctx, message, internalTransaction, *log, router)
				if err != nil {
					return nil, err
				}

				internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransfer.Tag, filter.ExchangeLiquidity, internalTransfer.Type)
				internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransaction.Tag, filter.ExchangeLiquidity, internalTransaction.Type)
				internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
			case uniswap.EventHashBurnV2:
				internalTransfer, err := i.handleUniswapV2Burn(ctx, message, internalTransaction, *log, router)
				if err != nil {
					return nil, err
				}

				internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransfer.Tag, filter.ExchangeLiquidity, internalTransfer.Type)
				internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransaction.Tag, filter.ExchangeLiquidity, internalTransaction.Type)
				internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
			// Uniswap Protocol V3
			case uniswap.EventHashMintV3:
				internalTransfer, err := i.handleUniswapV3Mint(ctx, message, internalTransaction, *log, router)
				if err != nil {
					return nil, err
				}

				internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransfer.Tag, filter.ExchangeLiquidity, internalTransfer.Type)
				internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransaction.Tag, filter.ExchangeLiquidity, internalTransaction.Type)
				internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
			case uniswap.EventHashBurnV3:
				internalTransfer, err := i.handleUniswapV3Burn(ctx, message, internalTransaction, *log, router)
				if err != nil {
					return nil, err
				}

				internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransfer.Tag, filter.ExchangeLiquidity, internalTransfer.Type)
				internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransaction.Tag, filter.ExchangeLiquidity, internalTransaction.Type)
				internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
			case uniswap.EventHashCollectV3:
				// TODO Fee = Collect - Remove
				if lo.ContainsBy(receipt.Logs, func(log *types.Log) bool {
					return log.Topics[0] == uniswap.EventHashBurnV3
				}) {
					continue
				}

				internalTransfer, err := i.handleUniswapV3Collect(ctx, message, internalTransaction, *log, router)
				if err != nil {
					return nil, err
				}

				internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransfer.Tag, filter.ExchangeLiquidity, internalTransfer.Type)
				internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransaction.Tag, filter.ExchangeLiquidity, internalTransaction.Type)
				internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
			}
		}

		if len(internalTransaction.Transfers) == 0 {
			continue
		}

		internalTransaction.Platform = router.Name

		internalTransactions = append(internalTransactions, internalTransaction)
	}

	return internalTransactions, nil
}

func (i *internal) handleUniswapV2Mint(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV2(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV2(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseMint(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityAdd, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(event.Sender.String()),
		AddressTo:       strings.ToLower(log.Address.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV2Burn(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV2(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV2(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseBurn(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityRemove, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.To.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV3Mint(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV3(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV3(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseMint(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityAdd, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(event.Sender.String()),
		AddressTo:       strings.ToLower(log.Address.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV3Burn(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV3(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV3(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseBurn(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityRemove, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.Owner.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV3Collect(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV3(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV3(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseCollect(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityCollect, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.Recipient.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) buildTokenPairV2(ctx context.Context, network string, poolContract *uniswap.PoolV2) (*token.ERC20, *token.ERC20, error) {
	tokenLeftAddress, err := poolContract.Token0(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}

	tokenRightAddress, err := poolContract.Token1(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}

	tokenLeft, err := i.tokenClient.ERC20(ctx, network, tokenLeftAddress.String())
	if err != nil {
		return nil, nil, err
	}

	tokenRight, err := i.tokenClient.ERC20(ctx, network, tokenRightAddress.String())
	if err != nil {
		return nil, nil, err
	}

	return tokenLeft, tokenRight, nil
}

func (i *internal) buildTokenPairV3(ctx context.Context, network string, poolContract *uniswap.PoolV3) (*token.ERC20, *token.ERC20, error) {
	tokenLeftAddress, err := poolContract.Token0(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}

	tokenRightAddress, err := poolContract.Token1(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}

	tokenLeft, err := i.tokenClient.ERC20(ctx, network, tokenLeftAddress.String())
	if err != nil {
		return nil, nil, err
	}

	tokenRight, err := i.tokenClient.ERC20(ctx, network, tokenRightAddress.String())
	if err != nil {
		return nil, nil, err
	}

	return tokenLeft, tokenRight, nil
}

func (i *internal) buildLiquidityMetadata(ctx context.Context, router Router, action string, tokenMap map[*token.ERC20]*big.Int) (json.RawMessage, error) {
	liquidityMetadata := metadata.Liquidity{
		Protocol: router.Protocol,
		Action:   action,
		Tokens:   make([]metadata.Token, 0),
	}

	for internalToken, value := range tokenMap {
		internalTokenValue := decimal.NewFromBigInt(value, 0)

		internalTokenDisplay := internalTokenValue.Shift(-int32(internalToken.Decimals))

		liquidityMetadata.Tokens = append(liquidityMetadata.Tokens, metadata.Token{
			Name:            internalToken.Name,
			Symbol:          internalToken.Symbol,
			Decimals:        internalToken.Decimals,
			Image:           internalToken.Logo,
			Standard:        protocol.TokenStandardERC20,
			ContractAddress: internalToken.ContractAddress,
			Value:           &internalTokenValue,
			ValueDisplay:    &internalTokenDisplay,
		})
	}

	return json.Marshal(&liquidityMetadata)
}

func (i *internal) Jobs() []worker.Job {
	return nil
}

func New(databaseClient *gorm.DB, ethereumClientMap map[string]*ethclient.Client) worker.Worker {
	return &internal{
		tokenClient:       token.New(databaseClient, ethereumClientMap),
		ethereumClientMap: ethereumClientMap,
	}
}
