package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

type CrossbellProfilestruct struct {
	Type             string   `json:"type"`
	Avatars          []string `json:"avatars"`
	ConnectedAvatars []string `json:"connected_avatars"`
	Name             string   `json:"name"`
	Bio              string   `json:"bio"`
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

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)

	tempStructure := &CrossbellProfilestruct{}
	if err := json.Unmarshal(profileMetadata, &tempStructure); err != nil {
		return nil, err
	}
	fmt.Println(profileMetadata, tempStructure)

	profile := &model.Profile{
		Address:    transfer.AddressFrom,
		Platform:   transfer.Platform,
		Network:    transfer.Network,
		Source:     transfer.Network,
		ExpireAt:   time.Unix(0, 0),
		SourceData: profileMetadata,
	}

	profile.Name = tempStructure.Name
	profile.Handle = tempStructure.Name
	profile.Bio = tempStructure.Bio

	if len(tempStructure.Avatars) > 0 {
		for _, avatar := range tempStructure.Avatars {
			profile.ProfileUris = append(profile.ProfileUris, avatar)
		}
	}

	if len(tempStructure.ConnectedAvatars) > 0 {
		for _, avatar := range tempStructure.ConnectedAvatars {
			profile.SocialUris = append(profile.SocialUris, avatar)
		}
	}

	p.databaseClient.Model(&model.Profile{}).Create(profile)

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

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)

	return &transfer, nil
}
