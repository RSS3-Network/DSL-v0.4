package token

import (
	"errors"

	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

var (
	ErrorTokenNotFound         = errors.New("token not found")
	ErrorNetworkNotSupported   = errors.New("network not supported")
	ErrorInvalidMetadataFormat = errors.New("invalid metadata format")
)

type Client struct {
	databaseClient    *gorm.DB
	ethereumClientMap map[string]*ethclient.Client
}

func New(databaseClient *gorm.DB, ethereumClientMap map[string]*ethclient.Client) *Client {
	return &Client{
		databaseClient:    databaseClient,
		ethereumClientMap: ethereumClientMap,
	}
}
