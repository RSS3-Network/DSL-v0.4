package moralis

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
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
	internalNFTTransfers, err := d.handleEthereumTokenTransfers(ctx, message)
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
	// TODO
	return nil, nil
}

func (d *Datasource) handleEthereumTokenTransfers(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	// TODO
	return nil, nil
}

func (d *Datasource) handleEthereumNFTTransfers(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	address := common.HexToAddress(message.Address)

	// Get nft transfers from Moralis
	nftTransfers, response, err := d.moralisClient.GetNFTTransfers(ctx, address, &moralis.GetNFTTransfersOption{
		Chain: protocol.NetworkToID(message.Network),
	})
	if err != nil {
		return nil, err
	}

	// Iterate through the next page of data
	for i := 1; int64(len(nftTransfers)) < response.Total && i < MaxPage; i++ {
		var internalNFTTransfers []moralis.NFTTransfer

		internalNFTTransfers, response, err = d.moralisClient.GetNFTTransfers(ctx, address, &moralis.GetNFTTransfersOption{
			Chain:  protocol.NetworkToID(message.Network),
			Cursor: response.Cursor,
		})

		if err != nil {
			return nil, err
		}

		nftTransfers = append(nftTransfers, internalNFTTransfers...)
	}

	transfers := make([]model.Transfer, 0)

	for _, nftTransfer := range nftTransfers {
		timestamp, err := time.Parse(time.RFC3339, nftTransfer.BlockTimestamp)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, model.Transfer{
			// TODO
			Timestamp: timestamp,
		})
	}

	return transfers, nil
}

func New(key string) datasource.Datasource {
	moralisClient := moralis.NewClient(key)

	return &Datasource{
		moralisClient: moralisClient,
	}
}
