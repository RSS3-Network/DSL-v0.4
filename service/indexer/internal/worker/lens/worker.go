package lens

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/lens"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shurcooL/graphql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	Name = protocol.PlatfromLens
)

var _ worker.Worker = (*service)(nil)

type service struct {
	databaseClient *gorm.DB
	lensClient     *lens.Client
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		lensClient:     lens.NewClient(),
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

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	// read the last cursor value from the database
	var lensCursor model.LensCursor

	if err := s.databaseClient.Model((*model.LensCursor)(nil)).
		Where(map[string]interface{}{
			"address": strings.ToLower(message.Address),
		}).First(&lensCursor).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// if it's a new address, the cursor will be empty
	if lensCursor.Address == "" {
		lensCursor = model.LensCursor{
			Address: strings.ToLower(message.Address),
			Cursor:  "{}",
		}
	}

	options := lens.Options{
		Address: graphql.String(strings.ToLower(message.Address)),
		Cursor:  graphql.String(lensCursor.Cursor),
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

		metadataModel.Lens = &metadata.Lens{
			Type:    string(publication.Type),
			Content: string(publication.Metadata.Description),
			Media:   media,
		}

		if publication.Type != "Post" {
			metadataModel.Lens.Target = publication.Target
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
			Platform:    Name,
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

	// update the latest cursor value in the database
	lensCursor.Cursor = string(options.Cursor)

	if err := s.databaseClient.Model((*model.LensCursor)(nil)).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "address"}},
			DoUpdates: clause.AssignmentColumns([]string{"cursor"}),
		}).
		Create(lensCursor).Error; err != nil {
		return nil, err
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
