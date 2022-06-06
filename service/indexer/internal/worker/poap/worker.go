package poap

import (
	"context"
	"encoding/json"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/poap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
)

var _ worker.Worker = (*service)(nil)

type service struct{}

func (s *service) Name() string {
	return "poap"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkXDAI,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {
		if transfer.Source != s.Name() {
			continue
		}

		var metadataModel metadata.Metadata

		if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
			return nil, err
		}

		action := poap.Action{}

		if err := json.Unmarshal(transfer.SourceData, &action); err != nil {
			return nil, err
		}

		metadataModel.POAP = &metadata.POAP{
			ID:          action.Event.ID,
			Name:        action.Event.Name,
			ImageURL:    action.Event.ImageURL,
			Description: action.Event.Description,
			Year:        action.Event.Year,
			StartDate:   action.Event.StartDate,
			EndDate:     action.Event.EndDate,
			ExpiryDate:  action.Event.ExpiryDate,
			Supply:      action.Event.Supply,
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

func New() worker.Worker {
	return &service{}
}
