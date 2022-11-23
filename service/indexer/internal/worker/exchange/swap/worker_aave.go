package swap

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/aave"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

func (s *service) handleAAVEMint(ctx context.Context, message *protocol.Message, log types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	pool, err := aave.NewToken(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new pool: %w", err)
	}

	event, err := pool.ParseMint(log)
	if err != nil {
		return nil, fmt.Errorf("parse mint: %w", err)
	}

	tokenFrom, err := pool.UNDERLYINGASSETADDRESS(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("get underlying asset address: %w", err)
	}

	// Calculate the value of token from
	tokenFromValue, exists := tokenMap[tokenFrom]
	if !exists {
		tokenFromValue = decimal.Zero.BigInt()
	}

	tokenTo := log.Address

	// Calculate the value of token to
	tokenToValue, exists := tokenMap[tokenTo]
	if !exists {
		tokenToValue = decimal.Zero.BigInt()
	}

	tokenMap[tokenFrom] = tokenFromValue.Sub(tokenFromValue, event.Value)
	tokenMap[tokenTo] = tokenToValue.Add(tokenToValue, event.Value)

	zap.L().Debug(
		"swap by mint in aave",
		zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
		zap.Stringer("token_from", tokenFrom), zap.Stringer("token_from_value", event.Value),
		zap.Stringer("token_to", tokenTo), zap.Stringer("token_to_value", event.Value),
	)

	return tokenMap, nil
}

func (s *service) handleAAVEBurn(ctx context.Context, message *protocol.Message, log types.Log, tokenMap map[common.Address]*big.Int, ethereumClient *ethclient.Client) (map[common.Address]*big.Int, error) {
	pool, err := aave.NewToken(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new pool: %w", err)
	}

	event, err := pool.ParseBurn(log)
	if err != nil {
		return nil, fmt.Errorf("parse burn: %w", err)
	}

	tokenFrom, err := pool.UNDERLYINGASSETADDRESS(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("get underlying asset address: %w", err)
	}

	// Calculate the value of token from
	tokenFromValue, exists := tokenMap[tokenFrom]
	if !exists {
		tokenFromValue = decimal.Zero.BigInt()
	}

	tokenTo := log.Address

	// Calculate the value of token to
	tokenToValue, exists := tokenMap[tokenTo]
	if !exists {
		tokenToValue = decimal.Zero.BigInt()
	}

	tokenMap[tokenFrom] = tokenFromValue.Add(tokenFromValue, event.Value)
	tokenMap[tokenTo] = tokenToValue.Sub(tokenToValue, event.Value)

	zap.L().Debug(
		"swap by burn token in aave",
		zap.Stringer("transaction_hash", log.TxHash), zap.String("network", message.Network),
		zap.Stringer("token_from", tokenFrom), zap.Stringer("token_from_value", event.Value),
		zap.Stringer("token_to", tokenTo), zap.Stringer("token_to_value", event.Value),
	)

	return tokenMap, nil
}
