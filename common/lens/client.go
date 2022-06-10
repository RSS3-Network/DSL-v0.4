package lens

import (
	"context"
	"net/http"
	"net/url"

	graphqlx "github.com/naturalselectionlabs/pregod/common/lens/graphql"
	"github.com/shurcooL/graphql"
)

const (
	EndpointScheme = "https"
	EndpointHost   = "api.lens.dev"
	EndpointPath   = "/graphql"
)

type Client struct {
	graphqlClient *graphql.Client
}

// types for GetProfiles
type (
	EthereumAddress     string
	ProfileQueryRequest struct {
		OwnedBy []EthereumAddress `json:"ownedBy"`
	}
)

func (c *Client) GetProfiles(ctx context.Context, address string) ([]string, error) {
	var query struct {
		Profiles struct {
			Items []struct {
				ID graphql.String `json:"id"`
			} `graphql:"items"`
		} `graphql:"profiles(request: $request )"`
	}

	variable := EthereumAddress(address)

	variableMap := map[string]interface{}{
		"request": ProfileQueryRequest{
			OwnedBy: []EthereumAddress{variable},
		},
	}

	if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
		return nil, err
	}

	result := make([]string, len(query.Profiles.Items))

	for i, item := range query.Profiles.Items {
		result[i] = string(item.ID)
	}

	return result, nil
}

// types for GetPublications
type (
	PublicationTypes         string
	PublicationsQueryRequest struct {
		ProfileId        graphql.String     `json:"profileId"`
		PublicationTypes []PublicationTypes `json:"publicationTypes"`
	}
)

func (c *Client) GetPublications(ctx context.Context, profile string) ([]graphqlx.Publication, error) {
	var query struct {
		Publications struct {
			Items []struct {
				Post    graphqlx.Publication `graphql:"... on Post"`
				Comment graphqlx.Publication `graphql:"... on Comment"`
			} `graphql:"items"`
		} `graphql:"publications(request: $request)"`

		// `graphql:"publications(request: { profileId: \"0x0d\", publicationTypes: [POST,COMMENT] })"`
	}

	variable := PublicationsQueryRequest{
		ProfileId:        graphql.String(profile),
		PublicationTypes: []PublicationTypes{"POST", "COMMENT"},
	}

	variableMap := map[string]interface{}{
		"request": variable,
	}
	if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
		return nil, err
	}

	result := make([]graphqlx.Publication, len(query.Publications.Items))

	for i, item := range query.Publications.Items {
		// both item.Post and item.Comment contain the same data
		// here we only want to append one of them
		result[i] = item.Post
	}

	return result, nil
}

func NewClient() *Client {
	endpointURL := url.URL{
		Scheme: EndpointScheme,
		Host:   EndpointHost,
		Path:   EndpointPath,
	}

	client := &Client{
		graphqlClient: graphql.NewClient(endpointURL.String(), http.DefaultClient),
	}

	return client
}
