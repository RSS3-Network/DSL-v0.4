package friendtech

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/collectibe"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/friendtech"
	friendtechClient "github.com/naturalselectionlabs/pregod/common/datasource/friendtech"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/collectible/friendtech/job"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ worker.Worker = &service{}

type service struct {
	client      *friendtechClient.Client
	tokenClient *token.Client
}

func (s *service) Name() string {
	return protocol.PlatformFriendTech
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkBase,
	}
}

func (s *service) Initialize(_ context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("worker_friend.tech")
	_, trace := tracer.Start(ctx, "worker_friend.tech:Handle")

	defer opentelemetry.Log(trace, transactions, data, err)

	return s.handleEthereum(ctx, message, transactions)
}

func (s *service) handleEthereum(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("worker_friend.tech")
	_, trace := tracer.Start(ctx, "worker_friend.tech:handleEthereum")

	defer opentelemetry.Log(trace, transactions, data, err)

	internalTransactionMap := make(map[string]model.Transaction)
	opt := lop.NewOption().WithConcurrency(ethereum.RPCMaxConcurrency)

	var mu sync.Mutex

	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		if !strings.EqualFold(transaction.AddressTo, friendtech.AddressFriendTech.String()) {
			return
		}

		if transaction.Transfers, err = s.handleEthereumTransaction(ctx, message, &transaction); err != nil {
			loggerx.Global().Warn("failed to handle ethereum transaction", zap.Error(err), zap.String("network", message.Network), zap.String("transaction_hash", transaction.Hash), zap.String("address", message.Address))

			return
		}

		transaction.Tag, transaction.Type = filter.UpdateTagAndType(filter.TagCollectible, transaction.Tag, filter.CollectibleTrade, transaction.Type)

		transaction.Owner = transaction.AddressFrom
		transaction.Platform = s.Name()

		mu.Lock()
		internalTransactionMap[transaction.Hash] = transaction
		mu.Unlock()
	}, opt)

	// Lay the map flat
	internalTransactions := make([]model.Transaction, 0)

	for _, transaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, transaction)
	}

	return internalTransactions, nil
}

func (s *service) handleEthereumTransaction(ctx context.Context, message *protocol.Message, transaction *model.Transaction) (internalTransfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_friend.tech")
	_, trace := tracer.Start(ctx, "worker_friend.tech:handleEthereumTransaction")

	defer opentelemetry.Log(trace, transaction, internalTransfers, err)

	ethereumClient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, errors.New("not supported network")
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source data: %w", err)
	}

	for _, log := range sourceData.Receipt.Logs {
		// Filter anonymous log
		if len(log.Topics) == 0 {
			continue
		}

		var transfer *model.Transfer

		switch log.Topics[0] {
		case friendtech.EventTrade:
			transfer, err = s.buildEthereumTransfer(ctx, message, *log, transaction, ethereumClient)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		if transfer != nil {
			internalTransfers = append(internalTransfers, *transfer)
		}

	}

	return internalTransfers, nil
}

func (s *service) buildEthereumTransfer(ctx context.Context, message *protocol.Message, log types.Log, transaction *model.Transaction, ethereumClient *ethclient.Client) (*model.Transfer, error) {
	friendTechContract, err := friendtech.NewFriendTech(log.Address, ethereumClient)
	if err != nil {
		return nil, err
	}

	event, err := friendTechContract.ParseTrade(log)
	if err != nil {
		return nil, err
	}

	var user *friendtechClient.UserResponse

	// get subject metadata
	user, err = s.client.GetUserMeta(ctx, strings.ToLower(event.Subject.String()))

	if err != nil || user == nil {
		// search user from db
		var item collectibe.FriendTech

		if err := database.Global().
			Model(&collectibe.FriendTech{}).
			First(&item, "address", strings.ToLower(event.Subject.String())).Error; err != nil {
			logrus.Errorf("not found friend tech item, db error: %v", err)
			return nil, err
		}

		user = &friendtechClient.UserResponse{
			ID:              item.ID,
			TwitterUsername: item.TwitterUsername,
			TwitterName:     item.TwitterName,
			Address:         item.Address,
			TwitterPfpURL:   item.TwitterPfpUrl,
		}
	}

	nativeToken, err := s.tokenClient.Native(ctx, message.Network)
	if err != nil {
		return nil, err
	}

	costValue := decimal.NewFromBigInt(event.EthAmount, 0)
	costValueDisplay := costValue.Shift(-int32(nativeToken.Decimals))

	friendTechMeta := &metadata.Token{
		ID:              strconv.FormatInt(user.ID, 10),
		Name:            user.TwitterUsername,
		Symbol:          user.TwitterName,
		ContractAddress: user.Address,
		Collection:      "friend.tech: Shares",
		Image:           user.TwitterPfpURL,
		Standard:        protocol.TokenStandardERC1155,
		Value:           lo.ToPtr(decimal.NewFromBigInt(event.ShareAmount, 0)),
		ValueDisplay:    lo.ToPtr(decimal.NewFromBigInt(event.ShareAmount, 0)),
		Cost: &metadata.Token{
			Name:         nativeToken.Name,
			Symbol:       nativeToken.Symbol,
			Decimals:     nativeToken.Decimals,
			Standard:     protocol.TokenStandardNative,
			Image:        nativeToken.Logo,
			Value:        &costValue,
			ValueDisplay: &costValueDisplay,
		},
	}

	ftMetadata, err := json.Marshal(friendTechMeta)
	if err != nil {
		return nil, err
	}
	var (
		addressFrom string
		addressTo   string
	)
	if event.IsBuy {
		addressFrom, addressTo = strings.ToLower(event.Subject.String()), strings.ToLower(event.Trader.String())
	} else {
		addressFrom, addressTo = strings.ToLower(event.Trader.String()), strings.ToLower(event.Subject.String())
	}

	transfer := &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     addressFrom,
		AddressTo:       addressTo,
		Metadata:        ftMetadata,
		Network:         message.Network,
		Platform:        s.Name(),
		Tag:             filter.TagCollectible,
		Type:            filter.CollectibleTrade,
		Source:          protocol.SourceOrigin,
		RelatedUrls: []string{
			ethereum.BuildScanURL(message.Network, transaction.Hash),
		},
	}

	return transfer, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		job.New(),
	}
}

func New() worker.Worker {
	svc := service{
		client:      friendtechClient.NewClient(),
		tokenClient: token.New(),
	}
	return &svc
}
