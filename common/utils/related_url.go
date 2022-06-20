package utils

import "github.com/naturalselectionlabs/pregod/common/protocol"

// Returns related urls based on the network and contract tx hash.
func GetTxHashURL(
	network string,
	transactionHash string,
) string {
	switch network {
	case protocol.NetworkEthereum:
		return "https://etherscan.io/tx/" + (transactionHash)

	case protocol.NetworkPolygon:
		return "https://polygonscan.com/tx/" + (transactionHash)

	case protocol.NetworkBinanceSmartChain:
		return "https://bscscan.com/tx/" + (transactionHash)

	case protocol.NetworkAvalanche:
		return "https://avascan.info/blockchain/c/tx/" + (transactionHash)
	case protocol.NetworkFantom:
		return "https://ftmscan.com/tx/" + (transactionHash)
	case protocol.NetworkZkSync:
		return "https://zkscan.io/explorer/transactions/" + (transactionHash)
	default:
		return ""
	}
}
