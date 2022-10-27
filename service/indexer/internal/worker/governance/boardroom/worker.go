package boardroom

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/boardroom"
	bdModel "github.com/naturalselectionlabs/pregod/common/worker/boardroom/model"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const (
	Name = "Boardroom"
)

type service struct {
	boardroomClient *boardroom.Client
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("boardroom_worker")
	_, trace := tracer.Start(ctx, "boardroom_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	currTransactions := transactions

	proposals, err := s.boardroomClient.GetProposalsByAddress(ctx, message.Address)
	if err != nil {
		return currTransactions, nil
	}

	if len(proposals) > 0 {
		for _, proposal := range proposals {
			newTransaction, err := s.getProposal(proposal, message)
			if err != nil {
				loggerx.Global().Warn("[boardroom worker] failed to get to formatted Proposal", zap.Error(err))
				continue
			}
			currTransactions = append(currTransactions, *newTransaction)
		}
	}

	votes, err := s.boardroomClient.GetVotesByAddress(ctx, message.Address)
	if err != nil {
		return currTransactions, nil
	}

	if len(votes) > 0 {
		for _, vote := range votes {
			newTransaction, err := s.getVote(vote, message)
			if err != nil {
				loggerx.Global().Warn("[boardroom worker] failed to get to formatted Vote", zap.Error(err))
				continue
			}
			currTransactions = append(currTransactions, *newTransaction)
		}
	}

	return currTransactions, nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func (s *service) getProposal(proposal bdModel.Proposal, message *protocol.Message) (*model.Transaction, error) {
	formattedProposal, err := s.formatProposal(&proposal)
	if err != nil {
		loggerx.Global().Warn("[boardroom worker] failed to format Proposal", zap.Error(err))
		return nil, nil
	}

	rawMetadata, err := json.Marshal(formattedProposal)
	if err != nil {
		loggerx.Global().Warn("[boardroom worker] failed to marshal rawMetadata", zap.Error(err))
		return nil, nil
	}

	var platform string
	var relatedUrls []string

	if strings.EqualFold(proposal.Adapter, protocol.PlatformSnapshot) {
		platform = protocol.PlatformSnapshot
		relatedUrls = append(relatedUrls, proposal.ExternalUrl)
	} else {
		platform = protocol.PlatformOnchain
		url := fmt.Sprintf("https://boardroom.io/gitcoin/proposal/%s", proposal.RefId)
		relatedUrls = append(relatedUrls, url)
	}

	timpstamp, err := strconv.ParseInt(proposal.StartTimestamp, 10, 64)
	if err != nil {
		return nil, err
	}

	currTransaction := model.Transaction{
		Hash:        strings.ToLower(proposal.Id),
		BlockNumber: proposal.BlockNumber,
		Timestamp:   time.Unix(timpstamp, 0),
		Owner:       message.Address,
		AddressFrom: message.Address,
		Platform:    Name,
		Network:     message.Network,
		Source:      s.Name(),
		Tag:         filter.TagGovernance,
		Type:        filter.GovernancePropose,
		Transfers: []model.Transfer{
			{
				TransactionHash: strings.ToLower(proposal.Id),
				Tag:             filter.TagGovernance,
				Type:            filter.GovernancePropose,
				Timestamp:       time.Unix(timpstamp, 0),
				Index:           0, // cannot use -1
				AddressFrom:     message.Address,
				Metadata:        rawMetadata,
				Platform:        platform,
				Network:         message.Network,
				Source:          s.Name(),
				RelatedUrls:     relatedUrls,
			},
		},
	}

	return &currTransaction, nil
}

func (s *service) formatProposal(proposal *bdModel.Proposal) (*metadata.BrProposal, error) {
	sTimpstamp, err := strconv.ParseInt(proposal.StartTimestamp, 10, 64)
	if err != nil {
		return nil, err
	}
	eTimpstamp, err := strconv.ParseInt(proposal.EndTimestamp, 10, 64)
	if err != nil {
		return nil, err
	}
	return &metadata.BrProposal{
		Id:      proposal.Id,
		Title:   proposal.Title,
		Body:    proposal.Content,
		Options: proposal.Choices,
		StartAt: time.Unix(sTimpstamp, 0),
		EndAt:   time.Unix(eTimpstamp, 0),
	}, nil
}

func (s *service) getVote(vote bdModel.Vote, message *protocol.Message) (*model.Transaction, error) {
	proposal, err := s.boardroomClient.GetProposalsByRefId(context.Background(), vote.ProposalRefId)
	if err != nil {
		return nil, err
	}

	if proposal == nil {
		return nil, nil
	}

	formattedVote, err := s.formatVote(&vote, &proposal.Data)
	if err != nil {
		loggerx.Global().Warn("[boardroom worker] failed to format Vote", zap.Error(err))
		return nil, nil
	}

	rawMetadata, err := json.Marshal(formattedVote)
	if err != nil {
		loggerx.Global().Warn("[boardroom worker] failed to marshal rawMetadata", zap.Error(err))
		return nil, nil
	}

	var platform string
	var relatedUrls []string

	if strings.EqualFold(vote.Adapter, protocol.PlatformSnapshot) {
		platform = protocol.PlatformSnapshot
		relatedUrls = append(relatedUrls, proposal.Data.ExternalUrl)
	} else {
		platform = protocol.PlatformOnchain
		url := fmt.Sprintf("https://boardroom.io/gitcoin/proposal/%s", proposal.Data.RefId)
		relatedUrls = append(relatedUrls, url)
	}

	timpstamp, err := strconv.ParseInt(vote.Timestamp, 10, 64)
	if err != nil {
		return nil, err
	}
	currTransaction := model.Transaction{
		Hash:        vote.RefId,
		Timestamp:   time.Unix(timpstamp, 0),
		Owner:       message.Address,
		AddressFrom: message.Address,
		Platform:    s.Name(),
		Network:     message.Network,
		Source:      s.Name(),
		Tag:         filter.TagGovernance,
		Type:        filter.GovernanceVote,
		Transfers: []model.Transfer{
			{
				TransactionHash: vote.RefId,
				Tag:             filter.TagGovernance,
				Type:            filter.GovernanceVote,
				Timestamp:       time.Unix(vote.Time.Timestamp, 0),
				Index:           0, // cannot use -1
				AddressFrom:     message.Address,
				Metadata:        rawMetadata,
				Platform:        platform,
				Network:         message.Network,
				Source:          s.Name(),
				RelatedUrls:     relatedUrls,
			},
		},
	}

	return &currTransaction, nil
}

func (s *service) formatVote(vote *bdModel.Vote, proposal *bdModel.Proposal) (*metadata.BrVote, error) {
	sTimpstamp, err := strconv.ParseInt(proposal.StartTimestamp, 10, 64)
	if err != nil {
		return nil, err
	}
	eTimpstamp, err := strconv.ParseInt(proposal.EndTimestamp, 10, 64)
	if err != nil {
		return nil, err
	}
	brProposal := &metadata.BrProposal{
		Id:      proposal.Id,
		Title:   proposal.Title,
		Body:    proposal.Content,
		Options: proposal.Choices,
		StartAt: time.Unix(sTimpstamp, 0),
		EndAt:   time.Unix(eTimpstamp, 0),
	}

	return &metadata.BrVote{
		Choice:         strconv.Itoa(vote.Choice + 1),
		TypeOnPlatform: []string{filter.GovernanceVote},
		BrProposal:     brProposal,
	}, nil
}

func New() worker.Worker {
	return &service{
		boardroomClient: boardroom.NewClient(),
	}
}
