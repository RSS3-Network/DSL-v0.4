package ethereum

import (
	"fmt"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

func BuildScanURL(network string, transactionHash string) string {
	switch network {
	case protocol.NetworkEthereum:
		return fmt.Sprintf("https://etherscan.io/tx/%s", transactionHash)
	case protocol.NetworkPolygon:
		return fmt.Sprintf("https://polygonscan.com/tx/%s", transactionHash)
	case protocol.NetworkBinanceSmartChain:
		return fmt.Sprintf("https://bscscan.com/tx/%s", transactionHash)
	case protocol.NetworkXDAI:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tx/%s", transactionHash)
	case protocol.NetworkCrossbell:
		return fmt.Sprintf("https://crossbellscan.com/tx/%s", transactionHash)
	default:
		return ""
	}
}
