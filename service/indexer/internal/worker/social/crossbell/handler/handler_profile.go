package handler

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/nft"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/profile"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ Interface = (*profileHandler)(nil)

type profileHandler struct {
	profileContract *profile.Profile
	databaseClient  *gorm.DB
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

	profile := &model.Profile{
		Address:  transfer.AddressFrom,
		Platform: transfer.Platform,
		Network:  transfer.Network,
		Source:   transfer.Network,
	}

	if err = BuildProfileMetadata(profileMetadata, profile); err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(profile); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)

	p.databaseClient.Model(&model.Profile{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(profile)

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

	toProfileMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.ToProfileId)

	profile := &model.Profile{
		Address:  transfer.AddressFrom,
		Platform: transfer.Platform,
		Network:  transfer.Network,
		Source:   transfer.Network,
	}

	if err = BuildProfileMetadata(toProfileMetadata, profile); err != nil {
		return nil, err
	}

	characterOwner, err := p.profileContract.OwnerOf(&bind.CallOpts{}, event.ToProfileId)
	if err != nil {
		return nil, err
	}
	profile.Address = strings.ToLower(characterOwner.String())

	if transfer.Metadata, err = json.Marshal(profile); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialFollow, transfer.Type)

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

	toProfileMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.ToProfileId)

	profile := &model.Profile{
		Address:  transfer.AddressFrom,
		Platform: transfer.Platform,
		Network:  transfer.Network,
		Source:   transfer.Network,
	}

	if err = BuildProfileMetadata(toProfileMetadata, profile); err != nil {
		return nil, err
	}

	characterOwner, err := p.profileContract.OwnerOf(&bind.CallOpts{}, event.ToProfileId)
	if err != nil {
		return nil, err
	}
	profile.Address = strings.ToLower(characterOwner.String())

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialUnfollow, transfer.Type)

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

	if transfer.Metadata, err = json.Marshal(&metadata.Crossbell{
		Event: contract.EventNameSetCharacterUri,
		Character: &metadata.CrossbellCharacter{
			ID:       event.ProfileId,
			URI:      event.NewUri,
			Metadata: profileMetadata,
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)

	return &transfer, nil
}
