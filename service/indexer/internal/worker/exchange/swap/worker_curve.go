package swap

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/curve"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

func (s *service) handleCurve3PoolAddLiquidity(ctx context.Context, message *protocol.Message, log types.Log, relatedLogs []*types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	threePool, err := curve.NewThreePool(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new 3pool: %w", err)
	}

	eventAddLiquidity, err := threePool.ParseAddLiquidity(log)
	if err != nil {
		return nil, fmt.Errorf("parse add liquidity: %w", err)
	}

	var (
		tokenFrom, tokenTo           common.Address
		tokenFromValue, tokenToValue *big.Int
		exists                       bool
	)

	// Get the added coin id
	for index, amount := range eventAddLiquidity.TokenAmounts {
		if amount.Cmp(decimal.Zero.BigInt()) == 1 { // Need amount > 0
			// Get the token contract address by coin id
			tokenFrom, err = threePool.Coins(&bind.CallOpts{}, big.NewInt(int64(index)))
			if err != nil {
				return nil, fmt.Errorf("coins: %w", err)
			}

			// Calculate the value of token from
			tokenFromValue, exists = tokenMap[tokenFrom]
			if !exists {
				tokenFromValue = decimal.Zero.BigInt()
			}

			tokenMap[tokenFrom] = tokenFromValue.Sub(tokenFromValue, amount)

			// Find the related transfer from other logs
			for _, relatedLog := range relatedLogs {
				if relatedLog.Index > log.Index && len(relatedLog.Topics) == 3 /* With topic hash + sender + receiver */ && relatedLog.Topics[0] == erc20.EventHashTransfer {
					// Parse transfer event
					erc20Filterer, err := erc20.NewERC20Filterer(relatedLog.Address, ethereumClient)
					if err != nil {
						return nil, fmt.Errorf("new erc20 filterer: %w", err)
					}

					eventTransfer, err := erc20Filterer.ParseTransfer(*relatedLog)
					if err != nil {
						return nil, fmt.Errorf("parse transfer: %w", err)
					}

					// Matched transfer event
					if eventTransfer.From == eventAddLiquidity.Provider {
						tokenTo = relatedLog.Address

						// Calculate the value of token to
						tokenToValue, exists = tokenMap[tokenTo]
						if !exists {
							tokenToValue = decimal.Zero.BigInt()
						}

						tokenMap[tokenTo] = tokenToValue.Add(tokenToValue, eventTransfer.Value)

						zap.L().Debug(
							"swap by add liquidity in curve pool",
							zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
							zap.Stringer("token_from", tokenFrom), zap.Stringer("token_from_value", amount),
							zap.Stringer("token_to", tokenTo), zap.Stringer("token_to_value", eventTransfer.Value),
						)

						break
					}
				}
			}

			break
		}
	}

	return tokenMap, nil
}

func (s *service) handleCurve3PoolRemoveLiquidityOne(ctx context.Context, message *protocol.Message, log types.Log, relatedLogs []*types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	threePool, err := curve.NewThreePool(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new 3pool: %w", err)
	}

	eventRemoveLiquidityOne, err := threePool.ParseRemoveLiquidityOne(log)
	if err != nil {
		return nil, fmt.Errorf("parse remove liquidity one: %w", err)
	}

	registryExchange, err := curve.NewRegistryExchange(eventRemoveLiquidityOne.Provider, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new registry exchange: %w", err)
	}

	registryAddress, err := registryExchange.Registry(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("get registry address: %w", err)
	}

	registry, err := curve.NewRegistry(registryAddress, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new registry: %w", err)
	}

	var (
		tokenFrom, tokenTo           common.Address
		tokenFromValue, tokenToValue *big.Int
		exists                       bool
	)

	if tokenFrom, err = registry.GetLpToken(&bind.CallOpts{}, log.Address); err != nil {
		return nil, fmt.Errorf("get lp token: %w", err)
	}

	// Calculate the value of token to
	tokenFromValue, exists = tokenMap[tokenFrom]
	if !exists {
		tokenFromValue = decimal.Zero.BigInt()
	}

	tokenMap[tokenFrom] = tokenFromValue.Sub(tokenFromValue, eventRemoveLiquidityOne.TokenAmount)

	// Find the related transfer from other logs
	for _, relatedLog := range relatedLogs {
		if relatedLog.Index > log.Index && len(relatedLog.Topics) == 3 /* With topic hash + sender + receiver */ && relatedLog.Topics[0] == erc20.EventHashTransfer {
			// Parse transfer event
			erc20Filterer, err := erc20.NewERC20Filterer(relatedLog.Address, ethereumClient)
			if err != nil {
				return nil, fmt.Errorf("new erc20 filterer: %w", err)
			}

			eventTransfer, err := erc20Filterer.ParseTransfer(*relatedLog)
			if err != nil {
				return nil, fmt.Errorf("parse transfer: %w", err)
			}

			// Matched transfer event
			if eventTransfer.From == eventRemoveLiquidityOne.Provider {
				tokenTo = relatedLog.Address

				// Calculate the value of token to
				tokenToValue, exists = tokenMap[tokenTo]
				if !exists {
					tokenToValue = decimal.Zero.BigInt()
				}

				tokenMap[tokenTo] = tokenToValue.Add(tokenToValue, eventTransfer.Value)

				zap.L().Debug(
					"swap by remove liquidity one in curve pool",
					zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
					zap.Stringer("token_from", tokenFrom), zap.Stringer("token_from_value", eventRemoveLiquidityOne.TokenAmount),
					zap.Stringer("token_to", tokenTo), zap.Stringer("token_to_value", eventTransfer.Value),
				)

				break
			}
		}
	}

	return tokenMap, nil
}

func (s *service) handleCurve3PoolTokenExchange(ctx context.Context, message *protocol.Message, log types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	threePool, err := curve.NewThreePool(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new 3pool: %w", err)
	}

	eventTokenExchange, err := threePool.ParseTokenExchange(log)
	if err != nil {
		return nil, fmt.Errorf("parse token excahnge: %w", err)
	}

	tokenFrom, err := threePool.Coins(&bind.CallOpts{}, eventTokenExchange.SoldId)
	if err != nil {
		return nil, fmt.Errorf("get token from: %w", err)
	}

	// Calculate the value of token from
	tokenFromValue, exists := tokenMap[tokenFrom]
	if !exists {
		tokenFromValue = decimal.Zero.BigInt()
	}

	tokenTo, err := threePool.Coins(&bind.CallOpts{}, eventTokenExchange.BoughtId)
	if err != nil {
		return nil, fmt.Errorf("get token to: %w", err)
	}

	// Calculate the value of token to
	tokenToValue, exists := tokenMap[tokenTo]
	if !exists {
		tokenToValue = decimal.Zero.BigInt()
	}

	tokenMap[tokenFrom] = tokenFromValue.Sub(tokenFromValue, eventTokenExchange.TokensSold)
	tokenMap[tokenTo] = tokenToValue.Add(tokenToValue, eventTokenExchange.TokensBought)

	zap.L().Debug(
		"swap by token exchange in curve pool",
		zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
		zap.Stringer("token_from", tokenFrom), zap.Stringer("token_from_value", eventTokenExchange.TokensSold),
		zap.Stringer("token_to", tokenTo), zap.Stringer("token_to_value", eventTokenExchange.TokensBought),
	)

	return tokenMap, nil
}

func (s *service) handleCurve3PoolTokenExchange2(ctx context.Context, message *protocol.Message, log types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	threePool, err := curve.NewThreePool2(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new 3pool: %w", err)
	}

	eventTokenExchange, err := threePool.ParseTokenExchange(log)
	if err != nil {
		return nil, fmt.Errorf("parse token excahnge: %w", err)
	}

	tokenFrom, err := threePool.Coins(&bind.CallOpts{}, eventTokenExchange.SoldId)
	if err != nil {
		return nil, fmt.Errorf("get token from: %w", err)
	}

	// Calculate the value of token from
	tokenFromValue, exists := tokenMap[tokenFrom]
	if !exists {
		tokenFromValue = decimal.Zero.BigInt()
	}

	tokenTo, err := threePool.Coins(&bind.CallOpts{}, eventTokenExchange.BoughtId)
	if err != nil {
		return nil, fmt.Errorf("get token to: %w", err)
	}

	// Calculate the value of token to
	tokenToValue, exists := tokenMap[tokenTo]
	if !exists {
		tokenToValue = decimal.Zero.BigInt()
	}

	tokenMap[tokenFrom] = tokenFromValue.Sub(tokenFromValue, eventTokenExchange.TokensSold)
	tokenMap[tokenTo] = tokenToValue.Add(tokenToValue, eventTokenExchange.TokensBought)

	zap.L().Debug(
		"swap by token exchange in curve pool",
		zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
		zap.Stringer("token_from", tokenFrom), zap.Stringer("token_from_value", eventTokenExchange.TokensSold),
		zap.Stringer("token_to", tokenTo), zap.Stringer("token_to_value", eventTokenExchange.TokensBought),
	)

	return tokenMap, nil
}
