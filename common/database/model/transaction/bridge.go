package transaction

import "github.com/naturalselectionlabs/pregod/common/database/model/metadata"

type Bridge struct {
	Action        string         `json:"action"`
	TargetNetwork TargetNetwork  `json:"target_network"`
	Token         metadata.Token `json:"token"`
}

type TargetNetwork struct {
	Name    string `json:"name"`
	ChainID uint64 `json:"chain_id"`
	Symbol  string `json:"symbol"`
}
