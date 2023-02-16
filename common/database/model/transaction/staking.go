package transaction

import "github.com/naturalselectionlabs/pregod/common/database/model/metadata"

type Staking struct {
	Action string         `json:"action"`
	Token  metadata.Token `json:"token"`
	Period *Period        `json:"period,omitempty"`
}

type Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
