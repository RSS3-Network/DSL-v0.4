package spaceid

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/spaceid"
	spaceidcontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/spaceid/contracts"
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
	spaceidContract, err := spaceidcontract.NewSpaceid(spaceid.AddressSpaceID, c.ethClient)
	if err != nil {
		loggerx.Global().Error("failed to connect to bsc rpc:", zap.Error(err))

		return nil, err
	}

	namehash, _ := goens.NameHash(fmt.Sprintf("%s.addr.reverse", strings.TrimPrefix(address, "0x")))

	resolver, err := spaceidContract.Resolver(&bind.CallOpts{}, namehash)
	if err != nil {
		loggerx.Global().Error("failed to get space id resolver:", zap.Error(err))

		return nil, err
	}

	if resolver == ethereum.AddressGenesis {
		loggerx.Global().Error("the space id does not have a resolver: ", zap.Error(err))

		return nil, fmt.Errorf("the space id does not have a resolver")
	}

	spaceidResolver, err := spaceidcontract.NewResolver(resolver, c.ethClient)
	if err != nil {
		loggerx.Global().Error("failed to new to space id resolver contract:", zap.Error(err))

		return nil, err
	}

	handle, err := spaceidResolver.Name(&bind.CallOpts{}, namehash)
	if err != nil {
		loggerx.Global().Error("failed to get space id handle by address:", zap.Error(err))

		return nil, err
	}

	profile := &social.Profile{
		Address:     address,
		Network:     protocol.NetworkBinanceSmartChain,
		Platform:    protocol.PlatformSpaceID,
		Source:      protocol.PlatformSpaceID,
		Name:        handle,
		Handle:      handle,
		ProfileUris: []string{fmt.Sprintf("https://meta.image.space.id/image/mainnet/%v.svg", resolver)},
	}

	return profile, nil
}

func New() *Client {
	var ethClient *ethclient.Client

	ethClient, err := ethclientx.Global(protocol.NetworkBinanceSmartChain)
	if err != nil {
		loggerx.Global().Error("get ethclient error", zap.Error(err))

		return nil
	}

	return &Client{
		ethClient: ethClient,
	}
}
