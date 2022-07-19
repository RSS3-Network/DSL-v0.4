package lens

import (
	"context"
	"net/http"
	"net/url"

	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/lens/graphql"

	"github.com/hasura/go-graphql-client"
)

const (
	EndpointScheme = "https"
	EndpointHost   = "api.lens.dev"
	EndpointPath   = "/graphql"
	// EndpointLimit : the maximum number of items that can be returned in a single query
	EndpointLimit = 50
)

var PublicationType = []PublicationTypes{"POST", "COMMENT", "MIRROR"}

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

func (c *Client) GetProfiles(ctx context.Context, options *Options) ([]graphqlx.Profile, error) {
	var query struct {
		Profiles struct {
			Items []graphqlx.Profile `graphql:"items"`
		} `graphql:"profiles(request: $request )"`
	}

	variableMap := map[string]interface{}{
		"request": ProfileQueryRequest{
			OwnedBy: []EthereumAddress{EthereumAddress(options.Address)},
		},
	}

	if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
		return nil, err
	}

	return query.Profiles.Items, nil
}

// types for GetPublications and GetPublicationCount
type (
	PublicationTypes         string
	PublicationsQueryRequest struct {
		ProfileId        graphql.String     `json:"profileId"`
		PublicationTypes []PublicationTypes `json:"publicationTypes"`
		Cursor           graphql.String     `json:"cursor"`
		Limit            graphql.Int        `json:"limit"`
	}
)

func (c *Client) GetPublications(ctx context.Context, options *Options) ([]graphqlx.Publication, error) {
	variable := PublicationsQueryRequest{
		ProfileId:        options.Profile,
		PublicationTypes: PublicationType,
		Cursor:           options.Cursor,
		Limit:            graphql.Int(EndpointLimit),
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

			var query struct {
				Publications struct {
					Items    []graphqlx.Item
					PageInfo graphqlx.PageInfo
				} `graphql:"publications(request: $request)"`
			}

			if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
				return nil, err
			}

			for _, item := range query.Publications.Items {
				publication := graphqlx.Publication{
					Type:       item.Post.Type,
					ID:         item.Post.ID,
					RelatedURL: item.Post.RelatedURL,
					Platform:   item.Post.Platform,
					Metadata:   item.Post.Metadata,
					CreatedAt:  item.Post.CreatedAt,
				}
				// assign the right target to the publication
				if err := assignTarget(&item, &publication); err != nil {
					return nil, err
				}

				result = append(result, publication)
			}
			if query.Publications.PageInfo.Next != "" {
				options.Cursor = query.Publications.PageInfo.Next
			}
		}
	}

	return result, nil
}

func assignTarget(item *graphqlx.Item, publication *graphqlx.Publication) error {
	var target graphqlx.Post
	var err error

	// a target can either be a Post or a Comment
	switch publication.Type {
	case "Comment":
		if item.Comment.CommentOn.Post.Type == "Post" {
			target.ID = item.Comment.CommentOn.Post.ID
			target.Type = "Post"
			target.Metadata = item.Comment.CommentOn.Post.Metadata
			target.CreatedAt = item.Comment.CommentOn.Post.CreatedAt
			if err != nil {
				return err
			}
		} else {
			target.ID = item.Comment.CommentOn.Comment.ID
			target.Type = "Comment"
			target.Metadata = item.Comment.CommentOn.Comment.Metadata
			target.CreatedAt = item.Comment.CommentOn.Comment.CreatedAt
			if err != nil {
				return err
			}
		}

	case "Mirror":
		if item.Mirror.MirrorOf.Post.Type == "Post" {
			target.ID = item.Mirror.MirrorOf.Post.ID
			target.Type = "Post"
			target.Metadata = item.Mirror.MirrorOf.Post.Metadata
			target.CreatedAt = item.Mirror.MirrorOf.Post.CreatedAt
			if err != nil {
				return err
			}
		} else {
			target.ID = item.Mirror.MirrorOf.Comment.ID
			target.Type = "Post"
			target.Metadata = item.Mirror.MirrorOf.Comment.Metadata
			target.CreatedAt = item.Mirror.MirrorOf.Comment.CreatedAt
			if err != nil {
				return err
			}
		}
	}

	publication.Target = target
	return nil
}

func (c *Client) GetPublicationPageInfo(ctx context.Context, options *Options) error {
	var query struct {
		Publications struct {
			PageInfo graphqlx.PageInfo
		} `graphql:"publications(request: $request)"`
	}

	variable := PublicationsQueryRequest{
		ProfileId:        options.Profile,
		PublicationTypes: PublicationType,
		Cursor:           graphql.String("{}"),
		Limit:            graphql.Int(EndpointLimit),
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
	profiles, err := c.GetProfiles(ctx, options)
	if err != nil {
		return nil, err
	}

	result := make([]graphqlx.Publication, 0)

	for _, profile := range profiles {
		options.Profile = profile.ID
		publications, err := c.GetPublications(ctx, options)
		if err != nil {
			return nil, err
		}

		result = append(result, publications...)

	}

	return result, nil
}

type (
	FollowingRequest struct {
		Address EthereumAddress `json:"address"`
		Cursor  graphql.String  `json:"cursor"`
		Limit   graphql.Int     `json:"limit"`
	}
)

func (c *Client) GetFollowings(ctx context.Context, options *Options) ([]graphqlx.Profile, error) {
	variable := FollowingRequest{
		Address: EthereumAddress(options.Address),
		Cursor:  options.Cursor,
		Limit:   graphql.Int(EndpointLimit),
	}
	options.TotalCount = 0

	result := make([]graphqlx.Profile, 0)

	for i := 0; i <= int(options.TotalCount); i += EndpointLimit {
		variable.Cursor = options.Cursor
		variableMap := map[string]interface{}{
			"request": variable,
		}

		var query struct {
			Following struct {
				Items []struct {
					Profile graphqlx.Profile `graphql:"profile"`
				} `graphql:"items"`
				PageInfo graphqlx.PageInfo
			} `graphql:"following(request: $request)"`
		}

		if err := c.graphqlClient.Query(ctx, &query, variableMap); err != nil {
			return nil, err
		}

		for _, item := range query.Following.Items {
			result = append(result, item.Profile)
		}

		if len(query.Following.Items) > 0 {
			options.TotalCount = query.Following.PageInfo.TotalCount
			options.Cursor = query.Following.PageInfo.Next
		}
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
