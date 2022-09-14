package token

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

func (c *Client) ERC1155(context context.Context, network, contractAddress string, tokenID *big.Int) (*ERC1155, error) {
	ethclient, err := ethclientx.Global(network)
	if err != nil {
		return nil, err
	}

	erc1155Contract, err := erc1155.NewERC1155(common.HexToAddress(contractAddress), ethclient)
	if err != nil {
		return nil, err
	}

	result := ERC1155{
		ContractAddress: contractAddress,
		ID:              tokenID,
	}

	result.Name, _ = erc1155Contract.Name(&bind.CallOpts{})
	result.Symbol, _ = erc1155Contract.Symbol(&bind.CallOpts{})

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

	nft := NFT{
		Name:            metadata.Name,
		Collection:      erc1155.Name,
		Symbol:          erc1155.Symbol,
		Description:     metadata.Description,
		ContractAddress: erc1155.ContractAddress,
		ID:              tokenID,
		Image:           metadata.Image,
		ImageData:       metadata.ImageData,
		Attributes:      c.metadataToAttributes(metadata),
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
