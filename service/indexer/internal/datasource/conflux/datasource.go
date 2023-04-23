package conflux

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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

	// deduplicate transactions
	seenTxs := make(map[string]struct{})
	for _, transfer := range accountTxns.Data.List {
		seenTxs[strings.ToLower(transfer.TransactionHash)] = struct{}{}
	}

	// get transaction receipt
	receiptInfo := make(map[string]*confluxClient.ConfluxTransactionReceipt)

	for txHash := range seenTxs {
		// get transaction receipt
		receipt, err := d.client.GetTransactionReceipt(context.Background(), txHash)
		if err != nil {
			loggerx.Global().Error("getInternalTransaction error", zap.Error(err), zap.Any("txHash", txHash))
			continue
		}
		receiptInfo[strings.ToLower(txHash)] = receipt
	}

	type EpochTime struct {
		EpochNumber int64
		Timestamp   int64
	}

	// group by hash in transfers
	groupByTransfers := make(map[string][]model.Transfer)
	seenTxnsEpochTimes := make(map[string]EpochTime)
	for _, transfer := range accountTxns.Data.List {
		seenTxnsEpochTimes[strings.ToLower(transfer.TransactionHash)] = EpochTime{
			EpochNumber: transfer.EpochNumber,
			Timestamp:   transfer.Timestamp,
		}

		value, err := decimal.NewFromString(transfer.Amount)
		if err != nil {
			loggerx.Global().Error("convert amount error", zap.Error(err), zap.Any("transfer", transfer))
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
			loggerx.Global().Error("getInternalTransaction error", zap.Error(err), zap.Any("metadata", _metadata))
			continue
		}

		groupByTransfers[strings.ToLower(transfer.TransactionHash)] = append(groupByTransfers[strings.ToLower(transfer.TransactionHash)], model.Transfer{
			TransactionHash: transfer.TransactionHash,
			Timestamp:       time.Unix(transfer.Timestamp, 0),
			Index:           transfer.TxIndex,
			AddressFrom:     transfer.From,
			AddressTo:       transfer.To,
			Metadata:        metadataRaw,
			Network:         protocol.NetworkConflux,
			Source:          protocol.SourceConflux,
			RelatedUrls: []string{
				fmt.Sprintf("https://confluxscan.io/transaction/%s", transfer.TransactionHash),
			},
			Tag:  filter.TagTransaction,
			Type: filter.TransactionTransfer,
		})
	}

	// get transaction receipt
	txnsInfo := make(map[string]model.Transaction)
	for txHash := range seenTxs {
		// get transaction receipt
		receipt, ok := receiptInfo[txHash]
		if !ok {
			loggerx.Global().Error("ignore receipt", zap.Any("txHash", txHash))
			continue
		}

		epochTime, ok := seenTxnsEpochTimes[txHash]
		if !ok {
			loggerx.Global().Error("ignore epochTime", zap.Any("txHash", txHash))
			continue
		}

		txIndex, err := hexutil.DecodeBig(receipt.Index)
		if err != nil {
			loggerx.Global().Error("convert txIndex error", zap.Any("receipt", receipt))
			continue
		}

		gasFee, err := hexutil.DecodeBig(receipt.GasFee)
		if err != nil {
			loggerx.Global().Error("convert gasFee error", zap.Any("receipt", receipt))
			continue
		}

		txnsInfo[strings.ToLower(txHash)] = model.Transaction{
			BlockNumber: epochTime.EpochNumber,
			Index:       txIndex.Int64(),
			Hash:        receipt.TransactionHash,
			Timestamp:   time.Unix(epochTime.Timestamp, 0),
			Owner:       receipt.From,
			AddressFrom: receipt.From,
			AddressTo:   receipt.To,
			Network:     protocol.NetworkConflux,
			Source:      protocol.SourceConflux,
			Tag:         filter.TagTransaction,
			Type:        filter.TransactionTransfer,
			Success:     lo.ToPtr[bool](true),
			Fee:         lo.ToPtr[decimal.Decimal](decimal.New(gasFee.Int64(), 0).Shift(-ConfluxTokenDecimals)),
		}
	}

	// merge transaction and transfers
	for txHash, txInfo := range txnsInfo {
		txInfo.Transfers = groupByTransfers[txHash]
		internalTransactions = append(internalTransactions, txInfo)
	}

	return internalTransactions, nil
}

func New() datasource.Datasource {
	return &Datasource{
		client: confluxClient.NewClient(),
	}
}
