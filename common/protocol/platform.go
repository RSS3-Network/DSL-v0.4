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

	PlatformLens                 = "Lens"
	PlatformLensOrb              = "Orb"
	PlatformLensLenster          = "Lenster"
	PlatformLensLenstube         = "Lenstube"
	PlatformLensLenstubeBytes    = "lenstube-bytes"
	PlatformLensLensterCrowdfund = "Lenster Crowdfund"
	PlatformLensButtrfly         = "Buttrfly"

	PlatformMatters = "Matters"
	PlatformRara    = "RARA"

	// collectible
	PlatformPOAP              = "POAP"
	PlatformGalaxy            = "Galaxy"
	PlatformEns               = "ENS Registrar"
	PlatformSpaceID           = "Space ID"
	PlatformUnstoppableDomain = "Unstoppable"
	PlatformAvvy              = "Avvy"
	PlatformArb               = "Arb"
	PlatformSound             = "Sound"
	PlatformLink3             = "Link3"

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
	PlatformParaswap       = "Paraswap"
	PlatformDODO           = "DODO"
	PlatformBalancer       = "Balancer"
	PlatformVelodrome      = "Velodrome"
	PlatformOpenSea        = "OpenSea"
	PlatformGem            = "Gem"
	PlatformQuix           = "Quix"
	PlatformLooksRare      = "LooksRare"
	PlatformTofuNFT        = "tofuNFT"
	PlatformBlur           = "Blur"
	PlatformElement        = "Element"
	PlatformHop            = "Hop"
	PlatformCow            = "Cow"
	PlatformMaskNetwork    = "MaskNetwork"
	PlatformNSwap          = "NSwap"
	PlatformRSS3           = "RSS3"
	PlatformZora           = "Zora"
	PlatformNouns          = "Nouns"
	PlatformFoundation     = "Foundation"
	PlatformZerion         = "Zerion"
	PlatformBendDAO        = "BendDAO"
	PlatformFriendTech     = "friend.tech"

	// metaverse
	PlatformMars4      = "Mars4"
	PlatformAavegotchi = "Aavegotchi"
	PlatformCarv       = "Carv"
	PlatformPlanetIX   = "PlanetIX"

	PlatformGnosisSafe = "Gnosis Safe"
)

var PlatformList = map[string][]string{
	filter.TagSocial: {
		PlatformMirror,
		PlatformLens,
		PlatformLensOrb,
		PlatformLensLenster,
		PlatformCrossbell,
		PlatformCrossbellXLog,
		PlatformFarcaster,
		NetworkEIP1577,
		PlatformIQWiki,
		PlatformRara,
	},
	filter.TagCollectible: {
		PlatformPOAP,
		PlatformGalaxy,
		PlatformEns,
		PlatformOpenSea,
		PlatformGem,
		PlatformQuix,
		PlatformLooksRare,
		PlatformTofuNFT,
		PlatformBlur,
		PlatformElement,
		PlatformNSwap,
		PlatformSpaceID,
		PlatformUnstoppableDomain,
		PlatformAvvy,
		PlatformSound,
		PlatformZora,
		PlatformNouns,
		PlatformFoundation,
		PlatformLink3,
		PlatformFriendTech,
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
		PlatformQuickSwap,
		PlatformRainbow,
		PlatformSpookySwap,
		PlatformParaswap,
		PlatformDODO,
		PlatformVelodrome,
		PlatformZerion,
		PlatformBendDAO,
	},
}
