package job

import (
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"time"
)

type SnapshotNetworkJob struct {
}

func (job *SnapshotNetworkJob) Name() string {
	return "snapshot_network_job"
}

func (job *SnapshotNetworkJob) Spec() string {
	return "* /2880 * * * *" // 2 days
}

func (job *SnapshotNetworkJob) Timeout() time.Duration {
	return time.Minute * 5
}

func (job *SnapshotNetworkJob) Run(renewal worker.RenewalFunc) error {

}
