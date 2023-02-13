package token

import (
	"errors"
	"net/http"
)

var (
	ErrorTokenNotFound         = errors.New("token not found")
	ErrorNetworkNotSupported   = errors.New("network not supported")
	ErrorInvalidMetadataFormat = errors.New("invalid metadata format")
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{httpClient: &http.Client{}}
}
