package poap

type Action struct {
	Event struct {
		ID          int    `json:"id"`
		FancyID     string `json:"fancy_id"`
		Name        string `json:"name"`
		EventURL    string `json:"event_url"`
		ImageURL    string `json:"image_url"`
		Country     string `json:"country"`
		City        string `json:"city"`
		Description string `json:"description"`
		Year        int    `json:"year"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		ExpiryDate  string `json:"expiry_date"`
		Supply      int    `json:"supply"`
	} `json:"event"`
	TokenID string `json:"tokenId"`
	Owner   string `json:"owner"`
	Chain   string `json:"chain"`
	Created string `json:"created"`
}

type Token struct {
	Event struct {
		ID          int    `json:"id"`
		FancyID     string `json:"fancy_id"`
		Name        string `json:"name"`
		EventURL    string `json:"event_url"`
		ImageURL    string `json:"image_url"`
		Country     string `json:"country"`
		City        string `json:"city"`
		Description string `json:"description"`
		Year        int    `json:"year"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		ExpiryDate  string `json:"expiry_date"`
	} `json:"event"`
	TokenID string `json:"tokenId"`
	Owner   string `json:"owner"`
	Layer   string `json:"layer"`
}
