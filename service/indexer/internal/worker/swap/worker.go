package swap

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/cache"
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

		var toPool model.Swap

		toErr := cache.HGet(context.Background(), "swappools", transfer.AddressTo, &toPool)

		var fromPool model.Swap

		fromErr := cache.HGet(context.Background(), "swappools", transfer.AddressFrom, &fromPool)

		if (toErr != nil || toErr == redis.Nil) || (fromErr != nil || fromErr == redis.Nil) {
			var metadataModel metadata.Metadata

			if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
				return nil, err
			}

			if strings.EqualFold(transfer.AddressFrom, message.Address) {
				transfer.Type = "swap_in"
				metadataModel.Swap = &metadata.Swap{
					Name: fromPool.Name,
					Pool: transfer.AddressFrom,
				}

			} else {
				transfer.Type = "swap_out"
				metadataModel.Swap = &metadata.Swap{
					Name: toPool.Name,
					Pool: transfer.AddressTo,
				}
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
