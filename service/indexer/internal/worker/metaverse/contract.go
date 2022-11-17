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

var Mars4MintContractAddress = strings.ToLower("0xdf9aa1012fa49dc1d2a306e3e65ef1797d2b5fbb")

var Agent = map[string]string{
	Mars4MintContractAddress: "Mars4 Mint",
}

var RouterMap = map[string]MetaverseContract{
	Mars4MintContractAddress: {
		Network:  protocol.NetworkEthereum,
		Platform: protocol.PlatformMars4,
		Type:     filter.MetaverseMint,
	},
}
