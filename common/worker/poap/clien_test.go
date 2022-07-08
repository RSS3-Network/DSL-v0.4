package poap_test

import (
	"github.com/naturalselectionlabs/pregod/common/worker/poap"
	"testing"
)

var client = poap.New()

func TestName(t *testing.T) {
	actions, err := client.GetActions("0x000000A52a03835517E9d193B3c27626e1Bc96b1")
	if err != nil {
		t.Fatal(err)
	}

	for _, action := range actions {
		t.Log(action)
	}
}
