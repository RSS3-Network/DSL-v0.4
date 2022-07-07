package swap

import (
	"context"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	exchange2 "github.com/naturalselectionlabs/pregod/common/worker/exchange"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ worker.Job = (*Job)(nil)

type Job struct {
	databaseClient *gorm.DB
}

func (j *Job) Name() string {
	return "Update Swap Pools"
}

func (j *Job) Spec() string {
	return "0 0 * * *"
}

func (j *Job) Timeout() time.Duration {
	return time.Minute * 5
}

func (j *Job) Run(renewal worker.RenewalFunc) error {
	poolClient := exchange2.NewClient()
	lop.ForEach(exchange2.SwapProviders, func(provider exchange2.SwapProvider, k int) {
		lop.ForEach(provider.SwapPools, func(pool exchange2.SwapPool, i int) {
			result, err := poolClient.GetSwapPairs(context.Background(), provider.Name, pool)
			if err != nil {
				logrus.Errorln(provider.Name, err)
			} else {
				pools := make([]exchange.SwapPool, 0)

				for _, pair := range result {
					pools = append(pools, exchange.SwapPool{
						Source:          provider.Name,
						Network:         pool.Network,
						ContractAddress: string(pair.ID),
						Protocol:        pool.Protocol,
						Token0:          string(pair.Token0.Symbol),
						Token1:          string(pair.Token1.Symbol),
					})
				}

				if len(pools) > 0 {
					if err := j.databaseClient.
						Model((*exchange.SwapPool)(nil)).
						Clauses(clause.OnConflict{
							DoNothing: true,
						}).
						Create(pools).
						Error; err != nil {
						logrus.Errorln(err)
					}
				}
			}
		})
	})

	return nil
}
