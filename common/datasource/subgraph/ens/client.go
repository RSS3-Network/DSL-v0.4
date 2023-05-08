package ens

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Khan/genqlient/graphql"
	enssubgraph "github.com/naturalselectionlabs/pregod/common/datasource/subgraph/ens/graphql"
)

const (
	Endpoint = "https://api.thegraph.com/subgraphs/name/ensdomains/ens"
)

var _ Client = (*client)(nil)

type Client interface {
	GetEnsName(ctx context.Context, nameHash string) (*enssubgraph.GetEnsNameByNameHashResponse, error)
}

type client struct {
	graphqlClient graphql.Client
}

func (c *client) GetEnsName(ctx context.Context, nameHash string) (*enssubgraph.GetEnsNameByNameHashResponse, error) {
	response, err := enssubgraph.GetEnsNameByNameHash(ctx, c.graphqlClient, nameHash)
	if err != nil {
		return nil, fmt.Errorf("get ens name %s: %w", nameHash, err)
	}

	return response, nil
}

func New() Client {
	endpointURL, _ := url.Parse(Endpoint)

	var instance client

	instance.graphqlClient = graphql.NewClient(endpointURL.String(), http.DefaultClient)

	return &instance
}
