package liquidity

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/polygon"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
)

func (i *internal) handlePolygonStakingDelegatorClaimedRewards(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, fmt.Errorf("get ethereum client: %w", err)
	}

	staking, err := polygon.NewStaking(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new staking: %w", err)
	}

	event, err := staking.ParseDelegatorClaimedRewards(log)
	if err != nil {
		return nil, fmt.Errorf("parse delefator claimed rewards: %w", err)
	}

	tokenFrom, err := i.tokenClient.ERC20(ctx, message.Network, strings.ToLower("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0") /* MATIC */)
	if err != nil {
		return nil, fmt.Errorf("erc20: %w", err)
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityCollect, map[*token.ERC20]*big.Int{
		tokenFrom: event.Rewards,
	})
	if err != nil {
		return nil, fmt.Errorf("build liquidity metadata: %w", err)
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(event.User.String()),
		AddressTo:       strings.ToLower(log.Address.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handlePolygonStakingShareMinted(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, fmt.Errorf("get ethereum client: %w", err)
	}

	staking, err := polygon.NewStaking(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new staking: %w", err)
	}

	event, err := staking.ParseShareMinted(log)
	if err != nil {
		return nil, fmt.Errorf("parse delefator claimed rewards: %w", err)
	}

	tokenFrom, err := i.tokenClient.ERC20(ctx, message.Network, strings.ToLower("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0") /* MATIC */)
	if err != nil {
		return nil, fmt.Errorf("erc20: %w", err)
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityAdd, map[*token.ERC20]*big.Int{
		tokenFrom: event.Tokens,
	})
	if err != nil {
		return nil, fmt.Errorf("build liquidity metadata: %w", err)
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(event.User.String()),
		AddressTo:       strings.ToLower(log.Address.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}
