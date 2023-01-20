package mirror

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/worker/arweave"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	kuroraClient *kurora.Client
}

func (d *Datasource) Name() string {
	return protocol.SourceKurora
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkArweave,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	tracer := otel.Tracer("datasource_lens")
	ctx, trace := tracer.Start(ctx, "datasource_lens:Handle")
	internalTransactions := make([]model.Transaction, 0)
	var err error
	defer func() { opentelemetry.Log(trace, message, len(internalTransactions), err) }()

	address := common.HexToAddress(message.Address)
	query := kurora.DatasetMirrorEntryQuery{
		HeightFrom:  lo.ToPtr(decimal.NewFromInt(int64(message.BlockNumber))),
		Contributor: &address,
	}
	for {
		// Fetch mirror entries from kurora
		entries, err := d.kuroraClient.FetchDatasetMirrorEntries(ctx, query)
		if err != nil {
			loggerx.Global().Error("FetchDatasetLensEvents error", zap.Error(err), zap.Any("query", query))

			return nil, err
		}

		if len(entries) == 0 {
			break
		}

		for _, entry := range entries {
			// Invalid entry
			if len(entry.ContentDigital) == 0 {
				zap.L().Warn("invalid mirror entry", zap.String("transaction_id", entry.TransactionID))

				continue
			}

			// Provide a default value for Origin content digital
			if len(entry.OriginContentDigital) == 0 {
				entry.OriginContentDigital = entry.ContentDigital
			}

			owner := strings.ToLower(entry.Contributor.String())

			post := metadata.Post{
				Title: entry.Title.String(),
				Body:  entry.Content.String(),
				Author: []string{
					fmt.Sprintf("https://mirror.xyz/%v", owner),
				},
				OriginNoteID: entry.OriginContentDigital,
			}

			postMetadata, err := json.Marshal(post)
			if err != nil {
				zap.L().Warn("marshal post metadata", zap.Error(err), zap.Any("post", post))

				continue
			}

			filterType := filter.SocialPost

			if entry.ContentDigital != entry.OriginContentDigital {
				filterType = filter.SocialRevise
			}

			transaction := model.Transaction{
				BlockNumber: entry.Height.BigInt().Int64(),
				Hash:        entry.TransactionID,
				Owner:       strings.ToLower(entry.Contributor.String()),
				AddressFrom: owner,
				AddressTo:   arweave.AddressMirror,
				Platform:    protocol.PlatformMirror,
				Network:     protocol.NetworkArweave,
				Source:      protocol.SourceKurora,
				Timestamp:   entry.Timestamp,
				Tag:         filter.TagSocial,
				Type:        filterType,
				Transfers: []model.Transfer{
					// This is a virtual transfer
					{
						TransactionHash: entry.TransactionID,
						Timestamp:       entry.Timestamp,
						Index:           protocol.IndexVirtual,
						AddressFrom:     owner,
						AddressTo:       arweave.AddressMirror,
						Metadata:        postMetadata,
						Network:         protocol.NetworkArweave,
						Source:          protocol.SourceKurora,
						Platform:        protocol.PlatformMirror,
						RelatedUrls:     []string{fmt.Sprintf("https://mirror.xyz/%s/%s", owner, entry.OriginContentDigital)},
						Tag:             filter.TagSocial,
						Type:            filterType,
						// SourceData:      source, // Not required
					},
				},
			}

			internalTransactions = append(internalTransactions, transaction)
		}

		lastEntry, err := lo.Last(entries)
		if err == nil {
			query.Cursor = &lastEntry.TransactionID
		}
	}

	return internalTransactions, nil
}

func New(ctx context.Context, endpoint string) (datasource.Datasource, error) {
	var internalDatasource Datasource
	var err error

	internalDatasource.kuroraClient, err = kurora.Dial(ctx, endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error(" kurora Dial error", zap.Error(err))
		return nil, err
	}

	return &internalDatasource, nil
}
