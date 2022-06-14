package arweave

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/arweave"
	graphqlx "github.com/naturalselectionlabs/pregod/common/arweave/graphql"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/shopspring/decimal"
	"github.com/shurcooL/graphql"
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

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	address := common.NewMixedcaseAddress(common.HexToAddress(message.Address))

	transactions := make([]model.Transaction, 0)

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
			BlockNumber: decimal.NewFromInt32(int32(edge.Node.Block.Height)),
			Timestamp:   timestamp,
			Hash:        edge.Node.ID.(string),
			AddressFrom: string(edge.Node.Owner.Address),
			AddressTo:   addressTo,
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
			Transfers: []model.Transfer{
				// This is a virtual transfer
				{
					TransactionHash:     edge.Node.ID.(string),
					Timestamp:           timestamp,
					TransactionLogIndex: protocol.LogIndexVirtual,
					AddressFrom:         string(edge.Node.Owner.Address),
					AddressTo:           addressTo,
					Metadata:            metadata.Default,
					Network:             protocol.NetworkArweave,
					Source:              Source,
					SourceData:          sourceData,
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
