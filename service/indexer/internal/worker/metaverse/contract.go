package metaverse

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/aavegotchi"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/planet_ix"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

type MetaverseContract struct {
	Network  string
	Platform string
	Type     string
}

var (
	Mars4MintContractAddress        = common.HexToAddress("0xdf9aa1012fa49dc1d2a306e3e65ef1797d2b5fbb")
	CarvAchievementsContractAddress = common.HexToAddress("0xc2f24ffe96a69e381a747dc73fcd51492e29a0a4")
	CarvEventsContractAddress       = common.HexToAddress("0xdda56260fcb1b1c6ba1d84c5f99f5507d556a04b")
)

var Agent = map[common.Address]string{
	Mars4MintContractAddress:        "Mars4 Mint",
	CarvEventsContractAddress:       "Carv Events",
	CarvAchievementsContractAddress: "Carv Achievements",
}

var RouterMap = map[common.Address]MetaverseContract{
	Mars4MintContractAddress: {
		Network:  protocol.NetworkEthereum,
		Platform: protocol.PlatformMars4,
		Type:     filter.MetaverseMint,
	},
	CarvAchievementsContractAddress: {
		Network:  protocol.NetworkBinanceSmartChain,
		Platform: protocol.PlatformCarv,
		Type:     filter.MetaverseMint,
	},
	CarvEventsContractAddress: {
		Network:  protocol.NetworkBinanceSmartChain,
		Platform: protocol.PlatformCarv,
		Type:     filter.MetaverseMint,
	},
	aavegotchi.AavegotchiContractAddress: {
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformAavegotchi,
	},
	planet_ix.PlanetIXMarketplaceContractAddress: {
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformPlanetIX,
	},
}

var EventsRouterMap = map[common.Address]map[common.Hash]MetaverseContract{
	aavegotchi.AavegotchiContractAddress:         AavegotchiEvents,
	planet_ix.PlanetIXMarketplaceContractAddress: PlanetIXEvents,
}

var AavegotchiEvents = map[common.Hash]MetaverseContract{
	aavegotchi.EventERC1155ExecutedListing: {
		Type:     filter.MetaverseTrade,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformAavegotchi,
	},
	aavegotchi.EventERC1155ListingCancelled: {
		Type:     filter.MetaverseUnlist,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformAavegotchi,
	},
	aavegotchi.EventERC1155ListingAdd: {
		Type:     filter.MetaverseList,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformAavegotchi,
	},
	aavegotchi.EventERC721ExecutedListing: {
		Type:     filter.MetaverseTrade,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformAavegotchi,
	},
	aavegotchi.EventERC721ListingAdd: {
		Type:     filter.MetaverseList,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformAavegotchi,
	},
	aavegotchi.EventClaimAavegotchi: {
		Type:     filter.MetaverseClaim,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformAavegotchi,
	},
}

var PlanetIXEvents = map[common.Hash]MetaverseContract{
	planet_ix.EventOwnershipTransferred: {
		Type:     filter.MetaverseGift,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformPlanetIX,
	},
	planet_ix.EventSaleCancelled: {
		Type:     filter.MetaverseUnlist,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformPlanetIX,
	},
	planet_ix.EventSaleRequested: {
		Type:     filter.MetaverseList,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformPlanetIX,
	},
	planet_ix.EventPurchasedWithSignature: {
		Type:     filter.MetaverseTrade,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformPlanetIX,
	},
	planet_ix.EventPurchased: {
		Type:     filter.MetaverseTrade,
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformPlanetIX,
	},
}
