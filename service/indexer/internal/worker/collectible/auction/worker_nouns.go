package auction

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func (i *internal) handleNounsTransactions(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	tracer := otel.Tracer("worker_auction")
	_, span := tracer.Start(ctx, "worker_auction:handleNounsTransactions")

	transactions := make([]model.Transaction, 0)
	transactionsMap := make(map[common.Hash]*model.Transaction, 0)

	defer opentelemetry.Log(span, transactions, nil, nil)

	nounsQuery := kurora.DatasetNounsEventQuery{
		Owner: lo.ToPtr(common.HexToAddress(message.Address)),
	}

	if message.BlockNumber > 0 {
		nounsQuery.BlockNumberFrom = lo.ToPtr(uint64(message.BlockNumber))
	}

	for first := true; nounsQuery.Cursor != nil || first; first = false {
		events, err := i.kuroraClient.FetchDatasetNounsEvents(ctx, nounsQuery)
		if err != nil {
			return nil, fmt.Errorf("nouns transactions: %w", err)
		}

		var (
			limiter = make(chan struct{}, ThreadSize)
			nftMap  = make(map[string]*metadata.Token, 0)

			waitGroup sync.WaitGroup
			locker    sync.Mutex
		)

		for _, event := range events {
			limiter <- struct{}{}
			waitGroup.Add(1)

			go func(nftKey, address string, tokenId *big.Int) {
				defer func() {
					<-limiter
					waitGroup.Done()
				}()

				nft, err := i.tokenClient.NFTToMetadata(ctx, message.Network, address, tokenId)
				if err != nil {
					zap.L().Error("nft to metadata", zap.Error(err), zap.String("transaction_hash", nftKey), zap.String("contract_address", address), zap.Stringer("token_id", tokenId))

					return
				}

				locker.Lock()
				nftMap[nftKey] = nft
				locker.Unlock()
			}(fmt.Sprintf("%v-%v", event.TransactionHash.String(), event.LogIndex), event.NftAddress.String(), event.NftId.BigInt())
		}

		waitGroup.Wait()

		for _, event := range events {
			if _, ok := nftMap[fmt.Sprintf("%v-%v", event.TransactionHash.String(), event.LogIndex)]; !ok {
				continue
			}

			nft := nftMap[fmt.Sprintf("%v-%v", event.TransactionHash.String(), event.LogIndex)]

			nft.Action = event.EventType
			nft.StartTime = &event.StartTime
			nft.EndTime = &event.EndTime

			nativeToken, err := i.tokenClient.Native(context.Background(), message.Network)
			if err != nil {
				return nil, fmt.Errorf("get native token: %w", err)
			}

			costValue := decimal.NewFromBigInt(event.Amount.BigInt(), 0)
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
			case filter.CollectibleAuctionBid:
				from, to = strings.ToLower(event.From.String()), strings.ToLower(event.To.String())
			case filter.CollectibleAuctionFinalize:
				from, to = strings.ToLower(event.To.String()), strings.ToLower(event.From.String())
			}

			transfer := model.Transfer{
				TransactionHash: strings.ToLower(event.TransactionHash.String()),
				Timestamp:       event.Timestamp,
				Index:           int64(event.LogIndex),
				AddressFrom:     from,
				AddressTo:       to,
				Metadata:        metadataRaw,
				Network:         message.Network,
				Platform:        protocol.PlatformNouns,
				Source:          protocol.SourceKurora,
				Tag:             filter.TagCollectible,
				Type:            filter.CollectibleAuction,
				RelatedUrls: ethereum.BuildURL(
					[]string{utils.GetTxHashURL(message.Network, event.TransactionHash.String()), fmt.Sprintf("%s%v", "https://nouns.wtf/noun/", event.NftId)},
					ethereum.BuildTokenURL(message.Network, event.NftAddress.String(), event.NftId.String())...,
				),
			}

			if _, ok := transactionsMap[event.TransactionHash]; !ok {
				// append new tx
				transaction := &model.Transaction{
					BlockNumber: int64(event.BlockNumber),
					Timestamp:   event.Timestamp,
					Owner:       strings.ToLower(event.Owner.String()),
					Hash:        strings.ToLower(event.TransactionHash.String()),
					AddressFrom: strings.ToLower(event.From.String()),
					AddressTo:   strings.ToLower(event.To.String()),
					Network:     message.Network,
					Success:     lo.ToPtr(true),
					Source:      protocol.SourceKurora,
					Platform:    protocol.PlatformNouns,
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
		nounsQuery.Cursor = lo.ToPtr(kurora.LogCursor(lastEvent.TransactionHash, lastEvent.LogIndex))
	}

	for _, transaction := range transactionsMap {
		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}
