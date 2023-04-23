package aptos

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	aptosClient "github.com/naturalselectionlabs/pregod/common/datasource/aptos"
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
	client *aptosClient.Client
}

func (d *Datasource) Name() string {
	return protocol.SourceAptos
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkAptos,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	tracer := otel.Tracer("datasource_aptos")
	ctx, trace := tracer.Start(ctx, "datasource_aptos:Handle")
	internalTransactions := make([]model.Transaction, 0)
	var err error
	defer func() { opentelemetry.Log(trace, message, len(internalTransactions), err) }()

	// recover panic
	defer func() {
		if err := recover(); err != nil {
			loggerx.Global().Error("datasource_aptos panic", zap.Any("error", err))
		}
	}()

	query := aptosClient.GetAccountTransactionsParameter{
		Address: message.Address,
		Start:   decimal.NewFromInt(message.BlockNumber),
		Limit:   protocol.DatasourceLimit, // order by desc
	}

	result, err := d.client.GetAccountTractions(ctx, query)
	if err != nil {
		loggerx.Global().Error("GetAccountTractions error", zap.Error(err), zap.Any("query", query))

		return nil, err
	}

	for _, tx := range result {
		// (currently) only supports native token transfer
		if tx.Payload.Function != CoinTransferFunc && tx.Payload.Function != AccountTransferFunc {
			continue
		}

		if len(tx.Payload.Arguments) < 2 {
			continue
		}

		coinType := AptosCoin
		addressTo, _ := tx.Payload.Arguments[0].(string)
		value, _ := tx.Payload.Arguments[1].(string)

		if len(tx.Payload.TypeArguments) > 0 {
			coinType, _ = tx.Payload.TypeArguments[0].(string)
		}

		metadata, err := d.buildMetadata(coinType, value)
		if err != nil {
			loggerx.Global().Error("aptos datasource buildMetadata error", zap.Error(err))

			continue
		}

		metadataRaw, err := json.Marshal(metadata)
		if err != nil {
			loggerx.Global().Error("aptos metadata marshal error", zap.Error(err))

			continue
		}

		internalTransactions = append(internalTransactions, model.Transaction{
			BlockNumber: tx.Version.BigInt().Int64(),
			Index:       tx.SequenceNumber.BigInt().Int64(),
			Hash:        tx.Hash,
			Owner:       message.Address,
			AddressFrom: tx.Sender,
			AddressTo:   addressTo,
			Network:     protocol.NetworkAptos,
			Source:      protocol.SourceAptos,
			Timestamp:   time.UnixMicro(tx.Timestamp.BigInt().Int64()),
			Tag:         filter.TagTransaction,
			Type:        filter.TransactionTransfer,
			Success:     lo.ToPtr(tx.Success),
			Fee:         lo.ToPtr(tx.GasUsed.Shift(-int32(metadata.Decimals)).Mul(tx.GasUnitPrice)),
			Transfers: []model.Transfer{
				// This is a virtual transfer
				{
					TransactionHash: tx.Hash,
					Timestamp:       time.UnixMicro(tx.Timestamp.BigInt().Int64()),
					Index:           protocol.IndexVirtual,
					AddressFrom:     tx.Sender,
					AddressTo:       addressTo,
					Metadata:        metadataRaw,
					Network:         protocol.NetworkAptos,
					Source:          protocol.SourceAptos,
					RelatedUrls: []string{
						fmt.Sprintf("https://explorer.aptoslabs.com/txn/%v", tx.Hash),
					},
					Tag:  filter.TagTransaction,
					Type: filter.TransactionTransfer,
				},
			},
		})

		if len(internalTransactions) > datasource.DefaultLimit {
			break
		}
	}

	return internalTransactions, nil
}

func (d *Datasource) buildMetadata(coinType string, v string) (*metadata.Token, error) {
	value, err := decimal.NewFromString(v)
	if err != nil {
		return nil, err
	}

	coin, exist := coinMetadata[coinType]
	if !exist {
		return nil, fmt.Errorf("unsupported")
	}

	return &metadata.Token{
		Name:         coin.Name,
		Value:        lo.ToPtr(value),
		Symbol:       coin.Symbol,
		Decimals:     coin.Decimals,
		Standard:     coin.Standard,
		Image:        coin.Image,
		ValueDisplay: lo.ToPtr(value.Shift(-int32(coin.Decimals))),
	}, nil
}

func New() datasource.Datasource {
	return &Datasource{
		client: aptosClient.NewClient(),
	}
}
