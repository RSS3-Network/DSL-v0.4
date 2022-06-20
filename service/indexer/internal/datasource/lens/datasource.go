package lens

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/shurcooL/graphql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/lens"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
)

const (
	Source = "lens"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	lensClient     *lens.Client
	databaseClient *gorm.DB
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkArweave,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	transactions := make([]model.Transaction, 0)

	// read the last cursor value from the database
	var lensCursor model.LensCursor

	if err := d.databaseClient.Model((*model.LensCursor)(nil)).
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
		Address:    graphql.String(strings.ToLower(message.Address)),
		Cursor:     graphql.String(lensCursor.Cursor),
		TotalCount: 0,
	}

	result, err := d.lensClient.GetAllPublicationsByAddress(ctx, &options)
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

		if publication.Type == "Comment" {
			var target json.RawMessage

			switch publication.CommentOn.Post.Type {
			case "Post":
				target, err = json.Marshal(publication.CommentOn.Post)
				if err != nil {
					return nil, err
				}
			case "Comment":
				target, err = json.Marshal(publication.CommentOn.Comment)
				if err != nil {
					return nil, err
				}
			}
			metadataModel.Lens.Target = target
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
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
			Transfers: []model.Transfer{
				{
					Type:            Source,
					TransactionHash: string(publication.ID),
					Timestamp:       publication.CreatedAt,
					AddressFrom:     message.Address,
					AddressTo:       "",
					Metadata:        rawMetadata,
					Network:         protocol.NetworkPolygon,
					Platform:        string(publication.Platform),
					RelatedUrls:     []string{string(publication.RelatedURL)},
					Source:          d.Name(),
					SourceData:      sourceData,
				},
			},
		})
	}

	// update the latest cursor value in the database
	lensCursor.Cursor = string(options.Cursor)

	if err := d.databaseClient.Model((*model.LensCursor)(nil)).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "address"}},
			DoUpdates: clause.AssignmentColumns([]string{"cursor"}),
		}).
		Create(lensCursor).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func New(databaseClient *gorm.DB) datasource.Datasource {
	return &Datasource{
		lensClient:     lens.NewClient(),
		databaseClient: databaseClient,
	}
}
