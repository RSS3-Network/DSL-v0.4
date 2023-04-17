package name_service

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/ens/registrar"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
)

func (i *internal) handleENSNameRenewed(ctx context.Context, message *protocol.Message, transaction model.Transaction, log types.Log, platform string) (*model.Transfer, error) {
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	ensContract, err := registrar.NewRegistrarFilterer(log.Address, ethclient)
	if err != nil {
		return nil, err
	}

	event, err := ensContract.ParseNameRenewed(log)
	if err != nil {
		return nil, err
	}

	nativeToken, err := i.tokenClient.Native(ctx, message.Network)
	if err != nil {
		return nil, err
	}

	nameServiceMetadata, err := i.buildNameServiceMetadata(fmt.Sprintf("%s.eth", event.Name), filter.CollectibleEditRenew, nativeToken, event.Cost, event.Expires, "", "")
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Index:           int64(log.Index),
		AddressFrom:     strings.ToLower(transaction.AddressFrom),
		AddressTo:       strings.ToLower(log.Address.String()),
		Metadata:        nameServiceMetadata,
		Network:         transaction.Network,
		Platform:        platform,
		Source:          transaction.Source,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}, nil
}
