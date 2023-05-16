package lens

type DatasetMomokaTransactionPostData struct {
	Signature          string `json:"signature"`
	DataAvailabilityId string `json:"dataAvailabilityId"`
	Type               string `json:"type"`
	TimestampProofs    struct {
		Type       string `json:"type"`
		HashPrefix string `json:"hashPrefix"`
		Response   struct {
			Id                  string        `json:"id"`
			Timestamp           int64         `json:"timestamp"`
			Version             string        `json:"version"`
			Public              string        `json:"public"`
			Signature           string        `json:"signature"`
			DeadlineHeight      int           `json:"deadlineHeight"`
			Block               int           `json:"block"`
			ValidatorSignatures []interface{} `json:"validatorSignatures"`
		} `json:"response"`
	} `json:"timestampProofs"`
	ChainProofs struct {
		ThisPublication struct {
			Signature         string `json:"signature"`
			SignedByDelegate  bool   `json:"signedByDelegate"`
			SignatureDeadline int    `json:"signatureDeadline"`
			TypedData         struct {
				Types struct {
					PostWithSig []struct {
						Name string `json:"name"`
						Type string `json:"type"`
					} `json:"PostWithSig"`
				} `json:"types"`
				Domain struct {
					Name              string `json:"name"`
					Version           string `json:"version"`
					ChainId           int    `json:"chainId"`
					VerifyingContract string `json:"verifyingContract"`
				} `json:"domain"`
				Value struct {
					ProfileId               string `json:"profileId"`
					ContentURI              string `json:"contentURI"`
					CollectModule           string `json:"collectModule"`
					CollectModuleInitData   string `json:"collectModuleInitData"`
					ReferenceModule         string `json:"referenceModule"`
					ReferenceModuleInitData string `json:"referenceModuleInitData"`
					Nonce                   int    `json:"nonce"`
					Deadline                int    `json:"deadline"`
				} `json:"value"`
			} `json:"typedData"`
			BlockHash      string `json:"blockHash"`
			BlockNumber    int    `json:"blockNumber"`
			BlockTimestamp int    `json:"blockTimestamp"`
		} `json:"thisPublication"`
		Pointer interface{} `json:"pointer"`
	} `json:"chainProofs"`
	PublicationId string `json:"publicationId"`
	Event         struct {
		ProfileId                 string `json:"profileId"`
		PubId                     string `json:"pubId"`
		ContentURI                string `json:"contentURI"`
		CollectModule             string `json:"collectModule"`
		CollectModuleReturnData   string `json:"collectModuleReturnData"`
		ReferenceModule           string `json:"referenceModule"`
		ReferenceModuleReturnData string `json:"referenceModuleReturnData"`
		Timestamp                 int    `json:"timestamp"`
	} `json:"event"`
}

type DatasetMomokaTransactionCommentData struct {
	ChainProofs struct {
		Pointer struct {
			Location string `json:"location"`
			Type     string `json:"type"`
		} `json:"pointer"`
		ThisPublication struct {
			BlockHash         string `json:"blockHash"`
			BlockNumber       int    `json:"blockNumber"`
			BlockTimestamp    int    `json:"blockTimestamp"`
			Signature         string `json:"signature"`
			SignatureDeadline int    `json:"signatureDeadline"`
			SignedByDelegate  bool   `json:"signedByDelegate"`
			TypedData         struct {
				Domain struct {
					ChainId           int    `json:"chainId"`
					Name              string `json:"name"`
					VerifyingContract string `json:"verifyingContract"`
					Version           string `json:"version"`
				} `json:"domain"`
				Types struct {
					CommentWithSig []struct {
						Name string `json:"name"`
						Type string `json:"type"`
					} `json:"CommentWithSig"`
				} `json:"types"`
				Value struct {
					CollectModule           string `json:"collectModule"`
					CollectModuleInitData   string `json:"collectModuleInitData"`
					ContentURI              string `json:"contentURI"`
					Deadline                int    `json:"deadline"`
					Nonce                   int    `json:"nonce"`
					ProfileId               string `json:"profileId"`
					ProfileIdPointed        string `json:"profileIdPointed"`
					PubIdPointed            string `json:"pubIdPointed"`
					ReferenceModule         string `json:"referenceModule"`
					ReferenceModuleData     string `json:"referenceModuleData"`
					ReferenceModuleInitData string `json:"referenceModuleInitData"`
				} `json:"value"`
			} `json:"typedData"`
		} `json:"thisPublication"`
	} `json:"chainProofs"`
	DataAvailabilityId string `json:"dataAvailabilityId"`
	Event              struct {
		CollectModule             string `json:"collectModule"`
		CollectModuleReturnData   string `json:"collectModuleReturnData"`
		ContentURI                string `json:"contentURI"`
		ProfileId                 string `json:"profileId"`
		ProfileIdPointed          string `json:"profileIdPointed"`
		PubId                     string `json:"pubId"`
		PubIdPointed              string `json:"pubIdPointed"`
		ReferenceModule           string `json:"referenceModule"`
		ReferenceModuleData       string `json:"referenceModuleData"`
		ReferenceModuleReturnData string `json:"referenceModuleReturnData"`
		Timestamp                 int    `json:"timestamp"`
	} `json:"event"`
	PublicationId   string `json:"publicationId"`
	Signature       string `json:"signature"`
	TimestampProofs struct {
		HashPrefix string `json:"hashPrefix"`
		Response   struct {
			Block               int           `json:"block"`
			DeadlineHeight      int           `json:"deadlineHeight"`
			Id                  string        `json:"id"`
			Public              string        `json:"public"`
			Signature           string        `json:"signature"`
			Timestamp           int64         `json:"timestamp"`
			ValidatorSignatures []interface{} `json:"validatorSignatures"`
			Version             string        `json:"version"`
		} `json:"response"`
		Type string `json:"type"`
	} `json:"timestampProofs"`
	Type string `json:"type"`
}

type DatasetMomokaTransactionMirrorData struct {
	ChainProofs struct {
		Pointer struct {
			Location string `json:"location"`
			Type     string `json:"type"`
		} `json:"pointer"`
		ThisPublication struct {
			BlockHash         string `json:"blockHash"`
			BlockNumber       int    `json:"blockNumber"`
			BlockTimestamp    int    `json:"blockTimestamp"`
			Signature         string `json:"signature"`
			SignatureDeadline int    `json:"signatureDeadline"`
			SignedByDelegate  bool   `json:"signedByDelegate"`
			TypedData         struct {
				Domain struct {
					ChainId           int    `json:"chainId"`
					Name              string `json:"name"`
					VerifyingContract string `json:"verifyingContract"`
					Version           string `json:"version"`
				} `json:"domain"`
				Types struct {
					MirrorWithSig []struct {
						Name string `json:"name"`
						Type string `json:"type"`
					} `json:"MirrorWithSig"`
				} `json:"types"`
				Value struct {
					Deadline                int    `json:"deadline"`
					Nonce                   int    `json:"nonce"`
					ProfileId               string `json:"profileId"`
					ProfileIdPointed        string `json:"profileIdPointed"`
					PubIdPointed            string `json:"pubIdPointed"`
					ReferenceModule         string `json:"referenceModule"`
					ReferenceModuleData     string `json:"referenceModuleData"`
					ReferenceModuleInitData string `json:"referenceModuleInitData"`
				} `json:"value"`
			} `json:"typedData"`
		} `json:"thisPublication"`
	} `json:"chainProofs"`
	DataAvailabilityId string `json:"dataAvailabilityId"`
	Event              struct {
		ProfileId                 string `json:"profileId"`
		ProfileIdPointed          string `json:"profileIdPointed"`
		PubId                     string `json:"pubId"`
		PubIdPointed              string `json:"pubIdPointed"`
		ReferenceModule           string `json:"referenceModule"`
		ReferenceModuleData       string `json:"referenceModuleData"`
		ReferenceModuleReturnData string `json:"referenceModuleReturnData"`
		Timestamp                 int    `json:"timestamp"`
	} `json:"event"`
	PublicationId   string `json:"publicationId"`
	Signature       string `json:"signature"`
	TimestampProofs struct {
		HashPrefix string `json:"hashPrefix"`
		Response   struct {
			Block               int           `json:"block"`
			DeadlineHeight      int           `json:"deadlineHeight"`
			Id                  string        `json:"id"`
			Public              string        `json:"public"`
			Signature           string        `json:"signature"`
			Timestamp           int64         `json:"timestamp"`
			ValidatorSignatures []interface{} `json:"validatorSignatures"`
			Version             string        `json:"version"`
		} `json:"response"`
		Type string `json:"type"`
	} `json:"timestampProofs"`
	Type string `json:"type"`
}

type DatasetMomokaTransactionContent struct {
	AnimationUrl interface{} `json:"animation_url"`
	AppId        string      `json:"appId"`
	Attributes   []struct {
		DisplayType string `json:"displayType"`
		TraitType   string `json:"traitType"`
		Value       string `json:"value"`
	} `json:"attributes"`
	Content          string      `json:"content"`
	ContentWarning   interface{} `json:"contentWarning"`
	ExternalUrl      string      `json:"external_url"`
	Image            string      `json:"image"`
	ImageMimeType    string      `json:"imageMimeType"`
	Locale           string      `json:"locale"`
	MainContentFocus string      `json:"mainContentFocus"`
	Media            []struct {
		AltTag string `json:"altTag"`
		Item   string `json:"item"`
		Type   string `json:"type"`
	} `json:"media"`
	MetadataId string        `json:"metadata_id"`
	Name       string        `json:"name"`
	Tags       []interface{} `json:"tags"`
	Version    string        `json:"version"`
}
