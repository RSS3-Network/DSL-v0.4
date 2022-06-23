package dexpools

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/hasura/go-graphql-client"
	graphqlx "github.com/naturalselectionlabs/pregod/common/dexpools/graphql"
)

type Client struct {
	httpClient    *http.Client
	graphqlClient *graphql.Client
}

type GetQueryFun func() interface{}

// GetSwapPools returns all pools from a DEX
func (c *Client) GetSwapPools(ctx context.Context, swap SwapPool) ([]graphqlx.Pair, error) {
	// set the default limit to 6000
	if swap.Limit == 0 {
		swap.Limit = 6000
	}

	if swap.NonSubgraph {
		response, err := c.httpClient.Get(swap.Endpoint)
		if err != nil {
			return nil, err
		}

		defer func() {
			_ = response.Body.Close()
		}()

		// case by case handling of the response
		if strings.HasPrefix(swap.Name, "1inch") {
			type OneInch struct {
				Id     string `json:"pair"`
				Token0 struct {
					Name   string `json:"address"`
					Symbol string `json:"symbol"`
				} `json:"token0"`
				Token1 struct {
					Name   string `json:"address"`
					Symbol string `json:"symbol"`
				} `json:"token1"`
			}

			res := make([]OneInch, 0)

			if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
				return nil, err
			}

			result := make([]graphqlx.Pair, len(res))

			for i, pair := range res {
				result[i] = graphqlx.Pair{
					ID:     graphql.String(pair.Id),
					Token0: graphqlx.Token{Name: graphql.String(pair.Token0.Symbol), Symbol: graphql.String(pair.Token0.Symbol)},
					Token1: graphqlx.Token{Name: graphql.String(pair.Token1.Symbol), Symbol: graphql.String(pair.Token1.Symbol)},
				}
			}

			return result, nil
		}
	} else {
		c.graphqlClient = graphql.NewClient(swap.Endpoint, c.httpClient)
		switch swap.Protocol {
		case UniSwapV2:
			// nolint:gocritic // cannot dynamically create a new struct with different graphql tags
			if swap.OrderByVolumeUSD {
				getQueryFun := func() interface{} {
					var query struct {
						Pairs []graphqlx.Pair `graphql:"pairs(first: 1000, skip: $skip, orderBy: volumeUSD, orderDirection: desc)"`
					}

					return &query
				}

				logrus.Infof("UniSwapV21")
				return c.GetGraphQLResult(ctx, getQueryFun, swap.Limit)
			} else {
				getQueryFun := func() interface{} {
					var query struct {
						Pairs []graphqlx.Pair `graphql:"pairs(first: 1000, skip: $skip)"`
					}

					return &query
				}

				logrus.Infof("UniSwapV22")
				return c.GetGraphQLResult(ctx, getQueryFun, swap.Limit)
			}

		case UniSwapV3:
			// nolint:gocritic // cannot dynamically create a new struct with different graphql tags
			if swap.OrderByVolumeUSD {
				getQueryFun := func() interface{} {
					var query struct {
						Pairs []graphqlx.Pair `graphql:"pools(first: 1000, skip: $skip, orderBy: volumeUSD, orderDirection: desc)"`
					}

					return &query
				}

				return c.GetGraphQLResult(ctx, getQueryFun, swap.Limit)
			} else {
				getQueryFun := func() interface{} {
					var query struct {
						Pairs []graphqlx.Pair `graphql:"pools(first: 1000, skip: $skip)"`
					}

					return &query
				}

				return c.GetGraphQLResult(ctx, getQueryFun, swap.Limit)
			}

		}

	}

	return nil, nil
}

// GetGraphQLResult executes a GraphQL query and returns the result
func (c *Client) GetGraphQLResult(ctx context.Context, queryFun GetQueryFun, limit int) ([]graphqlx.Pair, error) {
	result := make([]graphqlx.Pair, 0)
	firstQuery := true

	// a subgraph only returns `limit` records, so i < limit/1000
	// a query returns 1000 records, so len(result)%1000 == 0
	// the first query needs to be sent anyway, so firstQuery = true
	for i := 0; (i < limit/1000 && len(result)%1000 == 0) || firstQuery; i++ {
		variableMap := map[string]interface{}{
			"skip": graphql.NewInt(graphql.Int(i * 1000)),
		}

		query := queryFun()

		firstQuery = false

		if err := c.graphqlClient.Query(ctx, query, variableMap); err != nil {
			return nil, err
		}

		// extract the result from the query without knowing the interface struct
		pairs := reflect.ValueOf(query).Elem().FieldByName("Pairs").Interface().([]graphqlx.Pair)

		result = append(result, pairs...)
	}

	return result, nil
}

func NewClient() *Client {
	client := &Client{
		httpClient: http.DefaultClient,
	}

	return client
}
