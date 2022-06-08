package zksync

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
	Scheme   = "https"
	Endpoint = "api.zksync.io"

	FromLatest = "latest"

	DirectionNewer = "newer"
	DirectionOlder = "older"
)

type Client struct {
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

type GetAddressTransactionsOption struct {
	From      string `url:"from"`
	Limit     int64  `url:"limit"`
	Direction string `url:"direction"`
}

func (c *Client) GetAddressTransactions(ctx context.Context, address common.Address, from string, limit int64, direction string) (*GetAccountTransactionList, *Response, error) {
	values, err := query.Values(&GetAddressTransactionsOption{
		From:      from,
		Limit:     limit,
		Direction: direction,
	})
	if err != nil {
		return nil, nil, err
	}

	requestURL := &url.URL{
		Scheme:   Scheme,
		Host:     Endpoint,
		Path:     fmt.Sprintf("/api/v0.2/accounts/%s/transactions", address.String()),
		RawQuery: values.Encode(),
	}

	httpRequest, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, _, err := c.DoRequest(ctx, httpRequest)
	if err != nil {
		return nil, nil, err
	}

	transactionList := GetAccountTransactionList{}

	if err := json.Unmarshal(response.Result, &transactionList); err != nil {
		return nil, nil, err
	}

	return &transactionList, response, nil
}

func (c *Client) GetTransactionData(ctx context.Context, transactionHash common.Hash) (*GetTransactionData, *Response, error) {
	requestURL := &url.URL{
		Scheme: Scheme,
		Host:   Endpoint,
		Path:   fmt.Sprintf("/api/v0.2/transactions/%s/data", transactionHash.String()),
	}

	httpRequest, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, _, err := c.DoRequest(ctx, httpRequest)
	if err != nil {
		return nil, nil, err
	}

	transactionData := GetTransactionData{}

	if err := json.Unmarshal(response.Result, &transactionData); err != nil {
		return nil, nil, err
	}

	return &transactionData, response, nil
}

func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}
