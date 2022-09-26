package eip1577

import "time"

type Result struct {
	Cid         string    `json:"cid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	Platform    string    `json:"platform"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Feeds       []struct {
		Cid       string    `json:"cid"`
		Path      string    `json:"path"`
		URL       string    `json:"url"`
		Title     string    `json:"title"`
		Summary   string    `json:"summary"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"feeds"`
}
