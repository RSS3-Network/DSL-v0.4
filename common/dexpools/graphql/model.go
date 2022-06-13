package graphqlx

import "github.com/hasura/go-graphql-client"

type Token struct {
	Name   graphql.String `json:"name"`
	Symbol graphql.String `json:"symbol"`
}

type Pair struct {
	ID     graphql.String `json:"id"`
	Token0 Token          `json:"token0"`
	Token1 Token          `json:"token1"`
}
