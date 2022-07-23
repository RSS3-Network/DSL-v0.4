package ens

import (
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	goens "github.com/wealdtech/go-ens/v3"
)

type Client struct {
	httpClient *http.Client
	ethClient  *ethclient.Client
}

type ENSProfile struct {
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Message     string `json:"message"`
}

func (c *Client) GetProfile(address string) (*model.Profile, error) {
	primaryENS, err := c.GetReverseResolve(address)
	if err != nil {
		return nil, err
	}

	profile := &model.Profile{
		Address:  address,
		Network:  protocol.NetworkEthereum,
		Platform: "ENS",
		Source:   "ENS",
		Name:     primaryENS,
		Handle:   primaryENS,
	}

	err = c.GetENSTextValue(primaryENS, profile)
	if err != nil {
		return nil, err
	}

	expiry, err := c.GetENSExpiry(primaryENS)
	if err != nil {
		return nil, err
	}

	profile.ExpireAt = &expiry

	return profile, nil
}

func (c *Client) GetReverseResolve(address string) (string, error) {
	target, err := goens.ReverseResolve(c.ethClient, common.HexToAddress(address))
	if err != nil {
		return "", err
	}

	return target, nil
}

func (c *Client) GetResolvedAddress(domain string) (string, error) {
	address, err := goens.Resolve(c.ethClient, domain)
	if err != nil {
		return "", err
	}
	return address.String(), nil
}

func (c *Client) GetENSExpiry(domain string) (time.Time, error) {
	name, err := goens.NewName(c.ethClient, domain)
	if err != nil {
		return time.Time{}, err
	}

	expiry, err := name.Expires()
	if err != nil {
		return time.Time{}, err
	}
	return expiry, nil
}

// GetENSTextValue reads the text record value for a given domain
// with the list of Global Keys and Service Keys, see: https://eips.ethereum.org/EIPS/eip-634
func (c *Client) GetENSTextValue(domain string, profile *model.Profile) error {
	r, err := goens.NewResolver(c.ethClient, domain)
	if err != nil {
		return err
	}

	for _, key := range getTextRecordKeyList() {
		text, err := r.Text(key)
		if err != nil {
			return err
		}

		if len(text) == 0 {
			continue
		}

		switch key {
		case "url":
			profile.URL = text
		case "avatar":
			profile.ProfileUris = append(profile.ProfileUris, text)
		case "description":
			profile.Bio = text
		case "com.discord", "com.github", "com.reddit", "com.twitter", "org.telegram":
			profile.SocialUris = append(profile.SocialUris, text)
		}
	}

	return nil
}

// returns a list of recommended keys for a given ENS domain, as per https://app.ens.domains/
// this is a combination of Global Keys and Service Keys, see: https://eips.ethereum.org/EIPS/eip-634
func getTextRecordKeyList() []string {
	// nolint:lll // this is a list of keys
	return []string{"email", "url", "avatar", "description", "notice", "keywords", "com.discord", "com.github", "com.reddit", "com.twitter", "org.telegram", "eth.ens.delegate"}
}

func New(ethRPCEndpoint string) *Client {
	var ethClient *ethclient.Client

	ethClient, err := ethclient.Dial(ethRPCEndpoint)
	if err != nil {
		return nil
	}

	return &Client{
		httpClient: http.DefaultClient,
		ethClient:  ethClient,
	}
}

func Resolve(ethRPCEndpoint string, input string) (string, error) {
	client := New(ethRPCEndpoint)

	var result string
	var err error

	if strings.HasSuffix(input, ".eth") {
		result, err = client.GetResolvedAddress(input)
	} else {
		result, err = client.GetReverseResolve(input)
	}

	if err != nil {
		return "", err
	}

	return result, nil
}
