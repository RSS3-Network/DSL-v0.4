package bridge

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/arbitrum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/hop"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/optimism"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/polygon"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

var UnsupportedPlatform = errors.New("unsupported platform")

var addressesMap = map[string][]common.Address{
	protocol.PlatformHop: {
		hop.AddressL2Bridge,
		hop.AddressL2AMMWrapperPolygonETH,
	},
	polygon.PlatformBridge: {
		polygon.AddressEtherePredicate,
		polygon.AddressERC20Predicate,
		polygon.AddressMintableERC20Predicate,
	},
	optimism.PlatformBridge: {
		optimism.AddressGateway,
	},
	arbitrum.PlatformInboxOne: {
		arbitrum.AddressInboxOne,
	},
	arbitrum.PlatformInboxNova: {
		arbitrum.AddressInboxNova,
	},
}

var platformMap = map[common.Address]string{}

func init() {
	for platform, addresses := range addressesMap {
		for _, address := range addresses {
			platformMap[address] = platform
		}
	}
}
