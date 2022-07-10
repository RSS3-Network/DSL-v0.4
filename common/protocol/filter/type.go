package filter

const (
	// transaction types:  transfer | mint | burn | cancel
	TransactionTransfer = "transfer"
	TransactionMint     = "mint"
	TransactionBurn     = "burn"
	TransactionSelf     = "self"

	// exchange types:  withdraw | deposit | swap
	ExchangeWithdraw = "withdraw"
	ExchangeDeposit  = "deposit"
	ExchangeSwap     = "swap"

	// NFT types: transfer | mint | burn | poap
	CollectibleTrade = "trade"
	CollectibleMint  = "mint"
	CollectibleBurn  = "burn"
	CollectiblePoap  = "POAP"

	// social types: post | comment | share (retweet) | profile | follow | unfollow | like
	SocialPost     = "post"
	SocialRevise   = "revise"
	SocialComment  = "comment"
	SocialShare    = "share"
	SocialProfile  = "profile"
	SocialFollow   = "follow"
	SocialUnfollow = "unfollow"
	SocialLike     = "like"

	// Governance types: propose | vote
	GovernancePropose = "propose"
	GovernanceVote    = "vote"

	// donation type: donate
	DonationLaunch = "launch"
	DonationDonate = "donate"
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
