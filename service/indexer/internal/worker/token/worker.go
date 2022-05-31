package token

import (
	"context"
	"embed"
	"encoding/csv"
	"encoding/json"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	moralisx "github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
)

var (
	_ worker.Worker = (*service)(nil)

	//go:embed asset/*
	asseFS embed.FS
)

type service struct {
	databaseClient *gorm.DB
}

func (s *service) Name() string {
	return "token"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
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

		// file.Close()

		wallentModels := make([]model.ExchangeWallet, 0)

		for i, record := range records {
			if i == 0 {
				continue
			}

			wallentModels = append(wallentModels, model.ExchangeWallet{
				WalletAddress: record[0],
				Name:          record[1],
				Source:        record[2],
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

			tokenValue := decimal.NewFromInt(1)

			metadataModel.Token = &metadata.Token{
				TokenAddress:  nftTransfer.TokenAddress,
				TokenStandard: strings.ToLower(nftTransfer.ContractType),
				TokenID:       &tokenID,
				TokenValue:    &tokenValue, // TODO ERC1155
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

			metadataModel.Token = &metadata.Token{
				TokenAddress:  tokenTransfer.Address,
				TokenStandard: "erc20",
				TokenValue:    &tokenValue,
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

			metadataModel.Token = &metadata.Token{
				TokenStandard: "native",
				TokenValue:    &tokenValue,
			}
		}

		rawMetadata, err := json.Marshal(metadataModel)
		if err != nil {
			return nil, err
		}

		transfer.Metadata = rawMetadata

		internalTransfers = append(internalTransfers, transfer)
	}

	return internalTransfers, nil
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
	}
}
