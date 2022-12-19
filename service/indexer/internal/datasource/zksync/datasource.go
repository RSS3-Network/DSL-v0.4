package zksync

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/zksync"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"go.opentelemetry.io/otel"
)

var _ datasource.Datasource = (*Datasource)(nil)

const (
	Limit = 100

	OperationTypeTransfer = "Transfer"

	StatusCommitted = "committed"
	StatusFinalized = "finalized"
	StatusRejected  = "rejected"
)

type Datasource struct {
	zksyncClient *zksync.Client
}

func (d *Datasource) Name() string {
	return protocol.SourceZksync
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkZkSync,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("zksync_datasource")
	_, trace := tracer.Start(ctx, "zksync_datasource:Handle")

	defer func() { opentelemetry.Log(trace, message, transactions, err) }()

	transactions = make([]model.Transaction, 0)

	address := common.HexToAddress(message.Address)

	from := zksync.FromLatest

	for {
		internalTransactions, _, err := d.zksyncClient.GetAddressTransactions(ctx, address, from, Limit, zksync.DirectionOlder)
		if err != nil {
			return nil, err
		}

		if len(internalTransactions.List) == 0 {
			break
		}

		for _, internalTransaction := range internalTransactions.List {
			if internalTransaction.CreatedAt.Before(message.Timestamp) {
				continue
			}

			if internalTransaction.Operation.From != "" && !strings.EqualFold(internalTransaction.Operation.From, message.Address) && !allowlist.AllowList.Contains(internalTransaction.Operation.From) {
				continue
			}

			transactionData, _, err := d.zksyncClient.GetTransactionData(ctx, common.HexToHash(internalTransaction.TransactionHash))
			if err != nil {
				return nil, err
			}

			// TODO only support `Transfer` now
			// https://docs.zksync.io/apiv02-docs/#transactions-api-v0.2-transactions-txhash-data
			if transactionData.Transaction.Operation.Type != OperationTypeTransfer {
				continue
			}

			// Mark the transaction successful or not
			sourceData, err := json.Marshal(transactionData)
			if err != nil {
				return nil, err
			}

			success := true

			if internalTransaction.Status == StatusCommitted {
				continue
			}

			if internalTransaction.Status != StatusFinalized {
				success = false
			}

			transactions = append(transactions, model.Transaction{
				BlockNumber: internalTransaction.BlockNumber,
				Timestamp:   internalTransaction.CreatedAt,
				Hash:        internalTransaction.TransactionHash,
				AddressFrom: transactionData.Transaction.Operation.From,
				AddressTo:   transactionData.Transaction.Operation.To,
				Network:     message.Network,
				Success:     &success,
				Source:      d.Name(),
				SourceData:  sourceData,
				Transfers: []model.Transfer{
					// This is a virtual transfer
					{
						TransactionHash: internalTransaction.TransactionHash,
						Timestamp:       internalTransaction.CreatedAt,
						Index:           protocol.IndexVirtual,
						AddressFrom:     transactionData.Transaction.Operation.From,
						AddressTo:       transactionData.Transaction.Operation.To,
						Metadata:        metadata.Default,
						Network:         message.Network,
						Source:          d.Name(),
						SourceData:      sourceData,
						RelatedUrls:     []string{utils.GetTxHashURL(protocol.NetworkZkSync, internalTransaction.TransactionHash)},
					},
				},
			})
		}

		// If the condition is met, then all data has been obtained
		if internalTransactions.Pagination.Count >= len(transactions) {
			break
		}

		if len(transactions) >= protocol.DatasourceLimit {
			break
		}

		// Get the hash of the last transaction in the array, used for paging
		from = internalTransactions.List[len(internalTransactions.List)-1].TransactionHash
	}

	return transactions, nil
}

func New() datasource.Datasource {
	return &Datasource{
		zksyncClient: zksync.New(),
	}
}
