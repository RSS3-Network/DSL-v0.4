package nft

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/nft/contract/nft"
)

const (
	NetworkEthereum  = "ethereum"
	NetworkPolygon   = "polygon"
	NetworkBinance   = "binance"
	NetworkCrossbell = "crossbell"

	MaxSize = 1024 * 8
)

var (
	ProjectID                  = ""
	ErrorInvalidMetadataFormat = errors.New("invalid metadata format")
)

func Initialize(projectID string) {
	ProjectID = projectID
}

func GetTokenMetadataByID(network string, contractAddress common.Address, tokenID *big.Int) ([]byte, error) {
	ethereumEndpoint := ""

	switch network {
	case NetworkEthereum:
		ethereumEndpoint = "https://eth.rss3.dev"
	case NetworkPolygon:
		ethereumEndpoint = fmt.Sprintf("https://polygon-mainnet.infura.io/v3/%s", ProjectID)
	case NetworkBinance:
		ethereumEndpoint = "https://bsc-dataseed.binance.org"
	case NetworkCrossbell:
		ethereumEndpoint = "https://rpc.crossbell.io"
	default:
		return nil, errors.New("network not support")
	}

	ethereumClient, err := ethclient.Dial(ethereumEndpoint)
	if err != nil {
		return nil, err
	}

	nftClient, err := nft.NewNFT(contractAddress, ethereumClient)
	if err != nil {
		return nil, err
	}

	var tokenURI string

	if tokenURI, err = nftClient.TokenURI(&bind.CallOpts{}, tokenID); err != nil {
		if tokenURI, err = nftClient.Uri(&bind.CallOpts{}, tokenID); err != nil {
			return nil, err
		}
	}

	return GetTokenMetadata(tokenURI)
}

func GetTokenMetadata(tokenURI string) (json.RawMessage, error) {
	if strings.HasPrefix(tokenURI, "ipfs://") {
		tokenURI = strings.Replace(tokenURI, "ipfs://", "https://rss3.infura-ipfs.io/ipfs/", 1)
	} else if strings.Contains(tokenURI, ";base64,") {
		contents := strings.Split(tokenURI, ";base64,")

		if len(contents) < 2 {
			return nil, ErrorInvalidMetadataFormat
		}

		return base64.StdEncoding.DecodeString(contents[1])
	}

	response, err := http.Get(tokenURI)
	if err != nil {
		return nil, errors.New("failed to get the url of the token metadata")
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
