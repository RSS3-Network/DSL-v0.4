package snapshot

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hasura/go-graphql-client"
	graphqlx "github.com/naturalselectionlabs/pregod/common/snapshot/graphql"
)

const (
	EndpointScheme = "https"
	EndpointHost   = "hub.snapshot.org"
	// EndpointScheme = "http"
	// EndpointHost   = "127.0.0.1:8080"
	EndpointPath = "/graphql"
)

type OrderDirection graphql.String

const (
	OrderDirectionAsc  OrderDirection = "asc"
	OrderDirectionDesc OrderDirection = "desc"
)

type Client struct {
	snapShotGraphqlUrl url.URL
	httpClient         *http.Client
	graphqlClient      *graphql.Client
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

	if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
		return nil, err
	}

	return query.Spaces, nil
}

type GetMultipleProposalsVariable struct {
	First          graphql.Int            `json:"first"`
	Skip           graphql.Int            `json:"skip"`
	Where          graphqlx.ProposalWhere `json:"where"`
	OrderBy        graphql.String         `json:"orderBy"`
	OrderDirection OrderDirection         `json:"orderDirection"`
}

func (c *Client) GetMultipleProposals(ctx context.Context, variable GetMultipleProposalsVariable) ([]graphqlx.Proposal, error) {
	var query struct {
		Proposals []graphqlx.Proposal `graphql:"proposals(first: $first, skip: $skip, where:$where, orderBy: $orderBy, orderDirection: $orderDirection)"`
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

	variableMap["where"] = variable.Where
	variableMap["orderDirection"] = variable.OrderDirection

	fmt.Printf("variableMap: %v\n", variableMap)

	if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
		return nil, err
	}

	return query.Proposals, nil
}

type GetMultipleVotesVariable struct {
	First          graphql.Int
	Skip           graphql.Int
	Where          graphqlx.VoteWhere
	OrderBy        graphql.String
	OrderDirection OrderDirection
}

func (c *Client) GetMultipleVotes(ctx context.Context, variable GetMultipleVotesVariable) ([]graphqlx.Vote, error) {
	var query struct {
		Votes []graphqlx.Vote `graphql:"votes(first: $first, skip: $skip, where:$where, orderBy: $orderBy, orderDirection: $orderDirection)"`
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

	variableMap["where"] = variable.Where
	variableMap["orderDirection"] = variable.OrderDirection

	if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
		return nil, err
	}

	return query.Votes, nil
}

func NewClient() *Client {
	client := &Client{
		httpClient: http.DefaultClient,
	}

	client.snapShotGraphqlUrl = url.URL{
		Scheme: EndpointScheme,
		Host:   EndpointHost,
		Path:   EndpointPath,
	}

	client.graphqlClient = graphql.NewClient(client.snapShotGraphqlUrl.String(), client.httpClient)

	return client
}
