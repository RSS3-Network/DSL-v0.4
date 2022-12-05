package mirror

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"

	"go.uber.org/zap"

	"go.opentelemetry.io/otel"
)

var (
	mirrorCacheKey = "crawler:mirror" // height:cursor
	AddressMirror  = "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c"
)

var _ crawler.Crawler = (*service)(nil)

type service struct {
	kuroraClient *kurora.Client
}

func New(config *config.Config) crawler.Crawler {
	var (
		err      error
		instance service
	)

	if instance.kuroraClient, err = kurora.Dial(context.Background(), config.RPC.PregodETL.KuroraV2.HTTP); err != nil {
		zap.L().Error("dial kurora", zap.Error(err), zap.String("endpoint", config.RPC.PregodETL.KuroraV2.HTTP))

		return nil
	}

	return &instance
}

func (s *service) Name() string {
	return protocol.PlatformMirror
}

func (s *service) Run() error {
	ctx := context.Background()

	var query kurora.DatasetMirrorEntryQuery

	for {
		value, err := cache.Global().Get(ctx, mirrorCacheKey).Result()
		if err != nil {
			zap.L().Error("get cursor of mirror crawler from cache", zap.Error(err))

			return fmt.Errorf("get cursor of mirror crawler from cache: %w", err)
		}

		if splits := strings.Split(value, ":"); len(value) == 2 {
			query.Cursor = lo.ToPtr(splits[1])
		}

		// Fetch the mirror entries and then build them as transactions
		transactions, err := s.buildTransactions(ctx, query)
		if err != nil {
			zap.L().Error("build transactions", zap.Error(err), zap.String("cursor", lo.FromPtr(query.Cursor)))

			return fmt.Errorf("build transactions: %w", err)
		}

		if len(transactions) == 0 {
			time.Sleep(5 * time.Minute)

			continue
		}

		// Store transactions to databse
		err = database.UpsertTransactions(ctx, transactions)
		if err != nil {
			continue
		}
	}
}

func (s *service) buildTransactions(ctx context.Context, query kurora.DatasetMirrorEntryQuery) (transactions []*model.Transaction, err error) {
	ctx, trace := otel.Tracer("crawler").Start(ctx, "mirror:buildTransactions")
	defer opentelemetry.Log(trace, nil, transactions, err)

	// Fetch mirror entries from kurora
	entries, err := s.kuroraClient.FetchDatasetMirrorEntries(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("fetch mirror entries: %w", err)
	}

	transactions = make([]*model.Transaction, 0, len(entries))

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

		address := strings.ToLower(entry.Contributor.String())

		post := metadata.Post{
			Title: entry.Title,
			Body:  entry.Content,
			Author: []string{
				fmt.Sprintf("https://mirror.xyz/%v", address),
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
			AddressFrom: address,
			AddressTo:   AddressMirror,
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
					AddressFrom:     address,
					AddressTo:       AddressMirror,
					Metadata:        postMetadata,
					Network:         protocol.NetworkArweave,
					Source:          protocol.SourceKurora,
					// SourceData:      source, // Not required
					Platform:    protocol.PlatformMirror,
					RelatedUrls: []string{fmt.Sprintf("https://mirror.xyz/%s/%s", address, entry.OriginContentDigital)},
					Tag:         filter.TagSocial,
					Type:        filterType,
				},
			},
		}

		transactions = append(transactions, &transaction)
	}

	lastEntry, err := lo.Last(entries)
	if err != nil {
		// No update or no eligible data
		return nil, nil
	}

	// Set cursor to cache
	if err := cache.Global().Set(ctx, mirrorCacheKey, fmt.Sprintf("%v:%v", lastEntry.Height, lastEntry.TransactionID), 7*24*time.Hour).Err(); err != nil {
		return nil, fmt.Errorf("set cursor to cache: %w", err)
	}

	return transactions, nil
}
