package swap

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
)

var _ worker.Worker = &service{}

type service struct{}

func (s *service) Name() string {
	return "swap"
}

func (s *service) Network() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {
		if transfer.Source != moralis.Source {
			continue
		}

		// TODO Refactor this demo
		if strings.ToLower(transfer.AddressTo) == "0x42f0530351471dab7ec968476d19bd36af9ec52d" || strings.ToLower(transfer.AddressFrom) == "0x42f0530351471dab7ec968476d19bd36af9ec52d" {
			var metadataModel metadata.Metadata

			if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
				return nil, err
			}

			if strings.EqualFold(transfer.AddressFrom, message.Address) {
				transfer.Type = "swap_out"
			} else {
				transfer.Type = "swap_in"
			}

			metadataModel.Swap = &metadata.Swap{
				Name: "UniSwap",
				Pool: "DAI/USDT",
			}

			rawMetadata, err := json.Marshal(metadataModel)
			if err != nil {
				return nil, err
			}

			transfer.Metadata = rawMetadata

			internalTransfers = append(internalTransfers, transfer)
		}
	}

	return internalTransfers, nil
}

func New() worker.Worker {
	return &service{}
}
