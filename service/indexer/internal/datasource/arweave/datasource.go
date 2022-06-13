package arweave

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/arweave"
	graphqlx "github.com/naturalselectionlabs/pregod/common/arweave/graphql"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/shopspring/decimal"
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

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	address := common.NewMixedcaseAddress(common.HexToAddress(message.Address))

	var query struct {
		TransactionConnection graphqlx.TransactionConnection `graphql:"transactions(owners: $owners, tags: $tags)"`
	}

	variable := arweave.GetTransactionsVariable{
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
					graphql.String(address.Address().String()),
				},
				TagOperator: graphqlx.TagOperatorEQ,
			},
		},
	}

	if err := d.arweaveClient.GetTransactions(ctx, &query, variable); err != nil {
		return nil, err
	}

	for _, edge := range query.TransactionConnection.Edges {
		sourceData, err := json.Marshal(edge)
		if err != nil {
			return nil, err
		}

		internalTransfers = append(internalTransfers, model.Transfer{
			TransactionHash:     edge.Node.ID.(string),
			Timestamp:           time.Unix(int64(edge.Node.Block.Timestamp), 0),
			TransactionLogIndex: decimal.Zero,
			AddressFrom:         string(edge.Node.Owner.Address),
			AddressTo:           "",
			Metadata:            metadata.Default,
			Network:             protocol.NetworkArweave,
			Source:              Source,
			SourceData:          sourceData,
		})
	}

	return internalTransfers, nil
}

func New() datasource.Datasource {
	return &Datasource{
		arweaveClient: arweave.NewClient(),
	}
}
