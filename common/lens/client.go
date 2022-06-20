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
	// EndpointLimit : the maximum number of items that can be returned in a single query
	EndpointLimit = 25
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

// types for GetPublications and GetPublicationCount
type (
	PublicationTypes         string
	PublicationsQueryRequest struct {
		ProfileId        graphql.String     `json:"profileId"`
		PublicationTypes []PublicationTypes `json:"publicationTypes"`
		Cursor           graphql.String     `json:"cursor"`
	}
)

func (c *Client) GetPublications(ctx context.Context, options *Options) ([]graphqlx.Publication, error) {
	var query struct {
		Publications struct {
			Items []struct {
				Post    graphqlx.Post        `graphql:"... on Post"`
				Comment graphqlx.Publication `graphql:"... on Comment"`
			}
			PageInfo graphqlx.PageInfo
		} `graphql:"publications(request: $request)"`
	}

	variable := PublicationsQueryRequest{
		ProfileId:        options.Profile,
		PublicationTypes: []PublicationTypes{"POST", "COMMENT"},
		Cursor:           options.Cursor,
	}

	if err := c.GetPublicationPageInfo(ctx, options); err != nil {
		return nil, err
	}

	result := make([]graphqlx.Publication, 0)

	if options.TotalCount > 0 {
		for i := 0; i < int(options.TotalCount); i += EndpointLimit {
			variable.Cursor = options.Cursor
			variableMap := map[string]interface{}{
				"request": variable,
			}
			if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
				return nil, err
			}

			for _, item := range query.Publications.Items {
				// both item.Post and item.Comment contain the same data
				// here we only want to append one of them
				result = append(result, item.Comment)
			}
			if query.Publications.PageInfo.Next != "" {
				options.Cursor = query.Publications.PageInfo.Next
			}
		}
	}

	return result, nil
}

func (c *Client) GetPublicationPageInfo(ctx context.Context, options *Options) error {
	var query struct {
		Publications struct {
			PageInfo graphqlx.PageInfo
		} `graphql:"publications(request: $request)"`
	}

	variable := PublicationsQueryRequest{
		ProfileId:        options.Profile,
		PublicationTypes: []PublicationTypes{"POST", "COMMENT"},
		Cursor:           graphql.String("{}"),
	}

	variableMap := map[string]interface{}{
		"request": variable,
	}

	if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
		return err
	}

	options.TotalCount = query.Publications.PageInfo.TotalCount

	return nil
}

func (c *Client) GetAllPublicationsByAddress(ctx context.Context, options *Options) ([]graphqlx.Publication, error) {
	profiles, err := c.GetProfiles(ctx, string(options.Address))
	if err != nil {
		return nil, err
	}

	result := make([]graphqlx.Publication, 0)

	for _, profile := range profiles {
		options.Profile = graphql.String(profile)
		publications, err := c.GetPublications(ctx, options)
		if err != nil {
			return nil, err
		}

		result = append(result, publications...)

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
