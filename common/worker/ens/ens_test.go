package ens

import (
	"errors"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"log"
	"strings"
	"testing"
)

func init() {
	rpcConfig := configx.RPC{
		General: configx.RPCNetwork{
			Ethereum: &configx.RPCEndpoint{
				// Goerli
				HTTP: "https://rpc.ankr.com/eth_goerli",
			},
		},
	}

	globalEthereumClientMap, err := ethclientx.Dial(&rpcConfig, []string{
		protocol.NetworkEthereum,
	})
	if err != nil {
		log.Fatal(err)
	}

	ethclientx.ReplaceGlobal(protocol.NetworkEthereum, globalEthereumClientMap[protocol.NetworkEthereum])
}

// karl.floersch.eth
func TestResolve(t *testing.T) {
	tests := []struct {
		input  string
		output string
		err    error
	}{
		{"brucexc.eth", "0x23c46e912b34C09c4bCC97F4eD7cDd762cee408A", nil},
		{"1.brucexc.eth", "0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0", nil},
		// wrong resovler
		{"2.brucexc.eth", "0x23c46e912b34C09c4bCC97F4eD7cDd762cee408A", nil},
		// no address
		{"3.brucexc.eth", "0x23c46e912b34C09c4bCC97F4eD7cDd762cee408A", nil},
		{"4.brucexc.eth", "0x23c46e912b34C09c4bCC97F4eD7cDd762cee408A", errors.New("unregistered name")},
	}

	for _, tt := range tests {
		output, err := Resolve(tt.input)
		if tt.err == nil {
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if !strings.EqualFold(tt.output, output) {
				t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, strings.ToLower(output), strings.ToLower(tt.output))
			}
		} else {
			if err == nil {
				t.Fatalf("missing expected error")
			}
			if tt.err.Error() != err.Error() {
				t.Errorf("unexpected error value %v", err)
			}
		}
	}
}
