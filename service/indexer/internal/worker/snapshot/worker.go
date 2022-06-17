package snapshot

import (
	"context"
	"encoding/json"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/snapshot/job"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
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

var snapShotWorker = "snapshot_worker"

func (s *service) Name() string {
	return "snapshot"
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

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	snapshotNetworkNum := snapshotNetworkNumMap[message.Network]

	votes, err := s.getSnapshotVotes(ctx, message.Address, message.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("[snapshot worker] failed to get snapshot votes: %w", err)
	}
	if len(votes) == 0 {
		return nil, nil
	}

	proposalIDSet := mapset.NewSet()
	proposalIDs := []string{}
	spaceIDSet := mapset.NewSet()
	spaceIDs := []string{}

	for _, vote := range votes {
		proposalIDSet.Add(vote.ProposalID)
		spaceIDSet.Add(vote.SpaceID)
	}

	for _, proposalNode := range proposalIDSet.ToSlice() {
		proposal, ok := proposalNode.(model.SnapshotProposal)
		if !ok {
			logrus.Warnf("[snapshot worker] failed to convert proposal node to snapshot proposal:%v", proposalNode)
		}
		proposalIDs = append(proposalIDs, proposal.ID)
	}

	for _, spaceNode := range spaceIDSet.ToSlice() {
		space, ok := spaceNode.(model.SnapshotSpace)
		if !ok {
			logrus.Warnf("[snapshot worker] failed to convert space node to snapshot space:%v", spaceNode)
		}
		spaceIDs = append(spaceIDs, space.ID)
	}

	proposalMap := s.getSnapshotProposals(ctx, proposalIDs)

	spaceMap := s.getSnapshotSpaces(ctx, spaceIDs, snapshotNetworkNum)

	for _, vote := range votes {
		transactions = append(transactions, model.Transaction{
			Hash:        vote.ID,
			Timestamp:   vote.DateCreated,
			AddressFrom: message.Address,
			Network:     message.Network,
			Transfers: []model.Transfer{
				{
					TransactionHash:     vote.ID,
					Timestamp:           vote.DateCreated,
					TransactionLogIndex: protocol.LogIndexVirtual,
					AddressFrom:         message.Address,
					Metadata:            json.RawMessage{},
					Network:             message.Network,
				},
			},
		})
	}

	return transactions, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&job.SnapshotSpaceJob{
			SnapshotJobBase: job.SnapshotJobBase{
				Name:           "snapshot_space_job",
				DatabaseClient: s.databaseClient,
				RedisClient:    s.redisClient,
				SnapshotClient: s.snapshotClient,
				Limit:          1000,
			},
		},
		&job.SnapshotProposalJob{
			SnapshotJobBase: job.SnapshotJobBase{
				Name:           "snapshot_proposal_job",
				DatabaseClient: s.databaseClient,
				RedisClient:    s.redisClient,
				SnapshotClient: s.snapshotClient,
				Limit:          2000,
			},
		},
		&job.SnapshotVoteJob{
			SnapshotJobBase: job.SnapshotJobBase{
				Name:           "snapshot_vote_job",
				DatabaseClient: s.databaseClient,
				RedisClient:    s.redisClient,
				SnapshotClient: s.snapshotClient,
				Limit:          100000,
			},
		},
	}
}

func (s *service) getSnapshotVotes(ctx context.Context, address string, timestamp time.Time) ([]model.SnapshotVote, error) {
	_, handlerSpan := otel.Tracer(snapShotWorker).Start(ctx, "get_snapshot_vote_info")
	defer handlerSpan.End()

	var snapshotVotes []model.SnapshotVote

	// TODO:from redis

	// from db
	if err := s.databaseClient.
		Model(&model.SnapshotVote{}).
		Where("voter = ?", address).
		Find(&snapshotVotes).Error; err != nil {
		return nil, err
	}

	return snapshotVotes, nil
}

func (s *service) getSnapshotProposals(ctx context.Context, proposals []string) (map[string]model.SnapshotProposal, error) {
	_, handlerSpan := otel.Tracer(snapShotWorker).Start(ctx, "get_snapshot_proposal_info")
	defer handlerSpan.End()

	var snapshotProposals []model.SnapshotProposal
	var snapshotProposalMap = make(map[string]model.SnapshotProposal)

	// TODO:from redis

	// from db
	if err := s.databaseClient.
		Model(&model.SnapshotProposal{}).
		Where("proposal_id IN (?)", proposals).
		Find(&snapshotProposals).Error; err != nil {
		return nil, err
	}

	for _, proposal := range snapshotProposals {
		snapshotProposalMap[proposal.ID] = proposal
	}

	return snapshotProposalMap, nil
}

func (s *service) getSnapshotSpaces(ctx context.Context, spaces []string, networkNum string) (map[string]model.SnapshotSpace, error) {
	_, handlerSpan := otel.Tracer(snapShotWorker).Start(ctx, "get_snapshot_space_info")
	defer handlerSpan.End()

	var snapshotSpaces []model.SnapshotSpace
	var snapshotSpaceMap = make(map[string]model.SnapshotSpace)

	// TODO:from redis

	// from db
	if err := s.databaseClient.
		Model(&model.SnapshotSpace{}).
		Where("space in (?)", spaces).
		Where("network = ?", networkNum).
		Find(&snapshotSpaces).Error; err != nil {
		return nil, err
	}

	for _, space := range snapshotSpaces {
		snapshotSpaceMap[space.ID] = space
	}

	return snapshotSpaceMap, nil
}

// TODO
func (s *service) setSnapshotVotesInCache(ctx context.Context) {}

// TODO
func (s *service) setSnapshotProposalsInCache(ctx context.Context) {}

// TODO
func (s *service) setSnapshotSpacesInCache(ctx context.Context) {}

func New(
	databaseClient *gorm.DB,
	redisClient *redis.Client) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		redisClient:    redisClient,
		snapshotClient: snapshot.NewClient(),
	}
}
