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
// https://etherscan.io/address/0x7d655c57f71464b6f83811c55d84009cd9f5221c
//go:generate abigen --abi ./gitcoin/gitcoin.abi --pkg gitcoin --type Gitcoin --out ./gitcoin/gitcoin.go
// https://polygonscan.com/token/0xdb46d1dc155634fbc732f92e853b10b288ad5a1d
//go:generate abigen --abi ./lens/hub.abi --pkg lens --type Hub --out ./lens/hub.go
// https://scan.crossbell.io/address/0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8
//go:generate abigen --abi ./crossbell/character.abi --pkg crossbell --type Character --out ./crossbell/character.go
// https://etherscan.io/address/0x00000000006c3852cbef3e08e8df289169ede581
//go:generate abigen --abi ./opensea/seaport.abi --pkg opensea --type Seaport --out ./opensea/seaport.go
// https://etherscan.io/address/0x59728544b08ab483533076417fbbb2fd0b17ce3a
//go:generate abigen --abi ./looksrare/exchange.abi --pkg looksrare --type Exchange --out ./looksrare/exchange.go
