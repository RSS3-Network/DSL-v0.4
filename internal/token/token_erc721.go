package token

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721/zora"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var ENSContractAddress = strings.ToLower("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85")

func (c *Client) ERC721(ctx context.Context, network, contractAddress string, tokenID *big.Int) (*ERC721, error) {
	if network == protocol.NetworkEthereum && strings.ToLower(contractAddress) == ENSContractAddress {
		return c.ERC721Ens(ctx, contractAddress, tokenID)
	}

	if network == protocol.NetworkEthereum && strings.EqualFold(contractAddress, zora.Address.String()) {
		return c.ERC721Zora(ctx, network, tokenID)
	}

	ethclient, err := ethclientx.Global(network)
	if err != nil {
		return nil, err
	}

	erc721Contract, err := erc721.NewERC721(common.HexToAddress(contractAddress), ethclient)
	if err != nil {
		return nil, err
	}

	result := ERC721{
		ContractAddress: strings.ToLower(contractAddress),
		ID:              tokenID,
	}

	if result.Name, err = erc721Contract.Name(&bind.CallOpts{}); err != nil {
		loggerx.Global().Named(contractAddress).Warn("Get NFT Name error", zap.Error(err))
		return nil, err
	}

	if result.Symbol, err = erc721Contract.Symbol(&bind.CallOpts{}); err != nil {
		loggerx.Global().Named(contractAddress).Warn("Get NFT Symbol error", zap.Error(err))
		return nil, err
	}

	if tokenID != nil {
		tokenURI, err := erc721Contract.TokenURI(&bind.CallOpts{}, tokenID)
		if err != nil {
			loggerx.Global().Named(contractAddress).Warn("Get NFT Uri error", zap.Error(err))
			return nil, err
		}

		if result.URI, err = c.URI(contractAddress, tokenID, tokenURI); err != nil {
			loggerx.Global().Named(contractAddress).Warn("Get NFT Name error", zap.Error(err))
			return nil, err
		}

		if result.Metadata, err = c.Metadata(ctx, result.URI); err != nil {
			loggerx.Global().Named(contractAddress).Warn("Get NFT Metadata error", zap.Error(err))
		}
	}

	var metadata Metadata

	if err := json.Unmarshal(result.Metadata, &metadata); err != nil {
		loggerx.Global().Named(contractAddress).Warn("Get NFT Metadata Unmarshal error", zap.Error(err))
	}

	return &result, nil
}

func (c *Client) ERC721Ens(ctx context.Context, contractAddress string, tokenID *big.Int) (*ERC721, error) {
	response := DomainsMetadata{}
	url := fmt.Sprintf("https://metadata.ens.domains/mainnet/%v/%v", ENSContractAddress, tokenID)

	if err := retry.Do(func() error {
		client := resty.New()

		_, err := client.R().SetResult(&response).Get(url)

		return err
	},
		retry.Attempts(2),
		retry.DelayType(func(_ uint, _ error, _ *retry.Config) time.Duration {
			delay, err := rand.Int(rand.Reader, big.NewInt(250))
			if err != nil {
				delay = big.NewInt(0)
			}

			return time.Second + time.Duration(delay.Int64())*time.Millisecond
		})); err != nil {
		logrus.Errorf("[token] EnsToNFT url: %v, err: %v", url, err)
		return nil, err
	}

	metadata := Metadata{
		Name:            response.Name,
		Description:     response.Description,
		Image:           response.Image,
		ImageURL:        response.Image,
		Attributes:      response.Attributes,
		BackgroundColor: response.BackgroundImage,
		ExternalLink:    response.URL,
		ExternalURL:     response.URL,
	}

	m, _ := json.Marshal(metadata)

	erc721 := ERC721{
		Name:            response.Name,
		Symbol:          "ENS",
		ContractAddress: strings.ToLower(contractAddress),
		ID:              tokenID,
		Metadata:        m,
	}

	return &erc721, nil
}

func (c *Client) ERC721Zora(ctx context.Context, network string, tokenID *big.Int) (*ERC721, error) {
	ethereumClient, err := ethclientx.Global(network)
	if err != nil {
		return nil, err
	}

	zoraContract, err := zora.NewZora(zora.Address, ethereumClient)
	if err != nil {
		return nil, err
	}

	result := ERC721{
		ContractAddress: strings.ToLower(zora.Address.String()),
		ID:              tokenID,
	}

	if result.Name, err = zoraContract.Name(&bind.CallOpts{}); err != nil {
		return nil, err
	}

	if result.Symbol, err = zoraContract.Symbol(&bind.CallOpts{}); err != nil {
		return nil, err
	}

	if result.URI, err = zoraContract.TokenMetadataURI(&bind.CallOpts{}, tokenID); err != nil {
		return nil, err
	}

	result.Metadata, err = c.Metadata(ctx, result.URI)
	if err != nil {
		return nil, err
	}

	var metadata Metadata

	if err := json.Unmarshal(result.Metadata, &metadata); err != nil {
		return nil, err
	}

	return &result, nil
}
