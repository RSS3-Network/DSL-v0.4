package alchemy

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/protocol"

	"github.com/ethereum/go-ethereum/rpc"
)

const (
	MethodGetAssetTransfers = "alchemy_getAssetTransfers"

	CategoryExternal = "external"
	CategoryInternal = "internal"
	CategoryERC20    = protocol.TokenStandardERC20
	CategoryERC721   = protocol.TokenStandardERC721
	CategoryERC1155  = protocol.TokenStandardERC1155

	MaxCount = 255
)

type GetAssetTransfersParameter struct {
	FromBlock         string   `json:"fromBlock,omitempty"`
	ToBlock           string   `json:"toBlock,omitempty"`
	FromAddress       string   `json:"fromAddress,omitempty"`
	ToAddress         string   `json:"toAddress,omitempty"`
	ContractAddresses []string `json:"contractAddresses,omitempty"`
	MaxCount          string   `json:"maxCount,omitempty"`
	ExcludeZeroValue  bool     `json:"excludeZeroValue"`
	PageKey           string   `json:"pageKey,omitempty"`
	WithMetadata      bool     `json:"withMetadata,omitempty"`
	Category          []string `json:"category,omitempty"`
}

type GetAssetTransfersResult struct {
	Transfers []struct {
		BlockNum        string      `json:"blockNum"`
		Hash            string      `json:"hash"`
		From            string      `json:"from"`
		To              string      `json:"to"`
		Value           float64     `json:"value"`
		Erc721TokenId   interface{} `json:"erc721TokenId"`
		Erc1155Metadata interface{} `json:"erc1155Metadata"`
		TokenId         interface{} `json:"tokenId"`
		Asset           string      `json:"asset"`
		Category        string      `json:"category"`
		RawContract     struct {
			Value   string      `json:"value"`
			Address interface{} `json:"address"`
			Decimal string      `json:"decimal"`
		}
	}
	PageKey string `json:"pageKey"`
}

func GetAssetTransfers(ctx context.Context, rpcClient *rpc.Client, parameter GetAssetTransfersParameter) (*GetAssetTransfersResult, error) {
	result := GetAssetTransfersResult{}

	if err := rpcClient.CallContext(ctx, &result, MethodGetAssetTransfers, parameter); err != nil {
		return nil, err
	}

	return &result, nil
}
