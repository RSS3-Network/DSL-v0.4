package partybid

const (
	Buy int = iota
	Auction
	Collection
	CollectionBatch

	BuyCrowdfundCreated                = "buy"
	CollectionBuyCrowdfundCreated      = "collection"
	RollingAuctionCrowdfundCreated     = "auction"
	CollectionBatchBuyCrowdfundCreated = "collection-batch"
)
