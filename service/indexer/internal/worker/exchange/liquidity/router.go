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

	routerAAVEV2 = Router{
		Name:     protocol.PlatformAAVE,
		Protocol: "AAVE V2",
	}

	routerAAVEV3 = Router{
		Name:     protocol.PlatformAAVE,
		Protocol: "AAVE V3",
	}

	routerMap = map[string]Router{
		// Uniswap V3
		strings.ToLower("0xC36442b4a4522E871399CD717aBDD847Ab11FE88"): routerUniswapV3, // Uniswap V3 Positions NFT
		// Pancake
		strings.ToLower("0x10ED43C718714eb63d5aA57B78B54704E256024E"): routerPancake, // Pancake Router
		// Sushi
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushi, // Sushi Router
		// AAVE
		strings.ToLower("0xcc9a0B7c43DC2a5F023Bb9b738E45B0Ef6B06E04"): routerAAVEV2, // AAVE WETH Gateway V2 Ethereum
		strings.ToLower("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"): routerAAVEV2, // AAVE Lending Pool V2
		strings.ToLower("0x9BdB5fcc80A49640c7872ac089Cc0e00A98451B6"): routerAAVEV3, // AAVE WETH Gateway V3 Polygon
		strings.ToLower("0x794a61358D6845594F94dc1DB02A252b5b4814aD"): routerAAVEV3, // AAVE Pool V3
	}
)
