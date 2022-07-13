package alchemy

import (
	"context"
	"errors"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/alchemy"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource_asset"
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

		for _, nft := range result.OwnedNFTs {
			asset := model.Asset{
				TokenAddress:  nft.Contract.Address,
				TokenID:       nft.ID.TokenID,
				TokenStandard: nft.ID.TokenMetadata.TokenType,
				Network:       message.Network,
				Owner:         message.Address,
				Source:        Source,
				Title:         nft.Title,
				Description:   nft.Description,
				Timestamp:     nft.TimeLastUpdated,
			}
			if len(nft.Media) > 1 {
				asset.FileURL = nft.Media[0].Raw
			}
			assets = append(assets, asset)
		}

		if result.PageKey == "" {
			break
		}

		parameter.PageKey = result.PageKey
	}

	return assets, err
}

func New(config *configx.RPC) (datasource_asset.Datasource, error) {
	internalDatasource := Datasource{
		rpcClientMap: map[string]*alchemy.Client{},
	}

	var err error

	if internalDatasource.rpcClientMap[protocol.NetworkEthereum], err = alchemy.NewNFTClient(protocol.NetworkEthereum, config.Alchemy.Ethereum); err != nil {
		return nil, err
	}
	if internalDatasource.rpcClientMap[protocol.NetworkPolygon], err = alchemy.NewNFTClient(protocol.NetworkPolygon, config.Alchemy.Polygon); err != nil {
		return nil, err
	}

	return &internalDatasource, nil
}
