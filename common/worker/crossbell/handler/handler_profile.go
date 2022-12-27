package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell/contract/profile"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"

	"go.opentelemetry.io/otel"
)

var _ Interface = (*profileHandler)(nil)

type profileHandler struct {
	profileContract *profile.Profile
	tokenClient     *token.Client
}

func (p *profileHandler) Handle(ctx context.Context, transaction *model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler_profile")

	ctx, snap := tracer.Start(ctx, "worker_crossbell_handler_profile:Handle")

	defer snap.End()

	var log types.Log

	if err := json.Unmarshal(transfer.SourceData, &log); err != nil {
		return nil, err
	}

	switch log.Topics[0] {
	case crossbell.EventHashProfileCreated:
		return p.handleProfileCreated(ctx, transaction, transfer, log)
	case crossbell.EventHashLinkProfile:
		return p.handleLinkProfile(ctx, transaction, transfer, log)
	case crossbell.EventHashUnlinkProfile:
		return p.handleUnLinkProfile(ctx, transaction, transfer, log)
	default:
		return nil, crossbell.ErrorUnknownEvent
	}
}

func (p *profileHandler) handleProfileCreated(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleProfileCreated")

	defer snap.End()

	event, err := p.profileContract.ParseProfileCreated(log)
	if err != nil {
		return nil, err
	}

	// Self-hosted IPFS files may be out of date
	erc721Token, err := p.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, crossbell.AddressCharacter.String(), event.ProfileId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Address:   strings.ToLower(event.To.String()),
		Platform:  protocol.PlatformCrossbell,
		Network:   transfer.Network,
		Source:    transfer.Network,
		Action:    filter.SocialCreate,
		Handle:    fmt.Sprintf("%v.csb", event.Handle),
		CreatedAt: time.Unix(event.Timestamp.Int64(), 0),
		URL:       fmt.Sprintf("https://crossbell.io/@%v", event.Handle),
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
		return nil, err
	}

	// transaction
	transaction.Owner = strings.ToLower(event.Creator.String())

	// transfer
	if transfer.Metadata, err = json.Marshal(profile); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}

func (p *profileHandler) handleLinkProfile(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleLinkProfile")

	defer snap.End()

	event, err := p.profileContract.ParseLinkProfile(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := p.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, crossbell.AddressCharacter.String(), event.ToProfileId)
	if err != nil {
		return nil, err
	}

	// profile address
	characterOwner, err := p.profileContract.OwnerOf(&bind.CallOpts{BlockNumber: big.NewInt(transaction.BlockNumber)}, event.ToProfileId)
	if err != nil {
		return nil, err
	}

	// profile handle
	handle, err := p.profileContract.GetHandle(&bind.CallOpts{BlockNumber: big.NewInt(transaction.BlockNumber)}, event.ToProfileId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Address:  strings.ToLower(characterOwner.String()),
		Handle:   handle,
		Platform: protocol.PlatformCrossbell,
		Network:  transfer.Network,
		Source:   transfer.Network,
		URL:      fmt.Sprintf("https://crossbell.io/@%v", handle),
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
		return nil, err
	}

	// transaction
	transaction.Owner = strings.ToLower(event.Account.String())

	// transfer
	if transfer.Metadata, err = json.Marshal(profile); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialFollow, transfer.Type)

	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}

func (p *profileHandler) handleUnLinkProfile(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleUnLinkProfile")

	defer snap.End()

	event, err := p.profileContract.ParseUnlinkProfile(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := p.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, crossbell.AddressCharacter.String(), event.ToProfileId)
	if err != nil {
		return nil, err
	}

	// profile address
	characterOwner, err := p.profileContract.OwnerOf(&bind.CallOpts{BlockNumber: big.NewInt(transaction.BlockNumber)}, event.ToProfileId)
	if err != nil {
		return nil, err
	}

	// profile handle
	handle, err := p.profileContract.GetHandle(&bind.CallOpts{BlockNumber: big.NewInt(transaction.BlockNumber)}, event.ToProfileId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Address:  strings.ToLower(characterOwner.String()),
		Handle:   fmt.Sprintf("%v.csb", handle),
		Platform: protocol.PlatformCrossbell,
		Network:  transfer.Network,
		Source:   transfer.Network,
		URL:      fmt.Sprintf("https://crossbell.io/@%v", handle),
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
		return nil, err
	}

	// transaction
	transaction.Owner = strings.ToLower(event.Account.String())

	if transfer.Metadata, err = json.Marshal(profile); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialUnfollow, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}

func (p *profileHandler) handleSetProfileUri(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleSetProfileUri")

	defer snap.End()

	event, err := p.profileContract.ParseSetProfileUri(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := p.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, crossbell.AddressCharacter.String(), event.ProfileId)
	if err != nil {
		return nil, err
	}

	// profile address
	characterOwner, err := p.profileContract.OwnerOf(&bind.CallOpts{BlockNumber: big.NewInt(transaction.BlockNumber)}, event.ProfileId)
	if err != nil {
		return nil, err
	}

	// profile handle
	handle, err := p.profileContract.GetHandle(&bind.CallOpts{BlockNumber: big.NewInt(transaction.BlockNumber)}, event.ProfileId)
	if err != nil {
		return nil, err
	}

	transaction.Owner = strings.ToLower(characterOwner.String())

	profile := &social.Profile{
		Address:  transaction.Owner,
		Handle:   fmt.Sprintf("%v.csb", handle),
		Platform: protocol.PlatformCrossbell,
		Network:  transfer.Network,
		Source:   transfer.Network,
		Action:   filter.SocialUpdate,
		URL:      fmt.Sprintf("https://crossbell.io/@%v", handle),
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
		return nil, err
	}

	// transaction
	transaction.Owner = strings.ToLower(characterOwner.String())

	// transfer
	if transfer.Metadata, err = json.Marshal(profile); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}
