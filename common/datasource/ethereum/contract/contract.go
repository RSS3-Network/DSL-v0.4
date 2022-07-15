package contract

// https://eips.ethereum.org/EIPS/eip-20
//go:generate abigen --abi ./erc20/erc20.abi --pkg erc20 --type ERC20 --out ./erc20/erc20.go
//https://etherscan.io/address/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
//go:generate abigen --abi ./erc20/weth/weth.abi --pkg weth --type WETH --out ./erc20/weth/weth.go
// https://eips.ethereum.org/EIPS/eip-721
//go:generate abigen --abi ./erc721/erc721.abi --pkg erc721 --type ERC721 --out ./erc721/erc721.go
// https://eips.ethereum.org/EIPS/eip-1155
//go:generate abigen --abi ./erc1155/erc1155.abi --pkg erc1155 --type ERC1155 --out ./erc1155/erc1155.go
// https://polygonscan.com/address/0x0000000000000000000000000000000000001010
//go:generate abigen --abi ./mrc20/mrc20.abi --pkg mrc20 --type MRC20 --out ./mrc20/mrc20.go
// https://etherscan.io/address/0xae461ca67b15dc8dc81ce7615e0320da1a9ab8d5
//go:generate abigen --abi ./uniswap/pool_v2.abi --pkg uniswap --type PoolV2 --out ./uniswap/pool_v2.go
// https://etherscan.io/address/0x5777d92f208679db4b9778590fa3cab3ac9e2168
//go:generate abigen --abi ./uniswap/pool_v3.abi --pkg uniswap --type PoolV3 --out ./uniswap/pool_v3.go
// https://etherscan.io/address/0xdef1c0ded9bec7f1a1670819833240f027b25eff
//go:generate abigen --abi ./zerox/zerox.abi --pkg zerox --type ZeroX --out ./zerox/zerox.go
// https://etherscan.io/address/0x11111112542d85b3ef69ae05771c2dccff4faa26
//go:generate abigen --abi ./oneinch/oneinch.abi --pkg oneinch --type Oneinch --out ./oneinch/oneinch.go
