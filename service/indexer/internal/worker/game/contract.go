package game 

import (
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)



type GameContract struct {
	Network         string 
	Platform        string 
	Type            string 
	EventHash string
}


var routerMap = map[string]GameContract{
	"0xdf9aa1012fa49dc1d2a306e3e65ef1797d2b5fbb": GameContract {
		Network: protocol.NetworkEthereum,
		Platform: protocol.PlatformMars4,
		Type: filter.GameMint,
	},
}