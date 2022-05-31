package thegraph

import (
	"context"
	"net/http"
	"net/url"

	"github.com/shurcooL/graphql"
)

const (
	EndpointScheme = "https"
	EndpointHost   = "api.thegraph.com"
)

type Client struct {
	httpClient    *http.Client
	graphqlClient *graphql.Client
}

func (c *Client) GetPools(ctx context.Context, subgraph string, query interface{}, variableMap map[string]interface{}) error {
	endpointURL := url.URL{
		Scheme: EndpointScheme,
		Host:   EndpointHost,
		Path:   subgraph,
	}

	c.graphqlClient = graphql.NewClient(endpointURL.String(), c.httpClient)

	if err := c.graphqlClient.Query(ctx, query, variableMap); err != nil {
		return err
	}

	return nil
}

func NewClient() *Client {
	client := &Client{
		httpClient: http.DefaultClient,
	}

	return client
}
