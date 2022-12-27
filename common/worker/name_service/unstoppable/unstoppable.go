package unstoppable

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/unstoppable_domains"
	udcontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/unstoppable_domains/contracts"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/httpx"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"
)

type Client struct {
	ethClient *ethclient.Client
}

func (c *Client) GetProfile(address string) (*social.Profile, error) {
	reverseRegistry, err := udcontract.NewUnstoppableDomains(unstoppable_domains.AddressUDPolygon, c.ethClient)
	if err != nil {
		loggerx.Global().Error("failed to initiate an unstoppable domains reverse registry: ", zap.Error(err))

		return nil, err
	}

	tokenURI, err := reverseRegistry.ReverseOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		loggerx.Global().Error("not a unstoppable domains user:", zap.Error(err))

		return nil, err
	}

	UDResult := &unstoppable_domains.UnstoppableDomain{}
	UDEndpoint := "metadata.unstoppabledomains.com"
	request := http.Request{Method: http.MethodGet, URL: &url.URL{Scheme: "https", Host: UDEndpoint, Path: fmt.Sprintf("/metadata/%s", tokenURI)}}

	err = httpx.DoRequest(context.Background(), http.DefaultClient, &request, &UDResult)
	if err != nil {
		loggerx.Global().Error("failed to request unstoppabledomains metadata:", zap.Error(err))

		return nil, err
	}

	if len(UDResult.Name) == 0 {
		return nil, fmt.Errorf("unstoppable unregistered")
	}

	profile := &social.Profile{
		Address:     address,
		Network:     protocol.NetworkPolygon,
		Platform:    protocol.PlatformUnstoppableDomain,
		Source:      protocol.PlatformUnstoppableDomain,
		Name:        UDResult.Name,
		Handle:      UDResult.Name,
		URL:         UDResult.ExternalUrl,
		ProfileUris: []string{UDResult.ImageUrl},
	}

	return profile, nil
}

func New() *Client {
	var ethClient *ethclient.Client

	ethClient, err := ethclientx.Global(protocol.NetworkPolygon)
	if err != nil {
		loggerx.Global().Error("get ethclient error", zap.Error(err))

		return nil
	}

	return &Client{
		ethClient: ethClient,
	}
}
