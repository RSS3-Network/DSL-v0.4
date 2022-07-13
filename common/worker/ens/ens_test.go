package ens_test

import (
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *ens.Client

var address = "0x827431510a5D249cE4fdB7F00C83a3353F471848"
var ensDomain = "henryqw.eth"

func init() {
	// TODO: load config
	var rpc = ""
	client, _ = ens.New(rpc)
}

func TestGetReverseResolve(t *testing.T) {
	result, err := client.GetReverseResolve(address)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, ensDomain)
}

func TestGetResolvedAddress(t *testing.T) {
	result, err := client.GetResolvedAddress(ensDomain)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, address)
}
