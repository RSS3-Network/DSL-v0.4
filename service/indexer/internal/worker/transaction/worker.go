package transaction

import (
	"context"
	"embed"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/datasource/blockscout"
	moralisx "github.com/naturalselectionlabs/pregod/common/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/common/datasource/nft"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	zksync2 "github.com/naturalselectionlabs/pregod/common/worker/zksync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	transaction2 "github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction/coinmarketcap"
	lop "github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	Name = "transaction"

	AddressGenesis = "0x0000000000000000000000000000000000000000"
)

var (
	_ worker.Worker = (*service)(nil)

	//go:embed asset/*
	asseFS embed.FS
)

type service struct {
	databaseClient *gorm.DB
	zksyncClient   *zksync2.Client
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

		wallentModels := make([]exchange.CexWallet, 0)

		for i, record := range records {
			if i == 0 {
				continue
			}

			wallentModels = append(wallentModels, exchange.CexWallet{
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
			Model((*exchange.CexWallet)(nil)).
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
	tracer := otel.Tracer("token_worker")
	ctx, handlerSpan := tracer.Start(ctx, "token_worker:Handle")

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
func (s *service) handleCrossbellAndXDAI(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("token_worker")
	_, trace := tracer.Start(ctx, "token_worker:handleCrossbellAndXDAI")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	// nft.GetMetadata start
	bsTransactions := []blockscout.Transaction{}
	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			sourceData := blockscout.Transaction{}
			if err := json.Unmarshal(transfer.SourceData, &sourceData); err != nil {
				return nil, err
			}
			if transfer.Source == "blockscout" && message.Network == protocol.NetworkCrossbell && sourceData.ContractAddress != "" {
				bsTransactions = append(bsTransactions, sourceData)
			}
		}
	}
	nftMetadata := lop.Map(bsTransactions, func(sourceData blockscout.Transaction, i int) metadata.Token {
		nftMetadata, err := nft.GetMetadata(
			message.Network,
			common.HexToAddress(sourceData.ContractAddress),
			sourceData.TokenID.BigInt(),
		)
		if err != nil { // print err but no return
			logrus.Errorf("%s: %v", message.Network, err)
		}
		return metadata.Token{
			TokenAddress:  sourceData.ContractAddress,
			TokenStandard: protocol.TokenStandardERC721,
			TokenID:       &sourceData.TokenID,
			TokenValue:    &sourceData.Value,
			NFTMetadata:   nftMetadata,
		}
	})
	// nft.GetMetadata end

	internalTransactionMap := make(map[string]model.Transaction)
	nftMetadataIndex := 0

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			if transfer.Source == "blockscout" {
				sourceData := blockscout.Transaction{}
				metadataModel := metadata.Metadata{}
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
					md := nftMetadata[nftMetadataIndex]
					nftMetadataIndex += 1
					metadataModel.Token = &md

					transfer.Tag = filter.UpdateTag(filter.TagCollectible, transfer.Tag)
				case message.Network == protocol.NetworkXDAI:
					var coinInfo *transaction2.CoinMarketCapCoinInfo
					var err error
					if sourceData.ContractAddress != "" {
						coinInfo, err = coinmarketcap.CachedGetCoinInfo(ctx, message.Network, sourceData.ContractAddress)
					} else {
						coinInfo, err = coinmarketcap.CachedGetCoinInfoByNetwork(ctx, message.Network, transfer.Index)
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

			transaction, transfer = s.updateType(transaction, transfer)

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

func (s *service) handleEthereum(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("token_worker")
	_, trace := tracer.Start(ctx, "token_worker:handleEthereum")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	// nft.GetMetadata start
	var nftTransfers []model.Transfer
	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			sourceDataMap := make(map[string]interface{})
			if err := json.Unmarshal(transfer.SourceData, &sourceDataMap); err != nil {
				return nil, err
			}
			if _, exist := sourceDataMap["contract_type"]; exist {
				nftTransfers = append(nftTransfers, transfer)
			}
		}
	}
	nftMetadata := lop.Map(nftTransfers, func(transfer model.Transfer, i int) metadata.Token {
		// NFT transfer
		nftTransfer := moralisx.NFTTransfer{}
		if err := json.Unmarshal(transfer.SourceData, &nftTransfer); err != nil {
			return metadata.Token{}
		}

		tokenID, err := decimal.NewFromString(nftTransfer.TokenId)
		if err != nil {
			return metadata.Token{}
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

		return metadata.Token{
			TokenAddress:  nftTransfer.TokenAddress,
			TokenStandard: strings.ToUpper(nftTransfer.ContractType),
			TokenID:       &tokenID,
			TokenValue:    &tokenValue,
			NFTMetadata:   nftMetadata,
		}
	})
	// nft.GetMetadata end

	internalTransactionMap := make(map[string]model.Transaction)
	nftMetadataIndex := 0

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			sourceDataMap := make(map[string]interface{})

			var wallet exchange.CexWallet

			if err := json.Unmarshal(transfer.SourceData, &sourceDataMap); err != nil {
				return nil, err
			}

			var metadataModel metadata.Metadata

			if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
				return nil, err
			}

			if _, exist := sourceDataMap["contract_type"]; exist {
				// NFT transfer
				md := nftMetadata[nftMetadataIndex]
				nftMetadataIndex += 1
				metadataModel.Token = &md

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
				if err := s.checkCexWallet(ctx, message.Address, &transfer, &wallet); err != nil {
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
				coinInfo, err := coinmarketcap.CachedGetCoinInfoByNetwork(ctx, message.Network, transfer.Index)
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
				if err := s.checkCexWallet(ctx, message.Address, &transfer, &wallet); err != nil {
					return nil, err
				}
			}

			rawMetadata, err := json.Marshal(metadataModel)
			if err != nil {
				return nil, err
			}

			transfer.Metadata = rawMetadata

			transaction, transfer = s.updateType(transaction, transfer)

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

func (s *service) handleZkSync(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("token_worker")
	_, trace := tracer.Start(ctx, "token_worker:handleZkSync")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		for _, transfer := range transaction.Transfers {
			data := zksync2.GetTransactionData{}

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
			} else { // transaction
				metadataModel.Token = &metadata.Token{
					TokenAddress:  tokenInfo.Address,
					TokenStandard: protocol.TokenStandardERC20,
					TokenValue:    &amount,
					Decimals:      tokenInfo.Decimals,
					Symbol:        tokenInfo.Symbol,
				}

				transfer.Tag = filter.UpdateTag(filter.TagTransaction, transfer.Tag)
			}

			transaction, transfer = s.updateType(transaction, transfer)

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

func keyOfCheckCexWallet(address string) string {
	return fmt.Sprintf("check_exchange_wallet.%s", address)
}

// Check address (from / to) is a WalletAddress. If true, update transfer
func (s *service) checkCexWallet(ctx context.Context, address string, transfer *model.Transfer, wallet *exchange.CexWallet) error {
	// get from redis cache (to)
	exists, err := cache.GetMsgPack(ctx, keyOfCheckCexWallet(transfer.AddressTo), wallet)
	if err != nil {
		return err
	}
	if !exists { // get from redis cache (from)
		if exists, err = cache.GetMsgPack(ctx, keyOfCheckCexWallet(transfer.AddressFrom), wallet); err != nil {
			return nil
		}
	}

	if !exists {
		err := s.databaseClient.Model((*exchange.CexWallet)(nil)).Where("wallet_address = ?", strings.ToLower(transfer.AddressTo)).Or("wallet_address = ?", strings.ToLower(transfer.AddressFrom)).First(&wallet).Error
		switch {
		case err == nil: // exists, set WalletAddress' cache
			if err := cache.SetMsgPack(ctx, keyOfCheckCexWallet(wallet.WalletAddress), wallet, 7*24*time.Hour); err != nil {
				return err
			}
		case errors.Is(err, gorm.ErrRecordNotFound): // not exists, set `from` and `to` address' cache (empty)
			if err := cache.SetMsgPack(ctx, keyOfCheckCexWallet(transfer.AddressFrom), wallet, 7*24*time.Hour); err != nil {
				return err
			}
			if err := cache.SetMsgPack(ctx, keyOfCheckCexWallet(transfer.AddressTo), wallet, 7*24*time.Hour); err != nil {
				return err
			}
		default: // other err
			return err
		}
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

func (s *service) updateType(transaction model.Transaction, transfer model.Transfer) (model.Transaction, model.Transfer) {
	switch {
	case transfer.AddressFrom == AddressGenesis:
		transfer.Type = filter.TransactionMint
		switch transfer.Tag {
		case filter.TagTransaction:
			transfer.Type = filter.TransactionMint
		case filter.TagCollectible:
			transfer.Type = filter.CollectibleMint
		}
	case transfer.AddressTo == AddressGenesis:
		transfer.Type = filter.TransactionBurn
		switch transfer.Tag {
		case filter.TagTransaction:
			transfer.Type = filter.TransactionBurn
		case filter.TagCollectible:
			transfer.Type = filter.CollectibleBurn
		}
	case transfer.AddressFrom == transfer.AddressTo:
		transfer.Type = filter.TransactionSelf
	case transfer.Tag == filter.TagExchange:
		transaction.Platform = transfer.Platform
	default:
		transfer.Type = filter.TransactionTransfer
		switch transfer.Tag {
		case filter.TagTransaction:
			transfer.Type = filter.TransactionTransfer
		case filter.TagCollectible:
			transfer.Type = filter.CollectibleTrade
		}
	}

	return transaction, transfer
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		zksyncClient:   zksync2.New(),
	}
}