package iqwiki

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/iqwiki"
	iqwiki_contract "github.com/naturalselectionlabs/pregod/common/worker/iqwiki/contract"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/iqwiki/graphql"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ crawler.Crawler = (*service)(nil)

type service struct {
	iqClient *iqwiki.Client
}

func New() crawler.Crawler {
	return &service{
		iqClient: iqwiki.NewClient(),
	}
}

func (s *service) Name() string {
	return protocol.PlatformIQWiki
}

func (s *service) Run() error {
	ctx := context.Background()
	// init cache cast number 0
	loggerx.Global().Info("iqwiki_crawler start to init Map Cache")
	iqwikiCacheMap, err := s.iqClient.GetIqwikiCacheMap()
	if err != nil && err != redis.Nil {
		loggerx.Global().Error("iqwiki_crawler fail to init Map Cache", zap.Error(err))
		return err
	}

	// get address cast number from db
	loggerx.Global().Info("iqwiki_crawler start to get db data")
	s.GetAddressCastNumFromDb(iqwikiCacheMap)

	for {
		loggerx.Global().Info("iqwiki_crawler start a new round ", zap.Int("cache user is:", len(iqwikiCacheMap)))

		for address, num := range iqwikiCacheMap {
			activityList, err := s.iqClient.GetUserActivityList(ctx, address)
			if err != nil {
				loggerx.Global().Warn("iqwiki_crawler get cast error, ", zap.Error(err))
				continue
			}

			if len(activityList) == num {
				continue
			}

			transactions, err := s.HandleTransactions(ctx, activityList, address)
			if err != nil {
				loggerx.Global().Warn("iqwiki_crawler handle transactions error, ", zap.Error(err))
				continue
			}

			internalTransactions, err := s.HandleTransfer(ctx, transactions)
			if err != nil {
				loggerx.Global().Warn("iqwiki_crawler handle transfer error, ", zap.Error(err))
				continue
			}

			err = database.UpsertTransactions(ctx, internalTransactions)
			if err != nil {
				loggerx.Global().Warn("iqwiki_crawler upsertTransactions error, ", zap.Error(err))
				continue
			}

			iqwiki.ReplaceGlobal(address, len(transactions))
		}
		loggerx.Global().Info("iqwiki_crawler end a round ", zap.Int("cache user is:", len(iqwikiCacheMap)))

		s.iqClient.UpdateIqwikiCacheMap()

		res, err := s.iqClient.GetIqwikiCacheMap()
		if err != nil {
			loggerx.Global().Error("iqwiki_crawler fail to get Map Cache", zap.Error(err))
		}

		err = s.iqClient.SetCurrentMap(ctx, res)
		if err != nil {
			loggerx.Global().Error("iqwiki_crawler fail to set Map Cache", zap.Error(err))
		}

		iqwikiCacheMap = res
	}
}

func (s *service) GetAddressCastNumFromDb(iqwikiCacheMap map[string]int) {
	var wg sync.WaitGroup

	ch := make(chan struct{}, 10)

	for userAddress := range iqwikiCacheMap {
		wg.Add(1)
		go func(address string) {
			defer wg.Done()
			ch <- struct{}{}
			total := int64(0)
			if err := database.Global().
				Model(&model.Transaction{}).
				Where("owner = ?", address).
				Where("address_from", strings.ToLower(iqwiki_contract.AddressEveripediaSign.String())).
				Where("address_to", strings.ToLower(iqwiki_contract.AddressEveripedia.String())).
				Where("platform = ?", protocol.PlatformIQWiki).
				Count(&total).Error; err != nil {
				loggerx.Global().Warn("iqwiki_crawler find transfer error, ", zap.Error(err))
			}

			iqwiki.ReplaceGlobal(address, int(total))
			<-ch
		}(userAddress)
	}
	wg.Wait()
}

func (s *service) getEveTransactions(address string) int64 {
	var result struct {
		BlockNumber int64 `gorm:"column:block_number"`
	}

	if err := database.Global().
		Model(&model.Transaction{}).
		Select("COALESCE(block_number, 0) AS block_number").
		Where("owner = ?", address).
		Where("address_from", strings.ToLower(iqwiki_contract.AddressEveripediaSign.String())).
		Where("address_to", strings.ToLower(iqwiki_contract.AddressEveripedia.String())).
		Where("success IS TRUE").
		Order("block_number DESC").
		Limit(1).
		Find(&result).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0
	}

	return result.BlockNumber
}

func (s *service) HandleTransactions(ctx context.Context, activityList []graphqlx.GetUserActivitiesActivitiesByUserActivity, address string) ([]model.Transaction, error) {
	tracer := otel.Tracer("iqwiki_crawler")
	_, trace := tracer.Start(ctx, "iqwiki_crawler:HandleTransactions")
	defer trace.End()

	transactions := make([]model.Transaction, 0)

	blockNum := s.getEveTransactions(strings.ToLower(address))

	for _, activity := range activityList {

		if activity.Block <= int(blockNum) {
			continue
		}
		var action string

		if activity.Type == graphqlx.StatusCreated {
			action = filter.SocialProfileCreate
		} else if activity.Type == graphqlx.StatusUpdated {
			action = filter.SocialProfileUpdate
		}

		for _, content := range activity.Content {
			sourceData, err := json.Marshal(content)
			if err != nil {
				continue
			}
			transactions = append(transactions, model.Transaction{
				BlockNumber: int64(activity.Block),
				Timestamp:   activity.Datetime,
				Hash:        content.TransactionHash,
				AddressFrom: strings.ToLower(iqwiki_contract.AddressEveripediaSign.String()),
				AddressTo:   strings.ToLower(iqwiki_contract.AddressEveripedia.String()),
				Owner:       strings.ToLower(activity.User.Id),
				Platform:    protocol.PlatformIQWiki,
				Network:     protocol.NetworkPolygon,
				Source:      s.Name(),
				SourceData:  sourceData,
				Tag:         filter.TagSocial,
				Type:        filter.SocialWiki,
				Transfers: []model.Transfer{
					// This is a virtual transfer
					{
						TransactionHash: content.TransactionHash,
						Timestamp:       activity.Datetime,
						Index:           protocol.IndexVirtual,
						AddressFrom:     strings.ToLower(iqwiki_contract.AddressEveripediaSign.String()),
						AddressTo:       strings.ToLower(iqwiki_contract.AddressEveripedia.String()),
						Metadata:        metadata.Default,
						Network:         protocol.NetworkPolygon,
						SourceData:      sourceData,
						Tag:             filter.TagSocial,
						Type:            action,
					},
				},
			})
		}

	}

	return transactions, nil
}

func (s *service) HandleTransfer(ctx context.Context, transactions []model.Transaction) ([]*model.Transaction, error) {
	tracer := otel.Tracer("crawler_iqwiki")

	_, snap := tracer.Start(ctx, "crawler_iqwiki:HandleTransfer")

	defer snap.End()

	internalTransactions := make([]*model.Transaction, 0)

	polygonClient, err := ethclientx.Global(protocol.NetworkPolygon)
	if err != nil {
		return nil, err
	}

	for _, transaction := range transactions {
		internalTransaction := transaction
		internalTransaction.Transfers = make([]model.Transfer, 0)

		addressTo := common.HexToAddress(transaction.AddressTo)

		switch addressTo {
		case iqwiki_contract.AddressEveripedia:
			receipt, err := polygonClient.TransactionReceipt(context.Background(), common.HexToHash(transaction.Hash))
			if err != nil {
				loggerx.Global().Warn("worker_iqwiki: failed to Get Receipt", zap.Error(err), zap.String("network", protocol.NetworkPolygon), zap.String("transaction_hash", transaction.Hash), zap.String("address", transaction.Owner))
				continue
			}

			for _, transfer := range transaction.Transfers {
				for _, log := range receipt.Logs {
					switch log.Topics[0] {
					case iqwiki_contract.EventPosted:
						transfer.Index = int64(log.Index)
						transfer.SourceData = transaction.SourceData
						if err := s.HandlePost(ctx, &transfer); err != nil {
							loggerx.Global().Warn("worker_iqwiki: failed to HandleTransfer", zap.Error(err), zap.String("network", protocol.NetworkPolygon), zap.String("transaction_hash", transaction.Hash), zap.String("address", transaction.Owner))

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

			internalTransactions = append(internalTransactions, &internalTransaction)
		default:
			continue
		}
	}

	return internalTransactions, nil
}

type Wiki struct {
	Content string `json:"content"`
}

func (s *service) HandlePost(ctx context.Context, transfer *model.Transfer) (err error) {
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
