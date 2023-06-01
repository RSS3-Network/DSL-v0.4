package graphqlx

import (
	"encoding/json"

	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/arweave/graphql"

	"github.com/hasura/go-graphql-client"
)

type GetSpacesVariable struct {
	Owners []graphql.String     `json:"owners"`
	Tags   []graphqlx.TagFilter `json:"tags"`
	Block  graphqlx.BlockFilter `json:"block"`
}

type SpaceStrategieParam struct {
	Symbol   graphql.String `json:"symbol"`
	Address  graphql.String `json:"address"`
	Decimals graphql.Int    `json:"decimals"`
}

type SpaceStrategie struct {
	Name   graphql.String  `json:"name"`
	Params json.RawMessage `json:"params"`
}

type SpaceFilter struct {
	MinScore    graphql.Float   `json:"minScore"`
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
	Filters    SpaceFilter      `json:"filters"`
	Plugins    json.RawMessage  `json:"plugins"`
}

type SpaceWhere struct {
	IdArray []graphql.String `json:"id_in"`
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
	Created  graphql.Int      `json:"created"`
}

type ProposalWhere struct {
	Space_in   []graphql.String `json:"space_in"`    // param: ["ens.eth"]
	State      graphql.String   `json:"state"`       // param: "active" or "closed"
	CreatedGte graphql.Int      `json:"created_gte"` // param: 1682170931
}

// Vote

type VoteProposal struct {
	Id graphql.String `json:"id"`
}

type VoteSpace struct {
	Id graphql.String `json:"id"`
}

type Vote struct {
	Id       graphql.String  `json:"id"`
	Voter    graphql.String  `json:"voter"`
	Created  graphql.Int     `json:"created"`
	Proposal VoteProposal    `json:"proposal"`
	Choice   json.RawMessage `json:"choice"`
	Space    VoteSpace       `json:"space"`
}

type VoteWhere struct {
	CreatedGte graphql.Int `json:"created_gte"` // param: 1682170931
}
