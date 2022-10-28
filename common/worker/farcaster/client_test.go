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

func TestGetActivityList(t *testing.T) {
	address := "0x70844dc4C5084a3fabbfede2b49C4038DDfE66A3"
	result, err := client.GetActivityList(context.Background(), address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(result))
}

func TestConvertFarcasterAdderess(t *testing.T) {
	address := "0x4114E33eb831858649ea3702E1C9a2db3f626446"
	result, err := client.ConvertFarcasterAdderess(context.Background(), address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetFarcasterCacheMap(t *testing.T) {
	result := client.GetFarcasterCacheMap()
	t.Log(len(result))
}
