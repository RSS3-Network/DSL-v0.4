package dexpools_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/dexpools"
)

var (
	client *dexpools.Client
)

func init() {
	client = dexpools.NewClient()
}

// this test might take a while to run (>1 minutes)
func TestGetSwapPools(t *testing.T) {
	for _, dex := range dexpools.SwapPools {
		t.Log("TestGetSwapPools, running: " + dex.Name)
		result, err := client.GetSwapPools(context.Background(), dex)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(len(result))
	}
}
