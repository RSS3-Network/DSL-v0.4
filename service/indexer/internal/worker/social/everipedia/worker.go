package everipedia

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/everipedia/graphql"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	everipedia "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/everipedia/contract"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ worker.Worker = (*internal)(nil)

const (
	SourceName = "everipedia"
)

type internal struct{}

func (i *internal) Name() string {
	return SourceName
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkPolygon,
	}
}

func (i *internal) Initialize(ctx context.Context) error {
	return nil
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_everipedia")
	_, span := tracer.Start(ctx, "worker_everipedia:Handle")

	defer span.End()

	internalTransactions := make([]model.Transaction, 0)

	polygonClient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	for _, transaction := range transactions {
		internalTransaction := transaction
		internalTransaction.Transfers = make([]model.Transfer, 0)

		addressTo := common.HexToAddress(transaction.AddressTo)

		switch addressTo {
		case everipedia.AddressEveripedia:
			receipt, err := polygonClient.TransactionReceipt(context.Background(), common.HexToHash(transaction.Hash))
			if err != nil {
				loggerx.Global().Warn("worker_farcaster: failed to Get Receipt", zap.Error(err), zap.String("network", message.Network), zap.String("transaction_hash", transaction.Hash), zap.String("address", message.Address))
				continue
			}

			for _, transfer := range transaction.Transfers {
				for _, log := range receipt.Logs {
					switch log.Topics[0] {
					case everipedia.EventPosted:
						transfer.Index = int64(log.Index)

						var data graphqlx.GetUserActivitiesActivitiesByUserActivityContentWiki
						transfer.Metadata = transaction.SourceData

						err := json.Unmarshal(transaction.SourceData, &data)
						if err != nil {
							continue
						}
						transfer.RelatedUrls = ethereum.BuildURL(
							[]string{
								ethereum.BuildScanURL(transaction.Network, transaction.Hash),
							}, fmt.Sprintf("https://iq.wiki/wiki/%s", data.Id),
							fmt.Sprintf("https://iq.wiki/account/%s", common.HexToAddress(transaction.Owner).String()),
							fmt.Sprintf("https://ipfs.rss3.page/ipfs/%s", data.Ipfs),
						)

					default:
						continue
					}
				}
				internalTransaction.Transfers = append(internalTransaction.Transfers, transfer)

				internalTransaction.Tag, internalTransaction.Type = filter.UpdateTagAndType(transfer.Tag, internalTransaction.Tag, transfer.Type, internalTransaction.Type)
			}

			internalTransactions = append(internalTransactions, internalTransaction)
		default:
			continue
		}
	}

	return internalTransactions, nil
}

func (i *internal) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &internal{}
}
