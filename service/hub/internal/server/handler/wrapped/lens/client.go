package lens

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Khan/genqlient/graphql"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
)

const (
	Scheme   = "https"
	Endpoint = "api.lens.dev"
)

type Client struct {
	GClient *graphql.Client
}

type Result struct {
	Follower  int `json:"follower"`
	Following int `json:"following"`
}

func (c *Client) GetFollowStat(ctx context.Context, result *model.SocialResult, address string) error {
	profile, err := GetDefaultProfile(ctx, *c.GClient, address)

	if err != nil {
		return err
	}

	profileId := profile.DefaultProfile.Id

	followers, err := GetFollowers(ctx, *c.GClient, profileId)
	if err != nil {
		return err
	}

	following, err := GetFollowing(ctx, *c.GClient, address)
	if err != nil {
		return err
	}

	result.Follower += int64(followers.Followers.PageInfo.TotalCount)
	result.Following += int64(following.Following.PageInfo.TotalCount)

	return nil
}

func NewClient() *Client {
	requestURL := &url.URL{
		Scheme: Scheme,
		Host:   Endpoint,
	}
	client := graphql.NewClient(requestURL.String(), http.DefaultClient)

	return &Client{
		GClient: &client,
	}
}
