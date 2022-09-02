package worker_test

import (
	"context"
	"time"

	"github.com/naturalselectionlabs/pregod/service/indexer/internal/internalModel"
)

const (
	Name = "example"
)

var _ internalModel.Job = (*example)(nil)

type example struct{}

func (e *example) Name() string {
	return Name
}

func (e *example) Spec() string {
	return "* * * * *"
}

func (e *example) Timeout() time.Duration {
	return time.Minute
}

func (e *example) Run(renewal internalModel.RenewalFunc) error {
	for {
		time.Sleep(time.Second)

		if err := renewal(context.Background(), time.Minute); err != nil {
			return err
		}
	}
}
