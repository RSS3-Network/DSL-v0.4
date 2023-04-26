package ens

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ENSContractAddress = common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85")

	EnsRegistrarController = common.HexToAddress("0x283af0b28c62c092c9727f1ee09c02ca627eb7f5")
	EnsBulkRenewal         = common.HexToAddress("0xff252725f6122a92551a5fa9a6b6bf10eb0be035")
	EnsBatchRenew          = common.HexToAddress("0xd24d9a7B765c2946bc6fBB84Df6319A31F023a1D")
	EnsBulkRegistrarV1     = common.HexToAddress("0x1d6552e8F46fD509f3918A174FE62C34b42564aE")
	EventNameRenewed       = common.BytesToHash(crypto.Keccak256([]byte("NameRenewed(string,bytes32,uint256,uint256)")))
	EventNameRegistered    = common.BytesToHash(crypto.Keccak256([]byte("NameRegistered(string,bytes32,address,uint256,uint256)")))

	EnsRegistrarControllerV2 = common.HexToAddress("0x253553366Da8546fC250F225fe3d25d0C782303b")
	EventNameRegisteredV2    = common.BytesToHash(crypto.Keccak256([]byte("NameRegistered(string,bytes32,address,uint256,uint256,uint256)")))

	EnsPublicResolver = common.HexToAddress("0x231b0ee14048e9dccd1d247744d114a4eb5e8e63")
	EventTextChanged  = common.BytesToHash(crypto.Keccak256([]byte("TextChanged(bytes32,string,string,string)")))

	EnsPublicResolver2 = common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41")
)
