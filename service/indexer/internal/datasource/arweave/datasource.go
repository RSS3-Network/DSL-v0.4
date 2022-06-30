package arweave

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/arweave"
	graphqlx "github.com/naturalselectionlabs/pregod/common/arweave/graphql"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"go.opentelemetry.io/otel"
)

const (
	Source = "arweave"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	arweaveClient *arweave.Client
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkArweave,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("arweave_datasource")
	_, trace := tracer.Start(ctx, "arweave_datasource:Handle")

	defer opentelemetry.Log(trace, message, transactions, err)

	address := common.NewMixedcaseAddress(common.HexToAddress(message.Address))

	transactions = make([]model.Transaction, 0)

	var query struct {
		TransactionConnection graphqlx.TransactionConnection `graphql:"transactions(owners: $owners, tags: $tags)"`
	}

	variable := d.buildGetTransactionsVariable(ctx, address.Address().String())

	if err := d.arweaveClient.GetTransactions(ctx, &query, variable); err != nil {
		return nil, err
	}

	for _, edge := range query.TransactionConnection.Edges {
		sourceData, err := json.Marshal(edge)
		if err != nil {
			return nil, err
		}

		timestamp := time.Unix(int64(edge.Node.Block.Timestamp), 0)

		// Mirror's transactions don't have a recipient address
		addressTo := ""

		transactions = append(transactions, model.Transaction{
			BlockNumber: int64(edge.Node.Block.Height),
			Timestamp:   timestamp,
			Hash:        strings.ToLower(edge.Node.ID.(string)),
			AddressFrom: strings.ToLower(string(edge.Node.Owner.Address)),
			AddressTo:   addressTo,
			Platform:    message.Network,
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
			Transfers: []model.Transfer{
				// This is a virtual transfer
				{
					TransactionHash: strings.ToLower(edge.Node.ID.(string)),
					Timestamp:       timestamp,
					Index:           protocol.IndexVirtual,
					AddressFrom:     strings.ToLower(string(edge.Node.Owner.Address)),
					AddressTo:       addressTo,
					Metadata:        metadata.Default,
					Platform:        message.Network,
					Network:         message.Network,
					Source:          Source,
					SourceData:      sourceData,
				},
			},
		})
	}

	return transactions, nil
}

func (d *Datasource) buildGetTransactionsVariable(ctx context.Context, address string) arweave.GetTransactionsVariable {
	return arweave.GetTransactionsVariable{
		Owners: []graphql.String{
			arweave.AddressMirror,
		},
		Tags: []graphqlx.TagFilter{
			{
				Name: "App-Name",
				Values: []graphql.String{
					"MirrorXYZ",
				},
				TagOperator: graphqlx.TagOperatorEQ,
			},
			{
				Name: "Contributor",
				Values: []graphql.String{
					graphql.String(address),
				},
				TagOperator: graphqlx.TagOperatorEQ,
			},
		},
	}
}

func New() datasource.Datasource {
	return &Datasource{
		arweaveClient: arweave.NewClient(),
	}
}
