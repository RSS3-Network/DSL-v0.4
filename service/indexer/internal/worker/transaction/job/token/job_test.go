package token

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/stretchr/testify/assert"

	"github.com/naturalselectionlabs/pregod/common/datasource/coingecko"
	"github.com/naturalselectionlabs/pregod/common/worker/zksync"

	"go.uber.org/ratelimit"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

var once sync.Once

func initialize(t *testing.T) {
	once.Do(func() {
		config.Initialize()

		logger, err := zap.NewDevelopment()
		assert.NoError(t, err, "initialize logger")

		zap.ReplaceGlobals(logger)

		ethereumClientMap, err := ethclientx.Dial(config.ConfigIndexer.RPC, protocol.EthclientNetworks)
		assert.NoError(t, err, "dial ethereum rpc endpoint")

		for network, ethereumClient := range ethereumClientMap {
			ethclientx.ReplaceGlobal(network, ethereumClient)
		}

		databaseClient, err := database.Dial(config.ConfigIndexer.Postgres.String(), true)
		assert.NoError(t, err, "dial database")

		database.ReplaceGlobal(databaseClient)
	})
}

func TestNew(t *testing.T) {
	initialize(t)

	testcases := []struct {
		name string
		want assert.ValueAssertionFunc
	}{
		{
			name: "default",
			want: assert.NotNil,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			job := New()

			testcase.want(t, job, fmt.Sprintf("nil job"))

			assert.NotEmpty(t, job.Name(), "empty name")
			assert.NotEmpty(t, job.Spec(), "empty spec")
			assert.NotZero(t, job.Timeout(), "zero timeout")
		})
	}
}

func TestJob_updateTokenListFromCoinGecko(t *testing.T) {
	initialize(t)

	type fields struct {
		coingeckoClient *coingecko.Client
		zksyncClient    *zksync.Client
		databaseClient  *gorm.DB
		rateLimiter     ratelimit.Limiter
	}

	type arguments struct {
		ctx      context.Context
		coinList []coingecko.Coin
	}

	testcases := []struct {
		name      string
		fields    fields
		arguments arguments
		want      assert.ErrorAssertionFunc
	}{
		{
			name: "default",
			fields: fields{
				coingeckoClient: coingecko.New(),
				databaseClient:  database.Global(),
				rateLimiter:     ratelimit.New(50, ratelimit.Per(time.Minute)),
			},
			arguments: arguments{
				ctx: context.Background(),
				coinList: []coingecko.Coin{
					{
						ID:     "weth",
						Name:   "Wrapped Ether",
						Symbol: "WETH",
						Platforms: map[string]string{
							"ethereum":    "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
							"polygon-pos": "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619",
							"xdai":        "0x6a023ccd1ff6f2045c3309768ead9e68f978f6e1",
							"avalanche":   "0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab",
						},
						Image: &coingecko.CoinImage{
							Large: "https://assets.coingecko.com/coins/images/2518/thumb/weth.png?1595348880",
						},
					},
				},
			},
			want: assert.NoError,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			job := &Job{
				coingeckoClient: testcase.fields.coingeckoClient,
				databaseClient:  testcase.fields.databaseClient,
				rateLimiter:     testcase.fields.rateLimiter,
			}

			testcase.want(t, job.buildTokenListFromCoinGecko(testcase.arguments.ctx, testcase.arguments.coinList), fmt.Sprintf("buildTokenListFromCoinGecko(%v, %v)", testcase.arguments.ctx, testcase.arguments.coinList))
		})
	}
}

func TestJob_updateTokenListFromZySync(t *testing.T) {
	initialize(t)

	type fields struct {
		coingeckoClient *coingecko.Client
		zksyncClient    *zksync.Client
		databaseClient  *gorm.DB
		rateLimiter     ratelimit.Limiter
	}

	type arguments struct {
		ctx       context.Context
		coinList  []coingecko.Coin
		tokenList []zksync.GetTokenListItem
	}

	testcases := []struct {
		name      string
		fields    fields
		arguments arguments
		want      assert.ValueAssertionFunc
		wantError assert.ErrorAssertionFunc
	}{
		{
			name: "default",
			fields: fields{
				coingeckoClient: coingecko.New(),
				zksyncClient:    zksync.New(),
				databaseClient:  database.Global(),
				rateLimiter:     ratelimit.New(50, ratelimit.Per(time.Minute)),
			},
			arguments: arguments{
				ctx: context.Background(),
				coinList: []coingecko.Coin{
					{
						ID:     "dai",
						Name:   "DAI",
						Symbol: "DAI",
						Platforms: map[string]string{
							"ethereum": "0x6b175474e89094c44da98b954eedeac495271d0f",
						},
						Image: &coingecko.CoinImage{
							Large: "https://assets.coingecko.com/coins/images/9956/large/4943.png",
						},
					},
				},
				tokenList: []zksync.GetTokenListItem{
					{
						ID:       1,
						Address:  "0x6b175474e89094c44da98b954eedeac495271d0f",
						Symbol:   "DAI",
						Decimals: 18,
						Kind:     zksync.KindERC20,
						IsNFT:    false,
					},
				},
			},
			want:      assert.NotEmpty,
			wantError: assert.NoError,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			job := &Job{
				coingeckoClient: testcase.fields.coingeckoClient,
				zksyncClient:    testcase.fields.zksyncClient,
				databaseClient:  testcase.fields.databaseClient,
				rateLimiter:     testcase.fields.rateLimiter,
			}

			tokens, err := job.buildTokenListFromZkSync(testcase.arguments.ctx, testcase.arguments.tokenList, testcase.arguments.coinList)
			if testcase.wantError(t, err, fmt.Sprintf("updateTokenListFromZkSync(%v, %v, %v)", testcase.arguments.ctx, testcase.arguments.tokenList, testcase.arguments.coinList)) {
				return
			}

			testcase.want(t, tokens, fmt.Sprintf("updateTokenListFromZkSync(%v, %v, %v)", testcase.arguments.ctx, testcase.arguments.tokenList, testcase.arguments.coinList))
		})
	}
}
