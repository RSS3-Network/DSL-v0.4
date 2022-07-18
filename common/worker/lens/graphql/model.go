package graphqlx

import (
	"time"

	"github.com/hasura/go-graphql-client"
)

type Profile struct {
	ID           graphql.String  `json:"id"`
	Name         graphql.String  `json:"name"`
	Handle       graphql.String  `json:"handle"`
	Bio          graphql.String  `json:"bio"`
	IsDefault    graphql.Boolean `json:"isDefault"`
	Picture      ProfileMedia    `json:"picture"`
	CoverPicture ProfileMedia    `json:"coverPicture"`
	Metadata     graphql.String  `json:"metadata"`
}

// Publication is a custom type to hold Post, Comment, and Mirror.
// Target is used to hold either CommentOn or MirrorOf
type Publication struct {
	Type       graphql.String `json:"type"`
	ID         graphql.String `json:"id"`
	RelatedURL graphql.String `json:"related_urls"`
	Platform   graphql.String `json:"platform"`
	CreatedAt  time.Time      `json:"createdAt"`
	Metadata   Metadata       `json:"metadata"`
	Target     Post           `json:"target"`
}

// Post is the basic type
type Post struct {
	Type       graphql.String `graphql:"__typename" json:"type"`
	ID         graphql.String `json:"id"`
	RelatedURL graphql.String `graphql:"onChainContentURI" json:"related_urls"`
	Platform   graphql.String `graphql:"appId" json:"platform"`
	CreatedAt  time.Time      `json:"createdAt"`
	Metadata   Metadata       `json:"metadata"`
}

// Comment has a CommentOn field on top of a Post.
type Comment struct {
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

// Mirror is equivalent to re-post, it has a MirrorOf field on top of a Post.
type Mirror struct {
	Type       graphql.String `graphql:"__typename" json:"type"`
	ID         graphql.String `json:"id"`
	RelatedURL graphql.String `graphql:"onChainContentURI" json:"related_urls"`
	Platform   graphql.String `graphql:"appId" json:"platform"`
	CreatedAt  time.Time      `json:"createdAt"`
	Metadata   Metadata       `json:"metadata"`
	MirrorOf   struct {
		Post    Post `graphql:"... on Post"`
		Comment Post `graphql:"... on Comment"`
	} `graphql:"mirrorOf"`
}

type Item struct {
	Post    Post    `graphql:"... on Post"`
	Comment Comment `graphql:"... on Comment"`
	Mirror  Mirror  `graphql:"... on Mirror"`
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

type NFTMedia struct {
	Uri             graphql.String  `json:"uri"`
	ChainId         graphql.Int     `json:"chainId"`
	ContractAddress graphql.String  `json:"contractAddress"`
	TokenId         graphql.String  `json:"tokenId"`
	Verified        graphql.Boolean `json:"verified"`
}

type ProfileMedia struct {
	Type     graphql.String `graphql:"__typename" json:"type"`
	MediaSet Media          `graphql:"... on MediaSet"`
	NFTImage NFTMedia       `graphql:"... on NftImage"`
}

type PageInfo struct {
	Next       graphql.String
	TotalCount graphql.Int
}
