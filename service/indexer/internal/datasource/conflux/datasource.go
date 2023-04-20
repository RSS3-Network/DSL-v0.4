package conflux

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/conflux"
	confluxClient "github.com/naturalselectionlabs/pregod/common/datasource/conflux"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	client *confluxClient.Client
}

func (d *Datasource) Name() string {
	return protocol.SourceConflux
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkConflux,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	tracer := otel.Tracer("datasource_conflux")
	ctx, trace := tracer.Start(ctx, "datasource_conflux:Handle")
	internalTransactions := make([]model.Transaction, 0)
	var err error
	defer func() { opentelemetry.Log(trace, message, len(internalTransactions), err) }()

	// recover panic
	defer func() {
		if err := recover(); err != nil {
			loggerx.Global().Error("datasource_conflux panic", zap.Any("error", err))
		}
	}()

	query := confluxClient.GetBlockTransactionsParameter{
		BlockNumber: message.BlockNumber,
	}
	confluxBlock, err := d.client.GetBlockTransactions(ctx, query)
	if err != nil {
		loggerx.Global().Error("GetBlockTransactions error", zap.Error(err), zap.Any("query", query))

		return nil, err
	}

	internalTransactions, err = d.getInternalTransaction(confluxBlock)
	if err != nil {
		loggerx.Global().Error("getInternalTransaction error", zap.Error(err), zap.Any("query", query))
		return nil, err
	}

	return internalTransactions, nil
}

func (d *Datasource) getInternalTransaction(confluxBlock *conflux.ConfluxBlock) ([]model.Transaction, error) {
	internalTransactions := make([]model.Transaction, 0)

	blockNumber, err := hexutil.DecodeBig(confluxBlock.BlockNumber)
	if err != nil {
		loggerx.Global().Error("hexutil.DecodeBig error", zap.Error(err), zap.Any("blockNumber", confluxBlock.BlockNumber))
		return nil, err
	}

	blockTimestamp, err := hexutil.DecodeBig(confluxBlock.Timestamp)
	if err != nil {
		loggerx.Global().Error("hexutil.DecodeBig error", zap.Error(err), zap.Any("timestamp", confluxBlock.Timestamp))
		return nil, err
	}

	for _, tx := range confluxBlock.Transactions {
		// ignore call transaction
		if tx.Data != "0x" {
			continue
		}
		txIndex, err := hexutil.DecodeBig(tx.TransactionIndex)
		if err != nil {
			loggerx.Global().Error("hexutil.DecodeBig error", zap.Error(err), zap.Any("txIndex", tx.TransactionIndex))
			return nil, err
		}

		txGasFee, err := hexutil.DecodeBig(tx.Receipt.GasFee)
		if err != nil {
			loggerx.Global().Error("hexutil.DecodeBig error", zap.Error(err), zap.Any("txGas", tx.Gas))
			return nil, err
		}

		valueBig, err := hexutil.DecodeBig(tx.Value)
		if err != nil {
			loggerx.Global().Error("hexutil.DecodeBig error", zap.Error(err), zap.Any("value", tx.Value))
			return nil, err
		}
		value := decimal.NewFromBigInt(valueBig, 0)

		_metadata := &metadata.Token{
			Name:         ConfluxTokenName,
			Value:        lo.ToPtr[decimal.Decimal](value),
			Symbol:       ConfluxTokenSymbol,
			Decimals:     ConfluxTokenDecimals,
			Standard:     ConfluxTokenStandard,
			Image:        ConfluxTokenImage,
			ValueDisplay: lo.ToPtr[decimal.Decimal](value.Shift(-ConfluxTokenDecimals)),
		}
		metadataRaw, err := json.Marshal(_metadata)
		if err != nil {
			loggerx.Global().Error("json.Marshal error", zap.Error(err), zap.Any("_metadata", _metadata))
			return nil, err
		}

		internalTransactions = append(internalTransactions, model.Transaction{
			BlockNumber: blockNumber.Int64(),
			Timestamp:   time.Unix(blockTimestamp.Int64(), 0),
			Hash:        tx.Hash,
			Index:       txIndex.Int64(),
			Owner:       tx.From,

			AddressFrom: tx.From,
			AddressTo:   tx.To,
			Network:     protocol.NetworkConflux,
			Source:      protocol.SourceConflux,
			Tag:         filter.TagTransaction,
			Type:        filter.TransactionTransfer,
			Success:     lo.ToPtr[bool](tx.Status == ConfluxStatusSuccess),
			Fee:         lo.ToPtr[decimal.Decimal](decimal.NewFromBigInt(txGasFee, 0).Shift(-ConfluxTokenDecimals)),

			Transfers: []model.Transfer{
				// This is a virtual transfer
				{
					TransactionHash: tx.Hash,
					Timestamp:       time.Unix(blockTimestamp.Int64(), 0),
					Index:           txIndex.Int64(),
					AddressFrom:     tx.From,
					AddressTo:       tx.To,
					Metadata:        metadataRaw,
					Network:         protocol.NetworkConflux,
					Source:          protocol.SourceConflux,
					RelatedUrls: []string{
						fmt.Sprintf("https://confluxscan.io/transaction/%s", tx.Hash),
					},
					Tag:  filter.TagTransaction,
					Type: filter.TransactionTransfer,
				},
			},
		})

	}
	return internalTransactions, nil
}

func New() datasource.Datasource {
	return &Datasource{
		client: confluxClient.NewClient(),
	}
}
