package erc1155

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	EventHashTransferSingle = common.BytesToHash(crypto.Keccak256([]byte("TransferSingle(address,address,address,uint256,uint256)")))
	EventHashTransferBatch  = common.BytesToHash(crypto.Keccak256([]byte("TransferBatch(address,address,address,uint256[],uint256[])")))
	EventHashApprovalForAll = common.BytesToHash(crypto.Keccak256([]byte("ApprovalForAll(address,address,bool)")))
	EventHashURI            = common.BytesToHash(crypto.Keccak256([]byte("URI(string,uint256)")))
)
