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
	"github.com/sirupsen/logrus"
)

const (
	Name = "blockscout"
)

var _ datasource.Datasource = (*Datasource)(nil)

type Datasource struct {
	blockscoutClientMap map[string]*blockscout.Client
}

func (d *Datasource) Name() string {
	return Name
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkEthereumClassic, protocol.NetworkXDAI, protocol.NetworkCrossbell,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	internalTransactionMap := make(map[string]model.Transaction)

	// Get transactions
	internalTransactions, err := d.handleTransactions(ctx, message)
	if err != nil {
		return nil, err
	}

	for _, internalTransaction := range internalTransactions {
		internalTransactionMap[internalTransaction.Hash] = internalTransaction
	}

	// Get token transfers
	internalTokenTransfers, err := d.handleTokenTransfers(ctx, message)
	if err != nil {
		logrus.Error(err)
		// return nil, err
	}

	for _, internalTokenTransfer := range internalTokenTransfers {
		value, exist := internalTransactionMap[internalTokenTransfer.TransactionHash]
		if !exist {
			continue
		}

		value.Transfers = append(value.Transfers, internalTokenTransfer)

		internalTransactionMap[internalTokenTransfer.TransactionHash] = value
	}

	transactions := make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		transactions = append(transactions, internalTransaction)
	}

	return transactions, nil
}

func (d *Datasource) handleTransactions(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	transactions := make([]model.Transaction, 0)

	// Use a different Client for different networks
	blockscoutClient := d.blockscoutClientMap[message.Network]

	internalTransactions, _, err := blockscoutClient.GetTransactionList(ctx, common.HexToAddress(message.Address), &blockscout.GetTransactionListOption{})
	if err != nil {
		return nil, err
	}

	for _, internalTransaction := range internalTransactions {
		sourceData, err := json.Marshal(internalTransaction)
		if err != nil {
			return nil, err
		}

		timestamp := time.Unix(internalTransaction.TimeStamp.BigInt().Int64(), 0)

		transactions = append(transactions, model.Transaction{
			Hash:        internalTransaction.Hash,
			BlockNumber: internalTransaction.BlockNumber.IntPart(),
			Timestamp:   timestamp,
			Index:       internalTransaction.TransactionIndex.IntPart(),
			AddressFrom: internalTransaction.From,
			AddressTo:   internalTransaction.To,
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
			Transfers: []model.Transfer{
				// This is a virtual transfer
				{
					TransactionHash:     internalTransaction.Hash,
					Timestamp:           timestamp,
					TransactionLogIndex: protocol.LogIndexVirtual,
					AddressFrom:         internalTransaction.From,
					AddressTo:           internalTransaction.To,
					Metadata:            metadata.Default,
					Network:             message.Network,
					Source:              d.Name(),
					SourceData:          sourceData,
					RelatedUrls:         GetTxRelatedURLs(message.Network, internalTransaction.Hash),
				},
			},
		})
	}

	return transactions, nil
}

func (d *Datasource) handleTokenTransfers(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	transfers := make([]model.Transfer, 0)

	// Use a different Client for different networks
	blockscoutClient := d.blockscoutClientMap[message.Network]

	internalTokenTransfers, _, err := blockscoutClient.GetTokenTransactionList(ctx, common.HexToAddress(message.Address), &blockscout.GetTokenTransactionListOption{})
	if err != nil {
		return nil, err
	}

	for _, internalTokenTransfer := range internalTokenTransfers {
		sourceData, err := json.Marshal(internalTokenTransfer)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, model.Transfer{
			TransactionHash:     internalTokenTransfer.Hash,
			Timestamp:           time.Unix(internalTokenTransfer.TimeStamp.BigInt().Int64(), 0),
			TransactionLogIndex: internalTokenTransfer.LogIndex,
			AddressFrom:         internalTokenTransfer.From,
			AddressTo:           internalTokenTransfer.To,
			Network:             message.Network,
			Source:              d.Name(),
			SourceData:          sourceData,
			RelatedUrls:         GetTxRelatedURLs(message.Network, internalTokenTransfer.Hash),
		})
	}

	return transfers, nil
}

// Returns related urls based on the network and contract tx hash.
func GetTxRelatedURLs(
	network string,
	transactionHash string,
) []string {
	var urls []string

	switch network {
	case protocol.NetworkCrossbell:
		urls = append(urls, "https://scan.crossbell.io/tx/"+transactionHash)
	case protocol.NetworkXDAI:
		urls = append(urls, "https://blockscout.com/xdai/mainnet/tx/"+transactionHash)
	case protocol.NetworkEthereum:
		urls = append(urls, "https://blockscout.com/eth/mainnet/tx/"+transactionHash)
	case protocol.NetworkEthereumClassic:
		urls = append(urls, "https://blockscout.com/etc/mainnet/tx/"+transactionHash)
	}

	return urls
}

func New() datasource.Datasource {
	return &Datasource{
		blockscoutClientMap: map[string]*blockscout.Client{
			protocol.NetworkEthereum:        blockscout.New(blockscout.EndpointDefault, blockscout.NetworkEthereum),
			protocol.NetworkEthereumClassic: blockscout.New(blockscout.EndpointDefault, blockscout.NetworkEthereumClassic),
			protocol.NetworkXDAI:            blockscout.New(blockscout.EndpointDefault, blockscout.NetworkXDAI),
			protocol.NetworkCrossbell:       blockscout.New(blockscout.EndpointCrossbell, blockscout.NetworkCrossbell),
		},
	}
}
