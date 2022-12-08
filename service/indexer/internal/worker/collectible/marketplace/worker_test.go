package marketplace

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var once sync.Once

func initialize(t *testing.T) {
	once.Do(func() {
		config.Initialize()

		ethereumClientMap, err := ethclientx.Dial(config.ConfigIndexer.RPC, protocol.EthclientNetworks)
		assert.NoError(t, err)

		for network, ethereumClient := range ethereumClientMap {
			ethclientx.ReplaceGlobal(network, ethereumClient)
		}
	})
}

func Test_internal_Handle(t *testing.T) {
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
					Address: "0x8365ae902329bbdc1c8c5dc1618a0ac2546d00b8", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x80280a1766c141befca028166b0473bab8a0a622fa87376ca3525f5d04700451
						Hash:        "0x80280a1766c141befca028166b0473bab8a0a622fa87376ca3525f5d04700451",
						BlockNumber: 16095505,
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
					if !assert.Equal(t, transaction.Platform, protocol.PlatformTofuNFT) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
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
