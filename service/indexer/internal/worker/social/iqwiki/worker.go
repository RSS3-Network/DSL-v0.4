package iqwiki

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	iqwiki "github.com/naturalselectionlabs/pregod/common/worker/iqwiki/contract"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/iqwiki/graphql"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ worker.Worker = (*internal)(nil)

const (
	SourceName = "IQ.Wiki"
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
		case iqwiki.AddressEveripedia:
			receipt, err := polygonClient.TransactionReceipt(context.Background(), common.HexToHash(transaction.Hash))
			if err != nil {
				loggerx.Global().Warn("worker_iqwiki: failed to Get Receipt", zap.Error(err), zap.String("network", message.Network), zap.String("transaction_hash", transaction.Hash), zap.String("address", message.Address))
				continue
			}

			for _, transfer := range transaction.Transfers {
				for _, log := range receipt.Logs {
					switch log.Topics[0] {
					case iqwiki.EventPosted:
						transfer.Index = int64(log.Index)
						transfer.SourceData = transaction.SourceData
						if err := i.HandleTransfer(ctx, &transfer); err != nil {
							loggerx.Global().Warn("worker_iqwiki: failed to HandleTransfer", zap.Error(err), zap.String("network", message.Network), zap.String("transaction_hash", transaction.Hash), zap.String("address", message.Address))

							continue
						}
					default:
						continue
					}
				}
				transfer.Type = filter.SocialWiki
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

func (i *internal) HandleTransfer(ctx context.Context, transfer *model.Transfer) (err error) {
	tracer := otel.Tracer("worker_iqwiki")
	_, trace := tracer.Start(ctx, "worker_iqwiki:HandleTransfer")

	defer func() { opentelemetry.Log(trace, transfer.AddressFrom, transfer.TransactionHash, err) }()

	var wiki graphqlx.GetUserActivitiesActivitiesByUserActivityContentWiki
	err = json.Unmarshal(transfer.SourceData, &wiki)
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("https://ipfs.rss3.page/ipfs/%s", wiki.Ipfs)
	response, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &json.RawMessage{}); err != nil {
		return err
	}

	var wikiContent Wiki
	if err = json.Unmarshal(data, &wikiContent); err != nil {
		return err
	}

	post := &metadata.Post{
		CreatedAt:      wiki.Created.Format(time.RFC3339),
		Author:         []string{wiki.Author.Id},
		Body:           wikiContent.Content,
		TypeOnPlatform: []string{transfer.Type},
		Title:          wiki.Title,
		Summary:        wiki.Summary,
		TargetURL:      fmt.Sprintf("https://iq.wiki/wiki/%s", wiki.Id),
	}
	transfer.Metadata, _ = json.Marshal(post)
	transfer.RelatedUrls = ethereum.BuildURL(
		[]string{
			ethereum.BuildScanURL(transfer.Network, wiki.TransactionHash),
		},
		fmt.Sprintf("https://iq.wiki/account/%s", common.HexToAddress(wiki.User.Id).String()),
	)
	return nil
}

type Wiki struct {
	Content string `json:"content"`
}

func (i *internal) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &internal{}
}
