package swap

import (
	"context"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/dexpools"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"

	lop "github.com/samber/lo/parallel"
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

func (j *Job) Run() {
	poolClient := dexpools.NewClient()

	lop.ForEach(dexpools.SwapPools, func(dex dexpools.SwapPool, i int) {
		result, err := poolClient.GetSwapPools(context.Background(), dex)
		if err != nil {
			logrus.Errorln(dex.Name, err)
		} else {
			pools := make([]model.SwapPool, 0)

			for _, pair := range result {
				pools = append(pools, model.SwapPool{
					Source:          dex.Name,
					Network:         dex.Network,
					ContractAddress: string(pair.ID),
					Protocol:        dex.Protocol,
					Token0:          string(pair.Token0.Symbol),
					Token1:          string(pair.Token1.Symbol),
				})
			}

			if len(pools) > 0 {
				if err := j.databaseClient.
					Model((*model.SwapPool)(nil)).
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

	logrus.Infoln("Job: Swap Pools finished")
}
