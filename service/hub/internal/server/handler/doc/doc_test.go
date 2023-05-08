package doc_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler/doc"
	"github.com/ysmood/got"
	"github.com/ysmood/gson"
)

func TestNotes(t *testing.T) {
	g := got.T(t)

	d := doc.New()

	out := g.ToJSONString(d.Generate())

	g.Regex(`transaction.Bridge`, out)
	g.Regex(`model.Transfer`, out)
	g.Regex(`metadata.Token`, out)
	g.Regex(`metadata.Swap`, out)
	g.Regex(`Farcaster`, out)
}

func TestBatchGetProfilesRequest(t *testing.T) {
	g := got.T(t)

	d := doc.New()

	val := gson.NewFrom(g.ToJSONString(d.Generate()))

	g.Eq(
		val.Get("components.schemas.BatchGetProfilesRequest.required").Val(),
		[]interface{}{"address"},
	)
}
