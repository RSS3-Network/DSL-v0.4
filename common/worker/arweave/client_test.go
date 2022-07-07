package arweave_test

import (
	"context"
	"github.com/naturalselectionlabs/pregod/common/worker/arweave"
	"github.com/naturalselectionlabs/pregod/common/worker/arweave/graphql"
	"testing"

	"github.com/hasura/go-graphql-client"
)

var (
	client *arweave.Client
)

func init() {
	client = arweave.NewClient()
}

func TestGetInfo(t *testing.T) {
	info, _, err := client.GetInfo(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(info)
}

func TestGetBlockByID(t *testing.T) {
	block, _, err := client.GetBlockByHeight(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(block)
}

func TestGetTransactions(t *testing.T) {
	var query struct {
		TransactionConnection graphqlx.TransactionConnection `graphql:"transactions(block: $block)"`
	}

	variable := arweave.GetTransactionsVariable{
		Block: &graphqlx.BlockFilter{
			Min: graphql.Int(0),
			Max: graphql.Int(100),
		},
	}

	if err := client.GetTransactions(context.Background(), &query, variable); err != nil {
		t.Fatal(err)
	}

	t.Log(query.TransactionConnection)
}
