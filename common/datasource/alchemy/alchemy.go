package alchemy

const (
	EndpointEthereum = "eth-mainnet.g.alchemy.com"
	EndpointPolygon  = "polygon-mainnet.g.alchemy.com"
	EndpointArbitrum = "arb-mainnet.g.alchemy.com"
	EndpointOptimism = "opt-mainnet.g.alchemy.com"

	MethodGetAssetTransfers = "alchemy_getAssetTransfers"
	MethodGetNFTs           = "getNFTs"

	CategoryExternal = "external"
	CategoryInternal = "internal"
	CategoryERC20    = "erc20"
	CategoryERC721   = "erc721"
	CategoryERC1155  = "erc1155"

	MaxCount = 1000
)
