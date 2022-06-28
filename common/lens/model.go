package lens

import (
	"github.com/shurcooL/graphql"
)

type Options struct {
	Address    graphql.String
	Profile    graphql.String
	Cursor     graphql.String
	TotalCount graphql.Int
}
