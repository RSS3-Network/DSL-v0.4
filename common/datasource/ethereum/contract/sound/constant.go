package sound

import (
	"github.com/ethereum/go-ethereum/common"
)

//go:generate abigen --abi ./contract/artist.abi --pkg contract --type Artist --out ./contract/artist.go
//go:generate abigen --abi ./contract/merkleDropMinter.abi --pkg contract --type MerkleDropMinter --out ./contract/merkleDropMinter.go

//https://etherscan.io/address/0xbaa4b1c49b84876332ddafc9d0cae5bcd439e12c

var (
	EventEditionPurchased = common.HexToHash("0xc3e4ad784e3889561b308df24921cf08a47410a8ed63b8afe80c50a002efb251")
	EventEditionCreated   = common.HexToHash("0xb56f9ba6a8af17a142f8ad372c6fc283e81d8c6193b86afe7f6ca901acd698f3")
	EventMint             = common.HexToHash("0x3d73a0206d94d61b7038abcd0eb766a5de22f9844b38a78449054d19a4f1b58a")

	MintContractAddress = common.HexToAddress("0xeae422887230C0FFB91Fd8F708f5fdD354C92F2F")
)
