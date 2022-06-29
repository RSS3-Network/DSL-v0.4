package handler

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
	"go.opentelemetry.io/otel"
)

var _ Interface = (*profile)(nil)

type profile struct {
	contract any
}

func (p *profile) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler_profileC")

	ctx, snap := tracer.Start(ctx, "worker_crossbell_handler_profileC:Handle")

	defer snap.End()

	var log types.Log

	if err := json.Unmarshal(transfer.SourceData, &log); err != nil {
		return nil, err
	}

	switch log.Topics[0] {
	case contract.EventHashProfileCreated:
		return p.handleProfileCreated(ctx, transfer, log)
	default:
		return nil, contract.ErrorUnknownUnknownEvent
	}
}

func (p *profile) handleProfileCreated(ctx context.Context, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler_profileC")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler_profileC:handleProfileCreated")

	defer snap.End()

	event := contract.EventProfileCreated{}

	event.ProfileID = big.NewInt(0)
	event.ProfileID.SetString(log.Topics[1].Hex(), 0)

	event.Creator = common.HexToAddress(log.Topics[2].Hex())
	event.To = common.HexToAddress(log.Topics[3].Hex())

	var err error

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{}); err != nil {
		return nil, err
	}

	return &transfer, nil
}
