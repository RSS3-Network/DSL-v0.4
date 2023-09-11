package friendtech

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Friend Tech https://basescan.org/address/0xcf205808ed36593aa40a44f10c7f7c2f67d4a4d4
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./FriendTech.abi --pkg friendtech --type FriendTech --out friend_tech.go

var (
	AddressFriendTech = common.HexToAddress("0xCF205808Ed36593aa40a44F10c7f7C2F67d4A4d4")

	EventTrade = crypto.Keccak256Hash([]byte("Trade(address,address,bool,uint256,uint256,uint256,uint256,uint256)"))
)
