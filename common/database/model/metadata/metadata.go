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
	Token  *Token  `json:"token,omitempty"`
	Swap   *Swap   `json:"swap,omitempty"`
	Mirror *Mirror `json:"mirror,omitempty"`
}

type Token struct {
	TokenAddress  string           `json:"token_address,omitempty"`
	TokenStandard string           `json:"token_standard"`
	TokenID       *decimal.Decimal `json:"token_id,omitempty"`
	TokenValue    *decimal.Decimal `json:"token_value"`

	Logo     string           `json:"logo"`
	Decimals uint8            `json:"decimals"`
	Name     string           `json:"name"`
	Symbol   string           `json:"symbol"`
	Supply   *decimal.Decimal `json:"suuply"`
}

type Mirror struct {
	ContentType           string          `json:"content_type"`
	Contributor           string          `json:"contributor"`
	ContentDigest         string          `json:"content_digest"`
	OriginalContentDigest string          `json:"original_content_digest,omitempty"`
	Content               json.RawMessage `json:"content"`
}

type Swap struct {
	Name string `json:"name"`
	Pool string `json:"pool"`
}
