package coingecko

type CoinListParameter struct {
	IncludePlatform bool `url:"include_platform,omitempty"`
}

type Coin struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Symbol    string            `json:"symbol"`
	Platforms map[string]string `json:"platforms,omitempty"`
	Image     *CoinImage        `json:"image,omitempty"`
}

type CoinImage struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type CoinParameter struct {
	Localization  string `url:"localization,omitempty"`
	Tickers       bool   `url:"tickers,omitempty"`
	MarketData    bool   `url:"market_data,omitempty"`
	CommunityData bool   `url:"community_data,omitempty"`
	DeveloperData bool   `url:"developer_data,omitempty"`
	Sparkline     bool   `url:"sparkline,omitempty"`
}
