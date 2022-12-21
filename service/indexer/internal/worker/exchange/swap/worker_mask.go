package swap

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/masknetwork"
	"github.com/naturalselectionlabs/pregod/common/protocol"

	"go.uber.org/zap"
)

func (s *service) handleMaskEventSwapSuccess(ctx context.Context, message *protocol.Message, log types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	filterer, err := masknetwork.NewHappyTokenPoolFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new filterer: %w", err)
	}

	event, err := filterer.ParseSwapSuccess(log)
	if err != nil {
		return nil, fmt.Errorf("parse swap success: %w", err)
	}

	tokenFromValue, exist := tokenMap[event.FromAddress]
	if !exist {
		tokenFromValue = big.NewInt(0)
	}

	tokenToValue, exist := tokenMap[event.ToAddress]
	if !exist {
		tokenToValue = big.NewInt(0)
	}

	tokenMap[event.FromAddress] = tokenFromValue.Sub(tokenFromValue, event.FromValue)
	tokenMap[event.ToAddress] = tokenToValue.Add(tokenToValue, event.ToValue)

	zap.L().Debug(
		"swap by swap in mask network",
		zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
		zap.Stringer("token_from", event.FromAddress), zap.Stringer("token_from_value", event.FromValue),
		zap.Stringer("token_to", event.ToAddress), zap.Stringer("token_to_value", event.ToValue),
	)

	return tokenMap, nil
}
