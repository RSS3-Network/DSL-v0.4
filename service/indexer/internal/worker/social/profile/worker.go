package profile

import (
	"context"

	"gorm.io/gorm/clause"

	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/sirupsen/logrus"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

const (
	Name = protocol.PlatformLens
)

var _ worker.Worker = (*service)(nil)

type service struct {
	databaseClient *gorm.DB
	ensClient      *ens.Client
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		ensClient:      ens.New(config.ConfigIndexer.RPC.General.Ethereum.HTTP),
	}
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkZkSync, protocol.NetworkCrossbell, protocol.NetworkXDAI,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("profile_worker")
	_, trace := tracer.Start(ctx, "profile_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	// get ENS profile
	profile, err := s.ensClient.GetProfile(message.Address)
	if err != nil {
		logrus.Errorf("[profile] Handle: ensClient.GetProfile error, %v", err)
		return transactions, nil
	}

	s.databaseClient.Model(&model.Profile{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(profile)

	return transactions, nil
}
