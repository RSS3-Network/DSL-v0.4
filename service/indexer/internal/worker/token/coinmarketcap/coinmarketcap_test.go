package coinmarketcap_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token/coinmarketcap"
	"github.com/stretchr/testify/assert"
)

func setup() {
	database.Dial("postgres://postgres:password@127.0.0.1:5432/pregod2", true)
	coinmarketcap.Init("11f16fe7-7036-42c7-8155-1524d74b05eb")
}

func teardown() {
	coinmarketcap.Init("")
}

func Test_CachedGetCoinInfo(t *testing.T) {
	setup()
	defer teardown()

	ctx := context.Background()
	rss3Address := "0xc98d64da73a6616c42117b582e832812e7b8d57f"

	// no cache
	info, err := coinmarketcap.CachedGetCoinInfo(ctx, "ethereum", rss3Address)
	assert.Nil(t, err)
	assert.EqualValues(t, "RSS3", info.Name)
	assert.EqualValues(t, "rss3", info.Slug)
	assert.EqualValues(t, "RSS3", info.Symbol)
	assert.EqualValues(t, 18, info.Decimals)
	assert.EqualValues(t, "token", info.Category)

	// cache exists
	info = &model.CoinMarketCapCoinInfo{}
	info, err = coinmarketcap.CachedGetCoinInfo(ctx, "ethereum", rss3Address)
	assert.Nil(t, err)
	assert.EqualValues(t, "https://s2.coinmarketcap.com/static/img/coins/64x64/17917.png", info.Logo)

	// cache exists (wrong apikey)
	coinmarketcap.Init("aaa")
	info, err = coinmarketcap.CachedGetCoinInfo(ctx, "ethereum", rss3Address)
	assert.Nil(t, err)
	assert.EqualValues(t, "RSS3", info.Name)

	// ETH address (on BSC)
	coinmarketcap.Init("11f16fe7-7036-42c7-8155-1524d74b05eb")
	ethBNBAddress := "0x2170ed0880ac9a755fd29b2688956bd959f933f8"
	info, err = coinmarketcap.CachedGetCoinInfo(ctx, "binance_smart_chain", ethBNBAddress)
	assert.Nil(t, err)
	assert.EqualValues(t, "Ethereum", info.Name)
	assert.EqualValues(t, "ethereum", info.Slug)
	assert.EqualValues(t, 18, info.Decimals)
	assert.EqualValues(t, "ETH", info.Symbol)
	assert.EqualValues(t, "coin", info.Category)
}

func Test_CachedGetCoinInfoByNetwork(t *testing.T) {
	setup()
	defer teardown()

	ctx := context.Background()

	// no cache
	info, err := coinmarketcap.CachedGetCoinInfoByNetwork(ctx, "ethereum")
	assert.Nil(t, err)
	assert.EqualValues(t, "Ethereum", info.Name)
	assert.EqualValues(t, "ethereum", info.Slug)
	assert.EqualValues(t, "ETH", info.Symbol)
	assert.EqualValues(t, 18, info.Decimals)
	assert.EqualValues(t, "coin", info.Category)

	// cache exists
	info, err = coinmarketcap.CachedGetCoinInfoByNetwork(ctx, "ethereum")
	assert.EqualValues(t, "Ethereum", info.Name)
	assert.EqualValues(t, "ethereum", info.Slug)
	assert.EqualValues(t, "ETH", info.Symbol)
	assert.EqualValues(t, 18, info.Decimals)
	assert.EqualValues(t, "coin", info.Category)


	// BSC
	info, err = coinmarketcap.CachedGetCoinInfoByNetwork(ctx, "binance_smart_chain")
	assert.Nil(t, err)
	assert.EqualValues(t, "BNB", info.Name)
	assert.EqualValues(t, "bnb", info.Slug)
	assert.EqualValues(t, 18, info.Decimals)
	assert.EqualValues(t, "BNB", info.Symbol)
	assert.EqualValues(t, "coin", info.Category)

	// polygon
	info, err = coinmarketcap.CachedGetCoinInfoByNetwork(ctx, "polygon")
	assert.Nil(t, err)
	assert.EqualValues(t, "Polygon", info.Name)
	assert.EqualValues(t, "polygon", info.Slug)
	assert.EqualValues(t, 18, info.Decimals)
	assert.EqualValues(t, "MATIC", info.Symbol)
	assert.EqualValues(t, "coin", info.Category)
}
