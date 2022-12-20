package liquidity

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lido"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
)

func (i *internal) handleLidoSubmitted(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, fmt.Errorf("get ethereum client: %w", err)
	}

	lidoContract, err := lido.NewETH(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new eth: %w", err)
	}

	event, err := lidoContract.ParseSubmitted(log)
	if err != nil {
		return nil, fmt.Errorf("parse submitted: %w", err)
	}

	tokenFrom, err := i.tokenClient.Native(ctx, message.Network)
	if err != nil {
		return nil, fmt.Errorf("native: %w", err)
	}

	tokenFromERC20 := token.ERC20{
		Name:     tokenFrom.Name,
		Symbol:   tokenFrom.Symbol,
		Decimals: tokenFrom.Decimals,
		Logo:     tokenFrom.Logo,
	}

	// Temporary conversion
	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityAdd, map[*token.ERC20]*big.Int{
		&tokenFromERC20: event.Amount,
	})
	if err != nil {
		return nil, fmt.Errorf("build liquidity metadata: %w", err)
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
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleLidoSubmitEvent(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, fmt.Errorf("get ethereum client: %w", err)
	}

	lidoContract, err := lido.NewMatic(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new matic: %w", err)
	}

	event, err := lidoContract.ParseSubmitEvent(log)
	if err != nil {
		return nil, fmt.Errorf("parse submit event: %w", err)
	}

	tokenFrom, err := i.tokenClient.ERC20(ctx, message.Network, strings.ToLower("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0") /* MATIC */)
	if err != nil {
		return nil, fmt.Errorf("erc20: %w", err)
	}

	// Temporary conversion
	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityAdd, map[*token.ERC20]*big.Int{
		tokenFrom: event.Amount,
	})
	if err != nil {
		return nil, fmt.Errorf("build liquidity metadata: %w", err)
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(event.From.String()),
		AddressTo:       strings.ToLower(log.Address.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}
