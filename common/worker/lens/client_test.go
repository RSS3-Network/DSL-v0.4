package lens

import (
	"context"
	"testing"
)

func TestGetProfiles(t *testing.T) {
	client := New()
	result, err := client.BatchGetProfileID(context.Background(),
		"0x7241dddec3a6af367882eaf9651b87e1c7549dff")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
