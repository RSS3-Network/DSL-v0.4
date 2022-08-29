package pregod_etl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/naturalselectionlabs/pregod/common/utils/httpx"
	"github.com/sirupsen/logrus"
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

func (c *Client) GetLogs(ctx context.Context, network string, parameter GetLogsRequest) (*GetLogsResponse, error) {
	values, err := query.Values(parameter)
	if err != nil {
		logrus.Errorf("[pregod_etl client] GetLogs: query.Values error, %v", err)

		return nil, err
	}

	url := &url.URL{
		Scheme:   "http",
		Host:     c.endpoint,
		Path:     fmt.Sprintf("/networks/%v/logs", network),
		RawQuery: values.Encode(),
	}

	request, err := httpx.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		logrus.Errorf("[pregod_etl client] GetLogs: NewRequest error, %v", err)

		return nil, err
	}

	var result *GetLogsResponse

	err = httpx.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		logrus.Errorf("[pregod_etl client] GetLogs: DoRequest error, %v", err)

		return nil, err
	}

	return result, nil
}
