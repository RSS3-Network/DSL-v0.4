package iqwiki

import (
	"context"
	"strings"
	"testing"
)

var (
	client = NewClient()
)

func TestGetUserActivityList(t *testing.T) {
	address := "0xFF1D1d9E1E3a88A1176870244b5a456FC42e5E90"
	resp, err := client.GetUserActivityList(context.Background(), strings.ToLower(address))

	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(resp))

}
