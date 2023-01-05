package bridge

import (
	"context"
	"sync"
	"testing"

	"github.com/k0kubun/pp/v3"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/polygon"
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

		cacheClient, err := cache.Dial(config.ConfigIndexer.Redis)
		assert.NoError(t, err)

		cache.ReplaceGlobal(cacheClient)

		ethereumClientMap, err := ethclientx.Dial(config.ConfigIndexer.RPC, protocol.EthclientNetworks)
		assert.NoError(t, err)

		for network, ethereumClient := range ethereumClientMap {
			ethclientx.ReplaceGlobal(network, ethereumClient)
		}
	})
}

func Test_worker_Handle(t *testing.T) {
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
			name: "hop bridge weth",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6f5c23d53e03762dc84f6643d08ebe0625202a93", // Unknown
					Network: protocol.NetworkOptimism,
				},
				transactions: []model.Transaction{
					{
						// https://optimistic.etherscan.io/tx/0x797e95562d85a11b601e4692429e49df971e955dd5aa23ef5623edc45fe0adce
						Hash:        "0x797e95562d85a11b601e4692429e49df971e955dd5aa23ef5623edc45fe0adce",
						BlockNumber: 32867815,
						Network:     protocol.NetworkOptimism,
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

					assert.Equal(t, transaction.Platform, protocol.PlatformHop)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "hop bridge usdc",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6f5c23d53e03762dc84f6643d08ebe0625202a93", // Unknown
					Network: protocol.NetworkOptimism,
				},
				transactions: []model.Transaction{
					{
						// https://optimistic.etherscan.io/tx/0x309aa2dee3fab8c2a372627e3ee0c13a048e72ec0374e2e2b02ec0a03b6eaf1b
						Hash:        "0x309aa2dee3fab8c2a372627e3ee0c13a048e72ec0374e2e2b02ec0a03b6eaf1b",
						BlockNumber: 30112272,
						Network:     protocol.NetworkOptimism,
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

					assert.Equal(t, transaction.Platform, protocol.PlatformHop)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "polygon",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x35fa50af5ab137e861dacd6aab5be2b5ad003607", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x67ab44e05164ad464d79009e39c38796553f4445addcdf4cf83a6b199fa29af6
						Hash:        "0x67ab44e05164ad464d79009e39c38796553f4445addcdf4cf83a6b199fa29af6",
						BlockNumber: 15767206,
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

					assert.Equal(t, transaction.Platform, polygon.PlatformBridge)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		//{
		//	name: "optimism native",
		//	arguments: arguments{
		//		ctx: context.Background(),
		//		message: &protocol.Message{
		//			Address: "0x777f4d384bc5dc2e4ef29f235ec95fc8efff8f35", // Unknown
		//			Network: protocol.NetworkEthereum,
		//		},
		//		transactions: []model.Transaction{
		//			{
		//				// https://etherscan.io/tx/0xeec6b2c3a7ab2f47f429014de4df2ef75a7b7bef606cc387c2494b8cf28bd622
		//				Hash:        "0xeec6b2c3a7ab2f47f429014de4df2ef75a7b7bef606cc387c2494b8cf28bd622",
		//				BlockNumber: 15733462,
		//				Network:     protocol.NetworkEthereum,
		//			},
		//		},
		//	},
		//	want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
		//		transactions, ok := i.([]model.Transaction)
		//		if !ok {
		//			return false
		//		}
		//
		//		assert.Len(t, transactions, 1)
		//
		//		for _, transaction := range transactions {
		//			_, _ = pp.Println(transaction)
		//
		//			assert.Equal(t, transaction.Platform, optimism.PlatformBridge)
		//		}
		//
		//		return false
		//	},
		//	wantErr: assert.NoError,
		//},
		//{
		//	name: "arbitrum one",
		//	arguments: arguments{
		//		ctx: context.Background(),
		//		message: &protocol.Message{
		//			Address: "0xaee4a9a854a28d70cb6531a8735e40f255aca61a", // Unknown
		//			Network: protocol.NetworkEthereum,
		//		},
		//		transactions: []model.Transaction{
		//			{
		//				// https://etherscan.io/tx/0xec53d0b79d6fe33d9ab54c00aab12b8af737d3dfddd0cf87ba271089e9dc07d3
		//				Hash:        "0xec53d0b79d6fe33d9ab54c00aab12b8af737d3dfddd0cf87ba271089e9dc07d3",
		//				BlockNumber: 15773866,
		//				Network:     protocol.NetworkEthereum,
		//			},
		//		},
		//	},
		//	want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
		//		transactions, ok := i.([]model.Transaction)
		//		if !ok {
		//			return false
		//		}
		//
		//		assert.Len(t, transactions, 1)
		//
		//		for _, transaction := range transactions {
		//			_, _ = pp.Println(transaction)
		//
		//			assert.Equal(t, transaction.Platform, arbitrum.PlatformInboxOne)
		//		}
		//
		//		return false
		//	},
		//	wantErr: assert.NoError,
		//},
		//{
		//	name: "arbitrum nova",
		//	arguments: arguments{
		//		ctx: context.Background(),
		//		message: &protocol.Message{
		//			Address: "0xaee4a9a854a28d70cb6531a8735e40f255aca61a", // Unknown
		//			Network: protocol.NetworkEthereum,
		//		},
		//		transactions: []model.Transaction{
		//			{
		//				// https://etherscan.io/tx/0xa8b2d5c889fbc233a9e7730129364959b05c0edc9048a3d1e997abd392cbb271
		//				Hash:        "0xa8b2d5c889fbc233a9e7730129364959b05c0edc9048a3d1e997abd392cbb271",
		//				BlockNumber: 15773792,
		//				Network:     protocol.NetworkEthereum,
		//			},
		//		},
		//	},
		//	want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
		//		transactions, ok := i.([]model.Transaction)
		//		if !ok {
		//			return false
		//		}
		//
		//		assert.Len(t, transactions, 1)
		//
		//		for _, transaction := range transactions {
		//			_, _ = pp.Println(transaction)
		//
		//			assert.Equal(t, transaction.Platform, arbitrum.PlatformInboxNova)
		//		}
		//
		//		return false
		//	},
		//	wantErr: assert.NoError,
		//},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			worker := New()

			assert.NotEmpty(t, worker.Name(), "worker name")
			assert.NotEmpty(t, worker.Networks(), "worker networks")
			assert.Empty(t, worker.Jobs(), "worker jobs")

			assert.NoError(t, worker.Initialize(testcase.arguments.ctx), "initialize worker")

			filledTransactions, err := ethereum.BuildTransactions(testcase.arguments.ctx, testcase.arguments.message, lo.ToSlicePtr(testcase.arguments.transactions))
			testcase.wantErr(t, err, "build transactions")

			internalTransactions := make([]model.Transaction, 0, len(filledTransactions))

			for _, filledTransaction := range filledTransactions {
				internalTransactions = append(internalTransactions, *filledTransaction)
			}

			transactions, err := worker.Handle(testcase.arguments.ctx, testcase.arguments.message, internalTransactions)
			testcase.wantErr(t, err, "worker handle")

			testcase.want(t, transactions, "worker handle")
		})
	}
}
