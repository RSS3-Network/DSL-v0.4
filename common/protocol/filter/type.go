package filter

const (
	// transaction types:  transfer | mint | burn | cancel
	TransactionTransfer string = "transfer"
	TransactionMint     string = "mint"
	TransactionBurn     string = "burn"
	TransactionSelf     string = "self"

	// exchange types:  withdraw | deposit | swap
	ExchangeWithdraw string = "withdraw"
	ExchangeDeposit  string = "deposit"
	ExchangeSwap     string = "swap"

	// NFT types: transfer | mint | burn | poap
	CollectibleTransfer string = "transfer"
	CollectibleTrade    string = "trade"
	CollectibleMint     string = "mint"
	CollectibleBurn     string = "burn"
	CollectiblePoap     string = "poap"

	// social types: post | comment | share (retweet) | profile | follow | unfollow | like
	SocialPost     string = "post"
	SocialRevise   string = "revise"
	SocialComment  string = "comment"
	SocialShare    string = "share"
	SocialProfile  string = "profile"
	SocialFollow   string = "follow"
	SocialUnfollow string = "unfollow"
	SocialLike     string = "like"

	// two sub-types for profile
	SocialProfileCreate string = "create"
	SocialProfileUpdate string = "update"

	// Governance types: propose | vote
	GovernancePropose string = "propose"
	GovernanceVote    string = "vote"

	// donation type: donate
	DonationLaunch string = "launch"
	DonationDonate string = "donate"
)

var ValidTypeMap = map[string][]string{
	TagTransaction: {TransactionTransfer, TransactionMint, TransactionBurn, TransactionSelf},
	TagExchange:    {ExchangeWithdraw, ExchangeDeposit, ExchangeSwap},
	TagCollectible: {CollectibleTrade, CollectibleMint, CollectibleBurn, CollectiblePoap},
	TagSocial:      {SocialPost, SocialRevise, SocialComment, SocialShare, SocialProfile, SocialFollow, SocialUnfollow, SocialLike},
	TagDonation:    {DonationLaunch, DonationDonate},
	TagGovernance:  {GovernancePropose, GovernanceVote},
}

func CheckTypeValid(tag string, transferType string) bool {
	if len(tag) == 0 {
		return false
	}

	validTypeList := ValidTypeMap[tag]
	for _, t := range validTypeList {
		if t == transferType {
			return true
		}
	}
	return false
}
