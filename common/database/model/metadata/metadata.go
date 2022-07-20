package metadata

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
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
	Token     *Token         `json:"transaction,omitempty"`
	Swap      *SwapPool      `json:"swap,omitempty"`
	POAP      *POAP          `json:"poap,omitempty"`
	Donation  *Donation      `json:"donation,omitempty"`
	SnapShot  *SnapShot      `json:"snapshot,omitempty"`
	Crossbell *Crossbell     `json:"crossbell,omitempty"`
	Post      *Post          `json:"content,omitempty"`
	Profile   *model.Profile `json:"profile,omitempty"`
	Vote      *Vote          `json:"vote,omitempty"`
	Proposal  *Proposal      `json:"proposal,omitempty"`
}

type Token struct {
	Name            string            `json:"name"`
	Symbol          string            `json:"symbol"`
	Decimals        uint8             `json:"decimals,omitempty"`
	Standard        string            `json:"standard"`
	ContractAddress string            `json:"contract_address,omitempty"`
	Image           string            `json:"image,omitempty"`
	ID              *big.Int          `json:"id,omitempty"`
	Value           *decimal.Decimal  `json:"value,omitempty"`
	Description     string            `json:"description,omitempty"`
	Attributes      map[string]string `json:"attributes,omitempty"`
	ExternalLink    string            `json:"external_link,omitempty"`
	AnimationURL    string            `json:"animation_url,omitempty"`
	Metadata        json.RawMessage   `json:"metadata,omitempty"`
}

type Swap struct {
	Protocol  string `json:"protocol"`
	TokenFrom Token  `json:"token_from"`
	TokenTo   Token  `json:"token_to"`
}

type SwapToken struct {
	Address  string          `json:"address"`
	Symbol   string          `json:"symbol"`
	Decimals uint8           `json:"decimals"`
	Value    decimal.Decimal `json:"value"`
}

type SwapPool struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Token0   string `json:"token0,omitempty"`
	Token1   string `json:"token1,omitempty"`
	Network  string `json:"network"`
	Protocol string `json:"protocol"`
}

type Donation struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Platform    string `json:"platform"`
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
	case *POAP:
		internalMetadataModel.POAP = model
	case *Donation:
		internalMetadataModel.Donation = model
	case *Crossbell:
		internalMetadataModel.Crossbell = model
	// social
	case *Post:
		internalMetadataModel.Post = model
	case *model.Profile:
		internalMetadataModel.Profile = model
	// governance
	case *Vote:
		internalMetadataModel.Vote = model
	case *Proposal:
		internalMetadataModel.Proposal = model
	default:
		return nil, ErrorUnsupportedType
	}

	return json.Marshal(internalMetadataModel)
}
