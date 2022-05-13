package response

import "github.com/shopspring/decimal"

type Response struct {
	Total  int64 `json:"total"`
	Result any   `json:"result"`
}

type Feed struct {
	Tags     []string `json:"tags"`
	Platform string   `json:"platform,omitempty"`
	Network  string   `json:"network"`
	Proof    string   `json:"proof"`
	Actions  []Action `json:"actions"`
}

type Action struct {
	Type          string          `json:"type"`
	From          string          `json:"from"`
	To            string          `json:"to"`
	Token         string          `json:"token"`
	TokenAddress  string          `json:"token_address"`
	TokenStandard string          `json:"token_standard"`
	TokenID       decimal.Decimal `json:"token_id"`
	TokenValue    decimal.Decimal `json:"token_value"`
}
