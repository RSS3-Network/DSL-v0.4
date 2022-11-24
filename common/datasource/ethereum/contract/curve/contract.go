package curve

import "github.com/ethereum/go-ethereum/crypto"

var (
	EventHashAddLiquidity       = crypto.Keccak256Hash([]byte("AddLiquidity(address,uint256[3],uint256[3],uint256,uint256)"))
	EventHashRemoveLiquidityOne = crypto.Keccak256Hash([]byte("RemoveLiquidityOne(address,uint256,uint256)"))
	EventHashTokenExchange      = crypto.Keccak256Hash([]byte("TokenExchange(address,int128,uint256,int128,uint256)"))
	EventHashTokenExchange2     = crypto.Keccak256Hash([]byte("TokenExchange(address,uint256,uint256,uint256,uint256)"))
)
