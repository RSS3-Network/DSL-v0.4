package swap

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/dodo"
	"github.com/naturalselectionlabs/pregod/common/protocol"

	"go.uber.org/zap"
)

func (s *service) handleDODO(ctx context.Context, message *protocol.Message, log types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	vendingMachine, err := dodo.NewVendingMachine(common.HexToAddress(log.Address.Hex()), ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("create kyber aggregation router: %w", err)
	}

	event, err := vendingMachine.ParseDODOSwap(log)
	if err != nil {
		return nil, fmt.Errorf("parse swapped event: %w", err)
	}

	token0Value, exist := tokenMap[event.FromToken]
	if !exist {
		token0Value = big.NewInt(0)
	}

	token1Value, exist := tokenMap[event.ToToken]
	if !exist {
		token1Value = big.NewInt(0)
	}

	tokenMap[event.FromToken] = token0Value.Sub(token0Value, event.FromAmount)
	tokenMap[event.ToToken] = token1Value.Add(token1Value, event.ToAmount)

	zap.L().Debug(
		"swap by swap in dodo",
		zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
		zap.Stringer("token_from", event.FromToken), zap.Stringer("token_from_value", event.FromAmount),
		zap.Stringer("token_to", event.ToToken), zap.Stringer("token_to_value", event.ToAmount),
	)

	return tokenMap, nil
}
