package alchemy

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/go-querystring/query"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	http_utils "github.com/naturalselectionlabs/pregod/common/utils/http"
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

type GetAssetTransfersParameter struct {
	FromBlock         string   `json:"fromBlock,omitempty"`
	ToBlock           string   `json:"toBlock,omitempty"`
	FromAddress       string   `json:"fromAddress,omitempty"`
	ToAddress         string   `json:"toAddress,omitempty"`
	ContractAddresses []string `json:"contractAddresses,omitempty"`
	MaxCount          string   `json:"maxCount,omitempty"`
	ExcludeZeroValue  bool     `json:"excludeZeroValue"`
	PageKey           string   `json:"pageKey,omitempty"`
	WithMetadata      bool     `json:"withMetadata,omitempty"`
	Category          []string `json:"category,omitempty"`
}

type GetAssetTransfersResult struct {
	Transfers []struct {
		BlockNum        string      `json:"blockNum"`
		Hash            string      `json:"hash"`
		From            string      `json:"from"`
		To              string      `json:"to"`
		Value           float64     `json:"value"`
		Erc721TokenId   interface{} `json:"erc721TokenId"`
		Erc1155Metadata interface{} `json:"erc1155Metadata"`
		TokenId         interface{} `json:"tokenId"`
		Asset           string      `json:"asset"`
		Category        string      `json:"category"`
		RawContract     struct {
			Value   string      `json:"value"`
			Address interface{} `json:"address"`
			Decimal string      `json:"decimal"`
		}
	}
	PageKey string `json:"pageKey"`
}

type GetNFTsParameter struct {
	Owner             string   `url:"owner,omitempty"`
	PageKey           string   `url:"pageKey,omitempty"`
	ContractAddresses []string `url:"contractAddresses,omitempty"`
	WithMetadata      bool     `url:"withMetadata"`
}

type GetNFTsResult struct {
	OwnedNFTs  []NFTResult `json:"ownedNfts"`
	PageKey    string      `json:"pageKey"`
	TotalCount int         `json:"totalCount"`
	BlockHash  string      `json:"blockHash"`
}

type NFTResult struct {
	Contract struct {
		Address string `json:"address"`
	} `json:"contract"`
	ID struct {
		TokenID       string `json:"tokenId"`
		TokenMetadata struct {
			TokenType string `json:"tokenType"`
		} `json:"tokenMetadata"`
	} `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TokenURI    struct {
		Raw     string `json:"raw"`
		Gateway string `json:"gateway"`
	} `json:"tokenUri"`
	Media []struct {
		Raw     string `json:"raw"`
		Gateway string `json:"gateway"`
	} `json:"media"`
	Metadata struct {
		Image      string      `json:"image"`
		Attributes interface{} `json:"attributes"`
	} `json:"metadata"`
	TimeLastUpdated time.Time `json:"timeLastUpdated"`
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

	request, err := http_utils.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	result := GetNFTsResult{}

	err = http_utils.DoRequest(ctx, c.httpClient, request, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
