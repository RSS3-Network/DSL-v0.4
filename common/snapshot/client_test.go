package snapshot_test

import (
	"context"
	"testing"

	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/snapshot"
	graphqlx "github.com/naturalselectionlabs/pregod/common/snapshot/graphql"
	"github.com/stretchr/testify/assert"
)

var (
	client *snapshot.Client
)

func init() {
	client = snapshot.NewClient()
}

func TestGetMultipleSpaces(t *testing.T) {

	variable := snapshot.GetMultipleSpacesVariable{
		First:          graphql.Int(20),
		Skip:           graphql.Int(0),
		OrderBy:        graphql.String("created"),
		OrderDirection: snapshot.OrderDirectionAsc,
	}

	spaces, err := client.GetMultipleSpaces(context.Background(), variable)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, len(spaces) > 0)
}

func TestGetMultipleProposals(t *testing.T) {

	variable := snapshot.GetMultipleProposalsVariable{
		First:          graphql.Int(20),
		Skip:           graphql.Int(0),
		OrderBy:        graphql.String("created"),
		OrderDirection: snapshot.OrderDirectionAsc,
		Where: &graphqlx.ProposalWhere{
			Space_in: []graphql.String{"ens.eth"},
			State:    graphql.String("closed"),
		},
	}

	proposals, err := client.GetMultipleProposals(context.Background(), variable)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, len(proposals) > 0)
	// t.Log(proposals)
}

func TestGetMultipleVotes(t *testing.T) {
	variable := snapshot.GetMultipleVotesVariable{
		First:          graphql.Int(1000),
		Skip:           graphql.Int(0),
		OrderBy:        graphql.String("created"),
		OrderDirection: snapshot.OrderDirectionAsc,
		Where: &graphqlx.VoteWhere{
			Proposal: graphql.String("QmPvbwguLfcVryzBRrbY4Pb9bCtxURagdv1XjhtFLf3wHj"),
		},
	}

	proposals, err := client.GetMultipleVotes(context.Background(), variable)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(proposals)
}
