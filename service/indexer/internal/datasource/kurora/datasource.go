package kurora

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/kurora/constant"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
)

var _ datasource.Datasource = (*Datasource)(nil)

type Datasource struct {
	kuroraClient *kurora.Client
}

func (d *Datasource) Name() string {
	return protocol.SourceKurora
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkCrossbell,
		protocol.NetworkArbitrum,
		protocol.NetworkOptimism,
		protocol.NetworkAvalanche,
		protocol.NetworkFantom,
		protocol.NetworkCelo,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (transactions []model.Transaction, err error) {
	ctx, trace := otel.Tracer("datasource_kurora").Start(ctx, "datasource_kurora:Handle")
	defer opentelemetry.Log(trace, message, transactions, err)

	unindexedTransactions := make([]*model.Transaction, 0)

	ethereumReceiptQuery := kurora.EthereumReceiptQuery{
		From:  lo.ToPtr(common.HexToAddress(message.Address)),
		Order: lo.ToPtr("desc"),
	}

	// Increase index by timestamp and block number
	if !message.Timestamp.IsZero() {
		ethereumReceiptQuery.TimestampTo = lo.ToPtr(message.Timestamp)
	}

	if message.BlockNumber > 0 {
		ethereumReceiptQuery.BlockNumberTo = lo.ToPtr(decimal.NewFromInt(message.BlockNumber))
	}

	// When preparing to support the Binance Smart Chain and Gnosis in the future,
	// there will probably be some errors here.
	// Binance Smart Chain network should be `binance_smart_chain` instead of `binance`,
	// and Gnosis is `gnosis` instead of `xdai`.
	kuroraNetwork := message.Network

	if kuroraNetwork == protocol.NetworkArbitrum {
		kuroraNetwork = constant.NetworkArbitrumOne.String()
	}

	network, err := constant.NetworkString(kuroraNetwork)
	if err != nil {
		return nil, fmt.Errorf("invalid network: %w", err)
	}

	for first := true; ethereumReceiptQuery.Cursor != nil || first; first = false {
		// The receipts are obtained here instead of the transactions,
		// because the receipts do not require additional requests to get the status of the transactions.
		// Only fetch transaction receipts where the `receipt.from` is `message.Address`.
		receipts, err := d.kuroraClient.FetchEthereumReceipts(ctx, network, ethereumReceiptQuery)
		if err != nil {
			return nil, fmt.Errorf("ethereum transactions: %w", err)
		}

		for _, receipt := range receipts {
			var addressTo string

			if receipt.To != nil {
				addressTo = strings.ToLower(receipt.To.String())
			}

			unindexedTransactions = append(unindexedTransactions, &model.Transaction{
				BlockNumber: receipt.BlockNumber.IntPart(),
				Timestamp:   receipt.Timestamp,
				Hash:        receipt.TransactionHash.String(),
				AddressFrom: strings.ToLower(receipt.From.String()),
				AddressTo:   addressTo,
				Network:     message.Network,
				Success:     lo.ToPtr(receipt.Status == types.ReceiptStatusSuccessful),
				Source:      d.Name(),
				Transfers: []model.Transfer{
					// This is a virtual transfer
					{
						TransactionHash: receipt.TransactionHash.String(),
						Timestamp:       receipt.Timestamp,
						Index:           protocol.IndexVirtual,
						AddressFrom:     strings.ToLower(receipt.From.String()),
						AddressTo:       addressTo,
						Metadata:        metadata.Default,
						Network:         message.Network,
						Source:          d.Name(),
						RelatedUrls: []string{
							utils.GetTxHashURL(message.Network, receipt.TransactionHash.String()),
						},
					},
				},
			})

			if len(unindexedTransactions) >= datasource.DatasourceLimit {
				break
			}
		}

		// Update cursor
		if len(receipts) == 0 {
			break
		}

		lastTransaction, _ := lo.Last(receipts)
		ethereumReceiptQuery.Cursor = lo.ToPtr(lastTransaction.TransactionHash.String())
	}

	return lo.Map(unindexedTransactions, func(transaction *model.Transaction, _ int) model.Transaction {
		return *transaction
	}), nil
}

func New(ctx context.Context, endpoint string) (datasource.Datasource, error) {
	kuroraClient, err := kurora.Dial(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("dial kurora: %w", err)
	}

	return &Datasource{
		kuroraClient: kuroraClient,
	}, nil
}
