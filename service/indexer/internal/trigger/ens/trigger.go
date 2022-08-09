package ens

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/trigger"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ trigger.Trigger = (*internal)(nil)

type internal struct {
	databaseClient *gorm.DB
	ensClient      *ens.Client
}

func (i *internal) Name() string {
	return "ens"
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkXDAI,
		protocol.NetworkCrossbell,
	}
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message) (err error) {
	tracer := otel.Tracer("profile_worker")
	_, span := tracer.Start(ctx, "profile_worker:Handle")

	defer opentelemetry.Log(span, nil, nil, err)

	// Get ENS profile
	profile, err := i.ensClient.GetProfile(message.Address)
	if err != nil {
		logrus.Errorf("[profile] Handle: ensClient.GetProfile error, %v", err)

		return err
	}

	return i.databaseClient.
		Model(&model.Profile{}).
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).
		Create(profile).
		Error
}

func New() trigger.Trigger {
	return &internal{
		databaseClient: database.Global(),
		ensClient:      ens.New(config.ConfigIndexer.RPC.General.Ethereum.HTTP),
	}
}