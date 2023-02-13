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
	case protocol.NetworkArbitrum:
		return fmt.Sprintf("https://arbiscan.io/tx/%s", transactionHash)
	case protocol.NetworkOptimism:
		return fmt.Sprintf("https://optimistic.etherscan.io/tx/%s", transactionHash)
	case protocol.NetworkAvalanche:
		return fmt.Sprintf("https://snowtrace.io/tx/%s", transactionHash)
	case protocol.NetworkCelo:
		return fmt.Sprintf("https://celoscan.io/tx/%s", transactionHash)
	case protocol.NetworkFantom:
		return fmt.Sprintf("https://ftmscan.com/tx/%s", transactionHash)
	default:
		return ""
	}
}

func BuildTokenURL(network, address, id string) []string {
	switch network {
	case protocol.NetworkEthereum:
		return []string{
			fmt.Sprintf("https://opensea.io/assets/%s/%s", address, id),
			fmt.Sprintf("https://looksrare.org/collections/%s/%s", address, id),
		}
	case protocol.NetworkPolygon:
		return []string{
			fmt.Sprintf("https://opensea.io/assets/matic/%s/%s", address, id),
		}
	case protocol.NetworkBinanceSmartChain:
		return []string{
			fmt.Sprintf("https://tofunft.com/nft/bsc/%s/%s", address, id),
		}
	case protocol.NetworkArbitrum:
		return []string{
			fmt.Sprintf("https://opensea.io/assets/arbitrum/%s/%s", address, id),
		}
	case protocol.NetworkOptimism:
		return []string{
			fmt.Sprintf("https://qx.app/asset/%s/%s", address, id),
		}
	case protocol.NetworkAvalanche:
		return []string{
			fmt.Sprintf("https://opensea.io/assets/avalanche/%s/%s", address, id),
		}
	default:
		return []string{}
	}
}

func BuildURL(old []string, new ...string) []string {
	if len(new) == 0 {
		return old
	}

	return append(old, new...)
}
