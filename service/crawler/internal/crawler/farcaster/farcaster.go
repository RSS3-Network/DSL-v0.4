package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/farcaster"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	lop "github.com/samber/lo/parallel"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ crawler.Crawler = (*service)(nil)

type service struct {
	farClient *farcaster.Client
}

func New() crawler.Crawler {
	return &service{
		farClient: farcaster.NewClient(),
	}
}

func (s *service) Name() string {
	return protocol.PlatformFarcaster
}

func (s *service) Run() error {
	ctx := context.Background()
	// init cache cast number 0
	loggerx.Global().Info("farcaster_crawler start to init Map Cache")
	farcasterCacheMap, err := s.farClient.GetFarcasterCacheMap()
	if err != nil && err != redis.Nil {
		loggerx.Global().Error("farcaster_crawler fail to init Map Cache", zap.Error(err))
		return err
	}

	// get address cast number from db
	loggerx.Global().Info("farcaster_crawler start to get db data")
	s.GetAddressCastNumFromDb(farcasterCacheMap)

	var wg sync.WaitGroup
	ch := make(chan struct{}, 5)
	for {
		loggerx.Global().Info("farcaster_crawler start a new round ", zap.Int("cache user is:", len(farcasterCacheMap)))

		for address, data := range farcasterCacheMap {
			wg.Add(1)
			go func(faAddress string, cacheData *farcaster.CacheAddress) {
				defer func() {
					<-ch
					wg.Done()
				}()
				ch <- struct{}{}
				castList, err := s.farClient.GetActivityList(ctx, faAddress)
				if err != nil {
					loggerx.Global().Warn("farcaster_crawler get cast error, ", zap.Error(err))
					return
				}

				if len(castList) == int(cacheData.CastNumber) {
					return
				}

				transactions, err := s.HandleTransactions(ctx, castList, cacheData.EvmAddress, cacheData.CastNumber)
				if err != nil {
					loggerx.Global().Warn("farcaster_crawler handle transactions error, ", zap.Error(err))
					return
				}

				internalTransactions, err := s.HandleTransfer(ctx, transactions)
				if err != nil {
					loggerx.Global().Warn("farcaster_crawler handle transfer error, ", zap.Error(err))
					return
				}

				err = database.UpsertTransactions(ctx, internalTransactions)
				if err != nil {
					loggerx.Global().Warn("farcaster_crawler upsertTransactions error, ", zap.Error(err))
					return
				}

				farcaster.ReplaceGlobal(faAddress, &farcaster.CacheAddress{
					EvmAddress: cacheData.EvmAddress,
					CastNumber: int64(len(castList)),
				})
			}(address, data)
		}
		wg.Wait()

		loggerx.Global().Info("farcaster_crawler end a new round ", zap.Int("cache user is:", len(farcasterCacheMap)))

		// add new users to cache map
		s.farClient.UpdateFarcasterCacheMap()
		loggerx.Global().Info("farcaster_crawler finish update cache map ", zap.Int("cache user is:", len(farcaster.FarcasterCacheMap)))

		// get cache map
		res, err := s.farClient.GetFarcasterCacheMap()
		if err != nil {
			loggerx.Global().Error("farcaster_crawler fail to get Map Cache", zap.Error(err))
		}

		if len(res) > 0 {
			err = s.farClient.SetCurrentMap(ctx, res)
			if err != nil {
				loggerx.Global().Error("farcaster_crawler fail to set Map Cache", zap.Error(err))
			}
		} else {
			res = farcaster.FarcasterCacheMap
		}

		farcasterCacheMap = res
	}
}

func (s *service) GetAddressCastNumFromDb(farcasterCacheMap map[string]*farcaster.CacheAddress) {
	var wg sync.WaitGroup

	ch := make(chan struct{}, 10)

	for address, cache := range farcasterCacheMap {
		wg.Add(1)
		go func(evmAddress string, faAddress string) {
			defer func() {
				<-ch
				wg.Done()
			}()
			ch <- struct{}{}
			total := int64(0)
			if err := database.Global().
				Model(&model.Transfer{}).
				Where("address_from = ?", evmAddress).
				Where("network = ?", protocol.NetworkFarcaster).
				Count(&total).Error; err != nil {
				loggerx.Global().Warn("farcaster_crawler find transactions error, ", zap.Error(err))
			}

			farcaster.ReplaceGlobal(faAddress, &farcaster.CacheAddress{
				EvmAddress: evmAddress,
				CastNumber: total,
			})
		}(cache.EvmAddress, address)
	}
	wg.Wait()
}

func (s *service) HandleTransactions(ctx context.Context, castList []farcaster.Cast, evmAddress string, castNum int64) ([]model.Transaction, error) {
	tracer := otel.Tracer("crawler_farcaster")
	_, trace := tracer.Start(ctx, "crawler_farcaster:HandleTransactions")
	defer trace.End()

	transactions := make([]model.Transaction, 0)

	for i := 0; i < len(castList)-int(castNum); i++ {
		cast := castList[i]
		sourceData, err := json.Marshal(cast)
		if err != nil {
			loggerx.Global().Warn("farcaster_crawler marshal error", zap.Error(err))
			return transactions, err
		}
		timestamp := time.UnixMilli(cast.Body.PublishedAt)

		var addressTo string
		if cast.Meta.ReplyParentUsername.Address != "" {
			addressTo = cast.Meta.ReplyParentUsername.Address
		} else {
			addressTo = evmAddress
		}

		addressTo = strings.ToLower(addressTo)

		hash := fmt.Sprintf("%s%d", cast.MerkleRoot, cast.Body.Sequence)

		transactions = append(transactions, model.Transaction{
			// use timestamp as block number, as there is no actual blocknumber on farcaster
			BlockNumber: cast.Body.PublishedAt,
			Timestamp:   timestamp,
			// use MerkleRoot as hash, as there is no actual hash on farcaster
			Hash:        hash,
			AddressFrom: evmAddress,
			AddressTo:   addressTo,
			// TODO: identify the correct platform
			Platform:   protocol.PlatformFarcaster,
			Network:    protocol.NetworkFarcaster,
			Source:     s.Name(),
			SourceData: sourceData,
			Transfers: []model.Transfer{
				{
					// This is a virtual transfer
					TransactionHash: hash,
					Timestamp:       timestamp,
					Index:           protocol.IndexVirtual,
					AddressFrom:     evmAddress,
					AddressTo:       strings.ToLower(cast.Body.Address),
					Metadata:        metadata.Default,
					Network:         protocol.NetworkFarcaster,
					Platform:        protocol.PlatformFarcaster,
					SourceData:      sourceData,
				},
			},
		})
	}

	return transactions, nil
}

func (s *service) HandleTransfer(ctx context.Context, transactions []model.Transaction) ([]*model.Transaction, error) {
	tracer := otel.Tracer("crawler_farcaster")

	_, snap := tracer.Start(ctx, "crawler_farcaster:HandleTransfer")

	defer snap.End()

	internalTransactions := make([]*model.Transaction, 0)

	var mu sync.Mutex

	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		// Retain the action model of the transfer type
		transferMap := make(map[int64]model.Transfer)

		for _, transfer := range transaction.Transfers {
			if err := s.HandlePost(ctx, &transfer); err != nil {
				loggerx.Global().Warn("[farcaster]: failed to HandlePost", zap.Error(err), zap.String("network", protocol.NetworkFarcaster), zap.String("transaction_hash", transaction.Hash), zap.String("address", transfer.AddressFrom))

				continue
			}
			transfer.AddressTo = transaction.AddressTo
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)
		transaction.Transfers = append(transaction.Transfers, transferMap[protocol.IndexVirtual])

		for _, transfer := range transaction.Transfers {
			transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
		}

		transaction.Owner = transaction.AddressFrom
		mu.Lock()
		internalTransactions = append(internalTransactions, &transaction)
		mu.Unlock()
	})

	return internalTransactions, nil
}

func (s *service) HandlePost(ctx context.Context, transfer *model.Transfer) (err error) {
	tracer := otel.Tracer("crawler_farcaster")
	_, trace := tracer.Start(ctx, "crawler_farcaster:HandlePost")

	defer func() { opentelemetry.Log(trace, transfer.AddressFrom, transfer.TransactionHash, err) }()

	var cast farcaster.Cast
	err = json.Unmarshal(transfer.SourceData, &cast)
	if err != nil {
		loggerx.Global().Warn("unable to unmarshal cast", zap.Error(err))
		return err
	}

	post := &metadata.Post{
		CreatedAt:      time.UnixMilli(cast.Body.PublishedAt).Format(time.RFC3339),
		Author:         []string{cast.Meta.DisplayName, transfer.AddressTo},
		Body:           cast.Body.Data.Text,
		TypeOnPlatform: []string{"cast"},
	}
	transfer.Metadata, _ = json.Marshal(post)
	transfer.Tag = filter.TagSocial
	transfer.Type = filter.SocialPost
	// TODO: farcaster does not have an API for individual posts, so we can't get the post URL
	transfer.RelatedUrls = []string{fmt.Sprintf("https://www.discove.xyz/casts/%s", cast.MerkleRoot)}
	return nil
}
