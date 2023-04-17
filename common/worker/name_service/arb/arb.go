package arb

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/spaceid"
	arbcontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/spaceid/arb"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	goens "github.com/wealdtech/go-ens/v3"
	"go.uber.org/zap"
)

type Client struct {
	ethClient *ethclient.Client
}

func (c *Client) GetProfile(address string) (*social.Profile, error) {
	arbContract, err := arbcontract.NewArbid(spaceid.AddressArb, c.ethClient)
	if err != nil {
		loggerx.Global().Error("failed to connect to arb rpc:", zap.Error(err))

		return nil, err
	}

	namehash, _ := goens.NameHash(fmt.Sprintf("%s.addr.reverse", strings.TrimPrefix(address, "0x")))

	resolver, err := arbContract.Resolver(&bind.CallOpts{}, namehash)
	if err != nil {
		loggerx.Global().Error("failed to get arb resolver:", zap.Error(err))

		return nil, err
	}

	if resolver == ethereum.AddressGenesis {
		loggerx.Global().Error("the arb does not have a resolver: ", zap.Error(err))

		return nil, fmt.Errorf("the arbdoes not have a resolver")
	}

	arbResolver, err := arbcontract.NewResolver(resolver, c.ethClient)
	if err != nil {
		loggerx.Global().Error("failed to new to arb resolver contract:", zap.Error(err))

		return nil, err
	}

	handle, err := arbResolver.Name(&bind.CallOpts{}, namehash)
	if err != nil {
		loggerx.Global().Error("failed to get arb handle by address:", zap.Error(err))

		return nil, err
	}

	label := NameLabel(handle)
	tokenID := TokenID(label)

	profile := &social.Profile{
		Address:     address,
		Network:     protocol.NetworkArbitrum,
		Platform:    protocol.PlatformArb,
		Source:      protocol.PlatformArb,
		Name:        handle,
		Handle:      handle,
		ProfileUris: []string{fmt.Sprintf("https://meta.image.arb.id/image/mainnet/%v.svg", tokenID)},
	}

	return profile, nil
}

func NameLabel(handle string) (node common.Hash) {
	name := strings.Split(handle, ".arb")[0]
	return crypto.Keccak256Hash([]byte(name))
}

func TokenID(label common.Hash) *big.Int {
	return new(big.Int).SetBytes(label.Bytes())
}

func New() *Client {
	var ethClient *ethclient.Client

	ethClient, err := ethclientx.Global(protocol.NetworkArbitrum)
	if err != nil {
		loggerx.Global().Error("get ethclient error", zap.Error(err))

		return nil
	}

	return &Client{
		ethClient: ethClient,
	}
}
