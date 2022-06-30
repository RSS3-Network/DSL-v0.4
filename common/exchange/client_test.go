package exchange_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/exchange"
)

var (
	client *exchange.Client
)

func init() {
	client = exchange.NewClient()
}

type mockInterface struct {
	wg        sync.WaitGroup
	CallCount int
}

func (m *mockInterface) exec() {
	m.wg.Done()
	m.CallCount += 1
}

func (m *mockInterface) currentCount() int {
	m.wg.Wait()
	return m.CallCount
}

// this test might take a while to run (>1 minutes)
func TestGetSwapPools(t *testing.T) {
	wg := &sync.WaitGroup{}
	for _, provider := range exchange.SwapProviders {
		provider := provider
		for _, pool := range provider.SwapPools {
			pool := pool
			wg.Add(1)
			go func(provider exchange.SwapProvider, pool exchange.SwapPool) {
				t.Log(fmt.Sprintf("TestGetSwapPools, running: %s %s %s", provider.Name, pool.Network, pool.Protocol))
				result, err := client.GetSwapPools(context.Background(), provider.Name, pool)
				if err != nil {
					t.Fatal(provider.Name, pool.Network, pool.Protocol, err)
				}
				t.Log(provider.Name, pool.Network, pool.Protocol, len(result))

				wg.Done()
			}(provider, pool)
		}
	}
	wg.Wait()

	// working:
	//for _, provider := range exchange.SwapProviders {
	//	for _, pool := range provider.SwapPools {
	//			t.Log(fmt.Sprintf("TestGetSwapPools, running: %s %s %s", provider.Name, pool.Network, pool.Protocol))
	//			result, err := client.GetSwapPools(context.Background(), provider.Name, pool)
	//			if err != nil {
	//				t.Fatal(provider.Name, pool.Network, pool.Protocol, err)
	//			}
	//			t.Log(provider.Name, pool.Network, pool.Protocol, len(result))
	//
	//	}
	//}

	//lop.ForEach(exchange.SwapProviders, func(provider exchange.SwapProvider, k int) {
	//	provider = provider
	//	lop.ForEach(provider.SwapPools, func(pool exchange.SwapPool, i int) {
	//
	//		t.Log(fmt.Sprintf("TestGetSwapPools, running: %s %s %s", provider.Name, pool.Network, pool.Protocol))
	//		result, err := client.GetSwapPools(context.Background(), provider.Name, pool)
	//		if err != nil {
	//			t.Fatal(provider.Name, pool.Network, pool.Protocol, err)
	//		}
	//
	//		t.Log(provider.Name, pool.Network, pool.Protocol, len(result))
	//	})
	//})

}
