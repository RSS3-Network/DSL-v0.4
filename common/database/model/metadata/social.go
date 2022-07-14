package metadata

import (
	"encoding/json"
)

// Post used for Post,Comment,Share
type Post struct {
	TypeOnPlatform string          `json:"type_on_platform,omitempty"`
	Title          string          `json:"title,omitempty"`
	Body           string          `json:"body"`
	Media          json.RawMessage `json:"media,omitempty"`
	Target         *Post           `json:"target,omitempty"`
}
