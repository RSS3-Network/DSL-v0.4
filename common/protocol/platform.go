package protocol

// NOTE: if you update the list, also update the list in the following files:
// - service/hub/internal/server/handler/handler_platform.go
const (
	// social
	PlatfromMirror    = "Mirror"
	PlatfromLens      = "Lens"
	PlatfromCrossbell = "Crossbell"

	// collectible
	PlatfromPOAP   = "POAP"
	PlatfromGalaxy = "Galaxy"

	// donation
	PlatfromGitcoin = "Gitcoin"

	// governance
	PlatfromSnapshot = "Snapshot"

	// exchange
	// dex
	PlatformUniswap     = "Uniswap"
	PlatformSushiswap   = "SushiSwap"
	PlatformPancakeswap = "PancakeSwap"
	Platform1inch       = "1inch"
	PlatformMetamask    = "MetaMask"
)
