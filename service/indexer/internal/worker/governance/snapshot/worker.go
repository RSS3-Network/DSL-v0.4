package snapshot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model/governance"

	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/snapshot"
	job2 "github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/governance/snapshot/job"

	"github.com/shopspring/decimal"

	mapset "github.com/deckarep/golang-set"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
)

const (
	Name = protocol.PlatfromSnapshot
)

// Need to go to the official website to get the corresponding network key
// to correspond to the corresponding network processing.
// https://github.com/snapshot-labs/snapshot.js/blob/master/src/networks.json
var snapshotNetworkNumMap = map[string]string{
	protocol.NetworkEthereum:          "1",
	protocol.NetworkEthereumClassic:   "61",
	protocol.NetworkBinanceSmartChain: "56",
	protocol.NetworkPolygon:           "137",
	protocol.NetworkXDAI:              "100",
	protocol.NetworkArbitrum:          "42161",
	protocol.NetworkOptimism:          "10",
	protocol.NetworkFantom:            "250",
	protocol.NetworkAvalanche:         "43113",
}

type service struct {
	databaseClient *gorm.DB
	redisClient    *redis.Client
	snapshotClient *snapshot.Client
}

type MetadataSpaceFilter struct {
	MinScore    decimal.Decimal `json:"minScore"`
	OnlyMembers bool            `json:"onlyMembers"`
}

type MetadataSpace struct {
	Id      string              `json:"id"`
	Name    string              `json:"name"`
	About   string              `json:"about"`
	Network string              `json:"network"`
	Symbol  string              `json:"symbol"`
	Filters MetadataSpaceFilter `json:"filters"`
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkEthereumClassic,
		protocol.NetworkBinanceSmartChain,
		protocol.NetworkPolygon,
		protocol.NetworkXDAI,
		protocol.NetworkArbitrum,
		protocol.NetworkOptimism,
		protocol.NetworkFantom,
		protocol.NetworkAvalanche,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("snapshot_worker")
	_, trace := tracer.Start(ctx, "snapshot_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	// Only some mainnets are currently supported
	snapshotNetworkNum := snapshotNetworkNumMap[message.Network]

	timeStamp, err := s.getLatestTimestamp(message)
	if err != nil {
		logrus.Errorf("failed to get latest timestamp: %s", err)
		return nil, err
	}

	currTransactions := transactions

	isVoteError := false
	votes, err := s.getSnapshotVotes(ctx, message.Address, timeStamp)
	if err != nil {
		logrus.Errorf("failed to get snapshot votes: %s", err)
		isVoteError = true
	}

	isProposalsByAuthorError := false
	proposalsByAuthorMap, err := s.getProposalsByAuthor(ctx, message.Address, timeStamp)
	if err != nil {
		logrus.Errorf("failed to get proposals by author: %s", err)
		isProposalsByAuthorError = true
	}

	if len(votes) == 0 && len(proposalsByAuthorMap) == 0 {
		return nil, nil
	}

	if isVoteError && isProposalsByAuthorError {
		return nil, errors.New("failed to get snapshot votes and proposals by author")
	}

	proposalIDSet := mapset.NewSet()
	var proposalIDs []string
	spaceIDSet := mapset.NewSet()
	var spaceIDs []string

	for _, vote := range votes {
		proposalIDSet.Add(vote.ProposalID)
		spaceIDSet.Add(vote.SpaceID)
	}

	for _, proposal := range proposalsByAuthorMap {
		proposalIDSet.Add(proposal.ID)
		spaceIDSet.Add(proposal.SpaceID)
	}

	for _, proposalNode := range proposalIDSet.ToSlice() {
		proposal, ok := proposalNode.(string)
		if !ok {
			logrus.Warnf("[snapshot worker] failed to convert proposal node to snapshot proposal:%v", proposalNode)
		}
		proposalIDs = append(proposalIDs, proposal)
	}

	for _, spaceNode := range spaceIDSet.ToSlice() {
		space, ok := spaceNode.(string)
		if !ok {
			logrus.Warnf("[snapshot worker] failed to convert space node to snapshot space:%v", spaceNode)
		}
		spaceIDs = append(spaceIDs, space)
	}

	proposalMap, err := s.getSnapshotProposals(ctx, proposalIDs)
	if err != nil {
		return nil, fmt.Errorf("[snapshot worker] failed to get snapshot proposals: %w", err)
	}

	spaceMap, err := s.getSnapshotSpaces(ctx, spaceIDs, snapshotNetworkNum)
	if err != nil {
		return nil, fmt.Errorf("[snapshot worker] failed to get snapshot spaces: %w", err)
	}

	if !isVoteError {
		for _, vote := range votes {
			newTransaction, err := s.getVote(vote, proposalMap, spaceMap, message)
			if err != nil {
				logrus.Errorf("failed to format vote: %s", err)
				continue
			}
			if newTransaction != nil {
				currTransactions = append(currTransactions, *newTransaction)
			}
		}
	}

	if !isProposalsByAuthorError {
		for _, proposal := range proposalsByAuthorMap {
			newTransaction, err := s.getProposal(proposal, spaceMap, message)
			if err != nil {
				logrus.Errorf("failed to format proposal: %s", err)
				continue
			}
			if newTransaction != nil {
				currTransactions = append(currTransactions, *newTransaction)
			}
		}
	}

	return currTransactions, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&job2.SnapshotSpaceJob{
			SnapshotJobBase: job2.SnapshotJobBase{
				Name:           "snapshot_space_job",
				DatabaseClient: s.databaseClient,
				RedisClient:    s.redisClient,
				SnapshotClient: s.snapshotClient,
				Limit:          1000,
				HighUpdateTime: time.Second,
				LowUpdateTime:  time.Minute * 5,
			},
		},
		&job2.SnapshotProposalJob{
			SnapshotJobBase: job2.SnapshotJobBase{
				Name:           "snapshot_proposal_job",
				DatabaseClient: s.databaseClient,
				RedisClient:    s.redisClient,
				SnapshotClient: s.snapshotClient,
				Limit:          2000,
				HighUpdateTime: time.Second,
				LowUpdateTime:  time.Minute * 5,
			},
		},
		&job2.SnapshotVoteJob{
			SnapshotJobBase: job2.SnapshotJobBase{
				Name:           "snapshot_vote_job",
				DatabaseClient: s.databaseClient,
				RedisClient:    s.redisClient,
				SnapshotClient: s.snapshotClient,
				Limit:          10000,
				HighUpdateTime: time.Second * 15,
				LowUpdateTime:  time.Minute * 5,
			},
		},
	}
}

func (s *service) getLatestTimestamp(message *protocol.Message) (time.Time, error) {
	var timestamp time.Time

	if err := s.databaseClient.
		Model((*model.Transaction)(nil)).
		Select("COALESCE(timestamp, 'epoch'::timestamp) AS timestamp").
		Where(map[string]interface{}{
			"address_from": message.Address,
			"network":      message.Network,
			"source":       s.Name(),
		}).
		Order("timestamp DESC").
		Limit(1).
		Pluck("timestamp", &timestamp).
		Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return time.Time{}, err
	}

	return timestamp, nil
}

func (s *service) getSnapshotVotes(ctx context.Context, address string, timestamp time.Time) (snapshotVotes []governance.SnapshotVote, err error) {
	tracer := otel.Tracer("snapshot_worker")
	_, trace := tracer.Start(ctx, "snapshot_worker:getSnapshotVotes")

	defer func() { opentelemetry.Log(trace, address, snapshotVotes, err) }()

	// from db
	if err := s.databaseClient.
		Model(&governance.SnapshotVote{}).
		Where("voter = ?", strings.ToLower(address)).
		Where("date_created >= ?", timestamp).
		Find(&snapshotVotes).Error; err != nil {
		return nil, err
	}

	return snapshotVotes, nil
}

func (s *service) getSnapshotProposals(ctx context.Context, proposals []string) (snapshotProposalMap map[string]governance.SnapshotProposal, err error) {
	tracer := otel.Tracer("snapshot_worker")
	_, trace := tracer.Start(ctx, "snapshot_worker:getSnapshotProposals")

	defer func() { opentelemetry.Log(trace, proposals, snapshotProposalMap, err) }()

	var snapshotProposals []governance.SnapshotProposal
	snapshotProposalMap = make(map[string]governance.SnapshotProposal)

	// from db
	if err := s.databaseClient.
		Model(&governance.SnapshotProposal{}).
		Where("id IN (?)", proposals).
		Find(&snapshotProposals).Error; err != nil {
		return nil, err
	}

	for _, proposal := range snapshotProposals {
		snapshotProposalMap[proposal.ID] = proposal
	}

	return snapshotProposalMap, nil
}

func (s *service) getSnapshotSpaces(ctx context.Context, spaces []string, networkNum string) (snapshotSpaceMap map[string]governance.SnapshotSpace, err error) {
	tracer := otel.Tracer("snapshot_worker")
	_, trace := tracer.Start(ctx, "snapshot_worker:getSnapshotSpaces")

	defer func() { opentelemetry.Log(trace, spaces, snapshotSpaceMap, err) }()

	var snapshotSpaces []governance.SnapshotSpace
	snapshotSpaceMap = make(map[string]governance.SnapshotSpace)

	// from db
	if err := s.databaseClient.
		Model(&governance.SnapshotSpace{}).
		Where("id in (?)", spaces).
		Find(&snapshotSpaces).Error; err != nil {
		return nil, err
	}

	for _, space := range snapshotSpaces {
		snapshotSpaceMap[space.ID] = space
	}

	return snapshotSpaceMap, nil
}

func (s *service) getProposalsByAuthor(ctx context.Context, author string, timestamp time.Time) (snapshotProposalMap map[string]governance.SnapshotProposal, err error) {
	tracer := otel.Tracer("snapshot_worker")
	_, trace := tracer.Start(ctx, "snapshot_worker:getProposalsByAuthor")

	defer func() { opentelemetry.Log(trace, author, snapshotProposalMap, err) }()

	var snapshotProposals []governance.SnapshotProposal
	snapshotProposalMap = make(map[string]governance.SnapshotProposal)

	// from db
	if err := s.databaseClient.
		Model(&governance.SnapshotProposal{}).
		Where("author IN (?)", author).
		Where("date_created >= ?", timestamp).
		Find(&snapshotProposals).Error; err != nil {
		return nil, err
	}

	for _, proposal := range snapshotProposals {
		snapshotProposalMap[proposal.ID] = proposal
	}

	return snapshotProposalMap, nil
}

func (s *service) getVote(
	vote governance.SnapshotVote,
	proposalMap map[string]governance.SnapshotProposal,
	spaceMap map[string]governance.SnapshotSpace,
	message *protocol.Message,
) (*model.Transaction, error) {
	var metadataModel metadata.Metadata

	proposal, ok := proposalMap[vote.ProposalID]
	if !ok {
		logrus.Warnf("[snapshot worker] failed to get proposal:%v", vote.ProposalID)
		return nil, nil
	}

	space, ok := spaceMap[vote.SpaceID]
	if !ok {
		logrus.Warnf("[snapshot worker] failed to get space:%v, network:%v", vote.SpaceID, message.Network)

		return nil, nil
	}

	formattedProposal, err := s.formatProposal(&proposal, &space)
	if err != nil {
		logrus.Warnf("[snapshot worker] failed to format proposal:%v", proposal.ID)
		return nil, nil
	}

	metadataModel.Vote = &metadata.Vote{
		TypeOnPlatform: []string{filter.GovernanceVote},
		Choice:         string(vote.Choice),
		Proposal:       formattedProposal,
	}

	rawMetadata, err := json.Marshal(metadataModel)
	if err != nil {
		logrus.Warnf("[snapshot worker] failed to marshal metadata:%v", err)
		return nil, nil
	}

	snapShotSourceData := &metadata.SnapShot{
		Proposal: proposal.Metadata,
		Space:    space.Metadata,
		Choice:   vote.Choice,
	}

	sourceData, err := json.Marshal(snapShotSourceData)
	if err != nil {
		logrus.Warnf("[snapshot worker] failed to marshal sourcedata:%v", err)
		return nil, nil
	}

	relatedUrl := "https://snapshot.org/#/" + vote.SpaceID + "/proposal/" + vote.ProposalID
	lowerAddress := message.Address

	currTransaction := model.Transaction{
		Hash:        vote.ID,
		Timestamp:   vote.DateCreated,
		AddressFrom: lowerAddress,
		Platform:    Name,
		Network:     message.Network,
		Source:      Name,
		SourceData:  sourceData,
		Tag:         filter.TagGovernance,
		Transfers: []model.Transfer{
			{
				TransactionHash: vote.ID,
				Tag:             filter.TagGovernance,
				Type:            filter.GovernanceVote,
				Timestamp:       vote.DateCreated,
				Index:           protocol.IndexVirtual,
				AddressFrom:     lowerAddress,
				Metadata:        rawMetadata,
				Platform:        Name,
				Network:         message.Network,
				Source:          Name,
				SourceData:      sourceData,
				RelatedUrls:     []string{relatedUrl},
			},
		},
	}

	return &currTransaction, nil
}

func (s *service) getProposal(
	proposal governance.SnapshotProposal,
	spaceMap map[string]governance.SnapshotSpace,
	message *protocol.Message,
) (*model.Transaction, error) {
	var metadataModel metadata.Metadata

	space, ok := spaceMap[proposal.SpaceID]
	if !ok {
		logrus.Warnf("[snapshot worker] failed to get space:%v, network:%v", proposal.SpaceID, message.Network)
		return nil, nil
	}

	formattedProposal, err := s.formatProposal(&proposal, &space)
	if err != nil {
		logrus.Warnf("[snapshot worker] failed to format proposal:%v", proposal.ID)
		return nil, nil
	}

	metadataModel.Proposal = formattedProposal

	rawMetadata, err := json.Marshal(metadataModel)
	if err != nil {
		logrus.Warnf("[snapshot worker] failed to marshal metadata:%v", err)
		return nil, nil
	}

	snapShotSourceData := metadata.SnapShot{
		Proposal: proposal.Metadata,
		Space:    space.Metadata,
	}

	sourceData, err := json.Marshal(snapShotSourceData)
	if err != nil {
		logrus.Warnf("[snapshot worker] failed to marshal sourcedata:%v", err)
		return nil, nil
	}

	relatedUrl := "https://snapshot.org/#/" + proposal.SpaceID + "/proposal/" + proposal.ID

	currTransaction := model.Transaction{
		Hash:        proposal.ID,
		Timestamp:   proposal.DateCreated,
		AddressFrom: message.Address,
		Platform:    Name,
		Network:     message.Network,
		Source:      s.Name(),
		SourceData:  sourceData,
		Tag:         filter.TagGovernance,
		Transfers: []model.Transfer{
			{
				TransactionHash: proposal.ID,
				Tag:             filter.TagGovernance,
				Type:            filter.GovernancePropose,
				Timestamp:       proposal.DateCreated,
				Index:           protocol.IndexVirtual,
				AddressFrom:     message.Address,
				Metadata:        rawMetadata,
				Platform:        Name,
				Network:         message.Network,
				Source:          s.Name(),
				SourceData:      sourceData,
				RelatedUrls:     []string{relatedUrl},
			},
		},
	}

	return &currTransaction, nil
}

// formatProposal format governance.SnapshotProposal into metadata.Proposal
func (s *service) formatProposal(
	proposal *governance.SnapshotProposal, space *governance.SnapshotSpace,
) (*metadata.Proposal, error) {
	var proposalDetail Proposal

	err := json.Unmarshal(proposal.Metadata, &proposalDetail)
	if err != nil {
		return nil, err
	}

	var spaceDetail Space

	err = json.Unmarshal(space.Metadata, &spaceDetail)
	if err != nil {
		return nil, err
	}

	return &metadata.Proposal{
		Id:      proposalDetail.Id,
		Title:   proposalDetail.Title,
		Body:    proposalDetail.Body,
		Options: proposalDetail.Choices,
		StartAt: time.Unix(int64(proposalDetail.Start), 0),
		EndAt:   time.Unix(int64(proposalDetail.End), 0),
		Organization: &metadata.Organization{
			Id:    spaceDetail.Id,
			Name:  spaceDetail.Name,
			About: spaceDetail.About,
		},
	}, nil
}

func New(
	databaseClient *gorm.DB,
	redisClient *redis.Client,
) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		redisClient:    redisClient,
		snapshotClient: snapshot.NewClient(),
	}
}
