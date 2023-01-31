package aptos

import (
	"context"
	"testing"
)

func TestGetAccountTractions(t *testing.T) {
	client := NewClient()

	parameter := GetAccountTransactionsParameter{
		Address: "0x76ae7e8b181237513e913e286708c1d9ccf638f675a19165b9c4c8472e62b0d7",
		Limit:   2,
	}

	result, err := client.GetAccountTractions(context.Background(), parameter)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
