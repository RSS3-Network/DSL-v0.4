package transaction

import (
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

var NativeTokenMap = map[string]transaction.Token{
	protocol.NetworkEthereum: {
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimal:  18,
		Standard: protocol.TokenStandardNative,
		Network:  protocol.NetworkEthereum,
	},
	protocol.NetworkPolygon: {
		Name:     "Polygon",
		Symbol:   "MATIC",
		Decimal:  18,
		Standard: protocol.TokenStandardNative,
		Network:  protocol.NetworkPolygon,
	},
	protocol.NetworkBinanceSmartChain: {
		Name:     "BNB",
		Symbol:   "BNB",
		Decimal:  18,
		Standard: protocol.TokenStandardNative,
		Network:  protocol.NetworkBinanceSmartChain,
	},
	protocol.NetworkCrossbell: {
		Name:     "Crossbell",
		Symbol:   "CSB",
		Decimal:  18,
		Standard: protocol.TokenStandardNative,
		Network:  protocol.NetworkCrossbell,
	},
	protocol.NetworkXDAI: {
		Name:     "xDAI",
		Symbol:   "XDAI",
		Decimal:  18,
		Standard: protocol.TokenStandardNative,
		Network:  protocol.NetworkXDAI,
	},
}
