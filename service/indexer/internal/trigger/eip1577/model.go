package eip1577

type Planet struct {
	About              string          `json:"about"`
	Articles           []PlanetArticle `json:"articles"`
	Created            float64         `json:"created"`
	GitHubUsername     string          `json:"githubUsername"`
	ID                 string          `json:"id"`
	IPNS               string          `json:"ipns"`
	Name               string          `json:"name"`
	PlausibleAPIServer string          `json:"plausibleAPIServer"`
	PlausibleDomain    string          `json:"plausibleDomain"`
	PlausibleEnabled   bool            `json:"plausibleEnabled"`
	TwitterUsername    string          `json:"twitterUsername"`
	Updated            float64         `json:"updated"`
}

type PlanetArticle struct {
	Attachments   []string `json:"attachments,omitempty"`
	Content       string   `json:"content"`
	Created       float64  `json:"created"`
	HasAudio      bool     `json:"hasAudio"`
	HasVideo      bool     `json:"hasVideo"`
	ID            string   `json:"id"`
	Link          string   `json:"link"`
	Title         string   `json:"title"`
	AudioFilename string   `json:"audioFilename,omitempty"`
	VideoFilename string   `json:"videoFilename,omitempty"`
}
