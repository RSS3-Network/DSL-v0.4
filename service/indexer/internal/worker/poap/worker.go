package poap

import (
	"context"
	"encoding/json"

	"github.com/naturalselectionlabs/pregod/common/blockscout"
	"github.com/naturalselectionlabs/pregod/common/constant"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/poap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	blockscoutdatasource "github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
)

const (
	Name = "poap"

	ContractAddress = "0x22c1f6050e56d2876009903609a2cc3fef83b415"
)

var _ worker.Worker = (*service)(nil)

type service struct {
	poapClient *poap.Client
}

func (s *service) Name() string {
	return Name
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
		if transfer.Source != blockscoutdatasource.Name {
			continue
		}

		dataSource := blockscout.Transaction{}
		if err := json.Unmarshal(transfer.SourceData, &dataSource); err != nil {
			return nil, err
		}

		if dataSource.ContractAddress != ContractAddress {
			continue
		}

		var metadataModel metadata.Metadata

		if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
			return nil, err
		}

		token, err := s.poapClient.GetToken(ctx, dataSource.TokenID.BigInt().Int64())
		if err != nil {
			return nil, err
		}

		metadataModel.POAP = &metadata.POAP{
			ID:          token.Event.ID,
			Name:        token.Event.Name,
			ImageURL:    token.Event.ImageURL,
			Description: token.Event.Description,
			Year:        token.Event.Year,
			StartDate:   token.Event.StartDate,
			EndDate:     token.Event.EndDate,
			ExpiryDate:  token.Event.ExpiryDate,
			TokenID:     token.TokenID,
		}

		rawMetadata, err := json.Marshal(metadataModel)
		if err != nil {
			return nil, err
		}

		transfer.Metadata = rawMetadata
		transfer.Tags = append(transfer.Tags, constant.TransferTagPoap.String())

		internalTransfers = append(internalTransfers, transfer)
	}

	return internalTransfers, nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &service{
		poapClient: poap.New(),
	}
}
