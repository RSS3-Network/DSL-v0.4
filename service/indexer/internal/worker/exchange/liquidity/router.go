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

	routerPancake = Router{
		Name:     protocol.PlatformPancakeswap,
		Protocol: "Uniswap V2",
	}

	routerSushi = Router{
		Name:     protocol.PlatformSushiswap,
		Protocol: "Uniswap V2",
	}

	routerMap = map[string]Router{
		// Uniswap V3
		strings.ToLower("0xC36442b4a4522E871399CD717aBDD847Ab11FE88"): routerUniswapV3, // Uniswap V3 Positions NFT
		// Pancake
		strings.ToLower("0x10ED43C718714eb63d5aA57B78B54704E256024E"): routerPancake, // Pancake Router
		// Sushi
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushi, // Sushi Router
	}
)
