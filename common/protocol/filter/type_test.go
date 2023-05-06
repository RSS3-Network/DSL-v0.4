package filter

import (
	"sort"
	"testing"

	"github.com/ysmood/got"
)

func TestValidateTypeMap(t *testing.T) {
	g := got.T(t)

	eq := func(tag string, list []string) {
		g.Helper()
		sort.Strings(validTypeMap[tag])
		sort.Strings(list)
		g.Eq(validTypeMap[tag], list)
	}

	eq(TagTransaction, []string{
		TransactionTransfer,
		TransactionBridge,
		TransactionMint,
		TransactionBurn,
		TransactionApproval,
		TransactionMultiSig,
	})

	eq(TagExchange, []string{
		ExchangeWithdraw,
		ExchangeDeposit,
		ExchangeSwap,
		ExchangeLiquidity,
	})

	eq(TagCollectible, []string{
		CollectibleTransfer,
		CollectibleAuction,
		CollectibleTrade,
		CollectibleMint,
		CollectibleBurn,
		CollectiblePoap,
		CollectibleApproval,
		CollectibleEdit,
	})

	eq(TagSocial, []string{
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
	})

	eq(TagDonation, []string{
		DonationLaunch,
		DonationDonate,
	})

	eq(TagGovernance, []string{
		GovernancePropose,
		GovernanceVote,
	})

	eq(TagMetaverse, []string{
		MetaverseMint,
		MetaverseTransfer,
		MetaverseTrade,
		MetaverseGift,
		MetaverseList,
		MetaverseUnlist,
		MetaverseClaim,
	})
}
