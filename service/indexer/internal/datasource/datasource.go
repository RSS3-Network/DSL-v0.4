package datasource

import (
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/config"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/alchemy"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/arweave"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/blockscout"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/moralis"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/pregod_etl/lens"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/internalModel"
)

var Datasources = []internalModel.Datasource{}

func Initialize() error {
	alchemyDatasource, err := alchemy.New(config.ConfigIndexer.RPC)
	if err != nil {
		return err
	}

	blockscoutDatasource, err := blockscout.New(config.ConfigIndexer.RPC)
	if err != nil {
		return err
	}

	lensDatasource, err := lens.New(config.ConfigIndexer.RPC)
	if err != nil {
		return err
	}

	Datasources = []internalModel.Datasource{
		alchemyDatasource,
		moralis.New(config.ConfigIndexer.Moralis.Key, config.ConfigIndexer.RPC),
		arweave.New(),
		blockscoutDatasource,
		zksync.New(),
		lensDatasource,
	}

	return nil
}
