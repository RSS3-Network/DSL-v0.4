package lens

import (
	"context"
	"encoding/json"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/lens"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/datasource"
)

const (
	Source = "lens"
)

var _ datasource.Datasource = &Datasource{}

type Datasource struct {
	lensClient *lens.Client
}

func (d *Datasource) Name() string {
	return Source
}

func (d *Datasource) Networks() []string {
	return []string{
		protocol.NetworkArweave,
	}
}

func (d *Datasource) Handle(ctx context.Context, message *protocol.Message) ([]model.Transfer, error) {
	internalTransfers := make([]model.Transfer, 0)

	result, err := d.lensClient.GetPublicationsByAddress(ctx, message.Address)
	if err != nil {
		return nil, err
	}

	for _, publication := range result {
		sourceData, err := json.Marshal(publication)
		if err != nil {
			return nil, err
		}

		var metadataModel metadata.Metadata

		media, err := json.Marshal(publication.Metadata.Media)
		if err != nil {
			return nil, err
		}

		metadataModel.Lens = &metadata.Lens{
			Type:    string(publication.Type),
			Content: string(publication.Metadata.Description),
			Media:   media,
		}

		rawMetadata, err := json.Marshal(metadataModel)
		if err != nil {
			return nil, err
		}

		internalTransfers = append(internalTransfers, model.Transfer{
			TransactionHash: string(publication.ID),
			Timestamp:       publication.CreatedAt,
			AddressFrom:     message.Address,
			AddressTo:       "",
			Metadata:        rawMetadata,
			Network:         protocol.NetworkPolygon,
			Source:          Source,
			SourceData:      sourceData,
		})
	}

	return internalTransfers, nil
}

func New() datasource.Datasource {
	return &Datasource{
		lensClient: lens.NewClient(),
	}
}
