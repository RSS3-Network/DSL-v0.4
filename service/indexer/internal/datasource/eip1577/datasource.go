package eip1577

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"github.com/shopspring/decimal"
	goens "github.com/wealdtech/go-ens/v3"
	"go.uber.org/zap"
)

type DomainHandleFunc func(ctx context.Context, address common.Address) (string, error)

var _ datasource.Datasource = (*Datasource)(nil)

type Datasource struct {
	employer *shedlock.Employer
}

func (d *Datasource) Name() string {
	return "eip1577"
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkEIP1577,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transaction, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
	}()

	lockerName := d.buildLockerName(message.Address)

	// Maximum index time is 3 minutes. The external indexer of EIP-1577 is not currently an asynchronous system,
	// so it should not be used as a basis for whether an address is successfully indexed, as this is very time intensive.
	if !d.employer.DoLock(lockerName, time.Minute*3) {
		loggerx.Global().Warn("eip1577 datasource is locked", zap.String("address", message.Address))

		return []model.Transaction{}, nil
	}

	defer d.employer.UnLock(lockerName)

	domainHandleFuncs := []DomainHandleFunc{
		d.handleENS,
		// If we need to support more domain services, add a new reverse resolution function here.
	}

	address := common.HexToAddress(message.Address)

	names := parallel.Map(domainHandleFuncs, func(domainHandleFunc DomainHandleFunc, _ int) string {
		name, err := domainHandleFunc(ctx, address)
		if err != nil {
			zap.L().Warn("failed to handel domain", zap.Error(err), zap.String("address", address.String()))

			return ""
		}

		return name
	})

	// parallel.MapWithError will cancel all possible successful requests,
	// but it's not appropriate here, so it will filter the failed requests after all of them have completed.
	names = lo.Filter(names, func(name string, _ int) bool {
		return !strings.EqualFold(name, "")
	})

	transactions := lo.Flatten(parallel.Map(names, func(name string, i int) []model.Transaction {
		transactions, err := d.buildTransactions(ctx, message, name)
		if err != nil {
			loggerx.Global().Warn("failed to build transactions", zap.Error(err))

			return []model.Transaction{}
		}

		return transactions
	}))

	return transactions, nil
}

func (d *Datasource) handleENS(ctx context.Context, address common.Address) (string, error) {
	ethereumClient, err := ethclientx.Global(protocol.NetworkEthereum)
	if err != nil {
		return "", err
	}

	name, err := goens.ReverseResolve(ethereumClient, address)
	if err != nil {
		return "", err
	}

	return name, nil
}

func (d *Datasource) buildTransactions(ctx context.Context, message *protocol.Message, name string) ([]model.Transaction, error) {
	var result Result

	response, err := resty.New().NewRequest().SetContext(ctx).SetResult(&result).Get(fmt.Sprintf("%s/domains/%s", config.ConfigIndexer.EIP1577.Endpoint, name))
	if err != nil {
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
			Title: feed.Title,
			Body:  feed.Summary,
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
			Source:      d.Name(),
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
					Source:          d.Name(),
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

func (d *Datasource) buildLockerName(address string) string {
	return fmt.Sprintf("eip1577:%s@datasource", address)
}

func New(employer *shedlock.Employer) datasource.Datasource {
	return &Datasource{
		employer: employer,
	}
}
