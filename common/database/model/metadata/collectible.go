package metadata

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Party struct {
	PartyAddress     string           `json:"address"`
	Name             string           `json:"name"`
	Symbol           string           `json:"symbol"`
	PartyType        string           `json:"type"`
	Creator          string           `json:"creator,omitempty"`
	NftContract      string           `json:"contract_address,omitempty"`
	TokenId          string           `json:"token_id,omitempty"`
	MarketWrapper    string           `json:"marketplace"`
	AuctionId        string           `json:"auction_id,omitempty"`
	ExpiredTime      *time.Time       `json:"expired_date,omitempty"`
	Deciders         []common.Address `json:"deciders,omitempty"`
	PartyStatus      uint8            `json:"status,omitempty"`
	Action           string           `json:"action,omitempty"`
	TotalContributed *Token           `json:"total_contributed_amount,omitempty"`
}

type PartyContribute struct {
	PartyInfo                       Party  `json:"crowdfund_info"`
	Action                          string `json:"action"`
	Contributor                     string `json:"-"`
	Amount                          *Token `json:"amount"`
	PreviousTotalContributedToParty *Token `json:"previous_total_contributed"`
	TotalFromContributor            *Token `json:"total_from_contributor"`
}

type PartyClaim struct {
	PartyInfo          Party  `json:"crowdfund_info"`
	Action             string `json:"action"`
	Contributor        string `json:"-"`
	TotalContributed   *Token `json:"total_contributed"`
	ExcessContribution *Token `json:"excess_contribution"`
	TokenAmount        *Token `json:"fractional_token_amount,omitempty"`
}

type PartyFinalize struct {
	PartyInfo        Party  `json:"crowdfund_info"`
	Action           string `json:"action"`
	Result           uint8  `json:"result"`
	TotalSpent       *Token `json:"total_spent"`
	Fee              *Token `json:"fee"`
	TotalContributed *Token `json:"total_contributed"`
}

type PartyBid struct {
	PartyInfo Party  `json:"crowdfund_info"`
	Action    string `json:"action"`
	BidAmount *Token `json:"bid_amount"`
}

type PartyExpire struct {
	PartyInfo   Party  `json:"crowdfund_info"`
	Action      string `json:"action"`
	TriggeredBy string `json:"-"`
}

type PartyBuy struct {
	PartyInfo        Party    `json:"crowdfund_info"`
	Action           string   `json:"action"`
	TokenId          *big.Int `json:"token_id,omitempty"`
	TriggeredBy      string   `json:"-"`
	TargetAddress    string   `json:"-"`
	EthSpent         *Token   `json:"eth_spent"`
	EthFeePaid       *Token   `json:"platform_fee"`
	TotalContributed *Token   `json:"total_contributed"`
}
