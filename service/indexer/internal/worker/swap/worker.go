package swap

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"gorm.io/gorm"
)

var _ worker.Worker = &service{}

type service struct {
	databaseClient *gorm.DB
}

func (s *service) Name() string {
	return "swap"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	var count int64

	if err := s.databaseClient.
		Model((*model.SwapPool)(nil)).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count == 0 {
		job := &Job{
			databaseClient: s.databaseClient,
		}

		job.Run()
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {
		if transfer.Source != moralis.Source {
			continue
		}

		var swapModel model.SwapPool

		if err := s.databaseClient.
			Model((*model.SwapPool)(nil)).
			Where(map[string]interface{}{
				"contract_address": transfer.AddressFrom,
				"network":          transfer.Network,
			}).
			First(&swapModel).
			Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		if swapModel.ContractAddress == "" {
			if err := s.databaseClient.
				Model((*model.SwapPool)(nil)).
				Where(map[string]interface{}{
					"contract_address": transfer.AddressTo,
					"network":          transfer.Network,
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

		metadataModel.Swap = &metadata.SwapPool{
			Name:     swapModel.Source,
			Network:  swapModel.Network,
			Token0:   swapModel.Token0,
			Token1:   swapModel.Token1,
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
	return []worker.Job{
		&Job{
			databaseClient: s.databaseClient,
		},
	}
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
	}
}
