package crossbell

type GetProfilesResponse struct {
	Count  int               `json:"count"`
	Cursor string            `json:"cursor"`
	List   []ProfileResponse `json:"list"`
}

type ProfileResponse struct {
	CharacterId int    `json:"characterId"`
	Handle      string `json:"handle"`
	Primary     bool   `json:"primary"`
	URI         string `json:"uri"`
	Metadata    struct {
		URI     string   `json:"uri"`
		Type    struct{} `json:"typeo"`
		Content struct{} `json:"content"`
	} `json:"metadata"`
	SocialToken            string `json:"socialToken"`
	Operator               string `json:"operator"`
	Owner                  string `json:"owner"`
	FromAddress            string `json:"fromAddress"`
	CreatedAt              string `json:"createdAt"`
	UpdatedAt              string `json:"updatedAt"`
	DeletedAt              string `json:"deletedAt"`
	TransactionHash        string `json:"transactionHash"`
	BlockNumber            int    `json:"blockNumber"`
	LogIndex               int    `json:"logIndex"`
	UpdatedTransactionHash string `json:"updatedTransactionHash"`
	UpdatedBlockNumber     int    `json:"updatedBlockNumber"`
	UpdatedLogIndex        int    `json:"updatedLogIndex"`
}

type CrossbellProfileStruct struct {
	Type             string   `json:"type"`
	Avatars          []string `json:"avatars"`
	ConnectedAvatars []string `json:"connected_avatars"`
	Name             string   `json:"name"`
	Bio              string   `json:"bio"`
}
