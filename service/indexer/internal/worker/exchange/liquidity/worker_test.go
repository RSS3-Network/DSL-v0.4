package liquidity

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/exchange/liquidity/job/polygonstaking"
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

	job := polygonstaking.New(cache.Global())

	err := job.Run(func(ctx context.Context, duration time.Duration) error {
		return nil
	})
	assert.NoError(t, err)

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
			name: "lido eth",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xa8bA7639edA5d91424667fB4157d9E88a465f192", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.com/tx/0x474b210dd10a870b3b6268184040d661ba0bcb4903b0b6d97ae9d0d2ec82d2e5
						Hash:        "0x474b210dd10a870b3b6268184040d661ba0bcb4903b0b6d97ae9d0d2ec82d2e5",
						BlockNumber: 16030947,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				for _, transaction := range transactions {
					assert.Equal(t, transaction.Platform, protocol.PlatformLido)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "lido matic",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x55323EA97b170Eb44da6C180D96a1f23874A1eE5", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.com/tx/0xdd89e12327ceb9e46d4daa4e7d862a8cd20d89691cbd5c306addf85287b39627
						Hash:        "0xdd89e12327ceb9e46d4daa4e7d862a8cd20d89691cbd5c306addf85287b39627",
						BlockNumber: 16030691,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				for _, transaction := range transactions {
					assert.Equal(t, transaction.Platform, protocol.PlatformLido)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "polygon staking",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xe37C7e8AB0AB2754Cb9b6Ee9F0E5480b1E0d71cA", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.com/tx/0xfe7d50eba76369215de756a29419f2569300f9efb7fc2f1f5c74662998b6c7e3
						Hash:        "0xfe7d50eba76369215de756a29419f2569300f9efb7fc2f1f5c74662998b6c7e3",
						BlockNumber: 16022915,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				for _, transaction := range transactions {
					assert.Equal(t, transaction.Platform, protocol.PlatformPolygonStaking)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "balancer add",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 Developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xbb0632bad860d04649c5807a0a41ee7eb8b35d5e5f504d52a0a64f6a6d19b0e5
						Hash:        "0xbb0632bad860d04649c5807a0a41ee7eb8b35d5e5f504d52a0a64f6a6d19b0e5",
						BlockNumber: 37011846,
						Network:     protocol.NetworkPolygon,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				for _, transaction := range transactions {
					assert.Equal(t, transaction.Platform, protocol.PlatformBalancer)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "spookyswap add",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 Developer
					Network: protocol.NetworkFantom,
				},
				transactions: []model.Transaction{
					{
						// https://ftmscan.com/tx/0x738d2fe1a4853127154689ef67cf0059ba4b0a13ab2842195c6bdb550143e1cf
						Hash:        "0x738d2fe1a4853127154689ef67cf0059ba4b0a13ab2842195c6bdb550143e1cf",
						BlockNumber: 52540540,
						Network:     protocol.NetworkFantom,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				for _, transaction := range transactions {
					assert.Equal(t, transaction.Platform, protocol.PlatformSpookySwap)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "Supply USDbC at AAVE Pool V3 on Base",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xd9d187a8126f5f9808f8d05c28db8311ec3f7078", // Unknown
					Network: protocol.NetworkBase,
				},
				transactions: []model.Transaction{
					{
						// https://basescan.org/tx/0xfccf602f6bdd6065cf396cc969d58a3eae69a702d2a9216fb243620fdbe8c773
						Hash:        "0xfccf602f6bdd6065cf396cc969d58a3eae69a702d2a9216fb243620fdbe8c773",
						BlockNumber: 3779647,
						Network:     protocol.NetworkBase,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				for _, transaction := range transactions {
					assert.Equal(t, transaction.Platform, protocol.PlatformAAVE)
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
			assert.NotEmpty(t, worker.Jobs(), "worker jobs")

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
