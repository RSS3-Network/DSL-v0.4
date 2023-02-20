package sound

import (
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"time"
)

var _ worker.Job = (*Job)(nil)

type Job struct {
	kuroraClient     *kurora.Client
	blockNumberLocal int64
}

func (j *Job) Name() string {
	return "sound"
}

func (j *Job) Spec() string {
	return "*/5 * * * *" // 5 min
}

func (j *Job) Timeout() time.Duration {
	return time.Minute * 4
}

func (j *Job) Run(renewal worker.RenewalFunc) error {
	// load sound contract

}
