package ethclientx

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/kurora/common/client/ethereum"
	configx "github.com/naturalselectionlabs/pregod/common/config"
)

var (
	globalLocker             sync.RWMutex
	globalEthereumClientMap  = make(map[string]*ethclient.Client)
	globalEthereumXClientMap = make(map[string]ethereum.Client)
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

func GlobalX(network string) (ethereum.Client, error) {
	globalLocker.RLock()
	defer globalLocker.RUnlock()

	client, exist := globalEthereumXClientMap[network]
	if !exist {
		return nil, fmt.Errorf("invalid network %s for ethclient", network)
	}

	return client, nil
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
		// Preferred use of WebSocket endpoints
		globalEthereumUrlMap[network] = config.General.WebSocket(network)

		if globalEthereumUrlMap[network] == "" {
			globalEthereumUrlMap[network] = config.General.HTTP(network)
		}

		// Ethereum standard client
		if ethereumClientMap[network], err = ethclient.Dial(globalEthereumUrlMap[network]); err != nil {
			return nil, fmt.Errorf("failed to dial %s %s: %w", network, globalEthereumUrlMap[network], err)
		}

		// Register global ethereum client
		registerGlobal := func() error {
			globalLocker.Lock()
			defer globalLocker.Unlock()

			// No transaction verification client
			if globalEthereumXClientMap[network], err = ethereum.Dial(context.TODO(), globalEthereumUrlMap[network]); err != nil {
				return fmt.Errorf("dial %s: %w", network, err)
			}

			return nil
		}

		if err = registerGlobal(); err != nil {
			return nil, fmt.Errorf("register global %s: %w", network, err)
		}
	}

	return ethereumClientMap, nil
}
