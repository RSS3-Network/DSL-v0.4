package ethclientx

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"golang.org/x/exp/slices"
)

var globalLocker sync.RWMutex

var globalEthereumClientMap = make(map[string]*ethclient.Client)

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
		if ethereumClientMap[protocol.NetworkEthereum], err = ethclient.Dial(config.General.Ethereum.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Polygon != nil && slices.Contains(networks, protocol.NetworkPolygon) {
		if ethereumClientMap[protocol.NetworkPolygon], err = ethclient.Dial(config.General.Polygon.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.BinanceSmartChain != nil && slices.Contains(networks, protocol.NetworkBinanceSmartChain) {
		if ethereumClientMap[protocol.NetworkBinanceSmartChain], err = ethclient.Dial(config.General.BinanceSmartChain.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.XDAI != nil && slices.Contains(networks, protocol.NetworkXDAI) {
		if ethereumClientMap[protocol.NetworkXDAI], err = ethclient.Dial(config.General.XDAI.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Crossbell != nil && slices.Contains(networks, protocol.NetworkCrossbell) {
		if ethereumClientMap[protocol.NetworkCrossbell], err = ethclient.Dial(config.General.Crossbell.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Optimism != nil && slices.Contains(networks, protocol.NetworkOptimism) {
		if ethereumClientMap[protocol.NetworkOptimism], err = ethclient.Dial(config.General.Optimism.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Avalanche != nil && slices.Contains(networks, protocol.NetworkAvalanche) {
		if ethereumClientMap[protocol.NetworkAvalanche], err = ethclient.Dial(config.General.Avalanche.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Celo != nil && slices.Contains(networks, protocol.NetworkCelo) {
		if ethereumClientMap[protocol.NetworkCelo], err = ethclient.Dial(config.General.Celo.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Fantom != nil && slices.Contains(networks, protocol.NetworkFantom) {
		if ethereumClientMap[protocol.NetworkFantom], err = ethclient.Dial(config.General.Fantom.HTTP); err != nil {
			return nil, err
		}
	}

	return ethereumClientMap, nil
}
