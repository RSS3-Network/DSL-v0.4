package thegraph_test

import (
	"context"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/thegraph"
	graphqlx "github.com/naturalselectionlabs/pregod/common/thegraph/graphql"
	"github.com/shurcooL/graphql"
	"testing"
)

var (
	client *thegraph.Client
)

func init() {
	client = thegraph.NewClient()
}

func TestGetUniswapV3Pools(t *testing.T) {
	var query struct {
		Pairs []graphqlx.Pair `graphql:"pools(first: 1000, skip: $skip, orderBy: volumeUSD, orderDirection: desc)"`
	}

	firstQuery := true

	// a subgraph only returns 5000 records, so i < 6
	// a query returns 1000 records, so len(query.Pairs) == 1000
	// the first query needs to be sent, so firstQuery = true
	for i := 0; (i < 6 && len(query.Pairs) == 1000) || firstQuery; i++ {
		firstQuery = false
		variableMap := map[string]interface{}{
			"skip": graphql.NewInt(graphql.Int(i * 1000)),
		}

		err := client.GetPools(context.Background(), protocol.UniSwapV3, &query, variableMap)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(len(query.Pairs))
	}

}

func TestGetUniswapV2Pools(t *testing.T) {
	var query struct {
		Pairs []graphqlx.Pair `graphql:"pairs(first: 1000, skip: $skip, orderBy: volumeUSD, orderDirection: desc)"`
	}

	firstQuery := true

	// a subgraph only returns 5000 records, so i < 6
	// a query returns 1000 records, so len(query.Pairs) == 1000
	// the first query needs to be sent, so firstQuery = true
	for i := 0; (i < 6 && len(query.Pairs) == 1000) || firstQuery; i++ {
		firstQuery = false
		variableMap := map[string]interface{}{
			"skip": graphql.NewInt(graphql.Int(i * 1000)),
		}

		err := client.GetPools(context.Background(), protocol.UniSwapV2, &query, variableMap)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(len(query.Pairs))
	}
}

func TestGetSushiSwapPools(t *testing.T) {
	var query struct {
		Pairs []graphqlx.Pair `graphql:"pairs(first: 1000, skip: $skip, orderBy: volumeUSD, orderDirection: desc)"`
	}

	firstQuery := true

	// a subgraph only returns 5000 records, so i < 6
	// a query returns 1000 records, so len(query.Pairs) == 1000
	// the first query needs to be sent, so firstQuery = true
	for i := 0; (i < 6 && len(query.Pairs) == 1000) || firstQuery; i++ {
		firstQuery = false
		variableMap := map[string]interface{}{
			"skip": graphql.NewInt(graphql.Int(i * 1000)),
		}

		err := client.GetPools(context.Background(), protocol.SushiSwap, &query, variableMap)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(len(query.Pairs))
	}
}
