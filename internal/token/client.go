package token

import (
	"errors"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrorTokenNotFound         = errors.New("token not found")
	ErrorNetworkNotSupported   = errors.New("network not supported")
	ErrorInvalidMetadataFormat = errors.New("invalid metadata format")
)

type Client struct {
	ethereumClientMap map[string]*ethclient.Client
}

func New(ethereumClientMap map[string]*ethclient.Client) *Client {
	return &Client{
		ethereumClientMap: ethereumClientMap,
	}
}
