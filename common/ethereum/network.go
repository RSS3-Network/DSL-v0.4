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
	default:
		return ""
	}
}
