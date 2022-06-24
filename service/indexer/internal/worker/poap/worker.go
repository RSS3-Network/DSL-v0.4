package poap

import (
	"context"
	"encoding/json"

	"github.com/naturalselectionlabs/pregod/common/protocol/filter"

	"github.com/naturalselectionlabs/pregod/common/blockscout"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/poap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
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

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			dataSource := blockscout.Transaction{}
			if err := json.Unmarshal(transfer.SourceData, &dataSource); err != nil {
				return nil, err
			}

			if dataSource.ContractAddress != ContractAddress {
				continue
			}

			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				internalTransactionMap[transaction.Hash] = transaction
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

			transfer.Platform = Name
			transfer.Metadata = rawMetadata
			transfer.Tag = filter.TagCollectible
			transfer.Type = filter.NFTPoap

			value.Transfers = append(value.Transfers, transfer)

			internalTransactionMap[value.Hash] = value
		}
	}

	internalTransfers := make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		internalTransaction.Platform = Name

		internalTransfers = append(internalTransfers, internalTransaction)
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
