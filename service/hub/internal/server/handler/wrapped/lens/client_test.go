package lens

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
)

var (
	client = NewClient()
)

func TestGetStat(t *testing.T) {
	address := "0x827431510a5d249ce4fdb7f00c83a3353f471848"
	result := &model.SocialResult{}

	err := client.GetFollowStat(context.Background(), result, address)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(result.Follower > 0)
}
