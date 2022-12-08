package marketplace

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/blur"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/gem"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/looksrare"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/opensea"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/quix"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/tofunft"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

var UnsupportedPlatform = errors.New("unsupported platform")

var platformMap = map[common.Address]string{
	opensea.AddressSeaport:          protocol.PlatformOpenSea,
	opensea.AddressWyvernExchangeV1: protocol.PlatformOpenSea,
	opensea.AddressWyvernExchangeV2: protocol.PlatformOpenSea,
	quix.AddressSeaport:             protocol.PlatformQuix,
	quix.AddressExchangeV5:          protocol.PlatformQuix,
	quix.AddressWrapperSeaportProxy: protocol.PlatformQuix,
	looksrare.AddressExchange:       protocol.PlatformLooksRare,
	gem.AddressSwap1:                protocol.PlatformGem,
	gem.AddressSwap2:                protocol.PlatformGem,
	uniswap.AddressUniversalRouter:  protocol.PlatformUniswap,
	tofunft.AddressMarketplace:      protocol.PlatformTofuNFT,
	blur.AddressMarketplace:         protocol.PlatformBlur,
	blur.AddressMarketplace2:        protocol.PlatformBlur,
}
