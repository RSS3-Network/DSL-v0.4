package job

import (
	"time"

	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/sirupsen/logrus"
)

var _ worker.Job = (*Example)(nil)

type Example struct{}

func (e *Example) Name() string {
	return "job_poap_example"
}

func (e *Example) Spec() string {
	return "* * * * *"
}

func (e *Example) Timeout() time.Duration {
	return time.Second * 10
}

func (e *Example) Run() {
	logrus.Println("hello world")
}
