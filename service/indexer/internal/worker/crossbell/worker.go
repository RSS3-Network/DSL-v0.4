package crossbell

import (
	"context"
	"embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/handler"
)

//go:generate abigen --abi ./contract/erc721.abi --pkg contract --type ERC721 --out ./contract/erc721.go

const (
	Name = "crossbell"

	Endpoint     = "https://rpc.crossbell.io"
	EndpointIPFS = "https://ipfs.io"

	ABINameWeb3EntryProfile = "contract/event.abi"
)

var (
	//go:embed contract/event.abi
	ContractABIFileSystem embed.FS
)

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

func (w *Worker) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range transactions {
		addressTo := common.HexToAddress(transaction.AddressTo)

		if addressTo != handler.AddressLinkListTokenProxy && addressTo != handler.AddressWeb3EntryProfileProxy {
			continue
		}

		transaction.Transfers = make([]model.Transfer, 0)

		receipt, err := w.ethereumClient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
		if err != nil {
			return nil, err
		}

		for _, log := range receipt.Logs {
			internalTransfer, err := w.handler.Handle(ctx, &transaction, log)
			if err != nil {
				return nil, err
			}

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
