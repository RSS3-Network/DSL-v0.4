package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const Endpoint = "api.coingecko.com"

type Client struct {
	httpClient *http.Client
}

func (c *Client) buildURL(path string) (*url.URL, error) {
	internalURL := url.URL{
		Scheme: "https",
		Host:   Endpoint,
	}

	internalURL.Scheme = "https"
	internalURL.Path = fmt.Sprintf("/api/v3%s", path)

	return &internalURL, nil
}

func (c *Client) CoinList(ctx context.Context, parameter CoinListParameter) ([]Coin, error) {
	values, err := query.Values(parameter)
	if err != nil {
		return nil, err
	}

	internalURL, err := c.buildURL("/coins/list")
	if err != nil {
		return nil, err
	}

	internalURL.RawQuery = values.Encode()

	response, err := c.httpClient.Get(internalURL.String())
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	var result []Coin

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) Coin(ctx context.Context, id string, parameter CoinParameter) (*Coin, error) {
	values, err := query.Values(parameter)
	if err != nil {
		return nil, err
	}

	internalURL, err := c.buildURL(fmt.Sprintf("/coins/%s", id))
	if err != nil {
		return nil, err
	}

	internalURL.RawQuery = values.Encode()

	response, err := c.httpClient.Get(internalURL.String())
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	var result Coin

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}
