package blockscout

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/blockscout"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/shopspring/decimal"
)

var _ datasource.Datasource = (*Datasource)(nil)

type Datasource struct {
	blockscoutClientMap map[string]*blockscout.Client
}

func (d *Datasource) Name() string {
	return "blocksout"
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkEthereumClassic, protocol.NetworkXDAI,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	blockscoutClient := d.blockscoutClientMap[message.Network]

	tokenTransactions, _, err := blockscoutClient.GetTokenTransactionList(ctx, common.HexToAddress(message.Address), &blockscout.GetTokenTransactionListOption{})
	if err != nil {
		return nil, err
	}

	tokenTransactionMap := map[string]struct{}{}
	for _, transaction := range tokenTransactions {
		tokenTransactionMap[transaction.Hash] = struct{}{}

		sourceData, err := json.Marshal(transaction)
		if err != nil {
			return nil, err
		}

		internalTransfers = append(internalTransfers, model.Transfer{
			TransactionHash:     transaction.Hash,
			Timestamp:           time.Unix(transaction.TimeStamp.BigInt().Int64(), 0),
			TransactionLogIndex: transaction.LogIndex,
			AddressFrom:         transaction.From,
			AddressTo:           transaction.To,
			Metadata:            metadata.Default,
			Network:             message.Network,
			Source:              d.Name(),
			SourceData:          sourceData,
		})
	}

	transactions, _, err := blockscoutClient.GetTransactionList(ctx, common.HexToAddress(message.Address), &blockscout.GetTransactionListOption{})
	if err != nil {
		return nil, err
	}

	for _, transaction := range transactions {
		if _, exist := tokenTransactionMap[transaction.Hash]; exist {
			continue
		}

		sourceData, err := json.Marshal(transaction)
		if err != nil {
			return nil, err
		}

		internalTransfers = append(internalTransfers, model.Transfer{
			TransactionHash:     transaction.Hash,
			Timestamp:           time.Unix(transaction.TimeStamp.BigInt().Int64(), 0),
			TransactionLogIndex: decimal.Zero,
			AddressFrom:         transaction.From,
			AddressTo:           transaction.To,
			Metadata:            metadata.Default,
			Network:             message.Network,
			Source:              d.Name(),
			SourceData:          sourceData,
		})
	}

	return internalTransfers, nil
}

func New() datasource.Datasource {
	return &Datasource{
		blockscoutClientMap: map[string]*blockscout.Client{
			protocol.NetworkEthereum:        blockscout.New(blockscout.NetworkEthereum),
			protocol.NetworkEthereumClassic: blockscout.New(blockscout.NetworkEthereumClassic),
			protocol.NetworkXDAI:            blockscout.New(blockscout.NetworkXDAI),
		},
	}
}
