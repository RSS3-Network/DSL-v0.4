package lens

import (
	"context"
	"math/big"
	"net/http"
	"net/url"

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

func (c *Client) GetProfiles(ctx context.Context, address string) ([]*big.Int, error) {
	req := map[string]interface{}{
		"request": ProfileQueryRequest{
			OwnedBy: []string{address},
		},
	}
	resp := &ProfileQueryResponse{}

	if err := c.graphqlClient.Query(ctx, &resp, req); err != nil {
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
			loggerx.Global().Warn("lens graphq client, decode response error", zap.String("id", string(item.ID)))

			continue
		}

		if id.Int64() > 0 {
			result = append(result, id)
		}
	}

	return result, nil
}
