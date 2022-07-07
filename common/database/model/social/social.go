package social

import "encoding/json"

// Content used for Post,Comment,Share
type Content struct {
	TypeOnPlatform string          `json:"type_on_platform,omitempty"`
	Title          string          `json:"title,omitempty"`
	Body           string          `json:"body"`
	Media          json.RawMessage `json:"media,omitempty"`
	Target         *Content        `json:"target,omitempty"`
}
