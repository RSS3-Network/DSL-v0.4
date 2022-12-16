package lens

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/naturalselectionlabs/pregod/common/database/model/social"
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

// profiles request
type ProfileQueryRequest struct {
	OwnedBy []string `json:"ownedBy"`
}

type ProfileQueryResponse struct {
	Profiles struct {
		Items []Profile `json:"items"`
	} `graphql:"profiles(request: $request )"`
}

// profile request
type SingleProfileQueryRequest struct {
	ID string `json:"profileId"`
}

type SingleProfileQueryResponse struct {
	Profile Profile `graphql:"profile(request: $request )"`
}

type Profile struct {
	ID      graphql.String `json:"id"`
	Name    graphql.String `json:"name"`
	Handle  graphql.String `json:"handle"`
	Bio     graphql.String `json:"bio"`
	OwnedBy graphql.String `json:"ownedBy"`
	Picture ProfilePicture `json:"picture"`
}

type ProfilePicture struct {
	MediaSet struct {
		Original struct {
			URL graphql.String `json:"url"`
		} `json:"original"`
	} `graphql:"... on MediaSet"`
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
	profileIDs, err := c.BatchGetProfileID(ctx, address)
	if err != nil {
		return nil, err
	}

	var profiles []*social.Profile
	var mutex sync.Mutex

	lop.ForEach(profileIDs, func(profileID *big.Int, i int) {
		profile, err := c.GetProfileByProfileID(ctx, profileID)
		if err != nil {
			loggerx.Global().Error("lens graphql query error", zap.String("id", profileID.String()), zap.Error(err))

			// contract
			profile, err = c.GetProfile(nil, "", profileID)
			if err != nil {
				loggerx.Global().Error("lens hubContract GetProfile error", zap.String("id", profileID.String()), zap.Error(err))

				return
			}
		}

		mutex.Lock()
		profiles = append(profiles, profile)
		mutex.Unlock()
	})

	return profiles, nil
}

func (c *Client) GetProfileByProfileID(ctx context.Context, profileID *big.Int) (*social.Profile, error) {
	profileIdHex := []byte(hexutil.EncodeBig(profileID))
	if len(profileIdHex)%2 == 1 {
		profileIdHex = append(profileIdHex[:2], append([]byte("0"), profileIdHex[2:]...)...)
	}

	req := map[string]interface{}{
		"request": SingleProfileQueryRequest{
			ID: string(profileIdHex),
		},
	}

	resp := &SingleProfileQueryResponse{}

	if err := c.graphqlClient.Query(ctx, &resp, req); err != nil {
		loggerx.Global().Error("lens graphql profile query error", zap.String("id", string(profileIdHex)), zap.Error(err))

		return nil, err
	}

	profile := &social.Profile{
		Address:  strings.ToLower(string(resp.Profile.OwnedBy)),
		Platform: protocol.PlatformLens,
		Network:  protocol.NetworkPolygon,
		Name:     string(resp.Profile.Name),
		Source:   protocol.PlatformLens,
		URL:      fmt.Sprintf("https://lenster.xyz/u/%v", string(resp.Profile.Handle)),
		Bio:      string(resp.Profile.Bio),
		Handle:   string(resp.Profile.Handle),
	}

	if url := string(resp.Profile.Picture.MediaSet.Original.URL); len(url) > 0 {
		profile.ProfileUris = append(profile.ProfileUris, ipfs.GetDirectURL(url))
	}

	return profile, nil
}
