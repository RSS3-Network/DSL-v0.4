package graphqlx

import (
	"time"

	"github.com/shurcooL/graphql"
)

// Publication the format for both Post and Comment
type Publication struct {
	Type      graphql.String `graphql:"__typename"`
	ID        graphql.String `graphql:"id"`
	CreatedAt time.Time      `graphql:"createdAt"`
	Metadata  Metadata       `graphql:"metadata"`
}

type Metadata struct {
	Description graphql.String `graphql:"description"`
	Media       []Media        `graphql:"media"`
}

type Media struct {
	Original struct {
		URL      graphql.String `graphql:"url"`
		MimeType graphql.String `graphql:"mimeType"`
	} `graphql:"original"`
}
