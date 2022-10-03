package metadata

// Post used for Post,Comment,Share
type Post struct {
	TypeOnPlatform []string `json:"type_on_platform,omitempty"`
	Title          string   `json:"title,omitempty"`
	Summary        string   `json:"summary"`
	Body           string   `json:"body"`
	Media          []Media  `json:"media,omitempty"`
	Target         *Post    `json:"target,omitempty"`
}

type Media struct {
	Address  string `json:"address"`
	MimeType string `json:"mime_type"`
}
