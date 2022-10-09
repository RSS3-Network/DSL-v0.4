package liquidity

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/aave"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

const Name = "liquidity"

var _ worker.Worker = (*internal)(nil)

type internal struct {
	tokenClient *token.Client
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
		ethclient, err := ethclientx.Global(message.Network)
		if err != nil {
			return nil, err
		}

		router, exists := routerMap[strings.ToLower(transaction.AddressTo)]
		if !exists {
			continue
		}

		internalTransaction := transaction
		internalTransaction.Transfers = make([]model.Transfer, 0)
		internalTransaction.Platform = router.Name

		receipt, err := ethclient.TransactionReceipt(context.Background(), common.HexToHash(transaction.Hash))
		if err != nil {
			return nil, err
		}

		for _, log := range receipt.Logs {
			var internalTransfer *model.Transfer

			switch log.Topics[0] {
			// Uniswap Protocol V2
			case uniswap.EventHashMintV2:
				internalTransfer, err = i.handleUniswapV2Mint(ctx, message, internalTransaction, *log, router)
			case uniswap.EventHashBurnV2:
				internalTransfer, err = i.handleUniswapV2Burn(ctx, message, internalTransaction, *log, router)
			// Uniswap Protocol V3
			case uniswap.EventHashMintV3:
				internalTransfer, err = i.handleUniswapV3Mint(ctx, message, internalTransaction, *log, router)
			case uniswap.EventHashBurnV3:
				internalTransfer, err = i.handleUniswapV3Burn(ctx, message, internalTransaction, *log, router)
			case uniswap.EventHashCollectV3:
				internalTransfer, err = i.handleUniswapV3Collect(ctx, message, internalTransaction, *log, router)
			case aave.EventHashSupplyV2:
				internalTransfer, err = i.handleAAVEV2Deposit(ctx, message, transaction, *log, router)
			case aave.EventHashBorrowV2:
				internalTransfer, err = i.handleAAVEV2Borrow(ctx, message, transaction, *log, router)
			case aave.EventHashRepayV2:
				internalTransfer, err = i.handleAAVEV2Repay(ctx, message, transaction, *log, router)
			case aave.EventHashWithdrawV2:
				internalTransfer, err = i.handleAAVEV2Withdraw(ctx, message, transaction, *log, router)
			case aave.EventHashSupplyV3:
				internalTransfer, err = i.handleAAVEV3Supply(ctx, message, transaction, *log, router)
			case aave.EventHashBorrowV3:
				internalTransfer, err = i.handleAAVEV3Borrow(ctx, message, transaction, *log, router)
			case aave.EventHashRepayV3:
				internalTransfer, err = i.handleAAVEV3Repay(ctx, message, transaction, *log, router)
			case aave.EventHashWithdrawV3:
				internalTransfer, err = i.handleAAVEV3Withdraw(ctx, message, transaction, *log, router)
			default:
				continue
			}

			if err != nil {
				if !errors.Is(err, ErrorInvalidNumberOfToken) {
					zap.L().Error("failed to handle log", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.Uint("log_index", log.Index))
				}

				continue
			}

			internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransfer.Tag, filter.ExchangeLiquidity, internalTransfer.Type)
			internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(filter.TagExchange, internalTransaction.Tag, filter.ExchangeLiquidity, internalTransaction.Type)
			internalTransaction.Transfers = append(internalTransaction.Transfers, *internalTransfer)
		}

		if len(internalTransaction.Transfers) == 0 {
			continue
		}

		internalTransaction.Platform = router.Name

		internalTransactions = append(internalTransactions, internalTransaction)
	}

	return internalTransactions, nil
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

func New() worker.Worker {
	return &internal{
		tokenClient: token.New(),
	}
}
