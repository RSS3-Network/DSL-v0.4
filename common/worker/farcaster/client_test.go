package farcaster_test

import (
	"context"
	"github.com/naturalselectionlabs/pregod/common/worker/farcaster"
	"testing"
)

var client = farcaster.NewClient()

func TestGetUserList(t *testing.T) {
	result, err := client.GetUserList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(result))
}
