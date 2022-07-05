package moralis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis/mrc20"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const (
	Source = "moralis"

	MaxPage = 5

	StatusFailed  = "0"
	StatusSuccess = "1"

	Endpoint = "https://polygon-mainnet.infura.io/v3/"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	moralisClient  *moralis.Client
	ethereumClient *ethclient.Client
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("moralis_datasource")
	_, trace := tracer.Start(ctx, "moralis_datasource:Handle")

	defer func() { opentelemetry.Log(trace, message, transactions, err) }()

	switch message.Network {
	case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain:
		return d.handleEthereum(ctx, message)
	default:
		return []model.Transaction{}, nil
	}
}

func (d *Datasource) handleEthereum(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("moralis_datasource")
	_, trace := tracer.Start(ctx, "moralis_datasource:handleEthereum")

	defer func() { opentelemetry.Log(trace, message, transactions, err) }()

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
	internalTokenTransfers, moralisTokenTransferMap, err := d.handleEthereumTokenTransfers(ctx, message)
	if err != nil {
		return nil, err
	}

	// Put token transfers into map
	for _, tokenTransfer := range internalTokenTransfers {
		transaction, exist := transactionMap[tokenTransfer.TransactionHash]
		if !exist {
			internalTransaction := moralisTokenTransferMap[tokenTransfer.TransactionHash]
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

			transaction = model.Transaction{
				BlockNumber: blockNumber.IntPart(),
				Timestamp:   timestamp,
				Hash:        internalTransaction.TransactionHash,
				Index:       0,
				AddressFrom: internalTransaction.FromAddress,
				AddressTo:   internalTransaction.ToAddress,
				Network:     message.Network,
				Source:      d.Name(),
				SourceData:  sourceData,
			}
		}

		transaction.Transfers = append(transaction.Transfers, tokenTransfer)

		transactionMap[tokenTransfer.TransactionHash] = transaction
	}

	// Get nft transfer for this address
	internalNFTTransfers, moralisNFTTransferMap, err := d.handleEthereumNFTTransfers(ctx, message)
	if err != nil {
		return nil, err
	}

	// Put nft transfers into map
	for _, nftTransfer := range internalNFTTransfers {
		transaction, exist := transactionMap[nftTransfer.TransactionHash]
		if !exist {
			internalTransaction := moralisNFTTransferMap[nftTransfer.TransactionHash]
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

			transaction = model.Transaction{
				BlockNumber: blockNumber.IntPart(),
				Timestamp:   timestamp,
				Hash:        internalTransaction.TransactionHash,
				Index:       internalTransaction.TransactionIndex.IntPart(),
				AddressFrom: internalTransaction.FromAddress,
				AddressTo:   internalTransaction.ToAddress,
				Network:     message.Network,
				Source:      d.Name(),
				SourceData:  sourceData,
			}
		}

		transaction.Transfers = append(transaction.Transfers, nftTransfer)

		transactionMap[nftTransfer.TransactionHash] = transaction
	}

	// Lay the map flat
	transactions = make([]model.Transaction, 0)

	for _, transaction := range transactionMap {
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (d *Datasource) handleEthereumTransactions(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	tracer := otel.Tracer("moralis_datasource")
	_, trace := tracer.Start(ctx, "moralis_datasource:handleEthereumTransactions")

	defer func() { opentelemetry.Log(trace, message, transactions, err) }()

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

		for _, nextInternalTransaction := range nextInternalTransactions { // nolint:gosimple
			internalTransactions = append(internalTransactions, nextInternalTransaction) // nolint:gosimple // TODO Filter data time
		}
	}

	transactions = make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactions {
		blockNumber, err := decimal.NewFromString(internalTransaction.BlockNumber)
		if err != nil {
			return nil, err
		}

		timestamp, err := time.Parse(time.RFC3339, internalTransaction.BlockTimestamp)
		if err != nil {
			return nil, err
		}

		index, err := decimal.NewFromString(internalTransaction.TransactionIndex)
		if err != nil {
			return nil, err
		}

		sourceData, err := json.Marshal(internalTransaction)
		if err != nil {
			return nil, err
		}

		// Mark the transaction successful or not
		success := true

		if internalTransaction.ReceiptStatus == StatusFailed {
			success = false
		}

		transaction := model.Transaction{
			BlockNumber: blockNumber.IntPart(),
			Timestamp:   timestamp,
			Hash:        internalTransaction.Hash,
			Index:       index.IntPart(),
			AddressFrom: internalTransaction.FromAddress,
			AddressTo:   internalTransaction.ToAddress,
			Network:     message.Network,
			Success:     &success,
			Source:      d.Name(),
			SourceData:  sourceData,
			Transfers: []model.Transfer{
				{
					// This is a virtual transfer
					TransactionHash: internalTransaction.Hash,
					Timestamp:       timestamp,
					Index:           protocol.IndexVirtual,
					AddressFrom:     internalTransaction.FromAddress,
					AddressTo:       internalTransaction.ToAddress,
					Metadata:        metadata.Default,
					Network:         message.Network,
					Source:          d.Name(),
					SourceData:      sourceData,
					RelatedUrls:     []string{GetTxHashURL(message.Network, internalTransaction.Hash)},
				},
			},
		}

		if message.Network == protocol.NetworkPolygon {
			internalTransfers, err := d.handleMRC20Transfers(ctx, transaction.Transfers[0])
			if err != nil {
				logrus.Errorf("[datasource] handleMRC20Transfers error: %v", err)
			}

			if len(internalTransfers) > 0 {
				transaction.Transfers = append(transaction.Transfers, internalTransfers...)
			}
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (d *Datasource) handleEthereumTokenTransfers(ctx context.Context, message *protocol.Message) (transfers []model.Transfer, moralisTokenTransferMap map[string]moralis.TokenTransfer, err error) {
	tracer := otel.Tracer("moralis_datasource")
	_, trace := tracer.Start(ctx, "moralis_datasource:handleEthereumTokenTransfers")

	defer func() { opentelemetry.Log(trace, message, transfers, err) }()

	address := common.HexToAddress(message.Address)

	// Get token transfers from Moralis
	internalTokenTransfers, response, err := d.moralisClient.GetTokenTransfers(ctx, address, &moralis.GetTokenTransfersOption{
		Chain: protocol.NetworkToID(message.Network),
	})
	if err != nil {
		return nil, nil, err
	}

	// Iterate through the next page of data
	for i := 0; int64(len(internalTokenTransfers)) < response.Total && i < MaxPage; i++ {
		var nextInternalTokenTransfers []moralis.TokenTransfer

		nextInternalTokenTransfers, response, err = d.moralisClient.GetTokenTransfers(ctx, address, &moralis.GetTokenTransfersOption{
			Chain:  protocol.NetworkToID(message.Network),
			Cursor: response.Cursor,
		})

		if err != nil {
			return nil, nil, err
		}

		internalTokenTransfers = append(internalTokenTransfers, nextInternalTokenTransfers...)
	}

	transfers = make([]model.Transfer, 0)
	moralisTokenTransferMap = map[string]moralis.TokenTransfer{}

	for i, internalTokenTransfer := range internalTokenTransfers {
		moralisTokenTransferMap[internalTokenTransfer.TransactionHash] = internalTokenTransfer

		timestamp, err := time.Parse(time.RFC3339, internalTokenTransfer.BlockTimestamp)
		if err != nil {
			return nil, nil, err
		}

		sourceData, err := json.Marshal(internalTokenTransfer)
		if err != nil {
			return nil, nil, err
		}

		transfers = append(transfers, model.Transfer{
			TransactionHash: internalTokenTransfer.TransactionHash,
			Timestamp:       timestamp,
			Index:           int64(i), // This is because Moralis don't provide log index
			AddressFrom:     internalTokenTransfer.FromAddress,
			AddressTo:       internalTokenTransfer.ToAddress,
			Metadata:        metadata.Default,
			Network:         message.Network,
			Source:          d.Name(),
			SourceData:      sourceData,
			RelatedUrls:     []string{GetTxHashURL(message.Network, internalTokenTransfer.TransactionHash)},
		})
	}

	return transfers, moralisTokenTransferMap, nil
}

func (d *Datasource) handleEthereumNFTTransfers(ctx context.Context, message *protocol.Message) (transfers []model.Transfer, moralisNFTTransferMap map[string]moralis.NFTTransfer, err error) {
	tracer := otel.Tracer("moralis_datasource")
	_, trace := tracer.Start(ctx, "moralis_datasource:handleEthereumNFTTransfers")

	defer func() { opentelemetry.Log(trace, message, transfers, err) }()

	address := common.HexToAddress(message.Address)

	// Get nft transfers from Moralis
	internalNFTTransfers, response, err := d.moralisClient.GetNFTTransfers(ctx, address, &moralis.GetNFTTransfersOption{
		Chain: protocol.NetworkToID(message.Network),
	})
	if err != nil {
		return nil, nil, err
	}

	// Iterate through the next page of data
	for i := 0; int64(len(internalNFTTransfers)) < response.Total && i < MaxPage; i++ {
		var nextInternalNFTTransfers []moralis.NFTTransfer

		nextInternalNFTTransfers, response, err = d.moralisClient.GetNFTTransfers(ctx, address, &moralis.GetNFTTransfersOption{
			Chain:  protocol.NetworkToID(message.Network),
			Cursor: response.Cursor,
		})

		if err != nil {
			return nil, nil, err
		}

		internalNFTTransfers = append(internalNFTTransfers, nextInternalNFTTransfers...)
	}

	// Moralis may return duplicate transfers data
	internalTransfersMap := make(map[string]map[int64]model.Transfer)
	moralisNFTTransferMap = make(map[string]moralis.NFTTransfer)

	for _, internalNFTTransfer := range internalNFTTransfers {
		moralisNFTTransferMap[internalNFTTransfer.TransactionHash] = internalNFTTransfer

		timestamp, err := time.Parse(time.RFC3339, internalNFTTransfer.BlockTimestamp)
		if err != nil {
			return nil, nil, err
		}

		sourceData, err := json.Marshal(internalNFTTransfer)
		if err != nil {
			return nil, nil, err
		}

		t := model.Transfer{
			TransactionHash: internalNFTTransfer.TransactionHash,
			Timestamp:       timestamp,
			Index:           internalNFTTransfer.LogIndex.IntPart(),
			AddressFrom:     internalNFTTransfer.FromAddress,
			AddressTo:       internalNFTTransfer.ToAddress,
			Metadata:        metadata.Default,
			Network:         message.Network,
			Source:          d.Name(),
			SourceData:      sourceData,
			RelatedUrls:     []string{utils.GetTxHashURL(message.Network, internalNFTTransfer.TransactionHash)},
		}

		// Data deduplication
		value, exist := internalTransfersMap[internalNFTTransfer.BlockHash]
		if exist {
			if _, exist := value[internalNFTTransfer.LogIndex.IntPart()]; exist {
				continue
			}
			internalTransfersMap[internalNFTTransfer.BlockHash][internalNFTTransfer.LogIndex.IntPart()] = t
		} else {
			internalTransfersMap[internalNFTTransfer.BlockHash] = map[int64]model.Transfer{
				internalNFTTransfer.LogIndex.IntPart(): t,
			}
		}
	}

	transfers = make([]model.Transfer, 0)

	for _, internalTransfers := range internalTransfersMap {
		for _, internalTransfer := range internalTransfers {
			transfers = append(transfers, internalTransfer)
		}
	}

	return transfers, moralisNFTTransferMap, nil
}

// Returns related urls based on the network and contract tx hash.
func GetTxRelatedURLs(
	network string,
	contractAddress string,
	tokenId string,
	transactionHash *string,
) []string {
	var urls []string
	if transactionHash != nil {
		urls = append(urls, GetTxHashURL(network, *transactionHash))
	}

	switch network {
	case protocol.NetworkEthereum:
		urls = append(urls, "https://etherscan.io/nft/"+contractAddress+"/"+tokenId)
		urls = append(urls, "https://opensea.io/assets/"+contractAddress+"/"+tokenId)
	case protocol.NetworkPolygon:
		urls = append(urls, "https://polygonscan.com/token/"+contractAddress)
		urls = append(urls, "https://opensea.io/assets/matic/"+contractAddress+"/"+tokenId)
	case protocol.NetworkBinanceSmartChain:
		urls = append(urls, "https://bscscan.com/nft/"+contractAddress+"/"+tokenId)
	case protocol.NetworkAvalanche:
	case protocol.NetworkFantom:
		urls = append(urls, "https://ftmscan.com/nft/"+contractAddress+"/"+tokenId)
	}

	return urls
}

// Returns related urls based on the network and contract tx hash.
func GetTxHashURL(
	network string,
	transactionHash string,
) string {
	switch network {
	case protocol.NetworkEthereum:
		return "https://etherscan.io/tx/" + (transactionHash)

	case protocol.NetworkPolygon:
		return "https://polygonscan.com/tx/" + (transactionHash)

	case protocol.NetworkBinanceSmartChain:
		return "https://bscscan.com/tx/" + (transactionHash)

	case protocol.NetworkAvalanche:
		return "https://avascan.info/blockchain/c/tx/" + (transactionHash)
	case protocol.NetworkFantom:
		return "https://ftmscan.com/tx/" + (transactionHash)
	case protocol.NetworkZkSync:
		return "https://zkscan.io/explorer/transactions/" + (transactionHash)
	default:
		return ""
	}
}

func New(moralisKey, infuraID string) datasource.Datasource {
	moralisClient := moralis.NewClient(moralisKey)

	ethereumClient, err := ethclient.Dial(Endpoint + infuraID)
	if err != nil {
		logrus.Errorf("[datasource] ethereum client error: %v", err)
	}

	return &Datasource{
		moralisClient:  moralisClient,
		ethereumClient: ethereumClient,
	}
}

func (d *Datasource) handleMRC20Transfers(ctx context.Context, virtualTransfer model.Transfer) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_crossbell")
	_, trace := tracer.Start(ctx, "worker_crossbell:handleReceipt")

	defer func() { opentelemetry.Log(trace, virtualTransfer.TransactionHash, transfers, err) }()

	receipt, err := d.ethereumClient.TransactionReceipt(ctx, common.HexToHash(virtualTransfer.TransactionHash))
	if err != nil {
		logrus.Errorf("[datasource] ethereum client TransactionReceipt error: %v", err)

		return nil, err
	}

	contract, err := mrc20.NewMRC20(mrc20.LogTransferContractAddress, d.ethereumClient)
	if err != nil {
		logrus.Errorf("[datasource] mrc20 NewMRC20 error: %v", err)

		return nil, err
	}

	for _, log := range receipt.Logs {
		logIndex := int64(log.Index)

		transfer := virtualTransfer

		if log.Topics[0] == mrc20.EventLogTransfer {
			event, err := contract.ParseLogTransfer(*log)
			if err != nil {
				zap.Error(err)

				return nil, err
			}

			transfer.AddressFrom = event.From.String()
			transfer.AddressTo = event.To.String()
			transfer.Index = logIndex
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}
