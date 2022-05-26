package protocol

const (
	NetworkEthereum          = "ethereum"
	NetworkBinanceSmartChain = "binance_smart_chain"
	NetworkPolygon           = "polygon"
	NetworkZkSync            = "zksync"
	NetworkXDAI              = "xdai"
	NetworkArweave           = "arweave"
)

func NetworkToID(networkName string) string {
	switch networkName {
	case NetworkEthereum:
		return "0x1"
	case NetworkBinanceSmartChain:
		return "0x38"
	case NetworkPolygon:
		return "0x89"
	default:
		return "0x0"
	}
}
