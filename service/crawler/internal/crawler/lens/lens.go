package lens

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/naturalselectionlabs/kurora/constant"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"

	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/common"

	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"

	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	worker "github.com/naturalselectionlabs/pregod/common/worker/lens"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

var (
	_ crawler.Crawler = (*service)(nil)

	lensLogsCacheKey   = "crawler:lens:%v"
	lensMomokaCacheKey = "crawler:lens:momoka"
)

type service struct {
	config       *config.Config
	kuroraClient *kurora.Client
	lensWorker   *worker.Client
}

func New(config *config.Config) (crawler.Crawler, error) {
	kuroraClient, err := kurora.Dial(context.Background(), config.Kurora.Endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		return nil, err
	}

	lensWorker, err := worker.New(kuroraClient)
	if err != nil {
		return nil, err
	}

	lensCrawler := &service{
		config:       config,
		lensWorker:   lensWorker,
		kuroraClient: kuroraClient,
	}

	return lensCrawler, nil
}

func (s *service) Name() string {
	return protocol.PlatformLens
}

func (s *service) Run() error {
	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(1)

	go s.HandlePolygonLensEvents(ctx)

	go s.HandleMomokaTransactions(ctx)

	wg.Wait()

	return nil
}

func (s *service) HandlePolygonLensEvents(ctx context.Context) {
	var wg sync.WaitGroup
	for eventHash, contractAddress := range lens.SupportLensEvents {
		wg.Add(1)
		go func(eventHash common.Hash, contractAddress common.Address) {
			defer wg.Done()

			for {
				// get lens logs
				transactions, cacheInfo, err := s.getPolygonLogs(ctx, eventHash, contractAddress)
				if err != nil {
					loggerx.Global().Error("failed to query lens", zap.Error(err))

					time.Sleep(1 * time.Minute)
					continue
				}

				if len(transactions) == 0 {
					time.Sleep(1 * time.Minute)

					continue
				}

				// build transaction
				message := &protocol.Message{
					Network: protocol.NetworkPolygon,
				}

				if transactions, err = ethereum.BuildTransactions(ctx, message, transactions); err != nil {
					loggerx.Global().Error("failed to build transactions, ", zap.Error(err))

					continue
				}

				// lens worker
				results := s.handlePolygonReceipts(ctx, transactions)

				// insert db
				err = database.UpsertTransactions(ctx, results, false)
				if err != nil {
					continue
				}

				// set cache
				cacheKey := fmt.Sprintf(lensLogsCacheKey, eventHash.String())
				cache.Global().Set(ctx, cacheKey, cacheInfo, 7*24*time.Hour)
			}
		}(eventHash, contractAddress)
	}
	wg.Wait()
}

func (s *service) getPolygonLogs(ctx context.Context, eventHash common.Hash, contractAddress common.Address) ([]*model.Transaction, string, error) {
	tracer := otel.Tracer("lens")
	_, trace := tracer.Start(ctx, "len:GetLensLogs")
	var internalTransactions []*model.Transaction
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

	query := kurora.DatasetLensEventQuery{
		Address:    &contractAddress,
		TopicFirst: &eventHash,
		Limit:      lo.ToPtr(100),
	}

	cacheKey := fmt.Sprintf(lensLogsCacheKey, eventHash.String())
	cacheInfo, _ := cache.Global().Get(ctx, cacheKey).Result()
	if cacheList := strings.Split(cacheInfo, ":"); len(cacheList) == 2 {
		blockNumber, _ := decimal.NewFromString(cacheList[0])
		query.BlockNumberFrom = &blockNumber
		query.Cursor = &cacheList[1]
	}

	loggerx.Global().Info("FetchDatasetLensEvents request", zap.Any("query", query))

	// get logs from kurora
	result, err := s.kuroraClient.FetchDatasetLensEvents(ctx, query)
	if err != nil {
		loggerx.Global().Error("FetchDatasetLensEvents error", zap.Error(err), zap.Any("query", query))
		return nil, "", err
	}

	transactionMap := make(map[string]*model.Transaction)
	for _, transfer := range result {
		transaction := &model.Transaction{
			BlockNumber: transfer.BlockNumber.BigInt().Int64(),
			Hash:        transfer.TransactionHash.String(),
			Index:       int64(transfer.TransactionIndex),
			Network:     protocol.NetworkPolygon,
			Platform:    protocol.PlatformLens,
			Transfers:   make([]model.Transfer, 0),
			Source:      protocol.SourceKurora,
		}

		transactionMap[transaction.Hash] = transaction
	}

	for _, transaction := range transactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	last, err := lo.Last(result)
	if err != nil {
		return nil, "", nil
	}

	cursor := kurora.LogCursor(last.TransactionHash, last.Index)
	cacheInfo = fmt.Sprintf("%v:%v", last.BlockNumber, cursor)

	return internalTransactions, cacheInfo, nil
}

func (s *service) handlePolygonReceipts(ctx context.Context, transactions []*model.Transaction) []*model.Transaction {
	var mu sync.Mutex
	var internalTransactions []*model.Transaction
	opt := lop.NewOption().WithConcurrency(10)
	lop.ForEach(transactions, func(transaction *model.Transaction, i int) {
		transferMap := make(map[int64]model.Transfer)
		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)

		// get receipt
		internalTransfers, err := s.lensWorker.HandleReceipt(ctx, transaction)
		if err != nil {
			logrus.Errorf("[lens worker] handleReceipt error, %v", err)

			return
		}

		transaction.Transfers = s.lensWorker.FilterLensTransfer(transaction.Owner, internalTransfers)

		for _, transfer := range transaction.Transfers {
			if transaction.Tag == "" {
				transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
			}
		}

		mu.Lock()
		internalTransactions = append(internalTransactions, transaction)
		mu.Unlock()
	}, opt)

	return internalTransactions
}

func (s *service) HandleMomokaTransactions(ctx context.Context) {
	for {
		// get arweave lens transactions
		momokaTransactions, cacheInfo, err := s.getMomokaTransactions(ctx)
		if err != nil {
			loggerx.Global().Error("failed to query momoka transaction", zap.Error(err))

			time.Sleep(1 * time.Minute)
			continue
		}

		// handle momoka transactions
		transactions := s.handleMomokaTransactions(ctx, momokaTransactions)

		// insert db
		err = database.UpsertTransactions(ctx, transactions, false)
		if err != nil {
			loggerx.Global().Error("failed to insert momoka transaction", zap.Error(err))

			continue
		}

		// set cache
		cache.Global().Set(ctx, lensMomokaCacheKey, cacheInfo, 7*24*time.Hour)
	}
}

func (s *service) getMomokaTransactions(ctx context.Context) ([]*kurora.DatasetMomokaTransaction, string, error) {
	// current block height
	cacheInfo, _ := cache.Global().Get(ctx, lensMomokaCacheKey).Result()
	blockHeight, _ := decimal.NewFromString(cacheInfo)
	if blockHeight.IsZero() {
		blockHeight = decimal.NewFromInt(1166650)
	}

	// query momoka latest block height
	latestBlockHeight, err := s.getMomokaLatestTransaction(ctx)
	if err != nil {
		return nil, "", err
	}

	if blockHeight.Cmp(*latestBlockHeight) == 0 {
		time.Sleep(1 * time.Minute)

		return nil, blockHeight.String(), nil
	}

	// query momoka transactions
	query := kurora.DatasetMomokaTransactionQuery{
		BlockHeightFrom: lo.ToPtr(blockHeight),
		BlockHeightTo:   lo.ToPtr(blockHeight.Add(decimal.NewFromInt(20))),
		Order:           lo.ToPtr("asc"),
	}

	if query.BlockHeightTo.GreaterThan(*latestBlockHeight) {
		query.BlockHeightTo = latestBlockHeight
	}

	transactions, err := s.kuroraClient.FetchDatasetMomokaTransactions(ctx, query)

	loggerx.Global().Info("FetchDatasetMomokaTransactions request", zap.Any("query", query), zap.Any("result", len(transactions)))

	if err != nil {
		loggerx.Global().Error("FetchEthereumTransactions error", zap.Error(err), zap.Any("query", query))
		return nil, "", err
	}

	return transactions, query.BlockHeightTo.String(), nil
}

func (s *service) getMomokaLatestTransaction(ctx context.Context) (*decimal.Decimal, error) {
	query := kurora.DatasetMomokaTransactionQuery{
		Order: lo.ToPtr("desc"),
		Limit: lo.ToPtr(1),
	}

	transaction, err := s.kuroraClient.FetchDatasetMomokaTransactions(ctx, query)
	if err != nil {
		loggerx.Global().Error("get latest transaction error", zap.Error(err), zap.Any("query", query))
		return nil, err
	}

	if len(transaction) == 0 {
		return nil, fmt.Errorf("no transaction found")
	}

	return &transaction[0].BlockHeight, nil
}

func (s *service) handleMomokaTransactions(ctx context.Context, momokaTransactions []*kurora.DatasetMomokaTransaction) []*model.Transaction {
	transactions := make([]*model.Transaction, 0)

	for _, momoka := range momokaTransactions {
		if momoka.ProfileID.IsZero() || momoka.PublicationID.IsZero() {
			continue
		}

		var transaction *model.Transaction
		var err error

		switch momoka.Type {
		case "POST_CREATED":
			transaction, err = s.handleMomokaPostCreated(ctx, momoka)
		case "COMMENT_CREATED":
			transaction, err = s.handleMomokaCommentCreated(ctx, momoka)
		case "MIRROR_CREATED":
			transaction, err = s.handleMomokaMirrorCreated(ctx, momoka)
		default:
			continue
		}

		if err != nil {
			loggerx.Global().Error("handle momoka transaction error", zap.Error(err), zap.Any("transaction", momoka))
			continue
		}

		transactions = append(transactions, transaction)
	}

	return transactions
}

func (s *service) handleMomokaPostCreated(ctx context.Context, momoka *kurora.DatasetMomokaTransaction) (*model.Transaction, error) {
	// parse post event
	var data DatasetMomokaTransactionPostData
	if err := json.Unmarshal(momoka.Data, &data); err != nil {
		return nil, err
	}

	// parse post content
	var content worker.LensContent
	if err := json.Unmarshal(momoka.Content, &content); err != nil {
		return nil, err
	}

	// get profile
	profile, err := s.lensWorker.GetProfile(momoka.ProfileID.BigInt())
	if err != nil {
		return nil, err
	}

	post := s.lensWorker.CreatePost(ctx, &content, profile.Handle)
	post.ProfileID = momoka.ProfileID.BigInt()
	post.PublicationID = momoka.PublicationID.BigInt()

	// build transaction
	return s.buildMomokaTransaction(momoka, profile.Address, filter.SocialPost, post, content.AppId, data.PublicationId), nil
}

func (s *service) handleMomokaCommentCreated(ctx context.Context, momoka *kurora.DatasetMomokaTransaction) (*model.Transaction, error) {
	// parse comment event
	var data DatasetMomokaTransactionCommentData
	if err := json.Unmarshal(momoka.Data, &data); err != nil {
		return nil, err
	}

	// parse comment content
	var content worker.LensContent
	if err := json.Unmarshal(momoka.Content, &content); err != nil {
		return nil, err
	}

	// get profile
	profile, err := s.lensWorker.GetProfile(momoka.ProfileID.BigInt())
	if err != nil {
		return nil, err
	}

	post := s.lensWorker.CreatePost(ctx, &content, profile.Handle)
	post.ProfileID = momoka.ProfileID.BigInt()
	post.PublicationID = momoka.PublicationID.BigInt()

	// get original post
	post.Target, err = s.getPublicationMetadata(ctx, data.Event.ProfileIdPointed, data.Event.PubIdPointed, data.ChainProofs.Pointer.Location)
	if err != nil {
		return nil, err
	}

	// build transaction
	return s.buildMomokaTransaction(momoka, profile.Address, filter.SocialComment, post, content.AppId, data.PublicationId), nil
}

func (s *service) handleMomokaMirrorCreated(ctx context.Context, momoka *kurora.DatasetMomokaTransaction) (*model.Transaction, error) {
	// parse mirror event
	var data DatasetMomokaTransactionMirrorData
	if err := json.Unmarshal(momoka.Data, &data); err != nil {
		return nil, err
	}

	// parse mirror content
	var content worker.LensContent
	if err := json.Unmarshal(momoka.Content, &content); err != nil {
		return nil, err
	}

	// get profile
	profile, err := s.lensWorker.GetProfile(momoka.ProfileID.BigInt())
	if err != nil {
		return nil, err
	}

	post := s.lensWorker.CreatePost(ctx, &content, profile.Handle)
	post.ProfileID = momoka.ProfileID.BigInt()
	post.PublicationID = momoka.PublicationID.BigInt()

	// get original post
	post.Target, err = s.getPublicationMetadata(ctx, data.Event.ProfileIdPointed, data.Event.PubIdPointed, data.ChainProofs.Pointer.Location)
	if err != nil {
		return nil, err
	}

	// build transaction
	return s.buildMomokaTransaction(momoka, profile.Address, filter.SocialShare, post, content.AppId, data.PublicationId), nil
}

func (s *service) getPublicationMetadata(ctx context.Context, profileIDStr string, pubIDStr string, pointer string) (*metadata.Post, error) {
	profileID, err := hexutil.DecodeBig(strings.Replace(profileIDStr, "0x0", "0x", 1))
	if err != nil {
		return nil, fmt.Errorf("decode profile id: %w", err)
	}

	pubID, err := hexutil.DecodeBig(strings.Replace(pubIDStr, "0x0", "0x", 1))
	if err != nil {
		return nil, fmt.Errorf("decode publication id: %w", err)
	}

	transactionID := strings.TrimPrefix(pointer, "ar://")

	// query publication content
	content, url, err := s.getPublicationContent(ctx, profileID, pubID, transactionID)
	if err != nil {
		loggerx.Global().Error("get publication content error", zap.Error(err), zap.Any("profileID", profileID), zap.Any("pubID", pubID), zap.String("transactionID", transactionID))

		return nil, err
	}

	// get profile
	profile, err := s.lensWorker.GetProfile(profileID)
	if err != nil {
		return nil, err
	}

	post := s.lensWorker.CreatePost(ctx, content, profile.Handle)
	post.ProfileID = profileID
	post.PublicationID = pubID
	post.TargetURL = url

	if !content.CreatedOn.IsZero() {
		post.CreatedAt = content.CreatedOn.Format(time.RFC3339)
	}

	return post, nil
}

func (s *service) getPublicationContent(ctx context.Context, profileID *big.Int, pubID *big.Int, transactionID string) (*worker.LensContent, string, error) {
	var content worker.LensContent

	// query momoka
	queryMomoka := kurora.DatasetMomokaTransactionQuery{
		TransactionID: &transactionID,
	}

	momokaTransactions, err := s.kuroraClient.FetchDatasetMomokaTransactions(ctx, queryMomoka)
	if err != nil {
		loggerx.Global().Error("get momoka transaction error", zap.Error(err), zap.Any("query", queryMomoka))

		return nil, "", err
	}

	if len(momokaTransactions) > 0 && len(momokaTransactions[0].Data) > 0 && len(momokaTransactions[0].Content) > 0 {
		var data DatasetMomokaTransactionPostData
		if err = json.Unmarshal(momokaTransactions[0].Data, &data); err != nil {
			loggerx.Global().Error("unmarshal momoka transaction data error", zap.Error(err), zap.Any("transaction", momokaTransactions))
			return nil, "", err
		}

		if err = json.Unmarshal(momokaTransactions[0].Content, &content); err != nil {
			loggerx.Global().Error("unmarshal momoka transaction content error", zap.Error(err), zap.Any("transaction", momokaTransactions))
			return nil, "", err
		}

		return &content, s.buildMomokaTransactionURL(data.PublicationId), nil
	}

	// query lens metadata
	queryMetadata := kurora.DatasetLensPublicationQuery{
		ProfileID:     lo.ToPtr(decimal.NewFromBigInt(profileID, 0)),
		PublicationID: lo.ToPtr(decimal.NewFromBigInt(pubID, 0)),
		Limit:         lo.ToPtr(1),
	}

	publication, err := s.kuroraClient.FetchDatasetLensPublications(ctx, queryMetadata)
	if err != nil {
		loggerx.Global().Error("get lens publication error", zap.Error(err), zap.Any("query", queryMetadata))

		return nil, "", err
	}

	if len(publication) > 0 {
		if err = json.Unmarshal(publication[0].Content, &content); err == nil {
			return &content, s.lensWorker.GetLensRelatedURL(ctx, profileID, pubID), nil
		}
	}

	return nil, "", fmt.Errorf("no content found")
}

func (s *service) buildMomokaTransaction(momoka *kurora.DatasetMomokaTransaction, owner string, socialType string, post *metadata.Post, app string, pubID string) *model.Transaction {
	metadataJson, _ := json.Marshal(post)

	return &model.Transaction{
		BlockNumber: momoka.BlockHeight.BigInt().Int64(),
		Timestamp:   momoka.Timestamp,
		Hash:        momoka.TransactionID,
		Owner:       owner,
		AddressFrom: owner,
		Network:     constant.NetworkArweave.String(),
		Platform:    protocol.PlatformLens,
		Tag:         filter.TagSocial,
		Type:        socialType,
		Success:     lo.ToPtr(true),
		Transfers: []model.Transfer{
			{
				TransactionHash: momoka.TransactionID,
				Timestamp:       momoka.Timestamp,
				BlockNumber:     momoka.BlockHeight.BigInt(),
				AddressFrom:     owner,
				Tag:             filter.TagSocial,
				Type:            socialType,
				Network:         constant.NetworkArweave.String(),
				Platform:        app,
				Metadata:        metadataJson,
				RelatedUrls: []string{
					fmt.Sprintf("https://momoka.lens.xyz/tx/%v", momoka.TransactionID),
					s.buildMomokaTransactionURL(pubID),
				},
			},
		},
	}
}

func (s *service) buildMomokaTransactionURL(pubID string) string {
	return fmt.Sprintf("https://lenster.xyz/posts/%v", pubID)
}
