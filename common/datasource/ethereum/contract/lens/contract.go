package lens

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./events.abi --pkg contract --type Events --out ./events.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./ILensHub.abi --pkg contract --type ILensHub --out ./iLensHub.go

var (
	HubProxyContractAddress     = common.HexToAddress("0xdb46d1dc155634fbc732f92e853b10b288ad5a1d")
	ProfileProxyContractAddress = common.HexToAddress("0x1eeC6ecCaA4625da3Fa6Cd6339DBcc2418710E8a")

	EventHashPostCreated          = crypto.Keccak256Hash([]byte("PostCreated(uint256,uint256,string,address,bytes,address,bytes,uint256)"))
	EventHashProfileCreated       = crypto.Keccak256Hash([]byte("ProfileCreated(uint256,address,address,string,string,address,bytes,string,uint256)"))
	EventHashCommentCreated       = crypto.Keccak256Hash([]byte("CommentCreated(uint256,uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,uint256)"))
	EventHashFollowNFTTransferred = crypto.Keccak256Hash([]byte("FollowNFTTransferred(uint256,uint256,address,address,uint256)"))
	EventHashMirrorCreated        = crypto.Keccak256Hash([]byte("MirrorCreated(uint256,uint256,uint256,uint256,bytes,address,bytes,uint256)"))
	EventHashFollowed             = crypto.Keccak256Hash([]byte("Followed(address,uint256[],bytes[],uint256)"))
)

var SupportLensEvents = map[common.Hash]common.Address{
	EventHashCommentCreated:       HubProxyContractAddress,
	EventHashPostCreated:          HubProxyContractAddress,
	EventHashProfileCreated:       HubProxyContractAddress,
	EventHashMirrorCreated:        HubProxyContractAddress,
	EventHashFollowed:             HubProxyContractAddress,
	EventHashFollowNFTTransferred: HubProxyContractAddress,
}

// note: kurora cannot query who the user has unfollowed
