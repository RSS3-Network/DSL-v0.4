package multisig

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

var once sync.Once

func initialize(t *testing.T) {
	once.Do(func() {
		config.Initialize()

		cacheClient, err := cache.Dial(config.ConfigIndexer.Redis)
		assert.NoError(t, err)

		cache.ReplaceGlobal(cacheClient)

		logger, _ := zap.NewDevelopment()
		zap.ReplaceGlobals(logger)

		ethereumClientMap, err := ethclientx.Dial(config.ConfigIndexer.RPC, protocol.EthclientNetworks)
		assert.NoError(t, err)

		for network, ethereumClient := range ethereumClientMap {
			ethclientx.ReplaceGlobal(network, ethereumClient)
		}
	})
}

func Test_multisign_Handle(t *testing.T) {
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
			name: "Gnosis Safe Create",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727A51CaeFCaF1Bc189A8316eA09f844644b195", // RSS3 Developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xef2a0e78c8852527d3895e1b46a4e3b3a4f00c9611fac5cfb8ba4a1bb72d7ef1
						Hash:        "0xef2a0e78c8852527d3895e1b46a4e3b3a4f00c9611fac5cfb8ba4a1bb72d7ef1",
						BlockNumber: 39050652,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformGnosisSafe)
				}

				return false
			},
		},
		{
			name: "Gnosis Safe Add Owner",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727A51CaeFCaF1Bc189A8316eA09f844644b195", // RSS3 Developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0x4461943cd415bb520ad0f9ebea7cee7ceed2485a4d21bb2f24cad59a7bb111a4
						Hash:        "0x4461943cd415bb520ad0f9ebea7cee7ceed2485a4d21bb2f24cad59a7bb111a4",
						BlockNumber: 39096407,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformGnosisSafe)
				}

				return false
			},
		},
		{
			name: "Gnosis Safe Execution",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727A51CaeFCaF1Bc189A8316eA09f844644b195", // RSS3 Developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0x1a295dba1c904cd716a33e4870965d484e3e8ea92fa7d76f3e02d607bad1879b
						Hash:        "0x1a295dba1c904cd716a33e4870965d484e3e8ea92fa7d76f3e02d607bad1879b",
						BlockNumber: 39050792,
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
					for _, transfer := range transaction.Transfers {
						fmt.Println(string(transfer.Metadata))
					}

					assert.Equal(t, transaction.Platform, protocol.PlatformGnosisSafe)
				}

				return false
			},
		},
		{
			name: "Gnosis Safe Reject Execution",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727A51CaeFCaF1Bc189A8316eA09f844644b195", // RSS3 Developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xfaee359eda358128ab07bb5c33538bf1e0e47562f55a30131ff241342e448c30
						Hash:        "0xfaee359eda358128ab07bb5c33538bf1e0e47562f55a30131ff241342e448c30",
						BlockNumber: 39096556,
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
					for _, transfer := range transaction.Transfers {
						fmt.Println(string(transfer.Metadata))
					}

					assert.Equal(t, transaction.Platform, protocol.PlatformGnosisSafe)
				}

				return false
			},
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
			assert.NoError(t, err, "build transactions")

			assert.NotZero(t, len(filledTransactions))

			internalTransactions := make([]model.Transaction, 0, len(filledTransactions))

			for _, filledTransaction := range filledTransactions {
				internalTransactions = append(internalTransactions, *filledTransaction)
			}

			transactions, err := worker.Handle(testcase.arguments.ctx, testcase.arguments.message, internalTransactions)
			assert.NoError(t, err, "worker handle")

			testcase.want(t, transactions, "worker handle")
		})
	}
}
