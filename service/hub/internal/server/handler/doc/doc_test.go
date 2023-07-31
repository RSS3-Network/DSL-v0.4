package doc_test

import (
	"testing"

	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler/doc"
	"github.com/ysmood/got"
	"github.com/ysmood/gson"

	"github.com/getkin/kin-openapi/openapi3"
)

var docData = doc.New().Generate()

func TestValidateSchema(t *testing.T) {
	g := got.T(t)

	loader := &openapi3.Loader{Context: g.Context()}
	doc, _ := loader.LoadFromData(g.ToJSON(docData).Bytes())

	g.E(doc.Validate(g.Context()))
}

func TestNotes(t *testing.T) {
	g := got.T(t)

	out := g.ToJSONString(docData)

	g.Regex(`transaction.Bridge`, out)
	g.Regex(`metadata.Token`, out)
	g.Regex(`metadata.Swap`, out)
	g.Regex(`Farcaster`, out)
}

func TestBatchGetProfilesRequest(t *testing.T) {
	g := got.T(t)

	val := gson.NewFrom(g.ToJSONString(docData))

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

func TestTransferTypeMetadata(t *testing.T) {
	g := got.T(t)

	g.Eq(
		gson.NewFrom(g.ToJSONString(docData)).Get("components.schemas.TransferTypes.enum.0.4").Val(),
		map[string]interface{} /* len=4 */ {
			"actions": []interface{} /* len=1 cap=1 */ {
				map[string]interface{} /* len=3 */ {
					"comment": "交易所提币",
					"examples": []interface{} /* len=1 cap=1 */ {
						map[string]interface{} /* len=2 */ {
							"text": `Withdrew 123ETH on xxxx`,
							"hash": "0x1d226709361694160082822cb0a0542aa1a45d04e54fbd00453d8400c3673705",
						},
					},
					"platforms": []interface{} /* len=1 cap=1 */ {
						"CEX",
					},
				},
			},
			"metadata": `github.com/naturalselectionlabs/pregod/common/database/model/metadata.Token`, /* len=75 */
			"tag":      "exchange",
			"type":     "withdraw",
		},
	)
}
