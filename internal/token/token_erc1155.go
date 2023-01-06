package token

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
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

	if tokenID != nil {
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
	}

	return &result, nil
}
