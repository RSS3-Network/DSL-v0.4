package filter

import (
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
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

type Criteria struct {
	Tag  string
	Type string
}

// MetadataMapping is used to declare the relationship between metadata, type, and tag,
// so that we can use it to generate doc or to validate metadata, etc.
var MetadataMapping = []struct {
	Metadata interface{}
	Criteria []Criteria
}{
	{
		&metadata.Token{}, []Criteria{
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
			{
				Tag:  TagCollectible,
				Type: CollectibleEdit,
			},
			{
				Tag:  TagMetaverse,
				Type: MetaverseMint,
			},
			{
				Tag:  TagMetaverse,
				Type: MetaverseTransfer,
			},
			{
				Tag:  TagMetaverse,
				Type: MetaverseTrade,
			},
			{
				Tag:  TagMetaverse,
				Type: MetaverseGift,
			},
			{
				Tag:  TagMetaverse,
				Type: MetaverseList,
			},
			{
				Tag:  TagMetaverse,
				Type: MetaverseUnlist,
			},
			{
				Tag:  TagMetaverse,
				Type: MetaverseClaim,
			},
		},
	},
	{
		&transaction.Bridge{}, []Criteria{
			{
				Tag:  TagTransaction,
				Type: TransactionBridge,
			},
		},
	},
	{
		&metadata.MultiSig{}, []Criteria{
			{
				Tag:  TagTransaction,
				Type: TransactionMultiSig,
			},
		},
	},
	{
		&metadata.Swap{}, []Criteria{
			{
				Tag:  TagExchange,
				Type: ExchangeSwap,
			},
		},
	},
	{
		&metadata.Liquidity{}, []Criteria{
			{
				Tag:  TagExchange,
				Type: ExchangeLiquidity,
			},
		},
	},
	{
		&metadata.Post{}, []Criteria{
			{
				Tag:  TagSocial,
				Type: SocialPost,
			},
			{
				Tag:  TagSocial,
				Type: SocialRevise,
			},
			{
				Tag:  TagSocial,
				Type: SocialComment,
			},
			{
				Tag:  TagSocial,
				Type: SocialShare,
			},
			{
				Tag:  TagSocial,
				Type: SocialLike,
			},
			{
				Tag:  TagSocial,
				Type: SocialMint,
			},
			{
				Tag:  TagSocial,
				Type: SocialWiki,
			},
			{
				Tag:  TagSocial,
				Type: SocialReward,
			},
		},
	},
	{
		&social.Profile{}, []Criteria{
			{
				Tag:  TagSocial,
				Type: SocialProxy,
			},
			{
				Tag:  TagSocial,
				Type: SocialProfile,
			},
			{
				Tag:  TagSocial,
				Type: SocialFollow,
			},
			{
				Tag:  TagSocial,
				Type: SocialUnfollow,
			},
		},
	},
	{
		&metadata.Donation{}, []Criteria{
			{
				Tag:  TagDonation,
				Type: DonationLaunch,
			},
			{
				Tag:  TagDonation,
				Type: DonationDonate,
			},
		},
	},
	{
		&metadata.SnapShot{}, []Criteria{
			{
				Tag:  TagGovernance,
				Type: GovernancePropose,
			},
		},
	},
	{
		&metadata.Vote{}, []Criteria{
			{
				Tag:  TagGovernance,
				Type: GovernanceVote,
			},
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
