package farcaster

type Cast struct {
	Body struct {
		Type        string `json:"type"`
		PublishedAt int64  `json:"publishedAt"`
		Sequence    int    `json:"sequence"`
		Username    string `json:"username"`
		Address     string `json:"address"`
		Data        struct {
			Text                  string `json:"text"`
			ReplyParentMerkleRoot string `json:"replyParentMerkleRoot"`
		} `json:"data"`
		// PrevMerkleRoot string `json:"prevMerkleRoot"`
	} `json:"body"`
	MerkleRoot string `json:"merkleRoot"`
	// Signature  string `json:"signature"`
	Meta struct {
		DisplayName string `json:"displayName"`
		Avatar      string `json:"avatar"`
		// IsVerifiedAvatar bool   `json:"isVerifiedAvatar"`
		// NumReplyChildren int    `json:"numReplyChildren"`
		// Reactions        struct {
		//	Count int    `json:"count"`
		//	Type  string `json:"type"`
		//	Self  bool   `json:"self"`
		// } `json:"reactions"`
		// Recasts struct {
		//	Count int  `json:"count"`
		//	Self  bool `json:"self"`
		// } `json:"recasts"`
		// Watches struct {
		//	Count int  `json:"count"`
		//	Self  bool `json:"self"`
		// } `json:"watches"`
		ReplyParentUsername struct {
			Address  string `json:"address"`
			Username string `json:"username"`
		} `json:"replyParentUsername"`
	} `json:"meta"`
	// Attachments struct {
	//	OpenGraph []interface{} `json:"openGraph"`
	// } `json:"attachments"`
}

type FarcasterUser struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Avatar   struct {
		Url        string `json:"url"`
		IsVerified bool   `json:"isVerified"`
	} `json:"avatar"`
	DisplayName       string `json:"displayName"`
	IsViewerFollowing bool   `json:"isViewerFollowing"`
	IsFollowingViewer bool   `json:"isFollowingViewer"`
	Profile           struct {
		Bio struct {
			Text     string        `json:"text"`
			Mentions []interface{} `json:"mentions"`
		} `json:"bio"`
	} `json:"profile"`
}

type Directory struct {
	Body struct {
		AddressActivityUrl string `json:"addressActivityUrl"`
		AvatarUrl          string `json:"avatarUrl"`
		DisplayName        string `json:"displayName"`
		ProofUrl           string `json:"proofUrl"`
		Timestamp          int64  `json:"timestamp"`
		Version            int    `json:"version"`
	} `json:"body"`
	MerkleRoot string `json:"merkleRoot"`
	Signature  string `json:"signature"`
}

type ConvertAddress struct {
	SignedMessage    string `json:"signedMessage"`
	SignerAddress    string `json:"signerAddress"`
	FarcasterAddress string `json:"farcasterAddress"`
	OriginalMessage  string `json:"originalMessage"`
}

type CacheAddress struct {
	EvmAddress string `json:"evmAddress"`
	CastNumber int64  `json:"castNumber"`
}
