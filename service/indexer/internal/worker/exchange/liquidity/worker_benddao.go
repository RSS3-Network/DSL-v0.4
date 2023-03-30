package liquidity

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/benddao"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
)

func (i *internal) handleBendDAODeposit(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform Platform) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := benddao.NewBendDAO(log.Address, ethclient)
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

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, platform, filter.ExchangeLiquiditySupply, map[*token.ERC20]*big.Int{
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
		AddressFrom:     strings.ToLower(event.OnBehalfOf.String()),
		AddressTo:       strings.ToLower(event.User.String()),
		Metadata:        liquidityMetadata,
		Network:         transaction.Network,
		Platform:        platform.Name,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}

func (i *internal) handleBendDAOWithdraw(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform Platform) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	poolContract, err := benddao.NewBendDAO(log.Address, ethclient)
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

	liquidityMetadata, err := i.buildLiquidityMetadata(ctx, platform, filter.ExchangeLiquidityWithdraw, map[*token.ERC20]*big.Int{
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
		Platform:        platform.Name,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}
