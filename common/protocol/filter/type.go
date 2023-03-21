package filter

import (
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
)

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

var ValidTypeMap = map[string][]string{
	TagTransaction: {
		TransactionTransfer,
		TransactionBridge,
		TransactionMint,
		TransactionBurn,
		TransactionApproval,
		TransactionMultiSig,
	},
	TagExchange: {
		ExchangeWithdraw,
		ExchangeDeposit,
		ExchangeSwap,
		ExchangeLiquidity,
	},
	TagCollectible: {
		CollectibleTransfer,
		CollectibleAuction,
		CollectibleTrade,
		CollectibleMint,
		CollectibleBurn,
		CollectiblePoap,
		CollectibleApproval,
		CollectibleEdit,
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
		SocialMint,
		SocialWiki,
		SocialReward,
		SocialProxy,
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

// MetadataTypeMap is used to declare the relationship between metadata, type, and tag,
// so that we can use it to generate doc or to validate metadata, etc.
var MetadataTypeMap = map[interface{}][]struct {
	Tag  string
	Type string
}{
	&metadata.Token{}: {
		{
			Tag:  TagTransaction,
			Type: TransactionTransfer,
		},
		{
			Tag:  TagTransaction,
			Type: TransactionMint,
		},
		{
			Tag:  TagTransaction,
			Type: TransactionBurn,
		},
		{
			Tag:  TagTransaction,
			Type: TransactionApproval,
		},
		{
			Tag:  TagExchange,
			Type: ExchangeWithdraw,
		},
		{
			Tag:  TagExchange,
			Type: ExchangeDeposit,
		},
		{
			Tag:  TagCollectible,
			Type: CollectibleTransfer,
		},
		{
			Tag:  TagCollectible,
			Type: CollectibleAuction,
		},
		{
			Tag:  TagCollectible,
			Type: CollectibleTrade,
		},
		{
			Tag:  TagCollectible,
			Type: CollectibleMint,
		},
		{
			Tag:  TagCollectible,
			Type: CollectibleBurn,
		},
		{
			Tag:  TagCollectible,
			Type: CollectiblePoap,
		},
		{
			Tag:  TagCollectible,
			Type: CollectibleApproval,
		},
	},
	&transaction.Bridge{}: {
		{
			Tag:  TagTransaction,
			Type: TransactionBridge,
		},
	},
	&metadata.MultiSig{}: {
		{
			Tag:  TagTransaction,
			Type: TransactionMultiSig,
		},
	},
	&metadata.Swap{}: {
		{
			Tag:  TagExchange,
			Type: ExchangeSwap,
		},
	},
	&metadata.Liquidity{}: {
		{
			Tag:  TagExchange,
			Type: ExchangeLiquidity,
		},
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
