package liquidity

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/database/model"
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
			name: "lido submit",
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
					for _, transfer := range transaction.Transfers {
						fmt.Println(string(transfer.Metadata))
					}
				}

				for _, transaction := range transactions {
					assert.Equal(t, transaction.Platform, protocol.PlatformLido)
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
