package gitcoin

type ApplicationResp struct {
	Signature   string `json:"signature"`
	Application struct {
		Round     string `json:"round"`
		Recipient string `json:"recipient"`
		Project   struct {
			LastUpdated int    `json:"lastUpdated"`
			CreatedAt   int64  `json:"createdAt"`
			ID          string `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Website     string `json:"website"`
			BannerImg   string `json:"bannerImg"`
			LogoImg     string `json:"logoImg"`
			MetaPtr     struct {
				Protocol string `json:"protocol"`
				Pointer  string `json:"pointer"`
			} `json:"metaPtr"`
		} `json:"project"`
	} `json:"application"`
}
