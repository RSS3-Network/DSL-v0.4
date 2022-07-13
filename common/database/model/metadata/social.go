package metadata

import (
	"encoding/json"
	"time"
)

// Post used for Post,Comment,Share
type Post struct {
	TypeOnPlatform string          `json:"type_on_platform,omitempty"`
	Title          string          `json:"title,omitempty"`
	Body           string          `json:"body"`
	Media          json.RawMessage `json:"media,omitempty"`
	Target         *Post           `json:"target,omitempty"`
}

type Profile struct {
	Name        string     `json:"name"`
	Handle      string     `json:"handle"`
	Bio         string     `json:"bio"`
	URL         string     `json:"url"`
	ExpireAt    *time.Time `json:"expire_at"`
	ProfileUris []string   `json:"profile_uri"`
	BannerUris  []string   `json:"banner_uri"`
	SocialUris  []string   `json:"social_uri"`
}
