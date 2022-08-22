package lens

import (
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

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
	Name = protocol.PlatformLens
)

var _ worker.Worker = (*service)(nil)

type service struct {
	ethereumClient *ethclient.Client
	httpClient     *http.Client
}

type LensContent struct {
	Description string `json:"description"`
	Content     string `json:"content"`
	ExternalURL string `json:"external_url"`
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

	opt := lop.NewOption().WithConcurrency(ethereum.RPCMaxConcurrency)
	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		addressTo := common.HexToAddress(transaction.AddressTo)
		if addressTo != lens.HubProxyContractAddress && addressTo != lens.ProfileProxyContractAddress {
			return
		}

		transaction.Platform = protocol.PlatformLens
		transaction.Owner = message.Address

		transferMap := make(map[int64]model.Transfer)

		for _, transfer := range transaction.Transfers {
			transferMap[transfer.Index] = transfer
		}

		// avoid data duplication
		transaction.Transfers = make([]model.Transfer, 0)
		if _, exists := transferMap[protocol.IndexVirtual]; exists {
			transaction.Transfers = append(transaction.Transfers, transferMap[protocol.IndexVirtual])
		}

		internalTransfers, err := s.handleReceipt(ctx, transaction)
		if err != nil {
			logrus.Errorf("[crossbell worker] handleReceipt error, %v", err)

			return
		}

		transaction.Transfers = append(transaction.Transfers, internalTransfers...)

		for _, transfer := range transaction.Transfers {
			transaction.Tag, transaction.Type = filter.UpdateTagAndType(transfer.Tag, transaction.Tag, transfer.Type, transaction.Type)
		}

		internalTransactions = append(internalTransactions, transaction)
	}, opt)

	return internalTransactions, nil
}

func (s *service) handleReceipt(ctx context.Context, transaction model.Transaction) (transfers []model.Transfer, err error) {
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
		switch log.Topics[0] {
		case lens.EventHashPostCreated:
			transfer, err := s.handlePostCreated(ctx, transaction, *log)
			if err != nil {
				loggerx.Global().Error("[lens worker] handleReceipt: handlePostCreated error", zap.Error(err))

				continue
			}
			transfers = append(transfers, transfer)

		case lens.EventHashCommentCreated:
			transfer, err := s.handleCommentCreated(ctx, transaction, *log)
			if err != nil {
				loggerx.Global().Error("[lens worker] handleReceipt: handleCommentCreated error", zap.Error(err))

				continue
			}
			transfers = append(transfers, transfer)

		case lens.EventHashProfileCreated:
			transfer, err := s.handleProfileCreated(ctx, transaction, *log)
			if err != nil {
				loggerx.Global().Error("[lens worker] handleReceipt: handleProfileCreated error", zap.Error(err))

				continue
			}
			transfers = append(transfers, transfer)

		}
	}

	return transfers, nil
}

func (s *service) handlePostCreated(ctx context.Context, transaction model.Transaction, log types.Log) (transfer model.Transfer, err error) {
	lensContract, err := contract.NewEvents(log.Address, s.ethereumClient)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleReceipt: new events error", zap.Error(err))

		return transfer, err
	}

	sourceData, err := json.Marshal(log)
	if err != nil {
		loggerx.Global().Error("marshal source data error", zap.Error(err))

		return transfer, err
	}

	event, err := lensContract.EventsFilterer.ParsePostCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleReceipt: ParsePostCreated error", zap.Error(err))

		return transfer, err
	}

	transfer = model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       time.Unix(event.Timestamp.Int64(), 0),
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     transaction.Owner,
		Network:         transaction.Network,
		Source:          ethereum.Source,
		SourceData:      sourceData,
	}

	// get content
	content, err := s.getContent(event.ContentURI)
	if err != nil {
		logrus.Errorf("[lens worker] Handle: getContent error, %v", err)
		return transfer, err
	}

	lensContent := LensContent{}

	if err = json.Unmarshal(content, &lensContent); err != nil {
		logrus.Errorf("[lens worker] Handle: json unmarshal error, %v", err)
		return transfer, err
	}

	rawMetadata, err := json.Marshal(&metadata.Post{
		Body: lensContent.Content,
	})
	if err != nil {
		return transfer, err
	}

	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialPost, transfer.Type)
	transfer.RelatedUrls = []string{utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash)}

	return transfer, nil
}

func (s *service) handleCommentCreated(ctx context.Context, transaction model.Transaction, log types.Log) (transfer model.Transfer, err error) {
	lensContract, err := contract.NewEvents(log.Address, s.ethereumClient)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: new events error", zap.Error(err))

		return transfer, err
	}

	sourceData, err := json.Marshal(log)
	if err != nil {
		loggerx.Global().Error("marshal source data error", zap.Error(err))

		return transfer, err
	}

	event, err := lensContract.EventsFilterer.ParseCommentCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleCommentCreated: ParsePostCreated error", zap.Error(err))

		return transfer, err
	}

	transfer = model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       time.Unix(event.Timestamp.Int64(), 0),
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     transaction.Owner,
		Network:         transaction.Network,
		Source:          ethereum.Source,
		SourceData:      sourceData,
	}

	// get content
	content, err := s.getContent(event.ContentURI)
	if err != nil {
		logrus.Errorf("[lens worker] handleCommentCreated: getContent error, %v", err)
		return transfer, err
	}

	lensContent := LensContent{}

	if err = json.Unmarshal(content, &lensContent); err != nil {
		logrus.Errorf("[lens worker] handleCommentCreated: json unmarshal error, %v", err)
		return transfer, err
	}

	rawMetadata, err := json.Marshal(&metadata.Post{
		Body: lensContent.Content,
	})
	if err != nil {
		return transfer, err
	}

	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialComment, transfer.Type)
	transfer.RelatedUrls = []string{utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash), lensContent.ExternalURL}

	return transfer, nil
}

func (s *service) handleProfileCreated(ctx context.Context, transaction model.Transaction, log types.Log) (transfer model.Transfer, err error) {
	lensContract, err := contract.NewEvents(log.Address, s.ethereumClient)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleProfileCreated: new events error", zap.Error(err))

		return transfer, err
	}

	sourceData, err := json.Marshal(log)
	if err != nil {
		loggerx.Global().Error("marshal source data error", zap.Error(err))

		return transfer, err
	}

	event, err := lensContract.EventsFilterer.ParseProfileCreated(log)
	if err != nil {
		loggerx.Global().Error("[lens worker] handleProfileCreated: ParseProfileCreated error", zap.Error(err))

		return transfer, err
	}

	transfer = model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       time.Unix(event.Timestamp.Int64(), 0),
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(event.Creator.String()),
		AddressTo:       strings.ToLower(event.To.String()),
		Network:         transaction.Network,
		Source:          ethereum.Source,
		SourceData:      sourceData,
	}

	profile := social.Profile{
		Address:  strings.ToLower(event.To.String()),
		Platform: transaction.Platform,
		Network:  transaction.Network,
		Source:   transaction.Platform,
		Type:     filter.SocialProfileCreate,
		URL:      ipfs.GetDirectURL(event.ImageURI),
		Handle:   event.Handle,
	}

	rawMetadata, err := json.Marshal(&profile)
	if err != nil {
		return transfer, err
	}

	transfer.Metadata = rawMetadata
	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfileCreate, transfer.Type)
	transfer.RelatedUrls = []string{utils.GetTxHashURL(protocol.NetworkPolygon, transfer.TransactionHash)}

	return transfer, nil
}

func (s *service) getContent(uri string) (json.RawMessage, error) {
	uri = ipfs.GetDirectURL(uri)
	response, err := s.httpClient.Get(uri)
	if err != nil {
		return nil, err
	}

	reader := response.Body

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	_ = reader.Close()

	return data, nil
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
