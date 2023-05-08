package matters

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var (
	_               crawler.Crawler = (*service)(nil)
	mattersCacheKey                 = "crawler:matters"
)

type service struct {
	config       *config.Config
	kuroraClient *kurora.Client
	tokenClient  *token.Client
}

func New(conf *config.Config) crawler.Crawler {
	return &service{
		config:      conf,
		tokenClient: token.New(),
	}
}

func (s *service) Name() string {
	return protocol.PlatformMatters
}

func (s *service) Network() string {
	return protocol.NetworkPolygon
}

func (s *service) Run() error {
	loggerx.Global().Info("matters: run")

	var err error
	ctx := context.Background()

	s.kuroraClient, err = kurora.Dial(ctx, s.config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error("matters: kurora.Dial error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))

		return err
	}

	for {
		transactions, cacheInfo, err := s.HandleKuroraEntries(ctx)
		if err != nil {
			loggerx.Global().Error("matters: HandleKuroraEntries error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))

			return err
		}

		if len(transactions) == 0 {
			time.Sleep(1 * time.Minute)

			continue
		}

		// insert db
		err = database.UpsertTransactions(ctx, transactions, true)
		if err != nil {
			continue
		}

		// set cache
		cache.Global().Set(ctx, mattersCacheKey, cacheInfo, 7*24*time.Hour)
	}
}

func (s *service) HandleKuroraEntries(ctx context.Context) ([]*model.Transaction, string, error) {
	tracer := otel.Tracer("matters")
	_, trace := tracer.Start(ctx, "matters:HandleKuroraEntries")
	var internalTransactions []*model.Transaction
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

	cursor, _ := cache.Global().Get(ctx, mattersCacheKey).Result()
	query := kurora.DatasetMattersEntryQuery{
		Limit: &DefaultLimit,
	}

	if len(cursor) > 0 {
		query.Cursor = &cursor
	}

	entries, err := s.kuroraClient.FetchDatasetMattersEntries(ctx, query)
	if err != nil {
		loggerx.Global().Error("matters: kuroraClient FetchDatasetMattersEntries error", zap.Error(err))
		return nil, cursor, err
	}

	loggerx.Global().Info("matters: kuroraClient FetchDatasetMattersEntries result", zap.Int("len", len(entries)), zap.String("cursor", cursor))

	articleQuery := kurora.DatasetMattersArticleQuery{
		State: lo.ToPtr("active"),
	}

	// handle transaction
	for _, entry := range entries {
		if len(entry.Author) == 0 {
			continue
		}

		entry.Author = strings.TrimPrefix(entry.Author, "@")

		transactionFrom := model.Transaction{
			BlockNumber: int64(entry.BlockNumber),
			Hash:        strings.ToLower(entry.TransactionHash.String()),
			Timestamp:   entry.Timestamp,
			Owner:       strings.ToLower(entry.From.String()),
			Network:     s.Network(),
			Tag:         filter.TagSocial,
			Type:        filter.SocialReward,
			AddressFrom: strings.ToLower(entry.From.String()),
			AddressTo:   AddressCuration,
			Success:     lo.ToPtr(true),
			Platform:    s.Name(),
			Source:      protocol.SourceKurora,
		}

		// handle transfer
		tokenMetadata, err := s.tokenClient.ERC20ToMetadata(ctx, s.Network(), strings.ToLower(entry.Token.String()))
		if err != nil {
			loggerx.Global().Warn("failed to handle token metadata", zap.Error(err), zap.String("network", s.Network()), zap.String("transaction_hash", transactionFrom.Hash), zap.String("address", entry.From.String()), zap.String("token", entry.Token.String()))

			continue
		}

		tokenValueDisplayTo := entry.Amount.Shift(-int32(tokenMetadata.Decimals))

		tokenMetadata.Value = &entry.Amount
		tokenMetadata.ValueDisplay = &tokenValueDisplayTo

		postOuter := &metadata.Post{
			Reward:         tokenMetadata,
			TypeOnPlatform: []string{"Curation"},
		}

		if len(entry.Author) > 0 && len(entry.Summary) > 0 {
			postOuter.Target = &metadata.Post{
				Title:          entry.Title,
				Summary:        entry.Summary,
				Body:           entry.ContentMarkdown,
				Author:         []string{fmt.Sprintf("%s%s", "https://matters.news/@", entry.Author), entry.Author},
				TypeOnPlatform: []string{"Post"},
			}
		}

		metadataRaw, err := json.Marshal(postOuter)
		if err != nil {
			loggerx.Global().Warn("failed to handle marshall", zap.Error(err), zap.String("network", s.Network()), zap.String("transaction_hash", transactionFrom.Hash), zap.String("address", entry.From.String()))

			continue
		}

		internalTransfer := model.Transfer{
			TransactionHash: strings.ToLower(entry.TransactionHash.String()),
			Timestamp:       entry.Timestamp,
			Tag:             filter.TagSocial,
			Type:            filter.SocialReward,
			Index:           int64(entry.LogIndex),
			Network:         s.Network(),
			AddressFrom:     strings.ToLower(entry.From.String()),
			AddressTo:       strings.ToLower(entry.To.String()),
			Platform:        s.Name(),
			Source:          protocol.SourceKurora,
			Metadata:        metadataRaw,
			RelatedUrls: []string{
				ethereum.BuildScanURL(s.Network(), strings.ToLower(entry.TransactionHash.String())),
				strings.ReplaceAll(entry.URI, "ipfs://", "https://ipfs.io/ipfs/"),
			},
		}

		articleQuery.Username = lo.ToPtr(entry.Author)
		articleQuery.Summary = lo.ToPtr(entry.Summary)

		articles, err := s.kuroraClient.FetchDatasetMattersArticles(ctx, articleQuery)
		if err != nil {
			loggerx.Global().Error("matters: kuroraClient FetchDatasetMattersArticles error", zap.Error(err))
		}

		if len(articles) > 0 {
			internalTransfer.RelatedUrls = append(internalTransfer.RelatedUrls, fmt.Sprintf("%s%s/%v", "https://matters.news/@", entry.Author, articles[0].ID))
		} else {
			loggerx.Global().Info("can not find article", zap.String("username", entry.Author), zap.String("summary", entry.Summary))
		}

		transactionFrom.Transfers = append(transactionFrom.Transfers, internalTransfer)

		internalTransactions = append(internalTransactions, &transactionFrom)

		transactionTo := transactionFrom

		transactionTo.Owner = strings.ToLower(entry.To.String())

		internalTransactions = append(internalTransactions, &transactionTo)
	}

	// set cache
	last, err := lo.Last(entries)
	if err == nil {
		cursor = kurora.LogCursor(last.TransactionHash, last.LogIndex)
	}

	return internalTransactions, cursor, nil
}
