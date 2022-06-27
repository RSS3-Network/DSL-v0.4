package swap

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"gorm.io/gorm"
)

const (
	Name = "swap"
)

var _ worker.Worker = &service{}

type service struct {
	employer       *shedlock.Employer
	databaseClient *gorm.DB
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	job := &Job{
		databaseClient: s.databaseClient,
	}

	return job.Run(func(ctx context.Context, duration time.Duration) error {
		return s.employer.Renewal(ctx, job.Name(), duration)
	})
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			// TODO ZkSync support
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

			transfer.Tag = filter.UpdateTag(filter.TagExchange, transfer.Tag)

			if transfer.Tag == filter.TagExchange {
				transfer.Type = filter.ExchangeSwap
			}

			transfer.Metadata = rawMetadata

			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Tag = filter.UpdateTag(transfer.Tag, value.Tag)
			value.Transfers = append(value.Transfers, transfer)

			internalTransactionMap[transaction.Hash] = value

			// transaction tag
			transaction.Tag = filter.UpdateTag(transfer.Tag, transaction.Tag)
		}
	}

	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&Job{
			databaseClient: s.databaseClient,
		},
	}
}

func New(employer *shedlock.Employer, databaseClient *gorm.DB) worker.Worker {
	return &service{
		employer:       employer,
		databaseClient: databaseClient,
	}
}
