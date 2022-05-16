package gitcoin

import (
	"context"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/sirupsen/logrus"
)

const (
	TagGitCoin = "gitcoin"
)

var grantMap = map[string]string{
	"0xde21f729137c5af1b01d73af1dc21effa2b8a0d6": "GitCoin Grants",
}

var _ worker.Worker = &Worker{}

type Worker struct{}

func (w *Worker) Handle(ctx context.Context, transfers []model.Transfer) ([]model.Transfer, error) {
	gitCoinTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {
		grantName, exist := grantMap[transfer.AddressTo]
		if !exist {
			continue
		}

		transfer.Tags = append(transfer.Tags, TagGitCoin)
		gitCoinTransfers = append(gitCoinTransfers, transfer)

		// TODO Save grant name to metadata
		logrus.Infoln(grantName)
	}

	return gitCoinTransfers, nil
}
