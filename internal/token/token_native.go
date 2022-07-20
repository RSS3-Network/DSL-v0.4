package token

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

var nativeMap = map[string]Native{
	protocol.NetworkEthereum: {
		Name:     "Ethereum", // https://www.coingecko.com/en/coins/ethereum
		Symbol:   "ETH",
		Decimals: 18,
		Logo:     "https://assets.coingecko.com/coins/images/279/large/ethereum.png",
	},
	protocol.NetworkPolygon: {
		Name:     "Polygon", // https://www.coingecko.com/en/coins/polygon
		Symbol:   "MATIC",
		Decimals: 18,
		Logo:     "https://assets.coingecko.com/coins/images/2713/large/matic-token-icon.png",
	},
	protocol.NetworkBinanceSmartChain: {
		Name:     "BNB", // https://www.coingecko.com/en/coins/bnb
		Symbol:   "BNB",
		Logo:     "https://assets.coingecko.com/coins/images/825/large/bnb-icon2_2x.png",
		Decimals: 18,
	},
	protocol.NetworkXDAI: {
		Name:     "XDAI", // https://www.coingecko.com/en/coins/xdai
		Symbol:   "XDAI",
		Logo:     "https://assets.coingecko.com/coins/images/11062/large/bnb-icon2_2x.png",
		Decimals: 18,
	},
	protocol.NetworkCrossbell: {
		Name:     "Crossbell",
		Symbol:   "CSB",
		Logo:     "", // TODO Add logo
		Decimals: 18,
	},
	protocol.NetworkZkSync: {
		Name:     "Ethereum", // https://www.coingecko.com/en/coins/ethereum
		Symbol:   "ETH",
		Logo:     "https://assets.coingecko.com/coins/images/279/large/ethereum.png",
		Decimals: 18,
	},
}

func (c *Client) Native(context context.Context, network string) (*Native, error) {
	token, exists := nativeMap[network]
	if !exists {
		return nil, ErrorTokenNotFound
	}

	return &token, nil
}
