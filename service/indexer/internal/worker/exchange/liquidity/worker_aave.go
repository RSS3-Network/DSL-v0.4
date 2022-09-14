package liquidity

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/aave"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
)

func (i *internal) handleAAVEV2Deposit(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := aave.NewPoolV2(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseDeposit(log)
	if err != nil {
		return nil, err
	}

	tokenSingle, err := i.tokenClient.ERC20(ctx, message.Network, event.Reserve.String())
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquiditySupply, map[*token.ERC20]*big.Int{
		tokenSingle: event.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.User.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleAAVEV2Borrow(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := aave.NewPoolV2(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseBorrow(log)
	if err != nil {
		return nil, err
	}

	tokenSingle, err := i.tokenClient.ERC20(ctx, message.Network, event.Reserve.String())
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityBorrow, map[*token.ERC20]*big.Int{
		tokenSingle: event.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.User.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleAAVEV2Repay(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := aave.NewPoolV2(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseRepay(log)
	if err != nil {
		return nil, err
	}

	tokenSingle, err := i.tokenClient.ERC20(ctx, message.Network, event.Reserve.String())
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityRepay, map[*token.ERC20]*big.Int{
		tokenSingle: event.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.User.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleAAVEV2Withdraw(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := aave.NewPoolV2(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseWithdraw(log)
	if err != nil {
		return nil, err
	}

	tokenSingle, err := i.tokenClient.ERC20(ctx, message.Network, event.Reserve.String())
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityWithdraw, map[*token.ERC20]*big.Int{
		tokenSingle: event.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.User.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleAAVEV3Supply(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := aave.NewPoolV3(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseSupply(log)
	if err != nil {
		return nil, err
	}

	tokenSingle, err := i.tokenClient.ERC20(ctx, message.Network, event.Reserve.String())
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquiditySupply, map[*token.ERC20]*big.Int{
		tokenSingle: event.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.User.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleAAVEV3Borrow(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := aave.NewPoolV3(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseBorrow(log)
	if err != nil {
		return nil, err
	}

	tokenSingle, err := i.tokenClient.ERC20(ctx, message.Network, event.Reserve.String())
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityBorrow, map[*token.ERC20]*big.Int{
		tokenSingle: event.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.User.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleAAVEV3Repay(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := aave.NewPoolV3(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseRepay(log)
	if err != nil {
		return nil, err
	}

	tokenSingle, err := i.tokenClient.ERC20(ctx, message.Network, event.Reserve.String())
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityRepay, map[*token.ERC20]*big.Int{
		tokenSingle: event.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.User.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleAAVEV3Withdraw(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := aave.NewPoolV3(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseWithdraw(log)
	if err != nil {
		return nil, err
	}

	tokenSingle, err := i.tokenClient.ERC20(ctx, message.Network, event.Reserve.String())
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityWithdraw, map[*token.ERC20]*big.Int{
		tokenSingle: event.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(log.Address.String()),
		AddressTo:       strings.ToLower(event.User.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}
