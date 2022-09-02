package alchemy

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/alchemy"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource_asset"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

const (
	Source = "alchemy"
)

var (
	ErrorUnsupportedNetwork       = errors.New("unsupported network")
	ErrorFailedToParseBlockNumber = errors.New("failed to parse block number")
)

var _ datasource_asset.Datasource = &Datasource{}

type Datasource struct {
	rpcClientMap map[string]*alchemy.Client // Alchemy
	tokenClient  *token.Client
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) (assets []model.Asset, err error) {
	tracer := otel.Tracer("datasource_nft_alchemy")
	_, trace := tracer.Start(ctx, "datasource_nft_alchemy:Handle")

	defer func() { opentelemetry.Log(trace, message, assets, err) }()

	assets, err = d.getNFTs(ctx, message)
	if err != nil {
		logrus.Errorf("[datasource_nft] Handle: getNFTs error, %v", err)
		return nil, err
	}

	return assets, nil
}

func (d *Datasource) getNFTs(ctx context.Context, message *protocol.Message) (assets []model.Asset, err error) {
	tracer := otel.Tracer("datasource_nft_alchemy")
	_, trace := tracer.Start(ctx, "datasource_nft_alchemy:getNFTs")

	defer func() { opentelemetry.Log(trace, message, assets, err) }()

	alchemyClient, exist := d.rpcClientMap[message.Network]
	if !exist {
		return nil, ErrorUnsupportedNetwork
	}

	parameter := alchemy.GetNFTsParameter{
		Owner:        message.Address,
		WithMetadata: true,
	}

	for {
		result, err := alchemyClient.GetNFTs(context.Background(), parameter)
		if err != nil {
			return nil, err
		}

		if len(result.OwnedNFTs) == 0 {
			break
		}

		lop.ForEach(result.OwnedNFTs, func(nft alchemy.NFTResult, k int) {
			if len(nft.Title) == 0 {
				return
			}

			asset := model.Asset{
				TokenAddress:  nft.Contract.Address,
				TokenStandard: nft.ID.TokenMetadata.TokenType,
				Network:       message.Network,
				Owner:         message.Address,
				Source:        Source,
				Title:         nft.Title,
				Description:   nft.Description,
				Timestamp:     nft.TimeLastUpdated,
			}

			// token id
			tokenID := big.NewInt(0)
			tokenID.SetString(nft.ID.TokenID, 0)
			asset.TokenID = tokenID.String()

			// attribute
			attributes := []metadata.TokenAttribute{}
			if list, ok := nft.Metadata.Attributes.([]metadata.TokenAttribute); ok {
				for _, attribute := range list {
					attributes = append(attributes, metadata.TokenAttribute{
						TraitType: attribute.TraitType,
						Value:     attribute.Value,
					})
				}
				asset.Attributes, _ = json.Marshal(attributes)
			}

			asset.Image = nft.Metadata.Image
			if strings.HasPrefix(asset.Image, "ipfs://") {
				asset.Image = ipfs.GetDirectURL(asset.Image)
			}

			asset.RelatedUrls = append(asset.RelatedUrls, ethereum.BuildTokenURL(message.Network, asset.TokenAddress, tokenID.String())...)

			assets = append(assets, asset)
		})

		if result.PageKey == "" {
			break
		}

		parameter.PageKey = result.PageKey
	}

	return assets, err
}

func New(config *configx.RPC, ethereumClientMap map[string]*ethclient.Client) (datasource_asset.Datasource, error) {
	internalDatasource := Datasource{
		rpcClientMap: map[string]*alchemy.Client{},
		tokenClient:  token.New(ethereumClientMap),
	}

	var err error

	if internalDatasource.rpcClientMap[protocol.NetworkEthereum], err = alchemy.NewClient(protocol.NetworkEthereum, config.Alchemy.Ethereum); err != nil {
		return nil, err
	}
	if internalDatasource.rpcClientMap[protocol.NetworkPolygon], err = alchemy.NewClient(protocol.NetworkPolygon, config.Alchemy.Polygon); err != nil {
		return nil, err
	}

	return &internalDatasource, nil
}
