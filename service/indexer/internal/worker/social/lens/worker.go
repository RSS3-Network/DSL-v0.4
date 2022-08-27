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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/datasource/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const (
	Name   = protocol.PlatformLens
	Source = "Lens ETL"
)

var _ worker.Worker = (*service)(nil)

type service struct {
	ethereumClient *ethclient.Client
	httpClient     *http.Client
}

type LensContent struct {
	Description string             `json:"description"`
	Content     string             `json:"content"`
	ExternalURL string             `json:"external_url"`
	AppId       string             `json:"appId"`
	Media       []LensContentMedia `json:"media"`
}

type LensContentMedia struct {
	Item string `json:"item"`
	Type string `json:"type"`
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkPolygon,
	}
}

func (s *service) Initialize(ctx context.Context) (err error) {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (internalTransactions []model.Transaction, err error) {
	tracer := otel.Tracer("worker_lens")
	_, trace := tracer.Start(ctx, "worker_len:Handle")

	defer func() { opentelemetry.Log(trace, transactions, internalTransactions, err) }()

	var mu sync.Mutex
	opt := lop.NewOption().WithConcurrency(ethereum.RPCMaxConcurrency)
	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		addressTo := common.HexToAddress(transaction.AddressTo)
		if addressTo != lens.HubProxyContractAddress && addressTo != lens.ProfileProxyContractAddress {
			return
		}

		transaction.Owner = message.Address
		transaction.Platform = s.Name()

		transferMap := make(map[int64]model.Transfer)
		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// Empty transfer data to avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)
		transaction.Transfers = append(transaction.Transfers, transferMap[protocol.IndexVirtual])

		// get receipt
		internalTransfers, err := s.handleReceipt(ctx, &transaction)
		if err != nil {
			logrus.Errorf("[lens worker] handleReceipt error, %v", err)

			return
		}

		transaction.Transfers = append(transaction.Transfers, internalTransfers...)

		for _, transfer := range transaction.Transfers {
			if transaction.Tag == "" {
				transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
			}
		}

		mu.Lock()
		internalTransactions = append(internalTransactions, transaction)
		mu.Unlock()
	}, opt)

	return internalTransactions, nil
}

func (s *service) handleReceipt(ctx context.Context, transaction *model.Transaction) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_lens")
	_, trace := tracer.Start(ctx, "worker_len:handleReceipt")

	defer func() { opentelemetry.Log(trace, transaction, transfers, err) }()

	// rpc
	receipt, err := s.ethereumClient.TransactionReceipt(ctx, common.HexToHash(transaction.Hash))
	if err != nil {
		logrus.Errorf("[lens worker] ethereumClient TransactionReceipt error, %v", err)

		return
	}

	// parse log
	for _, log := range receipt.Logs {

		lensContract, err := contract.NewEvents(log.Address, s.ethereumClient)
		if err != nil {
			loggerx.Global().Error("[lens worker] handleReceipt: new events error", zap.Error(err))

			continue
		}

		sourceData, err := json.Marshal(log)
		if err != nil {
			loggerx.Global().Error("marshal source data error", zap.Error(err))

			continue
		}

		// common attributes
		transfer := model.Transfer{
			TransactionHash: transaction.Hash,
			BlockNumber:     big.NewInt(transaction.BlockNumber),
			Index:           int64(log.Index),
			Network:         transaction.Network,
			AddressFrom:     transaction.Owner,
			SourceData:      sourceData,
			Source:          Source,
			Platform:        s.Name(),
		}

		var eventType string
		var handleErr error
		switch log.Topics[0] {
		case lens.EventHashPostCreated:
			handleErr = s.handlePostCreated(ctx, *lensContract, &transfer, *log)
			eventType = "handlePostCreated"
		case lens.EventHashCommentCreated:
			handleErr = s.handleCommentCreated(ctx, *lensContract, &transfer, *log)
			eventType = "handleCommentCreated"
		case lens.EventHashProfileCreated:
			handleErr = s.handleProfileCreated(ctx, *lensContract, transaction, &transfer, *log)
			eventType = "handleProfileCreated"
		default:
			continue
		}
		if handleErr != nil {
			loggerx.Global().Error(fmt.Sprintf("[lens worker] handleReceipt: %s error", eventType), zap.Error(handleErr))
			continue
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (s *service) handlePostCreated(ctx context.Context, lensContract contract.Events, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParsePostCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleReceipt: ParsePostCreated error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)

	// get content
	content, err := ipfs.GetFileByURL(event.ContentURI)
	if err != nil {
		logrus.Errorf("[lens worker] handleReceipt: getContent error, %v", err)
		return err
	}

	lensContent := LensContent{}
	if err = json.Unmarshal(content, &lensContent); err != nil {
		logrus.Errorf("[lens worker] handleReceipt: json unmarshal error, %v, json: %v", err, string(content))
		return err
	}

	post := &metadata.Post{
		Body: lensContent.Content,
	}
	for _, media := range lensContent.Media {
		post.Media = append(post.Media, metadata.Media{
			Address:  ipfs.GetDirectURL(media.Item),
			MimeType: media.Type,
		})
	}

	rawMetadata, err := json.Marshal(post)
	if err != nil {
		return err
	}

	transfer.Platform = lensContent.AppId
	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialPost, transfer.Type)
	transfer.RelatedUrls = []string{
		lensContent.ExternalURL,
		utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash),
	}

	return nil
}

func (s *service) handleCommentCreated(ctx context.Context, lensContract contract.Events, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParseCommentCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: ParsePostCreated error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)

	// get content
	content, err := ipfs.GetFileByURL(event.ContentURI)
	if err != nil {
		logrus.Errorf("[lens worker] handleCommentCreated: getContent error, %v", err)
		return err
	}

	lensContent := LensContent{}
	if err = json.Unmarshal(content, &lensContent); err != nil {
		logrus.Errorf("[lens worker] handleCommentCreated: json unmarshal error, %v, json: %v", err, string(content))
		return err
	}

	post := &metadata.Post{
		Body: lensContent.Content,
	}
	for _, media := range lensContent.Media {
		post.Media = append(post.Media, metadata.Media{
			Address:  ipfs.GetDirectURL(media.Item),
			MimeType: media.Type,
		})
	}

	// get content pointed
	post.Target, err = s.getContenPointed(ctx, event.ProfileId, event.PubId, s.ethereumClient)
	if err != nil {
		logrus.Errorf("[lens worker] handleCommentCreated: getContenPointed error, %v", err)
		return err
	}

	rawMetadata, err := json.Marshal(post)
	if err != nil {
		return err
	}

	transfer.Platform = lensContent.AppId
	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialComment, transfer.Type)
	transfer.RelatedUrls = []string{
		lensContent.ExternalURL,
		utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash),
	}

	return nil
}

func (s *service) handleProfileCreated(ctx context.Context, lensContract contract.Events, transaction *model.Transaction, transfer *model.Transfer, log types.Log) (err error) {
	event, err := lensContract.EventsFilterer.ParseProfileCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleProfileCreated: ParseProfileCreated error", zap.Error(err))

		return err
	}

	transfer.Timestamp = time.Unix(event.Timestamp.Int64(), 0)
	transfer.AddressFrom = strings.ToLower(event.Creator.String())
	transfer.AddressTo = strings.ToLower(event.To.String())

	profile := social.Profile{
		Address:  strings.ToLower(event.To.String()),
		Platform: Name,
		Network:  transaction.Network,
		Source:   transaction.Platform,
		Type:     filter.SocialProfileCreate,
		URL:      ipfs.GetDirectURL(event.ImageURI),
		Handle:   event.Handle,
	}

	rawMetadata, err := json.Marshal(&profile)
	if err != nil {
		return err
	}

	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfileCreate, transfer.Type)
	transfer.RelatedUrls = []string{
		fmt.Sprintf("https://lenster.xyz/u/%v", event.Handle),
		utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash),
	}

	transaction.Tag, transaction.Type = filter.UpdateTagAndType(filter.TagSocial, transaction.Tag, filter.SocialProfile, transaction.Type)

	return nil
}

func (s *service) getContenPointed(ctx context.Context, profileId *big.Int, pubId *big.Int, ethereumClient *ethclient.Client) (*metadata.Post, error) {
	// rpc
	iLensHub, err := contract.NewILensHub(lens.HubProxyContractAddress, ethereumClient)
	if err != nil {
		logrus.Errorf("[lens worker] NewILensHub error, %v", err)

		return nil, err
	}

	contentURI, err := iLensHub.GetContentURI(&bind.CallOpts{}, profileId, pubId)
	if err != nil {
		logrus.Errorf("[lens worker] iLensHub GetContentURI error, %v", err)

		return nil, err
	}

	// get content
	content, err := ipfs.GetFileByURL(contentURI)
	if err != nil {
		logrus.Errorf("[lens worker] getContenPointed: getContent error, %v", err)
		return nil, err
	}

	lensContent := LensContent{}
	if err = json.Unmarshal(content, &lensContent); err != nil {
		logrus.Errorf("[lens worker] getContenPointed: json unmarshal error, %v, json: %v", err, string(content))
		return nil, err
	}

	post := &metadata.Post{
		Body: lensContent.Content,
	}
	for _, media := range lensContent.Media {
		post.Media = append(post.Media, metadata.Media{
			Address:  ipfs.GetDirectURL(media.Item),
			MimeType: media.Type,
		})
	}

	return post, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{}
}

func New(ethereumClientMap map[string]*ethclient.Client) worker.Worker {
	return &service{
		ethereumClient: ethereumClientMap[protocol.NetworkPolygon],
		httpClient:     http.DefaultClient,
	}
}
