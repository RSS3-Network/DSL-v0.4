package token

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	moralisx "github.com/naturalselectionlabs/pregod/common/moralis"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shopspring/decimal"
)

var _ worker.Worker = &service{}

type service struct {
}

func (s *service) Name() string {
	return "token"
}

func (s *service) Network() []string {
	return []string{
		protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkBinanceSmartChain,
	}
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transfers []model.Transfer) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	for _, transfer := range transfers {
		if transfer.Source != moralis.Source {
			continue
		}

		sourceDataMap := make(map[string]interface{})

		if err := json.Unmarshal(transfer.SourceData, &sourceDataMap); err != nil {
			return nil, err
		}

		var metadataModel metadata.Metadata

		if err := json.Unmarshal(transfer.Metadata, &metadataModel); err != nil {
			return nil, err
		}

		if _, exist := sourceDataMap["contract_type"]; exist {
			// NFT transfer
			nftTransfer := moralisx.NFTTransfer{}
			if err := json.Unmarshal(transfer.SourceData, &nftTransfer); err != nil {
				return nil, err
			}

			tokenID, err := decimal.NewFromString(nftTransfer.TokenId)
			if err != nil {
				return nil, err
			}

			tokenValue := decimal.NewFromInt(1)

			metadataModel.Token = &metadata.Token{
				TokenAddress:  nftTransfer.TokenAddress,
				TokenStandard: strings.ToLower(nftTransfer.ContractType),
				TokenID:       &tokenID,
				TokenValue:    &tokenValue, // TODO ERC1155
			}
		} else if _, exist = sourceDataMap["address"]; exist {
			// Token transfer
			tokenTransfer := moralisx.TokenTransfer{}
			if err := json.Unmarshal(transfer.SourceData, &tokenTransfer); err != nil {
				return nil, err
			}

			tokenValue, err := decimal.NewFromString(tokenTransfer.Value)
			if err != nil {
				return nil, err
			}

			metadataModel.Token = &metadata.Token{
				TokenAddress:  tokenTransfer.Address,
				TokenStandard: "erc20",
				TokenValue:    &tokenValue,
			}
		} else {
			// Native transfer
			nativeTransfer := moralisx.Transaction{}
			if err := json.Unmarshal(transfer.SourceData, &nativeTransfer); err != nil {
				return nil, err
			}

			tokenValue, err := decimal.NewFromString(nativeTransfer.Value)
			if err != nil {
				return nil, err
			}

			metadataModel.Token = &metadata.Token{
				TokenStandard: "native",
				TokenValue:    &tokenValue,
			}
		}

		rawMetadata, err := json.Marshal(metadataModel)
		if err != nil {
			return nil, err
		}

		transfer.Metadata = rawMetadata

		internalTransfers = append(internalTransfers, transfer)
	}

	return internalTransfers, nil
}

func New() worker.Worker {
	return &service{}
}
