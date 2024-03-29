package token

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"reflect"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/metadata_url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"

	"go.uber.org/zap"
)

const MaxSize = 1024 * 8

type Native struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
	Logo     string `json:"logo"`
}

type ERC20 struct {
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        uint8  `json:"decimals"`
	ContractAddress string `json:"contract_address"`
	Logo            string `json:"logo"`
}

type ERC721 struct {
	Name            string          `json:"name"`
	Symbol          string          `json:"symbol"`
	ContractAddress string          `json:"contract_address"`
	ID              *big.Int        `json:"id"`
	Metadata        json.RawMessage `json:"metadata"`
	URI             string          `json:"uri"`
}

func (e *ERC721) ToNFT(tokenID *big.Int) (*NFT, error) {
	var metadata Metadata

	if err := json.Unmarshal(e.Metadata, &metadata); err != nil {
		loggerx.Global().Named(e.ContractAddress).Warn("Get NFT Metadata Unmarshal error", zap.Error(err))
		metadata = Metadata{}
	}

	if metadata.Name == "" {
		metadata.Name = e.Name
	}

	nft := NFT{
		Name:            metadata.Name,
		Collection:      e.Name,
		Symbol:          e.Symbol,
		Description:     metadata.Description,
		ContractAddress: e.ContractAddress,
		ID:              tokenID,
		Image:           metadata_url.GetDirectURL(metadata.Image),
		ImageData:       metadata_url.GetDirectURL(metadata.ImageData),
		Attributes:      MetadataToAttributes(metadata),
		Standard:        protocol.TokenStandardERC721,
		BackgroundColor: metadata.BackgroundColor,
		AnimationURL:    metadata.AnimationURL,
		ExternalLink:    metadata.ExternalLink,
		ExternalURL:     metadata.ExternalURL,
		YoutubeURL:      metadata.YoutubeURL,
	}

	if metadata.ExternalLink != "" {
		metadata.ExternalURL = metadata.ExternalLink
	}

	// POAP
	if metadata.ImageURL != "" {
		nft.Image = metadata.ImageURL
	}

	return &nft, nil
}

type ERC1155 struct {
	Name            string          `json:"name"`
	Symbol          string          `json:"symbol"`
	ContractAddress string          `json:"contract_address"`
	ID              *big.Int        `json:"id"`
	Metadata        json.RawMessage `json:"metadata"`
	URI             string          `json:"uri"`
}

func (e ERC1155) ToNFT(tokenID *big.Int) (*NFT, error) {
	var metadata Metadata

	if err := json.Unmarshal(e.Metadata, &metadata); err != nil {
		loggerx.Global().Named(e.ContractAddress).Warn("Get ERC1155 Metadata Unmarshal error", zap.Error(err))
		metadata = Metadata{}
	}

	if metadata.Name == "" {
		metadata.Name = e.Name
	}

	nft := NFT{
		Name:            metadata.Name,
		Collection:      e.Name,
		Symbol:          e.Symbol,
		Description:     metadata.Description,
		ContractAddress: e.ContractAddress,
		ID:              tokenID,
		Image:           metadata_url.GetDirectURL(metadata.Image),
		ImageData:       metadata_url.GetDirectURL(metadata.ImageData),
		Attributes:      MetadataToAttributes(metadata),
		Standard:        protocol.TokenStandardERC1155,
		BackgroundColor: metadata.BackgroundColor,
		AnimationURL:    metadata.AnimationURL,
		ExternalLink:    metadata.ExternalLink,
		ExternalURL:     metadata.ExternalURL,
		YoutubeURL:      metadata.YoutubeURL,
	}

	if metadata.ExternalLink != "" {
		metadata.ExternalURL = metadata.ExternalLink
	}

	return &nft, nil
}

type NFT struct {
	Name            string              `json:"name"`
	Collection      string              `json:"collection,omitempty"` // ERC-1155
	Symbol          string              `json:"symbol,omitempty"`
	Description     string              `json:"description"`
	ContractAddress string              `json:"contract_address"`
	ID              *big.Int            `json:"id"`
	Image           string              `json:"image"`
	ImageData       string              `json:"image_data,omitempty"`
	Attributes      []MetadataAttribute `json:"attributes"`
	Standard        string              `json:"standard"`
	BackgroundColor string              `json:"background_color,omitempty"`
	AnimationURL    string              `json:"animation_url,omitempty"`
	ExternalLink    string              `json:"external_link,omitempty"`
	ExternalURL     string              `json:"external_url,omitempty"`
	YoutubeURL      string              `json:"youtube_url,omitempty"`
}

func (n NFT) ToMetadata() (*metadata.Token, error) {
	id := ""

	if n.ID != nil {
		id = n.ID.String()
	}

	result := &metadata.Token{
		Name:            n.Name,
		Collection:      n.Collection,
		Symbol:          n.Symbol,
		Description:     n.Description,
		ID:              id,
		Image:           n.Image,
		ContractAddress: n.ContractAddress,
		AnimationURL:    n.AnimationURL,
		ExternalLink:    n.ExternalLink,
		Standard:        n.Standard,
	}

	for _, attribute := range n.Attributes {
		result.Attributes = append(result.Attributes, metadata.TokenAttribute{
			TraitType: attribute.TraitType,
			Value:     attribute.Value,
		})
	}

	if strings.HasPrefix(n.Image, "ipfs://") {
		result.Image = metadata_url.GetDirectURL(n.Image)
	}

	return result, nil
}

type Metadata struct {
	Name            string              `json:"name"`
	ImageData       string              `json:"image_data"`
	Title           string              `json:"title"`
	Description     string              `json:"description"`
	AnimationURL    string              `json:"animation_url"`
	Image           string              `json:"image"`
	ImageURL        string              `json:"image_url"` // POAP
	Type            string              `json:"type"`
	Attributes      []MetadataAttribute `json:"attributes"`
	Properties      map[string]any      `json:"properties"`
	Traits          []MetadataTrait     `json:"traits"`
	BackgroundColor string              `json:"background_color"`
	ExternalLink    string              `json:"external_link"`
	ExternalURL     string              `json:"external_url"`
	YoutubeURL      string              `json:"youtube_url"`
}

type MetadataAttribute struct {
	DisplayType string `json:"display_type,omitempty"`
	TraitType   string `json:"trait_type"`
	Value       any    `json:"value"`
}

type MetadataProperty struct {
	Type        string `json:"type"`
	Description any    `json:"description"`
}

type MetadataTrait struct {
	TraitType   string   `json:"trait_type"`
	Value       string   `json:"value"`
	DisplayType string   `json:"display_type"`
	MaxValue    string   `json:"max_value"`
	TraitCount  *big.Int `json:"trait_count"`
	Order       string   `json:"order"`
}

type DomainsMetadata struct {
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	Attributes      []MetadataAttribute `json:"attributes"`
	URL             string              `json:"url"`
	BackgroundImage string              `json:"background_image"`
	Image           string              `json:"image"`
}

func (c *Client) NFT(ctx context.Context, network, contractAddress string, tokenID *big.Int) (*NFT, error) {
	erc721, err := c.ERC721(ctx, network, contractAddress, tokenID)
	if err == nil {
		return erc721.ToNFT(tokenID)
	}

	erc1155, err := c.ERC1155(ctx, network, contractAddress, tokenID)
	if err == nil {
		return erc1155.ToNFT(tokenID)
	}

	return nil, ErrorTokenNotFound
}

func (c *Client) NFTToMetadata(context context.Context, network, contractAddress string, tokenID *big.Int) (*metadata.Token, error) {
	nft, err := c.NFT(context, network, contractAddress, tokenID)
	if err != nil {
		loggerx.Global().Error("get nft error", zap.Error(err))

		return nil, err
	}

	var tokenAttributes []metadata.TokenAttribute

	for _, attribute := range nft.Attributes {
		tokenAttributes = append(tokenAttributes, metadata.TokenAttribute{
			TraitType: attribute.TraitType,
			Value:     attribute.Value,
		})
	}

	id := ""
	if nft.ID != nil {
		id = nft.ID.String()
	}

	tokenMetadata := &metadata.Token{
		Name:            nft.Name,
		Collection:      nft.Collection,
		Symbol:          nft.Symbol,
		Standard:        nft.Standard,
		ContractAddress: nft.ContractAddress,
		Image:           nft.Image,
		ID:              id,
		Description:     nft.Description,
		Attributes:      tokenAttributes,
		ExternalLink:    nft.ExternalLink,
		ExternalURL:     nft.ExternalURL,
		AnimationURL:    nft.AnimationURL,
	}

	return tokenMetadata, nil
}

func MetadataToAttributes(metadata Metadata) []MetadataAttribute {
	attributeMap := make(map[string]any)

	for _, attribute := range metadata.Attributes {
		attributeMap[attribute.TraitType] = attribute.Value
		attributeMap[attribute.DisplayType] = attribute.DisplayType
	}

	var attributes []MetadataAttribute

	for _, trait := range metadata.Traits {
		attributeMap[trait.TraitType] = trait.Value
		attributeMap[trait.DisplayType] = trait.DisplayType
	}

	for key, value := range PropertiesToAttributes(metadata.Properties) {
		attributeMap[key] = value
	}

	for traitType, value := range attributeMap {
		if traitType == "" || value == "" {
			continue
		}

		attributes = append(attributes, MetadataAttribute{
			TraitType: traitType,
			Value:     value,
		})
	}

	return attributes
}

func PropertiesToAttributes(properties map[string]any) (attributeMap map[string]any) {
	attributeMap = map[string]any{}

	types := reflect.TypeOf(&MetadataProperty{})
	field1 := types.Elem().Field(1)
	description := field1.Tag.Get("json")

	for key, value := range properties {
		if value == nil {
			continue
		}

		t := reflect.TypeOf(value)

		if strings.EqualFold(t.Kind().String(), reflect.Map.String()) {
			if temp, ok := value.(map[string]any); ok {
				value = temp[description]
			}
		}

		attributeMap[key] = value
	}

	return attributeMap
}

func (c *Client) URI(contractAddress string, tokenID *big.Int, tokenURI string) (string, error) {
	if strings.Contains(tokenURI, ";base64,") {
		return tokenURI, nil
	}

	tokenURI = metadata_url.GetDirectURL(tokenURI)

	switch common.HexToAddress(contractAddress) {
	case common.HexToAddress("0x2953399124F0cBB46d2CbACD8A89cF0599974963"): // OpenSea Collections (OPENSTORE)
		// https://api.opensea.io/api/v2/metadata/matic/0x2953399124F0cBB46d2CbACD8A89cF0599974963/0x{id}
		fallthrough
	case common.HexToAddress("0x495f947276749Ce646f68AC8c248420045cb7b5e"): // OpenSea Shared Storefront (OPENSTORE)
		tokenURI = strings.Replace(tokenURI, "{id}", fmt.Sprintf("%x", tokenID), 1)
	default:
		tokenURI = strings.Replace(tokenURI, "{id}", tokenID.String(), 1)
	}

	return tokenURI, nil
}

func (c *Client) Metadata(ctx context.Context, tokenURI string) ([]byte, error) {
	if strings.Contains(tokenURI, ";base64,") {
		contents := strings.Split(tokenURI, ";base64,")

		if len(contents) < 2 {
			return nil, ErrorInvalidMetadataFormat
		}

		result, err := base64.StdEncoding.DecodeString(contents[1])
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	tokenURI = metadata_url.GetDirectURL(tokenURI)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, tokenURI, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	data, err := io.ReadAll(io.LimitReader(response.Body, MaxSize))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &json.RawMessage{}); err != nil {
		return nil, ErrorInvalidMetadataFormat
	}

	return data, nil
}
