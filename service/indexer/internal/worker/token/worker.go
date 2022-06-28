package token

import (
	"context"
	"embed"
	"encoding/csv"
	"encoding/json"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/blockscout"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	moralisx "github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/nft"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token/coinmarketcap"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	EmptyAddress = "0x0000000000000000000000000000000000000000"
	Name         = "token"
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
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
		protocol.NetworkBinanceSmartChain,
		protocol.NetworkZkSync,
		protocol.NetworkCrossbell,
		protocol.NetworkXDAI,
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

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_token")

	ctx, handlerSpan := tracer.Start(ctx, "handler")

	defer handlerSpan.End()

	ctx, databaseSpan := tracer.Start(ctx, "database") //nolint:ineffassign,staticcheck

	databaseSpan.End()

	switch message.Network {
	case protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain:
		return s.handleEthereum(ctx, message, transactions)
	case protocol.NetworkZkSync:
		return s.handleZkSync(ctx, message, transactions)
	case protocol.NetworkCrossbell, protocol.NetworkXDAI:
		return s.handleCrossbellAndXDAI(ctx, message, transactions)
	}

	return []model.Transaction{}, nil
}

// handle crossbell / XDAI NFT (link list / profile)
func (s *service) handleCrossbellAndXDAI(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			if transfer.Source == "blockscout" {
				sourceData := blockscout.Transaction{}
				metadataModel := metadata.Metadata{}
				if err := json.Unmarshal(transfer.SourceData, &sourceData); err != nil {
					return nil, err
				}
				if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
					return nil, err
				}

				switch {
				case message.Network == protocol.NetworkCrossbell && sourceData.ContractAddress == "":
					// crossbell empty: use CSB
					metadataModel.Token = &metadata.Token{
						TokenStandard: protocol.TokenStandardERC20,
						Logo:          "https://scan.crossbell.io/images/csb-yellow-no-bg.svg",
						Name:          "Crossbell",
						Symbol:        "CSB",
						Decimals:      18,
					}

					transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
				case message.Network == protocol.NetworkCrossbell && sourceData.ContractAddress != "":
					nftMetadata, err := nft.GetMetadata(
						message.Network,
						common.HexToAddress(sourceData.ContractAddress),
						sourceData.TokenID.BigInt(),
					)
					if err != nil { // print err but no return
						logrus.Errorf("%s: %v", message.Network, err)
					}

					metadataModel.Token = &metadata.Token{
						TokenAddress:  sourceData.ContractAddress,
						TokenStandard: protocol.TokenStandardERC721,
						TokenID:       &sourceData.TokenID,
						TokenValue:    &sourceData.Value,
						NFTMetadata:   nftMetadata,
					}

					transfer.Tag = filter.UpdateTag(filter.TagCollectible, transfer.Tag)
				case message.Network == protocol.NetworkXDAI:
					var coinInfo *model.CoinMarketCapCoinInfo
					var err error
					if sourceData.ContractAddress != "" {
						coinInfo, err = coinmarketcap.CachedGetCoinInfo(ctx, message.Network, sourceData.ContractAddress)
					} else {
						coinInfo, err = coinmarketcap.CachedGetCoinInfoByNetwork(ctx, message.Network)
					}
					if err != nil {
						logrus.Error(err)
					} else {
						metadataModel.Token = &metadata.Token{
							TokenAddress:  sourceData.ContractAddress,
							TokenStandard: protocol.TokenStandardERC20,
							TokenValue:    &sourceData.Value,
							Logo:          coinInfo.Logo,
							Name:          coinInfo.Name,
							Symbol:        coinInfo.Symbol,
							Decimals:      coinInfo.Decimals,
						}

						transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
					}
				}

				rawMetadata, err := json.Marshal(metadataModel)
				if err != nil {
					return nil, err
				}
				transfer.Metadata = rawMetadata
			}

			switch {
			case transfer.AddressFrom == EmptyAddress:
				transfer.Type = filter.TransactionMint
				if transfer.Tag == filter.TagCollectible {
					transfer.Type = filter.CollectibleMint
				}
			case transfer.AddressTo == EmptyAddress:
				transfer.Type = filter.TransactionBurn
				if transfer.Tag == filter.TagCollectible {
					transfer.Type = filter.CollectibleBurn
				}
			default:
				transfer.Type = filter.TransactionTransfer
				if transfer.Tag == filter.TagCollectible {
					transfer.Type = filter.CollectibleTrade
				}
			}

			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Tag = filter.UpdateTag(transfer.Tag, value.Tag)
			value.Transfers = append(value.Transfers, transfer)

			internalTransactionMap[transaction.Hash] = value

			// transaction tag
			transaction.Tag = filter.UpdateTag(transfer.Tag, transaction.Tag)
		}
	}

	internalTransactions := make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, internalTransaction)
	}

	return internalTransactions, nil
}

func (s *service) handleEthereum(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			sourceDataMap := make(map[string]interface{})

			var wallet model.ExchangeWallet

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
					TokenStandard: strings.ToUpper(nftTransfer.ContractType),
					TokenID:       &tokenID,
					TokenValue:    &tokenValue,
					NFTMetadata:   nftMetadata,
				}

				transfer.Tag = filter.UpdateTag(filter.TagCollectible, transfer.Tag)
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
					TokenStandard: protocol.TokenStandardERC20,
					TokenValue:    &tokenValue,
					Logo:          coinInfo.Logo,
					Name:          coinInfo.Name,
					Symbol:        coinInfo.Symbol,
					Decimals:      coinInfo.Decimals,
				}

				transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)

				// check for exchange transaction
				if err := s.checkExchangeWallet(strings.ToLower(message.Address), &transfer, &wallet); err != nil {
					return nil, err
				}
			} else {
				// Native transfer
				nativeTransfer := moralisx.Transaction{}
				if err := json.Unmarshal(transfer.SourceData, &nativeTransfer); err != nil {
					return nil, err
				}

				if len(nativeTransfer.Value) == 0 {
					nativeTransfer.Value = "18"
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
					TokenStandard: protocol.TokenStandardNative,
					TokenValue:    &tokenValue,
					Logo:          coinInfo.Logo,
					Name:          coinInfo.Name,
					Symbol:        coinInfo.Symbol,
					Decimals:      coinInfo.Decimals,
				}

				transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)

				// check for exchange transaction
				if err := s.checkExchangeWallet(message.Address, &transfer, &wallet); err != nil {
					return nil, err
				}
			}

			rawMetadata, err := json.Marshal(metadataModel)
			if err != nil {
				return nil, err
			}

			transfer.Metadata = rawMetadata

			switch {
			case transfer.AddressFrom == EmptyAddress:
				transfer.Type = filter.TransactionMint
				if transfer.Tag == filter.TagCollectible {
					transfer.Type = filter.CollectibleMint
				}
			case transfer.AddressTo == EmptyAddress:
				transfer.Type = filter.TransactionBurn
				if transfer.Tag == filter.TagCollectible {
					transfer.Type = filter.CollectibleBurn
				}
			case transfer.Tag == filter.TagExchange:
				transaction.Platform = transfer.Platform
			default:
				transfer.Type = filter.TransactionTransfer
				if transfer.Tag == filter.TagCollectible {
					transfer.Type = filter.CollectibleTrade
				}
			}

			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction
				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Tag = filter.UpdateTag(transfer.Tag, value.Tag)
			value.Transfers = append(value.Transfers, transfer)

			internalTransactionMap[transaction.Hash] = value

			// transaction tag
			transaction.Tag = filter.UpdateTag(transfer.Tag, transaction.Tag)
		}
	}

	internalTransactions := make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, internalTransaction)
	}

	return internalTransactions, nil
}

func (s *service) handleZkSync(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
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
					TokenStandard: protocol.TokenStandardERC721,
					TokenID:       &tokenID,
					TokenValue:    &amount,
					Symbol:        nftTokenInfo.Symbol,
					NFTMetadata:   nftTokenInfo.Bytes(),
				}

				transfer.Tag = filter.UpdateTag(filter.TagCollectible, transfer.Tag)
			} else { // token
				metadataModel.Token = &metadata.Token{
					TokenAddress:  tokenInfo.Address,
					TokenStandard: protocol.TokenStandardERC20,
					TokenValue:    &amount,
					Decimals:      tokenInfo.Decimals,
					Symbol:        tokenInfo.Symbol,
				}

				transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
			}
			switch {
			case transfer.AddressFrom == EmptyAddress:
				transfer.Type = filter.TransactionMint
				if transfer.Tag == filter.TagCollectible {
					transfer.Type = filter.CollectibleMint
				}
			case transfer.AddressTo == EmptyAddress:
				transfer.Type = filter.TransactionBurn
				if transfer.Tag == filter.TagCollectible {
					transfer.Type = filter.CollectibleBurn
				}
			default:
				transfer.Type = filter.TransactionTransfer
				if transfer.Tag == filter.TagCollectible {
					transfer.Type = filter.CollectibleTrade
				}
			}

			rawMetadata, err := json.Marshal(metadataModel)
			if err != nil {
				return nil, err
			}

			transfer.Metadata = rawMetadata

			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Tag = filter.UpdateTag(transfer.Tag, value.Tag)
			value.Transfers = append(value.Transfers, transfer)
			internalTransactionMap[transaction.Hash] = value

			// transaction tag
			transaction.Tag = filter.UpdateTag(transfer.Tag, transaction.Tag)
		}
	}

	internalTransactions := make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, internalTransaction)
	}

	return internalTransactions, nil
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

func (s *service) checkExchangeWallet(address string, transfer *model.Transfer, wallet *model.ExchangeWallet) error {
	if err := s.databaseClient.Model((*model.ExchangeWallet)(nil)).
		Where("wallet_address = ?", strings.ToLower(transfer.AddressTo)).Or("wallet_address = ?", strings.ToLower(transfer.AddressFrom)).First(&wallet).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if wallet.Name != "" {
		transfer.Tag = filter.UpdateTag(filter.TagExchange, transfer.Tag)
		transfer.Platform = wallet.Name

		if transfer.AddressTo == address {
			transfer.Type = filter.ExchangeWithdraw
		} else if transfer.AddressFrom == address {
			transfer.Type = filter.ExchangeDeposit
		}
	}

	return nil
}
