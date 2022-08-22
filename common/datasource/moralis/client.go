package moralis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-querystring/query"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/sirupsen/logrus"
)

const (
	Scheme   = "https"
	Endpoint = "deep-index.moralis.io"
)

type Client struct {
	key        string
	httpClient *http.Client
}

func (c *Client) NewRequest(method, rawURL string, body interface{}) (*http.Request, error) {
	var readWriter io.ReadWriter

	if body != nil {
		if err := json.NewEncoder(readWriter).Encode(body); err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, rawURL, readWriter)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", protocol.ContentTypeJSON)
	request.Header.Set("X-API-Key", c.key)

	return request, nil
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
		var p []byte
		httpResponse.Body.Read(p) // nolint:errcheck
		logrus.Errorf("moralis decode error: %v %s", httpResponse.StatusCode, string(p))
		return nil, httpResponse, err
	}

	return response, httpResponse, err
}

func (c *Client) GetTransactions(ctx context.Context, address common.Address, option *GetTransactionsOption) ([]Transaction, *Response, error) {
	values, err := query.Values(option)
	if err != nil {
		return nil, nil, err
	}

	requestURL := &url.URL{
		Scheme:   Scheme,
		Host:     Endpoint,
		Path:     fmt.Sprintf("/api/v2/%s", address),
		RawQuery: values.Encode(),
	}

	request, err := c.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, httpResponse, err := c.DoRequest(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	transactions := make([]Transaction, 0)

	if err := json.Unmarshal(response.Result, &transactions); err != nil {
		var p []byte
		_, _ = httpResponse.Body.Read(p)
		logrus.Error(string(p), httpResponse.StatusCode)
		return nil, nil, err
	}

	return transactions, response, nil
}

func (c *Client) GetTokenTransfers(ctx context.Context, address common.Address, option *GetTokenTransfersOption) ([]TokenTransfer, *Response, error) {
	values, err := query.Values(option)
	if err != nil {
		return nil, nil, err
	}

	requestURL := &url.URL{
		Scheme:   Scheme,
		Host:     Endpoint,
		Path:     fmt.Sprintf("/api/v2/%s/erc20/transfers", address),
		RawQuery: values.Encode(),
	}

	request, err := c.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, httpResponse, err := c.DoRequest(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	tokenTransfers := make([]TokenTransfer, 0)

	if err := json.Unmarshal(response.Result, &tokenTransfers); err != nil {
		var p []byte
		_, _ = httpResponse.Body.Read(p)
		logrus.Error(string(p), httpResponse.StatusCode)
		return nil, nil, err
	}

	return tokenTransfers, response, nil
}

func (c *Client) GetNFTTransfers(ctx context.Context, address common.Address, option *GetNFTTransfersOption) ([]NFTTransfer, *Response, error) {
	values, err := query.Values(option)
	if err != nil {
		return nil, nil, err
	}

	requestURL := &url.URL{
		Scheme:   Scheme,
		Host:     Endpoint,
		Path:     fmt.Sprintf("/api/v2/%s/nft/transfers", address),
		RawQuery: values.Encode(),
	}

	request, err := c.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	response, httpResponse, err := c.DoRequest(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	nftTransfers := make([]NFTTransfer, 0)

	if err := json.Unmarshal(response.Result, &nftTransfers); err != nil {
		var p []byte
		_, _ = httpResponse.Body.Read(p)
		logrus.Error(string(p), httpResponse.StatusCode)
		return nil, nil, err
	}

	return nftTransfers, response, nil
}

func NewClient(key string) *Client {
	return &Client{
		key:        key,
		httpClient: http.DefaultClient,
	}
}
