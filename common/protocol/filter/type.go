package filter

const (
	// transaction types:  transfer | mint | burn | cancel
	TransactionTransfer = "transfer"
	TransactionMint     = "mint"
	TransactionBurn     = "burn"
	TransactionCancel   = "cancel"

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
