package nftscan

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/httpx"
)

var EndpointMap = map[string]string{
	protocol.NetworkEthereum: "restapi.nftscan.com",
	protocol.NetworkPolygon:  "polygonapi.nftscan.com",
}

type Client struct {
	httpClient *http.Client
}

func (c *Client) HasCollection(ctx context.Context, network, address string) bool {
	if _, ok := EndpointMap[network]; !ok {
		return false
	}

	urlInfo := url.URL{
		Scheme: "https",
		Host:   EndpointMap[network],
		Path:   fmt.Sprintf("/api/v2/collections/%s", address),
	}

	request, err := httpx.NewRequest(http.MethodGet, urlInfo.String(), nil)

	request.Header.Add("X-API-KEY", "hbTSjJ01rungPJjmCA50qGew")

	if err != nil {
		return false
	}

	var result Result

	err = httpx.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		return false
	}

	if result.Code != http.StatusOK || result.Data == nil {
		return false
	}

	return true
}

func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}
