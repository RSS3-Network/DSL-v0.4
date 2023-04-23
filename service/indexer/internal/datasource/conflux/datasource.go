package conflux

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

	query := confluxClient.GetAccountParameter{
		MinEpochNumber: message.BlockNumber,
		Address:        message.Address,
		Limit:          conflux.MaxCount,
	}
	confluxAccountTxns, err := d.client.GetAccountCfxTransfers(ctx, query)
	if err != nil {
		loggerx.Global().Error("GetBlockTransactions error", zap.Error(err), zap.Any("query", query))
		return nil, err
	}

	internalTransactions, err = d.getInternalTransaction(confluxAccountTxns)
	if err != nil {
		loggerx.Global().Error("getInternalTransaction error", zap.Error(err), zap.Any("query", query))
		return nil, err
	}

	return internalTransactions, nil
}

func (d *Datasource) getInternalTransaction(accountTxns *conflux.ConfluxScanResp[conflux.ConfluxScanTransferBrief]) ([]model.Transaction, error) {
	internalTransactions := make([]model.Transaction, 0)

	for _, tx := range accountTxns.Data.List {
		receiptInfo, err := d.client.GetTransactionReceipt(context.Background(), tx.TransactionHash)
		if err != nil {
			loggerx.Global().Error("getInternalTransaction error", zap.Error(err), zap.Any("tx", tx))
			continue
		}
		gasFee, err := decimal.NewFromString(receiptInfo.GasFee)
		if err != nil {
			loggerx.Global().Error("getInternalTransaction error", zap.Error(err), zap.Any("value", receiptInfo.GasFee))
			continue
		}
		value, err := decimal.NewFromString(tx.Amount)
		if err != nil {
			loggerx.Global().Error("getInternalTransaction error", zap.Error(err), zap.Any("value", tx.Amount))
			continue
		}
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
			continue
		}
		internalTx := model.Transaction{
			BlockNumber: tx.EpochNumber,
			Index:       tx.TxIndex,
			Hash:        tx.TransactionHash,
			Timestamp:   time.Unix(tx.Timestamp, 0),
			Owner:       tx.From,
			AddressFrom: tx.From,
			AddressTo:   tx.To,
			Network:     protocol.NetworkConflux,
			Source:      protocol.SourceConflux,
			Tag:         filter.TagTransaction,
			Type:        filter.TransactionTransfer,
			Success:     lo.ToPtr[bool](true),
			Fee:         lo.ToPtr[decimal.Decimal](gasFee.Shift(-ConfluxTokenDecimals)),
		}
		// original transfer transaction
		internalTx.Transfers = []model.Transfer{
			{
				TransactionHash: tx.TransactionHash,
				Timestamp:       time.Unix(tx.Timestamp, 0),
				Index:           tx.TxIndex,
				AddressFrom:     tx.From,
				AddressTo:       tx.To,
				Metadata:        metadataRaw,
				Network:         protocol.NetworkConflux,
				Source:          protocol.SourceConflux,
				RelatedUrls: []string{
					fmt.Sprintf("https://confluxscan.io/transaction/%s", tx.TransactionHash),
				},
				Tag:  filter.TagTransaction,
				Type: filter.TransactionTransfer,
			},
		}
		internalTransactions = append(internalTransactions, internalTx)

	}

	return internalTransactions, nil
}

func New() datasource.Datasource {
	return &Datasource{
		client: confluxClient.NewClient(),
	}
}
