package moralis

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/allowlist"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
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
		protocol.NetworkBinanceSmartChain, // protocol.NetworkEthereum, protocol.NetworkPolygon,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("datasource_moralis")
	_, trace := tracer.Start(ctx, "datasource_moralis:Handle")

	defer func() { opentelemetry.Log(trace, message, transactions, err) }()

	switch message.Network {
	case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain:
		return d.handleEthereum(ctx, message)
	default:
		return []model.Transaction{}, nil
	}
}

func (d *Datasource) handleEthereum(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("datasource_moralis")
	_, span := tracer.Start(ctx, "datasource_moralis:handleEthereum")

	defer opentelemetry.Log(span, message, transactions, err)

	transactionsMap := make(map[string]*model.Transaction)

	nativeTransactions, err := d.handleEthereumTransactions(ctx, message)
	if err != nil {
		return nil, err
	}

	for _, nativeTransaction := range nativeTransactions {
		transactionsMap[nativeTransaction.Hash] = &nativeTransaction
	}

	tokenTransfers, err := d.handleEthereumTokenTransfers(ctx, message)
	if err != nil {
		return nil, err
	}

	for _, transfer := range tokenTransfers {
		sourceData, err := json.Marshal(transfer)
		if err != nil {
			return nil, err
		}

		transactionsMap[transfer.TransactionHash] = &model.Transaction{
			BlockNumber: transfer.BlockNumber.Int64(),
			Hash:        transfer.TransactionHash,
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
		}
	}

	nftTransfers, err := d.handleEthereumNFTTransfers(ctx, message)
	if err != nil {
		return nil, err
	}

	for _, transfer := range nftTransfers {
		sourceData, err := json.Marshal(transfer)
		if err != nil {
			return nil, err
		}

		transactionsMap[transfer.TransactionHash] = &model.Transaction{
			BlockNumber: transfer.BlockNumber.Int64(),
			Hash:        transfer.TransactionHash,
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
		}
	}

	internalTransactions := make([]*model.Transaction, 0)

	for _, transaction := range transactionsMap {
		if transaction.BlockNumber < message.BlockNumber {
			continue
		}

		if transaction.AddressFrom != "" && !strings.EqualFold(transaction.AddressFrom, message.Address) && !allowlist.AllowList.Contains(transaction.AddressFrom) {
			continue
		}

		internalTransactions = append(internalTransactions, transaction)
	}

	internalTransactions, err = ethereum.BuildTransactions(ctx, message, internalTransactions)
	if err != nil {
		return nil, err
	}

	for _, transaction := range internalTransactions {
		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}

func (d *Datasource) handleEthereumTransactions(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("datasource_moralis")
	_, span := tracer.Start(ctx, "datasource_moralis:handleEthereumTransactions")

	defer opentelemetry.Log(span, message, transactions, err)

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
			Chain:     protocol.NetworkToID(message.Network),
			Cursor:    response.Cursor,
			FromBlock: big.NewInt(message.BlockNumber).String(),
		})
		if err != nil {
			return nil, err
		}

		for _, nextInternalTransaction := range nextInternalTransactions { // nolint:gosimple
			internalTransactions = append(internalTransactions, nextInternalTransaction) // nolint:gosimple // TODO Filter data time
		}
	}

	transactions = make([]model.Transaction, 0)

	for _, transaction := range internalTransactions {
		sourceData, err := json.Marshal(transaction)
		if err != nil {
			return nil, err
		}

		blockNumber, err := decimal.NewFromString(transaction.BlockNumber)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, model.Transaction{
			BlockNumber: blockNumber.IntPart(),
			Hash:        transaction.Hash,
			AddressFrom: transaction.FromAddress,
			AddressTo:   transaction.ToAddress,
			Network:     message.Network,
			Source:      d.Name(),
			SourceData:  sourceData,
		})
	}

	return transactions, nil
}

func (d *Datasource) handleEthereumTokenTransfers(ctx context.Context, message *protocol.Message) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("datasource_moralis")
	_, span := tracer.Start(ctx, "datasource_moralis:handleEthereumTokenTransfers")

	defer opentelemetry.Log(span, message, transfers, err)

	address := common.HexToAddress(message.Address)

	// Get transaction transfers from Moralis
	internalTokenTransfers, response, err := d.moralisClient.GetTokenTransfers(ctx, address, &moralis.GetTokenTransfersOption{
		Chain:     protocol.NetworkToID(message.Network),
		FromBlock: big.NewInt(message.BlockNumber).String(),
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

	for _, transfer := range internalTokenTransfers {
		sourceData, err := json.Marshal(transfer)
		if err != nil {
			return nil, err
		}

		blockNumber, err := decimal.NewFromString(transfer.BlockNumber)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, model.Transfer{
			TransactionHash: transfer.TransactionHash,
			BlockNumber:     blockNumber.BigInt(),
			Metadata:        metadata.Default,
			Network:         message.Network,
			Source:          d.Name(),
			SourceData:      sourceData,
		})
	}

	return transfers, nil
}

func (d *Datasource) handleEthereumNFTTransfers(ctx context.Context, message *protocol.Message) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("datasource_moralis")
	_, span := tracer.Start(ctx, "datasource_moralis:handleEthereumNFTTransfers")

	defer opentelemetry.Log(span, message, transfers, err)

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

	for _, transfer := range internalNFTTransfers {
		sourceData, err := json.Marshal(transfer)
		if err != nil {
			return nil, err
		}

		blockNumber, err := decimal.NewFromString(transfer.BlockNumber)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, model.Transfer{
			TransactionHash: transfer.TransactionHash,
			BlockNumber:     blockNumber.BigInt(),
			Metadata:        metadata.Default,
			Network:         message.Network,
			Source:          d.Name(),
			SourceData:      sourceData,
		})
	}

	return transfers, nil
}

func New(moralisKey string) datasource.Datasource {
	moralisClient := moralis.NewClient(moralisKey)

	return &Datasource{
		moralisClient: moralisClient,
	}
}
