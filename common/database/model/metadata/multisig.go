package metadata

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type MultiSig struct {
	Action    string          `json:"action"`              // add_owner remove_owner / change_threshold / execution reject
	Success   bool            `json:"success,omitempty"`   // execution or rejection
	Owner     *common.Address `json:"owner,omitempty"`     // add_owner or remove_owner
	Threshold *big.Int        `json:"threshold,omitempty"` // change_threshold
	Vault     Vault           `json:"vault"`
}

type Vault struct {
	Address   common.Address   `json:"address"`
	Owners    []common.Address `json:"owners"`
	Threshold *big.Int         `json:"threshold"`
	Version   string           `json:"version"`
}
