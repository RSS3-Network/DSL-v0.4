package token

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"
)

var nativeMap = map[string]Native{
	protocol.NetworkEthereum: {
		Name:     "Ethereum", // https://www.coingecko.com/en/coins/ethereum
		Symbol:   "ETH",
		Decimals: 18,
		Logo:     "https://assets.coingecko.com/coins/images/279/large/ethereum.png",
	},
	protocol.NetworkPolygon: {
		Name:     "Polygon", // https://www.coingecko.com/en/coins/polygon
		Symbol:   "MATIC",
		Decimals: 18,
		Logo:     "https://assets.coingecko.com/coins/images/4713/large/matic-token-icon.png",
	},
	protocol.NetworkBinanceSmartChain: {
		Name:     "BNB", // https://www.coingecko.com/en/coins/bnb
		Symbol:   "BNB",
		Logo:     "https://assets.coingecko.com/coins/images/825/large/bnb-icon2_2x.png",
		Decimals: 18,
	},
	protocol.NetworkXDAI: {
		Name:     "XDAI", // https://www.coingecko.com/en/coins/xdai
		Symbol:   "XDAI",
		Logo:     "https://assets.coingecko.com/coins/images/11062/large/Identity-Primary-DarkBG.png",
		Decimals: 18,
	},
	protocol.NetworkCrossbell: {
		Name:     "Crossbell",
		Symbol:   "CSB",
		Logo:     "", // TODO Add logo
		Decimals: 18,
	},
	protocol.NetworkZkSync: {
		Name:     "Ethereum", // https://www.coingecko.com/en/coins/ethereum
		Symbol:   "ETH",
		Logo:     "https://assets.coingecko.com/coins/images/279/large/ethereum.png",
		Decimals: 18,
	},
	protocol.NetworkOptimism: {
		Name:     "Ethereum", // https://www.coingecko.com/en/coins/ethereum
		Symbol:   "ETH",
		Decimals: 18,
		Logo:     "https://assets.coingecko.com/coins/images/279/large/ethereum.png",
	},
	protocol.NetworkAvalanche: {
		Name:     "Avalanche", // https://www.coingecko.com/en/coins/avalanche
		Symbol:   "AVAX",
		Decimals: 18,
		Logo:     "https://assets.coingecko.com/coins/images/12559/large/coin-round-red.png",
	},
	protocol.NetworkCelo: {
		Name:     "Celo", // https://www.coingecko.com/en/coins/celo
		Symbol:   "CELO",
		Decimals: 18,
		Logo:     "https://assets.coingecko.com/coins/images/4001/large/Fantom.png",
	},
	protocol.NetworkFantom: {
		Name:     "Fantom", // https://www.coingecko.com/en/coins/fantom
		Symbol:   "FTM",
		Decimals: 18,
		Logo:     "https://assets.coingecko.com/coins/images/11090/large/icon-celo-CELO-color-500.png",
	},
}

func (c *Client) Native(context context.Context, network string) (*Native, error) {
	token, exists := nativeMap[network]
	if !exists {
		return nil, ErrorTokenNotFound
	}

	return &token, nil
}

func (c *Client) NatvieToMetadata(context context.Context, network string) (*metadata.Token, error) {
	native, err := c.Native(context, network)
	if err != nil {
		loggerx.Global().Error("get native token error", zap.Error(err))

		return nil, err
	}

	tokenMetadata := &metadata.Token{
		Name:     native.Name,
		Symbol:   native.Symbol,
		Decimals: native.Decimals,
		Standard: protocol.TokenStandardNative,
		Image:    native.Logo,
	}

	return tokenMetadata, nil
}
