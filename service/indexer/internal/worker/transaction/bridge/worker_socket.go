package bridge

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/socket"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

func (w *Worker) handleSocketSend(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("get ethereum client: %w", err)
	}

	filterer, err := socket.NewBridgeFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("build socket bridge contract: %w", err)
	}

	event, err := filterer.ParseSend(log)
	if err != nil {
		return nil, fmt.Errorf("parse socket send event: %w", err)
	}

	return w.buildTransfer(ctx, transaction, log, common.HexToAddress(transaction.AddressFrom), event.Receiver, protocol.PlatformSocket, event.DstChainId, &event.Token, event.Amount, filter.BridgeDeposit)
}
