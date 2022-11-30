package ethclientx

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"golang.org/x/exp/slices"
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

	if config.General.Ethereum != nil && slices.Contains(networks, protocol.NetworkEthereum) {
		globalEthereumUrlMap[protocol.NetworkEthereum] = config.General.Ethereum.HTTP
		if ethereumClientMap[protocol.NetworkEthereum], err = ethclient.Dial(config.General.Ethereum.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Polygon != nil && slices.Contains(networks, protocol.NetworkPolygon) {
		globalEthereumUrlMap[protocol.NetworkPolygon] = config.General.Polygon.HTTP
		if ethereumClientMap[protocol.NetworkPolygon], err = ethclient.Dial(config.General.Polygon.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.BinanceSmartChain != nil && slices.Contains(networks, protocol.NetworkBinanceSmartChain) {
		globalEthereumUrlMap[protocol.NetworkBinanceSmartChain] = config.General.BinanceSmartChain.HTTP
		if ethereumClientMap[protocol.NetworkBinanceSmartChain], err = ethclient.Dial(config.General.BinanceSmartChain.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.XDAI != nil && slices.Contains(networks, protocol.NetworkXDAI) {
		globalEthereumUrlMap[protocol.NetworkXDAI] = config.General.XDAI.HTTP
		if ethereumClientMap[protocol.NetworkXDAI], err = ethclient.Dial(config.General.XDAI.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Crossbell != nil && slices.Contains(networks, protocol.NetworkCrossbell) {
		globalEthereumUrlMap[protocol.NetworkCrossbell] = config.General.Crossbell.HTTP
		if ethereumClientMap[protocol.NetworkCrossbell], err = ethclient.Dial(config.General.Crossbell.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Optimism != nil && slices.Contains(networks, protocol.NetworkOptimism) {
		globalEthereumUrlMap[protocol.NetworkOptimism] = config.General.Optimism.HTTP
		if ethereumClientMap[protocol.NetworkOptimism], err = ethclient.Dial(config.General.Optimism.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Avalanche != nil && slices.Contains(networks, protocol.NetworkAvalanche) {
		globalEthereumUrlMap[protocol.NetworkAvalanche] = config.General.Avalanche.HTTP
		if ethereumClientMap[protocol.NetworkAvalanche], err = ethclient.Dial(config.General.Avalanche.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Celo != nil && slices.Contains(networks, protocol.NetworkCelo) {
		globalEthereumUrlMap[protocol.NetworkCelo] = config.General.Celo.HTTP
		if ethereumClientMap[protocol.NetworkCelo], err = ethclient.Dial(config.General.Celo.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Fantom != nil && slices.Contains(networks, protocol.NetworkFantom) {
		globalEthereumUrlMap[protocol.NetworkFantom] = config.General.Fantom.HTTP
		if ethereumClientMap[protocol.NetworkFantom], err = ethclient.Dial(config.General.Fantom.HTTP); err != nil {
			return nil, err
		}
	}

	return ethereumClientMap, nil
}
