package crossbell

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/worker/crossbell/handler"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type Client struct {
	httpClient *http.Client
	handler    handler.Interface
}

func (c *Client) GetProfile(address string) ([]*social.Profile, error) {
	requestURL := &url.URL{
		Scheme: "https",
		Host:   "indexer.crossbell.io",
		Path:   fmt.Sprintf("/v1/addresses/%v/characters", address),
	}

	httpResponse, err := c.httpClient.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	profiles := GetProfilesResponse{}

	if err := json.NewDecoder(httpResponse.Body).Decode(&profiles); err != nil {
		return nil, err
	}

	var result []*social.Profile
	var mutex sync.Mutex

	lop.ForEach(profiles.List, func(crossbell ProfileResponse, i int) {
		profile := &social.Profile{
			Address:  address,
			Network:  protocol.NetworkCrossbell,
			Platform: protocol.PlatformCrossbell,
			Source:   protocol.PlatformCrossbell,
			Handle:   fmt.Sprintf("%v.csb", crossbell.Handle),
		}

		content, err := ipfs.GetFileByURL(crossbell.URI)
		if err != nil {
			logrus.Errorf("[common] crossbell: ipfs.GetFileByURL err, %v, uri: %v", err, crossbell.URI)

			return
		}

		var metadata *metadata.Token
		if err := json.Unmarshal(content, &metadata); err != nil {
			logrus.Errorf("[common] crossbell: json.Unmarshal err, %v, uri: %v", err, crossbell.URI)

			return
		}

		if len(metadata.Image) > 0 {
			profile.Name = metadata.Name
			profile.ProfileUris = []string{ipfs.GetDirectURL(metadata.Image)}
			result = append(result, profile)
			return
		}

		var crossbellProfile *CrossbellProfileStruct
		if err := json.Unmarshal(content, &crossbellProfile); err != nil {
			logrus.Errorf("[common] crossbell: json.Unmarshal err, %v, uri: %v", err, crossbell.URI)

			return
		}

		profile.Name = crossbellProfile.Name
		profile.Bio = crossbellProfile.Bio
		for _, avatar := range crossbellProfile.Avatars {
			profile.ProfileUris = append(profile.ProfileUris, ipfs.GetDirectURL(avatar))
		}

		mutex.Lock()
		result = append(result, profile)
		mutex.Unlock()
	})

	return result, nil
}

func (c *Client) HandleReceipt(ctx context.Context, message *protocol.Message, transaction *model.Transaction, receipt *types.Receipt, transferMap map[int64]model.Transfer) ([]model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell")

	_, snap := tracer.Start(ctx, "worker_crossbell:handleReceipt")

	defer snap.End()

	internalTransfers := make([]model.Transfer, 0)

	for _, log := range receipt.Logs {
		logIndex := int64(log.Index)

		transfer, exist := transferMap[logIndex]
		if !exist {
			sourceData, err := json.Marshal(log)
			if err != nil {
				return nil, err
			}

			// Virtual transfer
			transfer = model.Transfer{
				TransactionHash: transaction.Hash,
				Timestamp:       transaction.Timestamp,
				Index:           logIndex,
				AddressFrom:     transaction.AddressFrom,
				AddressTo:       transaction.AddressTo,
				Metadata:        metadata.Default,
				Network:         message.Network,
				Source:          protocol.SourceOrigin,
				SourceData:      sourceData,
			}
		}

		internalTransfer, err := c.handler.Handle(ctx, transaction, transfer)
		if err != nil {
			if !errors.Is(err, crossbell.ErrorUnknownEvent) {
				loggerx.Global().Warn(
					"failed to handle crossbell transfer",
					zap.Error(err),
					zap.String("address", message.Address),
					zap.String("transaction_hash", transaction.Hash),
					zap.String("network", message.Network),
				)
			}

			continue
		}

		internalTransfer.Tag, internalTransfer.Type = filter.UpdateTagAndType(transfer.Tag, internalTransfer.Tag, transfer.Type, internalTransfer.Type)
		if internalTransfer.Tag == transfer.Tag {
			internalTransfer.Platform = transfer.Platform
		}

		internalTransfers = append(internalTransfers, *internalTransfer)
	}

	return internalTransfers, nil
}

func New() (*Client, error) {
	handler, err := handler.New()
	if err != nil {
		loggerx.Global().Error("crossbell handler init error", zap.Error(err))

		return nil, err
	}

	return &Client{
		httpClient: http.DefaultClient,
		handler:    handler,
	}, nil
}
