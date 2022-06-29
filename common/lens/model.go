package lens

import (
	"github.com/hasura/go-graphql-client"
)

type Options struct {
	Address    graphql.String
	Profile    graphql.String
	Cursor     graphql.String
	TotalCount graphql.Int
}
