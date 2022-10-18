package eip1577

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/eip1577"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	lop "github.com/samber/lo/parallel"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var _ crawler.Crawler = (*service)(nil)

type service struct {
	config    *config.Config
	ensClient *ens.Client
}

func New(config *config.Config) crawler.Crawler {
	return &service{
		config:    config,
		ensClient: ens.New(),
	}
}

func (s *service) Name() string {
	return protocol.NetworkEIP1577
}

func (s *service) Run() error {
	ctx := context.Background()
	domainList, err := s.GetEIP1577DomainList(ctx)
	if err != nil {
		loggerx.Global().Error("eip1577: GetEIP1577AddressList error", zap.Error(err))
		return err
	}

	opt := lop.NewOption().WithConcurrency(20)
	lop.ForEach(domainList, func(domain model.Domain, i int) {
		if err := s.HandleEIP1577(ctx, domain); err != nil {
			loggerx.Global().Error("eip1577: HandleEIP1577 error", zap.Error(err))
			return
		}
	}, opt)

	return nil
}

func (s *service) GetEIP1577DomainList(ctx context.Context) ([]model.Domain, error) {
	file, err := os.Open("./service/crawler/internal/crawler/eip1577/eip1577.csv")
	if err != nil {
		loggerx.Global().Error("eip1577: open file error", zap.Error(err))
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	domainList := []model.Domain{}
	for {
		row, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			loggerx.Global().Error("eip1577: read line error", zap.Error(err))
			continue
		}

		// domain, uri
		domain := row[0]

		// resolve
		address, err := s.ensClient.GetResolvedAddress(domain)
		if err != nil {
			loggerx.Global().Error("eip1577: domain resolve error", zap.Error(err))

			continue
		}

		if len(address) == 0 {
			continue
		}

		domainList = append(domainList, model.Domain{
			Name:         domain,
			AddressOwner: []byte(address),
		})
	}

	return domainList, nil
}

func (s *service) HandleEIP1577(ctx context.Context, domain model.Domain) error {
	message := &protocol.Message{
		Address: strings.ToLower(string(domain.AddressOwner)),
		Network: protocol.NetworkEIP1577,
	}

	// datasource
	internalTransactions, err := eip1577.GetEIP1577Transactions(ctx, message, domain.Name, s.config.EIP1577.Endpoint)
	if err != nil {
		loggerx.Global().Error("eip1577: eip1577 Datasource Handle error", zap.Error(err))

		return err
	}

	// insert db
	transaction := []*model.Transaction{}
	for i := range internalTransactions {
		transaction = append(transaction, &internalTransactions[i])
	}
	err = database.UpsertTransactions(ctx, transaction)
	if err != nil {
		loggerx.Global().Error("eip1577: eip1577 database UpsertTransactions error", zap.Error(err))

		return err
	}

	return nil
}
