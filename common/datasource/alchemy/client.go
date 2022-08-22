package alchemy

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/go-querystring/query"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/httpx"
)

var ErrorUnsupportedNetwork = errors.New("unsupported network")

type Client struct {
	network    string
	key        string
	httpClient *http.Client
	rpcClient  *rpc.Client
}

func NewClient(network, key string) (*Client, error) {
	client := Client{
		network:    network,
		key:        key,
		httpClient: http.DefaultClient,
	}

	rpcURL, err := client.buildURL(false)
	if err != nil {
		return nil, err
	}

	if client.rpcClient, err = rpc.Dial(rpcURL.String()); err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *Client) buildURL(nft bool) (*url.URL, error) {
	var internalURL url.URL

	internalURL.Scheme = "https"

	switch c.network {
	case protocol.NetworkEthereum:
		internalURL.Host = EndpointEthereum
	case protocol.NetworkPolygon:
		internalURL.Host = EndpointPolygon
	case protocol.NetworkArbitrum:
		internalURL.Host = EndpointArbitrum
	case protocol.NetworkOptimism:
		internalURL.Host = EndpointOptimism
	default:
		return nil, ErrorUnsupportedNetwork
	}

	if nft {
		internalURL.Path = fmt.Sprintf("/nft/v2/%s", c.key)
	} else {
		internalURL.Path = fmt.Sprintf("/v2/%s", c.key)
	}

	return &internalURL, nil
}

func (c *Client) GetAssetTransfers(ctx context.Context, parameter GetAssetTransfersParameter) (*GetAssetTransfersResult, error) {
	result := GetAssetTransfersResult{}

	if err := c.rpcClient.CallContext(ctx, &result, MethodGetAssetTransfers, parameter); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetNFTs(ctx context.Context, parameter GetNFTsParameter) (*GetNFTsResult, error) {
	values, err := query.Values(parameter)
	if err != nil {
		return nil, err
	}

	url, err := c.buildURL(true)
	if err != nil {
		return nil, err
	}
	url.Path += fmt.Sprintf("/%v", MethodGetNFTs)

	url.RawQuery = values.Encode()

	request, err := httpx.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	result := GetNFTsResult{}

	err = httpx.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
