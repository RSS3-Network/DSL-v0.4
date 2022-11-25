package swap

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/balancer"
	"github.com/naturalselectionlabs/pregod/common/protocol"

	"go.uber.org/zap"
)

func (s *service) handleBalancerSwap(ctx context.Context, message *protocol.Message, log types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	vault, err := balancer.NewVault(common.HexToAddress(log.Address.Hex()), ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("create kyber aggregation router: %w", err)
	}

	event, err := vault.ParseSwap(log)
	if err != nil {
		return nil, fmt.Errorf("parse swap: %w", err)
	}

	tokenFromValue, exists := tokenMap[event.TokenIn]
	if !exists {
		tokenFromValue = big.NewInt(0)
	}

	tokenToValue, exists := tokenMap[event.TokenOut]
	if !exists {
		tokenToValue = big.NewInt(0)
	}

	tokenMap[event.TokenIn] = tokenFromValue.Sub(tokenFromValue, event.AmountIn)
	tokenMap[event.TokenOut] = tokenToValue.Add(tokenToValue, event.AmountOut)

	zap.L().Debug(
		"swap by swap in balancer",
		zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
		zap.Stringer("token_from", event.TokenIn), zap.Stringer("token_from_value", event.AmountIn),
		zap.Stringer("token_to", event.TokenOut), zap.Stringer("token_to_value", event.AmountOut),
	)

	return tokenMap, nil
}
