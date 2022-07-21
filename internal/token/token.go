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
	"github.com/naturalselectionlabs/pregod/common/protocol"
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
	Image           string          `json:"image"`
	ID              *big.Int        `json:"id"`
	Metadata        json.RawMessage `json:"metadata"`
	URI             string          `json:"uri"`
}

type ERC1155 struct {
	Name            string          `json:"name"`
	ContractAddress string          `json:"contract_address"`
	Image           string          `json:"image"`
	ID              *big.Int        `json:"id"`
	Metadata        json.RawMessage `json:"metadata"`
	URI             string          `json:"uri"`
}

type NFT struct {
	Name        string          `json:"name"`
	Symbol      string          `json:"symbol,omitempty"`
	Description string          `json:"description"`
	ID          *big.Int        `json:"id"`
	Image       string          `json:"image"`
	Standard    string          `json:"standard"`
	Metadata    json.RawMessage `json:"metadata"`
}

type Metadata struct {
	Name         string                     `json:"name"`
	Title        string                     `json:"title"`
	Description  string                     `json:"description"`
	AnimationURL string                     `json:"animation_url"`
	Image        string                     `json:"image"`
	Type         string                     `json:"type"`
	Attributes   []MetadataAttribute        `json:"attributes"`
	Properties   map[string]json.RawMessage `json:"properties"`
	Traits       []MetadataTrait            `json:"traits"`
	ExternalLink string                     `json:"external_link"`
}

type MetadataAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
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
		var metadata Metadata

		if err := json.Unmarshal(erc721.Metadata, &metadata); err != nil {
			return nil, err
		}

		return &NFT{
			Name:     metadata.Name,
			Symbol:   erc721.Symbol,
			ID:       tokenID,
			Image:    erc721.Image,
			Standard: protocol.TokenStandardERC721,
			Metadata: erc721.Metadata,
		}, err
	}

	erc1155, err := c.ERC1155(ctx, network, contractAddress, tokenID)

	if err == nil {
		var metadata Metadata

		if erc1155.Metadata == nil {
			return nil, ErrorInvalidMetadataFormat
		}

		if err := json.Unmarshal(erc1155.Metadata, &metadata); err != nil {
			return nil, err
		}

		return &NFT{
			Name:     metadata.Name,
			ID:       tokenID,
			Image:    erc1155.Image,
			Standard: protocol.TokenStandardERC1155,
			Metadata: erc1155.Metadata,
		}, err
	}

	return nil, ErrorTokenNotFound
}

func (c *Client) URI(contractAddress string, tokenID *big.Int, tokenURI string) (string, error) {
	if strings.HasPrefix(tokenURI, "ipfs://") {
		// TODO Move it to config
		tokenURI = strings.Replace(tokenURI, "ipfs://", "https://ipfs.rss3.page/ipfs/", 1)
	} else if strings.Contains(tokenURI, ";base64,") {
		contents := strings.Split(tokenURI, ";base64,")

		if len(contents) < 2 {
			return "", ErrorInvalidMetadataFormat
		}

		result, err := base64.StdEncoding.DecodeString(contents[1])
		if err != nil {
			return "", err
		}

		return string(result), nil
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
