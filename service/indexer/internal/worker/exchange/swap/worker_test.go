package swap

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/k0kubun/pp/v3"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
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
					assert.Equal(t, transaction.Platform, protocol.Platform1inch)
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
			name: "uniswap v3 universal router",
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
						// https://polygonscan.com/tx/0x7ccf6a6a48297105e48f038284f38297d481306cbdd7f2ecf5371daccfe7ba49
						Hash:        "0x7ccf6a6a48297105e48f038284f38297d481306cbdd7f2ecf5371daccfe7ba49",
						BlockNumber: 38694790,
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
			name: "pancake swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x36e2146ed16cfc590305a0df7f48690ec7900e50",
					Network: protocol.NetworkBinanceSmartChain,
				},
				transactions: []model.Transaction{
					{
						// https://bscscan.com/tx/0x8e957afd052e02c96db25e48abfc2a48c95f6e55a5fede6d39f6541dff5fb1b3
						Hash:        "0x8e957afd052e02c96db25e48abfc2a48c95f6e55a5fede6d39f6541dff5fb1b3",
						BlockNumber: 12850354,
						Network:     protocol.NetworkBinanceSmartChain,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				for _, transaction := range transactions {
					var swap metadata.Swap
					assert.NoError(t, json.Unmarshal(transaction.Transfers[0].Metadata, &swap))

					_, _ = pp.Println(swap)

					assert.Equal(t, swap.TokenFrom.Symbol, "WBNB")
					assert.Equal(t, swap.TokenTo.Symbol, "YOOSHI")

					assert.Equal(t, transaction.Platform, protocol.PlatformPancakeswap)
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
		{
			name: "kyberswap swap",
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
						// https://polygonscan.com/tx/0xed0024fcf5b2ad9000f3bb8853fa7bf3da3e7f4e3b36035b0e31530d5394ed87
						Hash:        "0xed0024fcf5b2ad9000f3bb8853fa7bf3da3e7f4e3b36035b0e31530d5394ed87",
						BlockNumber: 35914497,
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
					assert.Equal(t, protocol.PlatformKyberSwap, transaction.Platform)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "spookswap swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 developer
					Network: protocol.NetworkFantom,
				},
				transactions: []model.Transaction{
					{
						// https://ftmscan.com/tx/0x5b0f1101ad37a5392bbac9d399dcd5572bc3fb91e04b65c2b2163675d5ff4e5d
						Hash:        "0x5b0f1101ad37a5392bbac9d399dcd5572bc3fb91e04b65c2b2163675d5ff4e5d",
						BlockNumber: 51285797,
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
			name: "metamask swap of curve",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x443e1f9b1c866e54e914822b7d3d7165edb6e9ea", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.com/tx/0xb82a38b608d67bb345c51186e3536ef70cfe6c7f903ad844db053c8e815236b1
						Hash:        "0xb82a38b608d67bb345c51186e3536ef70cfe6c7f903ad844db053c8e815236b1",
						BlockNumber: 16033860,
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
					assert.Equal(t, len(transaction.Transfers), 1)

					var swap metadata.Swap
					assert.NoError(t, json.Unmarshal(transaction.Transfers[0].Metadata, &swap))

					assert.Equal(t, swap.TokenFrom.Symbol, "rETH")
					assert.Equal(t, swap.TokenTo.Symbol, "USDC")

					assert.Equal(t, transaction.Platform, protocol.PlatformMetamask)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "dodo swap",
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
						// https://polygonscan.com/tx/0xd505554308ad1f2912b966b8f04620b7d8b823992209ae34849d7a03df12023c
						Hash:        "0xd505554308ad1f2912b966b8f04620b7d8b823992209ae34849d7a03df12023c",
						BlockNumber: 35915422,
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
					assert.Equal(t, transaction.Platform, protocol.PlatformDODO)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "velodrome swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 developer
					Network: protocol.NetworkOptimism,
				},
				transactions: []model.Transaction{
					{
						// https://optimistic.etherscan.io/tx/0x0eb28959b6a057e273b0b0797596888fe9d59a041a3061c0cb030b0bae6e0597
						Hash:        "0x0eb28959b6a057e273b0b0797596888fe9d59a041a3061c0cb030b0bae6e0597",
						BlockNumber: 41686215,
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
				transaction := transactions[0]

				assert.Equal(t, transaction.Platform, protocol.PlatformVelodrome)

				assert.Len(t, transaction.Transfers, 1)
				transfer := transaction.Transfers[0]

				var swap metadata.Swap
				assert.NoError(t, json.Unmarshal(transfer.Metadata, &swap))

				assert.Equal(t, swap.TokenFrom.Symbol, "OP")
				assert.Equal(t, swap.TokenTo.Symbol, "VELO")

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "curve swap of frax",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x934B510D4C9103E6a87AEf13b816fb080286D649", // sujiyan.eth
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.com/tx/0x5bc300abde5761d2cadcdccddda433ed01c5fa0328d0248c4b3e601ae382007d
						Hash:        "0x5bc300abde5761d2cadcdccddda433ed01c5fa0328d0248c4b3e601ae382007d",
						BlockNumber: 15939262,
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
					assert.Equal(t, len(transaction.Transfers), 1)

					var swap metadata.Swap
					assert.NoError(t, json.Unmarshal(transaction.Transfers[0].Metadata, &swap))

					assert.Equal(t, swap.TokenFrom.Symbol, "USDT")
					assert.Equal(t, swap.TokenTo.Symbol, "USDC")

					assert.Equal(t, transaction.Platform, protocol.PlatformCurve)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "curve swap of frax 2",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x2d75e77729112bffd82e1f4a6b90a04e38159c44", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.com/tx/0x81fb92ac06b7ae7f710cadeff194002eb854760bc047d841d3e8dd8d962b3971
						Hash:        "0x81fb92ac06b7ae7f710cadeff194002eb854760bc047d841d3e8dd8d962b3971",
						BlockNumber: 16033963,
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
					assert.Equal(t, len(transaction.Transfers), 1)

					var swap metadata.Swap
					assert.NoError(t, json.Unmarshal(transaction.Transfers[0].Metadata, &swap))

					assert.Equal(t, swap.TokenFrom.Symbol, "FRAX")
					assert.Equal(t, swap.TokenTo.Symbol, "WETH")

					assert.Equal(t, transaction.Platform, protocol.PlatformCurve)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		// There is currently no way to address the issue of dApp handling fees and native token transfer events.
		//{
		//	name: "curve swap of aave pool",
		//	fields: fields{
		//		employer: shedlock.New(),
		//	},
		//	arguments: arguments{
		//		ctx: context.Background(),
		//		message: &protocol.Message{
		//			Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // Unknown
		//			Network: protocol.NetworkPolygon,
		//		},
		//		transactions: []model.Transaction{
		//			{
		//				// https://polygonscan.com/tx/0x2994f2c7663b3f1dfaef9754ed14852f1e3fbe68fd167ba390e7625d1a5d129f
		//				Hash:        "0x2994f2c7663b3f1dfaef9754ed14852f1e3fbe68fd167ba390e7625d1a5d129f",
		//				BlockNumber: 35955144,
		//				Network:     protocol.NetworkPolygon,
		//			},
		//		},
		//	},
		//	want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
		//		transactions, ok := i.([]model.Transaction)
		//		if !ok {
		//			return false
		//		}
		//
		//		for _, transaction := range transactions {
		//			assert.Equal(t, len(transaction.Transfers), 1)
		//
		//			var swap metadata.Swap
		//			assert.NoError(t, json.Unmarshal(transaction.Transfers[0].Metadata, &swap))
		//
		//			assert.Equal(t, swap.TokenFrom.Symbol, "USDC")
		//			assert.Equal(t, swap.TokenTo.Symbol, "WETH")
		//
		//			assert.Equal(t, transaction.Platform, protocol.PlatformCurve)
		//		}
		//
		//		return false
		//	},
		//	wantErr: assert.NoError,
		//},
		{
			name: "paraswap of curve pool",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x0be55326919f08af4d14a42aafb3b68f95738355", // Unknown
					Network: protocol.NetworkEthereum,
				},
				transactions: []model.Transaction{
					{
						// https://etherscan.com/tx/0x7370e1fc9d69d097cb6775badbe59f3bc523df39cd5e5444aad96423ed11c710
						Hash:        "0x7370e1fc9d69d097cb6775badbe59f3bc523df39cd5e5444aad96423ed11c710",
						BlockNumber: 16013952,
						Network:     protocol.NetworkEthereum,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Equal(t, len(transactions), 1)

				for _, transaction := range transactions {
					assert.Equal(t, len(transaction.Transfers), 1)

					var swap metadata.Swap
					assert.NoError(t, json.Unmarshal(transaction.Transfers[0].Metadata, &swap))

					assert.Equal(t, swap.TokenFrom.Symbol, "3Crv")
					assert.Equal(t, swap.TokenTo.Symbol, "USDC")

					assert.Equal(t, transaction.Platform, protocol.PlatformParaswap)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "balancer swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 Developer
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0x20b7cc317ba1a44a83d34aa0258623c598592e025541fafb7efb0295a4799418
						Hash:        "0x20b7cc317ba1a44a83d34aa0258623c598592e025541fafb7efb0295a4799418",
						BlockNumber: 36002196,
						Network:     protocol.NetworkPolygon,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Equal(t, len(transactions), 1)

				for _, transaction := range transactions {
					assert.Equal(t, len(transaction.Transfers), 1)

					var swap metadata.Swap
					assert.NoError(t, json.Unmarshal(transaction.Transfers[0].Metadata, &swap))

					assert.Equal(t, swap.TokenFrom.Symbol, "WMATIC")
					assert.Equal(t, swap.TokenTo.Symbol, "USDC")

					assert.Equal(t, transaction.Platform, protocol.PlatformBalancer)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "cow swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x6727a51caefcaf1bc189a8316ea09f844644b195", // RSS3 Developer
					Network: protocol.NetworkXDAI,
				},
				transactions: []model.Transaction{
					{
						// https://gnosisscan.io/tx/0xe2c7b25de42c95e9f1ab4c6320e9d35f6ab8736ef3a11d4e07a81b67eccd18c9
						Hash:        "0xe2c7b25de42c95e9f1ab4c6320e9d35f6ab8736ef3a11d4e07a81b67eccd18c9",
						BlockNumber: 25408287,
						Network:     protocol.NetworkXDAI,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Equal(t, len(transactions), 1)

				for _, transaction := range transactions {
					assert.Equal(t, len(transaction.Transfers), 1)

					var swap metadata.Swap
					assert.NoError(t, json.Unmarshal(transaction.Transfers[0].Metadata, &swap))

					assert.Equal(t, swap.TokenFrom.Symbol, "WXDAI")
					assert.Equal(t, swap.TokenTo.Symbol, "USDC")

					assert.Equal(t, transaction.Platform, protocol.PlatformCow)
				}

				return false
			},
			wantErr: assert.NoError,
		},
		{
			name: "mask network swap",
			fields: fields{
				employer: shedlock.New(),
			},
			arguments: arguments{
				ctx: context.Background(),
				message: &protocol.Message{
					Address: "0x790116d0685eb197b886dacad9c247f785987a4a", // Unknown
					Network: protocol.NetworkPolygon,
				},
				transactions: []model.Transaction{
					{
						// https://polygonscan.com/tx/0xd18dc543bdaceda3c83727e529538f472e58c98981f0b97a0c2f20576f839951
						Hash:        "0xd18dc543bdaceda3c83727e529538f472e58c98981f0b97a0c2f20576f839951",
						BlockNumber: 36003655,
						Network:     protocol.NetworkPolygon,
					},
				},
			},
			want: func(t assert.TestingT, i interface{}, i2 ...interface{}) bool {
				transactions, ok := i.([]model.Transaction)
				if !ok {
					return false
				}

				assert.Equal(t, len(transactions), 1)

				for _, transaction := range transactions {
					assert.Equal(t, len(transaction.Transfers), 1)

					var swap metadata.Swap
					assert.NoError(t, json.Unmarshal(transaction.Transfers[0].Metadata, &swap))

					assert.Equal(t, swap.TokenFrom.Symbol, "MASK")
					assert.Equal(t, swap.TokenTo.Symbol, "TITAN")

					assert.Equal(t, transaction.Platform, protocol.PlatformMaskNetwork)
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

			assert.NotZero(t, len(filledTransactions))

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
