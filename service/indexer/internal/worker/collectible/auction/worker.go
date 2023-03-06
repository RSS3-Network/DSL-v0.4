package auction

import (
	"context"
	"sync"

	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"

	"go.uber.org/zap"
)

var _ worker.Worker = (*internal)(nil)

const (
	SourceName = "auction"
	ThreadSize = 200
)

type internal struct {
	tokenClient  *token.Client
	kuroraClient *kurora.Client
}

func (i *internal) Name() string {
	return SourceName
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
	}
}

func (i *internal) Initialize(ctx context.Context) error {
	return nil
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	var (
		resTransactions = make([]model.Transaction, 0)
		limiter         = make(chan struct{}, ThreadSize)

		waitGroup sync.WaitGroup
		locker    sync.Mutex
	)

	for _, platform := range platformList {

		limiter <- struct{}{}
		waitGroup.Add(1)

		go func(auctionPlatform string) {
			defer func() {
				<-limiter
				waitGroup.Done()
			}()

			var (
				internalTransactions []model.Transaction
				err                  error
			)

			switch auctionPlatform {
			case protocol.PlatformNouns:
				internalTransactions, err = i.handleNounsTransactions(ctx, message)
			case protocol.PlatformFoundation:
				internalTransactions, err = i.handleFoundationTransactions(ctx, message)
			default:
				return
			}

			if err != nil {
				zap.L().Warn("handle transactions", zap.Error(err), zap.String("platform", auctionPlatform))

				return
			}

			locker.Lock()
			resTransactions = append(resTransactions, internalTransactions...)
			locker.Unlock()
		}(platform)
	}

	waitGroup.Wait()

	return resTransactions, nil
}

func (i *internal) Jobs() []worker.Job {
	return nil
}

func New(kuroraClient *kurora.Client) worker.Worker {
	return &internal{
		tokenClient:  token.New(),
		kuroraClient: kuroraClient,
	}
}
