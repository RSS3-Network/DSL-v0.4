package model

import (
	"encoding/json"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
)

type WrappedResult struct {
	Social      SocialResult `json:"social"`
	Search      SearchResult `json:"search"`
	Gas         GasResult    `json:"gas"`
	Transaction TxResult     `json:"transaction"`
	NFT         NFTResult    `json:"nft"`
	DApp        DAppResult   `json:"dapp"`
	DeFi        DeFiResult   `json:"defi"`
}

type SocialResult struct {
	Post         int64  `json:"post"`
	Comment      int64  `json:"comment"`
	Follow       int64  `json:"follow"`
	LongestHash  string `json:"longest_hash"`
	ShortestHash string `json:"shortest_hash"`
	List         []DApp `json:"list" gorm:"-"`
}

type SearchResult struct {
	Count int64 `json:"count"`
}

type GasResult struct {
	Total       string `json:"total"`
	Highest     string `json:"highest"`
	HighestHash string `json:"highest_hash"`
}

type TxResult struct {
	Initiate []NetworkCount `json:"initiated"`
	Receive  []NetworkCount `json:"received"`
}

type NetworkCount struct {
	Network string `json:"network"`
	Total   int64  `json:"total"`
}

type NFTResult struct {
	Bought []metadata.Token `json:"bought"`
	Sold   []metadata.Token `json:"sold"`
	Total  int              `json:"total"`
	First  *NFTSingle       `json:"first"`
	Last   *NFTSingle       `json:"last"`
}

type NFT struct {
	Metadata  json.RawMessage `json:"metadata"`
	From      string          `json:"from"`
	To        string          `json:"to"`
	Timestamp time.Time       `json:"timestamp"`
}

type NFTSingle struct {
	Metadata  metadata.Token `json:"metadata"`
	Timestamp time.Time      `json:"timestamp"`
}

type DAppResult struct {
	List []DApp `json:"list"`
}

type DApp struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type DeFiResult struct {
	List      []DApp           `json:"list"`
	SwapPair  []SwapPair       `json:"swap_pair"`
	Bridge    []metadata.Token `json:"bridge"`
	Liquidity []metadata.Token `json:"liquidity"`
}

type SwapPair struct {
	From string `json:"from"`
	To   string `json:"to"`
}
