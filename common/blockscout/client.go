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

	Endpoint = "blockscout.com"

	NetworkXDAI            = "xdai"
	NetworkEthereum        = "eth"
	NetworkEthereumClassic = "etc"
)

type Client struct {
	network    string
	httpClient *http.Client
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

type GetTransactionListOption struct {
	Module          string `url:"module"`
	Action          string `url:"action"`
	Address         string `url:"address"`
	ContractAddress string `url:"contractaddress,omitempty"`
	Sort            string `url:"sort,omitempty"`
	StartBlock      int64  `url:"startblock,omitempty"`
	EndBlock        int64  `url:"endblock,omitempty"`
	Page            int64  `url:"page,omitempty"`
	Offset          int64  `url:"offset,omitempty"`
	FilterBy        string `url:"filterby,omitempty"`
	StartTimestamp  int64  `url:"starttimestamp,omitempty"`
	EndTimestamp    int64  `url:"endtimestamp,omitempty"`
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
		Host:     Endpoint,
		Path:     "/xdai/mainnet/api",
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
		Host:     Endpoint,
		Path:     fmt.Sprintf("/%s/mainnet/api", c.network),
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

type GetTransactionInfoOption struct {
	Module          string `url:"module"`
	Action          string `url:"action"`
	TransactionHash string `url:"txhash"`
}

func (c *Client) GetTransactionInfo(ctx context.Context, transactionHash string, option *GetTransactionInfoOption) (*TransactionInfo, *Response, error) {
	option.Module = "account"
	option.TransactionHash = transactionHash
	option.Action = "gettxinfo"

	values, err := query.Values(option)
	if err != nil {
		return nil, nil, err
	}

	requestURL := &url.URL{
		Scheme:   Scheme,
		Host:     Endpoint,
		Path:     fmt.Sprintf("/%s/mainnet/api", c.network),
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

	var transactionInfo TransactionInfo

	if err := json.Unmarshal(response.Result, &transactionInfo); err != nil {
		return nil, nil, err
	}

	return &transactionInfo, response, nil
}

func New(network string) *Client {
	return &Client{
		network:    network,
		httpClient: http.DefaultClient,
	}
}
