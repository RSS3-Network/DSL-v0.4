package nftscan

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

var endpointMap = map[string]string{
	protocol.NetworkEthereum:          "https://restapi.nftscan.com",
	protocol.NetworkBinanceSmartChain: "https://bnbapi.nftscan.com",
	protocol.NetworkPolygon:           "https://polygonapi.nftscan.com",
	protocol.NetworkArbitrum:          "https://arbitrumapi.nftscan.com",
	protocol.NetworkOptimism:          "https://optimismapi.nftscan.com",
	protocol.NetworkAvalanche:         "https://avaxapi.nftscan.com",
	protocol.NetworkFantom:            "https://fantomapi.nftscan.com",
	protocol.NetworkXDAI:              "https://gnosisapi.nftscan.com",
}

type Result struct {
	Code int `json:"code"`
	Msg  any `json:"msg"`
	Data any `json:"data"`
}

type Collection struct {
	ContractAddress string  `json:"contract_address"`
	ContractName    string  `json:"contract_name"`
	Symbol          string  `json:"symbol"`
	Description     string  `json:"description"`
	Assets          []Asset `json:"assets"`
}

type Asset struct {
	TokenID       decimal.Decimal `json:"token_id"`
	ERCType       string          `json:"erc_type"`
	Amount        decimal.Decimal `json:"amount"`
	Owner         string          `json:"owner"`
	OwnTimestamp  decimal.Decimal `json:"own_timestamp"`
	MintTimestamp decimal.Decimal `json:"mint_timestamp"`
	TokenURI      string          `json:"token_uri"`
	MetadataJSON  string          `json:"metadata_json"`
	Name          string          `json:"name"`
	ImageURI      string          `json:"image_uri"`
	Attributes    []Attribute     `json:"attributes"`
}

type Attribute struct {
	AttributeName  string `json:"attribute_name"`
	AttributeValue string `json:"attribute_value"`
	Percentage     string `json:"percentage"`
}

type Client interface {
	GetAllNFTsByAccount(ctx context.Context, network string, address string) ([]Collection, error)
}

var _ Client = (*client)(nil)

type client struct {
	httpClient *http.Client
	apiKey     string
}

func (c *client) GetAllNFTsByAccount(ctx context.Context, network string, address string) ([]Collection, error) {
	endpoint, exists := endpointMap[network]
	if !exists {
		return nil, fmt.Errorf("unsupported network %s", network)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/api/v2/account/own/all/%s?show_attribute=true", endpoint, address), nil)
	if err != nil {
		return nil, fmt.Errorf("create http request, %w", err)
	}

	var result []Collection

	if err := c.call(ctx, request, &result); err != nil {
		return nil, fmt.Errorf("call api: %w", err)
	}

	return result, nil
}

func (c *client) call(ctx context.Context, request *http.Request, data any) error {
	request.Header.Set("X-API-KEY", c.apiKey)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("do http request, %w", err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	result := Result{
		Data: data,
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}

	if result.Code != http.StatusOK {
		return fmt.Errorf("unexpected code %d: %s", result.Code, result.Msg)
	}

	return nil
}

func NewClient(ctx context.Context, apiKey string) (Client, error) {
	return &client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
	}, nil
}

func SupportedNetworks() []string {
	return lo.Keys(endpointMap)
}
