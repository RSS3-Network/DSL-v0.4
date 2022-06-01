package cexpools_test

import (
	"context"
	"github.com/naturalselectionlabs/pregod/common/cexpools"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"testing"
)

var (
	client *cexpools.Client
)

func init() {
	client = cexpools.NewClient()
}

// this test might take a while to run (>1 minutes)
func TestGetSwapPools(t *testing.T) {
	for _, dex := range protocol.SwapPools {
		t.Log("TestGetSwapPools, running: " + dex.Name)
		result, err := client.GetSwapPools(context.Background(), dex)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(len(result))
	}
}
