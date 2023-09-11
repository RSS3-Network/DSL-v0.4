package friendtech

type UserResponse struct {
	ID              int64  `json:"id"`
	Address         string `json:"address"`
	TwitterUsername string `json:"twitterUsername"`
	TwitterName     string `json:"twitterName"`
	TwitterPfpURL   string `json:"twitterPfpUrl"`
	TwitterUserID   string `json:"twitterUserId"`
	ShareSupply     int    `json:"shareSupply"`
	DisplayPrice    string `json:"displayPrice"`
}
