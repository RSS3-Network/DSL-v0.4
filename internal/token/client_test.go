package token_test

import (
	"context"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"log"
	"math/big"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/internal/token"
)

var tokenClient *token.Client

func init() {
	globalEthereumClientMap, err := ethclientx.Dial(&configx.RPC{
		General: configx.RPCNetwork{
			Ethereum: &configx.RPCEndpoint{
				HTTP: "https://shy-newest-brook.quiknode.pro/fac29f542ec6a0b43d53cff4f22be159172d1503/",
			},
		},
	}, protocol.EthclientNetworks)
	if err != nil {
		log.Fatal(err)
	}

	ethclientx.ReplaceGlobal("ethereum", globalEthereumClientMap["ethereum"])

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

// 0x712b9720b37bd206ed938a5fff4ca48cb89643ba 231
// 0x4A0d5a851BB8e8c3d580C41Ffb2f880e0A4b2AA6 164
func TestClient_ERC721(t *testing.T) {
	erc712, err := tokenClient.ERC721(context.Background(), protocol.NetworkEthereum, "0x4A0d5a851BB8e8c3d580C41Ffb2f880e0A4b2AA6", big.NewInt(164))
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

func TestClient_ERC721_ZORA(t *testing.T) {
	erc712, err := tokenClient.ERC721(context.Background(), protocol.NetworkEthereum, "0xabEFBc9fD2F806065b4f3C237d4b59D9A97Bcac7", big.NewInt(13737))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(erc712)
}
