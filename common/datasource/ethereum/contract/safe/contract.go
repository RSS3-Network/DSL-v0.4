package safe

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// https://github.com/safe-global/safe-deployments/blob/main/src/assets/v1.3.0/gnosis_safe.json
	AddressGnosisSafe = common.HexToAddress("0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552")

	EventHashAddedOwner       = crypto.Keccak256Hash([]byte("AddedOwner(address)"))
	EventHashRemovedOwner     = crypto.Keccak256Hash([]byte("RemovedOwner(address)"))
	EventHashChangedThreshold = crypto.Keccak256Hash([]byte("ChangedThreshold(uint256)"))
	EventHashExecutionSuccess = crypto.Keccak256Hash([]byte("ExecutionSuccess(bytes32,uint256)"))
	EventHashExecutionFailure = crypto.Keccak256Hash([]byte("ExecutionFailure(bytes32,uint256)"))
	EventHashSafeSetup        = crypto.Keccak256Hash([]byte("SafeSetup(address,address[],uint256,address,address)"))
)
