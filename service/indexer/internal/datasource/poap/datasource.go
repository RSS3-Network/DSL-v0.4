package poap

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/poap"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/shopspring/decimal"
)

const (
	Source = "poap"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	poapClient *poap.Client
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkXDAI,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	actions, err := d.poapClient.GetActions(message.Address)
	if err != nil {
		return nil, err
	}

	for _, action := range actions {
		timestamp, err := time.Parse("2006-01-02 15:04:05", action.Created)
		if err != nil {
			return nil, err
		}

		transactionLogIndex, err := decimal.NewFromString(action.TokenID)
		if err != nil {
			return nil, err
		}

		sourceData, err := json.Marshal(action)
		if err != nil {
			return nil, err
		}

		transactionHash := sha256.New()
		transactionHash.Write([]byte(action.TokenID))
		transactionHash.Write([]byte(action.Created))

		internalTransfers = append(internalTransfers, model.Transfer{
			TransactionHash:     hex.EncodeToString(transactionHash.Sum(nil)),
			Timestamp:           timestamp,
			Type:                "poap",
			TransactionLogIndex: transactionLogIndex,
			AddressTo:           action.Owner,
			Metadata:            nil,
			Network:             action.Chain,
			Source:              Source,
			SourceData:          sourceData,
		})
	}

	return internalTransfers, nil
}

func New() datasource.Datasource {
	return &Datasource{
		poapClient: poap.New(),
	}
}
