package nftscan

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ipfs/go-cid"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/nftscan"
	"github.com/naturalselectionlabs/pregod/common/metadata_url"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource_asset"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

var _ datasource_asset.Datasource = (*Datasource)(nil)

type Datasource struct {
	nftscanClient nftscan.Client
}

func (d *Datasource) Name() string {
	return protocol.SourceNFTScan
}

func (d *Datasource) Networks() []string {
	return nftscan.SupportedNetworks()
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Asset, error) {
	collections, err := d.nftscanClient.GetAllNFTsByAccount(ctx, message.Network, message.Address)
	if err != nil {
		return nil, err
	}

	result := make([]model.Asset, 0, len(collections))

	for _, collection := range collections {
		for _, asset := range collection.Assets {
			var metadata token.Metadata
			if err := json.Unmarshal([]byte(asset.MetadataJSON), &metadata); err != nil {
				zap.L().Error("unmarshal metadata error", zap.Error(err), zap.String("contract", collection.ContractAddress), zap.Stringer("token_id", asset.TokenID))

				continue
			}

			if metadata.Name == "" {
				metadata.Name = metadata.Title
			}

			if metadata.ImageURL == "" {
				metadata.ImageURL = metadata.Image
			}

			if metadata.Description == "" {
				metadata.Description = collection.Description
			}

			if asset.OwnTimestamp.Equal(decimal.Zero) {
				asset.OwnTimestamp = asset.MintTimestamp
			}

			if strings.HasPrefix(metadata.Image, "ipfs://") {
				metadata.ImageURL = metadata_url.GetDirectURL(metadata.ImageURL)
			}

			// QmSu2SMLH2RszGce4CcBkEwXxFHqwsjWAMq2WGNU5U9qx5 -> https://ipfs.io/ipfs/QmSu2SMLH2RszGce4CcBkEwXxFHqwsjWAMq2WGNU5U9qx5
			if _, err := cid.Parse(strings.Split(metadata.ImageURL, "/")[0]); err == nil {
				metadata.ImageURL = metadata_url.GetDirectURL(fmt.Sprintf("ipfs://%s", metadata.ImageURL))
			}

			result = append(result, model.Asset{
				Network:       message.Network,
				TokenAddress:  strings.ToLower(collection.ContractAddress),
				TokenID:       asset.TokenID.String(),
				TokenStandard: strings.ToUpper(asset.ERCType), // ERC721 or ERC1155
				Owner:         message.Address,
				Source:        protocol.SourceNFTScan,
				Title:         metadata.Name,
				Description:   metadata.Description,
				Image:         metadata.ImageURL,
				RelatedUrls:   ethereum.BuildTokenURL(message.Network, collection.ContractAddress, asset.TokenID.String()),
				Timestamp:     time.UnixMilli(asset.OwnTimestamp.IntPart()),
			})
		}
	}

	return result, nil
}

func New(ctx context.Context, config configx.NFTScan) (datasource_asset.Datasource, error) {
	internalDatasource := Datasource{
		nftscanClient: lo.Must(nftscan.NewClient(ctx, config.APIKey)),
	}

	return &internalDatasource, nil
}
