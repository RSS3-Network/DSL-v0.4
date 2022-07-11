package mirror

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type MirrorContent struct {
	Nft     json.RawMessage `json:"nft"`
	Wnft    json.RawMessage `json:"wnft"`
	Digest  string          `json:"digest"`
	Content struct {
		Body      string          `json:"body"`
		Title     string          `json:"title"`
		Timestamp decimal.Decimal `json:"timestamp"`
	} `json:"content"`
	Version        string          `json:"version"`
	Authorship     json.RawMessage `json:"authorship"`
	OriginalDigest string          `json:"originalDigest"`
}

type Mirror struct {
	ContentType           string        `json:"content_type"`
	Contributor           string        `json:"contributor"`
	ContentDigest         string        `json:"content_digest"`
	OriginalContentDigest string        `json:"original_content_digest,omitempty"`
	Content               MirrorContent `json:"content"`
}
