package worker

import (
	"context"
	"time"

	"github.com/naturalselectionlabs/pregod/common/shedlock"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/internalModel"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/collectible/marketplace"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/collectible/poap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/donation/gitcoin"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/exchange/liquidity"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/exchange/swap"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/governance/snapshot"
	lens_worker "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/lens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/mirror"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction"

	"go.uber.org/zap"
)

var Workers = []internalModel.Worker{}

func Initialize() error {
	swapWorker, err := swap.New(config.ConfigIndexer.RPC)
	if err != nil {
		return err
	}

	Workers = []internalModel.Worker{
		liquidity.New(ethereumClientMap),
		swapWorker,
		marketplace.New(ethereumClientMap),
		poap.New(ethereumClientMap),
		mirror.New(),
		gitcoin.New(ethereumClientMap),
		snapshot.New(),
		lens_worker.New(ethereumClientMap),
		transaction.New(ethereumClientMap),
	}

	for _, internalWorker := range Workers {
		loggerx.Global().Info("start initializing worker", zap.String("worker", internalWorker.Name()))

		startTime := time.Now()

		if err := internalWorker.Initialize(context.Background()); err != nil {
			return err
		}

		loggerx.Global().Info("initialize worker completion", zap.String("worker", internalWorker.Name()), zap.Duration("duration", time.Since(startTime)))

		if internalWorker.Jobs() == nil {
			continue
		}

		for _, job := range internalWorker.Jobs() {
			if err := shedlock.Global().AddJob(job.Name(), job.Spec(), job.Timeout(), NewCronJob(job)); err != nil {
				return err
			}
		}
	}

}
