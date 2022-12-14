package juicebox

import (
	"context"
	"sync"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
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

		ipfs.New(config.ConfigIndexer.RPC.IPFS.Internal)
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
			name: "launch",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0xc552be1ec0ad57c01303c21b7e2d6bdfd428ab65", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xfbbb8604264bde12d5616d7762dd40770ae8f2a6c1ad136ef4d5b6eee6b3bb30
						Hash:        "0xfbbb8604264bde12d5616d7762dd40770ae8f2a6c1ad136ef4d5b6eee6b3bb30",
						BlockNumber: 16170724,
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
					if !assert.Equal(t, transaction.Platform, protocol.PlatformJuiceBox) || !assert.Equal(t, transaction.Type, filter.DonationLaunch) {
						return false
					}
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "donate",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x0aa294a7a1fea425593454025bda9b02c5336b75", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x7534c63582086baeee47a2d6eb0a98a00debfe29bc6e9b15063fbfb6b9b9990f
						Hash:        "0x7534c63582086baeee47a2d6eb0a98a00debfe29bc6e9b15063fbfb6b9b9990f",
						BlockNumber: 16176607,
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
					if !assert.Equal(t, transaction.Platform, protocol.PlatformJuiceBox) || !assert.Equal(t, transaction.Type, filter.DonationDonate) {
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
