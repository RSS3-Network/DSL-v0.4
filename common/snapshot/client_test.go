package snapshot_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/common/snapshot"
	"github.com/shurcooL/graphql"
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
		OrderDirection: snapshot.DescriptionAsc,
	}

	spaces, err := client.GetMultipleSpaces(context.Background(), variable)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(spaces)
}
