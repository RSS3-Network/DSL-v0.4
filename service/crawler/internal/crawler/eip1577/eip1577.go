package eip1577

import (
	"bufio"
	"embed"
	"encoding/csv"
	"errors"
	"io"
	"strings"
	"sync"
	"time"
	_ "time/tzdata"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/eip1577"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service/ens"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type service struct {
	employer *shedlock.Employer
	job      crawler.Job
}

var (
	//go:embed eip1577.csv
	DomainFS embed.FS

	_ crawler.Crawler = (*service)(nil)
	_ crawler.Job     = (*EIP1577Job)(nil)

	ensClient *ens.Client
	conf      *config.Config
)

type EIP1577Job struct{}

func (job *EIP1577Job) Name() string {
	return "crawler_eip1577_cron"
}

func (job *EIP1577Job) Spec() string {
	return "CRON_TZ=Asia/Shanghai 0 10 * * *"
}

func (job *EIP1577Job) Timeout() time.Duration {
	return time.Hour * 5
}

func New(c *config.Config, employer *shedlock.Employer) crawler.Crawler {
	ensClient = ens.New()
	conf = c

	return &service{
		employer: employer,
		job:      &EIP1577Job{},
	}
}

func (s *service) Name() string {
	return protocol.NetworkEIP1577
}

func (s *service) Run() error {
	// job
	if err := s.employer.AddJob(s.job.Name(), s.job.Spec(), s.job.Timeout(), crawler.NewCronJob(s.employer, s.job)); err != nil {
		return err
	}

	return nil
}

func (job *EIP1577Job) Run(renewal crawler.RenewalFunc) error {
	loggerx.Global().Info("eip1577: job run")

	ctx := context.Background()
	file, err := DomainFS.Open("eip1577.csv")
	if err != nil {
		loggerx.Global().Error("eip1577: open file error", zap.Error(err))
		return err
	}
	defer file.Close()

	ch := make(chan struct{}, 30)
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
		address, err := ensClient.GetResolvedAddress(domain)
		if err != nil {
			loggerx.Global().Error("eip1577: domain resolve error", zap.Error(err))

			continue
		}

		if len(address) == 0 {
			continue
		}

		ch <- struct{}{}
		wg.Add(1)
		go func(domain string, address string) {
			defer wg.Done()
			for i := 0; i < 2; i++ {
				err := HandleEIP1577(ctx, domain, address)
				if err != nil {
					loggerx.Global().Error("eip1577: HandleEIP1577 error", zap.Error(err))
				}
			}
			<-ch
		}(domain, address)
	}

	wg.Wait()

	return nil
}

func HandleEIP1577(ctx context.Context, domain string, address string) error {
	loggerx.Global().Info("eip1577: start", zap.String("domain", domain), zap.String("address", address))

	c, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer func() {
		cancel()
	}()

	message := &protocol.Message{
		Address: strings.ToLower(address),
		Network: protocol.NetworkEIP1577,
	}

	// datasource
	internalTransactions, err := eip1577.GetEIP1577Transactions(c, message, domain, conf.EIP1577.Endpoint)
	if err != nil {
		loggerx.Global().Error("eip1577: eip1577 Datasource Handle error", zap.Error(err))

		return err
	}

	// insert db
	var transaction []*model.Transaction
	for i := range internalTransactions {
		transaction = append(transaction, &internalTransactions[i])
	}
	err = database.UpsertTransactions(ctx, transaction, false)
	if err != nil {
		loggerx.Global().Error("eip1577: eip1577 database UpsertTransactions error", zap.Error(err))

		return err
	}

	return nil
}
