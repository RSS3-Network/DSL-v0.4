package response

import (
	"encoding/json"
	"sort"
	"time"
)

type Response struct {
	Total  int64 `json:"total"`
	Result any   `json:"result"`
}

var _ sort.Interface = &Feeds{}

type Feeds []Feed

func (f Feeds) Len() int {
	return len(f)
}

func (f Feeds) Less(i, j int) bool {
	return f[i].Timestamp.After(f[j].Timestamp)
}

func (f Feeds) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type Feed struct {
	Timestamp time.Time `json:"timestamp"`
	Tag       string    `json:"tag"`
	Platform  string    `json:"platform,omitempty"`
	Network   string    `json:"network"`
	Proof     string    `json:"proof"`
	Actions   []Action  `json:"actions"`
}

type Action struct {
	Type     string          `json:"type"`
	From     string          `json:"from"`
	To       string          `json:"to"`
	Metadata json.RawMessage `json:"metadata"`
}
