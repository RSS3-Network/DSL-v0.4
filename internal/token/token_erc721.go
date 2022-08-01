package token

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

func (c *Client) ERC721(ctx context.Context, network, contractAddress string, tokenID *big.Int) (*ERC721, error) {
	ethereumClient, exists := c.ethereumClientMap[network]
	if !exists {
		return nil, ErrorNetworkNotSupported
	}

	erc721Contract, err := erc721.NewERC721(common.HexToAddress(contractAddress), ethereumClient)
	if err != nil {
		return nil, err
	}

	result := ERC721{
		ContractAddress: contractAddress,
		ID:              tokenID,
	}

	if result.Name, err = erc721Contract.Name(&bind.CallOpts{}); err != nil {
		return nil, err
	}

	if result.Symbol, err = erc721Contract.Symbol(&bind.CallOpts{}); err != nil {
		return nil, err
	}

	tokenURI, err := erc721Contract.TokenURI(&bind.CallOpts{}, tokenID)
	if err != nil {
		return nil, err
	}

	if result.URI, err = c.URI(contractAddress, tokenID, tokenURI); err != nil {
		return nil, err
	}

	result.Metadata, err = c.Metadata(result.URI)
	if err != nil {
		return nil, err
	}

	var metadata Metadata

	if err := json.Unmarshal(result.Metadata, &metadata); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) erc721ToNFT(erc721 *ERC721, tokenID *big.Int) (*NFT, error) {
	var metadata Metadata

	if err := json.Unmarshal(erc721.Metadata, &metadata); err != nil {
		return nil, err
	}

	nft := NFT{
		Name:            metadata.Name,
		Collection:      erc721.Name,
		Symbol:          erc721.Symbol,
		Description:     metadata.Description,
		ContractAddress: erc721.ContractAddress,
		ID:              tokenID,
		Image:           metadata.Image,
		ImageData:       metadata.ImageData,
		Attributes:      c.metadataToAttributes(metadata),
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
