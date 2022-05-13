package moralis

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/indexer"
	"github.com/sirupsen/logrus"
)

var _ indexer.Worker = &Indexer{}

type Indexer struct {
	config         *configx.Moralis
	databaseClient *database.Client
	moralisClient  *moralis.Client
}

func (i *Indexer) Initialize() error {
	i.moralisClient = moralis.NewClient(i.config.Key)

	return nil
}

func (i *Indexer) Handle(message *protocol.Message) error {
	// TODO Query latest timestamp
	transfers, _, err := i.moralisClient.GetNFTTransfers(context.Background(), common.HexToAddress(message.Address), nil)
	if err != nil {
		return err
	}

	switch message.Network {
	case protocol.NetworkEthereum:

	case protocol.NetworkPolygon:
	}

	for _, transfer := range transfers {
		logrus.Infoln(transfer)
	}

	return nil
}

func New(config *configx.Moralis) indexer.Worker {
	return &Indexer{
		config: config,
	}
}
