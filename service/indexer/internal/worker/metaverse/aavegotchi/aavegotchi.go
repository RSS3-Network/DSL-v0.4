package aavegotchi

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	contract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/aavegotchi"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type Handler struct{}

type actionTagAndType struct {
	Tag  string
	Type string
}

var actionRouter = map[common.Hash]actionTagAndType{
	contract.EventERC1155ExecutedListing: {
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseTrade,
	},
	contract.EventERC1155ListingCancelled: {
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseUnlist,
	},
	contract.EventERC1155ListingAdd: {
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseList,
	},
	contract.EventERC721ExecutedListing: {
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseTrade,
	},
	contract.EventERC721ListingAdd: {
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseList,
	},
	contract.EventClaimAavegotchi: {
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseClaim,
	},
}

func (h *Handler) Handle(ctx context.Context, transaction *model.Transaction) error {
	tracer := otel.Tracer("worker_metaverse")
	_, snap := tracer.Start(ctx, "worker_metaverse:aavegotchi")
	defer snap.End()

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		zap.L().Error("failed to unmarshal source data", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

		return err
	}

	var action actionTagAndType
	for _, log := range sourceData.Receipt.Logs {
		for _, topic := range log.Topics {
			if _, exist := actionRouter[topic]; exist {
				action = actionRouter[topic]
				goto update
			}
		}
	}

update:
	var internalTransfers []model.Transfer

	for _, transfer := range transaction.Transfers {
		if action.Tag != "" {
			transfer.Tag, transfer.Type = filter.UpdateTagAndType(action.Tag, transfer.Tag, action.Type, transfer.Type)
			transfer.Platform = protocol.PlatformAavegotchi
		}

		internalTransfers = append(internalTransfers, transfer)
	}

	transaction.Transfers = internalTransfers
	return nil
}
