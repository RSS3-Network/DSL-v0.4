package uniswap

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
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
