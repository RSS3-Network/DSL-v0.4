package crossbell

import (
	"context"
	"embed"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/handler"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

//go:generate abigen --abi ./contract/erc721.abi --pkg contract --type ERC721 --out ./contract/erc721.go

const (
	Name = protocol.PlatfromCrossbell

	Endpoint = "https://rpc.crossbell.io"

	ABINameWeb3EntryProfile = "contract/event.abi"
)

//go:embed contract/event.abi
var ContractABIFileSystem embed.FS

var _ worker.Worker = (*Worker)(nil)

type Worker struct {
	ethereumClient      *ethclient.Client
	abiWeb3EntryProfile abi.ABI
	handler             handler.Interface
}

func (w *Worker) Name() string {
	return Name
}

func (w *Worker) Networks() []string {
	return []string{
		protocol.NetworkCrossbell,
	}
}

func (w *Worker) Initialize(ctx context.Context) (err error) {
	if w.ethereumClient, err = ethclient.Dial(Endpoint); err != nil {
		return err
	}

	web3EntryProfileABIFile, err := ContractABIFileSystem.Open(ABINameWeb3EntryProfile)
	if err != nil {
		return err
	}

	if w.abiWeb3EntryProfile, err = abi.JSON(web3EntryProfileABIFile); err != nil {
		return err
	}

	if w.handler, err = handler.New(w.ethereumClient, w.abiWeb3EntryProfile); err != nil {
		return err
	}

	return nil
}

func (w *Worker) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (internalTransactions []model.Transaction, err error) {
	tracer := otel.Tracer("crossbell_worker")
	_, trace := tracer.Start(ctx, "crossbell_worker:Handle")

	defer opentelemetry.Log(trace, message, internalTransactions, err)

	internalTransactions = make([]model.Transaction, 0)

	for _, transaction := range transactions {
		addressTo := common.HexToAddress(transaction.AddressTo)

		// Processing of transactions for contracts Profile and LinkList only
		if addressTo != handler.AddressProfileProxy && addressTo != handler.AddressLinkListTokenProxy {
			continue
		}

		transaction.Platform = Name

		// Retain the action model of the transfer type
		transferMap := make(map[int64]model.Transfer)

		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)

		// Get the raw data directly via Ethereum RPC node
		receipt, err := w.ethereumClient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
		if err != nil {
			return nil, err
		}

		for _, log := range receipt.Logs {
			logIndex := int64(log.Index)

			transfer, exist := transferMap[logIndex]
			if !exist {
				sourceData, err := json.Marshal(log)
				if err != nil {
					return nil, err
				}

				// Virtual transfer
				transfer = model.Transfer{
					TransactionHash: transaction.Hash,
					Timestamp:       transaction.Timestamp,
					Index:           logIndex,
					AddressFrom:     transaction.AddressFrom,
					Metadata:        metadata.Default,
					Network:         message.Network,
					Source:          protocol.SourceOrigin,
					SourceData:      sourceData,
				}
			}

			internalTransfer, err := w.handler.Handle(ctx, transaction, transfer)
			if err != nil {
				if errors.Is(err, handler.ErrorUnknownUnknownEvent) {
					continue
				}

				logrus.Error(err)

				continue
			}

			transfer.Platform = Name

			transaction.Transfers = append(transaction.Transfers, *internalTransfer)
		}

		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (w *Worker) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &Worker{}
}
