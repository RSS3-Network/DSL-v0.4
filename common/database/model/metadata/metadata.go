package metadata

import (
	"encoding/json"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var Default json.RawMessage

func init() {
	if err := json.Unmarshal([]byte("{}"), &Default); err != nil {
		logrus.Fatalln(err)
	}
}

type Metadata struct {
	Token *Token `json:"token,omitempty"`
	Swap  *Swap  `json:"swap,omitempty"`
}

type Token struct {
	TokenAddress  string          `json:"token_address"`
	TokenStandard string          `json:"token_standard"`
	TokenID       decimal.Decimal `json:"token_id,omitempty"`
	TokenValue    decimal.Decimal `json:"token_value"`
}

type Swap struct {
	Name string `json:"name"`
	Pool string `json:"pool"`
}
