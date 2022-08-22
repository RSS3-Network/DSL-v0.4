package lens

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/lens/contract"
	"github.com/naturalselectionlabs/pregod/common/datasource/pregod_etl"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const (
	Source = "pregod_etl"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	ethereumClientMap map[string]*ethclient.Client  // QuickNode
	clientMap         map[string]*pregod_etl.Client // PregodETL
}

func (d *Datasource) Name() string {
	return Source
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

	ethereumClient, exist := d.ethereumClientMap[message.Network]
	if !exist {
		return nil, errors.New("unsupported network")
	}

	// get profileid by address
	profileID, err := d.getDefaultProfile(ctx, message, ethereumClient)
	if err != nil {
		return nil, err
	}

	// get transaction
	transactionMap, err := d.getLensTransferHashes(ctx, message, profileID)
	if err != nil {
		return nil, err
	}

	transactions := make([]*model.Transaction, 0)
	for _, transaction := range transactionMap {
		internalTransaction := transaction

		// if internalTransaction.BlockNumber < message.BlockNumber {
		// 	continue
		// }

		if internalTransaction.AddressFrom != "" && !strings.EqualFold(internalTransaction.AddressFrom, message.Address) {
			continue
		}

		transactions = append(transactions, &internalTransaction)
	}

	// build transaction
	if transactions, err = ethereum.BuildTransactions(ctx, message, transactions, ethereumClient); err != nil {
		loggerx.Global().Error("failed to build transactions", zap.Error(err))

		return nil, err
	}

	for _, transaction := range transactions {
		if transaction != nil {
			internalTransactions = append(internalTransactions, *transaction)
		}
	}

	return internalTransactions, nil
}

func (d *Datasource) getDefaultProfile(ctx context.Context, message *protocol.Message, ethereumClient *ethclient.Client) (*big.Int, error) {
	tracer := otel.Tracer("datasource_lens")
	_, trace := tracer.Start(ctx, "datasource_lens:getDefaultProfile")
	var profileID *big.Int
	var err error
	defer func() { opentelemetry.Log(trace, message, profileID, err) }()

	// rpc
	iLensHub, err := contract.NewILensHub(lens.HubProxyContractAddress, ethereumClient)
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

	var internalTransactions []model.Transaction
	hash := common.HexToHash(hexutil.EncodeBig(profileID))

	var wg sync.WaitGroup
	for eventHash, contractAddress := range lens.SupportLensEvents {
		wg.Add(1)
		go func(eventHash common.Hash, contractAddress common.Address) {
			defer wg.Done()
			parameter := pregod_etl.GetLogsRequest{
				Address:     contractAddress.String(),
				TopicSecond: hash.String(),
				TopicFirst:  eventHash.String(),
			}

			result, err := d.clientMap[message.Network].GetLogs(ctx, parameter)
			if err != nil {
				return
			}

			for _, transfer := range result {
				transaction := model.Transaction{
					BlockNumber: transfer.BlockNumber.BigInt().Int64(),
					Hash:        transfer.TransactionHash,
					Index:       transfer.TransactionIndex.BigInt().Int64(),
					Network:     message.Network,
					Source:      d.Name(),
					Transfers:   make([]model.Transfer, 0),
					Owner:       strings.ToLower(message.Address),
				}

				internalTransactions = append(internalTransactions, transaction)
			}
		}(eventHash, contractAddress)
	}
	wg.Wait()

	for _, transaction := range internalTransactions {
		internalTransactionMap[transaction.Hash] = transaction
	}

	return internalTransactionMap, nil
}

func New(config *configx.RPC) (datasource.Datasource, error) {
	internalDatasource := Datasource{
		ethereumClientMap: map[string]*ethclient.Client{},
		clientMap:         map[string]*pregod_etl.Client{},
	}

	var err error

	if internalDatasource.ethereumClientMap[protocol.NetworkEthereum], err = ethclient.Dial(config.General.Ethereum.HTTP); err != nil {
		logrus.Errorf("[datasource_lens] ethclient.Dial error, %v", err)
		return nil, err
	}

	if internalDatasource.ethereumClientMap[protocol.NetworkPolygon], err = ethclient.Dial(config.General.Polygon.HTTP); err != nil {
		logrus.Errorf("[datasource_lens] ethclient.Dial error, %v", err)
		return nil, err
	}

	if internalDatasource.clientMap[protocol.NetworkPolygon], err = pregod_etl.NewClient(protocol.NetworkPolygon, config.PregodETL.Polygon.HTTP); err != nil {
		logrus.Errorf("[datasource_lens] pregod_etl.NewClient error, %v", err)
		return nil, err
	}

	return &internalDatasource, nil
}