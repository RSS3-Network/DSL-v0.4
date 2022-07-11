package poap

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	Endpoint = "https://api.poap.xyz"
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) GetActions(address string) ([]Action, error) {
	response, err := c.httpClient.Get(fmt.Sprintf("%s/actions/scan/%s", Endpoint, address))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	actions := make([]Action, 0)

	if err := json.NewDecoder(response.Body).Decode(&actions); err != nil {
		return nil, err
	}

	return actions, nil
}

func (c *Client) GetToken(ctx context.Context, id int64) (*Token, error) {
	response, err := c.httpClient.Get(fmt.Sprintf("%s/token/%d", Endpoint, id))
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	token := Token{}

	if err := json.NewDecoder(response.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}
