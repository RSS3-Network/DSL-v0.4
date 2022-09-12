package lens

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

func TestPostCreated(t *testing.T) {
	ethereumClientMap, err := ethclientx.Dial(&configx.RPC{
		General: configx.RPCNetwork{
			Polygon: &configx.RPCEndpoint{
				HTTP: "",
			},
		},
	}, []string{"polygon"})
	if err != nil {
		t.Fatal(err)
	}

	worker := New(ethereumClientMap[protocol.NetworkPolygon])
	message := &protocol.Message{
		Address: "0xd1feccf6881970105dfb2b654054174007f0e07e",
	}
	transactions := []model.Transaction{
		// comment created
		{
			Hash:        "0xea39efc5a1315ffe8656ef442b406fdd737eef0f6c8aec442e87018d063027d0",
			AddressFrom: "0xd1feccf6881970105dfb2b654054174007f0e07e",
			AddressTo:   "0xdb46d1dc155634fbc732f92e853b10b288ad5a1d",
		},
	}

	data, err := worker.Handle(context.Background(), message, transactions)
	out, _ := json.Marshal(data)
	fmt.Println(string(out))
	fmt.Println(err)
}
