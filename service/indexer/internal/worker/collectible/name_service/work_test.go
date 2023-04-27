package name_service

import (
	"context"
	"strings"
	"sync"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
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
			name: "ens renew v1",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x19c2ec706e7245c32a6918406ca4dcc6c8a8a6fc",
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x1441ad3e2a5cc5ca26c29134afafd8191800b038057a1b28f5533244a3209842
						Hash:        "0x1441ad3e2a5cc5ca26c29134afafd8191800b038057a1b28f5533244a3209842",
						BlockNumber: 17020460,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformEns)
					assert.Equal(t, transaction.Type, filter.CollectibleEdit)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "ens renew v2",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x9ca30eaae1b6d8fe2391ef9180d7def2a6fb2b9b",
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xbf83bc960e06bb5b49ecb32e42135999410cf1a0ab1589eb7d21587f110476b9
						Hash:        "0xbf83bc960e06bb5b49ecb32e42135999410cf1a0ab1589eb7d21587f110476b9",
						BlockNumber: 17127800,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformEns)
					assert.Equal(t, transaction.Type, filter.CollectibleEdit)
					assert.Equal(t, transaction.AddressTo, strings.ToLower(ens.EnsRegistrarControllerV2.String()))
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "ens registrar v1",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x934b510d4c9103e6a87aef13b816fb080286d649",
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xeb33f04cc32907fc95e5a7772b814d29baf3b76949a44e6210e0614fae0abcfd
						Hash:        "0xeb33f04cc32907fc95e5a7772b814d29baf3b76949a44e6210e0614fae0abcfd",
						BlockNumber: 16999317,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformEns)
					assert.Equal(t, transaction.Type, filter.CollectibleMint)
					assert.Equal(t, transaction.AddressTo, strings.ToLower(ens.EnsRegistrarController.String()))
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "ens registrar v2",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x225114369fa6b133400a4bb86d08c62044387266",
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0x4f36e2f576778a5bd8509101b90877bd9961c77999503745627eac479c5acca3
						Hash:        "0x4f36e2f576778a5bd8509101b90877bd9961c77999503745627eac479c5acca3",
						BlockNumber: 17127792,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformEns)
					assert.Equal(t, transaction.Type, filter.CollectibleMint)
					assert.Equal(t, transaction.AddressTo, strings.ToLower(ens.EnsRegistrarControllerV2.String()))
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "ens wrapper",
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x934b510d4c9103e6a87aef13b816fb080286d649",
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.io/tx/0xbb73f618294dde570537b9445c6df4145402e916f00f34626b22c78541ee9004
						Hash:        "0xbb73f618294dde570537b9445c6df4145402e916f00f34626b22c78541ee9004",
						BlockNumber: 17114288,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformEns)
					assert.Equal(t, transaction.Type, filter.CollectibleEdit)
					assert.Equal(t, transaction.AddressTo, strings.ToLower(ens.EnsNameWrapper.String()))
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
