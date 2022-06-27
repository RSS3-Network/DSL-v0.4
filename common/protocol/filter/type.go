package filter

const (
	// transaction types: mint | transfer | burn
	TransactionMint     = "mint"
	TransactionTransfer = "transfer"
	TransactionBurn     = "burn"
	TransactionCancel   = "cancel"

	// exchange types: transfer | withdraw | deposit | swap
	ExchangeTransfer = "transfer"
	ExchangeWithdraw = "withdraw"
	ExchangeDeposit  = "deposit"
	ExchangeSwapIn   = "swap"
	ExchangeSwapOut  = "swap"

	// NFT types: transfer | mint | burn | poap
	NFTTransfer = "transfer"
	NFTMint     = "mint"
	NFTBurn     = "burn"
	NFTPoap     = "POAP"

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
	DonationDonate = "donate"
)
