package ethclientx

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	configx "github.com/naturalselectionlabs/pregod/common/config"
)

var (
	globalLocker            sync.RWMutex
	globalEthereumClientMap = make(map[string]*ethclient.Client)
)

var (
	globalUrlLocker      sync.RWMutex
	globalEthereumUrlMap = make(map[string]string)
)

func GlobalUrl(network string) (string, error) {
	globalUrlLocker.RLock()

	defer globalUrlLocker.RUnlock()

	url, exist := globalEthereumUrlMap[network]
	if !exist {
		return "", fmt.Errorf("invalid network %s for Url", url)
	}

	return url, nil
}

func Global(network string) (*ethclient.Client, error) {
	globalLocker.RLock()

	defer globalLocker.RUnlock()

	client, exist := globalEthereumClientMap[network]
	if !exist {
		return nil, fmt.Errorf("invalid network %s for ethclient", network)
	}

	return client, nil
}

func ReplaceGlobal(network string, ethereumClient *ethclient.Client) {
	globalLocker.Lock()

	defer globalLocker.Unlock()

	globalEthereumClientMap[network] = ethereumClient
}

func Dial(config *configx.RPC, networks []string) (ethereumClientMap map[string]*ethclient.Client, err error) {
	ethereumClientMap = make(map[string]*ethclient.Client)

	for _, network := range networks {
		globalEthereumUrlMap[network] = config.General.WebSocket(network)
		if globalEthereumUrlMap[network] == "" {
			globalEthereumUrlMap[network] = config.General.HTTP(network)
		}
		if ethereumClientMap[network], err = ethclient.Dial(globalEthereumUrlMap[network]); err != nil {
			return nil, fmt.Errorf("failed to dial %s: %w", network, err)
		}
	}

	return ethereumClientMap, nil
}
