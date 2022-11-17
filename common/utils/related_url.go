package utils

import (
	"fmt"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

// Returns related urls based on the network and contract tx hash.
func GetTxHashURL(network string, transactionHash string) string {
	switch network {
	case protocol.NetworkEthereum:
		return fmt.Sprintf("https://etherscan.io/tx/%s", transactionHash)
	case protocol.NetworkPolygon:
		return fmt.Sprintf("https://polygonscan.com/tx/%s", transactionHash)
	case protocol.NetworkBinanceSmartChain:
		return fmt.Sprintf("https://bscscan.com/tx/%s", transactionHash)
	case protocol.NetworkAvalanche:
		return fmt.Sprintf("https://snowtrace.io/tx/%s", transactionHash)
	case protocol.NetworkFantom:
		return fmt.Sprintf("https://ftmscan.com/tx/%s", transactionHash)
	case protocol.NetworkOptimism:
		return fmt.Sprintf("https://optimistic.etherscan.io/tx/%s", transactionHash)
	case protocol.NetworkZkSync:
		return fmt.Sprintf("https://zkscan.io/explorer/transactions/%s", transactionHash)
	default:
		return ""
	}
}
