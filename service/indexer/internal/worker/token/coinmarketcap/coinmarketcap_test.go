package coinmarketcap_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/cache"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token/coinmarketcap"
	"github.com/stretchr/testify/assert"
)

func setup() {
	cache.Dial(&configx.Redis{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	coinmarketcap.Init("11f16fe7-7036-42c7-8155-1524d74b05eb")
}

func teardown() {
	coinmarketcap.Init("")
	cache.Clear(context.Background())
	cache.Close()
}

func Test_CachedGetCoinInfo(t *testing.T) {
	setup()
	defer teardown()

	ctx := context.Background()
	rss3Address := "0xc98d64da73a6616c42117b582e832812e7b8d57f"

	// no cache
	cache.Clear(context.Background())
	info, err := coinmarketcap.CachedGetCoinInfo(ctx, "ethereum", rss3Address)
	assert.Nil(t, err)
	assert.EqualValues(t, "RSS3", info.Name)
	assert.EqualValues(t, "rss3", info.Slug)
	assert.EqualValues(t, "RSS3", info.Symbol)
	assert.EqualValues(t, 18, info.Decimals)
	assert.EqualValues(t, "token", info.Category)
	assert.EqualValues(t, 112722403.98, *info.SelfReportedCirculatingSupply)

	// cache exists
	info = &coinmarketcap.CoinInfo{}
	exists, err := cache.GetMsgPack(ctx, "getcoininfo.address."+rss3Address, info)
	assert.Nil(t, err)
	assert.True(t, exists)
	info, err = coinmarketcap.CachedGetCoinInfo(ctx, "ethereum", rss3Address)
	assert.Nil(t, err)
	assert.EqualValues(t, 112722403.98, *info.SelfReportedCirculatingSupply)
	assert.EqualValues(t, "https://s2.coinmarketcap.com/static/img/coins/64x64/17917.png", info.Logo)

	// cache exists (wrong apikey)
	coinmarketcap.Init("aaa")
	info, err = coinmarketcap.CachedGetCoinInfo(ctx, "ethereum", rss3Address)
	assert.Nil(t, err)
	assert.EqualValues(t, "RSS3", info.Name)

	// ETH BNB address
	coinmarketcap.Init("11f16fe7-7036-42c7-8155-1524d74b05eb")
	ethBNBAddress := "0x2170ed0880ac9a755fd29b2688956bd959f933f8"
	info, err = coinmarketcap.CachedGetCoinInfo(ctx, "binance_smart_chain", ethBNBAddress)
	assert.Nil(t, err)
	assert.EqualValues(t, "Ethereum", info.Name)
	assert.EqualValues(t, "ethereum", info.Slug)
	assert.EqualValues(t, 18, info.Decimals)
	assert.EqualValues(t, "ETH", info.Symbol)
	assert.EqualValues(t, "coin", info.Category)
	assert.Nil(t, info.SelfReportedCirculatingSupply)
}
