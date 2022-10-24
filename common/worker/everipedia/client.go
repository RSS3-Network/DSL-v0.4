package everipedia

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Khan/genqlient/graphql"
	"github.com/ethereum/go-ethereum/common"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/everipedia/graphql"
)

const (
	Scheme   = "https"
	Endpoint = "graph.everipedia.org"
	Path     = "graphql"
	Limit    = 30
)

type Client struct {
	GClient *graphql.Client
}

func (c *Client) GetUserActivityList(ctx context.Context, address string) ([]graphqlx.GetUserActivitiesActivitiesByUserActivity, error) {
	var activityList []graphqlx.GetUserActivitiesActivitiesByUserActivity
	userId := common.HexToAddress(address).String()
	offset := 0
	index := 0

	for {
		resp, err := graphqlx.GetUserActivities(ctx, *c.GClient, userId, Limit, offset)
		if err != nil {
			return nil, err
		}
		activityList = append(activityList, resp.ActivitiesByUser...)
		if len(resp.ActivitiesByUser) < Limit {
			break
		}
		index++
		offset = Limit * index
	}

	return activityList, nil
}

func NewClient() *Client {
	requestURL := &url.URL{
		Scheme: Scheme,
		Host:   Endpoint,
		Path:   Path,
	}
	client := graphql.NewClient(requestURL.String(), http.DefaultClient)

	return &Client{
		GClient: &client,
	}
}
