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
		return fmt.Sprintf("https://scan.crossbell.io/tx/%s", transactionHash)
	default:
		return ""
	}
}

func BuildTokenURL(network, address, id string) string {
	switch network {
	case protocol.NetworkEthereum:
		return fmt.Sprintf("https://opensea.io/assets/%s/%s", address, id)
	case protocol.NetworkPolygon:
		return fmt.Sprintf("https://opensea.io/assets/matic/%s/%s", address, id)
	case protocol.NetworkBinanceSmartChain:
		return fmt.Sprintf("https://tofunft.com/nft/bsc/%s/%s", address, id)
	default:
		return ""
	}
}

func BuildURL(urls []string, url string) []string {
	if url == "" {
		return urls
	}

	return append(urls, url)
}
