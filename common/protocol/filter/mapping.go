package filter

import (
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
)

// TagTypeCombo is a legal combination of a tag and type.
type TagTypeCombo struct {
	Tag  string
	Type string
}

// MetadataMapping is used to declare the relationship between metadata, type, and tag,
// so that we can use it to generate doc or to validate metadata, etc.
var MetadataMapping = []struct {
	Metadata      interface{}
	TagTypeCombos []TagTypeCombo
}{
	{
		&metadata.Token{}, []TagTypeCombo{
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
		&transaction.Bridge{}, []TagTypeCombo{
			{
				Tag:  TagTransaction,
				Type: TransactionBridge,
			},
		},
	},
	{
		&metadata.MultiSig{}, []TagTypeCombo{
			{
				Tag:  TagTransaction,
				Type: TransactionMultiSig,
			},
		},
	},
	{
		&metadata.Swap{}, []TagTypeCombo{
			{
				Tag:  TagExchange,
				Type: ExchangeSwap,
			},
		},
	},
	{
		&metadata.Liquidity{}, []TagTypeCombo{
			{
				Tag:  TagExchange,
				Type: ExchangeLiquidity,
			},
		},
	},
	{
		&metadata.Post{}, []TagTypeCombo{
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
		&social.Profile{}, []TagTypeCombo{
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
		&metadata.Donation{}, []TagTypeCombo{
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
		&metadata.SnapShot{}, []TagTypeCombo{
			{
				Tag:  TagGovernance,
				Type: GovernancePropose,
			},
		},
	},
	{
		&metadata.Vote{}, []TagTypeCombo{
			{
				Tag:  TagGovernance,
				Type: GovernanceVote,
			},
		},
	},
	{
		&metadata.NameService{}, []TagTypeCombo{
			{
				Tag:  TagCollectible,
				Type: CollectibleEdit,
			},
		},
	},
}

var validTypeMap = func() map[string][]string {
	typeMap := map[string][]string{}

	for _, m := range MetadataMapping {
		for _, c := range m.TagTypeCombos {
			typeMap[c.Tag] = append(typeMap[c.Tag], c.Type)
		}
	}

	return typeMap
}()

func CheckTypeValid(tag string, transferType string) bool {
	if len(tag) == 0 {
		return false
	}

	validTypeList := validTypeMap[tag]
	for _, t := range validTypeList {
		if t == transferType {
			return true
		}
	}
	return false
}
