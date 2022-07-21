package poap

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const (
	Name = "poap"
)

var ContractAddress = common.HexToAddress("0x22C1f6050E56d2876009903609a2cC3fEf83B415")

var _ worker.Worker = (*service)(nil)

type service struct {
	tokenClient *token.Client
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkXDAI,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("poap_worker")
	_, trace := tracer.Start(ctx, "poap_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			if transfer.Index == protocol.IndexVirtual {
				continue
			}

			dataSource := types.Log{}
			if err := json.Unmarshal(transfer.SourceData, &dataSource); err != nil {
				logger.Global().Error("failed to unmarshal source data", zap.Error(err), zap.String("transaction_hash", transaction.Hash), zap.String("source_data", string(transfer.SourceData)))

				return nil, err
			}

			if dataSource.Address != ContractAddress {
				continue
			}

			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction
				value.Transfers = make([]model.Transfer, 0)

				internalTransactionMap[transaction.Hash] = transaction
			}

			var metadataModel metadata.Metadata

			if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
				return nil, err
			}

			erc721Filterer, err := erc721.NewERC721Filterer(dataSource.Address, nil)
			if err != nil {
				return nil, err
			}

			event, err := erc721Filterer.ParseTransfer(dataSource)
			if err != nil {
				return nil, err
			}

			nft, err := s.tokenClient.NFT(ctx, message.Network, ContractAddress.String(), event.TokenId)
			if err != nil {
				return nil, err
			}

			var tokenMetadata metadata.Token

			tokenMetadata.Name = nft.Name
			tokenMetadata.Symbol = nft.Symbol
			tokenMetadata.Description = nft.Description
			tokenMetadata.ID = nft.ID
			tokenMetadata.AnimationURL = nft.AnimationURL
			tokenMetadata.ExternalLink = nft.ExternalLink
			tokenMetadata.Standard = nft.Standard

			// The `home_url` and `year` fields are ignored for now

			for _, attribute := range nft.Attributes {
				tokenMetadata.Attributes = append(tokenMetadata.Attributes, metadata.TokenAttribute{
					TraitType: attribute.TraitType,
					Value:     attribute.Value,
				})
			}

			tokenMetadata.Metadata = nft.Metadata

			rawMetadata, err := json.Marshal(tokenMetadata)
			if err != nil {
				return nil, err
			}

			// Transfer
			transfer.Metadata = rawMetadata
			transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagCollectible, transfer.Tag, filter.CollectiblePoap, transfer.Type)
			if transfer.Tag == filter.TagCollectible {
				transfer.Platform = Name
			}

			// Transaction
			value.Tag, value.Type = filter.UpdateTagAndType(transfer.Tag, value.Tag, transfer.Type, value.Type)
			if value.Tag == transfer.Tag {
				value.Platform = transfer.Platform
			}

			value.Transfers = append(value.Transfers, transfer)

			internalTransactionMap[value.Hash] = value
		}
	}

	internalTransfers := make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		internalTransfers = append(internalTransfers, internalTransaction)
	}

	return internalTransfers, nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func New(ethereumClientMap map[string]*ethclient.Client) worker.Worker {
	return &service{
		tokenClient: token.New(nil, ethereumClientMap),
	}
}
