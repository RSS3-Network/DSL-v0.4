package exchange_test

import (
	"context"
	"github.com/naturalselectionlabs/pregod/common/exchange"
	lop "github.com/samber/lo/parallel"
	"testing"
)

var (
	client *exchange.Client
)

func init() {
	client = exchange.NewClient()
}

// this test might take a while to run (>1 minutes)
func TestGetSwapPools(t *testing.T) {
	lop.ForEach(exchange.SwapProviders, func(provider exchange.SwapProvider, k int) {
		lop.ForEach(provider.SwapPools, func(pool exchange.SwapPool, i int) {
			result, err := client.GetSwapPairs(context.Background(), provider.Name, pool)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(len(result))
		})
	})
}
