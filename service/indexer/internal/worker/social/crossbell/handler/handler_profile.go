package handler

import (
	"context"
	"encoding/json"

	"github.com/naturalselectionlabs/pregod/common/datasource/nft"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/profile"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"go.opentelemetry.io/otel"
)

var _ Interface = (*profileHandler)(nil)

type profileHandler struct {
	profileContract *profile.Profile
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
	case contract.EventHashLinkProfile:
		return p.handleLinkProfile(ctx, transaction, transfer, log)
	case contract.EventHashUnlinkProfile:
		return p.handleUnLinkProfile(ctx, transaction, transfer, log)
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
		Event: contract.EventNameCharacterCreated,
		Character: &metadata.CrossbellCharacter{
			ID:       event.ProfileId,
			Metadata: profileMetadata,
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag = filter.UpdateTag(filter.TagSocial, transfer.Tag)
	if transfer.Tag == filter.TagSocial {
		transfer.Type = filter.SocialProfile
	}

	return &transfer, nil
}

func (p *profileHandler) handleLinkProfile(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleLinkProfile")

	defer snap.End()

	event, err := p.profileContract.ParseLinkProfile(log)
	if err != nil {
		return nil, err
	}

	fromProfileMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.FromProfileId)
	toProfileMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.ToProfileId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNameLinkCharacter,
		Link: &metadata.CrossbellLink{
			Type: contract.LinkTypeMap[event.LinkType],
			From: &metadata.CrossbellCharacter{
				ID:       event.FromProfileId,
				Metadata: fromProfileMetadata,
			},
			To: &metadata.CrossbellCharacter{
				ID:       event.ToProfileId,
				Metadata: toProfileMetadata,
			},
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag = filter.UpdateTag(filter.TagSocial, transfer.Tag)
	if transfer.Tag == filter.TagSocial {
		transfer.Type = filter.SocialFollow
	}

	return &transfer, nil
}

func (p *profileHandler) handleUnLinkProfile(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleUnLinkProfile")

	defer snap.End()

	event, err := p.profileContract.ParseUnlinkProfile(log)
	if err != nil {
		return nil, err
	}

	fromProfileMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.FromProfileId)
	toProfileMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.ToProfileId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNameUnlinkCharacter,
		Link: &metadata.CrossbellLink{
			Type: contract.LinkTypeMap[event.LinkType],
			From: &metadata.CrossbellCharacter{
				ID:       event.FromProfileId,
				Metadata: fromProfileMetadata,
			},
			To: &metadata.CrossbellCharacter{
				ID:       event.ToProfileId,
				Metadata: toProfileMetadata,
			},
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag = filter.UpdateTag(filter.TagSocial, transfer.Tag)
	if transfer.Tag == filter.TagSocial {
		transfer.Type = filter.SocialUnfollow
	}

	return &transfer, nil
}

func (p *profileHandler) handleSetProfileUri(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleSetProfileUri")

	defer snap.End()

	event, err := p.profileContract.ParseSetProfileUri(log)
	if err != nil {
		return nil, err
	}

	profileMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.ProfileId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNameSetCharacterUri,
		Character: &metadata.CrossbellCharacter{
			ID:       event.ProfileId,
			URI:      event.NewUri,
			Metadata: profileMetadata,
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag = filter.UpdateTag(filter.TagSocial, transfer.Tag)
	if transfer.Tag == filter.TagSocial {
		transfer.Type = filter.SocialProfile
	}

	return &transfer, nil
}
