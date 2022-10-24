package everipedia

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/everipedia"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/everipedia/graphql"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	everipedia_contract "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/everipedia/contract"
	"go.opentelemetry.io/otel"
)

type Datasource struct {
	everipediaClient *everipedia.Client
}

func (d *Datasource) Name() string {
	return protocol.PlatformEveripedia
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkPolygon,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("Everipedia_datasource")
	_, trace := tracer.Start(ctx, "Everipedia_datasource:Handle")

	defer func() { opentelemetry.Log(trace, message, transactions, err) }()

	activityList, err := d.everipediaClient.GetUserActivityList(ctx, message.Address)
	if err != nil {
		return nil, err
	}

	for _, activity := range activityList {
		var action string

		if activity.Type == graphqlx.StatusCreated {
			action = filter.SocialProfileCreate
		} else if activity.Type == graphqlx.StatusUpdated {
			action = filter.SocialProfileUpdate
		}

		for _, content := range activity.Content {
			sourceData, err := json.Marshal(content)
			if err != nil {
				continue
			}
			transactions = append(transactions, model.Transaction{
				BlockNumber: int64(activity.Block),
				Timestamp:   activity.Datetime,
				// use MerkleRoot as hash, as there is no actual hash on farcaster
				Hash:        content.TransactionHash,
				AddressFrom: strings.ToLower(everipedia_contract.AddressEveripediaSign.String()),
				AddressTo:   strings.ToLower(everipedia_contract.AddressEveripedia.String()),
				Owner:       strings.ToLower(activity.User.Id),
				Platform:    protocol.PlatformEveripedia,
				Network:     message.Network,
				Source:      d.Name(),
				SourceData:  sourceData,
				Tag:         filter.TagSocial,
				Type:        action,
				Transfers: []model.Transfer{
					// This is a virtual transfer
					{
						TransactionHash: content.TransactionHash,
						Timestamp:       activity.Datetime,
						Index:           protocol.IndexVirtual,
						AddressFrom:     strings.ToLower(everipedia_contract.AddressEveripediaSign.String()),
						AddressTo:       strings.ToLower(everipedia_contract.AddressEveripedia.String()),
						Metadata:        metadata.Default,
						Network:         message.Network,
						SourceData:      sourceData,
						Tag:             filter.TagSocial,
						Type:            action,
					},
				},
			})
		}

	}

	return transactions, nil
}

func New() datasource.Datasource {
	return &Datasource{
		everipediaClient: everipedia.NewClient(),
	}
}
