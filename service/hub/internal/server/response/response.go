package response

import (
	"encoding/json"
)

type Response struct {
	Total  int64 `json:"total"`
	Result any   `json:"result"`
}

type Feed struct {
	Tag      string   `json:"tag"`
	Platform string   `json:"platform,omitempty"`
	Network  string   `json:"network"`
	Proof    string   `json:"proof"`
	Actions  []Action `json:"actions"`
}

type Action struct {
	Type     string          `json:"type"`
	From     string          `json:"from"`
	To       string          `json:"to"`
	Metadata json.RawMessage `json:"metadata"`
}
