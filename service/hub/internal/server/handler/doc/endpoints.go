package doc

import (
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/NaturalSelectionLabs/jschema"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
)

type Endpoint struct {
	Method string
	Path   string
	Params interface{}
	Result interface{}
}

var endpoints = []Endpoint{
	{http.MethodGet, handler.PathGetNotes, model.GetRequest{}, []dbModel.Transaction{}},
	{http.MethodGet, handler.PathGetAssets, model.GetAssetRequest{}, []dbModel.Asset{}},
	{http.MethodGet, handler.PathGetExchanges, model.GetExchangeRequest{}, []ExchangeResult{}},
	{http.MethodGet, handler.PathGetPlatformList, model.GetPlatformRequest{}, []model.PlatformResult{}},
	{http.MethodGet, handler.PathGetProfiles, model.GetRequest{}, []social.Profile{}},
	{http.MethodGet, handler.PathGetNameResolve, model.GetNameResolveRequest{}, dbModel.NameServiceResult{}},
	{http.MethodGet, handler.PathGetTransaction, model.GetTransactionRequest{}, dbModel.Transaction{}},
	{http.MethodGet, handler.PathGetMastodon, model.GetRequest{}, []dbModel.Transaction{}},
	{http.MethodGet, handler.PathGetNotesByPlatform, model.GetNotesByPlatformRequest{}, []dbModel.Transaction{}},

	{http.MethodPost, handler.PathBatchGetSocialNotes, model.BatchGetSocialNotesRequest{}, []dbModel.Transaction{}},
	{http.MethodPost, handler.PathBatchGetNotes, model.BatchGetNotesRequest{}, []dbModel.Transaction{}},
	{http.MethodPost, handler.PathBatchGetProfiles, model.BatchGetProfilesRequest{}, []social.Profile{}},

	{http.MethodPost, handler.PathPostAPIKey, model.PostAPIKeyRequest{}, dbModel.APIKey{}},
	{http.MethodGet, handler.PathGetAPIKey, model.GetAPIKeyRequest{}, dbModel.APIKey{}},
}

func (d *Doc) endpoints() Obj {
	paths := Obj{}

	type schema struct {
		Params    []Obj `json:"parameters,omitempty"`
		ReqBody   Obj   `json:"requestBody,omitempty"`
		Responses Obj   `json:"responses,omitempty"`
	}

	for _, e := range endpoints {
		def := schema{
			Params:    d.params(e),
			ReqBody:   d.reqBody(e),
			Responses: d.resBody(e),
		}
		paths[toURITemplate(e.Path)] = Obj{strings.ToLower(e.Method): def}
	}

	return paths
}

func (d *Doc) params(e Endpoint) []Obj {
	if e.Method != http.MethodGet {
		return nil
	}

	t := reflect.TypeOf(e.Params)

	df := []Obj{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		name := f.Tag.Get("param")
		in := "path"
		query := f.Tag.Get("query")

		if query != "" {
			name = query
			in = "query"
		}

		p := Obj{
			"name":        name,
			"in":          in,
			"schema":      d.schemas.DefineT(f.Type),
			"description": f.Tag.Get("description"),
		}

		if echoRequired(f.Tag) {
			p["required"] = true
		}

		df = append(df, p)
	}

	return df
}

func (d *Doc) reqBody(e Endpoint) Obj {
	if e.Method != http.MethodPost {
		return nil
	}

	d.schemas.Define(e.Params)
	required := []string{}

	t := reflect.TypeOf(e.Params)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if echoRequired(f.Tag) {
			jt := jschema.ParseJSONTag(f.Tag)
			required = append(required, jt.Name)
		}
	}

	d.schemas.PeakSchema(e.Params).Required = required

	return Obj{
		"content": Obj{
			"application/json": Obj{
				"schema": d.schemas.PeakSchema(e.Params),
			},
		},
		"required": true,
	}
}

func (d *Doc) resBody(e Endpoint) Obj {
	res := d.schemas.Define(e.Result)

	if reflect.TypeOf(e.Result).Kind() == reflect.Slice {
		res = d.schemas.PeakSchema(model.Response{}).Clone()
		res.Properties["result"] = d.schemas.Define(e.Result)
	}

	return Obj{
		"200": Obj{
			"description": "Response",
			"content": Obj{
				"application/json": Obj{
					"schema": res,
				},
			},
		},
	}
}

var uriReg = regexp.MustCompile(`\:([a-z_]+)`)

// Convert to URI template format
func toURITemplate(p string) string {
	return uriReg.ReplaceAllString(p, "{$1}")
}

func echoRequired(tag reflect.StructTag) bool {
	v := tag.Get("validate")
	return strings.Contains(v, "required") && !strings.Contains(v, "required_with")
}
