package bridge

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/arbitrum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/hop"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/optimism"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/polygon"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/socket"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

var UnsupportedPlatform = errors.New("unsupported platform")

var addressesMap = map[string][]common.Address{
	protocol.PlatformHop: {
		hop.AddressBridgeL1EthereumUSDC,
		hop.AddressBridgeL1EthereumUSDT,
		hop.AddressBridgeL1EthereumMATIC,
		hop.AddressBridgeL1EthereumDAI,
		hop.AddressBridgeL1EthereumETH,
		hop.AddressBridgeL1EthereumHOP,
		hop.AddressBridgeL1EthereumSNX,
		hop.AddressBridgeL2GnosisUSDC,
		hop.AddressBridgeL2GnosisUSDT,
		hop.AddressBridgeL2GnosisMATIC,
		hop.AddressBridgeL2GnosisDAI,
		hop.AddressBridgeL2GnosisETH,
		hop.AddressBridgeL2GnosisHOP,
		hop.AddressBridgeL2PolygonUSDC,
		hop.AddressBridgeL2PolygonUSDT,
		hop.AddressBridgeL2PolygonMATIC,
		hop.AddressBridgeL2PolygonDAI,
		hop.AddressBridgeL2PolygonETH,
		hop.AddressBridgeL2PolygonHOP,
		hop.AddressBridgeL2OptimismUSDC,
		hop.AddressBridgeL2OptimismUSDT,
		hop.AddressBridgeL2OptimismDAI,
		hop.AddressBridgeL2OptimismETH,
		hop.AddressBridgeL2OptimismHOP,
		hop.AddressBridgeL2ArbitrumUSDC,
		hop.AddressBridgeL2ArbitrumUSDT,
		hop.AddressBridgeL2ArbitrumDAI,
		hop.AddressBridgeL2ArbitrumETH,
		hop.AddressBridgeL2ArbitrumHOP,
		hop.AddressBridgeL2ArbitrumSNX,
		hop.AddressAMMWrapperL2GnosisUSDC,
		hop.AddressAMMWrapperL2PolygonUSDC,
		hop.AddressAMMWrapperL2OptimismUSDC,
		hop.AddressAMMWrapperL2ArbitrumUSDC,
		hop.AddressAMMWrapperL2GnosisUSDT,
		hop.AddressAMMWrapperL2PolygonUSDT,
		hop.AddressAMMWrapperL2OptimismUSDT,
		hop.AddressAMMWrapperL2ArbitrumUSDT,
		hop.AddressAMMWrapperL2GnosisMATIC,
		hop.AddressAMMWrapperL2PolygonMATIC,
		hop.AddressAMMWrapperL2GnosisDAI,
		hop.AddressAMMWrapperL2PolygonDAI,
		hop.AddressAMMWrapperL2OptimismDAI,
		hop.AddressAMMWrapperL2ArbitrumDAI,
		hop.AddressAMMWrapperL2GnosisETH,
		hop.AddressAMMWrapperL2PolygonETH,
		hop.AddressAMMWrapperL2OptimismETH,
		hop.AddressAMMWrapperL2ArbitrumETH,
		hop.AddressAMMWrapperL2OptimismSNX,
	},
	polygon.PlatformBridge: {
		polygon.AddressEtherePredicate,
		polygon.AddressERC20Predicate,
		polygon.AddressMintableERC20Predicate,
		polygon.AddressBridge,
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
	protocol.PlatformSocket: {
		socket.AddressRegistry,
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
