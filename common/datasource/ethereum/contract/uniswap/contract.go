package uniswap

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	ProtocolUniswapV2 = "Uniswap V2"
	ProtocolUniswapV3 = "Uniswap V3"
)

var (
	// https://github.com/Uniswap/universal-router/tree/main/deploy-addresses
	AddressUniversalRouter         = common.HexToAddress("0xEf1c6E67703c7BD7107eed8303Fbe6EC2554BF6B")
	AddressUniversalRouterArbitrum = common.HexToAddress("0x4C60051384bd2d3C01bfc845Cf5F4b44bcbE9de5")
	AddressUniversalRouterCelo     = common.HexToAddress("0xC73d61d192FB994157168Fb56730FdEc64C9Cb8F")
	AddressUniversalRouterOptimism = common.HexToAddress("0xb555edF5dcF85f42cEeF1f3630a52A108E55A654")
	AddressUniversalRouterPolygon  = common.HexToAddress("0x4C60051384bd2d3C01bfc845Cf5F4b44bcbE9de5")

	EventHashTransfer            = common.BytesToHash(crypto.Keccak256([]byte("Transfer(address,address,uint256)")))
	EventHashSwapV2              = common.BytesToHash(crypto.Keccak256([]byte("Swap(address,uint256,uint256,uint256,uint256,address)")))
	EventHashSwapV3              = common.BytesToHash(crypto.Keccak256([]byte("Swap(address,address,int256,int256,uint160,uint128,int24)")))
	EventHashApproval            = common.BytesToHash(crypto.Keccak256([]byte("Approval(address,address,uint256)")))
	EventHashMintV2              = common.BytesToHash(crypto.Keccak256([]byte("Mint(address,uint256,uint256)")))
	EventHashBurnV2              = common.BytesToHash(crypto.Keccak256([]byte("Burn(address,uint256,uint256,address)")))
	EventHashMintV3              = common.BytesToHash(crypto.Keccak256([]byte("Mint(address,address,int24,int24,uint128,uint256,uint256)")))
	EventHashBurnV3              = common.BytesToHash(crypto.Keccak256([]byte("Burn(address,int24,int24,uint128,uint256,uint256)")))
	EventHashIncreaseLiquidityV3 = common.BytesToHash(crypto.Keccak256([]byte("IncreaseLiquidity(uint256,uint128,uint256,uint256)")))
	EventHashDecreaseLiquidityV3 = common.BytesToHash(crypto.Keccak256([]byte("DecreaseLiquidity(uint256,uint128,uint256,uint256)")))
	EventHashCollectV3           = common.BytesToHash(crypto.Keccak256([]byte("Collect(uint256,address,uint256,uint256)")))
)
