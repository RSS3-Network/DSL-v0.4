package lens

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	lenscontract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	lop "github.com/samber/lo/parallel"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/shurcooL/graphql"
	"go.uber.org/zap"
)

type Client struct {
	graphqlClient *graphql.Client
	httpClient    *http.Client
}

func New() *Client {
	endpointURL := url.URL{
		Scheme: "https",
		Host:   "api.lens.dev",
		Path:   "/graphql",
	}

	client := &Client{
		graphqlClient: graphql.NewClient(endpointURL.String(), http.DefaultClient),
		httpClient:    http.DefaultClient,
	}
	return client
}

// profile request
type ProfileQueryRequest struct {
	OwnedBy []string `json:"ownedBy"`
}

type ProfileQueryResponse struct {
	Profiles struct {
		Items []struct {
			ID graphql.String `json:"id"`
		} `graphql:"items"`
	} `graphql:"profiles(request: $request )"`
}

func (c *Client) BatchGetProfileID(ctx context.Context, address string) ([]*big.Int, error) {
	req := map[string]interface{}{
		"request": ProfileQueryRequest{
			OwnedBy: []string{address},
		},
	}
	resp := &ProfileQueryResponse{}

	if err := c.graphqlClient.Query(ctx, &resp, req); err != nil {
		loggerx.Global().Error("lens graphql query error", zap.String("address", address), zap.Error(err))

		return nil, err
	}

	result := []*big.Int{}
	for _, item := range resp.Profiles.Items {
		s := []byte(item.ID)
		if s[2] == '0' {
			s = append(s[:2], s[3:]...)
		}

		id, err := hexutil.DecodeBig(string(s))
		if err != nil {
			loggerx.Global().Error("lens graphq client, decode response error", zap.String("id", string(item.ID)), zap.Error(err))

			continue
		}

		if id.Int64() > 0 {
			result = append(result, id)
		}
	}

	return result, nil
}

func (c *Client) BatchGetProfiles(ctx context.Context, address string) ([]*social.Profile, error) {
	ethereumClient, err := ethclientx.Global(protocol.NetworkPolygon)
	if err != nil {
		loggerx.Global().Error("get ethclientx error", zap.Error(err))

		return nil, err
	}

	lensHubContract, err := lenscontract.NewHub(lens.HubProxyContractAddress, ethereumClient)
	if err != nil {
		loggerx.Global().Error("lens new hub error", zap.Error(err))

		return nil, err
	}

	profileIDs, err := c.BatchGetProfileID(ctx, address)
	if err != nil {
		return nil, err
	}

	profiles := lop.Map(profileIDs, func(profileID *big.Int, i int) *social.Profile {
		result, err := lensHubContract.GetProfile(&bind.CallOpts{}, profileID)
		if err != nil {
			loggerx.Global().Error("lens hubContract GetProfile error", zap.String("id", profileID.String()), zap.Error(err))

			return nil
		}

		profile := &social.Profile{
			Address:     address,
			Network:     protocol.NetworkPolygon,
			Platform:    protocol.PlatformLens,
			Source:      protocol.PlatformLens,
			Name:        result.Handle,
			Handle:      result.Handle,
			URL:         fmt.Sprintf("https://lenster.xyz/u/%v", result.Handle),
			ProfileUris: []string{ipfs.GetDirectURL(result.ImageURI)},
		}

		return profile
	})

	return profiles, nil
}
