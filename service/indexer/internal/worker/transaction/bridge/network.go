package bridge

import (
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
)

const (
	ChainIDEthereum          = 1
	ChainIDOptimism          = 10
	ChainIDBinanceSmartChain = 56
	ChainIDGnosis            = 100
	ChainIDPolygon           = 137
	ChainIDFantom            = 250
	ChainIDCelo              = 42220
	ChainIDArbitrumOne       = 42161
	ChainIDArbitrumNova      = 42170
	ChainIDAvalanche         = 43114
)

var networkMap = map[uint64]transaction.TargetNetwork{
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
	ChainIDBinanceSmartChain: {
		Name:    "Binance Smart Chain",
		ChainID: 56,
		Symbol:  "BNB",
	},
	ChainIDGnosis: {
		Name:    "Gnosis",
		ChainID: 100,
		Symbol:  "XDAI",
	},
	ChainIDPolygon: {
		Name:    "Polygon",
		ChainID: 137,
		Symbol:  "MATIC",
	},
	ChainIDFantom: {
		Name:    "Fantom",
		ChainID: 250,
		Symbol:  "FTM",
	},
	ChainIDCelo: {
		Name:    "Celo",
		ChainID: 42220,
		Symbol:  "CELO",
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
	ChainIDAvalanche: {
		Name:    "Avalanche",
		ChainID: 43114,
		Symbol:  "AVAX",
	},
}
