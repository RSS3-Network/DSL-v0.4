package friendtech

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"
)

const (
	NodeEndpoint = "https://prod-api.kosetto.com"
)

type Client struct {
	restyClient *resty.Client
}

func (c *Client) GetUserMeta(ctx context.Context, address string) (*UserResponse, error) {
	var result UserResponse

	_, err := c.restyClient.R().SetPathParams(
		map[string]string{
			"address": address,
		},
	).SetContext(ctx).SetResult(&result).Get(NodeEndpoint + "/users/{address}")
	if err != nil {
		return nil, err
	}

	if result.TwitterUserID == "" && result.Address == "" {
		return nil, nil
	}

	return &result, nil
}

func (c *Client) GetUserMetaByID(ctx context.Context, id int64) (*UserResponse, error) {
	var result UserResponse

	r, err := c.restyClient.R().SetPathParams(
		map[string]string{
			"id": strconv.FormatInt(id, 10),
		},
	).SetContext(ctx).SetResult(&result).Get(NodeEndpoint + "/users/by-id/{id}")
	if err != nil {
		return nil, err
	}

	if r.RawResponse.StatusCode != 200 {
		var message MessageResponse
		_ = json.Unmarshal(r.Body(), &message)
		if message.Message == "Address/User not found." {
			result.ID = -1
		}
	}

	return &result, nil
}

func NewClient() *Client {
	return &Client{
		restyClient: resty.New().SetRetryCount(10).SetTimeout(time.Second * 3).SetRetryAfter(func(c *resty.Client, r *resty.Response) (time.Duration, error) {
			loggerx.Global().Warn("retry request friend.tech public api", zap.String("url", r.String()))
			return time.Millisecond * 500, nil
		}),
	}
}
