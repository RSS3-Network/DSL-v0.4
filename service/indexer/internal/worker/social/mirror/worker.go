package mirror

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	arweave2 "github.com/naturalselectionlabs/pregod/common/worker/arweave"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/arweave/graphql"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

const (
	Name = protocol.PlatfromMirror
)

var _ worker.Worker = &service{}

type service struct {
	arweaveClient *arweave2.Client
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

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("mirror_worker")
	_, trace := tracer.Start(ctx, "mirror_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			transactionEdge := graphqlx.TransactionEdge{}

			if err := json.Unmarshal(transfer.SourceData, &transactionEdge); err != nil {
				return nil, err
			}

			// It may be a transaction for another dApp or a normal transfer
			if transactionEdge.Node.Owner.Address != arweave2.AddressMirror {
				continue
			}

			transfer.Platform = Name

			var metadataModel metadata.Metadata

			if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
				return nil, err
			}

			mirrorMetadata := Mirror{}

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

			// Get the article text content
			content, err := s.getContent(ctx, transactionEdge.Node.ID.(string))
			if err != nil {
				logrus.Errorf("[mirror_worker] Handle: getContent error, %v", err)
				return nil, err
			}

			mirrorContent := MirrorContent{}

			if err = json.Unmarshal(content, &mirrorContent); err != nil {
				logrus.Errorf("[mirror_worker] Handle: json unmarshal error, %v", err)
				return nil, err
			}

			metadataModel.Content = &social.Content{
				Title: mirrorContent.Content.Title,
				Body:  mirrorContent.Content.Body,
			}

			rawMetadata, err := json.Marshal(metadataModel)
			if err != nil {
				return nil, err
			}

			transfer.AddressTo = strings.ToLower(string(transactionEdge.Node.Owner.Address))
			transfer.AddressFrom = strings.ToLower(mirrorMetadata.Contributor)
			transfer.Metadata = rawMetadata
			transfer.Tag = filter.UpdateTag(filter.TagSocial, transfer.Tag)
			transfer.RelatedUrls = append(
				transfer.RelatedUrls,
				fmt.Sprintf("https://mirror.xyz/%s/%s", transfer.AddressFrom, mirrorMetadata.ContentDigest),
			)

			if transfer.Tag == filter.TagSocial {
				transfer.Type = filter.SocialPost
			}

			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction
				value.AddressTo = strings.ToLower(string(transactionEdge.Node.Owner.Address))
				value.AddressFrom = strings.ToLower(mirrorMetadata.Contributor)

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Tag = filter.UpdateTag(transfer.Tag, value.Tag)
			value.Transfers = append(value.Transfers, transfer)

			// parse timestamp
			if mirrorContent.Content.Timestamp.BigInt().Int64() > 0 {
				value.Timestamp = time.Unix(mirrorContent.Content.Timestamp.BigInt().Int64(), 0)
			}

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
		arweaveClient: arweave2.NewClient(),
	}
}
