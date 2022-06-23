package graphqlx

import "github.com/hasura/go-graphql-client"

type TagOperator graphql.String

const (
	TagOperatorEQ  TagOperator = "EQ"
	TagOperatorNEQ TagOperator = "NEQ"
)

type TransactionConnection struct {
	PageInfo PageInfo          `json:"pageInfo"`
	Edges    []TransactionEdge `json:"edges"`
}

type PageInfo struct {
	HasNextPage graphql.Boolean `json:"hasNextPage"`
}

type TransactionEdge struct {
	Cursor graphql.String `json:"cursor"`
	Node   Transaction    `json:"node"`
}

type Transaction struct {
	ID        graphql.ID     `json:"id"`
	Anchor    graphql.String `json:"anchor"`
	Signature graphql.String `json:"signature"`
	Recipient graphql.String `json:"recipient"`
	Owner     Owner          `json:"owner"`
	Fee       Amount         `json:"fee"`
	Quantity  Amount         `json:"quantity"`
	Data      MetaData       `json:"data"`
	Tags      []Tag          `json:"tags"`
	Block     Block          `json:"block"`
	Parent    Parent         `json:"parent"`
	BundledIn Bundle         `json:"bundledIn"`
}

type Owner struct {
	Address graphql.String `json:"address"`
	Key     graphql.String `json:"key"`
}

type Amount struct {
	Winston graphql.String `json:"winston"`
	AR      graphql.String `json:"ar"`
}

type MetaData struct {
	Size graphql.String `json:"size"`
	Type graphql.String `json:"type"`
}

type Tag struct {
	Name  graphql.String `json:"name"`
	Value graphql.String `json:"value"`
}

type Block struct {
	ID        graphql.ID  `json:"id"`
	Timestamp graphql.Int `json:"timestamp"`
	Height    graphql.Int `json:"height"`
	Previous  graphql.ID  `json:"previous"`
}

type Parent struct {
	ID graphql.ID `json:"id"`
}

type Bundle struct {
	ID graphql.ID `json:"id"`
}

type BlockFilter struct {
	Min graphql.Int `json:"min"`
	Max graphql.Int `json:"max"`
}

type TagFilter struct {
	Name        graphql.String   `json:"name"`
	Values      []graphql.String `json:"values"`
	TagOperator TagOperator      `json:"op"`
}
