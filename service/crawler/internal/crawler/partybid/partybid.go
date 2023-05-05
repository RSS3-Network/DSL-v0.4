package partybid

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
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
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type service struct {
	config       *config.Config
	kuroraClient *kurora.Client
	tokenClient  *token.Client
}

var (
	_ crawler.Crawler = (*service)(nil)

	partybidCacheKey = "crawler:partybid"
)

func New(config *config.Config) crawler.Crawler {
	return &service{
		config:      config,
		tokenClient: token.New(),
	}
}

func (s *service) Name() string {
	return protocol.PlatformPartyBid
}

func (s *service) Network() string {
	return protocol.NetworkEthereum
}

func (s *service) Run() error {
	loggerx.Global().Info("partybid: run")

	ctx := context.Background()
	var err error

	s.kuroraClient, err = kurora.Dial(ctx, s.config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error("partyBid: kurora.Dial error", zap.Error(err), zap.String("endpoint", s.config.Kurora.Endpoint))

		return err
	}

	for {
		transactions, cacheInfo, err := s.handlePartyBidEvents(ctx)
		if err != nil {
			loggerx.Global().Error("partyBid: handlePartyBidEvents error", zap.Error(err))

			return err
		}

		if len(transactions) == 0 {
			time.Sleep(1 * time.Minute)

			continue
		}

		// insert db
		err = database.UpsertTransactions(ctx, transactions, false)
		if err != nil {
			continue
		}

		// set cache
		cache.Global().Set(ctx, partybidCacheKey, cacheInfo, 7*24*time.Hour)
	}
}

type nftMetadata struct {
	Address  common.Address
	TokenId  *big.Int
	Metadata *metadata.Token
}

func (s *service) handlePartyBidEvents(ctx context.Context) ([]*model.Transaction, string, error) {
	tracer := otel.Tracer("partybid")
	_, trace := tracer.Start(ctx, "partyBid:handlePartyBidEvents")
	var (
		transactions []*model.Transaction
		err          error
	)

	transactionMap := make(map[string]*model.Transaction, 0)

	defer func() { opentelemetry.Log(trace, nil, transactions, err) }()

	cursor, _ := cache.Global().Get(ctx, partybidCacheKey).Result()

	query := kurora.DatasetPartyBidDAOEventsQuery{
		Limit: lo.ToPtr(100),
	}

	if len(cursor) > 0 {
		query.Cursor = &cursor
	}

	events, err := s.kuroraClient.FetchDatasetPartyBidEvents(ctx, query)
	if err != nil {
		loggerx.Global().Error("FetchDatasetPartyBidEvents error", zap.Error(err), zap.Any("query", query))

		return nil, "", err
	}

	loggerx.Global().Info("partybid: kuroraClient FetchDatasetPartyBidEvents result", zap.Int("len", len(events)), zap.String("cursor", cursor))

	var (
		metaDataMap = make(map[string]*nftMetadata, 0)
		nftList     = make([]*nftMetadata, 0)
		locker      sync.Mutex
	)

	for _, event := range events {
		key := fmt.Sprintf("%v-%v", event.NftAddress.String(), event.NftId.BigInt())
		if _, ok := metaDataMap[key]; ok || (event.Type == Liquidity && event.EventType != LoadCreate) {
			continue
		}

		nft := &nftMetadata{event.NftAddress, event.NftId.BigInt(), nil}

		metaDataMap[key] = nft

		nftList = append(nftList, nft)
	}

	opt := lop.NewOption().WithConcurrency(10)
	lop.ForEach(nftList, func(nftData *nftMetadata, i int) {
		address := nftData.Address.String()
		tokenId := nftData.TokenId

		metaData, err := s.tokenClient.NFTToMetadata(ctx, s.Network(), address, tokenId)
		if err != nil {
			zap.L().Error("nft to metadata", zap.Error(err), zap.String("contract_address", address), zap.Stringer("token_id", tokenId))

			return
		}

		locker.Lock()
		metaDataMap[fmt.Sprintf("%v-%v", address, tokenId)].Metadata = metaData
		locker.Unlock()
	}, opt)

	for _, event := range events {
		key := fmt.Sprintf("%v-%v", event.NftAddress.String(), event.NftId.BigInt())

		var (
			feedTag   string
			feedType  string
			addressTo string
		)

		transfers := make([]model.Transfer, 0)

		switch event.Type {
		}

		transaction := &model.Transaction{
			BlockNumber: int64(event.BlockNumber),
			Timestamp:   event.Timestamp,
			Owner:       strings.ToLower(event.From.String()),
			Hash:        strings.ToLower(event.TransactionHash.String()),
			AddressFrom: strings.ToLower(event.From.String()),
			AddressTo:   addressTo,
			Network:     s.Network(),
			Success:     lo.ToPtr(true),
			Platform:    s.Name(),
			Source:      protocol.SourceKurora,
			Tag:         feedTag,
			Type:        feedType,
			Transfers:   transfers,
		}

		if _, ok := transactionMap[strings.ToLower(event.TransactionHash.String())]; !ok {
			transactionMap[strings.ToLower(event.TransactionHash.String())] = transaction
		} else {
			transactionMap[strings.ToLower(event.TransactionHash.String())].Transfers = append(transactionMap[strings.ToLower(event.TransactionHash.String())].Transfers, transfers...)
		}
	}

	last, err := lo.Last(events)
	if err == nil {
		cursor = kurora.LogCursor(last.TransactionHash, last.LogIndex)
	}

	for _, tx := range transactionMap {
		transactions = append(transactions, tx)
	}

	return transactions, cursor, nil
}

func (s *service) buildTransfers(transfers []model.Transfer, hash, from, to, feedTag, feedType, nftAddress string, nftId decimal.Decimal, timestamp time.Time, index int64, metadata json.RawMessage, partyBidUrl bool) []model.Transfer {
	transfer := model.Transfer{
		TransactionHash: hash,
		Timestamp:       timestamp,
		Index:           index,
		AddressFrom:     from,
		AddressTo:       to,
		Metadata:        metadata,
		Network:         s.Network(),
		Platform:        s.Name(),
		Source:          protocol.SourceKurora,
		Tag:             feedTag,
		Type:            feedType,
		RelatedUrls: ethereum.BuildURL(
			[]string{utils.GetTxHashURL(s.Network(), hash)},
		),
	}

	if nftAddress != "" {
		transfer.RelatedUrls = append(transfer.RelatedUrls, ethereum.BuildTokenURL(s.Network(), nftAddress, nftId.String())...)
	}

	// if partyBidUrl {
	// 	transfer.RelatedUrls = append(transfer.RelatedUrls, fmt.Sprintf("%s/%s/%v/", "https://www.partybid.app/party/", strings.ToLower(nftAddress), nftId))
	// }

	transfers = append(transfers, transfer)

	return transfers
}
