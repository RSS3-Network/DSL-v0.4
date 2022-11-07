package handler

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/profile"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"
)

var _ Interface = (*profileHandler)(nil)

type profileHandler struct {
	profileContract *profile.Profile
	tokenClient     *token.Client
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
	erc721Token, err := p.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, contract.AddressCharacter.String(), event.ProfileId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Address:   transfer.AddressFrom,
		Platform:  protocol.PlatformCrossbell,
		Network:   transfer.Network,
		Source:    transfer.Network,
		Type:      filter.SocialCreate,
		Handle:    event.Handle,
		CreatedAt: time.Unix(event.Timestamp.Int64(), 0),
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(profile); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	database.Global().Model(&social.Profile{}).Clauses(clause.OnConflict{
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

	erc721Token, err := p.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, contract.AddressCharacter.String(), event.ToProfileId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Platform: protocol.PlatformCrossbell,
		Network:  transfer.Network,
		Source:   transfer.Network,
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
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

	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

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

	erc721Token, err := p.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, contract.AddressCharacter.String(), event.ToProfileId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Platform: protocol.PlatformCrossbell,
		Network:  transfer.Network,
		Source:   transfer.Network,
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
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

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialUnfollow, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

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

	erc721Token, err := p.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, contract.AddressCharacter.String(), event.ProfileId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Address:  transfer.AddressFrom,
		Platform: protocol.PlatformCrossbell,
		Network:  transfer.Network,
		Source:   transfer.Network,
		Type:     filter.SocialUpdate,
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(profile); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}
