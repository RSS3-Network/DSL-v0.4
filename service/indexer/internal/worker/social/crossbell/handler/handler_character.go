package handler

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"

	"github.com/naturalselectionlabs/pregod/common/datasource/ipfs"
	"github.com/naturalselectionlabs/pregod/common/datasource/nft"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/character"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/periphery"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"go.opentelemetry.io/otel"
)

var _ Interface = (*characterHandler)(nil)

type characterHandler struct {
	characterContract *character.Character
	peripheryContract *periphery.Periphery
	profileHandler    *profileHandler
	databaseClient    *gorm.DB
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
	characterMetadata, _ := nft.GetMetadata(protocol.NetworkCrossbell, contract.AddressCharacter, event.CharacterId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNameCharacterCreated,
		Character: &metadata.CrossbellCharacter{
			ID:       event.CharacterId,
			Metadata: characterMetadata,
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)

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

	characterMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.CharacterId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNameSetHandle,
		Character: &metadata.CrossbellCharacter{
			ID:       event.CharacterId,
			Metadata: characterMetadata,
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)

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

	// Self-hosted IPFS files may be out of date
	contentData, _ := ipfs.GetFileByURL(ctx, note.ContentUri)

	var noteMetadata json.RawMessage

	if err := json.Unmarshal(contentData, &noteMetadata); err != nil {
		return nil, err
	}

	uri, err := c.peripheryContract.GetLinkingAnyUri(&bind.CallOpts{}, event.LinkKey)
	if err != nil {
		return nil, err
	}

	characterMetadata, _ := nft.GetMetadata(protocol.NetworkCrossbell, contract.AddressCharacter, event.NoteId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNamePostNote,
		Character: &metadata.CrossbellCharacter{
			ID:       event.CharacterId,
			Metadata: characterMetadata,
		},
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

	fromCharacterMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.FromCharacterId)
	toCharacterMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.ToCharacterId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNameLinkCharacter,
		Link: &metadata.CrossbellLink{
			Type: contract.LinkTypeMap[event.LinkType],
			From: &metadata.CrossbellCharacter{
				ID:       event.FromCharacterId,
				Metadata: fromCharacterMetadata,
			},
			To: &metadata.CrossbellCharacter{
				ID:       event.ToCharacterId,
				Metadata: toCharacterMetadata,
			},
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialFollow, transfer.Type)

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

	fromCharacterMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.FromCharacterId)
	toCharacterMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.ToCharacterId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNameUnlinkCharacter,
		Link: &metadata.CrossbellLink{
			Type: contract.LinkTypeMap[event.LinkType],
			From: &metadata.CrossbellCharacter{
				ID:       event.FromCharacterId,
				Metadata: fromCharacterMetadata,
			},
			To: &metadata.CrossbellCharacter{
				ID:       event.ToCharacterId,
				Metadata: toCharacterMetadata,
			},
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialUnfollow, transfer.Type)

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

	characterMetadata, _ := nft.GetMetadata(protocol.PlatfromCrossbell, contract.AddressCharacter, event.CharacterId)

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
		Event: contract.EventNameSetCharacterUri,
		Character: &metadata.CrossbellCharacter{
			ID:       event.CharacterId,
			URI:      event.NewUri,
			Metadata: characterMetadata,
		},
	}); err != nil {
		return nil, err
	}

	transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagSocial, transfer.Tag, filter.SocialProfile, transfer.Type)

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
	contentData, _ := ipfs.GetFileByURL(ctx, note.ContentUri)

	var noteMetadata json.RawMessage

	if err := json.Unmarshal(contentData, &noteMetadata); err != nil {
		return nil, err
	}

	uri, err := c.peripheryContract.GetLinkingAnyUri(&bind.CallOpts{}, note.LinkKey)
	if err != nil {
		return nil, err
	}

	if transfer.Metadata, err = metadata.BuildMetadataRawMessage(transfer.Metadata, &metadata.Crossbell{
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

	return &transfer, nil
}
