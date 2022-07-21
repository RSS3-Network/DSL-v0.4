package token

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

func (c *Client) ERC1155(context context.Context, network, contractAddress string, tokenID *big.Int) (*ERC1155, error) {
	ethereumClient, exists := c.ethereumClientMap[network]
	if !exists {
		return nil, ErrorNetworkNotSupported
	}

	erc1155Contract, err := erc1155.NewERC1155(common.HexToAddress(contractAddress), ethereumClient)
	if err != nil {
		return nil, err
	}

	result := ERC1155{
		ContractAddress: contractAddress,
		ID:              tokenID,
	}

	tokenURI, err := erc1155Contract.Uri(&bind.CallOpts{}, tokenID)
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

func (c *Client) erc1155ToNFT(erc1155 *ERC1155, tokenID *big.Int) (*NFT, error) {
	var metadata Metadata

	if erc1155.Metadata == nil {
		return nil, ErrorInvalidMetadataFormat
	}

	if err := json.Unmarshal(erc1155.Metadata, &metadata); err != nil {
		return nil, err
	}

	return &NFT{
		Name:            metadata.Name,
		Description:     metadata.Description,
		ID:              tokenID,
		ContractAddress: erc1155.ContractAddress,
		Image:           metadata.Image,
		Attributes:      c.metadataToAttributes(metadata),
		Standard:        protocol.TokenStandardERC1155,
		Metadata:        erc1155.Metadata,
		AnimationURL:    metadata.AnimationURL,
		ExternalLink:    metadata.ExternalLink,
	}, nil
}