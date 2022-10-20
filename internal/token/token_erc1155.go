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
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"

	"go.uber.org/zap"
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
		loggerx.Global().Named(contractAddress).Warn("Get ERC1155 Metadata error", zap.Error(err))
	}

	var metadata Metadata

	if err := json.Unmarshal(result.Metadata, &metadata); err != nil {
		loggerx.Global().Named(contractAddress).Warn("Get ERC1155 Metadata Unmarshal error", zap.Error(err))
	}

	return &result, nil
}

func (c *Client) erc1155ToNFT(erc1155 *ERC1155, tokenID *big.Int) (*NFT, error) {
	var metadata Metadata

	if err := json.Unmarshal(erc1155.Metadata, &metadata); err != nil {
		loggerx.Global().Named(erc1155.ContractAddress).Warn("Get ERC1155 Metadata Unmarshal error", zap.Error(err))
		metadata = Metadata{}
	}

	if metadata.Name == "" {
		metadata.Name = erc1155.Name
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
