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
	NetworkConflux           = "conflux"
	NetworkBaseGoerli        = "base_goerli"
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
	NetworkCelo,
	NetworkAvalanche,
	NetworkCrossbell,
	NetworkEIP1577,
	NetworkFarcaster,
	NetworkAptos,
	NetworkConflux,
	NetworkBaseGoerli,
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
	NetworkBaseGoerli,
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
	case NetworkBaseGoerli:
		return "0x14a33"
	default:
		return "0x0"
	}
}

func IdToNetwork(chainId string) string {
	switch chainId {
	case "0x1":
		return NetworkEthereum
	case "0x38":
		return NetworkBinanceSmartChain
	case "0x89":
		return NetworkPolygon
	case "0x118":
		return NetworkZkSync
	case "0x64":
		return NetworkXDAI
	case "0xA4B1":
		return NetworkArbitrum
	case "0xA":
		return NetworkOptimism
	case "0xFA":
		return NetworkFantom
	case "0xA4EC":
		return NetworkCelo
	case "0xA86A":
		return NetworkAvalanche
	case "0x14a33":
		return NetworkBaseGoerli
	default:
		return ""
	}
}
