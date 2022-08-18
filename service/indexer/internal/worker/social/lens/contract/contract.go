package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate abigen --abi ./events.abi --pkg contract --type Events --out ./events.go
//go:generate abigen --abi ./ILensHub.abi --pkg contract --type ILensHub --out ./iLensHub.go

var (
	ContractAddress = common.HexToAddress("0xdb46d1dc155634fbc732f92e853b10b288ad5a1d")

	EventHashPostCreated    = crypto.Keccak256Hash([]byte("PostCreated(uint256,uint256,string,address,bytes,address,bytes,uint256)"))
	EventHashProfileCreated = crypto.Keccak256Hash([]byte("ProfileCreated(uint256,address,address,string,string,address,bytes,string,uint256)"))
	EventHashCommentCreated = crypto.Keccak256Hash([]byte("CommentCreated(uint256,uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,uint256)"))
	EventHashFollow         = crypto.Keccak256Hash([]byte("FollowNFTTransferred(uint256,uint256,address,address,uint256)"))
)