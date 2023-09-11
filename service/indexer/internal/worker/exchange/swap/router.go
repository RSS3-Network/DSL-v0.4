package swap

import (
	"strings"

	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/balancer"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/masknetwork"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/spookyswap"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/treaderjoe"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
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

	oneinchV5 = Router{
		Name:     protocol.Platform1inch,
		Protocol: "1inch Aggregation Protocol V5",
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

	quickSwapV2 = Router{
		Name:     protocol.PlatformQuickSwap,
		Protocol: "QuickSwap V2",
	}

	quickSwapV3 = Router{
		Name:     protocol.PlatformQuickSwap,
		Protocol: "QuickSwap V3",
	}

	rainbow = Router{
		Name:     protocol.PlatformRainbow,
		Protocol: "1inch Aggregation Protocol", // V4 or V5
	}

	kyberSwapAggregator = Router{
		Name:     protocol.PlatformKyberSwap,
		Protocol: "Kyberswap Aggregator",
	}

	kyberSwapElastic = Router{
		Name:     protocol.PlatformKyberSwap,
		Protocol: "KyberSwap Elastic",
	}

	kyberSwapClassic = Router{
		Name:     protocol.PlatformKyberSwap,
		Protocol: "KyberSwap Classic",
	}

	spookySwap = Router{
		Name:     protocol.PlatformSpookySwap,
		Protocol: "Uniswap V2",
	}

	dodoV2 = Router{
		Name:     protocol.PlatformDODO,
		Protocol: "DODO V2",
	}

	velodromeSwap = Router{
		Name:     protocol.PlatformVelodrome,
		Protocol: "Velodrome",
	}

	curveQuickSwap = Router{
		Name:     protocol.PlatformCurve,
		Protocol: "Curve Fi",
	}

	balancerSwap = Router{
		Name:     protocol.PlatformBalancer,
		Protocol: "Balancer",
	}

	cowSwap = Router{
		Name:     protocol.PlatformCow,
		Protocol: "Cow Protocol",
	}

	maskNetwork = Router{
		Name:     protocol.PlatformMaskNetwork,
		Protocol: "MaskNetwork ITO",
	}

	zerion = Router{
		Name:     protocol.PlatformZerion,
		Protocol: "Zerion",
	}

	routerMap = map[string]Router{
		// Uniswap V2
		strings.ToLower("0xf164fC0Ec4E93095b804a4795bBe1e041497b92a"): routerUniswapV2, // Uniswap V2 1
		strings.ToLower("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"): routerUniswapV2, // Uniswap V2 2
		// Uniswap V3
		// https://docs.uniswap.org/protocol/reference/deployments
		strings.ToLower("0xE592427A0AEce92De3Edee1F18E0157C05861564"): routerUniswapV3, // Uniswap V3 1
		strings.ToLower("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"): routerUniswapV3, // Uniswap V3 2
		// Uniswap V3 Universal Router
		strings.ToLower(uniswap.AddressUniversalRouter.String()):         routerUniswapV3, // Ethereum
		strings.ToLower(uniswap.AddressUniversalRouterArbitrum.String()): routerUniswapV3, // Arbitrum
		strings.ToLower(uniswap.AddressUniversalRouterCelo.String()):     routerUniswapV3, // Celo
		strings.ToLower(uniswap.AddressUniversalRouterOptimism.String()): routerUniswapV3, // Optimism
		strings.ToLower(uniswap.AddressUniversalRouterPolygon.String()):  routerUniswapV3, // Polygon
		strings.ToLower(uniswap.AddressUniversalRouterBase.String()):     routerUniswapV3, // Base
		// SushiSwap
		// https://docs.sushi.com/docs/Developers/Deployment%20Addresses
		strings.ToLower("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"): routerSushiSwap, // SushiSwap Ethereum
		strings.ToLower("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): routerSushiSwap, // SushiSwap Polygon, Binanace Smart Chain, Avalanche, Gnosis/XDAI
		strings.ToLower("0xbE811A0D44E2553d25d11CB8DC0d3F0D0E6430E6"): routerSushiSwap, // SushiSwap Optimism
		strings.ToLower("0x1421bDe4B10e8dd459b3BCb598810B1337D56842"): routerSushiSwap, // SushiSwap Celo
		// PancakeSwap
		// https://docs.pancakeswap.finance/code/smart-contracts/pancakeswap-exchange/router-v2
		strings.ToLower("0x10ED43C718714eb63d5aA57B78B54704E256024E"): pancakeSwap, // PancakeSwap V2 Binanace Smart Chain
		strings.ToLower("0xEfF92A263d31888d860bD50809A8D171709b7b1c"): pancakeSwap, // PancakeSwap V2 Ethereum
		// https://docs.pancakeswap.finance/code/smart-contracts/pancakeswap-exchange/smart-router
		strings.ToLower("0x64D74e1EAAe3176744b5767b93B7Bee39Cf7898F"): pancakeSwap, // PancakeSwap Smart Router
		// 1inch
		// https://etherscan.io/accounts/label/1inch-exchange
		strings.ToLower("0x11111254369792b2ca5d084ab5eea397ca8fa48b"): oneinchV1, // 1inch V1
		strings.ToLower("0x111111125434b319222cdbf8c261674adb56f3ae"): oneinchV2, // 1inch V2
		strings.ToLower("0x11111112542d85b3ef69ae05771c2dccff4faa26"): oneinchV3, // 1inch V3
		strings.ToLower("0x1111111254fb6c44bac0bed2854e76f90643097d"): oneinchV4, // 1inch V4
		strings.ToLower("0x1111111254EEB25477B68fb85Ed929f73A960582"): oneinchV5, // 1inch V5
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
		// QuickSwap
		// https://docs.quickswap.exchange/reference/smart-contracts/router02
		strings.ToLower("0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff"): quickSwapV2, // Market V2
		strings.ToLower("0xf5b509bB0909a69B1c207E495f687a596C168E12"): quickSwapV3, // Market V3
		// Rainbow
		// https://learn.rainbow.me/swap-with-confidence-with-rainbow
		strings.ToLower("0x00000000009726632680FB29d3F7A9734E3010E2"): rainbow, // Rainbow
		// KyberSwap
		// https://docs.kyberswap.com/
		strings.ToLower("0x617Dee16B86534a5d792A4d7A62FB491B544111E"): kyberSwapAggregator, // Meta Aggregation Router
		strings.ToLower("0x6131B5fae19EA4f9D964eAc0408E4408b66337b5"): kyberSwapElastic,    // Meta Aggregation Router V2
		strings.ToLower("0xC1e7dFE73E1598E3910EF4C7845B68A9Ab6F4c83"): kyberSwapElastic,    // Elastic
		strings.ToLower("0x1c87257f5e8609940bc751a07bb085bb7f8cdbe6"): kyberSwapClassic,    // Classic Ethereum
		strings.ToLower("0xEaE47c5D99f7B31165a7f0c5f7E0D6afA25CFd55"): kyberSwapClassic,    // Classic Arbitrum
		strings.ToLower("0x546C79662E028B661dFB4767664d0273184E4dD1"): kyberSwapClassic,    // Classic Polygon
		strings.ToLower("0x78df70615ffc8066cc0887917f2Cd72092C86409"): kyberSwapClassic,    // Classic Binanace Smart Chain
		strings.ToLower("0x8Efa5A9AD6D594Cf76830267077B78cE0Bc5A5F8"): kyberSwapClassic,    // Classic Avalanche
		strings.ToLower("0x5d5A5a0a465129848c2549669e12cDC2f8DE039A"): kyberSwapClassic,    // Classic Fantom
		strings.ToLower("0xEaE47c5D99f7B31165a7f0c5f7E0D6afA25CFd55"): kyberSwapClassic,    // Classic Optimism
		// SpookySwap
		// https://docs.spooky.fi/Resources/contracts
		strings.ToLower(spookyswap.AddressLiquidityBrewer.String()): spookySwap, // SpookySwap Liquidity Brewer
		strings.ToLower(spookyswap.AddressRouter.String()):          spookySwap, // SpookySwap Router
		// DODO
		// https://docs.dodoex.io/english/developers/contracts-address/ethereum
		strings.ToLower("0xa2398842F37465f89540430bDC00219fA9E4D28a"): dodoV2, // DODO V2 RouteProxy Ethereum
		strings.ToLower("0xa356867fDCEa8e71AEaF87805808803806231FdC"): dodoV2, // DODO V2 RouteProxy Ethereum 2
		strings.ToLower("0x6B3D817814eABc984d51896b1015C0b89E9737Ca"): dodoV2, // DODO V2 RouteProxy Binance Smart Chain
		strings.ToLower("0x8F8Dd7DB1bDA5eD3da8C9daf3bfa471c12d58486"): dodoV2, // DODO V2 RouteProxy Binance Smart Chain 2
		strings.ToLower("0x2fA4334cfD7c56a0E7Ca02BD81455205FcBDc5E9"): dodoV2, // DODO V2 RouteProxy Polygon
		strings.ToLower("0xa222e6a71D1A1Dd5F279805fbe38d5329C1d0e70"): dodoV2, // DODO V2 RouteProxy Polygon 2
		strings.ToLower("0x3B6067D4CAa8A14c63fdBE6318F27A0bBc9F9237"): dodoV2, // DODO V2 RouteProxy Arbitrum
		strings.ToLower("0x88CBf433471A0CD8240D2a12354362988b4593E5"): dodoV2, // DODO V2 RouteProxy Arbitrum 2
		strings.ToLower("0x409E377A7AfFB1FD3369cfc24880aD58895D1dD9"): dodoV2, // DODO V2 RouteProxy Avalanche
		strings.ToLower("0x2cD18557E14aF72DAA8090BcAA95b231ffC9ea26"): dodoV2, // DODO V2 RouteProxy Avalanche 2
		strings.ToLower("0x7950dc01542efe1c03aea610472e3b565b53f64a"): dodoV2, // DODO V2 RouteProxy Optimism
		strings.ToLower("0xfD9D2827AD469B72B69329dAA325ba7AfbDb3C98"): dodoV2, // DODO V2 RouteProxy Optimism 2
		// Velodrome
		// https://docs.velodrome.finance/security#contract-addresses
		strings.ToLower("0x9c12939390052919aF3155f41Bf4160Fd3666A6f"): velodromeSwap, // Velodrome Router
		// Curve QuickSwap
		strings.ToLower("0x55B916Ce078eA594c10a874ba67eCc3d62e29822"): curveQuickSwap, // Ethereum
		strings.ToLower("0xa522deb6F17853F3a97a65d0972a50bDC3B1AFFF"): curveQuickSwap, // Polygon
		strings.ToLower("0x16243caB3aC4d8eE8df7660a525F7F7539962468"): curveQuickSwap, // Fantom
		strings.ToLower("0x890f4e345B1dAED0367A877a1612f86A1f86985f"): curveQuickSwap, // Avalanche
		strings.ToLower("0xE6358f6a45B502477e83CC1CDa759f540E4459ee"): curveQuickSwap, // Gnosis
		// Balancer
		strings.ToLower(balancer.AddressVault.String()): balancerSwap,
		// Cow
		strings.ToLower("0x9008D19f58AAbD9eD0D60971565AA8510560ab41"): cowSwap,
		// MaskNetwork
		strings.ToLower(masknetwork.AddressHappyTokenPool.String()): maskNetwork,
		// Zerion
		strings.ToLower("0xd7F1Dd5D49206349CaE8b585fcB0Ce3D96f1696F"): zerion,
	}
)
