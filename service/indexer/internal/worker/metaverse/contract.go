package metaverse

import (
	"strings"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

type MetaverseContract struct {
	Network   string
	Platform  string
	Type      string
	EventHash string
}

var (
	Mars4MintContractAddress        = strings.ToLower("0xdf9aa1012fa49dc1d2a306e3e65ef1797d2b5fbb")
	AavegotchiContractAddress       = strings.ToLower("0x86935f11c86623dec8a25696e1c19a8659cbf95d")
	CarvAchievementsContractAddress = strings.ToLower("0xc2f24ffe96a69e381a747dc73fcd51492e29a0a4")
	CarvEventsContractAddress       = strings.ToLower("0xdda56260fcb1b1c6ba1d84c5f99f5507d556a04b")
)

var Agent = map[string]string{
	Mars4MintContractAddress:        "Mars4 Mint",
	CarvEventsContractAddress:       "Carv Events",
	CarvAchievementsContractAddress: "Carv Achievements",
}

var RouterMap = map[string]MetaverseContract{
	Mars4MintContractAddress: {
		Network:  protocol.NetworkEthereum,
		Platform: protocol.PlatformMars4,
		Type:     filter.MetaverseMint,
	},
	AavegotchiContractAddress: {
		Network:  protocol.NetworkPolygon,
		Platform: protocol.PlatformAavegotchi,
		Type:     filter.MetaverseMint,
	},
	CarvAchievementsContractAddress: {
		Network:  protocol.NetworkBinanceSmartChain,
		Platform: protocol.PlatformCarv,
	},
	CarvEventsContractAddress: {
		Network:  protocol.NetworkBinanceSmartChain,
		Platform: protocol.PlatformCarv,
		Type:     filter.MetaverseMint,
	},
}
