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
	//var wg sync.WaitGroup
	//for _, provider := range exchange.SwapProviders {
	//	provider := provider
	//	for _, pool := range provider.SwapPools {
	//		wg.Add(1)
	//		pool := pool
	//		t.Log(fmt.Sprintf("TestGetSwapPools, running: %s %s %s", provider.Name, pool.Network, pool.Protocol))
	//		go func(provider exchange.SwapProvider, pool exchange.SwapPool) {
	//			result, err := client.GetSwapPairs(context.Background(), provider.Name, pool)
	//			if err != nil {
	//				t.Fatal(provider.Name, pool.Network, pool.Protocol, err)
	//			}
	//			t.Log(provider.Name, pool.Network, pool.Protocol, len(result))
	//			wg.Done()
	//		}(provider, pool)
	//	}
	//}
	//
	//wg.Wait()

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
