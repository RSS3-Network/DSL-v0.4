package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/naturalselectionlabs/pregod/common/database"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
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

type Cast struct {
	Hash                 common.Hash
	ThreadHash           common.Hash
	ParentHash           common.Hash
	AuthorFid            int64
	AuthorUsername       string
	AuthorDisplayname    string
	AuthorPfpUrl         string
	Text                 string
	PublishedAt          time.Time
	RepliesCount         int64
	ReactionsCount       int64
	RecastsCount         int64
	WatchesCount         int64
	ParentAuthorFid      int64
	ParentAuthorUsername string
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
		transactions, err := s.GetKuroraCasts(ctx)
		if err != nil {
			loggerx.Global().Error("farcaster: GetKuroraLogs error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))
			return err
		}

		err = database.UpsertTransactions(ctx, transactions)
		if err != nil {
			continue
		}

		if len(transactions) == 0 {
			time.Sleep(2 * time.Hour)

			continue
		}
	}
}

func (s *service) GetKuroraCasts(ctx context.Context) ([]*model.Transaction, error) {
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
		return nil, err
	}

	loggerx.Global().Info("farcaster: kuroraClient FetchEthereumLogs result", zap.Int("len", len(casts)), zap.String("cursor", cursor))

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

			queryProfile := kurora.DatasetFarcasterProfileQuery{
				Fid: lo.ToPtr(cast.AuthorFid),
			}

			profiles, err := s.kuroraClient.FetchDatasetFarcasterProfiles(ctx, queryProfile)
			if err != nil {
				loggerx.Global().Error("farcaster: kuroraClient Fetch Profiles error", zap.Error(err))
				return
			}

			if len(profiles) == 0 || len(profiles[0].SignerAddresses) == 0 {
				return
			}

			user := profiles[0]

			// post
			if cast.Hash == cast.ThreadHash {
				post := &metadata.Post{
					CreatedAt:      cast.PublishedAt.Format(time.RFC3339),
					Author:         []string{cast.AuthorUsername, strings.ToLower(user.Address.String())},
					Body:           cast.Text,
					TypeOnPlatform: []string{"cast"},
				}
				metadataPost, _ := json.Marshal(post)

				for index, evmAddress := range user.SignerAddresses {
					hash := fmt.Sprintf("%s%d", cast.Hash.String(), index)
					transaction := model.Transaction{
						// use timestamp as block number, as there is no actual blocknumber on farcaster
						BlockNumber: cast.PublishedAt.UnixMilli(),
						Timestamp:   cast.PublishedAt,
						// use MerkleRoot as hash, as there is no actual hash on farcaster
						Hash:        hash,
						AddressFrom: strings.ToLower(user.Address.String()),
						AddressTo:   strings.ToLower(user.Address.String()),
						Owner:       strings.ToLower(evmAddress),
						Tag:         filter.TagSocial,
						Type:        filter.SocialPost,
						Platform:    protocol.PlatformFarcaster,
						Network:     protocol.NetworkFarcaster,
						Source:      protocol.SourceKurora,
						Transfers: []model.Transfer{
							{
								TransactionHash: hash,
								Timestamp:       cast.PublishedAt,
								Index:           protocol.IndexVirtual,
								AddressFrom:     strings.ToLower(user.Address.String()),
								AddressTo:       strings.ToLower(user.Address.String()),
								Tag:             filter.TagSocial,
								Type:            filter.SocialPost,
								Metadata:        metadataPost,
								Network:         protocol.NetworkFarcaster,
								Platform:        protocol.PlatformFarcaster,
								Source:          protocol.SourceKurora,
								RelatedUrls:     []string{fmt.Sprintf("https://www.discove.xyz/casts/%s", cast.ParentHash)},
							},
						},
					}
					mu.Lock()
					transactions = append(transactions, &transaction)
					mu.Unlock()
				}
			} else {
				// comment
				targetCastQuery := kurora.DatasetFarcasterCastQuery{
					Hash: lo.ToPtr(cast.ParentHash),
				}
				targetCasts, err := s.kuroraClient.FetchDatasetFarcasterCasts(ctx, targetCastQuery)
				if err != nil {
					loggerx.Global().Error("farcaster: kuroraClient Fetch parent Casts error", zap.Error(err), zap.String("parent hash", cast.ParentHash.String()))
					return
				}

				if len(targetCasts) == 0 {

					loggerx.Global().Warn("farcaster: kuroraClient Fetch Casts no parent hash", zap.String("hash", cast.Hash.String()), zap.String("parent hash", cast.ParentHash.String()))
					return
				}

				parent := targetCasts[0]

				queryParentProfile := kurora.DatasetFarcasterProfileQuery{
					Fid: lo.ToPtr(targetCasts[0].AuthorFid),
				}

				parentProfiles, err := s.kuroraClient.FetchDatasetFarcasterProfiles(ctx, queryParentProfile)
				if err != nil {
					loggerx.Global().Error("farcaster: kuroraClient Fetch Profiles error", zap.Error(err))
					return
				}

				if len(parentProfiles) == 0 || len(parentProfiles[0].SignerAddresses) == 0 {
					return
				}

				parentUser := parentProfiles[0]

				post := &metadata.Post{
					Target: &metadata.Post{
						CreatedAt:      parent.PublishedAt.Format(time.RFC3339),
						Author:         []string{parent.AuthorUsername, strings.ToLower(parentUser.Address.String())},
						Body:           parent.Text,
						TypeOnPlatform: []string{"cast"},
					},
					CreatedAt:      cast.PublishedAt.Format(time.RFC3339),
					Author:         []string{cast.AuthorUsername, strings.ToLower(user.Address.String())},
					Body:           cast.Text,
					TypeOnPlatform: []string{"cast"},
				}
				metadataPost, _ := json.Marshal(post)
				for index, evmAddress := range user.SignerAddresses {
					hash := fmt.Sprintf("%s%d", cast.Hash.String(), index)
					transaction := model.Transaction{
						// use timestamp as block number, as there is no actual blocknumber on farcaster
						BlockNumber: cast.PublishedAt.UnixMilli(),
						Timestamp:   cast.PublishedAt,
						// use MerkleRoot as hash, as there is no actual hash on farcaster
						Hash:        hash,
						AddressFrom: strings.ToLower(user.Address.String()),
						AddressTo:   strings.ToLower(parentUser.Address.String()),
						Owner:       strings.ToLower(evmAddress),
						Tag:         filter.TagSocial,
						Type:        filter.SocialComment,
						Platform:    protocol.PlatformFarcaster,
						Network:     protocol.NetworkFarcaster,
						Source:      s.Name(),
						Transfers: []model.Transfer{
							{
								TransactionHash: hash,
								Timestamp:       cast.PublishedAt,
								Index:           protocol.IndexVirtual,
								AddressFrom:     strings.ToLower(user.Address.String()),
								AddressTo:       strings.ToLower(parentUser.Address.String()),
								Tag:             filter.TagSocial,
								Type:            filter.SocialComment,
								Metadata:        metadataPost,
								Network:         protocol.NetworkFarcaster,
								Platform:        protocol.PlatformFarcaster,
								Source:          s.Name(),
								RelatedUrls:     []string{fmt.Sprintf("https://www.discove.xyz/casts/%s", cast.ParentHash)},
							},
						},
					}
					mu.Lock()
					transactions = append(transactions, &transaction)
					mu.Unlock()
				}
			}
		}(cast)
	}
	wg.Wait()

	if len(casts) == 0 {
		return transactions, nil
	}
	// set cache
	last, _ := lo.Last(casts)

	cursor = last.Hash.String()
	cache.Global().Set(ctx, farcasterCacheKey, cursor, 7*24*time.Hour)

	return transactions, nil
}
