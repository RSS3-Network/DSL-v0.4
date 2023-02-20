package music

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/sound"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/sound/contract"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	lop "github.com/samber/lo/parallel"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var (
	_       worker.Worker = (*service)(nil)
	network []string
)

type service struct{}

type SoundContractMetadata struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ExternalLink string `json:"external_link"`
}

func (s *service) Name() string {
	return "metaverse"
}

func (s *service) Networks() []string {
	return network
}

func (s *service) Initialize(ctx context.Context) error {
	for _, contract := range EventsRouterMap {
		network = append(network, contract.Network)
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("music_worker")
	_, handlerSpan := tracer.Start(ctx, "music_worker:Handle")

	defer handlerSpan.End()

	internalTransactions := make([]model.Transaction, 0)

	lop.ForEach(transactions, func(transaction model.Transaction, i int) {
		var action ContractAction
		var sourceData ethereum.SourceData
		if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
			zap.L().Error("failed to unmarshal source data", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("network", transaction.Network))

			return
		}

		if sourceData.Receipt == nil {
			return
		}

		for _, log := range sourceData.Receipt.Logs {
			// Filter anonymous log
			if len(log.Topics) == 0 {
				continue
			}

			for _, topic := range log.Topics {
				if contractAction, exist := EventsRouterMap[topic]; exist && contractAction.Network == message.Network {

					// check
					if contractAction.ContractAddress != nil && strings.ToLower(contractAction.ContractAddress.String()) != transaction.AddressTo {
						continue
					}

					if contractAction.Platform == protocol.PlatformSound && contractAction.ContractAddress == nil {
						if err := s.CheckSoundContract(ctx, log.Address); err != nil {
							continue
						}
					}

					action = contractAction
					goto update
				}
			}
		}

	update:
		for index, transfer := range transaction.Transfers {
			if action.Type != "" {
				transaction.Transfers[index].Tag, transaction.Transfers[index].Type = filter.UpdateTagAndType(action.Tag, transaction.Transfers[index].Tag, action.Type, transaction.Transfers[index].Type)

				// metadata
				var nftMetadata metadata.Token
				if err := json.Unmarshal(transfer.Metadata, &nftMetadata); err != nil {
					loggerx.Global().Error("nft metadata unmarshal error", zap.Error(err))

					continue
				}
				nftMetadata.Action = action.Action
				metadataJson, err := json.Marshal(nftMetadata)
				if err != nil {
					loggerx.Global().Error("nft metadata unmarshal error", zap.Error(err))

					continue
				}
				transaction.Transfers[index].Metadata = metadataJson
			}

			transaction.Tag, transaction.Type = filter.UpdateTagAndType(transaction.Transfers[index].Tag, transaction.Tag, transaction.Transfers[index].Type, transaction.Type)

			if transaction.Transfers[index].Type == filter.CollectibleMusic {
				transaction.Transfers[index].Platform, transaction.Platform = action.Platform, action.Platform
			}
		}

		internalTransactions = append(internalTransactions, transaction)
	})

	return internalTransactions, nil
}

func (s *service) CheckSoundContract(ctx context.Context, contractAddress common.Address) error {
	// rpc
	ethclient, err := ethclientx.Global(protocol.NetworkEthereum)
	if err != nil {
		return err
	}

	soundContract, err := contract.NewArtist(contractAddress, ethclient)
	if err != nil {
		loggerx.Global().Error("new artist contract error", zap.Error(err))

		return err
	}

	name, err := soundContract.Name(&bind.CallOpts{})
	if err != nil {
		loggerx.Global().Error("get artist name error", zap.Error(err))

		return err
	}

	if len(name) == 0 {
		return fmt.Errorf("artist is nil")
	}

	return nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{}
}

func New() worker.Worker {
	return &service{}
}
