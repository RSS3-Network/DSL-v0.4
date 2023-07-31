package marketplace

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/k0kubun/pp/v3"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/metadata_url"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

var once sync.Once

func initialize(t *testing.T) {
	once.Do(func() {
		logger, err := zap.NewDevelopment()
		assert.NoError(t, err)

		zap.ReplaceGlobals(logger)

		config.Initialize()

		ethereumClientMap, err := ethclientx.Dial(config.ConfigIndexer.RPC, protocol.EthclientNetworks)
		assert.NoError(t, err)

		for network, ethereumClient := range ethereumClientMap {
			ethclientx.ReplaceGlobal(network, ethereumClient)
		}

		metadata_url.New("https://ipfs.rss3.page/ipfs/")
	})
}

func Test_internal_Handle(t *testing.T) {
	t.Parallel()

	initialize(t)

	type arguments struct {
		ctx          context.Context
		message      *protocol.Message
		transactions []model.Transaction
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      assert.ValueAssertionFunc
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "opensea seaport",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xca5982fa4ddc6967e2a8784ae0227116fbe1a08f", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xbcac0c843560334e98a61b8c4120802a48572ba5b82edaa43d35962f85c7e97a
						Hash:        "0xbcac0c843560334e98a61b8c4120802a48572ba5b82edaa43d35962f85c7e97a",
						BlockNumber: 16095724,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformOpenSea) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "opensea seaport 2",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x409e69490da698cf2f9a928d96e769c59a5ffbbb", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x30d3791752c1d486e630b009aa83b4d9bc9a551531d6245396e5806a7e730c81
						Hash:        "0x30d3791752c1d486e630b009aa83b4d9bc9a551531d6245396e5806a7e730c81",
						BlockNumber: 16382596,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformOpenSea) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "opensea seaport sweep",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x0fefed77bb715e96f1c35c1a4e0d349563d6f6c0", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xe4690049758d03cdba2556220e85c034e44b604af5ca5b55fe0a8d2143f7873f
						Hash:        "0xe4690049758d03cdba2556220e85c034e44b604af5ca5b55fe0a8d2143f7873f",
						BlockNumber: 16393668,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformOpenSea) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "opensea seaport zero dollar purchase",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xc5bacbb46b18cd035c1063d16d32c6bdcf541de8", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x7df97b3ef1d5d5b0fe913a40488975c004c27933abbe7acdad04e5f1b6657a7a
						Hash:        "0x7df97b3ef1d5d5b0fe913a40488975c004c27933abbe7acdad04e5f1b6657a7a",
						BlockNumber: 15323187,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformOpenSea) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "OpenSea Seaport 1.4",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x8d5925baf21f25e17a365736f4190739f31bbd9f", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x5fd16ff092743fbde52afdaa770268a98abab96341462662794565ae3736ff79
						Hash:        "0x5fd16ff092743fbde52afdaa770268a98abab96341462662794565ae3736ff79",
						BlockNumber: 16866977,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformOpenSea) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "opensea wyvern exchange v1",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x085729eec537e05e173dd919939980aa652f8c2f", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x787c8681ae5c77a75877f07afe9bd7ae5676a6ad8ae45d730778798c5d2e6a3a
						Hash:        "0x787c8681ae5c77a75877f07afe9bd7ae5676a6ad8ae45d730778798c5d2e6a3a",
						BlockNumber: 14275750,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformOpenSea) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "opensea wyvern exchange v2",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x71eb086e8cb1075d9f056d09485fd0b2f7b2ec9c", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x78eec96d55d4ea07d65d95cac55aa46f51f6d4d04d9d83dc5b9b8603b5a26eee
						Hash:        "0x78eec96d55d4ea07d65d95cac55aa46f51f6d4d04d9d83dc5b9b8603b5a26eee",
						BlockNumber: 15040006,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformOpenSea) {
						return false
					}

					assert.Len(t, transaction.Transfers, 1)

					for _, transfer := range transaction.Transfers {
						var token metadata.Token
						assert.NoError(t, json.Unmarshal(transfer.Metadata, &token))

						assert.Equal(t, token.Symbol, "ENS")
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "looksrare bid",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x71d1988c74a2321a4e71b99490a4d61c72fe2c96", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c
						Hash:        "0xceac035ba1023ba84b1c3a457170bf6dd3fe892e9718699eb88c4624efd4625c",
						BlockNumber: 16529836,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformLooksRare) {
						return false
					}

					assert.Len(t, transaction.Transfers, 1)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "looksrare ask",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x53b620a822984de20bd147319747b314e2a901ba", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01
						Hash:        "0xdfa1cdcbaf519757c74bd7a01ce3eed846c8866fd944207e277422067e2cdf01",
						BlockNumber: 16095612,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformLooksRare) {
						return false
					}

					assert.Len(t, transaction.Transfers, 1)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "LooksRare Aggregator execute",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6d0267156f1c6ce44caa4bf129b76009d3d41830", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d
						Hash:        "0x6d18b33acdbfba717ffe8c11831cbe7fda5a1dd2b7e90f6d2ae3af1b503fac1d",
						BlockNumber: 16912582,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformLooksRare) {
						return false
					}

					assert.Len(t, transaction.Transfers, 1)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "quix seaport",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xe33453f0b34265a3865ee20da9472480a587ff2f", // Unknown
					Network: protocol.NetworkOptimism,
				},
				transactions: []model.Transaction{
					{
						// https://optimistic.etherscan.io/tx/0x79d5c5892e2224fa4ac2fb8379074d389681036f67946845c7c15a6f441ae76c
						Hash:        "0x79d5c5892e2224fa4ac2fb8379074d389681036f67946845c7c15a6f441ae76c",
						BlockNumber: 38911944,
						Network:     protocol.NetworkOptimism,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformQuix) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "quix exchange v5",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1", // kallydev.eth
					Network: protocol.NetworkOptimism,
				},
				transactions: []model.Transaction{
					{
						// https://optimistic.etherscan.io/tx/0x56acc166518b208370cd7aa70596f1aaa2aff2e66d8a30486ea908a8913ddd1b
						Hash:        "0x56acc166518b208370cd7aa70596f1aaa2aff2e66d8a30486ea908a8913ddd1b",
						BlockNumber: 16920252,
						Network:     protocol.NetworkOptimism,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformQuix) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "gem",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x8c227e0627dd6fcc548e34793927d658cf34d11e", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x89d5d1b2ac9a2c7e6a5f3c9ed91f1366e3488936d8c6c674056e7c19011b7b01
						Hash:        "0x89d5d1b2ac9a2c7e6a5f3c9ed91f1366e3488936d8c6c674056e7c19011b7b01",
						BlockNumber: 16095458,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					assert.Len(t, transaction.Transfers, 2)

					// This transaction purchased two NFTs from both OpenSea and LooksRare through Gem
					var hasOpenSea, hasLooksRare bool

					for _, transfer := range transaction.Transfers {
						if transfer.Platform == protocol.PlatformOpenSea {
							hasOpenSea = true
						}

						if transfer.Platform == protocol.PlatformLooksRare {
							hasLooksRare = true
						}
					}

					assert.True(t, hasOpenSea && hasLooksRare)
				}

				for _, transaction := range transactions {
					if !assert.Equal(t, transaction.Platform, protocol.PlatformGem) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "uniswap",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xe242ba8cb95b3c52a8b19c830c55a62a7bd03370", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x442e3fbb366d9aadf01b59acdf8fa1431f798fce0296c6c6633540f038ecb579
						Hash:        "0x442e3fbb366d9aadf01b59acdf8fa1431f798fce0296c6c6633540f038ecb579",
						BlockNumber: 16094957,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					assert.Len(t, transaction.Transfers, 8)
				}

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformUniswap) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "uniswap 2",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x1cb1fa7d604e06cd8c596b5b7bcaaf5c5fdefd53", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xca38b066d755cad06be0443888726452fe55db56bf09a7d8d2b8ded2388330f2
						Hash:        "0xca38b066d755cad06be0443888726452fe55db56bf09a7d8d2b8ded2388330f2",
						BlockNumber: 16131568,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.NotEmpty(t, transactions)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformUniswap) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "tofunft",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xe4d8e15066353b8dfd538cd16015103da10f71df", // Unknown
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xd537e26e9e0d8c412cb5ec4fb68a6f346531e1a5a4751cd87690a436d9433459
						Hash:        "0xd537e26e9e0d8c412cb5ec4fb68a6f346531e1a5a4751cd87690a436d9433459",
						BlockNumber: 36529583,
						Network:     protocol.NetworkPolygon,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					assert.Len(t, transaction.Transfers, 1)
				}

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformTofuNFT) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "blur",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x182f73a479de9e9e093d98487b99e0aa8e92f286", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x83fabcfe282fd2b58564f8b4cd69ad4b52ae878fe66de7f121906c0c17db0ce3
						Hash:        "0x83fabcfe282fd2b58564f8b4cd69ad4b52ae878fe66de7f121906c0c17db0ce3",
						BlockNumber: 16107791,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					assert.Len(t, transaction.Transfers, 4)
				}

				for _, transaction := range transactions {
					if !assert.Equal(t, transaction.Platform, protocol.PlatformBlur) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "Blur bulk execute",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xf293436e8b4f15715d1058d58d026fcb9121e558", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x65824b7f59e7db0ba6f9380547bdb01902b23ded76597850f8208e20538899c7
						Hash:        "0x65824b7f59e7db0ba6f9380547bdb01902b23ded76597850f8208e20538899c7",
						BlockNumber: 16865308,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					if !assert.Equal(t, transaction.Platform, protocol.PlatformBlur) {
						return false
					}

					assert.NotEmpty(t, len(transaction.Transfers))
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "element erc-721",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xb3aA403D0d57C13683B2Fd6f6989FcDfD6b2a731", // Unknown
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xec397b30dd6340eb5c704cc5a722c9385b79d93f37ffd308762e96daa1381696
						Hash:        "0xec397b30dd6340eb5c704cc5a722c9385b79d93f37ffd308762e96daa1381696",
						BlockNumber: 36573921,
						Network:     protocol.NetworkPolygon,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					assert.Len(t, transaction.Transfers, 1)
				}

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					if !assert.Equal(t, transaction.Platform, protocol.PlatformElement) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "element erc-1155",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xF09b308d15E45D5dFa6bEad7046721d5FAB5a9C7", // Unknown
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xe07d0d3fa76d7383d84ae59bae77c6ebed29851568570bee97818f7c4abfa728
						Hash:        "0xe07d0d3fa76d7383d84ae59bae77c6ebed29851568570bee97818f7c4abfa728",
						BlockNumber: 36567107,
						Network:     protocol.NetworkPolygon,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					_, _ = pp.Println(transaction)

					assert.Len(t, transaction.Transfers, 1)
				}

				for _, transaction := range transactions {
					if !assert.Equal(t, transaction.Platform, protocol.PlatformElement) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "nswap match",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xf15d6d4e89a757223708de8a1e753a20834a6a17", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xc5347a6f415b16511ac433b122aea1f10c36cd9069dde4577dbd9ebaa69a41cb
						Hash:        "0xc5347a6f415b16511ac433b122aea1f10c36cd9069dde4577dbd9ebaa69a41cb",
						BlockNumber: 16573933,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					zap.L().Info("", zap.Any("transaction", transaction))

					assert.Len(t, transaction.Transfers, 1)
				}

				for _, transaction := range transactions {
					if !assert.Equal(t, transaction.Platform, protocol.PlatformNSwap) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "zora ask",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xb4fe3175e95c862837e72befe95cb09599ec8a22", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x65b9afe83ce9868810e722217f7b71c8b03eb2bf6187736cc77a8e4475a5b453
						Hash:        "0x65b9afe83ce9868810e722217f7b71c8b03eb2bf6187736cc77a8e4475a5b453",
						BlockNumber: 16712826,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					zap.L().Info("", zap.Any("transaction", transaction))

					assert.Len(t, transaction.Transfers, 1)
				}

				for _, transaction := range transactions {
					if !assert.Equal(t, transaction.Platform, protocol.PlatformZora) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "foundation buy",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x2b615388171efa3814d0ba59c9433189d5f9c5f5", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x51f18408dc529ab73542b904befa7750424ff14dc1a85167bfe2c5764771e7cf
						Hash:        "0x51f18408dc529ab73542b904befa7750424ff14dc1a85167bfe2c5764771e7cf",
						BlockNumber: 16752548,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					zap.L().Info("", zap.Any("transaction", transaction))

					assert.Len(t, transaction.Transfers, 1)
				}

				for _, transaction := range transactions {
					if !assert.Equal(t, transaction.Platform, protocol.PlatformFoundation) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "foundation offer",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x0c5b0a9a9ceca8566b9d866eb5c34359ab3943eb", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x3cf496b22dec1a969d30bcb0a7d8e28034ae1f9e6e13fcbc832c187b857f847c
						Hash:        "0x3cf496b22dec1a969d30bcb0a7d8e28034ae1f9e6e13fcbc832c187b857f847c",
						BlockNumber: 16752515,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					zap.L().Info("", zap.Any("transaction", transaction))

					assert.Len(t, transaction.Transfers, 1)
				}

				for _, transaction := range transactions {
					if !assert.Equal(t, transaction.Platform, protocol.PlatformFoundation) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "Accept an offer on Blur",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x79C4213a328E3B4F1D87b4953C14759399dB25E2", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xc62137ad7e20dbe70b01c87fcfd03721031979f38a47278bbc24557937e9d8cf
						Hash:        "0xc62137ad7e20dbe70b01c87fcfd03721031979f38a47278bbc24557937e9d8cf",
						BlockNumber: 17391315,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Len(t, transactions, 1)

				for _, transaction := range transactions {
					zap.L().Info("", zap.Any("transaction", transaction))

					assert.Len(t, transaction.Transfers, 1)
				}

				for _, transaction := range transactions {
					if !assert.Equal(t, transaction.Platform, protocol.PlatformBlur) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			worker := New()

			assert.NotEmpty(t, worker.Name(), "worker name")
			assert.NotEmpty(t, worker.Networks(), "worker networks")
			assert.Empty(t, worker.Jobs(), "worker jobs")

			assert.NoError(t, worker.Initialize(testcase.arguments.ctx), "initialize worker")

			filledTransactions, err := ethereum.BuildTransactions(testcase.arguments.ctx, testcase.arguments.message, lo.ToSlicePtr(testcase.arguments.transactions))
			if !testcase.wantErr(t, err, "build transactions") {
				return
			}

			internalTransactions := make([]model.Transaction, 0, len(filledTransactions))

			for _, filledTransaction := range filledTransactions {
				internalTransactions = append(internalTransactions, *filledTransaction)
			}

			transactions, err := worker.Handle(testcase.arguments.ctx, testcase.arguments.message, internalTransactions)
			if !testcase.wantErr(t, err, "worker handle") {
				return
			}

			testcase.want(t, transactions, "worker handle")
		})
	}
}
