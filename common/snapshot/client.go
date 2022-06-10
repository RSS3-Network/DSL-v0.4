package snapshot

import (
	"net/http"

	"github.com/shurcooL/graphql"
)

const (
	EndpointScheme = "https"
	EndpointHost   = "hub.snapshot.org"
	EndpointPath   = "/graphql"
)

type Client struct {
	httpClient    *http.Client
	graphqlClient *graphql.Client
}

func (c *Client) GetMultipleSpaces() {

}

func (c *Client) GetMultipleProposals(name string) {

}

func (c *Client) GetMultipleVotes() {

}
