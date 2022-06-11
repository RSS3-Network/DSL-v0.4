package graphqlx

import (
	"time"

	"github.com/shurcooL/graphql"
)

// Publication the format for both Post and Comment
type Publication struct {
	Type      graphql.String `graphql:"__typename"`
	ID        graphql.String `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	Metadata  Metadata       `json:"metadata"`
}

type Metadata struct {
	Description graphql.String `json:"description"`
	Media       []Media        `json:"media"`
}

type Media struct {
	Original struct {
		URL      graphql.String `json:"url"`
		MimeType graphql.String `json:"mimeType"`
	} `graphql:"original"`
}
