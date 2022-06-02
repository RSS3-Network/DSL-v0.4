package arweave

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	graphqlx "github.com/naturalselectionlabs/pregod/common/arweave/graphql"
	"github.com/shurcooL/graphql"
)

const (
	EndpointScheme = "https"
	EndpointHost   = "arweave.net"
	EndpointPath   = "/graphql"
)

type Client struct {
	httpClient    *http.Client
	graphqlClient *graphql.Client
}

func (c *Client) GetInfo(_ context.Context) (*Info, *http.Response, error) {
	requestURL := &url.URL{
		Scheme: EndpointScheme,
		Host:   EndpointHost,
		Path:   "/info",
	}

	response, err := c.httpClient.Get(requestURL.String())
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	info := Info{}

	if err := json.NewDecoder(response.Body).Decode(&info); err != nil {
		return nil, nil, err
	}

	return &info, response, nil
}

func (c *Client) GetBlockByHeight(_ context.Context, height int64) (*Block, *http.Response, error) {
	requestURL := &url.URL{
		Scheme: EndpointScheme,
		Host:   EndpointHost,
		Path:   fmt.Sprintf("/block/height/%d", height),
	}

	response, err := c.httpClient.Get(requestURL.String())
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	block := Block{}

	if err := json.NewDecoder(response.Body).Decode(&block); err != nil {
		return nil, nil, err
	}

	return &block, response, nil
}

type GetTransactionsVariable struct {
	Owners []graphql.String      `json:"owners"`
	Tags   []graphqlx.TagFilter  `json:"tags"`
	Block  *graphqlx.BlockFilter `json:"block"`
}

func (c *Client) GetTransactions(ctx context.Context, query interface{}, variable GetTransactionsVariable) error {
	variableMap := map[string]interface{}{}

	if variable.Owners != nil && len(variable.Owners) > 0 {
		variableMap["owners"] = variable.Owners
	}

	if variable.Tags != nil && len(variable.Tags) > 0 {
		variableMap["tags"] = variable.Tags
	}

	if variable.Block != nil {
		variableMap["block"] = variable.Block
	}

	if err := c.graphqlClient.Query(ctx, query, variableMap); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetFile(ctx context.Context, hash string) (io.ReadCloser, error) {
	requestURL := &url.URL{
		Scheme: EndpointScheme,
		Host:   EndpointHost,
		Path:   fmt.Sprintf("/%s", hash),
	}

	response, err := c.httpClient.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func NewClient() *Client {
	client := &Client{
		httpClient: http.DefaultClient,
	}

	endpointURL := url.URL{
		Scheme: EndpointScheme,
		Host:   EndpointHost,
		Path:   EndpointPath,
	}

	client.graphqlClient = graphql.NewClient(endpointURL.String(), client.httpClient)

	return client
}
