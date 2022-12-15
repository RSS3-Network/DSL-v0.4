package bridge

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/hop"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

var tokenMap = map[common.Address]common.Address{
	hop.AddressL2AMMWrapperPolygonETH: common.HexToAddress("0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619"),
}

func (w *Worker) handleHopTransferSent(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	filterer, err := hop.NewL2BridgeFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new l2 bridge filterer: %w", err)
	}

	ammWrapperAddress := common.HexToAddress(transaction.AddressTo)

	tokenAddress, exists := tokenMap[ammWrapperAddress]
	if !exists {
		return nil, fmt.Errorf("invalid token address: %s", log.Address)
	}

	event, err := filterer.ParseTransferSent(log)
	if err != nil {
		return nil, fmt.Errorf("parse transfer sent event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, common.HexToAddress(transaction.AddressFrom), event.Recipient, platform, event.ChainId.Uint64(), &tokenAddress, event.Amount, filter.TransactionBridge)
}
