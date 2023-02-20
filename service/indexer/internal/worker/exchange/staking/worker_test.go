package staking

import (
	"context"
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

func TestStaking_Handle(t *testing.T) {
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
			name: "RSS3 Staking Withdraw",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x2a03278590cd1962de28f9abc855cf3774fe3eb6", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xe4382ed9eac102711cc8ef53e3768b69d9196d9e479a89bd4e1d39df06a018d7
						Hash:        "0xe4382ed9eac102711cc8ef53e3768b69d9196d9e479a89bd4e1d39df06a018d7",
						BlockNumber: 16617711,
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
					zap.L().Debug("", zap.Any("transaction", transaction))

					assert.Equal(t, transaction.Platform, protocol.PlatformRSS3)
				}

				return false
			},
		},
		{
			name: "RSS3 Staking Collect",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x2a03278590cd1962de28f9abc855cf3774fe3eb6", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b
						Hash:        "0x46cf455d0500f0b0979401ef1192af508890c2f888bc0dc6c006b49f65a37b6b",
						BlockNumber: 16617713,
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
					zap.L().Debug("", zap.Any("transaction", transaction))

					assert.Equal(t, transaction.Platform, protocol.PlatformRSS3)
				}

				return false
			},
		},
		{
			name: "RSS3 Staking Deposit",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x1b861760ade296abe523c594118ef812208194ce", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f
						Hash:        "0x85ded9199c3c8a46bcb1e431e3e05f62fe014930c0a13a029a78e3a18061011f",
						BlockNumber: 16577896,
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
					zap.L().Debug("", zap.Any("transaction", transaction))

					assert.Equal(t, transaction.Platform, protocol.PlatformRSS3)
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
