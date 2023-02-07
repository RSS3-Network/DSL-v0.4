package middlewarex

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service"
)

func init() {
	rpcConfig := configx.RPC{
		General: configx.RPCNetwork{
			Ethereum: &configx.RPCEndpoint{
				WebSocket: "https://rpc.ankr.com/eth",
			},
			Polygon: &configx.RPCEndpoint{
				HTTP: "https://rpc.ankr.com/polygon",
			},
			Crossbell: &configx.RPCEndpoint{
				HTTP: "https://rpc.crossbell.io",
			},
			BinanceSmartChain: &configx.RPCEndpoint{
				HTTP: "https://rpc.ankr.com/bsc",
			},
			Avalanche: &configx.RPCEndpoint{
				HTTP: "https://api.avax.network/ext/bc/C/rpc",
			},
		},
	}

	globalEthereumClientMap, err := ethclientx.Dial(&rpcConfig, []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
		protocol.NetworkCrossbell,
		protocol.NetworkBinanceSmartChain,
		protocol.NetworkAvalanche,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for network, ethereumClient := range globalEthereumClientMap {
		ethclientx.ReplaceGlobal(network, ethereumClient)
	}
}

func TestResolveAddress(t *testing.T) {

	type arguments struct {
		ns             string
		ignoreContract bool
	}

	tests := []struct {
		name   string
		input  arguments
		output string
		err    error
	}{
		{
			name:   "unregister eth",
			input:  arguments{"qwerfdsazxcv.eth", true},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrUnregisterName),
		},
		{
			name:   "unregister csb",
			input:  arguments{"qwerfdsazxcv.csb", true},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrUnregisterName),
		},
		{
			name:   "unregister lens",
			input:  arguments{"qwerfdsazxcv.lens", true},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrUnregisterName),
		},
		{
			name:   "unregister bnb",
			input:  arguments{"qwerfdsazxcv.bnb", true},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrUnregisterName),
		},
		{
			name:   "unregister bit",
			input:  arguments{"qwerfdsazxcv.bit", true},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrUnregisterName),
		},
		{
			name:   "unregister avax",
			input:  arguments{"qwerfdsazxcv.avax", true},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrUnregisterName),
		},
		{
			name:   "unregister UnstoppableDomains .crypto",
			input:  arguments{"qwerfdsazxcv.crypto", true},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrUnregisterName),
		},
		{
			name:   "unregister UnstoppableDomains .nft",
			input:  arguments{"qwerfdsazxcv.nft", true},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrUnregisterName),
		},
		{
			name:   "unsupport name service .xxx",
			input:  arguments{"qwerfdsazxcv.xxx", true},
			output: "",
			err:    fmt.Errorf(".%s %s, %s", "xxx", name_service.ErrNotSupportNS, name_service.ReferDoc),
		},
		{
			name:   "invalid evm address",
			input:  arguments{"qwerfdsazxc", true},
			output: "qwerfdsazxc",
			err:    nil,
		},
		{
			name: "contract address",
			// USDT
			input:  arguments{"0xdAC17F958D2ee523a2206206994597C13D831ec7", false},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrNotSupportContract),
		},
		{
			name:   "not configure resovler avax",
			input:  arguments{"muafnft.avax", false},
			output: "",
			err:    fmt.Errorf("%s", name_service.ErrNotResolver),
		},

		// normall
		{
			name:   "resolve eth",
			input:  arguments{"vitalik.eth", false},
			output: "0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
			err:    nil,
		},
		{
			name:   "resolve csb",
			input:  arguments{"brucexx.csb", false},
			output: "0x23c46e912b34C09c4bCC97F4eD7cDd762cee408A",
			err:    nil,
		},
		{
			name:   "resolve lens",
			input:  arguments{"diygod.lens", false},
			output: "0xc8b960d09c0078c18dcbe7eb9ab9d816bcca8944",
			err:    nil,
		},
		{
			name:   "resolve bnb",
			input:  arguments{"vitalik.bnb", false},
			output: "0xfFD1a6CAbaa5B35Bf4B5B595C4493949Fa80E3b7",
			err:    nil,
		},
		{
			name:   "resolve bit",
			input:  arguments{"bigben.bit", false},
			output: "0xccbc49ecd4b3b38ef76c285ccad1bca946791a89",
			err:    nil,
		},
		{
			name:   "resolve avax",
			input:  arguments{"avvydomains.avax", false},
			output: "0x9BC4e7C1Fa4Ca66f6B2F4B6F446Dad80Ec541983",
			err:    nil,
		},
		{
			name:   "resolve UnstoppableDomains .crypto",
			input:  arguments{"brad.crypto", false},
			output: "0x8aad44321a86b170879d7a244c1e8d360c99dda8",
			err:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c = echo.New()
			output, err := ResolveAddress(c.NewContext(&http.Request{}, nil), tt.input.ns, tt.input.ignoreContract)
			if tt.err == nil {
				if err != nil {
					t.Fatalf("unexpected error %v", err)
				}
				if !strings.EqualFold(tt.output, output) {
					t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, output, tt.output)
				}
			} else {
				if err == nil {
					t.Fatalf("missing expected error")
				}
				if tt.err.Error() != err.Error() {
					t.Errorf("Failure: unexpected error value %v, (expected %v)\n", err, tt.err)
				}
			}
		})
	}
}
