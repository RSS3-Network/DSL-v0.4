package metadata

import (
	"time"
)

type Vote struct {
	TypeOnPlatform []string  `json:"type_on_platform,omitempty"`
	Choice         string    `json:"choice"`
	Proposal       *Proposal `json:"proposal"`
}

type Proposal struct {
	TypeOnPlatform []string      `json:"type_on_platform,omitempty"`
	Id             string        `json:"id"`
	Title          string        `json:"title,omitempty"`
	Body           string        `json:"body,omitempty"`
	Options        []string      `json:"options"`
	StartAt        time.Time     `json:"start_at"`
	EndAt          time.Time     `json:"end_at"`
	Organization   *Organization `json:"organization,omitempty"`
}

type Organization struct {
	TypeOnPlatform []string `json:"type_on_platform,omitempty"`
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	About          string   `json:"about,omitempty"`
}
