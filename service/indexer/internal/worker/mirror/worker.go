package mirror

import (
	"context"
	"encoding/json"
	"io"

	"github.com/naturalselectionlabs/pregod/common/arweave"
	graphqlx "github.com/naturalselectionlabs/pregod/common/arweave/graphql"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
)

const (
	Name = "mirror"
)

var _ worker.Worker = &service{}

type service struct {
	arweaveClient *arweave.Client
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkArweave,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			transactionEdge := graphqlx.TransactionEdge{}

			if err := json.Unmarshal(transfer.SourceData, &transactionEdge); err != nil {
				return nil, err
			}

			// It may be a transaction for another dApp or a normal transfer
			if transactionEdge.Node.Owner.Address != arweave.AddressMirror {
				continue
			}

			transfer.Platform = Name

			var metadataModel metadata.Metadata

			if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
				return nil, err
			}

			mirrorMetadata := metadata.Mirror{}

			for _, tag := range transactionEdge.Node.Tags {
				switch tag.Name {
				case "Content-Type":
					mirrorMetadata.ContentType = string(tag.Value)
				case "Contributor":
					mirrorMetadata.Contributor = string(tag.Value)
				case "Content-Digest":
					mirrorMetadata.ContentDigest = string(tag.Value)
				case "Original-Content-Digest":
					mirrorMetadata.OriginalContentDigest = string(tag.Value)
				}
			}

			var err error

			// Get the article text content
			if mirrorMetadata.Content, err = s.getContent(ctx, transactionEdge.Node.ID.(string)); err != nil {
				return nil, err
			}

			metadataModel.Mirror = &mirrorMetadata

			rawMetadata, err := json.Marshal(metadataModel)
			if err != nil {
				return nil, err
			}

			transfer.Metadata = rawMetadata
			transfer.Tag = filter.TagSocial
			transfer.Type = filter.SocialPost

			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Transfers = append(value.Transfers, transfer)
			internalTransactionMap[transaction.Hash] = value
		}
	}

	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range internalTransactionMap {
		transaction.Platform = Name

		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func (s *service) getContent(ctx context.Context, transactionHash string) (json.RawMessage, error) {
	reader, err := s.arweaveClient.GetFile(ctx, transactionHash)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	_ = reader.Close()

	return data, nil
}

func New() worker.Worker {
	return &service{
		arweaveClient: arweave.NewClient(),
	}
}
