package swap

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/protocol"

	"go.uber.org/zap"
)

func (s *service) handleUniswapV2(ctx context.Context, message *protocol.Message, log types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	uniswapPoolV2Contract, err := uniswap.NewPoolV2(log.Address, ethereumClient)
	if err != nil {
		return nil, err
	}

	event, err := uniswapPoolV2Contract.ParseSwap(log)
	if err != nil {
		return nil, err
	}

	token0, err := uniswapPoolV2Contract.Token0(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	token1, err := uniswapPoolV2Contract.Token1(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	token0Value, exist := tokenMap[token0]
	if !exist {
		token0Value = big.NewInt(0)
	}

	token1Value, exist := tokenMap[token1]
	if !exist {
		token1Value = big.NewInt(0)
	}

	if event.Amount1Out.Cmp(big.NewInt(0)) == 1 {
		// Swap token0 to token1
		tokenMap[token0] = token0Value.Sub(token0Value, event.Amount0In)
		tokenMap[token1] = token1Value.Add(token1Value, event.Amount1Out)

		zap.L().Debug(
			"swap by swap in uniswap v2",
			zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
			zap.Stringer("token_from", token0), zap.Stringer("token_from_value", event.Amount0In),
			zap.Stringer("token_to", token1), zap.Stringer("token_to_value", event.Amount1Out),
		)
	} else {
		// Swap token1 to token0
		tokenMap[token0] = token0Value.Add(token0Value, event.Amount0Out)
		tokenMap[token1] = token1Value.Sub(token1Value, event.Amount1In)

		zap.L().Debug(
			"swap by swap in uniswap v2",
			zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
			zap.Stringer("token_from", token1), zap.Stringer("token_from_value", event.Amount1In),
			zap.Stringer("token_to", token0), zap.Stringer("token_to_value", event.Amount0Out),
		)
	}

	return tokenMap, nil
}

func (s *service) handleUniswapV3(ctx context.Context, message *protocol.Message, log types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	uniswapPoolContact, err := uniswap.NewPoolV3(common.HexToAddress(log.Address.Hex()), ethereumClient)
	if err != nil {
		return nil, err
	}

	event, err := uniswapPoolContact.ParseSwap(log)
	if err != nil {
		return nil, err
	}

	token0, err := uniswapPoolContact.Token0(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	token1, err := uniswapPoolContact.Token1(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	token0Value, exist := tokenMap[token0]
	if !exist {
		token0Value = big.NewInt(0)
	}

	token1Value, exist := tokenMap[token1]
	if !exist {
		token1Value = big.NewInt(0)
	}

	if event.Amount0.Cmp(big.NewInt(0)) == 1 {
		zap.L().Debug(
			"swap by swap in uniswap v3",
			zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
			zap.Stringer("token_from", token0), zap.Stringer("token_from_value", event.Amount0),
			zap.Stringer("token_to", token1), zap.Stringer("token_to_value", event.Amount1),
		)
	} else {
		zap.L().Debug(
			"swap by swap in uniswap v3",
			zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
			zap.Stringer("token_from", token1), zap.Stringer("token_from_value", event.Amount1),
			zap.Stringer("token_to", token0), zap.Stringer("token_to_value", event.Amount0),
		)
	}

	tokenMap[token0] = token0Value.Sub(token0Value, event.Amount0)
	tokenMap[token1] = token1Value.Sub(token1Value, event.Amount1)

	return tokenMap, nil
}
