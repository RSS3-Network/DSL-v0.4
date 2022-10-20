package eip1577

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/eip1577"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	goens "github.com/wealdtech/go-ens/v3"
	"go.uber.org/zap"
)

type DomainHandleFunc func(ctx context.Context, address common.Address) (string, error)

var _ datasource.Datasource = (*Datasource)(nil)

type Datasource struct {
	employer *shedlock.Employer
}

func (d *Datasource) Name() string {
	return protocol.NetworkEIP1577
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
		transactions, err := eip1577.GetEIP1577Transactions(ctx,
			message,
			name,
			config.ConfigIndexer.EIP1577.Endpoint)
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

func (d *Datasource) buildLockerName(address string) string {
	return fmt.Sprintf("eip1577:%s@datasource", address)
}

func New(employer *shedlock.Employer) datasource.Datasource {
	return &Datasource{
		employer: employer,
	}
}
