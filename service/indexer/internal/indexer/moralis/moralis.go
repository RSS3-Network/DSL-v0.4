package moralis

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/indexer"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ indexer.Worker = &Indexer{}

type Indexer struct {
	DatabaseClient *gorm.DB
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
	tx := i.DatabaseClient.Begin()
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
		tokenValue, err := decimal.NewFromString(transfer.Amount)
		if err != nil {
			return err
		}

		tokenID, err := decimal.NewFromString(transfer.TokenId)
		if err != nil {
			return err
		}

		if err := tx.
			Model((*model.Transfer)(nil)).
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			Create(&model.Transfer{
				TransactionHash:     transfer.TransactionHash,
				TransactionLogIndex: transfer.LogIndex,
				AddressFrom:         transfer.FromAddress,
				AddressTo:           transfer.ToAddress,
				TokenAddress:        transfer.TokenAddress,
				TokenStandard:       "erc721",
				TokenValue:          tokenValue,
				TokenID:             tokenID,
				Network:             "ethereum",
			}).
			Error; err != nil {
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	spanDatabaseHandler.End()

	return nil
}
