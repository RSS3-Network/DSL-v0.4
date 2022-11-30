package metadata

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
}

type Media struct {
	Address  string `json:"address"`
	MimeType string `json:"mime_type"`
}

type Curation struct {
	Target Post   `json:"target"`
	Reward Token  `json:"reward"`
	From   string `json:"from"`
	To     string `json:"to"`
}
