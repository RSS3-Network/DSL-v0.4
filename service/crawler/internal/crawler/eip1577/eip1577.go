package eip1577

import (
	"bufio"
	"embed"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/eip1577"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var (
	//go:embed eip1577.csv
	DomainFS embed.FS

	_ crawler.Crawler = (*service)(nil)
)

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
	file, err := os.Open("./service/crawler/internal/crawler/eip1577/eip1577.csv")
	if err != nil {
		loggerx.Global().Error("eip1577: open file error", zap.Error(err))
		return err
	}
	defer file.Close()

	ch := make(chan struct{}, 5)
	reader := csv.NewReader(bufio.NewReader(file))

	var wg sync.WaitGroup
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

		fmt.Println(address)

		ch <- struct{}{}
		wg.Add(1)
		go func(domain string, address string) {
			defer wg.Done()
			err := s.HandleEIP1577(ctx, domain, address)
			if err != nil {
				loggerx.Global().Error("eip1577: HandleEIP1577 error", zap.Error(err))
			}
			<-ch
		}(domain, address)
	}

	wg.Wait()

	return nil
}

func (s *service) HandleEIP1577(ctx context.Context, domain string, address string) error {
	loggerx.Global().Info("eip1577: start ", zap.String("domain", domain), zap.String("address", address))

	message := &protocol.Message{
		Address: strings.ToLower(address),
		Network: protocol.NetworkEIP1577,
	}

	// datasource
	internalTransactions, err := eip1577.GetEIP1577Transactions(ctx, message, domain, s.config.EIP1577.Endpoint)
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
