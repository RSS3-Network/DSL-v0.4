package filter

const (
	// transaction types
	TransactionTransfer string = "transfer"
	TransactionBridge   string = "bridge"
	TransactionMint     string = "mint"
	TransactionBurn     string = "burn"
	TransactionApproval string = "approval"
	TransactionMultiSig string = "multisig"

	// Actions for TransactionApproval and CollectibleApproval
	ActionApprove string = "approve"
	ActionRevoke  string = "revoke"

	ActionAddOwner        string = "add_owner"
	ActionRemoveOwner     string = "remove_owner"
	ActionChangeThreshold string = "change_threshold"
	ActionExecution       string = "execution"
	ActionRejection       string = "rejection"
	ActionCreate          string = "create"

	// exchange types
	ExchangeWithdraw  string = "withdraw"
	ExchangeDeposit   string = "deposit"
	ExchangeSwap      string = "swap"
	ExchangeLiquidity string = "liquidity"
	ExchangeStaking   string = "staking"

	ActionStakingStake   string = "stake"
	ActionStakingUnstake string = "unstake"
	ActionStakingClaim   string = "claim"

	// Bridge types
	BridgeWithdraw string = "withdraw"
	BridgeDeposit  string = "deposit"

	// actions for Exchange-Liquidity
	ExchangeLiquidityAdd      string = "add"
	ExchangeLiquidityRemove   string = "remove"
	ExchangeLiquidityCollect  string = "collect"
	ExchangeLiquiditySupply   string = "supply"
	ExchangeLiquidityBorrow   string = "borrow"
	ExchangeLiquidityRepay    string = "repay"
	ExchangeLiquidityWithdraw string = "withdraw"

	// NFT types
	CollectibleTransfer string = "transfer"
	CollectibleTrade    string = "trade"
	CollectibleMint     string = "mint"
	CollectibleBurn     string = "burn"
	CollectiblePoap     string = "poap"
	CollectibleApproval string = "approval"
	CollectibleMusic    string = "music"
	CollectibleAuction  string = "auction"
	CollectibleEdit     string = "edit"

	// actions for Collectible-Auction
	CollectibleAuctionCreate     string = "create"
	CollectibleAuctionBid        string = "bid"
	CollectibleAuctionCancel     string = "cancel"
	CollectibleAuctionFinalize   string = "finalize"
	CollectibleAuctionInvalidate string = "invalidate"
	CollectibleAuctionUpdate     string = "update"

	// actions for Collectible-Edit
	CollectibleEditRenew string = "renew"
	CollectibleEditText  string = "text"
	CollectibleEditWrap  string = "wrap"

	// sub-types for Music
	CollectibleMusicBuyEdition string = "buy"
	CollectibleMusicRelease    string = "release"
	CollectibleMusicMint       string = "mint"

	// social types
	SocialPost     string = "post"
	SocialRevise   string = "revise"
	SocialComment  string = "comment"
	SocialShare    string = "share"
	SocialProfile  string = "profile"
	SocialFollow   string = "follow"
	SocialUnfollow string = "unfollow"
	SocialLike     string = "like"
	SocialMint     string = "mint"
	SocialWiki     string = "wiki"
	SocialReward   string = "reward"
	SocialProxy    string = "proxy"

	// actions for Social-Profile
	SocialCreate  string = "create"
	SocialUpdate  string = "update"
	SocialAppoint string = "appoint"
	SocialRemove  string = "remove"

	// Governance types
	GovernancePropose string = "propose"
	GovernanceVote    string = "vote"

	// donation type
	DonationLaunch string = "launch"
	DonationDonate string = "donate"

	// metaverse type
	MetaverseMint     string = "mint"
	MetaverseTransfer string = "transfer"
	MetaverseTrade    string = "trade"
	MetaverseGift     string = "gift"
	MetaverseList     string = "list"
	MetaverseUnlist   string = "unlist"
	MetaverseClaim    string = "claim"
)
