package balancer

import "github.com/ethereum/go-ethereum/crypto"

var EventHashSwap = crypto.Keccak256Hash([]byte("Swap(bytes32,address,address,uint256,uint256)"))
