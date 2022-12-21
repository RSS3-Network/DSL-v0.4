package avvy

import (
	"fmt"
	"strconv"
	"time"

	"github.com/avvydomains/golang-client/avvy"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"

	"go.uber.org/zap"
)

type Client struct {
	avvyClient *avvy.Client
}

func (c *Client) GetProfile(address string) (*social.Profile, error) {
	var handle string

	hash, success1 := c.avvyClient.ReverseResolve(c.avvyClient.RECORDS["EVM"], address)
	if success1 {
		name, success2 := c.avvyClient.LookupHash(hash)
		if success2 {
			handle = name
		} else {
			loggerx.Global().Error("the avvy does not find name hash")

			return nil, fmt.Errorf("the avvy does not find name hash")

		}
	} else {
		loggerx.Global().Error("the avvy does not have a resolver")

		return nil, fmt.Errorf("the avvy does not have a resolve")

	}

	expiry := c.avvyClient.GetExpiry(hash)

	expireAt := time.Unix(expiry.Int64(), 0)

	profile := &social.Profile{
		Address:  address,
		Network:  protocol.NetworkAvalanche,
		Platform: protocol.PlatformAvvy,
		Source:   protocol.PlatformAvvy,
		Name:     handle,
		Handle:   handle,
		ExpireAt: &expireAt,
	}

	return profile, nil
}

func New() *Client {
	chainId, _ := strconv.ParseInt(protocol.NetworkToID(protocol.NetworkAvalanche)[2:], 16, 64)

	avvyClient := new(avvy.Client)

	networkUrl, err := ethclientx.GlobalUrl(protocol.NetworkAvalanche)
	if err != nil {
		loggerx.Global().Error("get avvy client error", zap.Error(err))

		return nil
	}

	avvyClient.Init(networkUrl, int(chainId))
	return &Client{
		avvyClient: avvyClient,
	}
}
