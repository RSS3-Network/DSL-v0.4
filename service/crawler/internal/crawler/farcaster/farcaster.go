package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"

	"go.uber.org/zap"
)

var (
	_                 crawler.Crawler = (*service)(nil)
	farcasterCacheKey                 = "crawler:farcaster"
)

type service struct {
	config       *config.Config
	kuroraClient *kurora.Client
}

func New(conf *config.Config) crawler.Crawler {
	return &service{
		config: conf,
	}
}

func (s *service) Name() string {
	return protocol.PlatformFarcaster
}

func (s *service) Network() string {
	return protocol.NetworkFarcaster
}

func (s *service) Run() error {
	loggerx.Global().Info("farcaster: run")

	var err error
	ctx := context.Background()

	s.kuroraClient, err = kurora.Dial(ctx, s.config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error("farcaster: kurora.Dial error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))

		return err
	}

	for {
		transactions, cacheInfo, err := s.GetKuroraCasts(ctx)
		if err != nil {
			loggerx.Global().Error("farcaster: GetKuroraCasts error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))
			return err
		}

		if len(transactions) == 0 {
			time.Sleep(5 * time.Minute)

			continue
		}

		err = database.UpsertTransactions(ctx, transactions, true)
		if err != nil {
			continue
		}

		cache.Global().Set(ctx, farcasterCacheKey, cacheInfo, 0)
	}
}

func (s *service) GetKuroraCasts(ctx context.Context) ([]*model.Transaction, string, error) {
	var err error
	transactions := make([]*model.Transaction, 0)

	cursor, _ := cache.Global().Get(ctx, farcasterCacheKey).Result()
	query := kurora.DatasetFarcasterCastQuery{}

	if len(cursor) > 0 {
		query.Cursor = &cursor
	}

	casts, err := s.kuroraClient.FetchDatasetFarcasterCasts(ctx, query)
	if err != nil {
		loggerx.Global().Error("farcaster: kuroraClient FetchDatasetFarcasterCasts error", zap.Error(err), zap.String("cursor", cursor))
		return nil, cursor, err
	}

	loggerx.Global().Info("farcaster: kuroraClient FetchDatasetFarcasterCasts result", zap.Int("len", len(casts)), zap.String("cursor", cursor))

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	for _, rawCast := range casts {
		wg.Add(1)

		cast := Cast{
			Hash:                 rawCast.Hash,
			ThreadHash:           rawCast.ThreadHash,
			ParentHash:           rawCast.ParentHash,
			AuthorFid:            rawCast.AuthorFid,
			AuthorUsername:       rawCast.AuthorUsername,
			AuthorDisplayname:    rawCast.AuthorDisplayname,
			AuthorPfpUrl:         rawCast.AuthorPfpUrl,
			Text:                 rawCast.Text,
			PublishedAt:          rawCast.PublishedAt,
			RepliesCount:         rawCast.RepliesCount,
			ReactionsCount:       rawCast.ReactionsCount,
			RecastsCount:         rawCast.RecastsCount,
			WatchesCount:         rawCast.WatchesCount,
			ParentAuthorFid:      rawCast.ParentAuthorFid,
			ParentAuthorUsername: rawCast.ParentAuthorUsername,
		}

		go func(cast Cast) {
			defer func() {
				wg.Done()
			}()

			txs, err := s.buildTransactions(ctx, cast)
			if err != nil {
				return
			}

			mu.Lock()
			transactions = append(transactions, txs...)
			mu.Unlock()
		}(cast)
	}

	wg.Wait()

	if len(casts) == 0 {
		return transactions, cursor, nil
	}
	// set cache
	last, _ := lo.Last(casts)

	cursor = last.Hash.String()

	return transactions, cursor, nil
}

func (s *service) buildTransactions(ctx context.Context, cast Cast) ([]*model.Transaction, error) {
	internalTransactions := make([]*model.Transaction, 0)

	queryProfile := kurora.DatasetFarcasterProfileQuery{
		Fid: lo.ToPtr(cast.AuthorFid),
	}

	profiles, err := s.kuroraClient.FetchDatasetFarcasterProfiles(ctx, queryProfile)
	if err != nil {
		loggerx.Global().Error("farcaster: kuroraClient Fetch Profiles error", zap.Error(err))
		return nil, err
	}

	if len(profiles) == 0 || len(profiles[0].SignerAddresses) == 0 {
		return nil, nil
	}

	user := profiles[0]

	post := &metadata.Post{
		CreatedAt:      cast.PublishedAt.Format(time.RFC3339),
		Author:         []string{cast.AuthorUsername, strings.ToLower(user.Address.String())},
		Body:           cast.Text,
		TypeOnPlatform: []string{"cast"},
	}

	if cast.Hash == cast.ThreadHash {
		// post
		metadataPost, _ := json.Marshal(post)
		for _, evmAddress := range user.SignerAddresses {
			transaction := &model.Transaction{
				// use timestamp as block number, as there is no actual blocknumber on farcaster
				BlockNumber: cast.PublishedAt.UnixMilli(),
				Timestamp:   cast.PublishedAt,
				// use MerkleRoot as hash, as there is no actual hash on farcaster
				Hash:        cast.Hash.String(),
				AddressFrom: strings.ToLower(user.Address.String()),
				AddressTo:   strings.ToLower(user.Address.String()),
				Owner:       strings.ToLower(evmAddress),
				Tag:         filter.TagSocial,
				Type:        filter.SocialPost,
				Platform:    s.Name(),
				Network:     s.Network(),
				Source:      protocol.SourceKurora,
				Transfers: []model.Transfer{
					{
						TransactionHash: cast.Hash.String(),
						Timestamp:       cast.PublishedAt,
						Index:           protocol.IndexVirtual,
						AddressFrom:     strings.ToLower(user.Address.String()),
						AddressTo:       strings.ToLower(user.Address.String()),
						Tag:             filter.TagSocial,
						Type:            filter.SocialPost,
						Metadata:        metadataPost,
						Network:         s.Network(),
						Platform:        s.Name(),
						Source:          protocol.SourceKurora,
						RelatedUrls:     []string{fmt.Sprintf("https://www.discove.xyz/casts/%s", strings.ToLower(common.HexToAddress(cast.Hash.String()).String()))},
					},
				},
			}

			internalTransactions = append(internalTransactions, transaction)
		}
	} else {
		var (
			targetHash common.Hash
			socialType string
		)

		if cast.ParentHash == ethereum.HashGenesis {
			// recast
			targetHash = cast.ThreadHash
			socialType = filter.SocialShare
		} else {
			// comment
			targetHash = cast.ParentHash
			socialType = filter.SocialComment
		}

		targetCastQuery := kurora.DatasetFarcasterCastQuery{
			Hash: lo.ToPtr(targetHash),
		}

		targetCasts, err := s.kuroraClient.FetchDatasetFarcasterCasts(ctx, targetCastQuery)
		if err != nil {
			loggerx.Global().Error("farcaster: kuroraClient Fetch Casts error", zap.Error(err), zap.String(socialType, targetHash.String()))
			return nil, err
		}

		if len(targetCasts) == 0 {

			loggerx.Global().Warn("farcaster: kuroraClient Fetch Casts no hash", zap.String(socialType, targetHash.String()))
			return nil, nil
		}

		target := targetCasts[0]

		queryTargetProfile := kurora.DatasetFarcasterProfileQuery{
			Fid: lo.ToPtr(targetCasts[0].AuthorFid),
		}

		targetProfiles, err := s.kuroraClient.FetchDatasetFarcasterProfiles(ctx, queryTargetProfile)
		if err != nil {
			loggerx.Global().Error("farcaster: kuroraClient Fetch Profiles error", zap.Error(err))
			return nil, err
		}

		if len(targetProfiles) == 0 || len(targetProfiles[0].SignerAddresses) == 0 {
			return nil, nil
		}

		targetUser := targetProfiles[0]

		post.Target = &metadata.Post{
			CreatedAt:      target.PublishedAt.Format(time.RFC3339),
			Author:         []string{target.AuthorUsername, strings.ToLower(targetUser.Address.String())},
			Body:           target.Text,
			TypeOnPlatform: []string{"cast"},
		}

		metadataPost, _ := json.Marshal(post)

		for _, evmAddress := range user.SignerAddresses {
			transaction := &model.Transaction{
				BlockNumber: cast.PublishedAt.UnixMilli(),
				Timestamp:   cast.PublishedAt,
				Hash:        cast.Hash.String(),
				AddressFrom: strings.ToLower(user.Address.String()),
				AddressTo:   strings.ToLower(targetUser.Address.String()),
				Owner:       strings.ToLower(evmAddress),
				Tag:         filter.TagSocial,
				Type:        socialType,
				Platform:    s.Name(),
				Network:     s.Network(),
				Source:      protocol.SourceKurora,
				Transfers: []model.Transfer{
					{
						TransactionHash: cast.Hash.String(),
						Timestamp:       cast.PublishedAt,
						Index:           protocol.IndexVirtual,
						AddressFrom:     strings.ToLower(user.Address.String()),
						AddressTo:       strings.ToLower(targetUser.Address.String()),
						Tag:             filter.TagSocial,
						Type:            socialType,
						Metadata:        metadataPost,
						Network:         s.Network(),
						Platform:        s.Name(),
						Source:          protocol.SourceKurora,
						RelatedUrls:     []string{fmt.Sprintf("https://www.discove.xyz/casts/%s", strings.ToLower(common.HexToAddress(cast.Hash.String()).String()))},
					},
				},
			}
			internalTransactions = append(internalTransactions, transaction)
		}
	}

	return internalTransactions, nil
}
