package benddao

const (
	Auction int = iota
	Liquidity

	LoadCreate    = "create"
	LoadRepay     = "repay"
	LoadBorrow    = "borrow"
	LoadAuction   = "bid"
	LoadLiquidate = "finalize"
)
