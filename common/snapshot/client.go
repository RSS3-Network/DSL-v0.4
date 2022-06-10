package snapshot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shurcooL/graphql"
)

const (
	EndpointScheme = "https"
	EndpointHost   = "hub.snapshot.org"
	EndpointPath   = "/graphql"
)

type Description string

const (
	DescriptionAsc  Description = "asc"
	DescriptionDesc Description = "desc"
)

type Client struct {
	httpClient    *http.Client
	graphqlClient *graphql.Client
}

type GetMultipleSpacesVariable struct {
	First          graphql.Int
	Skip           graphql.Int
	OrderBy        graphql.String
	OrderDirection Description
}

func (c *Client) GetMultipleSpaces(ctx context.Context, variable GetMultipleSpacesVariable) error {
	variableMap := map[string]interface{}{}

	if variable.First > 0 {
		variableMap["first"] = variable.First
	} else {
		return fmt.Errorf("variable 'First' must be greater than 0")
	}

	if variable.Skip >= 0 {
		variableMap["first"] = variable.First
	} else {
		return fmt.Errorf("variable 'Skip' must be greater than 0")
	}

	if variable.OrderBy != "" {
		variableMap["first"] = variable.First
	} else {
		return fmt.Errorf("variable 'OrderBy' must not be nil")
	}

	variableMap["OrderDirection"] = variable.OrderDirection

	if err := c.graphqlClient.Query(ctx, query, variableMap); err != nil {
		return err
	}

	return nil
}

type GetMultipleProposalsVariable struct {
	First          graphql.Int
	Skip           graphql.Int
	WhereSpaceIn   []graphql.String
	WhereState     graphql.String
	orderBy        graphql.String
	orderDirection Description
}

func (c *Client) GetMultipleProposals(name string) {

}

type GetMultipleVotesVariable struct {
	First           graphql.Int
	Skip            graphql.Int
	WhereProposalIn []graphql.String
	orderBy         graphql.String
	orderDirection  Description
}

func (c *Client) GetMultipleVotes() {

}
