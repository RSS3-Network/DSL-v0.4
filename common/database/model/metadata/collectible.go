package metadata

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type Party struct {
	PartyAddress  string           `json:"party_address"`
	Name          string           `json:"party_name"`
	Symbol        string           `json:"symbol"`
	PartyType     string           `json:"party_type"`
	Creator       string           `json:"creator,omitempty"`
	NftContract   string           `json:"nft_contract_address,omitempty"`
	TokenId       *big.Int         `json:"nft_token_id,omitempty"`
	MarketWrapper string           `json:"market_wrapper"`
	AuctionId     *big.Int         `json:"auction_id,omitempty"`
	ExpiredTime   *time.Time       `json:"expired_date,omitempty"`
	Deciders      []common.Address `json:"deciders,omitempty"`
	PartyStatus   uint8            `json:"party_status,omitempty"`
	Action        string           `json:"action,omitempty"`
}

type PartyContribute struct {
	PartyInfo                       Party           `json:"party_info"`
	Action                          string          `json:"action"`
	Contributor                     string          `json:"contributor"`
	Amount                          decimal.Decimal `json:"amount"`
	PreviousTotalContributedToParty decimal.Decimal `json:"previous_total_contributed_to_party"`
	TotalFromContributor            decimal.Decimal `json:"total_from_contributor"`
}

type PartyClaim struct {
	PartyInfo          Party           `json:"party_info"`
	Action             string          `json:"action"`
	Contributor        string          `json:"contributor"`
	TotalContributed   decimal.Decimal `json:"total_contributed"`
	ExcessContribution decimal.Decimal `json:"excess_contribution"`
	TokenAmount        decimal.Decimal `json:"fractional_token_amount"`
}

type PartyFinalize struct {
	PartyInfo        Party           `json:"party_info"`
	Action           string          `json:"action"`
	Result           uint8           `json:"result"`
	TotalSpent       decimal.Decimal `json:"total_spent"`
	Fee              decimal.Decimal `json:"fee"`
	TotalContributed decimal.Decimal `json:"total_contributed"`
}

type PartyBid struct {
	PartyInfo Party           `json:"party_info"`
	Action    string          `json:"action"`
	BidAmount decimal.Decimal `json:"bid_amount"`
}

type PartyExpire struct {
	PartyInfo   Party  `json:"party_info"`
	Action      string `json:"action"`
	TriggeredBy string `json:"trigger"`
}

type PartyBuy struct {
	PartyInfo        Party           `json:"party_info"`
	Action           string          `json:"action"`
	TokenId          *big.Int        `json:"token_id,omitempty"`
	TriggeredBy      string          `json:"trigger"`
	TargetAddress    string          `json:"target_address"`
	EthSpent         decimal.Decimal `json:"eth_spent"`
	EthFeePaid       decimal.Decimal `json:"eth_fee_paid"`
	TotalContributed decimal.Decimal `json:"total_contributed"`
}
