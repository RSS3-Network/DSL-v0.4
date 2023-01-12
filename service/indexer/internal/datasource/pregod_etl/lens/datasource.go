package lens

import (
	"context"
	"math/big"
	"net/http"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common/hexutil"
	kurora "github.com/naturalselectionlabs/kurora/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	lens_comm "github.com/naturalselectionlabs/pregod/common/worker/lens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	kuroraClient *kurora.Client
	lensClient   *lens_comm.Client
}

func (d *Datasource) Name() string {
	return protocol.SourceKurora
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkPolygon,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	tracer := otel.Tracer("datasource_lens")
	ctx, trace := tracer.Start(ctx, "datasource_lens:Handle")
	internalTransactions := make([]model.Transaction, 0)
	var err error
	defer func() { opentelemetry.Log(trace, message, len(internalTransactions), err) }()

	// get profileid by address
	profileIDList, err := d.lensClient.BatchGetProfileID(ctx, message.Address)
	if err != nil {
		return nil, err
	}

	if len(profileIDList) == 0 {
		logrus.Infof("[datasource_lens] Handle GetProfiles is nil, address: %v", message.Address)

		return nil, nil
	}

	// get transaction
	for _, profileID := range profileIDList {
		transactionMap, err := d.getLensTransferHashes(ctx, message, profileID)
		if err != nil {
			return nil, err
		}

		transactions := make([]*model.Transaction, 0)
		for _, transaction := range transactionMap {
			internalTransaction := transaction

			if internalTransaction.Source != protocol.SourceKurora && internalTransaction.AddressFrom != "" && !strings.EqualFold(internalTransaction.AddressFrom, message.Address) {
				continue
			}

			transactions = append(transactions, &internalTransaction)
		}

		// build transaction
		if transactions, err = ethereum.BuildTransactions(ctx, message, transactions); err != nil {
			loggerx.Global().Error("failed to build transactions", zap.Error(err))

			return nil, err
		}

		for _, transaction := range transactions {
			if transaction != nil {
				internalTransactions = append(internalTransactions, *transaction)
			}
		}
	}

	return internalTransactions, nil
}

func (d *Datasource) GetDefaultProfile(ctx context.Context, message *protocol.Message) (*big.Int, error) {
	tracer := otel.Tracer("datasource_lens")
	_, trace := tracer.Start(ctx, "datasource_lens:getDefaultProfile")
	var profileID *big.Int
	var err error
	defer func() { opentelemetry.Log(trace, message, profileID, err) }()

	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	// rpc
	iLensHub, err := contract.NewILensHub(lens.HubProxyContractAddress, ethclient)
	if err != nil {
		logrus.Errorf("[datasource_lens] NewILensHub error, %v", err)

		return nil, err
	}

	profileID, err = iLensHub.DefaultProfile(&bind.CallOpts{}, common.HexToAddress(message.Address))
	if err != nil {
		logrus.Errorf("[datasource_lens] iLensHub DefaultProfile error, %v", err)

		return nil, err
	}

	return profileID, nil
}

func (d *Datasource) getLensTransferHashes(ctx context.Context, message *protocol.Message, profileID *big.Int) (map[string]model.Transaction, error) {
	tracer := otel.Tracer("datasource_lens")
	ctx, trace := tracer.Start(ctx, "datasource_lens:getLensTransferHashes")
	internalTransactionMap := make(map[string]model.Transaction)
	var err error
	defer func() { opentelemetry.Log(trace, profileID, internalTransactionMap, err) }()

	hash := common.HexToHash(hexutil.EncodeBig(profileID))

	var wg sync.WaitGroup
	var mu sync.Mutex
	for eventHash, contractAddress := range lens.SupportLensEvents {
		wg.Add(1)
		go func(eventHash common.Hash, contractAddress common.Address) {
			defer wg.Done()

			query := kurora.DatasetLensEventQuery{
				TopicFirst:  &eventHash,
				TopicSecond: &hash,
				Address:     &contractAddress,
				// BlockNumberFrom: message.BlockNumber,
			}

			if eventHash == lens.EventHashFollowed {
				address := common.HexToHash(message.Address)
				query.TopicSecond = &address
			}

			loggerx.Global().Info("query kurora FetchDatasetLensEvents", zap.Any("query", query))

			for {
				result, err := d.kuroraClient.FetchDatasetLensEvents(ctx, query)
				if err != nil {
					loggerx.Global().Error("FetchDatasetLensEvents error", zap.Error(err), zap.Any("query", query))
					return
				}

				if len(result) == 0 {
					return
				}

				for _, transfer := range result {
					mu.Lock()
					internalTransactionMap[transfer.TransactionHash.String()] = model.Transaction{
						BlockNumber: transfer.BlockNumber.BigInt().Int64(),
						Hash:        transfer.TransactionHash.String(),
						Index:       int64(transfer.TransactionIndex),
						Network:     message.Network,
						Transfers:   make([]model.Transfer, 0),
						Source:      protocol.SourceKurora,
					}
					mu.Unlock()
				}

				cursor := kurora.LogCursor(result[len(result)-1].TransactionHash, result[len(result)-1].Index)
				query.Cursor = &cursor
			}
		}(eventHash, contractAddress)
	}
	wg.Wait()

	return internalTransactionMap, nil
}

func New(ctx context.Context, endpoint string) (datasource.Datasource, error) {
	internalDatasource := Datasource{
		lensClient: lens_comm.New(),
	}

	var err error

	internalDatasource.kuroraClient, err = kurora.Dial(ctx, endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error(" kurora Dial error", zap.Error(err))
		return nil, err
	}

	return &internalDatasource, nil
}
