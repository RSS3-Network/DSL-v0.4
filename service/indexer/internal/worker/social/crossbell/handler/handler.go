package handler

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/character"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/linklist"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/periphery"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/social/crossbell/contract/profile"
	"go.opentelemetry.io/otel"
)

type Interface interface {
	Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error)
}

var _ Interface = (*handler)(nil)

type handler struct {
	characterHandler Interface
	linkListHandler  Interface
}

type CrossbellProfileStruct struct {
	Type             string   `json:"type"`
	Avatars          []string `json:"avatars"`
	ConnectedAvatars []string `json:"connected_avatars"`
	Name             string   `json:"name"`
	Bio              string   `json:"bio"`
}

type CrossbellPostStruct struct {
	Tags        []string `json:"tags"`
	Authors     []string `json:"authors"`
	Content     string   `json:"content"`
	Attachments []struct {
		Address     string `json:"address"`
		MimeType    string `json:"mime_type"`
		SizeInBytes int    `json:"size_in_bytes"`
	} `json:"attachments"`
	DatePublished string   `json:"date_published"`
	Title         string   `json:"title"`
	Sources       []string `json:"sources"`

	// special fields for xLog
	ExternalUrls []string `json:"external_urls"`
	Summary      string   `json:"summary"`
}

func (h *handler) Handle(ctx context.Context, transaction model.Transaction, transfer model.Transfer) (*model.Transfer, error) {
	tracer := otel.Tracer("worker_crossbell_handle")

	_, span := tracer.Start(ctx, "worker_crossbell_handle:Handle")

	defer span.End()

	// Transfer type actions should be handled by the transaction worker
	if transfer.Source != protocol.SourceOrigin {
		return &transfer, nil
	}

	var log types.Log

	if err := json.Unmarshal(transfer.SourceData, &log); err != nil {
		return nil, err
	}

	switch common.HexToAddress(transaction.AddressTo) {
	case contract.AddressCharacter:
		return h.characterHandler.Handle(ctx, transaction, transfer)
	case contract.AddressLinkList:
		return h.linkListHandler.Handle(ctx, transaction, transfer)
	default:
		return nil, contract.ErrorUnknownContractAddress
	}
}

func New(ethereumClient *ethclient.Client) (Interface, error) {
	profileContract, err := profile.NewProfile(contract.AddressCharacter, ethereumClient)
	if err != nil {
		return nil, err
	}

	characterContract, err := character.NewCharacter(contract.AddressCharacter, ethereumClient)
	if err != nil {
		return nil, err
	}

	peripheryContract, err := periphery.NewPeriphery(contract.AddressPeriphery, ethereumClient)
	if err != nil {
		return nil, err
	}

	linkListContract, err := linklist.NewLinkList(contract.AddressLinkList, ethereumClient)
	if err != nil {
		return nil, err
	}

	tokenClient := token.New()

	return &handler{
		characterHandler: &characterHandler{
			characterContract: characterContract,
			peripheryContract: peripheryContract,
			profileHandler: &profileHandler{
				profileContract: profileContract,
				tokenClient:     tokenClient,
			},
			tokenClient: tokenClient,
		},
		linkListHandler: &linkListHandler{
			linkListContract: linkListContract,
		},
	}, nil
}

func BuildProfileMetadata(profileMetadata []byte, profile *social.Profile) error {
	tempStructure := &CrossbellProfileStruct{}
	if err := json.Unmarshal(profileMetadata, &tempStructure); err != nil {
		return err
	}

	profile.Name = tempStructure.Name
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
	return nil
}
