package ens

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	EnsRegistrarController = common.HexToAddress("0x283af0b28c62c092c9727f1ee09c02ca627eb7f5")
	EnsBulkRenewal         = common.HexToAddress("0xff252725f6122a92551a5fa9a6b6bf10eb0be035")
	EnsBatchRenew          = common.HexToAddress("0xd24d9a7B765c2946bc6fBB84Df6319A31F023a1D")
	EventNameRenewed       = common.BytesToHash(crypto.Keccak256([]byte("NameRenewed(string,bytes32,uint256,uint256)")))

	EnsPublicResolver = common.HexToAddress("0x231b0ee14048e9dccd1d247744d114a4eb5e8e63")
	EventTextChanged  = common.BytesToHash(crypto.Keccak256([]byte("TextChanged(bytes32,string,string,string)")))

	EnsPublicResolver2 = common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41")
)
