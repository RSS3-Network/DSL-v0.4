package doc_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler/doc"
	"github.com/ysmood/got"
)

func TestNotes(t *testing.T) {
	g := got.T(t)

	d := doc.New()

	out := g.ToJSONString(d.Generate())

	g.WriteFile("tmp/out.json", out)

	g.Regex(`transaction.Bridge`, out)
	g.Regex(`model.Transfer`, out)
	g.Regex(`metadata.Token`, out)
	g.Regex(`metadata.Swap`, out)
}
