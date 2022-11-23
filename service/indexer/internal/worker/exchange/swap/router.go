package swap

import (
	"strings"

	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/treaderjoe"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Router struct {
	Name     string
	Protocol string
}

var (
	routerUniswapV2 = Router{
		Name:     protocol.PlatformUniswap,
		Protocol: "Uniswap V2",
	}

	routerUniswapV3 = Router{
		Name:     protocol.PlatformUniswap,
		Protocol: "Uniswap V3",
	}

	routerSushiSwap = Router{
		Name:     protocol.PlatformSushiswap,
		Protocol: "Uniswap V3",
	}

	pancakeSwap = Router{
		Name:     protocol.PlatformPancakeswap,
		Protocol: "Uniswap V2",
	}

	oneinchV1 = Router{
		Name:     protocol.Platform1inch,
		Protocol: "1inch Aggregation Protocol V1",
	}

	oneinchV2 = Router{
		Name:     protocol.Platform1inch,
		Protocol: "1inch Aggregation Protocol V2",
	}

	oneinchV3 = Router{
		Name:     protocol.Platform1inch,
		Protocol: "1inch Aggregation Protocol V3",
	}

	oneinchV4 = Router{
		Name:     protocol.Platform1inch,
		Protocol: "1inch Aggregation Protocol V4",
	}

	metamask = Router{
		Name:     protocol.PlatformMetamask,
		Protocol: "Metamask V1",
	}

	zeroXV4 = Router{
		Name:     protocol.Platform0x,
		Protocol: "0x Protocol V4",
	}

	zeroXV3 = Router{
		Name:     protocol.Platform0x,
		Protocol: "0x Protocol V3",
	}

	paraswap = Router{
		Name:     protocol.PlatformParaswap,
		Protocol: "Paraswap v5",
	}

	traderJoeV2 = Router{
		Name:     protocol.PlatformTraderJoe,
		Protocol: "Joe V2",
	}

	curveQuickSwap = Router{
		Name:     protocol.PlatformCurve,
		Protocol: "Curve Fi",
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
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushiSwap, // SushiSwap Polygon, Binanace Smart Chain, Avalanche, Gnosis/XDAI
		strings.ToLower("0xbE811A0D44E2553d25d11CB8DC0d3F0D0E6430E6"): routerSushiSwap, // SushiSwap Optimism
		strings.ToLower("0x1421bDe4B10e8dd459b3BCb598810B1337D56842"): routerSushiSwap, // SushiSwap Celo
		// PancakeSwap
		// https://docs.pancakeswap.finance/code/smart-contracts/pancakeswap-exchange/router-v2
		strings.ToLower("0x10ED43C718714eb63d5aA57B78B54704E256024E"): pancakeSwap, // PancakeSwap V2
		// 1inch
		// https://etherscan.io/accounts/label/1inch-exchange
		strings.ToLower("0x11111254369792b2ca5d084ab5eea397ca8fa48b"): oneinchV1, // 1inch V1
		strings.ToLower("0x111111125434b319222cdbf8c261674adb56f3ae"): oneinchV2, // 1inch V2
		strings.ToLower("0x11111112542d85b3ef69ae05771c2dccff4faa26"): oneinchV3, // 1inch V3
		strings.ToLower("0x1111111254fb6c44bac0bed2854e76f90643097d"): oneinchV4, // 1inch V4
		// Metamask
		// https://etherscan.io/accounts/label/metamask
		strings.ToLower("0x881d40237659c251811cec9c364ef91dc08d300c"): metamask, // Metamask Ethereum
		strings.ToLower("0x1a1ec25DC08e98e5E93F1104B5e5cdD298707d31"): metamask, // Metamask Polygon
		// 0x
		// https://docs.0x.org/developer-resources/contract-addresses
		strings.ToLower("0xDef1C0ded9bec7F1a1670819833240f027b25EfF"): zeroXV4, // 0x Exchange V4
		strings.ToLower("0x61935CbDd02287B511119DDb11Aeb42F1593b7Ef"): zeroXV3, // 0x Exchange V3
		// Paraswap
		// https://developers.paraswap.network/smart-contracts
		strings.ToLower("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57"): paraswap, // protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain, protocol.NetworkAvalanche, protocol.NetworkFantom, protocol.NetworkArbitrum, protocol.NetworkOptimism
		// TraderJoe
		// https://docs.traderjoexyz.com/deployment-addresses/avalanche
		strings.ToLower(treaderjoe.AddressRouter.String()): traderJoeV2,
		// Curve QuickSwap
		strings.ToLower("0x55B916Ce078eA594c10a874ba67eCc3d62e29822"): curveQuickSwap, // Ethereum
		strings.ToLower("0xa522deb6F17853F3a97a65d0972a50bDC3B1AFFF"): curveQuickSwap, // Polygon
		strings.ToLower("0x16243caB3aC4d8eE8df7660a525F7F7539962468"): curveQuickSwap, // Fantom
		strings.ToLower("0x890f4e345B1dAED0367A877a1612f86A1f86985f"): curveQuickSwap, // Avalanche
		strings.ToLower("0xE6358f6a45B502477e83CC1CDa759f540E4459ee"): curveQuickSwap, // Gnosis
	}
)
