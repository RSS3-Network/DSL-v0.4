package action

const (
	// transaction types: mint | transfer | burn
	TransactionMint     = "Mint"
	TransactionTransfer = "Transfer"
	TransactionBurn     = "Burn"

	// exchange types: transfer | withdraw | deposit | swap in | swap out
	ExchangeTransfer = "Transfer"
	ExchangeWithdraw = "Withdraw"
	ExchangeDeposit  = "Deposit"
	ExchangeSwap     = "Swap"

	// NFT types: transfer | mint | burn | poap
	NFTTransfer = "Transfer"
	NFTMint     = "Mint"
	NFTBurn     = "Burn"
	NFTPoap     = "POAP"

	// social types: post | comment | share (retweet) | profile | follow | unfollow | like
	SocialPost     = "Post"
	SocialComment  = "Comment"
	SocialShare    = "Share"
	SocialProfile  = "Profile"
	SocialFollow   = "Follow"
	SocialUnfollow = "Unfollow"
	SocialLike     = "Like"

	// voting types: propose | vote
	VotePropose = "Propose"
	VoteVote    = "Vote"
)
