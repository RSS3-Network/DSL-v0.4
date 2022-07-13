package ens_test

import (
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var client *ens.Client

var address = "0x827431510a5D249cE4fdB7F00C83a3353F471848"
var ensDomain = "henryqw.eth"
var expiry = time.Unix(1956691715, 0)

func init() {
	// TODO: load config
	var rpc = ""
	client = ens.New(rpc)
}
func TestGetProfile(t *testing.T) {
	result, err := client.GetProfile(address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
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

func TestGetENSExpiry(t *testing.T) {
	result, err := client.GetENSExpiry(ensDomain)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, result, expiry)
}
