package metadata

import (
	"math/big"
)

type MultiSig struct {
	Action    string   `json:"action"`              // add_owner remove_owner / change_threshold / execution reject
	Success   bool     `json:"success,omitempty"`   // execution or rejection
	Owner     *string  `json:"owner,omitempty"`     // add_owner or remove_owner
	Threshold *big.Int `json:"threshold,omitempty"` // change_threshold
	Vault     Vault    `json:"vault"`
}

type Vault struct {
	Address   string   `json:"address"`
	Owners    []string `json:"owners"`
	Threshold *big.Int `json:"threshold"`
	Version   string   `json:"version"`
}
