package protocol

const (
	NetworkEthereum          = "ethereum"
	NetworkEthereumClassic   = "ethereum_classic"
	NetworkBinanceSmartChain = "binance_smart_chain"
	NetworkPolygon           = "polygon"
	NetworkZkSync            = "zksync"
	NetworkXDAI              = "xdai"
	NetworkArweave           = "arweave"
	NetworkArbitrum          = "arbitrum"
	NetworkOptimism          = "optimism"
	NetworkFantom            = "fantom"
	NetworkCelo              = "celo"
	NetworkAvalanche         = "avalanche"
	NetworkCrossbell         = "crossbell"
	NetworkEIP1577           = "EIP-1577"
	NetworkFarcaster         = "farcaster"
	NetworkAptos             = "aptos"
)

var SupportNetworks = []string{
	NetworkEthereum,
	NetworkEthereumClassic,
	NetworkBinanceSmartChain,
	NetworkPolygon,
	NetworkZkSync,
	NetworkXDAI,
	NetworkArweave,
	NetworkArbitrum,
	NetworkOptimism,
	NetworkFantom,
	NetworkAvalanche,
	NetworkCrossbell,
	NetworkFarcaster,
	NetworkAptos,
}

var EthclientNetworks = []string{
	NetworkEthereum,
	NetworkBinanceSmartChain,
	NetworkPolygon,
	NetworkXDAI,
	NetworkCrossbell,
	NetworkArbitrum,
	NetworkOptimism,
	NetworkAvalanche,
	NetworkCelo,
	NetworkFantom,
}

func NetworkToID(networkName string) string {
	switch networkName {
	case NetworkEthereum:
		return "0x1"
	case NetworkBinanceSmartChain:
		return "0x38"
	case NetworkPolygon:
		return "0x89"
	case NetworkZkSync:
		return "0x118"
	case NetworkXDAI:
		return "0x64"
	case NetworkArbitrum:
		return "0xA4B1"
	case NetworkOptimism:
		return "0xA"
	case NetworkFantom:
		return "0xFA"
	case NetworkCelo:
		return "0xA4EC"
	case NetworkAvalanche:
		return "0xA86A"
	default:
		return "0x0"
	}
}
