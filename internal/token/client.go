package token

import (
	"errors"
)

var (
	ErrorTokenNotFound         = errors.New("token not found")
	ErrorNetworkNotSupported   = errors.New("network not supported")
	ErrorInvalidMetadataFormat = errors.New("invalid metadata format")
)

type Client struct{}

func New() *Client {
	return &Client{}
}
