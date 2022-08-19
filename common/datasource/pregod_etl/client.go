package pregod_etl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	http_utils "github.com/naturalselectionlabs/pregod/common/utils/http"
)

type Client struct {
	network    string
	endpoint   string
	httpClient *http.Client
}

func NewClient(network string, endpoint string) (*Client, error) {
	client := Client{
		network:    network,
		endpoint:   endpoint,
		httpClient: http.DefaultClient,
	}

	return &client, nil
}

func (c *Client) GetLogs(ctx context.Context, parameter GetLogsRequest) ([]EthereumLog, error) {
	values, err := query.Values(parameter)
	if err != nil {
		return nil, err
	}

	url := &url.URL{
		Scheme:   "http",
		Host:     c.endpoint,
		Path:     fmt.Sprintf("/networks/%v/logs", c.network),
		RawQuery: values.Encode(),
	}

	request, err := http_utils.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	result := []EthereumLog{}

	err = http_utils.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
