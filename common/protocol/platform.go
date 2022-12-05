package protocol

import "github.com/naturalselectionlabs/pregod/common/protocol/filter"

const (
	// social
	PlatformMirror         = "Mirror"
	PlatformCrossbell      = "Crossbell"
	PlatformCrossbellXLog  = "xLog"
	PlatformCrossbellXCast = "xCast"
	PlatformCrossbellXSync = "xSync"
	PlatformFarcaster      = "Farcaster"
	PlatformIQWiki         = "IQ.Wiki"

	PlatformLens              = "Lens"
	PlatformLensLenster       = "Lenster"
	PlatformLensLenstube      = "Lenstube"
	PlatformLensLenstubeBytes = "lenstube-bytes"

	PlatformMatters = "Matters"

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
	PlatformUniswap        = "Uniswap"
	PlatformSushiswap      = "SushiSwap"
	PlatformPancakeswap    = "PancakeSwap"
	Platform1inch          = "1inch"
	PlatformMetamask       = "MetaMask"
	Platform0x             = "0x"
	PlatformAAVE           = "AAVE"
	PlatformCurve          = "Curve"
	PlatformLido           = "Lido"
	PlatformPolygonStaking = "Polygon Staking"
	PlatformTraderJoe      = "TraderJoe"
	PlatformRainbow        = "Rainbow"
	PlatformQuickSwap      = "QuickSwap"
	PlatformKyberSwap      = "KyberSwap"
	PlatformSpookySwap     = "SpookySwap"
	PlatformDODO           = "DODO"
	PlatformBalancer       = "Balancer"
	PlatformVelodrome      = "Velodrome"
	PlatformOpenSea        = "OpenSea"
	PlatformGem            = "Gem"
	PlatformQuix           = "Quix"
	PlatformLooksRare      = "LooksRare"

	// metaverse
	PlatformMars4      = "Mars4"
	PlatformAavegotchi = "Aavegotchi"
	PlatformCarv       = "Carv"

	PlatformParaswap = "Paraswap"
)

var PlatformList = map[string][]string{
	filter.TagSocial: {
		PlatformMirror,
		PlatformLens,
		PlatformLensLenster,
		PlatformCrossbell,
		PlatformCrossbellXLog,
		PlatformFarcaster,
		NetworkEIP1577,
		PlatformIQWiki,
	},
	filter.TagCollectible: {
		PlatformPOAP,
		PlatformGalaxy,
		PlatformEns,
	},
	filter.TagDonation: {
		PlatformGitcoin,
	},
	filter.TagGovernance: {
		PlatformSnapshot,
	},
	filter.TagExchange: {
		PlatformUniswap,
		PlatformSushiswap,
		PlatformPancakeswap,
		Platform1inch,
		PlatformMetamask,
		Platform0x,
		PlatformAAVE,
		PlatformCurve,
		PlatformTraderJoe,
	},
}
