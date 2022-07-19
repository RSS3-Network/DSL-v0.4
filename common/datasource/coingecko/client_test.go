package coingecko_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/datasource/coingecko"
)

var coingeckoClient = coingecko.New()

func TestCoinList(t *testing.T) {
	coinList, err := coingeckoClient.CoinList(context.Background(), coingecko.CoinListParameter{})
	if err != nil {
		t.Fatal(err)
	}

	for _, coin := range coinList {
		t.Log(coin.ID, coin.Name, coin.Symbol)
	}
}

func TestCoin(t *testing.T) {
	coin, err := coingeckoClient.Coin(context.Background(), "rss3", coingecko.CoinParameter{
		Tickers:       false,
		MarketData:    false,
		CommunityData: false,
		DeveloperData: false,
		Sparkline:     false,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(coin)
}
