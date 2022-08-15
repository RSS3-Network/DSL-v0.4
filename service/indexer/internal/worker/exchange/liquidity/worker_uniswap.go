package liquidity

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
)

func (i *internal) handleUniswapV2Mint(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV2(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV2(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseMint(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityAdd, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
	})
	if err != nil {
		return nil, err
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
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV2Burn(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV2(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV2(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseBurn(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityRemove, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
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
		AddressTo:       strings.ToLower(event.To.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV3Mint(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV3(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV3(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseMint(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityAdd, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
	})
	if err != nil {
		return nil, err
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
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV3Burn(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV3(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV3(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseBurn(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityRemove, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
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
		AddressTo:       strings.ToLower(event.Owner.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV3Collect(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, router Router) (*model.Transfer, error) {
	poolContract, err := uniswap.NewPoolV3(log.Address, i.ethereumClientMap[message.Network])
	if err != nil {
		return nil, err
	}

	tokenLeft, tokenRight, err := i.buildTokenPairV3(ctx, message.Network, poolContract)
	if err != nil {
		return nil, err
	}

	event, err := poolContract.ParseCollect(log)
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, router, filter.ExchangeLiquidityCollect, map[*token.ERC20]*big.Int{
		tokenLeft:  event.Amount0,
		tokenRight: event.Amount1,
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
		AddressTo:       strings.ToLower(event.Recipient.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        router.Name,
		Source:          ethereum.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) buildTokenPairV2(ctx context.Context, network string, poolContract *uniswap.PoolV2) (*token.ERC20, *token.ERC20, error) {
	tokenLeftAddress, err := poolContract.Token0(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}

	tokenRightAddress, err := poolContract.Token1(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}

	tokenLeft, err := i.tokenClient.ERC20(ctx, network, tokenLeftAddress.String())
	if err != nil {
		return nil, nil, err
	}

	tokenRight, err := i.tokenClient.ERC20(ctx, network, tokenRightAddress.String())
	if err != nil {
		return nil, nil, err
	}

	return tokenLeft, tokenRight, nil
}

func (i *internal) buildTokenPairV3(ctx context.Context, network string, poolContract *uniswap.PoolV3) (*token.ERC20, *token.ERC20, error) {
	tokenLeftAddress, err := poolContract.Token0(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}

	tokenRightAddress, err := poolContract.Token1(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}

	tokenLeft, err := i.tokenClient.ERC20(ctx, network, tokenLeftAddress.String())
	if err != nil {
		return nil, nil, err
	}

	tokenRight, err := i.tokenClient.ERC20(ctx, network, tokenRightAddress.String())
	if err != nil {
		return nil, nil, err
	}

	return tokenLeft, tokenRight, nil
}
