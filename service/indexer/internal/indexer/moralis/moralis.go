package moralis

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/indexer"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/sdk/trace"
)

var _ indexer.Worker = &Indexer{}

type Indexer struct {
	DatabaseClient *database.Client
	MoralisClient  *moralis.Client
	TracerProvider *trace.TracerProvider
}

func (i *Indexer) Handle(ctx context.Context, message *protocol.Message) error {
	tracer := i.TracerProvider.Tracer("worker_moralis")

	ctx, spanWorkerHandler := tracer.Start(ctx, "worker")
	defer spanWorkerHandler.End()

	ctx, spanMoralisHandler := tracer.Start(ctx, "moralis")

	// TODO Query latest timestamp
	transfers, _, err := i.MoralisClient.GetNFTTransfers(context.Background(), common.HexToAddress(message.Address), nil)
	if err != nil {
		return err
	}

	spanMoralisHandler.End()

	ctx, spanDatabaseHandler := tracer.Start(ctx, "postgres")
	tx, err := i.DatabaseClient.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	switch message.Network {
	case protocol.NetworkEthereum:

	case protocol.NetworkPolygon:
	}

	for _, transfer := range transfers {
		if err := tx.Transfer.
			Create().
			SetTransactionHash(transfer.TransactionHash).
			SetTransactionLogIndex(transfer.LogIndex).
			SetAddressFrom(transfer.FromAddress).
			SetAddressTo(transfer.ToAddress).
			SetTokenAddress(transfer.TokenAddress).
			SetTokenID(transfer.TokenId).
			Exec(ctx); err != nil {
			logrus.Errorln(err)

			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	spanDatabaseHandler.End()

	return nil
}
