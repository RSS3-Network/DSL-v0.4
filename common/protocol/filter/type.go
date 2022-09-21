package filter

const (
	// transaction types
	TransactionTransfer string = "transfer"
	TransactionMint     string = "mint"
	TransactionBurn     string = "burn"

	// exchange types
	ExchangeWithdraw  string = "withdraw"
	ExchangeDeposit   string = "deposit"
	ExchangeSwap      string = "swap"
	ExchangeLiquidity string = "liquidity"

	// sub-types for Exchange-Liquidity
	ExchangeLiquidityAdd      string = "add"
	ExchangeLiquidityRemove   string = "remove"
	ExchangeLiquidityCollect  string = "collect"
	ExchangeLiquiditySupply   string = "supply"
	ExchangeLiquidityBorrow   string = "borrow"
	ExchangeLiquidityRepay    string = "repay"
	ExchangeLiquidityWithdraw string = "withdraw"

	// NFT types
	CollectibleTransfer  string = "transfer"
	CollectibleTrade     string = "trade"
	CollectibleMint      string = "mint"
	CollectibleBurn      string = "burn"
	CollectiblePoap      string = "poap"
	CollectibleCrowdFund string = "crowdfund"

	PartyBid        string = "bid"
	PartyBuy        string = "buy"
	PartyCollection string = "collection"

	// sub-types for CollectibleCrowdFunding
	PartyBidStart      string = "start"
	PartyBidContribute string = "contribute"
	PartyBidBid        string = "bid"
	PartyBidBuy        string = "buy"
	PartyBidFinalize   string = "finalize"
	PartyBidClaim      string = "claim"
	PartyBidExpire     string = "expire"

	// social types
	SocialPost     string = "post"
	SocialRevise   string = "revise"
	SocialComment  string = "comment"
	SocialShare    string = "share"
	SocialProfile  string = "profile"
	SocialFollow   string = "follow"
	SocialUnfollow string = "unfollow"
	SocialLike     string = "like"

	// sub-types for Social-Profile
	SocialProfileCreate string = "create"
	SocialProfileUpdate string = "update"

	// Governance types
	GovernancePropose string = "propose"
	GovernanceVote    string = "vote"

	// donation type
	DonationLaunch string = "launch"
	DonationDonate string = "donate"
)

var ValidTypeMap = map[string][]string{
	TagTransaction: {TransactionTransfer, TransactionMint, TransactionBurn},
	TagExchange:    {ExchangeWithdraw, ExchangeDeposit, ExchangeSwap, ExchangeLiquidity},
	TagCollectible: {CollectibleTransfer, CollectibleMint, CollectibleBurn, CollectiblePoap, CollectibleCrowdFund},
	TagSocial:      {SocialPost, SocialRevise, SocialComment, SocialShare, SocialProfile, SocialFollow, SocialUnfollow, SocialLike},
	TagDonation:    {DonationLaunch, DonationDonate},
	TagGovernance:  {GovernancePropose, GovernanceVote},
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
