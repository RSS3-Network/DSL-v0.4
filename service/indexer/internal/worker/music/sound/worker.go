package sound

import (
	"context"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.uber.org/zap"
	"net/http"
)

var (
	_ worker.Worker = (*service)(nil)

	soundMap = make(map[string]struct{})
)

type service struct {
	kuroraClient *kurora.Client
}

func (s *service) Name() string {
	return "sound"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&Job{},
	}
}

func New(endpoint string) (worker.Worker, error) {
	sound := &service{}

	var err error
	sound.kuroraClient, err = kurora.Dial(context.Background(), endpoint, kurora.WithHTTPClient(http.DefaultClient))
	if err != nil {
		loggerx.Global().Error(" kurora Dial error", zap.Error(err))
		return nil, err
	}

	return sound, nil
}
