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
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
)

var _ worker.Worker = &service{}

type service struct {
	arweaveClient *arweave.Client
}

func (s *service) Name() string {
	return "mirror"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkArweave,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {
		transactionEdge := graphqlx.TransactionEdge{}

		if err := json.Unmarshal(transfer.SourceData, &transactionEdge); err != nil {
			return nil, err
		}

		if transactionEdge.Node.Owner.Address != arweave.AddressMirror {
			continue
		}

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

		reader, err := s.arweaveClient.GetFile(ctx, transactionEdge.Node.ID.(string))
		if err != nil {
			return nil, err
		}

		data, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		_ = reader.Close()

		mirrorMetadata.Content = data

		metadataModel.Mirror = &mirrorMetadata

		rawMetadata, err := json.Marshal(metadataModel)
		if err != nil {
			return nil, err
		}

		transfer.Metadata = rawMetadata

		internalTransfers = append(internalTransfers, transfer)
	}

	return internalTransfers, nil
}

func New() worker.Worker {
	return &service{
		arweaveClient: arweave.NewClient(),
	}
}
