package blockscout

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"go.opentelemetry.io/otel"
)

var _ datasource.Datasource = (*Datasource)(nil)

type Datasource struct {
	blockscoutClientMap map[string]*blockscout.Client
}

func (d *Datasource) Name() string {
	return protocol.SourceBlockscout
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkXDAI,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("datasource_blockscout")
	_, trace := tracer.Start(ctx, "datasource_blockscout:Handle")

	defer opentelemetry.Log(trace, message, transactions, err)

	transactionMap := make(map[string]*model.Transaction)

	nativeTransactions, err := d.handleTransactions(ctx, message)
	if err != nil {
		return nil, err
	}

	for _, transaction := range nativeTransactions {
		internalTransaction := transaction

		transactionMap[transaction.Hash] = &internalTransaction
	}

	tokenTransfers, err := d.handleTokenTransfers(ctx, message)
	if err != nil {
		return nil, err
	}

	for _, transfer := range tokenTransfers {
		transactionMap[transfer.TransactionHash] = &model.Transaction{
			BlockNumber: transfer.BlockNumber.Int64(),
			Hash:        transfer.TransactionHash,
			Network:     message.Network,
			Source:      d.Name(),
		}
	}

	unindexedTransactions := make([]*model.Transaction, 0)

	for _, transaction := range transactionMap {
		if transaction.AddressFrom != "" && !strings.EqualFold(transaction.AddressFrom, message.Address) && !allowlist.AllowList.Contains(transaction.AddressFrom) {
			continue
		}

		unindexedTransactions = append(unindexedTransactions, transaction)
	}

	// indexedTransactions, err := ethereum.BuildTransactions(ctx, message, unindexedTransactions)
	// if err != nil {
	//	 return nil, err
	// }

	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range unindexedTransactions {
		internalTransactions = append(internalTransactions, *transaction)
	}

	return internalTransactions, nil
}

func (d *Datasource) handleTransactions(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("datasource_blockscout")
	_, span := tracer.Start(ctx, "datasource_blockscout:handleTransactions")

	defer opentelemetry.Log(span, message, transactions, err)

	transactions = make([]model.Transaction, 0)

	// Use a different Client for different networks
	blockscoutClient := d.blockscoutClientMap[message.Network]

	internalTransactions, _, err := blockscoutClient.GetTransactionList(ctx, common.HexToAddress(message.Address), &blockscout.GetTransactionListOption{
		StartBlock: message.BlockNumber,
	})
	if err != nil {
		return nil, err
	}

	for _, internalTransaction := range internalTransactions {
		sourceData, err := json.Marshal(internalTransaction)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, model.Transaction{
			Hash:        internalTransaction.Hash,
			BlockNumber: internalTransaction.BlockNumber.IntPart(),
			AddressFrom: strings.ToLower(internalTransaction.From),
			AddressTo:   strings.ToLower(internalTransaction.To),
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
		})
	}

	return transactions, err
}

func (d *Datasource) handleTokenTransfers(ctx context.Context, message *protocol.Message) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("datasource_blockscout")
	_, span := tracer.Start(ctx, "datasource_blockscout:handleTokenTransfers")

	defer opentelemetry.Log(span, message, transfers, err)

	// Use a different Client for different networks
	blockscoutClient := d.blockscoutClientMap[message.Network]

	internalTokenTransfers, _, err := blockscoutClient.GetTokenTransactionList(ctx, common.HexToAddress(message.Address), &blockscout.GetTokenTransactionListOption{
		StartBlock: message.BlockNumber,
	})
	if err != nil {
		return nil, err
	}

	for _, internalTokenTransfer := range internalTokenTransfers {
		sourceData, err := json.Marshal(internalTokenTransfer)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, model.Transfer{
			TransactionHash: internalTokenTransfer.Hash,
			BlockNumber:     internalTokenTransfer.BlockNumber.BigInt(),
			Network:         message.Network,
			Source:          d.Name(),
			SourceData:      sourceData,
		})
	}

	return transfers, nil
}

func New() (datasource.Datasource, error) {
	return &Datasource{
		blockscoutClientMap: map[string]*blockscout.Client{
			protocol.NetworkEthereum:        blockscout.New(blockscout.EndpointDefault, blockscout.NetworkEthereum),
			protocol.NetworkEthereumClassic: blockscout.New(blockscout.EndpointDefault, blockscout.NetworkEthereumClassic),
			protocol.NetworkXDAI:            blockscout.New(blockscout.EndpointDefault, blockscout.NetworkXDAI),
			protocol.NetworkCrossbell:       blockscout.New(blockscout.EndpointCrossbell, blockscout.NetworkCrossbell),
		},
	}, nil
}
