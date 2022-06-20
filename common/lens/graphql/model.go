package graphqlx

import (
	"time"

	"github.com/shurcooL/graphql"
)

type Post struct {
	Type       graphql.String `graphql:"__typename" json:"type"`
	ID         graphql.String `json:"id"`
	RelatedURL graphql.String `graphql:"onChainContentURI" json:"related_urls"`
	Platform   graphql.String `graphql:"appId" json:"platform"`
	CreatedAt  time.Time      `json:"createdAt"`
	Metadata   Metadata       `json:"metadata"`
}

type Publication struct {
	Type       graphql.String `graphql:"__typename" json:"type"`
	ID         graphql.String `json:"id"`
	RelatedURL graphql.String `graphql:"onChainContentURI" json:"related_urls"`
	Platform   graphql.String `graphql:"appId" json:"platform"`
	CreatedAt  time.Time      `json:"createdAt"`
	Metadata   Metadata       `json:"metadata"`
	CommentOn  struct {
		Post    Post `graphql:"... on Post"`
		Comment Post `graphql:"... on Comment"`
	} `graphql:"commentOn"`
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

type PageInfo struct {
	Next       graphql.String
	TotalCount graphql.Int
}
