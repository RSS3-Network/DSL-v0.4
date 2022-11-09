package metaverse

import (
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

type MetaverseContract struct {
	Network   string
	Platform  string
	Type      string
	EventHash string
}

var routerMap = map[string]MetaverseContract{
	"0xdf9aa1012fa49dc1d2a306e3e65ef1797d2b5fbb": {
		Network:  protocol.NetworkEthereum,
		Platform: protocol.PlatformMars4,
		Type:     filter.MetaverseMint,
	},
}
