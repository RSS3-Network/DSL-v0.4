package metadata

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Party struct {
	PartyAddress  common.Address   `json:"party_address"`
	Name          string           `json:"name"`
	Symbol        string           `json:"symbol"`
	PartyType     string           `json:"party_type"`
	Creator       common.Address   `json:"creator"`
	NftContract   common.Address   `json:"nft_contract,omitempty"`
	TokenId       *big.Int         `json:"token_id,omitempty"`
	MarketWrapper string           `json:"market_wrapper"`
	AuctionId     *big.Int         `json:"auction_id,omitempty"`
	MaxPrice      *big.Int         `json:"max_price,omitempty"`
	ExpiredTime   *big.Int         `json:"expired_time,omitempty"`
	Deciders      []common.Address `json:"deciders,omitempty"`
	PartyStatus   uint8            `json:"party_status,omitempty"`
	Action        string           `json:"action"`
}

type PartyContribute struct {
	PartyInfo                       Party    `json:"party_info"`
	Contributor                     string   `json:"contributor"`
	Amount                          *big.Int `json:"amount"`
	PreviousTotalContributedToParty *big.Int `json:"previous_total_contributed_to_party"`
	TotalFromContributor            *big.Int `json:"total_from_contributor"`
}

type PartyClaim struct {
	PartyInfo          Party          `json:"party_info"`
	Contributor        common.Address `json:"contributor"`
	TotalContributed   *big.Int       `json:"total_contributed"`
	ExcessContribution *big.Int       `json:"excess_contribution"`
	TokenAmount        *big.Int       `json:"token_amount"`
}

type PartyFinalize struct {
	PartyInfo        Party    `json:"party_info"`
	Result           uint8    `json:"result"`
	TotalSpent       *big.Int `json:"total_spent"`
	Fee              *big.Int `json:"fee"`
	TotalContributed *big.Int `json:"total_contributed"`
}

type PartyBid struct {
	PartyInfo Party    `json:"party_info"`
	BidAmount *big.Int `json:"bid_amount"`
}

type PartyExpire struct {
	PartyInfo   Party          `json:"party_info"`
	TriggeredBy common.Address `json:"triggered_by"`
}

type PartyBuy struct {
	PartyInfo        Party          `json:"party_info"`
	TokenId          *big.Int       `json:"token_id,omitempty"`
	TriggeredBy      common.Address `json:"triggered_by"`
	TargetAddress    common.Address `json:"target_address"`
	EthSpent         *big.Int       `json:"eth_spent"`
	EthFeePaid       *big.Int       `json:"eth-fee-paid"`
	TotalContributed *big.Int       `json:"total_contributed"`
}
