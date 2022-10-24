package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	Scheme   = "https"
	Endpoint = "guardian.farcaster.xyz"
	//	https://guardian.farcaster.xyz/origin/address_activity/0x012D3606bAe7aebF03a04F8802c561330eAce70A
)

type Cast struct {
	Body struct {
		Type        string `json:"type"`
		PublishedAt int64  `json:"publishedAt"`
		Sequence    int    `json:"sequence"`
		Username    string `json:"username"`
		Address     string `json:"address"`
		Data        struct {
			Text                  string `json:"text"`
			ReplyParentMerkleRoot string `json:"replyParentMerkleRoot"`
		} `json:"data"`
		// PrevMerkleRoot string `json:"prevMerkleRoot"`
	} `json:"body"`
	MerkleRoot string `json:"merkleRoot"`
	// Signature  string `json:"signature"`
	Meta struct {
		DisplayName string `json:"displayName"`
		Avatar      string `json:"avatar"`
		// IsVerifiedAvatar bool   `json:"isVerifiedAvatar"`
		// NumReplyChildren int    `json:"numReplyChildren"`
		// Reactions        struct {
		//	Count int    `json:"count"`
		//	Type  string `json:"type"`
		//	Self  bool   `json:"self"`
		// } `json:"reactions"`
		// Recasts struct {
		//	Count int  `json:"count"`
		//	Self  bool `json:"self"`
		// } `json:"recasts"`
		// Watches struct {
		//	Count int  `json:"count"`
		//	Self  bool `json:"self"`
		// } `json:"watches"`
		ReplyParentUsername struct {
			Address  string `json:"address"`
			Username string `json:"username"`
		} `json:"replyParentUsername"`
	} `json:"meta"`
	// Attachments struct {
	//	OpenGraph []interface{} `json:"openGraph"`
	// } `json:"attachments"`
}

type Directory struct {
	Body struct {
		AddressActivityUrl string `json:"addressActivityUrl"`
		AvatarUrl          string `json:"avatarUrl"`
		DisplayName        string `json:"displayName"`
		ProofUrl           string `json:"proofUrl"`
		Timestamp          int64  `json:"timestamp"`
		Version            int    `json:"version"`
	} `json:"body"`
	MerkleRoot string `json:"merkleRoot"`
	Signature  string `json:"signature"`
}

type Client struct {
	httpClient *http.Client
}

func (c *Client) GetActivityList(ctx context.Context, address string) ([]Cast, error) {
	activityUrl, err := c.GetFarcasterAddress(ctx, address)
	if err != nil {
		return nil, err
	}

	httpResponse, err := c.httpClient.Get(activityUrl)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	var activityList []Cast

	if err := json.NewDecoder(httpResponse.Body).Decode(&activityList); err != nil {
		return nil, err
	}

	return activityList, nil
}

func (c *Client) GetFarcasterAddress(ctx context.Context, address string) (string, error) {
	requestURL := &url.URL{
		Scheme: Scheme,
		Host:   Endpoint,
		Path:   fmt.Sprintf("/origin/directory/%s", address),
	}

	httpResponse, err := c.httpClient.Get(requestURL.String())
	if err != nil {
		return "", err
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()
	// non-200 means the address is not registered
	// or the server is down
	if httpResponse.StatusCode != http.StatusOK {
		return "", nil
	}

	var directory Directory

	if err := json.NewDecoder(httpResponse.Body).Decode(&directory); err != nil {
		return "", err
	}

	return directory.Body.AddressActivityUrl, nil
}

func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}
