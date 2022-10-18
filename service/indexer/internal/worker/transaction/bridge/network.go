package bridge

import "github.com/naturalselectionlabs/pregod/common/database/model/metadata"

const (
	ChainIDEthereum     = 1
	ChainIDOptimism     = 10
	ChainIDPolygon      = 137
	ChainIDArbitrumOne  = 42161
	ChainIDArbitrumNova = 42170
)

var networkMap = map[uint64]metadata.Network{
	ChainIDEthereum: {
		Name:    "Ethereum",
		ChainID: 1,
		Symbol:  "ETH",
	},
	ChainIDOptimism: {
		Name:    "Optimism",
		ChainID: 10,
		Symbol:  "ETH",
	},
	ChainIDPolygon: {
		Name:    "Polygon",
		ChainID: 137,
		Symbol:  "MATIC",
	},
	ChainIDArbitrumOne: {
		Name:    "Arbitrum One",
		ChainID: 42161,
		Symbol:  "ETH",
	},
	ChainIDArbitrumNova: {
		Name:    "Arbitrum Nova",
		ChainID: 42170,
		Symbol:  "ETH",
	},
}
