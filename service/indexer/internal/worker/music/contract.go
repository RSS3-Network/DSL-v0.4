package music

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/sound"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

type ContractAction struct {
	Network         string
	Platform        string
	Tag             string
	Type            string
	Action          string
	ContractAddress *common.Address
}

var EventsRouterMap = map[common.Hash]ContractAction{
	sound.EventEditionCreated: {
		Network:  protocol.NetworkEthereum,
		Platform: protocol.PlatformSound,
		Tag:      filter.TagCollectible,
		Type:     filter.CollectibleMusic,
		Action:   filter.CollectibleMusicRelease,
	},
	sound.EventEditionPurchased: {
		Network:  protocol.NetworkEthereum,
		Platform: protocol.PlatformSound,
		Tag:      filter.TagCollectible,
		Type:     filter.CollectibleMusic,
		Action:   filter.CollectibleMusicBuyEdition,
	},
	sound.EventMint: {
		Network:         protocol.NetworkEthereum,
		Platform:        protocol.PlatformSound,
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleMusic,
		Action:          filter.CollectibleMusicMint,
		ContractAddress: &sound.MintContractAddress,
	},
}
