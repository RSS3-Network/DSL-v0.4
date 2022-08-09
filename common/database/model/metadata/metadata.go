package metadata

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var Default json.RawMessage

var ErrorUnsupportedType = errors.New("unsupported metadata type")

func init() {
	if err := json.Unmarshal([]byte("{}"), &Default); err != nil {
		logrus.Fatalln(err)
	}
}

type Token struct {
	Name            string           `json:"name"`
	Collection      string           `json:"collection,omitempty"` // ERC-1155
	Symbol          string           `json:"symbol"`
	Decimals        uint8            `json:"decimals,omitempty"`
	Standard        string           `json:"standard"`
	ContractAddress string           `json:"contract_address,omitempty"`
	Image           string           `json:"image,omitempty"`
	ID              string           `json:"id,omitempty"`
	Value           *decimal.Decimal `json:"value,omitempty"`
	Cost            *Token           `json:"cost,omitempty"` // TODO Differentiate between UMS
	Description     string           `json:"description,omitempty"`
	Attributes      []TokenAttribute `json:"attributes,omitempty"`
	ExternalLink    string           `json:"external_link,omitempty"`
	ExternalURL     string           `json:"external_url,omitempty"`
	AnimationURL    string           `json:"animation_url,omitempty"`
}

type TokenAttribute struct {
	TraitType string `json:"trait_type"`
	Value     any    `json:"value"`
}

type Swap struct {
	Protocol  string `json:"protocol"`
	TokenFrom Token  `json:"from"`
	TokenTo   Token  `json:"to"`
}

type Donation struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Platform    string `json:"platform"`
	Token       Token  `json:"token"`
}

type POAP struct {
	ID          int    `json:"id"`
	FancyID     string `json:"fancy_id,omitempty"`
	Name        string `json:"name"`
	EventURL    string `json:"event_url,omitempty"`
	ImageURL    string `json:"image_url"`
	Country     string `json:"country,omitempty"`
	City        string `json:"city,omitempty"`
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
	Event     string              `json:"event"`
	Link      *CrossbellLink      `json:"link,omitempty"`
	Character *CrossbellCharacter `json:"character,omitempty"`
	Note      *CrossbellNote      `json:"note,omitempty"`
}

type CrossbellLink struct {
	Type string `json:"type"`
	From any    `json:"from"`
	To   any    `json:"to"`
}

type CrossbellCharacter struct {
	ID       *big.Int        `json:"id"`
	URI      string          `json:"uri,omitempty"`
	Metadata json.RawMessage `json:"metadata"`
}

type CrossbellNote struct {
	ID           *big.Int        `json:"id"`
	LinkItemType common.Hash     `json:"link_item_type"`
	LinkKey      common.Hash     `json:"link_key"`
	Link         any             `json:"link,omitempty"`
	ContentURI   string          `json:"content_uri"`
	LinkModule   common.Address  `json:"link_module"`
	MintModule   common.Address  `json:"mint_module"`
	MintNFT      common.Address  `json:"mint_nft"`
	Deleted      bool            `json:"deleted"`
	Locked       bool            `json:"locked"`
	Metadata     json.RawMessage `json:"metadata"`
}
