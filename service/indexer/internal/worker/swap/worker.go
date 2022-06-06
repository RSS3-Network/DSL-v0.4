package swap

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/dexpools"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ worker.Worker = &service{}

type service struct {
	databaseClient *gorm.DB
	poolClient     *dexpools.Client
}

func (s *service) Name() string {
	return "swap"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}
}

// Initialize TODO: convert into a cron job
func (s *service) Initialize(ctx context.Context) error {
	for _, dex := range protocol.SwapPools {
		result, err := s.poolClient.GetSwapPools(context.Background(), dex)
		if err != nil {
			return err
		}

		pools := make([]model.Swap, len(result))

		for i, pair := range result {
			pools[i] = model.Swap{
				Source:          dex.Name,
				Network:         dex.Network,
				ContractAddress: string(pair.ID),
				Protocol:        dex.Protocol,
			}
		}

		if err := s.databaseClient.
			Model((*model.Swap)(nil)).
			Clauses(clause.OnConflict{
				DoNothing: true,
			}).
			Create(&pools).
			Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {
		if transfer.Source != moralis.Source {
			continue
		}

		var swapModel model.Swap

		if err := s.databaseClient.
			Model((*model.Swap)(nil)).
			Where(map[string]interface{}{
				"contract_address": transfer.AddressFrom,
			}).
			First(&swapModel).
			Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		if swapModel.ContractAddress == "" {
			if err := s.databaseClient.
				Model((*model.Swap)(nil)).
				Where(map[string]interface{}{
					"contract_address": transfer.AddressTo,
				}).
				First(&swapModel).
				Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}

		if swapModel.ContractAddress == "" {
			continue
		}

		var metadataModel metadata.Metadata
		if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
			return nil, err
		}

		if strings.EqualFold(transfer.AddressFrom, message.Address) {
			transfer.Type = "swap_in"
		} else {
			transfer.Type = "swap_out"
		}

		metadataModel.Swap = &metadata.Swap{
			Name:     swapModel.Source,
			Network:  swapModel.Network,
			Protocol: swapModel.Protocol,
		}

		rawMetadata, err := json.Marshal(metadataModel)
		if err != nil {
			return nil, err
		}

		transfer.Metadata = rawMetadata

		internalTransfers = append(internalTransfers, transfer)
	}

	return internalTransfers, nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		poolClient:     dexpools.NewClient(),
	}
}
