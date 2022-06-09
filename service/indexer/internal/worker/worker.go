package worker

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Worker interface {
	Name() string
	Networks() []string
	Initialize(ctx context.Context) error
	// 清洗数据
	Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error)
	// 定时任务
	Jobs() []Job
}
