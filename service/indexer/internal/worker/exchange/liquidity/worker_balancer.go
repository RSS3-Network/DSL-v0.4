package liquidity

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/balancer"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/shopspring/decimal"
)

func (i *internal) handleBalancerPoolBalanceChanged(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform Platform) (*model.Transfer, error) {
	if log.Address != balancer.AddressVault {
		return nil, fmt.Errorf("unexpected address: %s", log.Address)
	}

	ethereumClient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, fmt.Errorf("get ethereum client: %w", err)
	}

	contract, err := balancer.NewVault(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new vault filterer: %w", err)
	}

	event, err := contract.ParsePoolBalanceChanged(log)
	if err != nil {
		return nil, fmt.Errorf("parse pool balance changed: %w", err)
	}

	tokenMap := map[*token.ERC20]*big.Int{}

	action := filter.ExchangeLiquidityAdd

	for index, address := range event.Tokens {
		switch event.Deltas[index].Cmp(decimal.Zero.BigInt()) {
		case -1:
			action = filter.ExchangeLiquidityRemove
		case 0:
			continue
		}

		erc20, err := i.tokenClient.ERC20(ctx, message.Network, address.String())
		if err != nil {
			return nil, fmt.Errorf("get token: %w", err)
		}

		tokenMap[erc20] = event.Deltas[index]
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, platform, action, tokenMap)
	if err != nil {
		return nil, fmt.Errorf("build liquidity metadata: %w", err)
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(event.LiquidityProvider.String()),
		AddressTo:       strings.ToLower(log.Address.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        platform.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}
