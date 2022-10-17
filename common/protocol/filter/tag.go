package filter

const (
	TagTransaction string = "transaction"
	TagExchange    string = "exchange"
	TagBridge      string = "bridge"
	TagCollectible string = "collectible"
	TagSocial      string = "social"
	TagDonation    string = "donation"
	TagGovernance  string = "governance"
)

var ValidTag = map[string]string{
	TagTransaction: TagTransaction,
	TagExchange:    TagExchange,
	TagBridge:      TagBridge,
	TagCollectible: TagCollectible,
	TagSocial:      TagSocial,
	TagDonation:    TagDonation,
	TagGovernance:  TagGovernance,
}

var TagPriority = map[string]int{
	TagTransaction: 1,
	TagExchange:    2,
	TagBridge:      2,
	TagCollectible: 3,
	TagSocial:      4,
	TagDonation:    4,
	TagGovernance:  4,
}

func UpdateTagAndType(targetTag string, currentTag string, targetType string, currentType string) (string, string) {
	if TagPriority[targetTag] > TagPriority[currentTag] {
		return targetTag, targetType
	}
	return currentTag, currentType
}

func UpdateTag(targetTag string, currentTag string) string {
	if TagPriority[targetTag] > TagPriority[currentTag] {
		return targetTag
	}
	return currentTag
}
