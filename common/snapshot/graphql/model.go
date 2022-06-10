package graphqlx

import "github.com/shopspring/decimal"

// Space

type SpaceStrategieParam struct {
	Symbol   string          `json:"symbol"`
	Address  string          `json:"address"`
	Decimals decimal.Decimal `json:"decimals"`
}

type SpaceStrategie struct {
	Name   string              `json:"name"`
	Params SpaceStrategieParam `json:"params"`
}

type SpaceFilter struct {
	MinScore    decimal.Decimal `json:"minScore"`
	OnlyMembers bool            `json:"onlyMembers"`
}

type Space struct {
	Id         string           `json:"id"`
	Name       string           `json:"name"`
	About      string           `json:"about"`
	Network    string           `json:"network"`
	Symbol     string           `json:"symbol"`
	Members    []string         `json:"members"`
	Strategies []SpaceStrategie `json:"strategies"`
	Admins     []string         `json:"admins"`
	Filters    []SpaceFilter    `json:"filters"`
	// plugins
}

// Proposal

type ProposalSpace struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Proposal struct {
	Id       string          `json:"id"`
	Title    string          `json:"title"`
	Body     string          `json:"body"`
	Choices  []string        `json:"choices"`
	Start    decimal.Decimal `json:"start"`
	End      decimal.Decimal `json:"end"`
	Snapshot string          `json:"snapshot"`
	State    string          `json:"state"`
	Author   string          `json:"author"`
	Space    ProposalSpace   `json:"space"`
}

// Vote

type VoteProposal struct {
	Id string `json:"id"`
}

type VoteSpace struct {
	Id string `json:"id"`
}

type Vote struct {
	Id       string          `json:"id"`
	Voter    string          `json:"voter"`
	Created  decimal.Decimal `json:"created"`
	Proposal VoteProposal    `json:"proposal"`
	Choice   decimal.Decimal `json:"choice"`
	Space    VoteSpace       `json:"space"`
}
