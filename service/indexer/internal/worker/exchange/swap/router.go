package swap

import (
	"strings"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Router struct {
	Name     string
	Protocol string
}

var (
	routerUniswapV2 = Router{
		Name:     protocol.PlatformUniswap,
		Protocol: "UniSwapV2",
	}

	routerUniswapV3 = Router{
		Name:     protocol.PlatformUniswap,
		Protocol: "UniSwapV2",
	}

	routerSushiSwap = Router{
		Name:     protocol.PlatformSushiswap,
		Protocol: "UniSwapV3",
	}

	pancakeSwap = Router{
		Name:     protocol.PlatformPancakeswap,
		Protocol: "UniSwapV2",
	}

	oneinchV1 = Router{
		Name:     protocol.Platform1inch,
		Protocol: "Aggregation Protocol V1",
	}

	oneinchV2 = Router{
		Name:     protocol.Platform1inch,
		Protocol: "Aggregation Protocol V2",
	}

	oneinchV3 = Router{
		Name:     protocol.Platform1inch,
		Protocol: "Aggregation Protocol V3",
	}

	oneinchV4 = Router{
		Name:     protocol.Platform1inch,
		Protocol: "Aggregation Protocol V4",
	}

	routerMap = map[string]Router{
		// Uniswap V2
		strings.ToLower("0xf164fC0Ec4E93095b804a4795bBe1e041497b92a"): routerUniswapV2, // Uniswap V2 1
		strings.ToLower("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"): routerUniswapV2, // Uniswap V2 2
		// Uniswap V3
		// https://docs.uniswap.org/protocol/reference/deployments
		strings.ToLower("0xE592427A0AEce92De3Edee1F18E0157C05861564"): routerUniswapV3, // Uniswap V3 1
		strings.ToLower("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"): routerUniswapV3, // Uniswap V3 2
		// SushiSwap
		// https://docs.sushi.com/docs/Developers/Deployment%20Addresses
		strings.ToLower("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"): routerSushiSwap, // SushiSwap Ethereum
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushiSwap, // SushiSwap Polygon
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushiSwap, // SushiSwap Binance Smart Chain
		// PancakeSwap
		// https://docs.pancakeswap.finance/code/smart-contracts/pancakeswap-exchange/router-v2
		strings.ToLower("0x10ED43C718714eb63d5aA57B78B54704E256024E"): pancakeSwap, // PancakeSwap V2
		// 1inch
		// https://etherscan.io/accounts/label/1inch-exchange
		//strings.ToLower("0x11111254369792b2ca5d084ab5eea397ca8fa48b"): oneinchV1,  // 1inch V1
		//strings.ToLower("0x111111125434b319222cdbf8c261674adb56f3ae"): oneinchV2,  // 1inch V2
		//strings.ToLower("0x11111112542d85b3ef69ae05771c2dccff4faa26"): oneinchV3, // 1inch V3
		//strings.ToLower("0x1111111254fb6c44bac0bed2854e76f90643097d"): oneinchV4, // 1inch V4
	}
)
