package boardroom_test

import (
	"context"
	"github.com/naturalselectionlabs/pregod/common/worker/boardroom"
	"testing"
)

var (
	client *boardroom.Client
)

func init() {
	client = boardroom.NewClient()
}

func TestGetProposalsByRefId(t *testing.T) {
	refID := "cHJvcG9zYWw6Z2l0Y29pbjpzbmFwc2hvdDoweGVjYzk3MzZkYzRmOWU2ZWI5NTRiNjdhMGQ3MzAxZWNiY2M0OGJiZDVhZDMxZmJkMjhmMWU3NGFkZTUwMGM0Nzk="
	res, err := client.GetProposalsByRefId(context.Background(), refID)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetProposalsByAddress(t *testing.T) {
	//address := "0xDD659911EcBD4458db07Ee7cDdeC79bf8F859AbC"
	address := "0x5e349eca2dc61aBCd9dD99Ce94d04136151a09Ee"
	res, err := client.GetProposalsByAddress(context.Background(), address)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(res))
}

func TestGetVotesByAddress(t *testing.T) {
	address := "0x5e349eca2dc61aBCd9dD99Ce94d04136151a09Ee"
	res, err := client.GetVotesByAddress(context.Background(), address)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(res))
}
