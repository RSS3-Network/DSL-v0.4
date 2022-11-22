package swap

import (
	"context"
	"sync"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
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

func Test_service_Handle(t *testing.T) {
	initialize(t)

	type fields struct {
		employer *shedlock.Employer
	}

	type arguments struct {
		ctx          context.Context
		message      *protocol.Message
		transactions []model.Transaction
	}

	testcases := []struct {
		name      string
		fields    fields
		arguments arguments
		want      assert.ValueAssertionFunc
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "traderjoe v2 swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // PreGod developer
					Network: protocol.NetworkAvalanche,
				},
				transactions: []model.Transaction{
					{
						// https://snowtrace.io/tx/0x8883c85ae121d7d243b913be99229a866975feca5096c7cfbeeaf691f57d0090
						Hash:        "0x8883c85ae121d7d243b913be99229a866975feca5096c7cfbeeaf691f57d0090",
						BlockNumber: 22483181,
						Network:     protocol.NetworkAvalanche,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				for _, transaction := range transactions {
					assert.Equal(t, transaction.Platform, protocol.PlatformTraderJoe)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "1inch swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // PreGod developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xb5538f83132b0a0d6ae0dc157c1adee18b3fd6db5e4a3d13cb7e89968084a040
						Hash:        "0xb5538f83132b0a0d6ae0dc157c1adee18b3fd6db5e4a3d13cb7e89968084a040",
						BlockNumber: 35711268,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformUniswap)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "uniswap v2 swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x38b0dc2d8af9f32f72372ba6955e16dc0aa369e1", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.com/tx/0x0f3bd94cf92e3807b7986e2eeecb757488139ca2735f50fdb45e02a189a78d6e
						Hash:        "0xc798307fdd83d4d052366c2644f74af5c067d95e98d95d152b6e78b19fd6227e",
						BlockNumber: 15976050,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformUniswap)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "uniswap v3 swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x000000a52a03835517e9d193b3c27626e1bc96b1", // kallydev.eth
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0x01ffb94195c10106abd060553ff742179f84954ba6a467e559e148a431bccf71
						Hash:        "0x01ffb94195c10106abd060553ff742179f84954ba6a467e559e148a431bccf71",
						BlockNumber: 29736278,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformUniswap)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "sushiswap swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0x8146a2c177a879adebb54462682ca9077ce0ced0e3a90950f373df97983c387f
						Hash:        "0x8146a2c177a879adebb54462682ca9077ce0ced0e3a90950f373df97983c387f",
						BlockNumber: 35714096,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformSushiswap)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "quickswap v2 swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0x6eb6b493925fa16e3d0f8b595cf4b5c47053d57f11ba4069880889fe1d765293
						Hash:        "0x6eb6b493925fa16e3d0f8b595cf4b5c47053d57f11ba4069880889fe1d765293",
						BlockNumber: 35913650,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformQuickSwap)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "quickswap v3 swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xb1dc5166f1b99fd69b9eebcd2ccfbdf650b70abfc302e9919c799471de56efe4
						Hash:        "0xb1dc5166f1b99fd69b9eebcd2ccfbdf650b70abfc302e9919c799471de56efe4",
						BlockNumber: 35913664,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformQuickSwap)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "rainbow swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xa01f6fd8d82f8da3103326fdf45dae9b66d92a3fbe90e390e0fe2aeeeae71a46
						Hash:        "0xa01f6fd8d82f8da3103326fdf45dae9b66d92a3fbe90e390e0fe2aeeeae71a46",
						BlockNumber: 35905238,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformRainbow)
				}

				return false
			},
			wantErr: assert.NoError,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			worker, err := New(testcase.fields.employer)
			testcase.wantErr(t, err, "new worker")

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
