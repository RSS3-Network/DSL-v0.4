package filter

const (
	TagTransaction string = "transaction"
	TagExchange    string = "exchange"
	TagStaking     string = "staking"
	TagCollectible string = "collectible"
	TagSocial      string = "social"
	TagDonation    string = "donation"
	TagGovernance  string = "governance"
	TagMetaverse   string = "metaverse"
)

var ValidTag = map[string]string{
	TagTransaction: TagTransaction,
	TagStaking:     TagStaking,
	TagExchange:    TagExchange,
	TagCollectible: TagCollectible,
	TagSocial:      TagSocial,
	TagDonation:    TagDonation,
	TagGovernance:  TagGovernance,
	TagMetaverse:   TagMetaverse,
}

var TagPriority = map[string]int{
	TagTransaction: 1,
	TagStaking:     2,
	TagExchange:    3,
	TagCollectible: 4,
	TagSocial:      5,
	TagDonation:    5,
	TagGovernance:  5,
	TagMetaverse:   5,
}

func UpdateTagAndType(targetTag string, currentTag string, targetType string, currentType string) (string, string) {
	if TagPriority[targetTag] >= TagPriority[currentTag] {
		return targetTag, targetType
	}

	return currentTag, currentType
}

func UpdateTag(targetTag string, currentTag string) string {
	if TagPriority[targetTag] >= TagPriority[currentTag] {
		return targetTag
	}

	return currentTag
}
