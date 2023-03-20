package token

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/blur"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/samber/lo"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

func (c *Client) ERC20(ctx context.Context, network string, contractAddress string) (*ERC20, error) {
	// Special handle non-standard tokens
	// nolint:gocritic
	switch {
	case network == protocol.NetworkEthereum && strings.EqualFold(contractAddress, blur.AddressPool.String()):
		return &ERC20{
			Name:            "Blur Pool",
			Symbol:          "Blur Pool", // No symbol function implemented
			Decimals:        18,
			ContractAddress: contractAddress,
		}, nil
	}

	var token model.Token

	// Get token from database
	if database.Global() != nil {
		err := database.Global().
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
			return lo.ToPtr(tokenToERC20(token)), nil
		}
	}

	result := &ERC20{
		ContractAddress: contractAddress,
	}

	ethclient, err := ethclientx.Global(network)
	if err != nil {
		return nil, fmt.Errorf("get ethclient: %w", err)
	}

	erc20Contract, err := erc20.NewERC20(common.HexToAddress(result.ContractAddress), ethclient)
	if err != nil {
		return nil, fmt.Errorf("new erc20 contract: %w", err)
	}

	if result.Name, err = erc20Contract.Name(&bind.CallOpts{}); err != nil {
		return nil, fmt.Errorf("get erc20 token name: %w", err)
	}

	if result.Symbol, err = erc20Contract.Symbol(&bind.CallOpts{}); err != nil {
		return nil, fmt.Errorf("get erc20 token symbol: %w", err)
	}

	if result.Decimals, err = erc20Contract.Decimals(&bind.CallOpts{}); err != nil {
		return nil, fmt.Errorf("get erc20 token decimals: %w", err)
	}

	return result, nil
}

func (c *Client) ERC20ToMetadata(context context.Context, network, contractAddress string) (*metadata.Token, error) {
	erc20, err := c.ERC20(context, network, contractAddress)
	if err != nil {
		loggerx.Global().Error("get erc20 token error", zap.Error(err))

		return nil, err
	}

	tokenMetadata := &metadata.Token{
		Name:            erc20.Name,
		Symbol:          erc20.Symbol,
		Decimals:        erc20.Decimals,
		Standard:        protocol.TokenStandardERC20,
		ContractAddress: erc20.ContractAddress,
		Image:           erc20.Logo,
	}

	return tokenMetadata, nil
}

func tokenToERC20(token model.Token) ERC20 {
	return ERC20{
		Name:            token.Name,
		Symbol:          token.Symbol,
		Decimals:        token.Decimal,
		ContractAddress: token.ContractAddress,
		Logo:            token.Logo,
	}
}
