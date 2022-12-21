package liquidity

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
)

var ErrorInvalidNumberOfToken = errors.New("invalid number of token")

func (i *internal) handleUniswapV2Mint(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform Platform) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := uniswap.NewPoolV2(log.Address, ethclient)
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

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, platform, filter.ExchangeLiquidityAdd, map[*token.ERC20]*big.Int{
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
		Platform:        platform.Name,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV2Burn(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform Platform) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := uniswap.NewPoolV2(log.Address, ethclient)
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

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, platform, filter.ExchangeLiquidityRemove, map[*token.ERC20]*big.Int{
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
		Source:          transaction.Source,
		Platform:        platform.Name,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV3Mint(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform Platform) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := uniswap.NewPoolV3(log.Address, ethclient)
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

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, platform, filter.ExchangeLiquidityAdd, map[*token.ERC20]*big.Int{
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
		Source:          transaction.Source,
		Platform:        platform.Name,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV3Burn(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform Platform) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := uniswap.NewPoolV3(log.Address, ethclient)
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

	if event.Amount0.Cmp(big.NewInt(0)) == 0 && event.Amount1.Cmp(big.NewInt(0)) == 0 {
		return nil, ErrorInvalidNumberOfToken
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, platform, filter.ExchangeLiquidityRemove, map[*token.ERC20]*big.Int{
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
		Source:          transaction.Source,
		Platform:        platform.Name,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleUniswapV3Collect(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform Platform) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	positionContract, err := uniswap.NewPosition(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := positionContract.ParseCollect(log)
	if err != nil {
		return nil, err
	}

	positions, err := positionContract.Positions(&bind.CallOpts{}, event.TokenId)
	if err != nil {
		return nil, err
	}

	tokenLeft, err := i.tokenClient.ERC20(ctx, message.Network, positions.Token0.String())
	if err != nil {
		return nil, err
	}

	tokenRight, err := i.tokenClient.ERC20(ctx, message.Network, positions.Token1.String())
	if err != nil {
		return nil, err
	}

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, platform, filter.ExchangeLiquidityCollect, map[*token.ERC20]*big.Int{
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
		Source:          transaction.Source,
		Platform:        platform.Name,
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
