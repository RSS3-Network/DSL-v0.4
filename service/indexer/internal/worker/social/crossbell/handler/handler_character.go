package handler

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/character"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/periphery"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm/clause"
)

var _ Interface = (*characterHandler)(nil)

type characterHandler struct {
	characterContract *character.Character
	peripheryContract *periphery.Periphery
	profileHandler    *profileHandler
	tokenClient       *token.Client
}

func (c *characterHandler) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	var log types.Log

	if err := json.Unmarshal(transfer.SourceData, &log); err != nil {
		return nil, err
	}

	switch log.Topics[0] {
	case contract.EventHashCharacterCreated, contract.EventHashProfileCreated:
		// Broken change
		if transaction.BlockNumber >= contract.BrokenBlockNumber {
			return c.handleCharacterCreated(ctx, transaction, transfer, log)
		}

		return c.profileHandler.handleProfileCreated(ctx, transaction, transfer, log)
	case contract.EventHashSetHandle:
		return c.handleSetHandle(ctx, transaction, transfer, log)
	case contract.EventHashPostNote:
		return c.handlePostNote(ctx, transaction, transfer, log)
	case contract.EventHashLinkCharacter, contract.EventHashLinkProfile:
		// Broken change
		if transaction.BlockNumber >= contract.BrokenBlockNumber {
			return c.handleLinkCharacter(ctx, transaction, transfer, log)
		}

		return c.profileHandler.handleLinkProfile(ctx, transaction, transfer, log)
	case contract.EventHashUnlinkCharacter, contract.EventHashUnlinkProfile:
		// Broken change
		if transaction.BlockNumber >= contract.BrokenBlockNumber {
			return c.handleUnLinkCharacter(ctx, transaction, transfer, log)
		}

		return c.profileHandler.handleUnLinkProfile(ctx, transaction, transfer, log)
	case contract.EventHashSetCharacterUri, contract.EventHashSetProfileUri:
		// Broken change
		if transaction.BlockNumber >= contract.BrokenBlockNumber {
			return c.handleSetCharacterUri(ctx, transaction, transfer, log)
		}

		return c.profileHandler.handleSetProfileUri(ctx, transaction, transfer, log)
	case contract.EventHashSetNoteUri:
		return c.handleSetNoteUri(ctx, transaction, transfer, log)
	default:
		return nil, contract.ErrorUnknownEvent
	}
}

func (c *characterHandler) handleCharacterCreated(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleProfileCreated")

	defer snap.End()

	event, err := c.characterContract.ParseCharacterCreated(log)
	if err != nil {
		return nil, err
	}

	// Self-hosted IPFS files may be out of date
	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, contract.AddressCharacter.String(), event.CharacterId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Address: transfer.AddressFrom,
		// TODO: use appId from CSB
		// Platform: transfer.Platform,
		Platform: transfer.Network,
		Network:  transfer.Network,
		Source:   transfer.Network,
		Type:     filter.SocialProfileCreate,
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

func (c *characterHandler) handleSetHandle(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleSetHandle")

	defer snap.End()

	event, err := c.characterContract.ParseSetHandle(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, contract.AddressCharacter.String(), event.CharacterId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Address: transfer.AddressFrom,
		// TODO: use appId from CSB
		// Platform: transfer.Platform,
		Platform: transfer.Network,
		Network:  transfer.Network,
		Source:   transfer.Network,
		Type:     filter.SocialProfileUpdate,
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

func (c *characterHandler) handlePostNote(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handlePostNote")

	defer snap.End()

	var err error

	event, err := c.characterContract.ParsePostNote(log)
	if err != nil {
		return nil, err
	}

	note, err := c.characterContract.GetNote(&bind.CallOpts{}, event.CharacterId, event.NoteId)
	if err != nil {
		return nil, err
	}

	transfer.RelatedUrls = []string{note.ContentUri}

	// Self-hosted IPFS files may be out of date
	contentData, _ := ipfs.GetFileByURL(note.ContentUri)

	postOriginal := CrossbellPostStruct{}

	if err := json.Unmarshal(contentData, &postOriginal); err != nil {
		return nil, err
	}

	post := &metadata.Post{
		TypeOnPlatform: []string{contract.EventNamePostNote},
		Title:          postOriginal.Content,
		Body:           postOriginal.Content,
	}

	for _, attachment := range postOriginal.Attachments {
		post.Media = append(post.Media, metadata.Media{
			Address:  attachment.Address,
			MimeType: attachment.MimeType,
		})
	}

	if transfer.Metadata, err = json.Marshal(post); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialPost, transfer.Type)

	return &transfer, nil
}

func (c *characterHandler) handleLinkCharacter(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleLinkCharacter")

	defer snap.End()

	event, err := c.characterContract.ParseLinkCharacter(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, contract.AddressCharacter.String(), event.ToCharacterId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		// TODO: use appId from CSB
		// Platform: transfer.Platform,
		Platform: transfer.Network,
		Network:  transfer.Network,
		Source:   transfer.Network,
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
		return nil, err
	}

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.ToCharacterId)
	if err != nil {
		return nil, err
	}
	profile.Address = strings.ToLower(characterOwner.String())

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialFollow, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}

func (c *characterHandler) handleUnLinkCharacter(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleUnLinkCharacter")

	defer snap.End()

	event, err := c.characterContract.ParseLinkCharacter(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, contract.AddressCharacter.String(), event.ToCharacterId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		// TODO: use appId from CSB
		// Platform: transfer.Platform,
		Platform: transfer.Network,
		Network:  transfer.Network,
		Source:   transfer.Network,
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
		return nil, err
	}

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.ToCharacterId)
	if err != nil {
		return nil, err
	}
	profile.Address = strings.ToLower(characterOwner.String())

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialUnfollow, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}

func (c *characterHandler) handleSetCharacterUri(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleSetCharacterUri")

	defer snap.End()

	event, err := c.characterContract.ParseSetCharacterUri(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, contract.AddressCharacter.String(), event.CharacterId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Address: transfer.AddressFrom,
		// TODO: use appId from CSB
		// Platform: transfer.Platform,
		Platform: transfer.Network,
		Network:  transfer.Network,
		Source:   transfer.Network,
		Type:     filter.SocialProfileUpdate,
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

func (c *characterHandler) handleSetNoteUri(ctx context.Context, transaction model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleSetNoteUri")

	defer snap.End()

	var err error

	event, err := c.characterContract.ParseSetNoteUri(log)
	if err != nil {
		return nil, err
	}

	note, err := c.characterContract.GetNote(&bind.CallOpts{}, event.CharacterId, event.NoteId)
	if err != nil {
		return nil, err
	}

	// Self-hosted IPFS files may be out of date
	contentData, _ := ipfs.GetFileByURL(note.ContentUri)

	var noteMetadata json.RawMessage

	if err := json.Unmarshal(contentData, &noteMetadata); err != nil {
		return nil, err
	}

	uri, err := c.peripheryContract.GetLinkingAnyUri(&bind.CallOpts{}, note.LinkKey)
	if err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(&metadata.Crossbell{
		Event: contract.EventNameSetNoteUri,
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

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialPost, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}
