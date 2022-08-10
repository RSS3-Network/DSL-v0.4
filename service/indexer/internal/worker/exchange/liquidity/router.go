package liquidity

import (
	"strings"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Router struct {
	Name     string
	Protocol string
}

var (
	routerUniswapV3 = Router{
		Name:     protocol.PlatformUniswap,
		Protocol: "Uniswap V3",
	}

	routerMap = map[string]Router{
		// Uniswap V3
		strings.ToLower("0xC36442b4a4522E871399CD717aBDD847Ab11FE88"): routerUniswapV3, // Uniswap V3 Positions NFT
	}
)
