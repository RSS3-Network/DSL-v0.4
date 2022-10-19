package eip1577

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func GetEIP1577Transactions(ctx context.Context, message *protocol.Message, name string, endpoint string) ([]model.Transaction, error) {
	var result Result

	url := fmt.Sprintf("%s/domains/%s", endpoint, name)
	response, err := resty.New().NewRequest().SetContext(ctx).SetResult(&result).Get(url)
	if err != nil {
		loggerx.Global().Error("eip1577 api error", zap.Error(err))
		return nil, err
	}

	if response.StatusCode() >= http.StatusBadRequest && response.StatusCode() < http.StatusNetworkAuthenticationRequired {
		loggerx.Global().Warn("eip1577 notes are not found", zap.String("address", message.Address), zap.String("status", response.Status()))

		return []model.Transaction{}, nil
	}

	transactions := make([]model.Transaction, 0, len(result.Feeds))

	for _, feed := range result.Feeds {
		// Use domain name and path as hash to cope with page content updates
		hash := common.Hash(sha256.Sum256([]byte(fmt.Sprintf("%s%s", name, feed.Path)))).String()
		success := true

		metadataRaw, err := json.Marshal(metadata.Post{
			Title:   feed.Title,
			Summary: feed.Summary,
			Body:    feed.Content,
		})
		if err != nil {
			return nil, err
		}

		sourceData, err := json.Marshal(feed)
		if err != nil {
			return nil, err
		}

		transaction := model.Transaction{
			BlockNumber: 0,
			Timestamp:   feed.CreatedAt,
			Hash:        hash,
			Index:       0,
			Owner:       message.Address,
			AddressFrom: message.Address,
			Network:     protocol.NetworkEIP1577,
			Platform:    result.Platform,
			Source:      protocol.NetworkEIP1577,
			Tag:         filter.TagSocial,
			Type:        filter.SocialPost,
			Success:     &success,
			Transfers: []model.Transfer{
				{
					TransactionHash: hash,
					Timestamp:       time.Time{},
					BlockNumber:     decimal.Zero.BigInt(),
					Tag:             filter.TagSocial,
					Type:            filter.SocialPost,
					Index:           0,
					AddressFrom:     message.Address,
					Metadata:        metadataRaw,
					Network:         protocol.NetworkEIP1577,
					Platform:        result.Platform,
					Source:          protocol.NetworkEIP1577,
					SourceData:      sourceData,
					RelatedUrls: []string{
						feed.URL,
					},
				},
			},
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
