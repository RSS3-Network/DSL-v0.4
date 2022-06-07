package zksync

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/shopspring/decimal"
)

var _ datasource.Datasource = (*Datasource)(nil)

type Datasource struct {
	zksyncClient *zksync.Client
}

func (d *Datasource) Name() string {
	return "zksync"
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkZkSync,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	from := zksync.FromLatest

	for {
		transactionList, _, err := d.zksyncClient.GetAddressTransactions(
			ctx,
			common.HexToAddress(message.Address),
			from,
			100,
			zksync.DirectionOlder,
		)

		if err != nil {
			return nil, err
		}

		if len(transactionList.List) == 0 {
			break
		}

		for _, item := range transactionList.List {
			transactionData, _, err := d.zksyncClient.GetTransactionData(ctx, common.HexToHash(item.TransactionHash))
			if err != nil {
				return nil, err
			}

			sourceData, err := json.Marshal(transactionData)
			if err != nil {
				return nil, err
			}

			internalTransfers = append(internalTransfers, model.Transfer{
				TransactionHash:     item.TransactionHash,
				Timestamp:           item.CreatedAt,
				TransactionLogIndex: decimal.Zero,
				AddressFrom:         transactionData.Transaction.Operation.From,
				AddressTo:           transactionData.Transaction.Operation.To,
				Metadata:            metadata.Default,
				Network:             message.Network,
				Source:              d.Name(),
				SourceData:          sourceData,
			})
		}

		if transactionList.Pagination.Count >= len(internalTransfers) {
			break
		}

		latestTransactionHash := transactionList.List[len(transactionList.List)-1].TransactionHash

		from = latestTransactionHash
	}

	return internalTransfers, nil
}

func New() datasource.Datasource {
	return &Datasource{
		zksyncClient: zksync.New(),
	}
}
