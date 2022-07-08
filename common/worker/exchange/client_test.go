package exchange_test

import (
	"context"
	exchange2 "github.com/naturalselectionlabs/pregod/common/worker/exchange"
	lop "github.com/samber/lo/parallel"
	"testing"
)

var (
	client *exchange2.Client
)

func init() {
	client = exchange2.NewClient()
}

// this test might take a while to run (>1 minutes)
func TestGetSwapPools(t *testing.T) {
	errs := make(chan error, 100)

	lop.ForEach(exchange2.SwapProviders, func(provider exchange2.SwapProvider, k int) {
		lop.ForEach(provider.SwapPools, func(pool exchange2.SwapPool, i int) {
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
