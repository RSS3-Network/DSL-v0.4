package protocol

// NOTE: if you update the list, also update the list in the following files:
// - service/hub/internal/server/handler/handler_platform.go
const (
	// social
	PlatformMirror        = "Mirror"
	PlatformLens          = "Lens"
	PlatformLenster       = "Lenster"
	PlatformCrossbell     = "Crossbell"
	PlatformCrossbellXlog = "xLog"

	// collectible
	PlatformPOAP   = "POAP"
	PlatformGalaxy = "Galaxy"
	PlatformEns    = "ENS Registrar"

	// donation
	PlatformGitcoin = "Gitcoin"

	// governance
	PlatformSnapshot = "Snapshot"

	// exchange
	// dex
	PlatformUniswap     = "Uniswap"
	PlatformSushiswap   = "SushiSwap"
	PlatformPancakeswap = "PancakeSwap"
	Platform1inch       = "1inch"
	PlatformMetamask    = "MetaMask"
	Platform0x          = "0x"
	PlatformAAVE        = "AAVE"
	PlatformCurve       = "Curve"
)
