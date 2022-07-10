package filter

const (
	TagTransaction = "transaction"
	TagExchange    = "exchange"
	TagCollectible = "collectible"
	TagSocial      = "social"
	TagDonation    = "donation"
	TagGovernance  = "governance"
)

var TagPriority = map[string]int{
	TagTransaction: 1,
	TagExchange:    2,
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
