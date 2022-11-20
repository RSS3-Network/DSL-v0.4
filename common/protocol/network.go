package protocol

type Network struct {
	ChainID string
	Name    string
	IsEVM   bool
}

const (
	NetworkArbitrum          = "arbitrum"
	NetworkArweave           = "arweave"
	NetworkAvalanche         = "avalanche"
	NetworkBinanceSmartChain = "binance_smart_chain"
	NetworkCelo              = "celo"
	NetworkCrossbell         = "crossbell"
	NetworkEIP1577           = "EIP-1577"
	NetworkEthereum          = "ethereum"
	NetworkEthereumClassic   = "ethereum_classic"
	NetworkFantom            = "fantom"
	NetworkFarcaster         = "farcaster"
	NetworkOptimism          = "optimism"
	NetworkPolygon           = "polygon"
	NetworkXDAI              = "xdai"
	NetworkZkSync            = "zksync"
)

// SupportedNetworks all supported networks
var SupportedNetworks = getNetworks(true, false)

// EVMNetworks all supported EVM compatible networks
var EVMNetworks = getNetworks(false, true)

// NonEVMNetworks all supported Non-EVM compatible networks
var NonEVMNetworks = getNetworks(false, false)

var Networks = map[string]Network{
	NetworkArbitrum: {
		ChainID: "0xA4B1",
		Name:    NetworkArbitrum,
		IsEVM:   false,
	},
	NetworkArweave: {
		ChainID: "0x0",
		Name:    NetworkArweave,
		IsEVM:   false,
	},
	NetworkAvalanche: {
		ChainID: "0xA86A",
		Name:    NetworkAvalanche,
		IsEVM:   true,
	},
	NetworkBinanceSmartChain: {
		ChainID: "0x38",
		Name:    NetworkBinanceSmartChain,
		IsEVM:   true,
	},
	NetworkCelo: {
		ChainID: "0xA4EC",
		Name:    NetworkCelo,
		IsEVM:   true,
	},
	NetworkCrossbell: {
		ChainID: "0x0E99",
		Name:    NetworkCrossbell,
		IsEVM:   true,
	},
	NetworkEIP1577: {
		ChainID: "0x0",
		Name:    NetworkEIP1577,
		IsEVM:   false,
	},
	NetworkEthereum: {
		ChainID: "0x1",
		Name:    NetworkEthereum,
		IsEVM:   true,
	},
	NetworkEthereumClassic: {
		ChainID: "0x3D",
		Name:    NetworkEthereumClassic,
		IsEVM:   true,
	},
	NetworkFantom: {
		ChainID: "0xFA",
		Name:    NetworkFantom,
		IsEVM:   true,
	},
	NetworkFarcaster: {
		ChainID: "0x0",
		Name:    NetworkFarcaster,
		IsEVM:   false,
	},
	NetworkOptimism: {
		ChainID: "0xA",
		Name:    NetworkOptimism,
		IsEVM:   true,
	},
	NetworkPolygon: {
		ChainID: "0x89",
		Name:    NetworkPolygon,
		IsEVM:   true,
	},
	NetworkXDAI: {
		ChainID: "0x64",
		Name:    NetworkXDAI,
		IsEVM:   true,
	},
	NetworkZkSync: {
		ChainID: "0x118",
		Name:    NetworkZkSync,
		IsEVM:   false,
	},
}

func getNetworks(all bool, evmOnly bool) []string {
	var keys []string
	for k := range Networks {
		if all || Networks[k].IsEVM == evmOnly {
			keys = append(keys, k)
			continue
		}
	}
	return keys
}

func NetworkToID(networkName string) string {
	return Networks[networkName].ChainID
}

// MergeNetworks merge valid networks with NonEVMNetworks
func MergeNetworks(networks []string) []string {
	result := NonEVMNetworks
	result = append(result, networks...)
	return result
}
