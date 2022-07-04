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

const (
	SwapTypeRouter = "router"
	SwapTypePool   = "pool"
)

type SwapPool struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Token0   string `json:"token0,omitempty"`
	Token1   string `json:"token1,omitempty"`
	Network  string `json:"network"`
	Protocol string `json:"protocol"`
}

type Gitcoin struct {
	ID          int    `json:"id"`
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

type Lens struct {
	Type    string          `json:"type"`
	Content string          `json:"content"`
	Media   json.RawMessage `json:"media"`
	Target  json.RawMessage `json:"target,omitempty"`
}

func BuildMetadataRawMessage(metadataRawMessage json.RawMessage, metadataModel any) (json.RawMessage, error) {
	var internalMetadataModel Metadata

	if err := json.Unmarshal(metadataRawMessage, &internalMetadataModel); err != nil {
		return nil, err
	}

	switch model := metadataModel.(type) {
	case *Token:
		internalMetadataModel.Token = model
	case *SwapPool:
		internalMetadataModel.Swap = model
	case *Mirror:
		internalMetadataModel.Mirror = model
	case *POAP:
		internalMetadataModel.POAP = model
	case *Gitcoin:
		internalMetadataModel.Gitcoin = model
	case *Crossbell:
		internalMetadataModel.Crossbell = model
	case *Lens:
		internalMetadataModel.Lens = model
	default:
		return nil, ErrorUnsupportedType
	}

	return json.Marshal(internalMetadataModel)
}
