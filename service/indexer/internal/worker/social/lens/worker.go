package lens

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	lensClient "github.com/naturalselectionlabs/pregod/common/worker/lens"

	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	Name = protocol.PlatfromLens
)

var _ worker.Worker = (*service)(nil)

type service struct {
	databaseClient *gorm.DB
	lensClient     *lensClient.Client
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		lensClient:     lensClient.NewClient(),
	}
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkPolygon,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("lens_worker")
	_, trace := tracer.Start(ctx, "lens_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	// read the last cursor value from the database
	var lastLens model.Transaction
	var cursor string

	if err := s.databaseClient.Model((*model.Transaction)(nil)).
		Where(map[string]interface{}{
			"address_from": message.Address,
			"source":       Name,
		}).Order("timestamp DESC").First(&lastLens).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// if it's a new address, the cursor will be empty
	if len(lastLens.Hash) == 0 {
		cursor = "{}"
	} else {
		cursor = fmt.Sprintf(`{"entityIdentifier":"%s","timestamp":%d,"cursorDirection":"BEFORE"}`, lastLens.Hash, lastLens.Timestamp.Unix())
	}

	options := lensClient.Options{
		Address: graphql.String(message.Address),
		Cursor:  graphql.String(cursor),
	}

	// handle publications
	result, err := s.lensClient.GetAllPublicationsByAddress(ctx, &options)
	if err != nil {
		return nil, err
	}

	// returns when no results are found
	if len(result) == 0 {
		return transactions, nil
	}

	for _, publication := range result {
		sourceData, err := json.Marshal(publication)
		if err != nil {
			return nil, err
		}

		var metadataModel metadata.Metadata

		media, err := json.Marshal(publication.Metadata.Media)
		if err != nil {
			return nil, err
		}

		metadataModel.Post = &metadata.Post{
			TypeOnPlatform: string(publication.Type),
			Body:           string(publication.Metadata.Description),
			Media:          media,
		}

		if publication.Type != "Post" {
			media, err := json.Marshal(publication.Target.Metadata.Media)
			if err != nil {
				return nil, err
			}

			targetContent := metadata.Post{
				TypeOnPlatform: string(publication.Target.Type),
				Body:           string(publication.Target.Metadata.Description),
				Media:          media,
			}

			metadataModel.Post.Target = &targetContent
		}

		rawMetadata, err := json.Marshal(metadataModel)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, model.Transaction{
			Timestamp:   publication.CreatedAt,
			Hash:        string(publication.ID),
			AddressFrom: message.Address,
			AddressTo:   "",
			Tag:         filter.TagSocial,
			Network:     message.Network,
			Platform:    string(publication.Platform),
			Source:      s.Name(),
			SourceData:  sourceData,
			Transfers: []model.Transfer{
				{
					Type:            getType(string(publication.Type)),
					Tag:             filter.TagSocial,
					TransactionHash: string(publication.ID),
					Timestamp:       publication.CreatedAt,
					AddressFrom:     message.Address,
					AddressTo:       "",
					Metadata:        rawMetadata,
					Network:         protocol.NetworkPolygon,
					Platform:        string(publication.Platform),
					RelatedUrls:     []string{string(publication.RelatedURL)},
					Source:          s.Name(),
					SourceData:      sourceData,
				},
			},
		})
	}

	return transactions, nil
}

func getType(pubType string) string {
	switch pubType {
	case "Post":
		return filter.SocialPost
	case "Comment":
		return filter.SocialComment
	case "Mirror":
		return filter.SocialShare
	}

	return filter.SocialPost
}
