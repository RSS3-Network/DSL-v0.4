package iqwiki

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/iqwiki"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/iqwiki/graphql"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	iqwiki_contract "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/iqwiki/contract"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type Datasource struct {
	iqwikiClient *iqwiki.Client
}

func (d *Datasource) Name() string {
	return protocol.PlatformIQWiki
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkPolygon,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("IQ.Wiki_datasource")
	_, trace := tracer.Start(ctx, "IQ.Wiki_datasource:Handle")

	defer func() { opentelemetry.Log(trace, message, transactions, err) }()

	activityList, err := d.iqwikiClient.GetUserActivityList(ctx, message.Address)
	if err != nil {
		return nil, err
	}
	blockNum := d.getEveTransactions(strings.ToLower(message.Address))
	if message.Reindex {
		blockNum = 0
	}

	for _, activity := range activityList {

		if activity.Block <= int(blockNum) {
			continue
		}
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
				Hash:        content.TransactionHash,
				AddressFrom: strings.ToLower(iqwiki_contract.AddressEveripediaSign.String()),
				AddressTo:   strings.ToLower(iqwiki_contract.AddressEveripedia.String()),
				Owner:       strings.ToLower(activity.User.Id),
				Platform:    protocol.PlatformIQWiki,
				Network:     message.Network,
				Source:      d.Name(),
				SourceData:  sourceData,
				Tag:         filter.TagSocial,
				Type:        filter.SocialWiki,
				Transfers: []model.Transfer{
					// This is a virtual transfer
					{
						TransactionHash: content.TransactionHash,
						Timestamp:       activity.Datetime,
						Index:           protocol.IndexVirtual,
						AddressFrom:     strings.ToLower(iqwiki_contract.AddressEveripediaSign.String()),
						AddressTo:       strings.ToLower(iqwiki_contract.AddressEveripedia.String()),
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

func (d *Datasource) getEveTransactions(address string) int64 {
	var result struct {
		BlockNumber int64 `gorm:"column:block_number"`
	}

	if err := database.Global().
		Model(&model.Transaction{}).
		Select("COALESCE(block_number, 0) AS block_number").
		Where("owner = ?", address).
		Where("address_from", strings.ToLower(iqwiki_contract.AddressEveripediaSign.String())).
		Where("address_to", strings.ToLower(iqwiki_contract.AddressEveripedia.String())).
		Where("success IS TRUE").
		Order("block_number DESC").
		Limit(1).
		Find(&result).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0
	}

	return result.BlockNumber
}

func New() datasource.Datasource {
	return &Datasource{
		iqwikiClient: iqwiki.NewClient(),
	}
}
