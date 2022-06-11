package graphqlx

import (
	"github.com/shopspring/decimal"
	"github.com/shurcooL/graphql"
)

// Space

type SpaceStrategieParam struct {
	Symbol   graphql.String  `json:"symbol"`
	Address  graphql.String  `json:"address"`
	Decimals decimal.Decimal `json:"decimals"`
}

type SpaceStrategie struct {
	Name   graphql.String      `json:"name"`
	Params SpaceStrategieParam `json:"params"`
}

type SpaceFilter struct {
	MinScore    decimal.Decimal `json:"minScore"`
	OnlyMembers graphql.Boolean `json:"onlyMembers"`
}

type Space struct {
	Id         graphql.String   `json:"id"`
	Name       graphql.String   `json:"name"`
	About      graphql.String   `json:"about"`
	Network    graphql.String   `json:"network"`
	Symbol     graphql.String   `json:"symbol"`
	Members    []graphql.String `json:"members"`
	Strategies []SpaceStrategie `json:"strategies"`
	Admins     []graphql.String `json:"admins"`
	Filters    []SpaceFilter    `json:"filters"`
	// plugins
}

// Proposal

type ProposalSpace struct {
	Id   graphql.String `json:"id"`
	Name graphql.String `json:"name"`
}

type Proposal struct {
	Id       graphql.String   `json:"id"`
	Title    graphql.String   `json:"title"`
	Body     graphql.String   `json:"body"`
	Choices  []graphql.String `json:"choices"`
	Start    graphql.Int      `json:"start"`
	End      graphql.Int      `json:"end"`
	Snapshot graphql.String   `json:"snapshot"`
	State    graphql.String   `json:"state"`
	Author   graphql.String   `json:"author"`
	Space    ProposalSpace    `json:"space"`
}

// Vote

type VoteProposal struct {
	Id graphql.String `json:"id"`
}

type VoteSpace struct {
	Id graphql.String `json:"id"`
}

type Vote struct {
	Id       graphql.String `json:"id"`
	Voter    graphql.String `json:"voter"`
	Created  graphql.Int    `json:"created"`
	Proposal VoteProposal   `json:"proposal"`
	Choice   graphql.Int    `json:"choice"`
	Space    VoteSpace      `json:"space"`
}
