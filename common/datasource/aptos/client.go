package aptos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/naturalselectionlabs/pregod/common/utils/httpx"
)

const (
	Endpoint = "fullnode.mainnet.aptoslabs.com"
	MaxCount = 1000
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) GetAccountTractions(ctx context.Context, parameter GetAccountTransactionsParameter) ([]GetAccountTransactionsResult, error) {
	values, err := query.Values(parameter)
	if err != nil {
		return nil, err
	}

	url := url.URL{
		Scheme:   "https",
		Host:     Endpoint,
		Path:     fmt.Sprintf("/v1/accounts/%v/transactions", parameter.Address),
		RawQuery: values.Encode(),
	}

	request, err := httpx.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var result []GetAccountTransactionsResult

	err = httpx.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}
