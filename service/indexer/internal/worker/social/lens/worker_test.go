package lens

import (
	"context"
	"encoding/json"
	"testing"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/stretchr/testify/assert"
)

func TestPostCreated(t *testing.T) {
	rpcConfig := configx.RPC{
		General: configx.RPCNetwork{
			Polygon: &configx.RPCEndpoint{
				HTTP: "https://rpc.rss3.dev/networks/polygon",
			},
		},
	}

	ethereumClientMap, err := ethclientx.Dial(&rpcConfig, []string{
		protocol.NetworkPolygon,
	})
	assert.NoError(t, err, "failed to create ethereum client")

	ethclientx.ReplaceGlobal(protocol.NetworkPolygon, ethereumClientMap[protocol.NetworkPolygon])

	ipfs.New("https://ipfs.rss3.page/ipfs/")

	worker := New()
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
	assert.NoError(t, err, "failed to handle message")

	out, err := json.Marshal(data)
	assert.NoError(t, err, "failed to marshal data")

	t.Log(string(out))
}
