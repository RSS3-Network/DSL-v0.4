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
			// TODO Uniswap Protocol V2
			// Uniswap Protocol V3
			case uniswap.EventHashMintV3:
				internalTransfer, err := i.handleUniswapV3Mint(ctx, message, internalTransaction, *log, router)
				if err != nil {
					return nil, err
				}

				internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransfer.Tag, filter.ExchangeAddLiquidity, internalTransfer.Type)
				internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransaction.Tag, filter.ExchangeAddLiquidity, internalTransaction.Type)
				internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
			case uniswap.EventHashBurnV3:
				internalTransfer, err := i.handleUniswapV3Burn(ctx, message, internalTransaction, *log, router)
				if err != nil {
					return nil, err
				}

				internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransfer.Tag, filter.ExchangeRemoveLiquidity, internalTransfer.Type)
				internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransaction.Tag, filter.ExchangeRemoveLiquidity, internalTransaction.Type)
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

				internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransfer.Tag, filter.ExchangeCollect, internalTransfer.Type)
				internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransaction.Tag, filter.ExchangeCollect, internalTransaction.Type)
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

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, tokenLeft, event.Amount0, tokenRight, event.Amount1)
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

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, tokenLeft, event.Amount0, tokenRight, event.Amount1)
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

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, tokenLeft, event.Amount0, tokenRight, event.Amount1)
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

func (i *internal) buildLiquidityMetadata(ctx context.Context, router Router, tokenLeft *token.ERC20, tokenLeftValue *big.Int, tokenRight *token.ERC20, tokenRightValue *big.Int) (json.RawMessage, error) {
	internalTokenLeftValue := decimal.NewFromBigInt(tokenLeftValue, 0)
	internalTokenRightValue := decimal.NewFromBigInt(tokenRightValue, 0)

	return json.Marshal(&metadata.Liquidity{
		Protocol: router.Protocol,
		TokenLeft: metadata.Token{
			Name:            tokenLeft.Name,
			Symbol:          tokenLeft.Symbol,
			Decimals:        tokenLeft.Decimals,
			Image:           tokenLeft.Logo,
			Standard:        protocol.TokenStandardERC20,
			ContractAddress: tokenLeft.ContractAddress,
			Value:           &internalTokenLeftValue,
		},
		TokenRight: metadata.Token{
			Name:            tokenLeft.Name,
			Symbol:          tokenRight.Symbol,
			Decimals:        tokenRight.Decimals,
			Image:           tokenRight.Logo,
			Standard:        protocol.TokenStandardERC20,
			ContractAddress: tokenRight.ContractAddress,
			Value:           &internalTokenRightValue,
		},
	})
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
