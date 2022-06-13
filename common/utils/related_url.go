package utils

// Returns related urls based on the network and contract tx hash.
func GetTxHashURL(
	network string,
	transactionHash string,
) string {
	switch network {
	case "ethereum":
		return "https://etherscan.io/tx/" + (transactionHash)

	case "polygon":
		return "https://polygonscan.com/tx/" + (transactionHash)

	case "bnb":
		return "https://bscscan.com/tx/" + (transactionHash)

	case "avalanche":
		return "https://avascan.info/blockchain/c/tx/" + (transactionHash)
	case "fantom":
		return "https://ftmscan.com/tx/" + (transactionHash)
	case "zksync":
		return "https://zkscan.io/explorer/transactions/" + (transactionHash)
	default:
		return ""
	}
}
