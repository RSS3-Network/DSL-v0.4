package mirror

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/pregod_etl"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

var (
	_ crawler.Crawler = (*service)(nil)

	mirrorCacheKey = "crawler:mirror" // mirror -> height:cursor
	MirrorAddress  = "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c"
)

type service struct {
	etlClient *pregod_etl.Client
}

func New(config *config.Config) crawler.Crawler {
	var err error
	crawler := &service{}

	if crawler.etlClient, err = pregod_etl.NewClient(protocol.NetworkArweave, config.RPC.PregodETL.Kurora.HTTP); err != nil {
		logrus.Error("[lens] pregod_etl.NewClient error, ", err)
		return nil
	}

	return crawler
}

func (s *service) Name() string {
	return protocol.PlatformMirror
}

func (s *service) Run() error {
	ctx := context.Background()

	for {
		// get mirror transactions
		transactions, err := s.getMirrorTransactions(ctx)
		if err != nil {
			logrus.Error("[mirror] Run: getMirrorTransactions error, ", err)

			continue
		}

		if len(transactions) == 0 {
			time.Sleep(10 * time.Minute)

			continue
		}

		// deduplicate data
		transactions, err = database.DeduplicateTransactions(ctx, transactions)
		if err != nil || len(transactions) == 0 {
			continue
		}

		// insert db
		err = database.UpsertTransactions(ctx, transactions)
		if err != nil {
			continue
		}
	}
}

func (s *service) getMirrorTransactions(ctx context.Context) ([]*model.Transaction, error) {
	tracer := otel.Tracer("mirror")
	_, trace := tracer.Start(ctx, "mirror:GetMirrorTransactions")
	internalTransactions := []*model.Transaction{}
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

	parameter := pregod_etl.GetArweaveTransactionsRequest{
		Owner: MirrorAddress,
		Limit: 100,
		Order: "asc",
	}

	cacheInfo, _ := cache.Global().Get(ctx, mirrorCacheKey).Result()
	if cacheList := strings.Split(cacheInfo, ":"); len(cacheList) == 2 {
		parameter.BlockHeightFrom, _ = decimal.NewFromString(cacheList[0])
		parameter.Cursor = cacheList[1]
	}

	// get logs from pregod_etl
	result, err := s.etlClient.GetArweaveTransactions(ctx, parameter)
	if err != nil {
		logrus.Error("[mirror] getMirrorTransactions: etlClient GetArweaveTransactions error, ", err)

		return nil, err
	}

	for _, transaction := range result.Result {
		data := pregod_etl.ArweaveData{}
		if err := json.Unmarshal([]byte(transaction.Data), &data); err != nil {
			logrus.Errorf("[mirror] getMirrorTransactions: json unmarshal error, %v", err)
			return nil, err
		}

		metadata, err := json.Marshal(&metadata.Post{
			Title: data.Content.Title,
			Body:  data.Content.Body,
		})
		if err != nil {
			logrus.Errorf("[mirror] getMirrorTransactions: json unmarshal error, %v", err)
			continue
		}

		source, err := json.Marshal(&transaction)
		if err != nil {
			logrus.Errorf("[mirror] getMirrorTransactions: json unmarshal error, %v", err)
			continue
		}

		var contentDigest, originDigest string
		for _, tag := range transaction.Tags {
			switch tag.Name {
			case "Content-Digest":
				contentDigest = tag.Value
			case "Original-Content-Digest":
				originDigest = tag.Value
			}
		}
		if len(originDigest) == 0 {
			originDigest = contentDigest
		}

		address := strings.ToLower(data.Authorship.Contributor)
		internalTransaction := &model.Transaction{
			BlockNumber: transaction.Height.BigInt().Int64(),
			Hash:        transaction.ID,
			Owner:       address,
			AddressFrom: address,
			AddressTo:   strings.ToLower(MirrorAddress),
			Platform:    protocol.PlatformMirror,
			Network:     protocol.NetworkArweave,
			Source:      protocol.SourcePregodETL,
			Timestamp:   time.Unix(data.Content.Timestamp.BigInt().Int64(), 0),
			Tag:         filter.TagSocial,
			Type:        filter.SocialPost,
			Transfers: []model.Transfer{
				// This is a virtual transfer
				{
					TransactionHash: transaction.ID,
					Timestamp:       time.Unix(data.Content.Timestamp.BigInt().Int64(), 0),
					Index:           protocol.IndexVirtual,
					AddressFrom:     address,
					AddressTo:       strings.ToLower(MirrorAddress),
					Metadata:        metadata,
					Network:         protocol.NetworkArweave,
					Source:          protocol.SourcePregodETL,
					SourceData:      source,
					Platform:        protocol.PlatformMirror,
					RelatedUrls:     []string{fmt.Sprintf("https://mirror.xyz/%s/%s", address, originDigest)},
					Tag:             filter.TagSocial,
					Type:            filter.SocialPost,
				},
			},
		}

		internalTransactions = append(internalTransactions, internalTransaction)
	}

	// set cache
	cacheInfo = fmt.Sprintf("%v:%v", result.Result[len(result.Result)-1].Height, result.Cursor)
	cache.Global().Set(ctx, mirrorCacheKey, cacheInfo, 7*24*time.Hour)

	return internalTransactions, nil
}
