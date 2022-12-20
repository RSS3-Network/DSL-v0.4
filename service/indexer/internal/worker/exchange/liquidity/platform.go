package liquidity

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/balancer"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lido"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Platform struct {
	Name     string
	Protocol string
}

var (
	platformUniswapV3 = Platform{
		Name:     protocol.PlatformUniswap,
		Protocol: uniswap.ProtocolUniswapV3,
	}

	platformPancake = Platform{
		Name:     protocol.PlatformPancakeswap,
		Protocol: uniswap.ProtocolUniswapV2,
	}

	platformSushi = Platform{
		Name:     protocol.PlatformSushiswap,
		Protocol: uniswap.ProtocolUniswapV2,
	}

	platformAAVEV2 = Platform{
		Name:     protocol.PlatformAAVE,
		Protocol: "AAVE V2",
	}

	platformAAVEV3 = Platform{
		Name:     protocol.PlatformAAVE,
		Protocol: "AAVE V3",
	}

	platformCurve = Platform{
		Name:     protocol.PlatformCurve,
		Protocol: "Curve",
	}

	platformLido = Platform{
		Name:     protocol.PlatformLido,
		Protocol: "Lido",
	}

	platformPolygonStaking = Platform{
		Name:     protocol.PlatformPolygonStaking,
		Protocol: "Polygon Staking",
	}

	platformBalancer = Platform{
		Name:     protocol.PlatformBalancer,
		Protocol: "Balancer",
	}

	platformMap = map[common.Address]Platform{
		// Uniswap V3
		common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88"): platformUniswapV3, // Uniswap V3 Positions NFT
		// PancakeSwap
		common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E"): platformPancake, // PancakeSwap
		// SushiSwap
		common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"): platformSushi, // SushiSwap
		// AAVE
		common.HexToAddress("0xcc9a0B7c43DC2a5F023Bb9b738E45B0Ef6B06E04"): platformAAVEV2, // AAVE WETH Gateway V2 Ethereum
		common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"): platformAAVEV2, // AAVE Lending Pool V2
		common.HexToAddress("0x9BdB5fcc80A49640c7872ac089Cc0e00A98451B6"): platformAAVEV3, // AAVE WETH Gateway V3 Polygon
		common.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"): platformAAVEV3, // AAVE Pool V3
		// Curve
		common.HexToAddress("0x1d8b86e3D88cDb2d34688e87E72F388Cb541B7C8"): platformCurve, // Curve Polygon ATriCrypto3 Zap
		// Lido
		lido.AddressETH:   platformLido, // Lido stETH Proxy
		lido.AddressMatic: platformLido, // Lido stMATIC Proxy
		// Balancer
		balancer.AddressVault: platformBalancer, // Balancer Vault
	}
)
