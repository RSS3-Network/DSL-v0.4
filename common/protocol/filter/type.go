package filter

const (
	// transaction types
	TransactionTransfer string = "transfer"
	TransactionBridge   string = "bridge"
	TransactionMint     string = "mint"
	TransactionBurn     string = "burn"
	TransactionApproval string = "approval"

	// actiosn for TransactionApproval and CollectibleApproval
	ActionApprove string = "approve"
	ActionRevoke  string = "revoke"

	// exchange types
	ExchangeWithdraw  string = "withdraw"
	ExchangeDeposit   string = "deposit"
	ExchangeSwap      string = "swap"
	ExchangeLiquidity string = "liquidity"

	// Bridge types
	BridgeWithdraw string = "withdraw"
	BridgeDeposit  string = "deposit"

	// sub-types for Exchange-Liquidity
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

	// sub-types for Social-Profile
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

var ValidTypeMap = map[string][]string{
	TagTransaction: {
		TransactionTransfer,
		TransactionBridge,
		TransactionMint,
		TransactionBurn,
	},
	TagExchange: {
		ExchangeWithdraw,
		ExchangeDeposit,
		ExchangeSwap,
		ExchangeLiquidity,
	},
	TagCollectible: {
		CollectibleTransfer,
		CollectibleMint,
		CollectibleBurn,
		CollectiblePoap,
	},
	TagSocial: {
		SocialPost,
		SocialRevise,
		SocialComment,
		SocialShare,
		SocialProfile,
		SocialFollow,
		SocialUnfollow,
		SocialLike,
		SocialWiki,
		SocialReward,
	},
	TagDonation: {
		DonationLaunch,
		DonationDonate,
	},
	TagGovernance: {
		GovernancePropose,
		GovernanceVote,
	},
	TagMetaverse: {
		MetaverseMint,
		MetaverseTransfer,
		MetaverseTrade,
		MetaverseGift,
		MetaverseList,
		MetaverseUnlist,
		MetaverseClaim,
	},
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
