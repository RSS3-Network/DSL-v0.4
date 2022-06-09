package nft

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/nft"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/sirupsen/logrus"
)

var _ worker.Worker = &service{}

type service struct {
	nftClient *nft.Client
}

func (s *service) Name() string {
	return "nft"
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error) {
	logrus.Info("nft handle")
	internalTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {

		var metadataModel metadata.Metadata

		if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
			return nil, err
		}

		// 0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85 is ENS token address
		if metadataModel.Token.TokenAddress == "0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85" {
			continue
		}

		if metadataModel.Token.TokenStandard != "erc721" &&
			metadataModel.Token.TokenStandard != "erc1155" {
			continue
		}

		logrus.Infof("tokenid: %+v", metadataModel.Token)

		tokenAddress := metadataModel.Token.TokenAddress
		tokenID := metadataModel.Token.TokenID
		tokenIDNum := big.NewInt(1)
		tokenIDNum.SetString(tokenID.String(), 0)

		nftMetadataByte, err := s.nftClient.GetMetadata("", message.Network, common.HexToAddress(tokenAddress), tokenIDNum)
		if err != nil {
			logrus.Warn(err)

			continue
		}

		logrus.Debugf("nftMetadataByte: %s", nftMetadataByte)

		metadataModel.NFT = nftMetadataByte

	}

	return internalTransfers, nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func New(projectId string) worker.Worker {
	return &service{
		nftClient: nft.NewClient(projectId),
	}
}
