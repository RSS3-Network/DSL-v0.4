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
	errs := make(chan error, 100)

	lop.ForEach(exchange.SwapProviders, func(provider exchange.SwapProvider, k int) {
		lop.ForEach(provider.SwapPools, func(pool exchange.SwapPool, i int) {
			result, err := client.GetSwapPairs(context.Background(), provider.Name, pool)
			if err != nil {
				t.Log(err)
				errs <- err
			}
			t.Log(len(result))
		})
	})

	if len(errs) != 0 {
		for err := range errs {
			t.Fatal(err)
		}
	}
}
