package model

type WrappedResult struct {
	Social SocialResult `json:"social"`
	Search SearchResult `json:"search"`
}

type SocialResult struct {
	Post         int64  `json:"post"`
	Comment      int64  `json:"comment"`
	Follow       int64  `json:"follow"`
	LongestHash  string `json:"longest_hash"`
	ShortestHash string `json:"shortest_hash"`
}

type SearchResult struct {
	Count int64 `json:"count"`
}
