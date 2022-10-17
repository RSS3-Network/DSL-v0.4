package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/shopspring/decimal"
)

const (
	ChainIDEthereum     = 1
	ChainIDOptimism     = 10
	ChainIDPolygon      = 137
	ChainIDArbitrumOne  = 42161
	ChainIDArbitrumNova = 42170
)

func (w *Worker) buildTransfer(ctx context.Context, transaction model.Transaction, log types.Log, from, to common.Address, platform string, chainID uint64, tokenAddress *common.Address, tokenValue *big.Int, transferType string) (*model.Transfer, error) {
	var (
		tokenMetadata *metadata.Token
		err           error
	)

	if tokenAddress == nil {
		tokenMetadata, err = w.tokenClient.NatvieToMetadata(ctx, transaction.Network)
		if err != nil {
			return nil, fmt.Errorf("failed to get token metadata: %w", err)
		}
	} else {
		tokenMetadata, err = w.tokenClient.ERC20ToMetadata(ctx, transaction.Network, tokenAddress.String())
		if err != nil {
			return nil, fmt.Errorf("failed to get token metadata: %w", err)
		}
	}

	internalTokenValue := decimal.NewFromBigInt(tokenValue, 0)
	tokenMetadata.Value = &internalTokenValue
	tokenDisplay := internalTokenValue.Shift(-int32(tokenMetadata.Decimals))
	tokenMetadata.ValueDisplay = &tokenDisplay

	metadataRaw, err := json.Marshal(metadata.Bride{
		ChainID: chainID,
		Token:   *tokenMetadata,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Tag:             filter.TagBridge,
		Type:            transferType,
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(from.String()),
		AddressTo:       strings.ToLower(to.String()),
		Metadata:        metadataRaw,
		Network:         transaction.Network,
		Platform:        platform,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}
