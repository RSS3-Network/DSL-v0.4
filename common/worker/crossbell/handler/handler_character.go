package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell/contract/character"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell/contract/periphery"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"

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

func (c *characterHandler) Handle(ctx context.Context, transaction *model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	var log types.Log

	if err := json.Unmarshal(transfer.SourceData, &log); err != nil {
		return nil, err
	}

	// Default platform
	transfer.Platform = protocol.PlatformCrossbell

	switch log.Topics[0] {
	case crossbell.EventHashCharacterCreated, crossbell.EventHashProfileCreated:
		// Broken change
		if transaction.BlockNumber >= crossbell.BrokenBlockNumber {
			return c.handleCharacterCreated(ctx, transaction, transfer, log)
		}

		return c.profileHandler.handleProfileCreated(ctx, transaction, transfer, log)
	case crossbell.EventHashSetHandle:
		return c.handleSetHandle(ctx, transaction, transfer, log)
	case crossbell.EventHashPostNote:
		return c.handlePostNote(ctx, transaction, transfer, log)
	case crossbell.EventHashLinkCharacter, crossbell.EventHashLinkProfile:
		// Broken change
		if transaction.BlockNumber >= crossbell.BrokenBlockNumber {
			return c.handleLinkCharacter(ctx, transaction, transfer, log)
		}

		return c.profileHandler.handleLinkProfile(ctx, transaction, transfer, log)
	case crossbell.EventHashUnlinkCharacter, crossbell.EventHashUnlinkProfile:
		// Broken change
		if transaction.BlockNumber >= crossbell.BrokenBlockNumber {
			return c.handleUnLinkCharacter(ctx, transaction, transfer, log)
		}

		return c.profileHandler.handleUnLinkProfile(ctx, transaction, transfer, log)
	case crossbell.EventHashSetCharacterUri, crossbell.EventHashSetProfileUri:
		// Broken change
		if transaction.BlockNumber >= crossbell.BrokenBlockNumber {
			return c.handleSetCharacterUri(ctx, transaction, transfer, log)
		}

		return c.profileHandler.handleSetProfileUri(ctx, transaction, transfer, log)
	case crossbell.EventHashSetNoteUri:
		return c.handleSetNoteUri(ctx, transaction, transfer, log)
	case crossbell.EventHashMintNote:
		return c.handleMintNote(ctx, transaction, transfer, log)
	default:
		return nil, crossbell.ErrorUnknownEvent
	}
}

func (c *characterHandler) handleCharacterCreated(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleProfileCreated")

	defer snap.End()

	event, err := c.characterContract.ParseCharacterCreated(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, crossbell.AddressCharacter.String(), event.CharacterId)
	if err != nil {
		return nil, err
	}

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.CharacterId)
	if err != nil {
		return nil, err
	}
	transaction.Owner = strings.ToLower(characterOwner.String())

	profile := &social.Profile{
		Address:  transaction.Owner,
		Platform: protocol.PlatformCrossbell,
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

	if transfer.RelatedUrls, err = c.buildRelatedUrls([]string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}, transfer.Platform, event.CharacterId, nil); err != nil {
		return nil, err
	}

	database.Global().Model(&social.Profile{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(profile)

	return &transfer, nil
}

func (c *characterHandler) handleSetHandle(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleSetHandle")

	defer snap.End()

	event, err := c.characterContract.ParseSetHandle(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, crossbell.AddressCharacter.String(), event.CharacterId)
	if err != nil {
		return nil, err
	}

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.CharacterId)
	if err != nil {
		return nil, err
	}
	transaction.Owner = strings.ToLower(characterOwner.String())

	profile := &social.Profile{
		Address:  transaction.Owner,
		Platform: protocol.PlatformCrossbell,
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

	if transfer.RelatedUrls, err = c.buildRelatedUrls([]string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}, transfer.Platform, event.CharacterId, nil); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (c *characterHandler) handlePostNote(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handlePostNote")

	defer snap.End()

	var err error

	event, err := c.characterContract.ParsePostNote(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse post note event: %w", err)
	}

	post, note, postOriginal, err := c.buildNoteMetadata(ctx, event.CharacterId, event.NoteId, crossbell.EventNamePostNote)
	if err != nil {
		return nil, fmt.Errorf("failed to build note metadata: %w", err)
	}

	// Comment
	if note.LinkItemType == crossbell.LinkItemTypeNote {
		ethereumClient, err := ethclientx.Global(transfer.Network)
		if err != nil {
			return nil, fmt.Errorf("failed to get ethereum client: %w", err)
		}

		periphery, err := periphery.NewPeriphery(crossbell.AddressPeriphery, ethereumClient)
		if err != nil {
			return nil, fmt.Errorf("failed to get periphery contract: %w", err)
		}

		targetNoteStruct, err := periphery.GetLinkingNote(&bind.CallOpts{}, note.LinkKey)
		if err != nil {
			return nil, fmt.Errorf("failed to get linking note: %w", err)
		}

		targetPost, targetNote, _, err := c.buildNoteMetadata(ctx, targetNoteStruct.CharacterId, targetNoteStruct.NoteId, crossbell.EventNamePostNote /* It may be a comment of a comment, but we can hardly judge it. */)
		if err != nil {
			return nil, fmt.Errorf("failed to build target note metadata: %w", err)
		}

		post.Target = targetPost
		post.Target.TargetURL = targetNote.ContentUri
	}

	transfer.RelatedUrls = []string{note.ContentUri}

	transfer.Platform = c.buildPlatform(postOriginal.Sources)

	if post.Target == nil {
		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialPost, transfer.Type)
	} else {
		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialComment, transfer.Type)
	}

	if transfer.RelatedUrls, err = c.buildRelatedUrls([]string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}, transfer.Platform, event.CharacterId, event.NoteId); err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(post); err != nil {
		return nil, err
	}

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.CharacterId)
	if err != nil {
		return nil, err
	}
	transaction.Owner = strings.ToLower(characterOwner.String())

	return &transfer, nil
}

func (c *characterHandler) buildNoteMetadata(ctx context.Context, characterID, noteID *big.Int, eventName string) (*metadata.Post, *character.DataTypesNote, *CrossbellPostStruct, error) {
	note, err := c.characterContract.GetNote(&bind.CallOpts{}, characterID, noteID)
	if err != nil {
		return nil, nil, nil, err
	}

	contentData, err := ipfs.GetFileByURL(note.ContentUri)
	if err != nil {
		return nil, nil, nil, err
	}

	ethereumClient, err := ethclientx.Global(protocol.NetworkCrossbell)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get ethereum client: %w", err)
	}

	characterContract, err := character.NewCharacter(crossbell.AddressCharacter, ethereumClient)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get character contract: %w", err)
	}

	handle, err := characterContract.GetHandle(&bind.CallOpts{}, characterID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get character handle: %w", err)
	}

	postOriginal := CrossbellPostStruct{}
	if err := json.Unmarshal(contentData, &postOriginal); err != nil {
		return nil, nil, nil, err
	}

	post := metadata.Post{
		TypeOnPlatform: []string{
			eventName,
		},
		Title: postOriginal.Title,
		Body:  postOriginal.Content,
		Author: []string{
			fmt.Sprintf("https://crossbell.io/@%s", handle),
		},
	}

	for _, attachment := range postOriginal.Attachments {
		post.Media = append(post.Media, metadata.Media{
			Address:  attachment.Address,
			MimeType: attachment.MimeType,
		})
	}

	return &post, &note, &postOriginal, nil
}

func (c *characterHandler) handleLinkCharacter(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleLinkCharacter")

	defer snap.End()

	event, err := c.characterContract.ParseLinkCharacter(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, crossbell.AddressCharacter.String(), event.ToCharacterId)
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

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.ToCharacterId)
	if err != nil {
		return nil, err
	}

	transaction.Owner = strings.ToLower(characterOwner.String())
	profile.Address = transaction.Owner

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialFollow, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}

func (c *characterHandler) handleUnLinkCharacter(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleUnLinkCharacter")

	defer snap.End()

	event, err := c.characterContract.ParseLinkCharacter(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, crossbell.AddressCharacter.String(), event.ToCharacterId)
	if err != nil {
		return nil, err
	}

	profile := &social.Profile{
		Address:  strings.ToLower(event.Account.String()),
		Platform: protocol.PlatformCrossbell,
		Network:  transfer.Network,
		Source:   transfer.Network,
		Type:     filter.SocialProfileUpdate,
	}

	if err = BuildProfileMetadata(erc721Token.Metadata, profile); err != nil {
		return nil, err
	}

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.ToCharacterId)
	if err != nil {
		return nil, err
	}

	transaction.Owner = strings.ToLower(characterOwner.String())

	profile.Address = transaction.Owner

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialUnfollow, transfer.Type)
	transfer.RelatedUrls = []string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}

	return &transfer, nil
}

func (c *characterHandler) handleSetCharacterUri(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleSetCharacterUri")

	defer snap.End()

	event, err := c.characterContract.ParseSetCharacterUri(log)
	if err != nil {
		return nil, err
	}

	erc721Token, err := c.tokenClient.ERC721(ctx, protocol.NetworkCrossbell, crossbell.AddressCharacter.String(), event.CharacterId)
	if err != nil {
		return nil, err
	}

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.CharacterId)
	if err != nil {
		return nil, err
	}

	transaction.Owner = strings.ToLower(characterOwner.String())

	profile := &social.Profile{
		Address:  transaction.Owner,
		Platform: protocol.PlatformCrossbell,
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

	if transfer.RelatedUrls, err = c.buildRelatedUrls([]string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}, transfer.Platform, event.CharacterId, nil); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (c *characterHandler) handleSetNoteUri(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleSetNoteUri")

	defer snap.End()

	var err error

	event, err := c.characterContract.ParseSetNoteUri(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SetNoteUri event: %w", err)
	}

	post, _, postOriginal, err := c.buildNoteMetadata(ctx, event.CharacterId, event.NoteId, crossbell.EventNameSetNoteUri)
	if err != nil {
		return nil, fmt.Errorf("build note metadata: %w", err)
	}

	transfer.Platform = c.buildPlatform(postOriginal.Sources)

	if transfer.Metadata, err = json.Marshal(post); err != nil {
		return nil, fmt.Errorf("marshal post: %w", err)
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialRevise, transfer.Type)

	if transfer.RelatedUrls, err = c.buildRelatedUrls([]string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}, transfer.Platform, event.CharacterId, event.NoteId); err != nil {
		return nil, fmt.Errorf("build related urls: %w", err)
	}

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.CharacterId)
	if err != nil {
		return nil, err
	}

	transaction.Owner = strings.ToLower(characterOwner.String())

	return &transfer, nil
}

func (c *characterHandler) handleMintNote(ctx context.Context, transaction *model.Transaction, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handler")

	_, snap := tracer.Start(ctx, "worker_crossbell_handler:handleMintNote")

	defer snap.End()

	var err error

	event, err := c.characterContract.ParseMintNote(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse event: %w", err)
	}

	post, _, postOriginal, err := c.buildNoteMetadata(ctx, event.CharacterId, event.NoteId, crossbell.EventNameMintNote)
	if err != nil {
		return nil, fmt.Errorf("build note metadata: %w", err)
	}

	transfer.Platform = c.buildPlatform(postOriginal.Sources)

	if transfer.Metadata, err = json.Marshal(post); err != nil {
		return nil, fmt.Errorf("failed to marshal post metadata: %w", err)
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialMint, transfer.Type)

	if transfer.RelatedUrls, err = c.buildRelatedUrls([]string{ethereum.BuildScanURL(transfer.Network, transfer.TransactionHash)}, transfer.Platform, event.CharacterId, event.NoteId); err != nil {
		return nil, err
	}

	characterOwner, err := c.characterContract.OwnerOf(&bind.CallOpts{}, event.CharacterId)
	if err != nil {
		return nil, err
	}

	transaction.Owner = strings.ToLower(characterOwner.String())

	return &transfer, nil
}

func (c *characterHandler) buildPlatform(sources []string) string {
	if len(sources) == 0 {
		return protocol.PlatformCrossbell
	}

	// Rewrite typos in sources
	for i, source := range sources {
		switch {
		case strings.EqualFold(source, "xlog"):
			sources[i] = protocol.PlatformCrossbellXLog
		case strings.EqualFold(source, "xcast"):
			sources[i] = protocol.PlatformCrossbellXCast
		}
	}

	return strings.Trim(sources[0], `\"`)
}

func (c *characterHandler) buildRelatedUrls(relatedUrls []string, platform string, characterID, noteID *big.Int) ([]string, error) {
	if noteID == nil {
		handle, err := c.characterContract.GetHandle(&bind.CallOpts{}, characterID)
		if err != nil {
			return nil, err
		}

		return append(relatedUrls, fmt.Sprintf("https://crossbell.io/@%s", handle)), nil
	}

	return append(relatedUrls, fmt.Sprintf("https://crossbell.io/notes/%d-%d", characterID, noteID)), nil
}
