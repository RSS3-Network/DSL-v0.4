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
		PlatformCrossbell,
		PlatformCrossbellXLog,
		PlatformCrossbellXCast,
		PlatformCrossbellXSync,
		PlatformFarcaster,
		PlatformIQWiki,

		PlatformLens,
		PlatformLensOrb,
		PlatformLensLenster,
		PlatformLensLenstube,
		PlatformLensLenstubeBytes,
		PlatformLensLensterCrowdfund,

		PlatformMatters,
		PlatformRara,
	},
	filter.TagCollectible: {
		PlatformPOAP,
		PlatformGalaxy,
		PlatformEns,
		PlatformSpaceID,
		PlatformUnstoppableDomain,
		PlatformAvvy,
		PlatformArb,
		PlatformSound,
		PlatformLink3,
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
		PlatformLido,
		PlatformPolygonStaking,
		PlatformTraderJoe,
		PlatformRainbow,
		PlatformQuickSwap,
		PlatformKyberSwap,
		PlatformSpookySwap,
		PlatformParaswap,
		PlatformDODO,
		PlatformBalancer,
		PlatformVelodrome,
		PlatformOpenSea,
		PlatformGem,
		PlatformQuix,
		PlatformLooksRare,
		PlatformTofuNFT,
		PlatformBlur,
		PlatformElement,
		PlatformHop,
		PlatformCow,
		PlatformMaskNetwork,
		PlatformNSwap,
		PlatformRSS3,
		PlatformZora,
		PlatformNouns,
		PlatformFoundation,
		PlatformZerion,
		PlatformBendDAO,
	},
	filter.TagMetaverse: {
		PlatformMars4,
		PlatformAavegotchi,
		PlatformCarv,
		PlatformPlanetIX,

		PlatformGnosisSafe,
	},
}
