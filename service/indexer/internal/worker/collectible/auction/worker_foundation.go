package auction

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/foundation"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func (i *internal) handleFoundationTransactions(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_auction")
	_, span := tracer.Start(ctx, "worker_auction:handleFoundationTransactions")

	transactions := make([]model.Transaction, 0)
	transactionsMap := make(map[common.Hash]*model.Transaction, 0)

	defer opentelemetry.Log(span, transactions, nil, nil)

	foundationQuery := kurora.DatasetFoundationEventQuery{
		From: lo.ToPtr(common.HexToAddress(message.Address)),
		Type: lo.ToPtr(0), // type action
	}

	if message.BlockNumber > 0 {
		foundationQuery.BlockNumberFrom = lo.ToPtr(uint64(message.BlockNumber))
	}

	for first := true; foundationQuery.Cursor != nil || first; first = false {
		events, err := i.kuroraClient.FetchDatasetFoundationEvents(ctx, foundationQuery)
		if err != nil {
			return nil, fmt.Errorf("nouns transactions: %w", err)
		}

		for _, event := range events {
			nft, err := i.tokenClient.NFTToMetadata(ctx, message.Network, event.NftAddress.String(), event.NftId.BigInt())
			if err != nil {
				zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", event.TransactionHash.String()), zap.Stringer("contract_address", event.NftAddress), zap.Stringer("token_id", event.NftId))

				return nil, fmt.Errorf("nft to metadata: %w", err)
			}

			nft.Action = event.EventType
			nft.StartTime = &event.Timestamp
			if !event.Expired.IsZero() {
				nft.EndTime = &event.Expired
			}

			nativeToken, err := i.tokenClient.Native(context.Background(), message.Network)
			if err != nil {
				return nil, fmt.Errorf("get native token: %w", err)
			}

			costValue := decimal.NewFromBigInt(event.CreatorAmountInETH.BigInt().Add(event.FeeInETH.BigInt(), event.OwnerInETH.BigInt()), 0)
			costValueDisplay := costValue.Shift(-int32(nativeToken.Decimals))

			cost := &metadata.Token{
				Name:         nativeToken.Name,
				Symbol:       nativeToken.Symbol,
				Decimals:     nativeToken.Decimals,
				Standard:     protocol.TokenStandardNative,
				Image:        nativeToken.Logo,
				Value:        &costValue,
				ValueDisplay: &costValueDisplay,
			}

			nft.Cost = cost

			metadataRaw, err := json.Marshal(nft)
			if err != nil {
				return nil, fmt.Errorf("marshal metadata: %w", err)
			}

			var from, to string

			switch event.EventType {
			case filter.CollectibleAuctionFinalize:
				from, to = strings.ToLower(event.Seller.String()), strings.ToLower(event.From.String())
			case filter.CollectibleAuctionBid:
				from, to = strings.ToLower(event.From.String()), strings.ToLower(event.Seller.String())
			case filter.CollectibleAuctionCreate, filter.CollectibleAuctionCancel, filter.CollectibleAuctionUpdate, filter.CollectibleAuctionInvalidate:
				from, to = strings.ToLower(event.From.String()), strings.ToLower(foundation.AddressFoundationMarket.String())
			}

			transfer := model.Transfer{
				TransactionHash: strings.ToLower(event.TransactionHash.String()),
				Timestamp:       event.Timestamp,
				Index:           int64(event.LogIndex),
				AddressFrom:     from,
				AddressTo:       to,
				Metadata:        metadataRaw,
				Network:         message.Network,
				Platform:        protocol.PlatformFoundation,
				Source:          protocol.SourceKurora,
				Tag:             filter.TagCollectible,
				Type:            filter.CollectibleAuction,
				RelatedUrls: ethereum.BuildURL(
					[]string{utils.GetTxHashURL(message.Network, event.TransactionHash.String()), fmt.Sprintf("%s/%s/%s/%v", "https://foundation.app/", event.Seller.String(), event.NftAddress.String(), event.NftId)},
					ethereum.BuildTokenURL(message.Network, event.NftAddress.String(), event.NftId.String())...,
				),
			}

			if _, ok := transactionsMap[event.TransactionHash]; !ok {
				// append new tx
				transaction := &model.Transaction{
					BlockNumber: int64(event.BlockNumber),
					Timestamp:   event.Timestamp,
					Owner:       strings.ToLower(event.From.String()),
					Hash:        strings.ToLower(event.TransactionHash.String()),
					AddressFrom: strings.ToLower(event.From.String()),
					AddressTo:   strings.ToLower(foundation.AddressFoundationMarket.String()),
					Network:     message.Network,
					Success:     lo.ToPtr(true),
					Source:      protocol.SourceKurora,
					Platform:    protocol.PlatformFoundation,
					Tag:         filter.TagCollectible,
					Type:        filter.CollectibleAuction,
					Transfers: []model.Transfer{
						// This is a real transfer
						transfer,
					},
				}

				transactionsMap[event.TransactionHash] = transaction

			} else {
				// append transfers
				transactionsMap[event.TransactionHash].Transfers = append(transactionsMap[event.TransactionHash].Transfers, transfer)
			}
		}

		if len(events) == 0 {
			break
		}

		lastEvent, _ := lo.Last(events)
		foundationQuery.Cursor = lo.ToPtr(kurora.LogCursor(lastEvent.TransactionHash, lastEvent.LogIndex))
	}

	for _, transaction := range transactionsMap {
		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}
