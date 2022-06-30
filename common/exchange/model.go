package exchange

import "github.com/naturalselectionlabs/pregod/common/protocol"

// UniSwapV2 uses `pairs`, UniSwapV3 uses `pools`
const (
	UniSwapV2 = "UniSwapV2"
	UniSwapV3 = "UniSwapV3"
)

// SwapPool
// OrderByVolumeUSD: if the pools should be ordered by `volumeUSD`,
// when there are more than `Limit` pools available from the endpoint,
// use OrderByVolumeUSD to get the top pools
// NonSubgraph: if the endpoint is not a graphql subgraph
// Limit: the number of pools to return, default is 6000
type SwapPool struct {
	Endpoint         string
	Network          string
	Protocol         string
	OrderByVolumeUSD bool
	NonSubgraph      bool
	Limit            int
}

type SwapProvider struct {
	Name      string
	SwapPools []SwapPool
}

var SwapProviders = []SwapProvider{
	// Endpoints for UniSwap
	{
		Name: "UniSwap",
		SwapPools: []SwapPool{
			{
				Endpoint:         "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2",
				Network:          protocol.NetworkEthereum,
				Protocol:         UniSwapV2,
				OrderByVolumeUSD: true,
			},
			{
				Endpoint:         "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3",
				Network:          protocol.NetworkEthereum,
				Protocol:         UniSwapV3,
				OrderByVolumeUSD: true,
			},
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-polygon",
				Network:  protocol.NetworkPolygon,
				Protocol: UniSwapV3,
			},
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/ianlapham/arbitrum-dev",
				Network:  protocol.NetworkArbitrum,
				Protocol: UniSwapV3,
			},
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/ianlapham/optimism-post-regenesis",
				Network:  protocol.NetworkOptimism,
				Protocol: UniSwapV3,
			},
		},
	},
	// Endpoints for SushiSwap
	{
		Name: "SushiSwap",
		SwapPools: []SwapPool{
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/sushiswap/bsc-exchange",
				Network:  protocol.NetworkBinanceSmartChain,
				Protocol: UniSwapV2,
			},
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/sushiswap/xdai-exchange",
				Network:  protocol.NetworkXDAI,
				Protocol: UniSwapV2,
			},
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/sushiswap/matic-exchange",
				Network:  protocol.NetworkPolygon,
				Protocol: UniSwapV2,
			},
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/sushiswap/arbitrum-exchange",
				Network:  protocol.NetworkArbitrum,
				Protocol: UniSwapV2,
			},
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/sushiswap/fantom-exchange",
				Network:  protocol.NetworkArbitrum,
				Protocol: UniSwapV2,
			},
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/sushiswap/avalanche-exchange",
				Network:  protocol.NetworkAvalanche,
				Protocol: UniSwapV2,
			},
		},
	},
	// Endpoints for Crypto.com
	{
		Name: "Crypto.com",
		SwapPools: []SwapPool{
			{
				Endpoint: "https://api.thegraph.com/subgraphs/name/crypto-com/stake-subgraph-v2",
				Network:  protocol.NetworkBinanceSmartChain,
				Protocol: UniSwapV2,
			},
		},
	},
	// Endpoints for 1inch.io
	{
		Name: "1inch.io",
		SwapPools: []SwapPool{
			{
				Endpoint:    "https://governance.1inch.io/v1.1/protocol/pairs",
				Network:     protocol.NetworkEthereum,
				Protocol:    UniSwapV2,
				NonSubgraph: true,
			},
			{
				Endpoint:    "https://governance.1inch.io/v1.2/56/protocol/pairs",
				Network:     protocol.NetworkBinanceSmartChain,
				Protocol:    UniSwapV2,
				NonSubgraph: true,
			},
		},
	},
	// Endpoints for PancakeSwap
	{
		Name: "PancakeSwap",
		SwapPools: []SwapPool{
			{
				Endpoint: "https://bsc.streamingfast.io/subgraphs/name/pancakeswap/exchange-v2",
				Network:  protocol.NetworkBinanceSmartChain,
				Protocol: UniSwapV2,
				Limit:    2000,
			},
		},
	},
}
