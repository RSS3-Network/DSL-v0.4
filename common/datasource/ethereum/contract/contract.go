package contract

import "embed"

//go:embed safe/GnosisSafe.abi
var ContractABIs embed.FS

// https://eips.ethereum.org/EIPS/eip-20
//go:generate abigen --abi ./erc20/erc20.abi --pkg erc20 --type ERC20 --out ./erc20/erc20.go
//https://etherscan.io/address/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
//go:generate abigen --abi ./erc20/weth/weth.abi --pkg weth --type WETH --out ./erc20/weth/weth.go
// https://eips.ethereum.org/EIPS/eip-721
//go:generate abigen --abi ./erc721/erc721.abi --pkg erc721 --type ERC721 --out ./erc721/erc721.go
// https://etherscan.io/token/0xabefbc9fd2f806065b4f3c237d4b59d9a97bcac7
//go:generate abigen --abi ./erc721/zora/zora.abi --pkg zora --type Zora --out ./erc721/zora/zora.go
// https://eips.ethereum.org/EIPS/eip-1155
//go:generate abigen --abi ./erc1155/erc1155.abi --pkg erc1155 --type ERC1155 --out ./erc1155/erc1155.go
// https://polygonscan.com/address/0x0000000000000000000000000000000000001010
//go:generate abigen --abi ./mrc20/mrc20.abi --pkg mrc20 --type MRC20 --out ./mrc20/mrc20.go
// https://etherscan.io/address/0xae461ca67b15dc8dc81ce7615e0320da1a9ab8d5
//go:generate abigen --abi ./uniswap/pool_v2.abi --pkg uniswap --type PoolV2 --out ./uniswap/pool_v2.go
// https://etherscan.io/address/0x5777d92f208679db4b9778590fa3cab3ac9e2168
//go:generate abigen --abi ./uniswap/pool_v3.abi --pkg uniswap --type PoolV3 --out ./uniswap/pool_v3.go
// https://etherscan.io/address/0xc36442b4a4522e871399cd717abdd847ab11fe88
//go:generate abigen --abi ./uniswap/position.abi --pkg uniswap --type Position --out ./uniswap/position.go
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
// https://etherscan.io/address/0x00000000006c3852cbef3e08e8df289169ede581
//go:generate abigen --abi ./opensea/wyvern_exchange.abi --pkg opensea --type WyvernExchange --out ./opensea/wyvern_exchange.go
// https://etherscan.io/address/0x59728544b08ab483533076417fbbb2fd0b17ce3a
//go:generate abigen --abi ./looksrare/exchange.abi --pkg looksrare --type Exchange --out ./looksrare/exchange.go
// https://etherscan.io/address/0xc6845a5c768bf8d7681249f8927877efda425baf
//go:generate abigen --abi ./aave/pool_v2.abi --pkg aave --type PoolV2 --alias DataTypes.ReserveConfigurationMap=DataTypes.ReserveConfigurationMapV2 --out ./aave/pool_v2.go
// https://polygonscan.com/address/0xdf9e4abdbd94107932265319479643d3b05809dc
//go:generate abigen --abi ./aave/pool_v3.abi --pkg aave --type PoolV3 --out ./aave/pool_v3.go
// https://polygonscan.com/address/0x1a13f4ca1d028320a707d99520abfefca3998b7f
//go:generate abigen --abi ./aave/token.abi --pkg aave --type Token --out ./aave/token.go
// https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1
//go:generate abigen --abi ./optimism/bridge.abi --pkg optimism --type Bridge --out ./optimism/bridge.go
// https://etherscan.io/address/0x499a865ac595e6167482d2bd5A224876baB85ab4
//go:generate abigen --abi ./polygon/predicate_ether.abi --pkg polygon --type EtherPredicate --out ./polygon/predicate_ether.go
// https://etherscan.io/address/0x40ec5b33f54e0e8a33a975908c5ba1c14e5bbbdf
//go:generate abigen --abi ./polygon/predicate_erc20.abi --pkg polygon --type ERC20Predicate --out ./polygon/predicate_erc20.go
// https://etherscan.io/address/0x9923263fA127b3d1484cFD649df8f1831c2A74e4
//go:generate abigen --abi ./polygon/predicate_erc20_mintable.abi --pkg polygon --type MintableERC20Predicate --out ./polygon/predicate_erc20_mintable.go
// https://etherscan.io/address/0x11cc04dd962e82d411587c56b815e8f8141eb7d5
//go:generate abigen --abi ./polygon/staking.abi --pkg polygon --type Staking --out ./polygon/staking.go
// https://optimistic.etherscan.io/address/0xa7078e74d1abcc57db40309cf2c680832d7cd1fc
//go:generate abigen --abi ./quix/seaport.abi --pkg quix --type SeaPort --out ./quix/seaport.go
// https://optimistic.etherscan.io/address/0x3f9da045b0f77d707ea4061110339c4ea8ecfa70
//go:generate abigen --abi ./quix/exchange_v5.abi --pkg quix --type ExchangeV5 --out ./quix/exchange_v5.go
// https://snowtrace.io/address/0xe3ffc583dc176575eea7fd9df2a7c65f7e23f4c3
//go:generate abigen --abi ./treaderjoe/router.abi --pkg treaderjoe --type Router --out ./treaderjoe/router.go
// https://polygonscan.com/address/0x581c7db44f2616781c86c331d31c1f09db87a746
// https://etherscan.io/address/0xbebc44782c7db0a1a60cb6fe97d0b483032ff1c7
//go:generate abigen --abi ./curve/3pool.abi --pkg curve --type ThreePool --out ./curve/3pool.go
// https://etherscan.io/address/0xd51a44d3fae010294c616388b506acda1bfaae46
//go:generate abigen --abi ./curve/3pool2.abi --pkg curve --type ThreePool2 --out ./curve/3pool2.go
// https://etherscan.io/address/0x90E00ACe148ca3b23Ac1bC8C240C2a7Dd9c2d7f5
//go:generate abigen --abi ./curve/registry.abi --pkg curve --type Registry --out ./curve/registry.go
// https://etherscan.io/address/0x55B916Ce078eA594c10a874ba67eCc3d62e29822
//go:generate abigen --abi ./curve/registry_exchange.abi --pkg curve --type RegistryExchange --out ./curve/registry_exchange.go
// https://etherscan.io/address/0xba12222222228d8ba445958a75a0704d566bf2c8
//go:generate abigen --abi ./balancer/vault.abi --pkg balancer --type Vault --out ./balancer/vault.go
// https://polygonscan.com/address/0x581c7db44f2616781c86c331d31c1f09db87a746#code
//go:generate abigen --abi ./dodo/machine.abi --pkg dodo --type VendingMachine --out ./dodo/machine.go
// https://etherscan.io/address/0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84
//go:generate abigen --abi ./lido/eth.abi --pkg lido --type ETH --out ./lido/eth.go
// https://etherscan.io/address/0x9c1563937145865308c8854e82f106775be28a05
//go:generate abigen --abi ./lido/matic.abi --pkg lido --type Matic --out ./lido/matic.go
// https://optimistic.etherscan.io/address/0x9c12939390052919aF3155f41Bf4160Fd3666A6f
//go:generate abigen --abi ./velodrome/router.abi --pkg velodrome --type Router --out ./velodrome/router.go

// tofuNFT
// Marketplace https://polygonscan.com/address/0x7bc8b1B5AbA4dF3Be9f9A32daE501214dC0E4f3f
//go:generate abigen --abi ./tofunft/marketplace.abi --pkg tofunft --type Marketplace --out ./tofunft/marketplace.go

// Blur
// BlurExchange https://etherscan.io/address/0x000000000000ad05ccc4f10045630fb830b95127
//go:generate abigen --abi ./blur/exchange.abi --pkg blur --type Exchange --out ./blur/exchange.go --alias execute=execute2

// Element
// ERC721OrdersFeature https://polygonscan.com/address/0xE89b615a0824286CE1cFed540f4EdDF40d2b40E3
//go:generate abigen --abi ./element/erc_721_orders_feature.abi --pkg element --type ERC721OrdersFeature --out ./element/erc_721_orders_feature.go
// ERC1155OrdersFeature https://polygonscan.com/address/0xA4807EF31298b121eF39Cd423A8Aa05A800B2Bb8
//go:generate abigen --abi ./element/erc_1155_orders_feature.abi --pkg element --type ERC1155OrdersFeature --out ./element/erc_1155_orders_feature.go

// Mask Network
// HappyTokenPool https://polygonscan.com/address/0xf9f7c1496c21bc0180f4b64dabe0754ebfc8a8c0
//go:generate abigen --abi ./masknetwork/happy_token_pool.abi --pkg masknetwork --type HappyTokenPool --out ./masknetwork/happy_token_pool.go

// Hop
// L1 bridge https://etherscan.io/address/0x3E4a3a4796d16c0Cd582C382691998f7c06420B6
//go:generate abigen --abi ./hop/l1_bridge.abi --pkg hop --type L1Bridge --out ./hop/l1_bridge.go
// L2 AMM Wapper https://polygonscan.com/address/0x9b0c694c6939b5ea9584e9b61c7815e8d97d9c63
//go:generate abigen --abi ./hop/l2_amm_wrapper.abi --pkg hop --type L2AMMWapper --out ./hop/l2_amm_wrapper.go

// NSwap
// NSwapExchange https://etherscan.io/address/0x080fa1fb48e0b1bd251348efd02c1e7a12a931ac
//go:generate abigen --abi ./nswap/NSwapExchange.abi --pkg nswap --type NSwapExchange --out ./nswap/NSwapExchange.go

// RSS3
// Staking https://etherscan.io/address/0x5301cbbedc048abac7e213184132cf982d593563
//go:generate abigen --abi ./rss3/staking.abi --pkg rss3 --type Staking --out ./rss3/staking.go

// Gnosis Safe
// Gnosis Safe V1.3 https://etherscan.io/address/0xd9db270c1b5e3bd161e8c8503c55ceabee709552
//go:generate abigen --abi ./safe/GnosisSafe.abi --pkg safe --type GnosisSafe --out ./safe/gnosisSafe.go

// ZORA
// Asks https://etherscan.io/address/0x6170b3c3a54c3d8c854934cbc314ed479b2b29a3
//go:generate abigen --abi ./zora/asks.abi --pkg zora --type Zora --out ./zora/asks.go
