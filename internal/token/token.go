package token

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
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

type ERC1155 struct {
	Name            string          `json:"name"`
	Symbol          string          `json:"symbol"`
	ContractAddress string          `json:"contract_address"`
	ID              *big.Int        `json:"id"`
	Metadata        json.RawMessage `json:"metadata"`
	URI             string          `json:"uri"`
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

type Metadata struct {
	Name            string                      `json:"name"`
	ImageData       string                      `json:"image_data"`
	Title           string                      `json:"title"`
	Description     string                      `json:"description"`
	AnimationURL    string                      `json:"animation_url"`
	Image           string                      `json:"image"`
	ImageURL        string                      `json:"image_url"` // POAP
	Type            string                      `json:"type"`
	Attributes      []MetadataAttribute         `json:"attributes"`
	Properties      map[string]MetadataProperty `json:"properties"`
	Traits          []MetadataTrait             `json:"traits"`
	BackgroundColor string                      `json:"background_color"`
	ExternalLink    string                      `json:"external_link"`
	ExternalURL     string                      `json:"external_url"`
	YoutubeURL      string                      `json:"youtube_url"`
}

type MetadataAttribute struct {
	DisplayType string `json:"display_type"`
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

func (c *Client) NFT(ctx context.Context, network, contractAddress string, tokenID *big.Int) (*NFT, error) {
	erc721, err := c.ERC721(ctx, network, contractAddress, tokenID)
	if err == nil {
		return c.erc721ToNFT(erc721, tokenID)
	}

	erc1155, err := c.ERC1155(ctx, network, contractAddress, tokenID)

	if err == nil {
		return c.erc1155ToNFT(erc1155, tokenID)
	}

	return nil, ErrorTokenNotFound
}

func (c *Client) metadataToAttributes(metadata Metadata) []MetadataAttribute {
	attributeMap := make(map[string]any)

	for _, attribute := range metadata.Attributes {
		attributeMap[attribute.TraitType] = attribute.Value
		attributeMap[attribute.DisplayType] = attribute.DisplayType
	}

	for key, value := range metadata.Properties {
		attributeMap[key] = value.Description
	}

	for _, trait := range metadata.Traits {
		attributeMap[trait.TraitType] = trait.Value
		attributeMap[trait.DisplayType] = trait.DisplayType
	}

	var attributes []MetadataAttribute

	for traitType, value := range attributeMap {
		attributes = append(attributes, MetadataAttribute{
			TraitType: traitType,
			Value:     value,
		})
	}

	return attributes
}

func (c *Client) URI(contractAddress string, tokenID *big.Int, tokenURI string) (string, error) {
	if strings.HasPrefix(tokenURI, "ipfs://") {
		// TODO Move it to config
		tokenURI = strings.Replace(tokenURI, "ipfs://", "https://ipfs.rss3.page/ipfs/", 1)
	} else if strings.Contains(tokenURI, ";base64,") {
		return tokenURI, nil
	}

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

func (c *Client) Metadata(tokenURI string) ([]byte, error) {
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

	response, err := http.Get(tokenURI)
	if err != nil {
		return nil, err
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
