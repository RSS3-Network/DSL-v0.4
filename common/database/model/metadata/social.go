package metadata

import "math/big"

// Post used for Post,Comment,Share
type Post struct {
	// TypeOnPlatform: used when the type on the platform is different from our type definition
	TypeOnPlatform []string `json:"type_on_platform,omitempty"`
	Action         string   `json:"action,omitempty"`
	CreatedAt      string   `json:"created_at,omitempty"`
	Author         []string `json:"author,omitempty"`
	Title          string   `json:"title,omitempty"`
	Summary        string   `json:"summary,omitempty"`
	Body           string   `json:"body,omitempty"`
	Media          []Media  `json:"media,omitempty"`
	Target         *Post    `json:"target,omitempty"`
	TargetURL      string   `json:"target_url,omitempty"`
	OriginNoteID   string   `json:"origin_note_id,omitempty"`
	Categories     []string `json:"categories,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	Reward         *Token   `json:"reward,omitempty"`
	PublicationID  *big.Int `json:"publication_id,omitempty"`
	ProfileID      *big.Int `json:"profile_id,omitempty"`
}

type Media struct {
	Address  string `json:"address"`
	MimeType string `json:"mime_type"`
}
