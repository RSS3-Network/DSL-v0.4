package crossbell

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) GetProfile(address string) ([]social.Profile, error) {
	requestURL := &url.URL{
		Scheme: "https",
		Host:   "indexer.crossbell.io",
		Path:   fmt.Sprintf("/v1/addresses/%v/characters", address),
	}

	httpResponse, err := c.httpClient.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	profiles := GetProfilesResponse{}

	if err := json.NewDecoder(httpResponse.Body).Decode(&profiles); err != nil {
		return nil, err
	}

	var result []social.Profile

	lop.ForEach(profiles.List, func(crossbell ProfileResponse, i int) {
		profile := social.Profile{
			Address:  address,
			Network:  protocol.NetworkCrossbell,
			Platform: protocol.PlatformCrossbell,
			Source:   protocol.PlatformCrossbell,
			Handle:   crossbell.Handle,
		}

		content, err := ipfs.GetFileByURL(crossbell.URI)
		if err != nil {
			logrus.Errorf("[common] crossbell: ipfs.GetFileByURL err, %v, uri: %v", err, crossbell.URI)

			return
		}

		var metadata *metadata.Token
		if err := json.Unmarshal(content, &metadata); err != nil {
			logrus.Errorf("[common] crossbell: json.Unmarshal err, %v, uri: %v, json: %v", err, crossbell.URI, content)

			return
		}

		if len(metadata.Image) > 0 {
			profile.Name = metadata.Name
			profile.ProfileUris = []string{ipfs.GetDirectURL(metadata.Image)}
			result = append(result, profile)
			return
		}

		var crossbellProfile *CrossbellProfileStruct
		if err := json.Unmarshal(content, &crossbellProfile); err != nil {
			logrus.Errorf("[common] crossbell: json.Unmarshal err, %v, uri: %v", err, crossbell.URI)

			return
		}

		profile.Name = crossbellProfile.Name
		profile.Bio = crossbellProfile.Bio
		for _, avatar := range crossbellProfile.Avatars {
			profile.ProfileUris = append(profile.ProfileUris, ipfs.GetDirectURL(avatar))
		}

		result = append(result, profile)
	})

	return result, nil
}

func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}
