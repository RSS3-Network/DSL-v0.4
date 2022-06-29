package handler

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/nft"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
	"go.opentelemetry.io/otel"
)

var (
	EventNameProfileCreated      = "ProfileCreated"
	EventNameSetHandle           = "SetHandle"
	EventNameSetPrimaryProfileId = "SetPrimaryProfileId"
	EventNameLinkProfile         = "LinkProfile"
	EventNameUnlinkProfile       = "UnlinkProfile"
	EventNameAttachLinklist      = "AttachLinkList"
	EventNameSetProfileUri       = "SetProfileUri"
	EventNamePostNote            = "PostNote"

	EventHashProfileCreated      = common.BytesToHash(crypto.Keccak256([]byte("ProfileCreated(uint256,address,address,string,uint256)")))
	EventHashSetHandle           = common.BytesToHash(crypto.Keccak256([]byte("SetHandle(address,uint256,string)")))
	EventHashSetPrimaryProfileId = common.BytesToHash(crypto.Keccak256([]byte("SetPrimaryProfileId(address,uint256)")))
	EventHashLinkProfile         = common.BytesToHash(crypto.Keccak256([]byte("LinkProfile(address,uint256,uint256,bytes32,uint256)")))
	EventHashUnlinkProfile       = common.BytesToHash(crypto.Keccak256([]byte("UnlinkProfile(address,uint256,uint256,bytes32)")))
	EventHashAttachLinklist      = common.BytesToHash(crypto.Keccak256([]byte("AttachLinklist(uint256,uint256,bytes32)")))
	EventHashSetProfileUri       = common.BytesToHash(crypto.Keccak256([]byte("SetProfileUri(uint256,string)")))
	EventHashPostNote            = common.BytesToHash(crypto.Keccak256([]byte("PostNote(uint256,uint256,bytes32,bytes32,bytes)")))

	LinkTypeFollow  = "follow"
	LinkTypeLike    = "like"
	LinkTypeComment = "comment"

	LinkTypeMap = map[common.Hash]string{
		common.BytesToHash(common.RightPadBytes([]byte(LinkTypeFollow), common.HashLength)):  LinkTypeFollow,
		common.BytesToHash(common.RightPadBytes([]byte(LinkTypeLike), common.HashLength)):    LinkTypeLike,
		common.BytesToHash(common.RightPadBytes([]byte(LinkTypeComment), common.HashLength)): LinkTypeComment,
	}
)

var _ Interface = (*profile)(nil)

type profile struct {
	contract *contract.ERC721
	abi      abi.ABI
}

func (p *profile) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	tracer := otel.Tracer("crossbell_handle_profile")
	ctx, trace := tracer.Start(ctx, "Handler")

	defer trace.End()

	var log types.Log

	if err := json.Unmarshal(transfer.SourceData, &log); err != nil {
		return nil, err
	}

	switch log.Topics[0] {
	case EventHashProfileCreated:
		return p.handleProfileCreated(ctx, transfer, log)
	case EventHashSetPrimaryProfileId:
		return p.handleSetPrimaryProfileId(ctx, transfer, log)
	case EventHashSetHandle:
		return p.handleSetHandle(ctx, transfer, log)
	case EventHashLinkProfile:
		return p.handleLinkProfile(ctx, transfer, log)
	case EventHashUnlinkProfile:
		return p.handleUnlinkProfile(ctx, transfer, log)
	case EventHashAttachLinklist:
		return p.handleAttachLinkList(ctx, transfer, log)
	case EventHashSetProfileUri:
		return p.handleSetProfileUri(ctx, transfer, log)
	case EventHashPostNote:
		return p.handlePostNote(ctx, transfer, log)
	default:
		return nil, ErrorUnknownUnknownEvent
	}
}

func (p *profile) handleProfileCreated(ctx context.Context, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tracer := otel.Tracer("crossbell_handle_profile")
	_, trace := tracer.Start(ctx, "handleProfileCreated")

	defer trace.End()

	profileCreated := contract.ProfileCreated{}

	if err := p.abi.UnpackIntoInterface(&profileCreated, EventNameProfileCreated, log.Data); err != nil {
		return nil, err
	}

	profileID := big.NewInt(0)
	profileID.SetString(log.Topics[1].Hex(), 0)

	addressCreator := common.HexToAddress(log.Topics[2].Hex())
	addressOwner := common.HexToAddress(log.Topics[3].Hex())

	metadataSelf, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, profileID)
	if err != nil {
		return nil, err
	}

	transfer.Type = EventNameProfileCreated
	transfer.AddressFrom = addressCreator.String()
	transfer.AddressTo = addressOwner.String()

	if transfer.Metadata, err = json.Marshal(&metadata.Metadata{
		Crossbell: &metadata.Crossbell{
			TokenID:  profileID,
			Metadata: metadataSelf,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (p *profile) handleSetHandle(ctx context.Context, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	profileID := big.NewInt(0)
	profileID.SetString(log.Topics[2].Hex(), 0)

	setHandlerData := contract.SetHandle{}

	if err := p.abi.UnpackIntoInterface(&setHandlerData, EventNameSetHandle, log.Data); err != nil {
		return nil, err
	}

	metadataSelf, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, profileID)
	if err != nil {
		return nil, err
	}

	transfer.Type = EventNameSetHandle

	if transfer.Metadata, err = json.Marshal(&metadata.Metadata{
		Crossbell: &metadata.Crossbell{
			TokenID:  profileID,
			Handle:   setHandlerData.NewHandle,
			Metadata: metadataSelf,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (p *profile) handleLinkProfile(ctx context.Context, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tokenIDFrom := big.NewInt(0)
	tokenIDFrom.SetString(log.Topics[2].Hex(), 0)

	tokenIDTo := big.NewInt(0)
	tokenIDTo.SetString(log.Topics[3].Hex(), 0)

	linkProfile := contract.LinkProfile{}
	if err := p.abi.UnpackIntoInterface(&linkProfile, EventNameLinkProfile, log.Data); err != nil {
		return nil, err
	}

	transfer.Type = EventNameLinkProfile

	metadataFrom, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, tokenIDFrom)
	if err != nil {
		return nil, err
	}

	metadataTo, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, tokenIDTo)
	if err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(&metadata.Metadata{
		Crossbell: &metadata.Crossbell{
			TokenIDFrom:  tokenIDFrom,
			TokenIDTo:    tokenIDTo,
			LinkType:     LinkTypeMap[linkProfile.LinkType],
			MetadataFrom: metadataFrom,
			MetadataTo:   metadataTo,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (p *profile) handleUnlinkProfile(ctx context.Context, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	tokenIDFrom := big.NewInt(0)
	tokenIDFrom.SetString(log.Topics[2].Hex(), 0)

	tokenIDTo := big.NewInt(0)
	tokenIDTo.SetString(log.Topics[3].Hex(), 0)

	unlinkProfile := contract.UnlinkProfile{}
	if err := p.abi.UnpackIntoInterface(&unlinkProfile, EventNameUnlinkProfile, log.Data); err != nil {
		return nil, err
	}

	transfer.Type = EventNameUnlinkProfile

	metadataFrom, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, tokenIDFrom)
	if err != nil {
		return nil, err
	}

	metadataTo, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, tokenIDTo)
	if err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(&metadata.Metadata{
		Crossbell: &metadata.Crossbell{
			TokenIDFrom:  tokenIDFrom,
			TokenIDTo:    tokenIDTo,
			LinkType:     LinkTypeMap[unlinkProfile.LinkType],
			MetadataFrom: metadataFrom,
			MetadataTo:   metadataTo,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (p *profile) handlePostNote(ctx context.Context, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	profileID := big.NewInt(0)
	profileID.SetString(log.Topics[1].Hex(), 0)

	noteID := big.NewInt(0)
	noteID.SetString(log.Topics[2].Hex(), 0)

	transfer.Type = filter.SocialPost

	profileMetadata, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, profileID)
	if err != nil {
		return nil, err
	}

	noteMetadata, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, noteID)
	if err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(&metadata.Metadata{
		Crossbell: &metadata.Crossbell{
			TokenIDFrom:  profileID,
			TokenIDTo:    noteID,
			MetadataFrom: profileMetadata,
			MetadataTo:   noteMetadata,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (p *profile) handleSetPrimaryProfileId(ctx context.Context, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	transfer.Type = EventNameSetPrimaryProfileId

	profileIDTo := big.NewInt(0)
	profileIDTo.SetString(log.Topics[2].Hex(), 0)

	profileIDFrom := big.NewInt(0)
	profileIDFrom.SetString(log.Topics[3].Hex(), 0)

	metadataFrom, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, profileIDFrom)
	if err != nil {
		return nil, err
	}

	metadataTo, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, profileIDTo)
	if err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(&metadata.Metadata{
		Crossbell: &metadata.Crossbell{
			ProfileIDFrom: profileIDFrom,
			ProfileIDTo:   profileIDTo,
			MetadataFrom:  metadataFrom,
			MetadataTo:    metadataTo,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (p *profile) handleAttachLinkList(ctx context.Context, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	transfer.Type = EventNameAttachLinklist

	linklistID := big.NewInt(0)
	linklistID.SetString(log.Topics[1].Hex(), 0)

	profileID := big.NewInt(0)
	profileID.SetString(log.Topics[2].Hex(), 0)

	linkType := LinkTypeMap[log.Topics[3]]

	profileMetadata, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, profileID)
	if err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(&metadata.Metadata{
		Crossbell: &metadata.Crossbell{
			ProfileID:  profileID,
			LinkListID: linklistID,
			LinkType:   linkType,
			Metadata:   profileMetadata,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (p *profile) handleSetProfileUri(ctx context.Context, transfer model.Transfer, log types.Log) (*model.Transfer, error) {
	transfer.Type = EventNameSetProfileUri

	profileID := big.NewInt(0)
	profileID.SetString(log.Topics[1].Hex(), 0)

	setProfileUri := contract.SetProfileUri{}
	if err := p.abi.UnpackIntoInterface(&setProfileUri, EventNameSetProfileUri, log.Data); err != nil {
		return nil, err
	}

	profileMetadata, err := nft.GetMetadata(nft.NetworkCrossbell, AddressProfileProxy, profileID)
	if err != nil {
		return nil, err
	}

	if transfer.Metadata, err = json.Marshal(&metadata.Metadata{
		Crossbell: &metadata.Crossbell{
			TokenID:  profileID,
			URI:      setProfileUri.NewUri,
			Metadata: profileMetadata,
		},
	}); err != nil {
		return nil, err
	}

	return &transfer, nil
}
