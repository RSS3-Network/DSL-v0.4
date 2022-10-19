package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/farcaster"
	"go.opentelemetry.io/otel"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
)

type Datasource struct {
	farcasterClient *farcaster.Client
}

func (d *Datasource) Name() string {
	return protocol.NetworkFarcaster
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkFarcaster,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("farcaster_datasource")
	_, trace := tracer.Start(ctx, "farcaster_datasource:Handle")

	defer func() { opentelemetry.Log(trace, message, transactions, err) }()

	activityList, err := d.farcasterClient.GetActivityList(ctx, message.Address)
	if err != nil {
		return nil, err
	}

	for _, cast := range activityList {
		sourceData, err := json.Marshal(cast)
		if err != nil {
			return nil, err
		}
		timestamp := time.UnixMilli(cast.Body.PublishedAt)

		var addressTo string
		if cast.Meta.ReplyParentUsername.Address != "" {
			addressTo = cast.Meta.ReplyParentUsername.Address
		} else {
			addressTo = message.Address
		}

		hash := fmt.Sprintf("%s%d", cast.MerkleRoot, cast.Body.Sequence)

		transactions = append(transactions, model.Transaction{
			// use timestamp as block number, as there is no actual blocknumber on farcaster
			BlockNumber: cast.Body.PublishedAt,
			Timestamp:   timestamp,
			// use MerkleRoot as hash, as there is no actual hash on farcaster
			Hash:        hash,
			AddressFrom: message.Address,
			AddressTo:   addressTo,
			// TODO: identify the correct platform
			Platform:   message.Network,
			Network:    message.Network,
			Source:     d.Name(),
			SourceData: sourceData,
			Transfers: []model.Transfer{
				// This is a virtual transfer
				{
					TransactionHash: hash,
					Timestamp:       timestamp,
					Index:           protocol.IndexVirtual,
					AddressFrom:     message.Address,
					AddressTo:       addressTo,
					Metadata:        metadata.Default,
					Network:         message.Network,
					SourceData:      sourceData,
				},
			},
		})
	}

	return transactions, nil
}

func New() datasource.Datasource {
	return &Datasource{
		farcasterClient: farcaster.NewClient(),
	}
}
