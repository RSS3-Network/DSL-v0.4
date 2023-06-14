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

type Example struct {
	Text string `json:"text,omitempty"`
	Hash string `json:"hash,omitempty"`
}

// Action is the value of the action field on Metadata struct.
type Action struct {
	// Name is the action name, it's used to identify the action.
	Name string `json:"name,omitempty"`

	// These fields are used for documentation.
	Examples  []Example `json:"examples,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	Platforms []string  `json:"platforms,omitempty"`
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
		Actions: []Action{{
			Examples: []Example{
				{
					Text: "Transferred 123ETH to 0xxx…xx",
					Hash: "0x558981af55e59fc42da7684a3b803b7e078b4b233d6e0be6ea22dd6b3263b98d", // Transfer
				},
				{
					Hash: "0x74a2e680f4dab8f840e52044b75890a329107faa9db5bb312284542953280d33", // Claim
				},
			},
		}},
	},
	{
		Tag:      filter.TagTransaction,
		Type:     filter.TransactionMint,
		Metadata: &metadata.Token{},
		Actions: []Action{{
			Examples: []Example{{
				Text: "Minted 123ETH",
				Hash: "0xd8acf4db6882fca667b84e44d5518a00b1afee9cd82891f1f712647aff88b1ef",
			}},
		}},
	},
	{
		Tag:      filter.TagTransaction,
		Type:     filter.TransactionBurn,
		Metadata: &metadata.Token{},
		Actions: []Action{{
			Examples: []Example{{
				Text: "Burned 123ETH",
				Hash: "0x1e31940b3b2fd992a4e65963dba2f2d51e92417c58e850e729de6dabf32342ea",
			}},
		}},
	},
	{
		Tag:  filter.TagTransaction,
		Type: filter.TransactionApproval,
		Actions: []Action{
			{
				Name: filter.ActionApprove,
				Examples: []Example{{
					Text: "Approved 1000 USDC to 0xxx…xx",
					Hash: "0x805b43d9f08352858effb9224df6ac43d1283ecaee115009707a9b2d23b722dc",
				}},
				Comment: "授权的代币 value 如果是 115792089237316195423570985008687907853269984665640564039457584007913129639935 可以换成 Infinite / Unlimited",
			},
			{
				Name: filter.ActionRevoke,
				Examples: []Example{{
					Text: "Revoked the approval of",
					Hash: "0x0f8888505c2dee133ab6574f28d1ff1e8b7854dea1a0f8ccfb9a39e3cd6a74b7",
				}},
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
				Examples: []Example{{
					Text: "Withdrew 123ETH on xxxx",
					Hash: "0x1d226709361694160082822cb0a0542aa1a45d04e54fbd00453d8400c3673705",
				}},
				Comment:   "交易所提币",
				Platforms: []string{CEX},
			},
		},
	},
	{
		Tag:      filter.TagExchange,
		Type:     filter.ExchangeDeposit,
		Metadata: &metadata.Token{},
		Actions: []Action{{
			Examples: []Example{{
				Text: "Deposited 123ETH on xxxx",
				Hash: "0x256255b6238215ffc347e2bf5e1ab5929ef428e35e3ea9d1adfd758e85f33842",
			}},
			Platforms: []string{CEX},
			Comment:   "交易所存币",
		}},
	},
	{
		Tag:      filter.TagCollectible,
		Type:     filter.CollectibleTransfer,
		Metadata: &metadata.Token{},
		Actions: []Action{{
			Examples: []Example{{
				Text: "Transferred an NFT to 0xxx…xx",
				Hash: "0x4440866ccab87ac40815f5c12fd6f824705c25931583db3534b6adeaeffb4fb9",
			}},
		}},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleAuction,
		Actions: []Action{
			{
				Name:      filter.CollectibleAuctionCreate,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora},
				Examples: []Example{{
					Text: "Created an auction on xxxx",
					Hash: "0xaa9bcf4aa0a11598610d4bb82c028dd80dad4f9740d95bf856332c642ede3b71",
				}},
			},
			{
				Name:      filter.CollectibleAuctionBid,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora, protocol.PlatformNouns},
				Examples: []Example{{
					Text: "Made a bid on xxxx",
					Hash: "0x3e55fa308f87b14f1f1abe23badf62731e547736819ed6c7e1e98c578d9c87aa",
				}},
			},
			{
				Name:      filter.CollectibleAuctionCancel,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora},
				Examples: []Example{{
					Text: "Canceled an auction on xxxx",
					Hash: "0x26e44a26e14cd20a0d4e1070d2c7bcab457c5271703d28daeac6a9f8b44eec9a",
				}},
			},
			{
				Name:      filter.CollectibleAuctionUpdate,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora},
				Examples: []Example{{
					Text: "Updated an auction on xxxx",
					Hash: "0x00d25b51a4af650b8a66421f55152383c84746639ea9a020daa2a5cabbddf555",
				}},
			},
			{
				Name:      filter.CollectibleAuctionFinalize,
				Platforms: []string{protocol.PlatformFoundation, protocol.PlatformZora, protocol.PlatformNouns},
				Examples: []Example{{
					Text: "Won an auction on xxxx",
					Hash: "0xe45780a6de45e7be064ab00a8c53a82a6637c8cbb3e0ee317f57f84b1734d31f",
				}},
			},
			{
				Name:      filter.CollectibleAuctionInvalidate,
				Platforms: []string{protocol.PlatformFoundation},
				Examples: []Example{{
					Text: "Invalidated an auction on xxxx",
					Hash: "0xb97bead4c050d14eec66a92b8e28d1bc5bdc4f12f8b8c197417f719a85b05262",
				}},
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
				Examples: []Example{{
					Text: "Bought an NFT from 0xxx…xx",
					Hash: "0xbf14c5b78920ca4c6cc7107c5901a5294fdbfaac0f2ce204f52f13e7e30f630e", // Bought"0x097c15e4ebe2cbaa8d0b565999795e5fefb71e3c5c7ef42e9ab7d92c2a25980d", // Sold
				}},
			},
		},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleMint,
		Actions: []Action{
			{
				Examples: []Example{{
					Text: "Minted an NFT",
					Hash: "0x2cf11fe89422e90d58b7d60b4c6c5b62c320df0b7c9579ad1941496971b8c393",
				}},
			},
		},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleBurn,
		Actions: []Action{
			{
				Examples: []Example{{
					Text: "Burned an NFT",
					Hash: "0x23b828aa286459d9d20915c7e4d272a59d9cf8da0495798a7613f69449bc3285",
				}},
			},
		},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleApproval,
		Actions: []Action{
			{
				Name: filter.ActionApprove,
				Examples: []Example{{
					Text: "Approved RSS3 Whitepaper collection to 0xxx…xx",
					Hash: "0xa5e6a2c53a74436bd6038c6091f26c77bf0443ef7185b65e8a00eea4657fda29",
				}},
			},
			{
				Name: filter.ActionRevoke,
				Examples: []Example{{
					Text: "Revoked the approval of from 0xxx…xx (address_to)",
					Hash: "0x6830623126117af0e853074850dda9d1e7b9acea9120a0e95e64f6fc2e6b42e4",
				}},
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
				Examples: []Example{{
					Hash: "0xfb1bce20c86df4e22ad3f9bc99843f88cd3ed10644867350a530dc4581f9bf07",
				}},
			},
			{
				Name:      filter.CollectibleEditText,
				Platforms: []string{protocol.PlatformEns},
				Examples: []Example{{
					Text: "Updated a text record for [ENS name]",
				}},
			},
			{
				Name:      filter.CollectibleEditWrap,
				Platforms: []string{protocol.PlatformEns},
				Examples: []Example{{
					Text: "Wrapped an ENS",
				}},
			},
		},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseMint,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi, protocol.PlatformCarv, protocol.PlatformMars4},
			Examples: []Example{{
				Text: "Minted an asset on xxxx",
				Hash: "0x90a0f56e53e0f5a817136288a618a78da201ce4506f825fd3e57dfe445ff5575",
			}},
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseTrade,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi, protocol.PlatformPlanetIX},
			Examples: []Example{{
				Text: "Bought an asset on xxx",
				Hash: "0x44bc70fd952f080a984a661edc9ac11e25805c8d3b1b9abc80a2a156f7b0a27f",
			}},
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseList,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi, protocol.PlatformPlanetIX},
			Examples: []Example{{
				Text: "Listed an asset on xxxx",
				Hash: "0x5de27b7666f9fde8caca7ed1ce29afe92abf2e4a446ddd5f20c824fc280fd6ec",
			}},
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseUnlist,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi, protocol.PlatformPlanetIX},
			Examples: []Example{{
				Text: "Unlisted an asset on xxxx",
				Hash: "0x499c6f89bcd1f8a18113a9d262a09c0a43998075f51cb1b98cd04ae9f4156221",
			}},
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagMetaverse,
		Type: filter.MetaverseClaim,
		Actions: []Action{{
			Platforms: []string{protocol.PlatformAavegotchi},
			Examples: []Example{{
				Text: "Burned an asset on xxx",
				Hash: "0xedb96d0420edcaa996d44df77d218d7af93b5cdaae84bdf985c9ec72afd09ce1",
			}},
		}},
		Metadata: &metadata.Token{},
	},
	{
		Tag:  filter.TagCollectible,
		Type: filter.CollectibleMusic,
		Actions: []Action{{
			Name:      filter.CollectibleMusicBuyEdition,
			Platforms: []string{protocol.PlatformSound},
			Examples: []Example{{
				Text: "Bought an NFT on xxxx",
				Hash: "0xefc1a6e1ff5159693902ebe72973ed2b1133e40b65ca91d2dce06a9845b5e30a",
			}},
		}, {
			Name:      filter.CollectibleMusicRelease,
			Platforms: []string{protocol.PlatformSound},
			Examples: []Example{{
				Text: "Released an NFT on xxxx",
				Hash: "0x6fbdd964339c937403ba22b7f37f754f48dc791d8393e535e41f783d789f6221",
			}},
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
				Examples: []Example{{
					Text: "Created a multisig transaction of 20 USDT from 0xff...ff to 0xff...ff on Gnosis Safe",
					Hash: "0x11db423456321efe84cd85a8374cf93bef65706e7a6422421a93a1f62b64d1d1",
				}},
				Comment: "add_owner remove_owner change_threshold execution rejection create",
			},
			{
				Name:      filter.ActionAddOwner,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Examples: []Example{{
					Text: "Added 0x33...22 to owners of 0xff...ff on Gnosis Safe",
					Hash: "0xfcc96dce4c316dee30870274df384ab8e91fd5aae9a5e2bdab82dd3809518096",
				}},
			},
			{
				Name:      filter.ActionRemoveOwner,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Examples: []Example{{
					Text: "Removed 0x33...22 from owners of 0xff...ff on Gnosis Safe",
					Hash: "0x17ab9f030d0cac696564b21c750280c6c15368005f32446a0ceda0c285e5fd6e",
				}},
			},
			{
				Name:      filter.ActionChangeThreshold,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Examples: []Example{{
					Text: "Changed the threshold of 0xff...ff to 3/5 on Gnosis Safe",
					Hash: "0xa1d4e12fee99247d29c521b0d70f408dc18fa0f2a954b1f526333e1bf535291e",
				}},
			},
			{
				Name:      filter.ActionRejection,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Examples: []Example{{
					Text: "Rejected a multisig transaction on Gnosis Safe of 20 USDT",
					Hash: "0x02cc01cdd0acd59c2a94ba4ff84e66f812c2c90b14bc251d8e6e973e81bb61b2",
				}},
			},
			{
				Name:      filter.ActionExecution,
				Platforms: []string{protocol.PlatformGnosisSafe},
				Examples: []Example{{
					Text: "Executed a multisig transaction on Gnosis Safe of 20 USDT",
					Hash: "0x15033f3128e0f6f21e9908039e0e7f4d9233b3e474b2cb5349a54b26f6e56944",
				}},
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
			Examples: []Example{{
				Text: "Swapped on xxxx",
				Hash: "0x0810db21380cec8f0afbf655917b2f0c64ba2035a7acf699fa5c3caa96abe298",
			}},
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
				Examples: []Example{{
					Text: "Added to liquidity on xxxx",
					Hash: "0xc52a62ad6997a4f67d5725b16b27b3fc5bf6bb4076d1c67f5877223c2b85f8a3",
				}},
				Comment: "添加流动性",
			},
			{
				Name: filter.ExchangeLiquidityRemove,
				Platforms: []string{
					protocol.PlatformBalancer, protocol.PlatformLido, protocol.PlatformPancakeswap,
					protocol.PlatformPolygonStaking, protocol.PlatformSushiswap, protocol.PlatformUniswap,
				},
				Examples: []Example{{
					Text: "Removed from liquidity on xxxx",
					Hash: "0x4501aac1bff46c1e92e5ef139feddc0751a043144590250876fe9c717a9fa460",
				}},
				Comment: "移除流动性",
			},
			{
				Name:      filter.ExchangeLiquidityCollect,
				Platforms: []string{protocol.PlatformPolygonStaking, protocol.PlatformPancakeswap},
				Examples: []Example{{
					Text: "Added to liquidity on xxxx",
					Hash: "0x4501aac1bff46c1e92e5ef139feddc0751a043144590250876fe9c717a9fa460",
				}},
				Comment: "领取收益",
			},
			{
				Name:      filter.ExchangeLiquiditySupply,
				Platforms: []string{protocol.PlatformAAVE, protocol.PlatformCurve},
				Examples: []Example{{
					Text: "Supplied liquidity on xxxx",
					Hash: "0x2c71dffa0622ad1244a65140c8b7ad2c561463d21ccffecd4d78881237762f56",
				}},
				Comment: "存款",
			},
			{
				Name:      filter.ExchangeLiquidityBorrow,
				Platforms: []string{protocol.PlatformAAVE},
				Examples: []Example{{
					Text: "Borrowed",
					Hash: "0xf4ae4547c9d8d3f248872628d6812eab3686dd74ef9ec9dcdc4b229b79920264",
				}},
				Comment: "借款",
			},
			{
				Name:      filter.ExchangeLiquidityRepay,
				Platforms: []string{protocol.PlatformAAVE},
				Examples: []Example{{
					Text: "Repaid",
					Hash: "0xb7299a906edacde17231ea28e5348b032a2442470c841773e1a296e43fb6fc29",
				}},
				Comment: "还款",
			},
			{
				Name:      filter.ExchangeLiquidityWithdraw,
				Platforms: []string{protocol.PlatformAAVE},
				Examples: []Example{{
					Text: "Withdrew ",
					Hash: "0x0b759af6219efce6dd23d1db86a26239f328f88d4bd84980b0ebeaf00d7950de",
				}},
				Comment: "提款",
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
				Examples: []Example{{
					Text: "Posted on xxxx",
					Hash: "0x000000000000000000000000108b4ccf9ced0d0975554dab0b11ef3263e5d6d6",
				}},
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
				Examples: []Example{{
					Text: "Revised a post on xxxx",
					Hash: "0x1f9586b323f432cd43a1de53a17bc655ce683327cd998334878eb262878f11a7",
				}},
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
				Examples: []Example{{
					Text: "Commented on xxxx",
					Hash: "0x549a8a2e362e647faac70c9f1595950aa35f4d2028738521320b6364f3e9823a",
				}},
			},
		},
		Metadata: &metadata.Post{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialShare,
		Actions: []Action{{
			Platforms: []string{
				protocol.PlatformCrossbell, protocol.PlatformFarcaster,
			},
			Examples: []Example{{
				Text: "Shared a post on xxxx",
				Hash: "0xc9deb029b752837d49265c83bc598d25a5301b28939822e97e12d1bb13be5a64",
			}},
		}},
		Metadata: &metadata.Post{},
	},
	{
		Tag:  filter.TagSocial,
		Type: filter.SocialMint,
		Actions: []Action{
			{
				// CrossSync, Crossbell, xLog, xSync
				Platforms: []string{protocol.PlatformCrossbell, protocol.PlatformCrossbellXLog, protocol.PlatformCrossbellXSync},
				Examples: []Example{{
					Text: "Minted a note on xxx",
					Hash: "0xa4074d5729d44fd1ad033420c6d424e8224eebd595123339eec36f732b30acb3",
				}},
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
				Examples: []Example{{
					Text: "Created a Wiki on xxxx",
					Hash: "0x6db0443bdf54852f3271791626e8617577bffa6e6e7acd8e7168abfd0c9758e8",
				}},
			},
			{
				Name:      filter.SocialRevise,
				Platforms: []string{protocol.PlatformIQWiki},
				Examples: []Example{{
					Text: "Revised a Wiki on xxxx",
					Hash: "0x3c589129bf3ea5e1339ea753d3191eb4523382e3e5bca7ead0d6d8426b31e5b3",
				}},
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
				Examples: []Example{{
					Text: `"(from==owner) Rewarded a post/
				(to==owner) Get a reward "`,
					Hash: "0x4088f9da039f07b0791e133ca3eefbca1ffe848413d110bfd8f367eaa43a7aef",
				}},
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
				Examples: []Example{{
					Text: "Appointed a proxy on xxxx",
					Hash: "0x6b7c2144e9146af7cd53cffa3b86aae97a48017c03160517a8e14d9482ba43c6",
				}},
			},
			{
				Name:      filter.SocialRemove,
				Platforms: []string{protocol.PlatformCrossbell},
				Examples: []Example{{
					Text: "Appointed a proxy on xxxx",
					Hash: "0x6b7c2144e9146af7cd53cffa3b86aae97a48017c03160517a8e14d9482ba43c6",
				}},
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
				Examples: []Example{{
					Text: "Created a profile on xxxx",
					Hash: "0x96c13d5b8191a6d530475c7d3863dffe644ba821e4aa6c91ac391173df28cd20",
				}},
			},
			{
				Name:      filter.SocialUpdate,
				Platforms: []string{protocol.PlatformCrossbell},
				Examples: []Example{{
					Text: "Updated a profile on xxx",
					Hash: "0x4d4d37a6e37affac633ca8fc1de82cd677afceffaea7060dcd4b9f20ee38e9b6",
				}},
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
				Examples: []Example{{
					Text: "Followed 0xxx…xx on xxxx",
					Hash: "0x0dbb33e0229350b37f0b8bbfef4fba6654db66ae24525525d20f72e42acc5c61",
				}},
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
				Examples: []Example{{
					Text: "Unfollowed 0xxx…xx on xxxx",
					Hash: "0xc9b8b658af20555163d073cb24f47f661b6244084aacc5f0b3a8255d4680c717",
				}},
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
				Examples: []Example{{
					Text: "Donated on xxxx",
					Hash: "0x07f0f0627ecd575844a8a74f3dadfc54789360e9cfbd027b95da8cb1f0034f7f",
				}},
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
				Examples: []Example{{
					Text: "Proposed on xxxx",
					Hash: "0xd810c4cf2f09737a6f833f1ec51eaa5504cbc0afeeb883a21a7e1c91c8a597e4",
				}},
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
				Examples: []Example{{
					Text: "Voted on xxxx",
					Hash: "qmb9avwnzshcd48p6p2u96xsbt15ydvnxfamtvcpacivwp",
				}},
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
				Examples: []Example{{
					Text: "Staked RSS3",
					Hash: "0x9ee469ae65656df6871d4d199787b65b57c85aa7073ebeb1f5bfde6e83618306",
				}},
			},
			{
				Name:      filter.ActionStakingUnstake,
				Platforms: []string{protocol.PlatformRSS3},
				Examples: []Example{{
					Text: "Unstaked RSS3",
					Hash: "0x14b943cfed9b0a4c7f81661cf74d8fdd46747395563acf34054cacac44c54eeb",
				}},
			},
			{
				Name:      filter.ActionStakingClaim,
				Platforms: []string{protocol.PlatformRSS3},
				Examples: []Example{{
					Text: "Claimed RSS3",
					Hash: "0x3e79c6731924f94f1bdb034c4816bb84ee282b00746ea02a3e477ac9744048f5",
				}},
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
