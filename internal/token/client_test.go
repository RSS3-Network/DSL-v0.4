package token_test

import (
	"context"
	"log"
	"math/big"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/ethclientx"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/internal/token"
)

var tokenClient *token.Client

func init() {
	ethereumClientMap, err := ethclientx.Dial(&configx.RPC{
		General: configx.RPCNetwork{
			Ethereum: &configx.RPCEndpoint{
				HTTP: "https://rpc.rss3.dev/networks/ethereum",
			},
			Polygon: &configx.RPCEndpoint{
				HTTP: "https://rpc.rss3.dev/networks/polygon",
			},
			BinanceSmartChain: &configx.RPCEndpoint{
				HTTP: "https://rpc.rss3.dev/networks/binance",
			},
			XDAI: &configx.RPCEndpoint{
				HTTP: "https://rpc.rss3.dev/networks/xdai",
			},
			Crossbell: &configx.RPCEndpoint{
				HTTP: "https://rpc.rss3.dev/networks/crossbell",
			},
		},
	}, protocol.EthclientNetworks)
	if err != nil {
		log.Fatal(err)
	}

	tokenClient = token.New()
}

func TestClient_Native(t *testing.T) {
	native, err := tokenClient.Native(context.Background(), protocol.NetworkEthereum)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(native)
}

func TestClient_ERC20(t *testing.T) {
	erc20, err := tokenClient.ERC20(context.Background(), protocol.NetworkEthereum, "0xc98d64da73a6616c42117b582e832812e7b8d57f")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(erc20)
}

func TestClient_ERC721(t *testing.T) {
	erc712, err := tokenClient.ERC721(context.Background(), protocol.NetworkEthereum, "0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85", big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(erc712)
}

func TestClient_ERC1155(t *testing.T) {
	tokenID := big.NewInt(0)
	tokenID.SetString("4452815761501376758038898720702591022279500679302039557361006834352129", 0)

	erc1155, err := tokenClient.ERC1155(context.Background(), protocol.NetworkPolygon, "0x2953399124f0cbb46d2cbacd8a89cf0599974963", tokenID)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", erc1155)
	t.Log(string(erc1155.Metadata))
}
