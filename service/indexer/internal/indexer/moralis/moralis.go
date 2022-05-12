package moralis

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/indexer"
	"github.com/sirupsen/logrus"
)

var _ indexer.Worker = &Indexer{}

type Indexer struct {
	config         *config.Config
	databaseClient *database.Client
	moralisClient  *moralis.Client
}

func (i *Indexer) Initialize() error {
	i.moralisClient = moralis.NewClient(i.config.Moralis.Key)

	return nil
}

func (i *Indexer) Handle(message *protocol.Message) error {
	transfers, _, err := i.moralisClient.GetTokenTransfers(context.Background(), common.HexToAddress(message.Address), nil)
	if err != nil {
		return err
	}

	for _, transfer := range transfers {
		logrus.Infoln(transfer)
	}

	return nil
}

func New(key string) indexer.Worker {
	return &Indexer{}
}
