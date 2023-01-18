package model

import (
	"time"

	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
)

const (
	DefaultLimit = 500

	// path
	GetNotes        = "/notes/"
	PostNotes       = "/notes"
	PostSocialNotes = "/notes/social"
	GetProfiles     = "/profiles/"
	PostProfiles    = "/profiles"
	GetNS           = "/ns/"

	EsIndex = "pregod-v1-visit-path"

	Subscribe   = "subscribe"
	Unsubscribe = "unsubscribe"
	Query       = "query"
)

type Response struct {
	Total         *int64            `json:"total,omitempty"`
	Cursor        string            `json:"cursor,omitempty"`
	Result        any               `json:"result,omitempty"`
	AddressStatus []dbModel.Address `json:"address_status,omitempty"`
	Message       string            `json:"message,omitempty"`
}

type GetRequest struct {
	Address   string    `param:"address" json:"address" validate:"required"`
	Limit     int       `query:"limit" json:"limit"`
	Cursor    string    `query:"cursor" json:"cursor"`
	Type      []string  `query:"type" json:"type"`
	Tag       []string  `query:"tag" json:"tag" validate:"required_with=Type"`
	Network   []string  `query:"network" json:"network"`
	Platform  []string  `query:"platform" json:"platform"`
	Timestamp time.Time `query:"timestamp" json:"timestamp"`
	Hash      string    `query:"hash" json:"hash"`
	// includes POAP in the response
	IncludePoap bool `query:"include_poap" json:"include_poap"`
	Refresh     bool `query:"refresh" json:"refresh"`
	Reindex     bool `query:"reindex" json:"reindex"`
	Page        int  `query:"page" json:"page"`
	QueryStatus bool `query:"query_status" json:"query_status"`
	// returns a count of transactions only
	CountOnly bool `query:"count_only" json:"count_only"`
}

type GetAssetRequest struct {
	Address      string   `param:"address" validate:"required"`
	Network      []string `query:"network"`
	TokenAddress string   `query:"token_address" validate:"required_with=TokenId"`
	TokenId      string   `query:"token_id"`
	Cursor       string   `query:"cursor"`
	Limit        int      `query:"limit"`
	Refresh      bool     `query:"refresh"`
	BlockSpam    *bool    `query:"block_spam"` // Default true
}

type BatchGetNotesRequest struct {
	Address        []string  `json:"address" validate:"required"`
	Type           []string  `json:"type"`
	Tag            []string  `json:"tag" validate:"required_with=Type"`
	Network        []string  `json:"network"`
	Platform       []string  `json:"platform"`
	Timestamp      time.Time `json:"timestamp"`
	Limit          int       `json:"limit"`
	Cursor         string    `json:"cursor"`
	Refresh        bool      `json:"refresh"`
	IncludePoap    bool      `json:"include_poap"`
	Page           int       `json:"page"`
	QueryStatus    bool      `json:"query_status"`
	CountOnly      bool      `json:"count_only"`
	IgnoreContract bool      `json:"ignore_contract"`
}

type BatchGetSocialNotesRequest struct {
	Address     []string  `json:"address" validate:"required"`
	Type        []string  `json:"type"`
	Network     []string  `json:"network"`
	Platform    []string  `json:"platform"`
	Timestamp   time.Time `json:"timestamp"`
	Limit       int       `json:"limit"`
	Cursor      string    `json:"cursor"`
	Page        int       `json:"page"`
	QueryStatus bool      `json:"query_status"`
	CountOnly   bool      `json:"count_only"`
}

type BatchGetProfilesRequest struct {
	Address  []string `json:"address" validate:"required"`
	Network  []string `json:"network"`
	Platform []string `json:"platform"`
	Refresh  bool     `query:"refresh"`
}

type APIKeyRequest struct {
	Address string `json:"address" validate:"required"`
}

// exchange
type CexResult struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Network string `json:"network"`
}

type DexResult struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Network string `json:"network"`
	Pair    string `json:"pair"`
}

type GetExchangeRequest struct {
	ExchangeType string   `param:"exchange_type"`
	Cursor       int      `query:"cursor"`
	Name         []string `query:"name"`
	Network      []string `query:"network"`
}

// platform
type GetPlatformRequest struct {
	PlatformType string   `param:"platform_type"`
	Network      []string `query:"network"`
}

type PlatformResult struct {
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Type    string `json:"type,omitempty"`
	Network string `json:"network,omitempty"`
}

type WebSocketRequest struct {
	Id         *int     `json:"id,omitempty"`
	Action     string   `json:"action"`
	AddressArr []string `json:"address"`
	ClientId   string   `json:"client_id"`
}

type WebsocketResponse struct {
	Id     int            `json:"id"`
	Status string         `json:"status"`
	Result map[string]any `json:"result"`
}

type Transactions []dbModel.Transaction

// Len()
func (t Transactions) Len() int {
	return len(t)
}

// Less()
func (t Transactions) Less(i, j int) bool {
	return t[i].Timestamp.Unix() > t[j].Timestamp.Unix() ||
		(t[i].Timestamp.Unix() == t[j].Timestamp.Unix() && t[i].Index > t[j].Index)
}

// Swap()
func (t Transactions) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
