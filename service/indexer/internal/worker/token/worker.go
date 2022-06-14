package token

import (
	"context"
	"embed"
	"encoding/csv"
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	moralisx "github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/nft"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token/coinmarketcap"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	_ worker.Worker = (*service)(nil)

	//go:embed asset/*
	asseFS embed.FS
)

type service struct {
	databaseClient *gorm.DB
	zksyncClient   *zksync.Client
}

func (s *service) Name() string {
	return "token"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain, protocol.NetworkZkSync,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	dir := "asset/cex_wallet"
	files, err := asseFS.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, path := range files {
		file, err := asseFS.Open(dir + "/" + path.Name())
		if err != nil {
			return err
		}

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			return err
		}

		file.Close()

		wallentModels := make([]model.ExchangeWallet, 0)

		for i, record := range records {
			if i == 0 {
				continue
			}

			wallentModels = append(wallentModels, model.ExchangeWallet{
				WalletAddress: record[0],
				Name:          record[1],
				Source:        record[2],
				Network:       record[3],
			})
		}

		if len(wallentModels) == 0 {
			return nil
		}

		if err := s.databaseClient.
			Model((*model.ExchangeWallet)(nil)).
			Clauses(clause.OnConflict{
				DoNothing: true,
			}).
			Create(wallentModels).
			Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error) {
	tracer := otel.Tracer("worker_token")

	ctx, handlerSpan := tracer.Start(ctx, "handler")

	defer handlerSpan.End()

	ctx, databaseSpan := tracer.Start(ctx, "database") //nolint:ineffassign,staticcheck

	// TODO

	databaseSpan.End()

	// ZkSync
	if message.Network == protocol.NetworkZkSync {
		return s.handleZkSync(ctx, message, transfers)
	}

	internalTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {
		if transfer.Source != moralis.Source {
			continue
		}

		sourceDataMap := make(map[string]interface{})

		if err := json.Unmarshal(transfer.SourceData, &sourceDataMap); err != nil {
			return nil, err
		}

		var metadataModel metadata.Metadata

		if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
			return nil, err
		}

		if _, exist := sourceDataMap["contract_type"]; exist {
			// NFT transfer
			nftTransfer := moralisx.NFTTransfer{}
			if err := json.Unmarshal(transfer.SourceData, &nftTransfer); err != nil {
				return nil, err
			}

			tokenID, err := decimal.NewFromString(nftTransfer.TokenId)
			if err != nil {
				return nil, err
			}

			nftMetadata, err := nft.GetMetadata(
				message.Network,
				common.HexToAddress(nftTransfer.TokenAddress),
				tokenID.BigInt(),
			)
			if err != nil { // print err but no return
				logrus.Errorf("%s: %v", message.Network, err)
			}

			tokenValue, err := decimal.NewFromString(nftTransfer.Amount)
			if err != nil {
				logrus.Error(err)
			}

			metadataModel.Token = &metadata.Token{
				TokenAddress:  nftTransfer.TokenAddress,
				TokenStandard: strings.ToLower(nftTransfer.ContractType),
				TokenID:       &tokenID,
				TokenValue:    &tokenValue,
				NFTMetadata:   nftMetadata,
			}
		} else if _, exist = sourceDataMap["address"]; exist {
			// Token transfer
			tokenTransfer := moralisx.TokenTransfer{}
			if err := json.Unmarshal(transfer.SourceData, &tokenTransfer); err != nil {
				return nil, err
			}

			tokenValue, err := decimal.NewFromString(tokenTransfer.Value)
			if err != nil {
				return nil, err
			}

			// get info from coinmarketcap
			coinInfo, err := coinmarketcap.CachedGetCoinInfo(ctx, message.Network, sourceDataMap["address"].(string))
			if err != nil {
				return nil, err
			}

			metadataModel.Token = &metadata.Token{
				TokenAddress:  tokenTransfer.Address,
				TokenStandard: "erc20",
				TokenValue:    &tokenValue,
				Logo:          coinInfo.Logo,
				Name:          coinInfo.Name,
				Symbol:        coinInfo.Symbol,
				Decimals:      coinInfo.Decimals,
			}
		} else {
			// Native transfer
			nativeTransfer := moralisx.Transaction{}
			if err := json.Unmarshal(transfer.SourceData, &nativeTransfer); err != nil {
				return nil, err
			}

			tokenValue, err := decimal.NewFromString(nativeTransfer.Value)
			if err != nil {
				return nil, err
			}

			// get info from coinmarketcap
			coinInfo, err := coinmarketcap.CachedGetCoinInfoByNetwork(ctx, message.Network)
			if err != nil {
				return nil, err
			}

			metadataModel.Token = &metadata.Token{
				TokenStandard: "native",
				TokenValue:    &tokenValue,
				Logo:          coinInfo.Logo,
				Name:          coinInfo.Name,
				Symbol:        coinInfo.Symbol,
				Decimals:      coinInfo.Decimals,
			}
		}

		rawMetadata, err := json.Marshal(metadataModel)
		if err != nil {
			return nil, err
		}

		transfer.Metadata = rawMetadata

		// action type
		switch {
		case !strings.EqualFold(sourceDataMap["from_address"].(string), message.Address): // TO == self
			transfer.Type = "receive"
		case !strings.EqualFold(sourceDataMap["to_address"].(string), message.Address): // FROM == self
			transfer.Type = "send"
		default: // FROM == TO == self
			transfer.Type = "cancel"
		}

		internalTransfers = append(internalTransfers, transfer)
	}

	return internalTransfers, nil
}

// TODO: only support `Transfer` now
// https://docs.zksync.io/apiv02-docs/#transactions-api-v0.2-transactions-txhash-data
func (s *service) handleZkSync(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {
		data := zksync.GetTransactionData{}
		if err := json.Unmarshal(transfer.SourceData, &data); err != nil {
			logrus.Error(err)
			return nil, err
		}

		tokenInfo, _, err := s.zksyncClient.GetToken(ctx, uint(data.Transaction.Operation.Token))
		if err != nil {
			logrus.Error(err)
		}

		var metadataModel metadata.Metadata

		amount, err := decimal.NewFromString(data.Transaction.Operation.Amount)
		if err != nil {
			logrus.Error(err)
		}

		// e.g. NFT-387049
		if strings.HasPrefix(tokenInfo.Symbol, "NFT") {
			nftTokenInfo, _, err := s.zksyncClient.GetNFTToken(ctx, uint(data.Transaction.Operation.Token))
			if err != nil {
				logrus.Error(err)
			}
			tokenID := decimal.NewFromInt(*nftTokenInfo.ID)
			metadataModel.Token = &metadata.Token{
				TokenAddress:  nftTokenInfo.Address,
				TokenStandard: "erc721",
				TokenID:       &tokenID,
				TokenValue:    &amount,
				Symbol:        nftTokenInfo.Symbol,
				NFTMetadata:   nftTokenInfo.Bytes(),
			}
		} else { // token
			tokenID := decimal.NewFromInt(*tokenInfo.ID)
			metadataModel.Token = &metadata.Token{
				TokenAddress:  tokenInfo.Address,
				TokenStandard: "erc20",
				TokenID:       &tokenID,
				TokenValue:    &amount,
				Decimals:      tokenInfo.Decimals,
				Symbol:        tokenInfo.Symbol,
			}
		}

		rawMetadata, err := json.Marshal(metadataModel)
		if err != nil {
			return nil, err
		}

		transfer.Metadata = rawMetadata

		// action type
		switch {
		case !strings.EqualFold(data.Transaction.Operation.From, message.Address): // TO == self
			transfer.Type = "receive"
		case !strings.EqualFold(data.Transaction.Operation.To, message.Address): // FROM == self
			transfer.Type = "send"
		default: // FROM == TO == self
			transfer.Type = "cancel"
		}

		internalTransfers = append(internalTransfers, transfer)

	}
	return internalTransfers, nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		zksyncClient:   zksync.New(),
	}
}
