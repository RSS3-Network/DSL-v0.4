package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

func New(config *configx.RPC) (ethereumClientMap map[string]*ethclient.Client, err error) {
	ethereumClientMap = make(map[string]*ethclient.Client)

	if config.General.Ethereum != nil {
		if ethereumClientMap[protocol.NetworkEthereum], err = ethclient.Dial(config.General.Ethereum.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Polygon != nil {
		if ethereumClientMap[protocol.NetworkPolygon], err = ethclient.Dial(config.General.Polygon.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.BinanceSmartChain != nil {
		if ethereumClientMap[protocol.NetworkBinanceSmartChain], err = ethclient.Dial(config.General.BinanceSmartChain.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.XDAI != nil {
		if ethereumClientMap[protocol.NetworkXDAI], err = ethclient.Dial(config.General.XDAI.HTTP); err != nil {
			return nil, err
		}
	}

	if config.General.Crossbell != nil {
		if ethereumClientMap[protocol.NetworkCrossbell], err = ethclient.Dial(config.General.Crossbell.HTTP); err != nil {
			return nil, err
		}
	}

	return ethereumClientMap, nil
}
