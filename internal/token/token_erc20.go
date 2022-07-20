package token

import (
	"context"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"gorm.io/gorm"
)

func (c *Client) ERC20(ctx context.Context, network string, contractAddress string) (*ERC20, error) {
	var token model.Token

	err := c.databaseClient.
		Model((*model.Token)(nil)).
		Where(map[string]interface{}{
			"network":          network,
			"contract_address": strings.ToLower(contractAddress),
		}).
		First(&token).
		Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &ERC20{
			Name:            token.Name,
			Symbol:          token.Symbol,
			Decimals:        token.Decimal, // TODO Rename it to decimals
			ContractAddress: token.ContractAddress,
			Logo:            token.Logo,
		}, nil
	}

	result := &ERC20{
		ContractAddress: contractAddress,
	}

	ethereumClient, exists := c.ethereumClientMap[network]
	if !exists {
		return nil, ErrorNetworkNotSupported
	}

	erc20Contract, err := erc20.NewERC20(common.HexToAddress(result.ContractAddress), ethereumClient)
	if err != nil {
		return nil, err
	}

	if result.Name, err = erc20Contract.Name(&bind.CallOpts{}); err != nil {
		return nil, err
	}

	if result.Symbol, err = erc20Contract.Symbol(&bind.CallOpts{}); err != nil {
		return nil, err
	}

	if result.Symbol, err = erc20Contract.Symbol(&bind.CallOpts{}); err != nil {
		return nil, err
	}

	if result.Decimals, err = erc20Contract.Decimals(&bind.CallOpts{}); err != nil {
		return nil, err
	}

	// TODO Insert token to database

	return result, nil
}
