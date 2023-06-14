package ens_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/datasource/subgraph/ens"
)

var client = ens.New()

func TestClient_GetIpfs(t *testing.T) {
	nameHash := "0xdf01a30ffe5d99c28bd52bc68a831033807c64191a7b7d4200af422ee4a693ca"

	resp, err := client.GetEnsName(context.Background(), nameHash)
	if err != nil {
		t.Error(err)
	}

	t.Log(resp)
}
