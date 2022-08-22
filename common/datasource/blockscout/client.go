package blockscout

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-querystring/query"
)

const (
	Scheme = "https"

	EndpointDefault   = "blockscout.com"
	EndpointCrossbell = "scan.crossbell.io"

	NetworkXDAI            = "/xdai/mainnet"
	NetworkEthereum        = "/eth/mainnet"
	NetworkEthereumClassic = "/etc/mainnet"
	NetworkCrossbell       = "" // Root path

	StatusFailed  = "0"
	StatusSuccess = "1"
)

type Client struct {
	network    string
	httpClient *http.Client
	endpoint   string
}

func (c *Client) DoRequest(_ context.Context, request *http.Request) (*Response, *http.Response, error) {
	httpResponse, err := c.httpClient.Do(request)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	response := &Response{}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, httpResponse, err
	}

	return response, httpResponse, err
}

func (c *Client) GetTransactionList(ctx context.Context, address common.Address, option *GetTransactionListOption) ([]Transaction, *Response, error) {
	option.Module = "account"
	option.Address = address.String()
	option.Action = "txlist"

	values, err := query.Values(option)
	if err != nil {
		return nil, nil, err
	}

	requestURL := &url.URL{
		Scheme:   Scheme,
		Host:     c.endpoint,
		Path:     fmt.Sprintf("%s/api", c.network),
		RawQuery: values.Encode(),
	}

	request, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, _, err := c.DoRequest(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	transactions := make([]Transaction, 0)

	if err := json.Unmarshal(response.Result, &transactions); err != nil {
		return nil, nil, err
	}

	return transactions, response, nil
}

type GetTokenTransactionListOption struct {
	Module          string `url:"module"`
	Action          string `url:"action"`
	Address         string `url:"address"`
	ContractAddress string `url:"contractaddress,omitempty"`
	Sort            string `url:"sort,omitempty"`
	StartBlock      int64  `url:"startblock,omitempty"`
	EndBlock        int64  `url:"endblock,omitempty"`
	Page            int64  `url:"page,omitempty"`
	Offset          int64  `url:"offset,omitempty"`
}

func (c *Client) GetTokenTransactionList(ctx context.Context, address common.Address, option *GetTokenTransactionListOption) ([]Transaction, *Response, error) {
	option.Module = "account"
	option.Address = address.String()
	option.Action = "tokentx"

	values, err := query.Values(option)
	if err != nil {
		return nil, nil, err
	}

	requestURL := &url.URL{
		Scheme:   Scheme,
		Host:     c.endpoint,
		Path:     fmt.Sprintf("%s/api", c.network),
		RawQuery: values.Encode(),
	}

	request, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, _, err := c.DoRequest(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	transactions := make([]Transaction, 0)

	if err := json.Unmarshal(response.Result, &transactions); err != nil {
		return nil, nil, err
	}

	return transactions, response, nil
}

func New(endpoint, network string) *Client {
	return &Client{
		endpoint:   endpoint,
		network:    network,
		httpClient: http.DefaultClient,
	}
}
