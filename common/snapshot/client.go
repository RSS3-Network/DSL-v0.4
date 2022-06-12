package snapshot

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	graphqlx "github.com/naturalselectionlabs/pregod/common/snapshot/graphql"
	"github.com/shurcooL/graphql"
)

const (
	// EndpointScheme = "https"
	// EndpointHost   = "hub.snapshot.org"
	EndpointScheme = "http"
	EndpointHost   = "127.0.0.1:8080"
	EndpointPath   = "/graphql"
)

type OrderDirection graphql.String

const (
	OrderDirectionAsc  OrderDirection = "asc"
	OrderDirectionDesc OrderDirection = "desc"
)

type Client struct {
	httpClient    *http.Client
	graphqlClient *graphql.Client
}

type GetMultipleSpacesVariable struct {
	First          graphql.Int    `json:"first"`
	Skip           graphql.Int    `json:"skip"`
	OrderBy        graphql.String `json:"orderBy"`
	OrderDirection OrderDirection `json:"orderDirection"`
}

func (c *Client) GetMultipleSpaces(ctx context.Context, variable GetMultipleSpacesVariable) ([]graphqlx.Space, error) {
	var query struct {
		Spaces []graphqlx.Space `graphql:"spaces(first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDirection)"`
	}

	variableMap := map[string]interface{}{}

	if variable.First > 0 {
		variableMap["first"] = variable.First
	} else {
		return nil, fmt.Errorf("variable 'First' must be greater than 0")
	}

	if variable.Skip >= 0 {
		variableMap["skip"] = variable.Skip
	} else {
		return nil, fmt.Errorf("variable 'Skip' must be greater than 0")
	}

	if variable.OrderBy != "" {
		variableMap["orderBy"] = variable.OrderBy
	} else {
		return nil, fmt.Errorf("variable 'OrderBy' must not be nil")
	}

	variableMap["orderDirection"] = variable.OrderDirection

	if err := c.graphqlClient.Query(ctx, query, variableMap); err != nil {
		return nil, err
	}

	return query.Spaces, nil
}

type GetMultipleProposalsVariable struct {
	First          graphql.Int
	Skip           graphql.Int
	WhereSpaceIn   []graphql.String
	WhereState     graphql.String
	orderBy        graphql.String
	orderDirection OrderDirection
}

func (c *Client) GetMultipleProposals(name string, variable GetMultipleProposalsVariable) error {
	return nil
}

type GetMultipleVotesVariable struct {
	First           graphql.Int
	Skip            graphql.Int
	WhereProposalIn []graphql.String
	orderBy         graphql.String
	orderDirection  OrderDirection
}

func (c *Client) GetMultipleVotes(variable GetMultipleVotesVariable) error {
	return nil
}

func NewClient() *Client {
	client := &Client{
		httpClient: http.DefaultClient,
	}

	endpointURL := url.URL{
		Scheme: EndpointScheme,
		Host:   EndpointHost,
		Path:   EndpointPath,
	}

	client.graphqlClient = graphql.NewClient(endpointURL.String(), client.httpClient)

	return client
}
