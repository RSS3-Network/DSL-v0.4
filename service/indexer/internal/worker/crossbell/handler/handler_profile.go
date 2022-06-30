package handler

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/nft"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract/character"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract/periphery"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract/profile"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

var _ Interface = (*profileHandler)(nil)

type profileHandler struct {
	profileContract   *profile.Profile
	characterContract *character.Character
	peripheryContract *periphery.Periphery
}

func (p *profileHandler) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler_profile")

	ctx, snap := tracer.Start(ctx, "worker_crossbell_handler_profile:Handle")

	defer snap.End()

	var log types.Log

	if err := json.Unmarshal(transfer.SourceData, &log); err != nil {
		return nil, err
	}

	switch log.Topics[0] {
	case contract.EventHashProfileCreated:
		return p.handleProfileCreated(ctx, transaction, transfer, log)
	case contract.EventHashPostNote:
		return p.handlePostNote(ctx, transaction, transfer, log)
	default:
		return nil, contract.ErrorUnknownEvent
	}
}

func (p *profileHandler) handleProfileCreated(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleProfileCreated")

	defer snap.End()

	event, err := p.profileContract.ParseProfileCreated(log)
	if err != nil {
		return nil, err
	}

	// Self-hosted IPFS files may be out of date
	profileMetadata, _ := nft.GetMetadata(protocol.NetworkCrossbell, contract.AddressCharacter, event.ProfileId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNameProfileCreated,
		Profile: &metadata.CrossbellProfile{
			ID:       event.ProfileId,
			Metadata: profileMetadata,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (p *profileHandler) handlePostNote(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handlePostNote")

	defer snap.End()

	var err error

	event, err := p.profileContract.ParsePostNote(log)
	if err != nil {
		return nil, err
	}

	logrus.Infof("%+v", event)

	note, err := p.profileContract.GetNote(&bind.CallOpts{}, event.ProfileId, event.NoteId)
	if err != nil {
		return nil, err
	}

	// Self-hosted IPFS files may be out of date
	contentData, _ := ipfs.GetFileByURL(ctx, note.ContentUri)

	logrus.Infof("%+v", contentData)

	var noteMetadata json.RawMessage

	if err := json.Unmarshal(contentData, &noteMetadata); err != nil {
		return nil, err
	}

	logrus.Infof("%+v", noteMetadata)

	uri, err := p.peripheryContract.GetLinkingAnyUri(&bind.CallOpts{}, event.LinkKey)
	if err != nil {
		return nil, err
	}

	logrus.Infof("%+v", uri)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNamePostNote,
		Note: &metadata.CrossbellNote{
			ID:           event.NoteId,
			LinkItemType: note.LinkItemType,
			LinkKey:      note.LinkKey,
			Link:         uri,
			ContentURI:   note.ContentUri,
			LinkModule:   note.LinkModule,
			MintModule:   note.MintModule,
			MintNFT:      note.MintNFT,
			Deleted:      note.Deleted,
			Locked:       note.Locked,
			Metadata:     noteMetadata,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}
