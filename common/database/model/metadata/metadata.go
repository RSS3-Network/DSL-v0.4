package metadata

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"math/big"
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
	CrossBell *CrossBell `json:"crossbell,omitempty"`
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

type CrossBell struct {
	Event          string          `json:"event"`
	ProfileID      *big.Int        `json:"profile_id"`
	Handle         string          `json:"handle,omitempty"`
	AddressCreator string          `json:"address_creator,omitempty"`
	AddressOwner   string          `json:"address_owner,omitempty"`
	Metadata       json.RawMessage `json:"metadata,omitempty"`
}
