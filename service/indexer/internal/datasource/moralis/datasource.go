package moralis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/shopspring/decimal"
)

const (
	Source = "moralis"

	MaxPage = 5
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	moralisClient *moralis.Client
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	switch message.Network {
	case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain:
		return d.handleEthereum(ctx, message)
	default:
		return []model.Transaction{}, nil
	}
}

func (d *Datasource) handleEthereum(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	transactionMap := make(map[string]model.Transaction)

	// Get transactions for this address
	internalTransactions, err := d.handleEthereumTransactions(ctx, message)
	if err != nil {
		return nil, err
	}

	// Cache to map to optimize transfers processing performance
	for _, internalTransaction := range internalTransactions {
		transactionMap[internalTransaction.Hash] = internalTransaction
	}

	// Get token transfer for this address
	internalTokenTransfers, err := d.handleEthereumTokenTransfers(ctx, message)
	if err != nil {
		return nil, err
	}

	// Put token transfers into map
	for _, tokenTransfer := range internalTokenTransfers {
		transaction, exist := transactionMap[tokenTransfer.TransactionHash]
		if !exist {
			continue
		}

		transaction.Transfers = append(transaction.Transfers, tokenTransfer)

		transactionMap[tokenTransfer.TransactionHash] = transaction
	}

	// Get nft transfer for this address
	internalNFTTransfers, err := d.handleEthereumNFTTransfers(ctx, message)
	if err != nil {
		return nil, err
	}

	// Put nft transfers into map
	for _, nftTransfer := range internalNFTTransfers {
		transaction, exist := transactionMap[nftTransfer.TransactionHash]
		if !exist {
			continue
		}

		transaction.Transfers = append(transaction.Transfers, nftTransfer)

		transactionMap[nftTransfer.TransactionHash] = transaction
	}

	// Lay the map flat
	transactions := make([]model.Transaction, 0)

	for _, transaction := range transactionMap {
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (d *Datasource) handleEthereumTransactions(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	address := common.HexToAddress(message.Address)

	// Get transactions from Moralis
	internalTransactions, response, err := d.moralisClient.GetTransactions(ctx, address, &moralis.GetTransactionsOption{
		Chain: protocol.NetworkToID(message.Network),
	})
	if err != nil {
		return nil, err
	}

	// Iterate through the next page of data
	for i := 0; int64(len(internalTransactions)) < response.Total && i < MaxPage; i++ {
		var nextInternalTransactions []moralis.Transaction

		nextInternalTransactions, response, err = d.moralisClient.GetTransactions(ctx, address, &moralis.GetTransactionsOption{
			Chain:  protocol.NetworkToID(message.Network),
			Cursor: response.Cursor,
		})
		if err != nil {
			return nil, err
		}

		internalTransactions = append(internalTransactions, nextInternalTransactions...)
	}

	transactions := make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactions {
		blockNumber, err := decimal.NewFromString(internalTransaction.BlockNumber)
		if err != nil {
			return nil, err
		}

		timestamp, err := time.Parse(time.RFC3339, internalTransaction.BlockTimestamp)
		if err != nil {
			return nil, err
		}

		sourceData, err := json.Marshal(internalTransaction)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, model.Transaction{
			BlockNumber: blockNumber,
			Timestamp:   timestamp,
			Hash:        internalTransaction.Hash,
			AddressFrom: internalTransaction.FromAddress,
			AddressTo:   internalTransaction.ToAddress,
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
			Transfers: []model.Transfer{
				// This is a virtual transfer
				{
					TransactionHash:     internalTransaction.Hash,
					Timestamp:           timestamp,
					TransactionLogIndex: protocol.LogIndexVirtual,
					AddressFrom:         internalTransaction.FromAddress,
					AddressTo:           internalTransaction.ToAddress,
					Metadata:            metadata.Default,
					Network:             message.Network,
					Source:              d.Name(),
					SourceData:          sourceData,
					RelatedUrls:         nil, // TODO
				},
			},
		})
	}

	return transactions, nil
}

func (d *Datasource) handleEthereumTokenTransfers(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	address := common.HexToAddress(message.Address)

	// Get token transfers from Moralis
	internalTokenTransfers, response, err := d.moralisClient.GetTokenTransfers(ctx, address, &moralis.GetTokenTransfersOption{
		Chain: protocol.NetworkToID(message.Network),
	})
	if err != nil {
		return nil, err
	}

	// Iterate through the next page of data
	for i := 0; int64(len(internalTokenTransfers)) < response.Total && i < MaxPage; i++ {
		var nextInternalTokenTransfers []moralis.TokenTransfer

		nextInternalTokenTransfers, response, err = d.moralisClient.GetTokenTransfers(ctx, address, &moralis.GetTokenTransfersOption{
			Chain:  protocol.NetworkToID(message.Network),
			Cursor: response.Cursor,
		})

		if err != nil {
			return nil, err
		}

		internalTokenTransfers = append(internalTokenTransfers, nextInternalTokenTransfers...)
	}

	transfers := make([]model.Transfer, 0)

	for i, internalTokenTransfer := range internalTokenTransfers {
		timestamp, err := time.Parse(time.RFC3339, internalTokenTransfer.BlockTimestamp)
		if err != nil {
			return nil, err
		}

		sourceData, err := json.Marshal(internalTokenTransfer)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, model.Transfer{
			TransactionHash:     internalTokenTransfer.TransactionHash,
			Timestamp:           timestamp,
			TransactionLogIndex: decimal.NewFromInt(int64(i)), // This is because Moralis don't provide log index
			AddressFrom:         internalTokenTransfer.FromAddress,
			AddressTo:           internalTokenTransfer.ToAddress,
			Metadata:            metadata.Default,
			Network:             message.Network,
			Source:              d.Name(),
			SourceData:          sourceData,
			RelatedUrls:         nil, // TODO
		})
	}

	return transfers, nil
}

func (d *Datasource) handleEthereumNFTTransfers(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	address := common.HexToAddress(message.Address)

	// Get nft transfers from Moralis
	internalNFTTransfers, response, err := d.moralisClient.GetNFTTransfers(ctx, address, &moralis.GetNFTTransfersOption{
		Chain: protocol.NetworkToID(message.Network),
	})
	if err != nil {
		return nil, err
	}

	// Iterate through the next page of data
	for i := 0; int64(len(internalNFTTransfers)) < response.Total && i < MaxPage; i++ {
		var nextInternalNFTTransfers []moralis.NFTTransfer

		nextInternalNFTTransfers, response, err = d.moralisClient.GetNFTTransfers(ctx, address, &moralis.GetNFTTransfersOption{
			Chain:  protocol.NetworkToID(message.Network),
			Cursor: response.Cursor,
		})

		if err != nil {
			return nil, err
		}

		internalNFTTransfers = append(internalNFTTransfers, nextInternalNFTTransfers...)
	}

	// Moralis may return duplicate transfers data
	internalTransfersMap := make(map[string]map[int64]model.Transfer)

	for _, internalNFTTransfer := range internalNFTTransfers {
		timestamp, err := time.Parse(time.RFC3339, internalNFTTransfer.BlockTimestamp)
		if err != nil {
			return nil, err
		}

		sourceData, err := json.Marshal(internalNFTTransfer)
		if err != nil {
			return nil, err
		}

		// Data deduplication
		value, exist := internalTransfersMap[internalNFTTransfer.BlockHash]
		if exist {
			if _, exist := value[internalNFTTransfer.LogIndex.IntPart()]; exist {
				continue
			}
		}

		internalTransfersMap[internalNFTTransfer.BlockHash] = map[int64]model.Transfer{
			internalNFTTransfer.LogIndex.IntPart(): {
				TransactionHash:     internalNFTTransfer.TransactionHash,
				Timestamp:           timestamp,
				TransactionLogIndex: internalNFTTransfer.LogIndex,
				AddressFrom:         internalNFTTransfer.FromAddress,
				AddressTo:           internalNFTTransfer.ToAddress,
				Metadata:            metadata.Default,
				Network:             message.Network,
				Source:              d.Name(),
				SourceData:          sourceData,
				RelatedUrls:         nil, // TODO
			},
		}
	}

	transfers := make([]model.Transfer, 0)

	for _, internalTransfers := range internalTransfersMap {
		for _, internalTransfer := range internalTransfers {
			transfers = append(transfers, internalTransfer)
		}
	}

	return transfers, nil
}

func New(key string) datasource.Datasource {
	moralisClient := moralis.NewClient(key)

	return &Datasource{
		moralisClient: moralisClient,
	}
}
