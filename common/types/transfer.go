package types

import (
	"encoding/json"
	"reflect"

	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

// TODO: expand this to a proper list
const CEX = "CEX"

// Action is the value of the action field on Metadata struct.
type Action struct {
	// Name is the action name, it's used to identify the action.
	Name string `json:"name,omitempty"`

	// These fields are used for documentation.
	Example   string   `json:"example,omitempty"`
	Comment   string   `json:"comment,omitempty"`
	Platforms []string `json:"platforms,omitempty"`
}

// MetadataMapping is used to declare the relationship between metadata, type, and tag,
// so that we can use it to generate doc or to validate metadata, etc.
type TransferType struct {
	Tag      string
	Type     string
	Actions  []Action
	Metadata interface{}
}

var _ json.Marshaler = (*TransferType)(nil)

func (tt *TransferType) MarshalJSON() ([]byte, error) {
	t := reflect.TypeOf(tt.Metadata).Elem()
	return json.Marshal(map[string]interface{}{
		"tag":      tt.Tag,
		"type":     tt.Type,
		"actions":  tt.Actions,
		"metadata": t.PkgPath() + "." + t.Name(),
	})
}

var TransferTypes = []TransferType{
	{
		Tag:      filter.TagTransaction,
		Type:     filter.TransactionTransfer,
		Metadata: &metadata.Token{},
		Actions: []Action{
			{
				Example: "Transferred 123ETH to 0xxx…xx",
			},
		},
	},
	{
		Tag:      filter.TagTransaction,
		Type:     filter.TransactionMint,
		Metadata: &metadata.Token{},
		Actions: []Action{
			{
				Example: "Minted 123ETH",
			},
		},
	},
	{
		Tag:      filter.TagTransaction,
		Type:     filter.TransactionBurn,
		Metadata: &metadata.Token{},
		Actions: []Action{
			{
				Example: "Burned 123ETH",
			},
		},
	},
	{
		Tag:  filter.TagTransaction,
		Type: filter.TransactionApproval,
		Actions: []Action{
			{
				Name:    filter.ActionApprove,
				Example: "Approved 1000 USDC to 0xxx…xx",
				Comment: "授权的代币 value 如果是 115792089237316195423570985008687907853269984665640564039457584007913129639935 可以换成 Infinite / Unlimited",
			},
			{
				Name:    filter.ActionRevoke,
				Example: "Revoked the approval of",
			},
		},
		Metadata: &metadata.Token{},
	},
	{
		Tag:      filter.TagExchange,
		Type:     filter.ExchangeWithdraw,
		Metadata: &metadata.Token{},
		Actions: []Action{
			{
				Example:   "Withdrew 123ETH on xxxx",
				Comment:   "交易所提币",
				Platforms: []string{CEX},
			},
		},
	},
	{
		Tag:      filter.TagExchange,
		Type:     filter.ExchangeDeposit,
		Metadata: &metadata.Token{},
		Actions: []Action{
			{
				Example:   "Deposited 123ETH on xxxx",
				Platforms: []string{CEX},
				Comment:   "交易所存币",
			},
		},
	},
	{
		Tag:      filter.TagCollectible,
		Type:     filter.CollectibleTransfer,
		Metadata: &metadata.Token{},
		Actions: []Action{
			{
				Example: "Transferred an NFT to 0xxx…xx",
			},
		},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleAuction,
		Actions: []Action{
			{
				Name:      filter.CollectibleAuctionCreate,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora},
				Example:   "Created an auction on xxxx",
			},
			{
				Name:      filter.CollectibleAuctionBid,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora, protocol.PlatformNouns},
				Example:   "Made a bid on xxxx",
			},
			{
				Name:      filter.CollectibleAuctionCancel,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora},
				Example:   "Canceled an auction on xxxx",
			},
			{
				Name:      filter.CollectibleAuctionUpdate,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora},
				Example:   "Updated an auction on xxxx",
			},
			{
				Name:      filter.CollectibleAuctionFinalize,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora, protocol.PlatformNouns},
				Example:   "Won an auction on xxxx",
			},
			{
				Name:      filter.CollectibleAuctionInvalidate,
				Platforms: []string{protocol.PlatformFoundation},
				Example:   "Invalidated an auction on xxxx",
			},
		},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleTrade,
		Actions: []Action{
			{
				Platforms: []string{
					protocol.PlatformBlur, protocol.PlatformElement, protocol.PlatformGem,
					protocol.PlatformLooksRare, protocol.PlatformNSwap, protocol.PlatformNouns, protocol.PlatformOpenSea,
					protocol.PlatformQuix, protocol.PlatformTofuNFT, protocol.PlatformUniswap, protocol.PlatformZora,
				},
				Example: "Bought an NFT from 0xxx…xx",
			},
		},
		Metadata: &metadata.Token{},
	},
	{
		Tag:      filter.TagCollectible,
		Type:     filter.CollectibleMint,
		Actions:  []Action{{Example: "Minted an NFT"}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:      filter.TagCollectible,
		Type:     filter.CollectibleBurn,
		Actions:  []Action{{Example: "Burned an NFT"}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleApproval,
		Actions: []Action{
			{
				Name:    filter.ActionApprove,
				Example: "Approved RSS3 Whitepaper collection to 0xxx…xx",
			},
			{
				Name:    filter.ActionRevoke,
				Example: "Revoked the approval of from 0xxx…xx (address_to)",
			},
		},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleEdit,
		Actions: []Action{
			{
				Name:      filter.CollectibleEditRenew,
				Platforms: []string{protocol.PlatformEns},
			},
			{
				Name:      filter.CollectibleEditText,
				Platforms: []string{protocol.PlatformEns},
				Example:   "Updated a text record for [ENS name]",
			},
			{
				Name:      filter.CollectibleEditWrap,
				Platforms: []string{protocol.PlatformEns},
				Example:   "Wrapped an ENS",
			},
		},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseMint,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi, protocol.PlatformCarv, protocol.PlatformMars4},
			Example:   "Minted an asset on xxxx",
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseTrade,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi, protocol.PlatformPlanetIX},
			Example:   "Bought an asset on xxx",
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseList,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi, protocol.PlatformPlanetIX},
			Example:   "Listed an asset on xxxx",
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseUnlist,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi, protocol.PlatformPlanetIX},
			Example:   "Unlisted an asset on xxxx",
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseClaim,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi},
			Example:   "Burned an asset on xxx",
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleMusic,
		Actions: []Action{{
			Name:      filter.CollectibleMusicBuyEdition,
			Platforms: []string{protocol.PlatformSound},
			Example:   "Bought an NFT on xxxx",
		}, {
			Name:      filter.CollectibleMusicRelease,
			Platforms: []string{protocol.PlatformSound},
			Example:   "Released an NFT on xxxx",
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagTransaction,
		Type: filter.TransactionMultiSig,
		Actions: []Action{
			{
				Name:      filter.ActionCreate,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Example:   "Created a multisig transaction of 20 USDT from 0xff...ff to 0xff...ff on Gnosis Safe",
				Comment:   "add_owner remove_owner change_threshold execution rejection create",
			},
			{
				Name:      filter.ActionAddOwner,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Example:   "Added 0x33...22 to owners of 0xff...ff on Gnosis Safe",
			},
			{
				Name:      filter.ActionRemoveOwner,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Example:   "Removed 0x33...22 from owners of 0xff...ff on Gnosis Safe",
			},
			{
				Name:      filter.ActionChangeThreshold,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Example:   "Changed the threshold of 0xff...ff to 3/5 on Gnosis Safe",
			},
			{
				Name:      filter.ActionRejection,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Example:   "Rejected a multisig transaction on Gnosis Safe of 20 USDT",
			},
			{
				Name:      filter.ActionExecution,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Example:   "Executed a multisig transaction on Gnosis Safe of 20 USDT",
			},
		},
		Metadata: &metadata.MultiSig{},
	},
	{
		Tag:  filter.TagExchange,
		Type: filter.ExchangeSwap,
		Actions: []Action{{
			Platforms: []string{
				protocol.Platform0x, protocol.Platform1inch, protocol.PlatformCow,
				protocol.PlatformCurve, protocol.PlatformDODO, protocol.PlatformKyberSwap,
				protocol.PlatformMetamask, protocol.PlatformPancakeswap, protocol.PlatformParaswap,
				protocol.PlatformQuickSwap, protocol.PlatformRainbow, protocol.PlatformSpookySwap,
				protocol.PlatformSushiswap, protocol.PlatformTraderJoe, protocol.PlatformUniswap,
				protocol.PlatformVelodrome, protocol.PlatformZerion,
			},
			Example: "Swapped on xxxx ",
		}},
		Metadata: &metadata.Swap{},
	},
	{
		Tag:  filter.TagExchange,
		Type: filter.ExchangeLiquidity,
		Actions: []Action{
			{
				Name: filter.ExchangeLiquidityAdd,
				Platforms: []string{
					protocol.PlatformBalancer, protocol.PlatformLido, protocol.PlatformPancakeswap,
					protocol.PlatformPolygonStaking, protocol.PlatformSushiswap, protocol.PlatformUniswap,
				},
				Example: "Added to liquidity on xxxx",
				Comment: "添加流动性",
			},
			{
				Name: filter.ExchangeLiquidityRemove,
				Platforms: []string{
					protocol.PlatformBalancer, protocol.PlatformLido, protocol.PlatformPancakeswap,
					protocol.PlatformPolygonStaking, protocol.PlatformSushiswap, protocol.PlatformUniswap,
				},
				Example: "Removed from liquidity on xxxx",
				Comment: "移除流动性",
			},
			{
				Name:      filter.ExchangeLiquidityCollect,
				Platforms: []string{protocol.PlatformPolygonStaking, protocol.PlatformPancakeswap},
				Example:   "Added to liquidity on xxxx",
				Comment:   "领取收益",
			},
			{
				Name:      filter.ExchangeLiquiditySupply,
				Platforms: []string{protocol.PlatformAAVE, protocol.PlatformCurve},
				Example:   "Supplied liquidity on xxxx",
				Comment:   "存款",
			},
			{
				Name:      filter.ExchangeLiquidityBorrow,
				Platforms: []string{protocol.PlatformAAVE},
				Example:   "Borrowed",
				Comment:   "借款",
			},
			{
				Name:      filter.ExchangeLiquidityRepay,
				Platforms: []string{protocol.PlatformAAVE},
				Example:   "Repaid",
				Comment:   "还款",
			},
			{
				Name:      filter.ExchangeLiquidityWithdraw,
				Platforms: []string{protocol.PlatformAAVE},
				Example:   "Withdrew ",
				Comment:   "提款",
			},
		},
		Metadata: &metadata.Liquidity{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialPost,
		Actions: []Action{
			{
				Platforms: []string{
					protocol.PlatformCrossbell, protocol.PlatformFarcaster,
					protocol.PlatformLensLenster, protocol.PlatformLensLenstube, protocol.PlatformMirror,
					protocol.PlatformLensOrb, protocol.PlatformCrossbellXCast, protocol.PlatformCrossbellXLog,
					protocol.PlatformCrossbellXSync,
				},
				Example: "Posted on xxxx",
			},
		},
		Metadata: &metadata.Post{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialRevise,
		Actions: []Action{
			{
				Platforms: []string{
					protocol.PlatformCrossbell, protocol.PlatformMirror,
				},
				Example: "Revised a post on xxxx",
			},
		},
		Metadata: &metadata.Post{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialComment,
		Actions: []Action{
			{
				Platforms: []string{protocol.PlatformFarcaster, protocol.PlatformRara},
				Example:   "Commented on xxxx",
			},
		},
		Metadata: &metadata.Post{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialShare,
		Actions: []Action{
			{
				Platforms: []string{
					protocol.PlatformCrossbell, protocol.PlatformFarcaster,
				},
				Example: "Shared a post on xxxx",
			},
		},
		Metadata: &metadata.Post{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialMint,
		Actions: []Action{
			{
				// CrossSync, Crossbell, xLog, xSync
				Platforms: []string{protocol.PlatformCrossbell, protocol.PlatformCrossbellXLog, protocol.PlatformCrossbellXSync},
				Example:   "Minted a note on xxx",
			},
		},
		Metadata: &metadata.Post{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialWiki,
		Actions: []Action{
			{
				Name:      filter.SocialCreate,
				Platforms: []string{protocol.PlatformIQWiki},
				Example:   "Created a Wiki on xxxx",
			},
			{
				Name:      filter.SocialRevise,
				Platforms: []string{protocol.PlatformIQWiki},
				Example:   "Revised a Wiki on xxxx",
			},
		},
		Metadata: &metadata.Post{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialReward,
		Actions: []Action{
			{
				Platforms: []string{protocol.PlatformMatters},
				Example: `"(from==owner) Rewarded a post/
				(to==owner) Get a reward "`,
			},
		},
		Metadata: &metadata.Post{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialProxy,
		Actions: []Action{
			{
				Name:      filter.SocialAppoint,
				Platforms: []string{protocol.PlatformCrossbell},
				Example:   "Appointed a proxy on xxxx",
			},
			{
				Name:      filter.SocialRemove,
				Platforms: []string{protocol.PlatformCrossbell},
				Example:   "Appointed a proxy on xxxx",
			},
		},
		Metadata: &social.Profile{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialProfile,
		Actions: []Action{
			{
				Name:      filter.SocialCreate,
				Platforms: []string{protocol.PlatformCrossbell, protocol.PlatformLens},
				Example:   "Created a profile on xxxx",
			},
			{
				Name:      filter.SocialUpdate,
				Platforms: []string{protocol.PlatformCrossbell},
				Example:   "Updated a profile on xxx",
			},
		},
		Metadata: &social.Profile{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialFollow,
		Actions: []Action{
			{
				Platforms: []string{protocol.PlatformCrossbell, protocol.PlatformLens},
				Example:   "Followed 0xxx…xx on xxxx",
			},
		},
		Metadata: &social.Profile{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialUnfollow,
		Actions: []Action{
			{
				Platforms: []string{protocol.PlatformCrossbell, protocol.PlatformLens},
				Example:   "Unfollowed 0xxx…xx on xxxx",
			},
		},
		Metadata: &social.Profile{},
	},
	{
		Tag:  filter.TagDonation,
		Type: filter.DonationDonate,
		Actions: []Action{
			{
				Platforms: []string{protocol.PlatformGitcoin},
				Example:   "Donated on xxxx",
			},
		},
		Metadata: &metadata.Donation{},
	},
	{
		Tag:  filter.TagGovernance,
		Type: filter.GovernancePropose,
		Actions: []Action{
			{
				Platforms: []string{protocol.PlatformSnapshot},
				Example:   "Proposed on xxxx",
			},
		},
		Metadata: &metadata.SnapShot{},
	},
	{
		Tag:  filter.TagGovernance,
		Type: filter.GovernanceVote,
		Actions: []Action{
			{
				Platforms: []string{protocol.PlatformSnapshot},
				Example:   "Voted on xxxx",
			},
		},
		Metadata: &metadata.Vote{},
	},
	{
		Tag:  filter.TagExchange,
		Type: filter.ExchangeStaking,
		Actions: []Action{
			{
				Name:      filter.ActionStakingStake,
				Platforms: []string{protocol.PlatformRSS3},
				Example:   "Staked RSS3",
			},
			{
				Name:      filter.ActionStakingUnstake,
				Platforms: []string{protocol.PlatformRSS3},
				Example:   "Unstaked RSS3",
			},
			{
				Name:      filter.ActionStakingClaim,
				Platforms: []string{protocol.PlatformRSS3},
				Example:   "Claimed RSS3",
			},
		},
		Metadata: &transaction.Staking{},
	},
	{
		Tag:  filter.TagTransaction,
		Type: filter.TransactionBridge,
		Actions: []Action{
			{
				Name: filter.BridgeDeposit,
			},
			{
				Name: filter.BridgeWithdraw,
			},
		},
		Metadata: &transaction.Bridge{},
	},
}

var validTypeMap = func() map[string][]string {
	typeMap := map[string][]string{}

	for _, t := range TransferTypes {
		typeMap[t.Tag] = append(typeMap[t.Tag], t.Type)
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
