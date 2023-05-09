package doc_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler/doc"
	"github.com/ysmood/got"
	"github.com/ysmood/gson"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestValidateSchema(t *testing.T) {
	g := got.T(t)

	loader := &openapi3.Loader{Context: g.Context()}
	doc, _ := loader.LoadFromData(g.ToJSON(doc.New().Generate()).Bytes())

	g.E(doc.Validate(g.Context()))
}

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

	val := gson.NewFrom(g.ToJSONString(doc.New().Generate()))

	g.Eq(
		val.Get("components.schemas.BatchGetProfilesRequest.required").Val(),
		[]interface{}{"address"},
	)
}

func TestStableSchemaOutput(t *testing.T) {
	g := got.T(t)

	g.Eq(
		g.ToJSONString(doc.New().Generate()),
		g.ToJSONString(doc.New().Generate()),
	)
}
