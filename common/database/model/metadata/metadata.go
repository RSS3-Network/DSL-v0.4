package metadata

import (
	"encoding/json"
	"math/big"

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
	Token     *Token     `json:"token,omitempty"`
	Swap      *SwapPool  `json:"swap,omitempty"`
	Mirror    *Mirror    `json:"mirror,omitempty"`
	POAP      *POAP      `json:"poap,omitempty"`
	Gitcoin   *Gitcoin   `json:"gitcoin,omitempty"`
	SnapShot  *SnapShot  `json:"snapshot,omitempty"`
	Crossbell *Crossbell `json:"crossbell,omitempty"`
	Lens      *Lens      `json:"lens,omitempty"`
}

type Token struct {
	TokenAddress  string           `json:"token_address,omitempty"`
	TokenStandard string           `json:"token_standard"`
	TokenID       *decimal.Decimal `json:"token_id,omitempty"`
	TokenValue    *decimal.Decimal `json:"token_value,omitempty"`

	Logo     string `json:"logo,omitempty"`
	Decimals uint8  `json:"decimals,omitempty"`
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`

	NFTMetadata json.RawMessage `json:"nft_metadata,omitempty"`
}

type Mirror struct {
	ContentType           string          `json:"content_type"`
	Contributor           string          `json:"contributor"`
	ContentDigest         string          `json:"content_digest"`
	OriginalContentDigest string          `json:"original_content_digest,omitempty"`
	Content               json.RawMessage `json:"content"`
}

type SwapPool struct {
	Name     string `json:"name"`
	Token0   string `json:"token0"`
	Token1   string `json:"token1"`
	Network  string `json:"network"`
	Protocol string `json:"protocol"`
}

type Gitcoin struct {
	Id          int    `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type POAP struct {
	ID          int    `json:"id"`
	FancyID     string `json:"fancy_id"`
	Name        string `json:"name"`
	EventURL    string `json:"event_url"`
	ImageURL    string `json:"image_url"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	ExpiryDate  string `json:"expiry_date"`
	TokenID     string `json:"token_id"`
}

type SnapShot struct {
	Proposal json.RawMessage `json:"proposal"`
	Space    json.RawMessage `json:"space"`
	Choice   json.RawMessage `json:"choice"`
}

type Crossbell struct {
	TokenID       *big.Int        `json:"token_id,omitempty"`
	TokenIDFrom   *big.Int        `json:"token_id_from,omitempty"`
	TokenIDTo     *big.Int        `json:"token_id_to,omitempty"`
	LinkListID    *big.Int        `json:"linklist_id,omitempty"`
	ProfileID     *big.Int        `json:"profile_id,omitempty"`
	ProfileIDFrom *big.Int        `json:"profile_id_from,omitempty"`
	ProfileIDTo   *big.Int        `json:"profile_id_to,omitempty"`
	LinkType      string          `json:"link_type,omitempty"`
	URI           string          `json:"uri,omitempty"`
	Handle        string          `json:"handle,omitempty"`
	Metadata      json.RawMessage `json:"metadata,omitempty"`
	MetadataFrom  json.RawMessage `json:"metadata_from,omitempty"`
	MetadataTo    json.RawMessage `json:"metadata_to,omitempty"`
}

type Lens struct {
	Type    string          `json:"type"`
	Content string          `json:"content"`
	Media   json.RawMessage `json:"media"`
	Target  json.RawMessage `json:"target,omitempty"`
}
