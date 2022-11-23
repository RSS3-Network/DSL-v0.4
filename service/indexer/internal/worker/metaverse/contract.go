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
	Mars4MintContractAddress  = strings.ToLower("0xdf9aa1012fa49dc1d2a306e3e65ef1797d2b5fbb")
	AavegotchiContractAddress = strings.ToLower("0x86935f11c86623dec8a25696e1c19a8659cbf95d")
)

var Agent = map[string]string{
	Mars4MintContractAddress: "Mars4 Mint",
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
	},
}
