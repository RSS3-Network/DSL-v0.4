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
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/crossbell/contract"
)

var (
	EventHashTransfer            = common.BytesToHash(crypto.Keccak256([]byte("Transfer(address,address,uint256)")))
	EventHashProfileCreated      = common.BytesToHash(crypto.Keccak256([]byte("ProfileCreated(uint256,address,address,string,uint256)")))
	EventHashSetHandle           = common.BytesToHash(crypto.Keccak256([]byte("SetHandle(address,uint256,string)")))
	EventHashSetPrimaryProfileId = common.BytesToHash(crypto.Keccak256([]byte("SetPrimaryProfileId(address,uint256)")))
	EventHashLinkProfile         = common.BytesToHash(crypto.Keccak256([]byte("LinkProfile(address,uint256,uint256,bytes32,uint256)")))
	EventHashUnlinkProfile       = common.BytesToHash(crypto.Keccak256([]byte("UnlinkProfile(address,uint256,uint256,bytes32)")))
	EventHashLinkAddress         = common.BytesToHash(crypto.Keccak256([]byte("LinkAddress(uint256,address,bytes32,uint256)")))
	EventHashAttachLinklist      = common.BytesToHash(crypto.Keccak256([]byte("AttachLinklist(uint256,uint256,bytes32)")))
	EventHashSetProfileUri       = common.BytesToHash(crypto.Keccak256([]byte("SetProfileUri(uint256,string)")))
)

var _ Interface = (*profile)(nil)

type profile struct {
	contract *contract.ERC721
	abi      abi.ABI
}

func (p *profile) Handle(ctx context.Context, transaction *model.Transaction, log *types.Log) (*model.Transfer, error) {
	switch log.Topics[0] {
	case EventHashTransfer:
		return p.handleTransfer(ctx, transaction, log)
	case EventHashProfileCreated:
		return p.handleProfileCreated(ctx, transaction, log)
	case EventHashSetProfileUri:
	case EventHashSetHandle:
		return p.handleSetHandle(ctx, transaction, log)
	case EventHashSetPrimaryProfileId:
	case EventHashLinkProfile:
	case EventHashUnlinkProfile:
	case EventHashLinkAddress:
	case EventHashAttachLinklist:
	default:
		return nil, ErrorUnknownUnknownEvent
	}
}

func (p *profile) handleTransfer(ctx context.Context, transaction *model.Transaction, log *types.Log) (*model.Transfer, error) {
	fromAddress := common.HexToAddress(log.Topics[1].Hex())
	toAddress := common.HexToAddress(log.Topics[2].Hex())

	profileID := big.NewInt(0)
	profileID.SetString(log.Topics[3].Hex(), 0)

	return &model.Transfer{
		Type:        EventNameTransfer,
		AddressFrom: fromAddress.String(),
		AddressTo:   toAddress.String(),
		Metadata:    metadataRaw,
	}, nil
}

func (p *profile) handleProfileCreated(ctx context.Context, transaction *model.Transaction, log *types.Log) (*model.Transfer, error) {
	profileCreatedData := contract.ProfileCreatedData{}

	if err := w.abiWeb3EntryProfile.UnpackIntoInterface(&profileCreatedData, EventNameProfileCreated, log.Data); err != nil {
		return nil, err
	}

	profileID := big.NewInt(0)
	profileID.SetString(log.Topics[1].Hex(), 0)

	creatorAddress := common.HexToAddress(log.Topics[2].Hex())
	toAddress := common.HexToAddress(log.Topics[3].Hex())

	tokenMetadataRaw, err := w.getTokenMetadata(ctx, profileID)
	if err != nil {
		return nil, err
	}

	metadataRaw, err := json.Marshal(&metadata.CrossBell{
		Event:          EventNameProfileCreated,
		ProfileID:      profileID,
		AddressCreator: creatorAddress.String(),
		AddressOwner:   toAddress.String(),
		Metadata:       tokenMetadataRaw,
	})
	if err != nil {
		return nil, err
	}

	return &model.Transfer{
		Type:     EventNameProfileCreated,
		Metadata: metadataRaw,
	}, nil
}

func (p *profile) handleSetHandle(ctx context.Context, transaction *model.Transaction, log *types.Log) (*model.Transfer, error) {
	account := common.HexToAddress(log.Topics[1].Hex())

	profileID := big.NewInt(0)
	profileID.SetString(log.Topics[2].Hex(), 0)

	setHandlerData := contract.SetHandle{}

	if err := w.abiWeb3EntryProfile.UnpackIntoInterface(&setHandlerData, EventNameSetHandle, log.Data); err != nil {
		return nil, err
	}

	metadataRaw, err := w.getTokenMetadata(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return &metadata.CrossBell{
		Event:        EventNameSetHandle,
		ProfileID:    profileID,
		Handle:       setHandlerData.NewHandle,
		AddressOwner: account.String(),
		Metadata:     metadataRaw,
	}, nil
}
