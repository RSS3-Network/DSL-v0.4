package rara

import (
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
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/nft"
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
	"golang.org/x/net/context"
)

var (
	_            crawler.Crawler = (*service)(nil)
	raraCacheKey                 = "crawler:rara"
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
	return protocol.PlatformRara
}

func (s *service) Network() string {
	return protocol.NetworkPolygon
}

func (s *service) Run() error {
	loggerx.Global().Info("rara: run")

	var err error
	ctx := context.Background()

	s.kuroraClient, err = kurora.Dial(ctx, s.config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error("rara: kurora.Dial error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))

		return err
	}

	for {
		transactions, err := s.HandleKuroraEntries(ctx)
		if err != nil {
			loggerx.Global().Error("rara: HandleKuroraEntries error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))

			return err
		}

		if len(transactions) == 0 {
			time.Sleep(1 * time.Hour)

			continue
		}

		// insert db
		err = database.UpsertTransactions(ctx, transactions, true)
		if err != nil {
			continue
		}
	}
}

func (s *service) HandleKuroraEntries(ctx context.Context) ([]*model.Transaction, error) {
	tracer := otel.Tracer("rara")
	_, trace := tracer.Start(ctx, "rara:HandleKuroraEntries")
	var internalTransactions []*model.Transaction
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

	cursor, _ := cache.Global().Get(ctx, raraCacheKey).Result()
	query := kurora.DatasetRaraEntryQuery{
		Limit: &DefaultLimit,
	}

	if len(cursor) > 0 {
		query.Cursor = &cursor
	}

	entries, err := s.kuroraClient.FetchDatasetRara(ctx, query)
	if err != nil {
		loggerx.Global().Error("rara: kuroraClient FetchDatasetMattersEntries error", zap.Error(err))
		return nil, err
	}

	loggerx.Global().Info("rara: kuroraClient FetchDatasetMattersEntries result", zap.Int("len", len(entries)), zap.String("cursor", cursor))

	// handle transaction
	for _, entry := range entries {
		transaction := model.Transaction{
			BlockNumber: entry.BlockNumber.CoefficientInt64(),
			Hash:        strings.ToLower(entry.TransactionHash.String()),
			Timestamp:   entry.Timestamp,
			Owner:       strings.ToLower(entry.From.String()),
			Network:     s.Network(),
			Tag:         filter.TagSocial,
			Type:        filter.SocialComment,
			AddressFrom: strings.ToLower(entry.From.String()),
			AddressTo:   strings.ToLower(entry.To.String()),
			Success:     lo.ToPtr(true),
			Platform:    s.Name(),
			Source:      protocol.SourceKurora,
		}

		network := protocol.IdToNetwork(fmt.Sprintf("%#x", entry.NftChainId.CoefficientInt64()))
		// handle transfer
		var err error

		post := metadata.Post{
			TargetURL:      ethereum.BuildTokenURL(network, entry.NftAddress.String(), entry.NftId.BigInt().String())[0],
			TypeOnPlatform: []string{filter.SocialPost},
		}

		switch entry.NftAddress {
		case nft.CryptoKitties:
			ckMetadata, err := nft.GetCkMetadata(ctx, entry.NftId)
			if err != nil {
				loggerx.Global().Warn("failed to handle NFT metadata", zap.Error(err), zap.String("network", network), zap.String("transaction_hash", transaction.Hash), zap.String("address", entry.From.String()), zap.String("token", entry.NftAddress.String()))

				continue
			}
			post.Title = ckMetadata.Name
			post.Body = ckMetadata.Bio
			post.Media = []metadata.Media{{Address: ckMetadata.Image, MimeType: "image/png"}}
		case nft.CryptoPunks:
			post.Title = nft.CryptoPunksName
			post.Body = nft.CryptoPunksDes
			post.Media = []metadata.Media{{Address: fmt.Sprintf("https://cryptopunks.app/cryptopunks/cryptopunk%s.png", entry.NftId.String()), MimeType: "image/png"}}
		default:
			nft, err := s.tokenClient.NFT(ctx, network, entry.NftAddress.String(), entry.NftId.BigInt())
			if err != nil {
				loggerx.Global().Warn("failed to handle NFT metadata", zap.Error(err), zap.String("network", network), zap.String("transaction_hash", transaction.Hash), zap.String("address", entry.From.String()), zap.String("token", entry.NftAddress.String()))

				continue
			}

			post.Title = nft.Name
			post.Body = nft.Description
			post.Media = []metadata.Media{{Address: nft.Image, MimeType: "image/png"}}
		}

		comment := &metadata.Post{
			Target:         &post,
			Body:           entry.Comment,
			Tags:           entry.Tags,
			Author:         []string{strings.ToLower(entry.From.String()), fmt.Sprintf("%s%s", "https://app.rara.social/profile/", strings.ToLower(entry.From.String()))},
			TypeOnPlatform: []string{filter.SocialComment},
		}

		metadataRaw, err := json.Marshal(comment)
		if err != nil {
			loggerx.Global().Warn("failed to handle marshall", zap.Error(err), zap.String("network", s.Network()), zap.String("transaction_hash", transaction.Hash), zap.String("address", entry.From.String()))

			continue
		}

		internalTransfer := model.Transfer{
			TransactionHash: strings.ToLower(entry.TransactionHash.String()),
			Timestamp:       entry.Timestamp,
			Tag:             filter.TagSocial,
			Type:            filter.SocialComment,
			Index:           int64(entry.LogIndex),
			Network:         s.Network(),
			AddressFrom:     strings.ToLower(entry.From.String()),
			AddressTo:       strings.ToLower(entry.To.String()),
			Platform:        s.Name(),
			Source:          protocol.SourceKurora,
			Metadata:        metadataRaw,
			RelatedUrls: []string{
				ethereum.BuildScanURL(s.Network(), strings.ToLower(entry.TransactionHash.String())),
				fmt.Sprintf("%s/%s/%s/%s", "https://app.rara.social/nft", entry.NftChainId, strings.ToLower(entry.NftAddress.String()), entry.NftId),
			},
		}

		transaction.Transfers = append(transaction.Transfers, internalTransfer)

		transaction.Owner = strings.ToLower(entry.From.String())

		internalTransactions = append(internalTransactions, &transaction)
	}

	// set cache
	last, err := lo.Last(entries)
	if err == nil {
		cursor = kurora.LogCursor(last.TransactionHash, last.LogIndex)
		cache.Global().Set(ctx, raraCacheKey, cursor, 7*24*time.Hour)
	}

	return internalTransactions, nil
}
